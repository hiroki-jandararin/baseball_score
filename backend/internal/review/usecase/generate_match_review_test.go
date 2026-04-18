package usecase

import (
	review "baseball-score-app/backend/internal/review/domain"
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type fakeMatchSummaryGenerator struct {
	summary review.MatchSummary
	err     error
	called  bool
	gotArgs struct {
		match   review.Match
		mvp     review.MVPResult
		players []review.PlayerReview
	}
}

type playerReviewCall struct {
	stats review.PlayerMatchStats
	isMVP bool
}

type fakePlayerReviewGenerator struct {
	reviews map[int]review.PlayerReview
	err     error
	called  bool
	calls   []playerReviewCall
}

func (f *fakePlayerReviewGenerator) GeneratePlayerReview(ctx context.Context, stats review.PlayerMatchStats, isMVP bool) (review.PlayerReview, error) {
	f.called = true
	f.calls = append(f.calls, playerReviewCall{
		stats: stats,
		isMVP: isMVP,
	})
	if f.err != nil {
		return review.PlayerReview{}, f.err
	}
	playerReview, ok := f.reviews[stats.PlayerID]
	if !ok {
		return playerReview, fmt.Errorf("no review for player id %d", stats.PlayerID)
	}
	return playerReview, nil
}

func (f *fakeMatchSummaryGenerator) GenerateMatchSummary(ctx context.Context, match review.Match, mvp review.MVPResult, players []review.PlayerReview) (review.MatchSummary, error) {
	f.called = true
	f.gotArgs.match = match
	f.gotArgs.mvp = mvp
	f.gotArgs.players = players
	if f.err != nil {
		return review.MatchSummary{}, f.err
	}
	return f.summary, nil
}

func TestMatchReviewServiceGenerateMatchReviewBuildsReview(t *testing.T) {
	t.Parallel()

	playerGenerator := &fakePlayerReviewGenerator{
		reviews: map[int]review.PlayerReview{
			1: {
				PlayerID:   1,
				PlayerName: "山田",
				Title:      "仕事人",
				Comment:    "堅実に仕事をした。",
				IsMVP:      false,
			},
			2: {
				PlayerID:   2,
				PlayerName: "佐藤",
				Title:      "本日の主役",
				Comment:    "勝負どころで決めた。",
				IsMVP:      true,
			},
		},
	}

	summaryGenerator := &fakeMatchSummaryGenerator{
		summary: review.MatchSummary{
			Headline: "終盤に決めた快勝",
			Summary:  "主役の一打で流れを引き寄せた。",
		},
	}

	service := MatchReviewService{
		playerReviewGenerator: playerGenerator,
		matchSummaryGenerator: summaryGenerator,
	}

	match := review.Match{
		ID:            10,
		MatchDate:     time.Date(2026, 4, 17, 0, 0, 0, 0, time.UTC),
		OpponentName:  "Rivals",
		TeamScore:     5,
		OpponentScore: 3,
		IsWin:         1,
		PlayerStats: []review.PlayerMatchStats{
			{PlayerID: 1, PlayerName: "山田", Hits: 1},
			{PlayerID: 2, PlayerName: "佐藤", Hits: 2, RBI: 1},
		},
	}

	got, err := service.GenerateMatchReview(context.Background(), match)
	require.NoError(t, err)

	expectedPlayers := []review.PlayerReview{
		{
			PlayerID:   1,
			PlayerName: "山田",
			Title:      "仕事人",
			Comment:    "堅実に仕事をした。",
			IsMVP:      false,
		},
		{
			PlayerID:   2,
			PlayerName: "佐藤",
			Title:      "本日の主役",
			Comment:    "勝負どころで決めた。",
			IsMVP:      true,
		},
	}

	assert.Equal(t, review.MatchReview{
		Match: review.MatchOverview{
			ID:            10,
			MatchDate:     match.MatchDate,
			OpponentName:  "Rivals",
			TeamScore:     5,
			OpponentScore: 3,
			IsWin:         1,
		},
		MVP: review.MVPResult{
			PlayerID:   2,
			PlayerName: "佐藤",
			Score:      7,
		},
		Summary: review.MatchSummary{
			Headline: "終盤に決めた快勝",
			Summary:  "主役の一打で流れを引き寄せた。",
		},
		Players: expectedPlayers,
	}, got)

	assert.True(t, playerGenerator.called)
	assert.Equal(t, []playerReviewCall{
		{
			stats: review.PlayerMatchStats{PlayerID: 1, PlayerName: "山田", Hits: 1},
			isMVP: false,
		},
		{
			stats: review.PlayerMatchStats{PlayerID: 2, PlayerName: "佐藤", Hits: 2, RBI: 1},
			isMVP: true,
		},
	}, playerGenerator.calls)

	assert.True(t, summaryGenerator.called)
	assert.Equal(t, match, summaryGenerator.gotArgs.match)
	assert.Equal(t, review.MVPResult{
		PlayerID:   2,
		PlayerName: "佐藤",
		Score:      7,
	}, summaryGenerator.gotArgs.mvp)
	assert.Equal(t, expectedPlayers, summaryGenerator.gotArgs.players)
}

func TestMatchReviewServiceGenerateMatchReviewPlayerReviewError(t *testing.T) {
	t.Parallel()

	playerGenerator := &fakePlayerReviewGenerator{
		err: assert.AnError,
	}

	summaryGenerator := &fakeMatchSummaryGenerator{}

	service := MatchReviewService{
		playerReviewGenerator: playerGenerator,
		matchSummaryGenerator: summaryGenerator,
	}

	match := review.Match{
		PlayerStats: []review.PlayerMatchStats{
			{PlayerID: 1, PlayerName: "山田"},
		},
	}

	got, err := service.GenerateMatchReview(context.Background(), match)
	require.Error(t, err)
	assert.ErrorIs(t, err, assert.AnError)
	assert.Equal(t, review.MatchReview{}, got)
	assert.False(t, summaryGenerator.called)
}

func TestMatchReviewServiceGenerateMatchReviewSummaryError(t *testing.T) {
	t.Parallel()

	playerGenerator := &fakePlayerReviewGenerator{
		reviews: map[int]review.PlayerReview{
			2: {
				PlayerID:   2,
				PlayerName: "佐藤",
				Title:      "本日の主役",
				Comment:    "勝負どころで決めた。",
				IsMVP:      true,
			},
		},
	}

	summaryGenerator := &fakeMatchSummaryGenerator{
		err: assert.AnError,
	}

	service := MatchReviewService{
		playerReviewGenerator: playerGenerator,
		matchSummaryGenerator: summaryGenerator,
	}

	match := review.Match{
		PlayerStats: []review.PlayerMatchStats{
			{PlayerID: 2, PlayerName: "佐藤", Hits: 2, RBI: 1},
		},
	}

	got, err := service.GenerateMatchReview(context.Background(), match)
	require.Error(t, err)
	assert.ErrorIs(t, err, assert.AnError)
	assert.Equal(t, review.MatchReview{}, got)
	assert.True(t, playerGenerator.called)
	assert.True(t, summaryGenerator.called)
}
