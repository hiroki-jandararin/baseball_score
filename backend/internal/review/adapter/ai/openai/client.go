package openai

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	platformconfig "baseball-score-app/backend/internal/platform/config"

	openai "github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
	"github.com/openai/openai-go/v3/responses"
	"github.com/openai/openai-go/v3/shared"
)

type Client struct {
	client openai.Client
	model  string
}

type GenerateReviewParams struct {
	InputJSON string
}

type GeneratedReview struct {
	RawResponse string
}


func NewClient(cfg platformconfig.OpenAI) (*Client, error) {
	if strings.TrimSpace(cfg.APIKey) == "" {
		return nil, fmt.Errorf("openai api key is required")
	}

	baseURL := strings.TrimRight(strings.TrimSpace(cfg.BaseURL), "/")
	if baseURL == "" {
		return nil, fmt.Errorf("openai base url is required")
	}

	timeout := cfg.Timeout
	if timeout <= 0 {
		timeout = 30 * time.Second
	}

	model := strings.TrimSpace(cfg.Model)
	if model == "" {
		model = "gpt-4.1-mini"
	}

	opts := []option.RequestOption{
		option.WithAPIKey(cfg.APIKey),
		option.WithBaseURL(baseURL),
		option.WithHTTPClient(&http.Client{Timeout: timeout}),
	}
	if cfg.Project != "" {
		opts = append(opts, option.WithProject(strings.TrimSpace(cfg.Project)))
	}

	return &Client{
		client: openai.NewClient(opts...),
		model:  model,
	}, nil
}

func (c *Client) GenerateMatchReview(ctx context.Context, params GenerateReviewParams) (GeneratedReview, error) {
	if strings.TrimSpace(params.InputJSON) == "" {
		return GeneratedReview{}, fmt.Errorf("input json is required")
	}

	resp, err := c.client.Responses.New(ctx, responses.ResponseNewParams{
		Model:        shared.ResponsesModel(c.model),
		Input:        responses.ResponseNewParamsInputUnion{OfString: openai.String(buildPrompt(params.InputJSON))},
		Text: responses.ResponseTextConfigParam{
			Format: responses.ResponseFormatTextConfigUnionParam{
				OfJSONObject: &shared.ResponseFormatJSONObjectParam{Type: "json_object"},
			},
		},
	})
	if err != nil {
		return GeneratedReview{}, fmt.Errorf("call openai responses api: %w", err)
	}

	outputText := strings.TrimSpace(resp.OutputText())
	if outputText == "" {
		return GeneratedReview{}, fmt.Errorf("openai response did not include output_text")
	}

	return GeneratedReview{RawResponse: outputText}, nil
}

func buildPrompt(inputJSON string) string {
	return "Match facts JSON:\n" + strings.TrimSpace(inputJSON)
}
