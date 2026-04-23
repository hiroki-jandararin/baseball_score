package usecase

import (
	"context"
	"testing"
	"time"

	review "baseball-score-app/backend/internal/review/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type fakeNewsMatchRepository struct {
	match      review.Match
	err        error
	called     bool
	gotMatchID int
}

func (f *fakeNewsMatchRepository) SaveMatch(ctx context.Context, match review.Match) (int, error) {
	return 0, nil
}

func (f *fakeNewsMatchRepository) FindMatchByID(ctx context.Context, matchID int) (review.Match, error) {
	f.called = true
	f.gotMatchID = matchID
	if f.err != nil {
		return review.Match{}, f.err
	}
	return f.match, nil
}

func TestNewNewsServiceUsesRepository(t *testing.T) {
	t.Parallel()

	const matchID = 10

	match := review.Match{
		ID:            matchID,
		OpponentName:  "東町ライバルズ",
		MatchDate:     time.Date(2026, 4, 18, 10, 0, 0, 0, time.UTC),
		IsWin:         1,
		TeamScore:     5,
		OpponentScore: 3,
		PlayerStats: []review.PlayerMatchStats{
			{PlayerID: 2, PlayerName: "佐藤", Hits: 2, RBI: 2},
		},
	}
	repo := &fakeNewsMatchRepository{match: match}

	service := NewNewsService(repo)
	got, err := service.GenerateArticle(context.Background(), matchID)
	require.NoError(t, err)

	assert.True(t, repo.called)
	assert.Equal(t, matchID, repo.gotMatchID)
	assert.Equal(t, "青空シャークスが東町ライバルズに5-3で勝利", got.Title)
}

func TestNewsServiceGenerateArticleBuildsArticle(t *testing.T) {
	t.Parallel()

	const matchID = 10

	match := review.Match{
		ID:            matchID,
		OpponentName:  "東町ライバルズ",
		MatchDate:     time.Date(2026, 4, 18, 10, 0, 0, 0, time.UTC),
		IsWin:         1,
		TeamScore:     5,
		OpponentScore: 3,
		PlayerStats: []review.PlayerMatchStats{
			{
				PlayerID:     1,
				PlayerName:   "山田",
				BattingOrder: 1,
				Position:     "CF",
				Hits:         1,
				Runs:         1,
			},
			{
				PlayerID:     2,
				PlayerName:   "佐藤",
				BattingOrder: 4,
				Position:     "1B",
				Hits:         2,
				RBI:          2,
			},
		},
	}

	repo := &fakeNewsMatchRepository{match: match}
	service := NewsService{repo: repo}

	got, err := service.GenerateArticle(context.Background(), matchID)
	require.NoError(t, err)

	assert.True(t, repo.called)
	assert.Equal(t, matchID, repo.gotMatchID)
	assert.Equal(t, NewsArticle{
		Title:   "青空シャークスが東町ライバルズに5-3で勝利",
		Source:  "Team Sports",
		Time:    "2026.04.18 10:00",
		Summary: "青空シャークスは東町ライバルズとの試合を5-3で制し、佐藤の活躍が光った。",
		Lead:    "1番山田が攻撃の起点を作り、MVPの佐藤が勝負どころで流れを引き寄せた。",
		Paragraphs: []string{
			"青空シャークスは東町ライバルズを相手に5-3で勝利した。",
			"佐藤は2安打2打点の内容で、チームの勝利に大きく貢献した。",
			"山田も出塁と得点で攻撃を支え、打線全体に流れを作った。",
		},
	}, got)
}
