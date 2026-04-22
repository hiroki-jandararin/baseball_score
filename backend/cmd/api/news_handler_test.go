package main

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"baseball-score-app/backend/internal/review/usecase"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type fakeNewsService struct {
	article    usecase.NewsArticle
	err        error
	called     bool
	gotMatchID int
}

func (f *fakeNewsService) GenerateArticle(ctx context.Context, matchID int) (usecase.NewsArticle, error) {
	f.called = true
	f.gotMatchID = matchID
	if f.err != nil {
		return usecase.NewsArticle{}, f.err
	}
	return f.article, nil
}

func TestNewsHandlerReturnsArticle(t *testing.T) {
	t.Parallel()

	expected := usecase.NewsArticle{
		Title:   "SharksがRivalsに5-3で勝利",
		Source:  "Team Sports",
		Time:    "2026.04.18 10:00",
		Summary: "SharksはRivalsとの試合を5-3で制し、佐藤の活躍が光った。",
		Lead:    "1番山田が攻撃の起点を作り、MVPの佐藤が勝負どころで流れを引き寄せた。",
		Paragraphs: []string{
			"SharksはRivalsを相手に5-3で勝利した。",
			"佐藤は2安打2打点の内容で、チームの勝利に大きく貢献した。",
			"山田も出塁と得点で攻撃を支え、打線全体に流れを作った。",
		},
	}
	service := &fakeNewsService{article: expected}
	handler := newNewsHandler(service)

	request := httptest.NewRequest(http.MethodGet, "/matches/10/news", nil)
	recorder := httptest.NewRecorder()

	handler.ServeHTTP(recorder, request)

	require.Equal(t, http.StatusOK, recorder.Code)
	assert.True(t, service.called)
	assert.Equal(t, 10, service.gotMatchID)
	assert.Contains(t, recorder.Header().Get("Content-Type"), "application/json")

	var got usecase.NewsArticle
	err := json.NewDecoder(recorder.Body).Decode(&got)
	require.NoError(t, err)
	assert.Equal(t, expected, got)
}
