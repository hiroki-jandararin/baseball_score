package usecase

import (
	"context"
	"testing"
	"time"

	review "baseball-score-app/backend/internal/review/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type fakeMatchRepository struct {
	savedMatch review.Match
	returned   int
	err        error
	called     bool
}

func (f *fakeMatchRepository) SaveMatch(ctx context.Context, match review.Match) (int, error) {
	f.called = true
	f.savedMatch = match
	if f.err != nil {
		return 0, f.err
	}
	return f.returned, nil
}

func (f *fakeMatchRepository) FindMatchByID(ctx context.Context, matchID int) (review.Match, error) {
	return review.Match{}, nil
}

func TestSaveMatchResultServiceSaveMatchResult(t *testing.T) {
	t.Parallel()

	input := review.Match{
		TeamID:        1,
		OpponentName:  "Rivals",
		MatchDate:     time.Date(2026, 4, 18, 10, 0, 0, 0, time.UTC),
		Location:      "Tokyo",
		IsWin:         1,
		TeamScore:     5,
		OpponentScore: 3,
		Note:          "good game",
		PlayerStats: []review.PlayerMatchStats{
			{
				PlayerID:   10,
				PlayerName: "山田",
				Hits:       2,
				RBI:        1,
			},
			{
				PlayerID:   20,
				PlayerName: "佐藤",
				Hits:       1,
				Runs:       2,
			},
		},
	}

	expectedID := 99

	repo := &fakeMatchRepository{returned: expectedID}
	service := SaveMatchResultService{repo: repo}

	got, err := service.SaveMatchResult(context.Background(), input)
	require.NoError(t, err)

	assert.True(t, repo.called)
	assert.Equal(t, expectedID, got)
	assert.Equal(t, input, repo.savedMatch)
}

func TestSaveMatchResultServiceSaveMatchResultReturnsError(t *testing.T) {
	t.Parallel()

	repo := &fakeMatchRepository{err: assert.AnError}
	service := SaveMatchResultService{repo: repo}

	got, err := service.SaveMatchResult(context.Background(), review.Match{})
	require.Error(t, err)
	assert.ErrorIs(t, err, assert.AnError)
	assert.Equal(t, 0, got)
}
