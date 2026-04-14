package openai

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	platformconfig "baseball-score-app/backend/internal/platform/config"
	"baseball-score-app/backend/internal/review/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewClientRejectsMissingAPIKey(t *testing.T) {
	_, err := NewClient(platformconfig.OpenAI{
		BaseURL: "https://api.openai.com/v1",
		Model:   "gpt-4.1-mini",
		Timeout: time.Second,
	})

	require.Error(t, err)
	assert.Contains(t, err.Error(), "api key")
}

func TestGenerateMatchReviewSendsExpectedRequest(t *testing.T) {
	t.Parallel()

	var authHeader string
	var projectHeader string
	var requestBody map[string]any

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		require.Equal(t, "/responses", r.URL.Path)

		authHeader = r.Header.Get("Authorization")
		projectHeader = r.Header.Get("OpenAI-Project")
		require.NoError(t, json.NewDecoder(r.Body).Decode(&requestBody))

		writeResponsesSuccess(t, w, `{"headline":"逆転勝利"}`)
	}))
	defer server.Close()

	client := newTestClient(t, server.URL, "test-key", "proj_123")

	_, err := client.GenerateMatchReview(context.Background(), GenerateReviewParams{
		InputJSON: `{"match":{"id":1}}`,
	})
	require.NoError(t, err)

	assert.Equal(t, "Bearer test-key", authHeader)
	assert.Equal(t, "proj_123", projectHeader)
	assert.Equal(t, "gpt-4.1-mini", requestBody["model"])
	assert.Equal(t, "Match facts JSON:\n{\"match\":{\"id\":1}}", requestBody["input"])

	textConfig, ok := requestBody["text"].(map[string]any)
	require.True(t, ok)
	formatConfig, ok := textConfig["format"].(map[string]any)
	require.True(t, ok)
	assert.Equal(t, "json_object", formatConfig["type"])
}

func TestGenerateMatchReviewReturnsOutputText(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		writeResponsesSuccess(t, w, `{"headline":"逆転勝利"}`)
	}))
	defer server.Close()

	client := newTestClient(t, server.URL, "test-key", "")

	result, err := client.GenerateMatchReview(context.Background(), GenerateReviewParams{
		InputJSON: `{"match":{"id":1}}`,
	})
	require.NoError(t, err)

	assert.Equal(t, `{"headline":"逆転勝利"}`, result.RawResponse)
}

func TestGenerateMatchReviewReturnsAPIError(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte(`{"error":{"message":"bad api key","type":"invalid_request_error"}}`))
	}))
	defer server.Close()

	client, err := NewClient(platformconfig.OpenAI{
		APIKey:  "bad-key",
		BaseURL: server.URL,
		Model:   "gpt-4.1-mini",
		Timeout: 2 * time.Second,
	})
	require.NoError(t, err)

	_, err = client.GenerateMatchReview(context.Background(), GenerateReviewParams{
		InputJSON: `{"match":{"id":1}}`,
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "bad api key")
}

func TestGeneratePlayerCommentSendsExpectedRequest(t *testing.T) {
	t.Parallel()

	var requestBody map[string]any

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		require.Equal(t, "/responses", r.URL.Path)
		require.NoError(t, json.NewDecoder(r.Body).Decode(&requestBody))

		writeResponsesSuccess(t, w, "今日はバットが火を吹いた。完全に主役。")
	}))
	defer server.Close()

	client := newTestClient(t, server.URL, "test-key", "")

	_, err := client.GeneratePlayerComment(context.Background(), review.PlayerMatchStats{
		PlayerID:   1,
		PlayerName: "山田",
		Hits:       3,
		RBI:        2,
		Runs:       1,
	})
	require.NoError(t, err)

	assert.Equal(t, "gpt-4.1-mini", requestBody["model"])
	assert.Equal(t, "Player result JSON:\n{\"player_id\":1,\"player_name\":\"山田\",\"hits\":3,\"rbi\":2,\"runs\":1}", requestBody["input"])
}

func TestGeneratePlayerCommentReturnsOutputText(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		writeResponsesSuccess(t, w, "今日はバットが火を吹いた。完全に主役。")
	}))
	defer server.Close()

	client := newTestClient(t, server.URL, "test-key", "")

	got, err := client.GeneratePlayerComment(context.Background(), review.PlayerMatchStats{
		PlayerID:   1,
		PlayerName: "山田",
		Hits:       3,
		RBI:        2,
		Runs:       1,
	})
	require.NoError(t, err)
	assert.Equal(t, "今日はバットが火を吹いた。完全に主役。", got)
}

func newTestClient(t *testing.T, baseURL, apiKey, project string) *Client {
	t.Helper()

	client, err := NewClient(platformconfig.OpenAI{
		APIKey:  apiKey,
		BaseURL: baseURL,
		Model:   "gpt-4.1-mini",
		Project: project,
		Timeout: 2 * time.Second,
	})
	require.NoError(t, err)

	return client
}

func writeResponsesSuccess(t *testing.T, w http.ResponseWriter, outputText string) {
	t.Helper()

	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write([]byte(`{
		"id":"resp_123",
		"object":"response",
		"created_at":1740000000,
		"model":"gpt-4.1-mini",
		"status":"completed",
		"output":[
			{
				"id":"msg_123",
				"type":"message",
				"role":"assistant",
				"status":"completed",
				"content":[
					{
						"type":"output_text",
						"text":` + strconv.Quote(outputText) + `,
						"annotations":[]
					}
				]
			}
		]
	}`))
	require.NoError(t, err)
}
