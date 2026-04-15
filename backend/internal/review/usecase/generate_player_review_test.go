package usecase

import (
	"context"
	"testing"

	review "baseball-score-app/backend/internal/review/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)


type fakePlayerCommentGenerator struct {
	comment  string
	err      error
	called   bool
	gotStats review.PlayerMatchStats
}

func (f *fakePlayerCommentGenerator) GeneratePlayerComment(ctx context.Context, stats review.PlayerMatchStats) (string, error) {
	f.called = true
	f.gotStats = stats
	if f.err != nil {
		return "", f.err
	}
	return f.comment, nil
}

func TestPlayerReviewServiceGenerateBuildsPlayerReview(t *testing.T) {
	t.Parallel()

	fake := &fakePlayerCommentGenerator{
		comment: "今日はバットが火を吹いた。完全に主役。",
	}

	service := PlayerReviewService{
		generator: fake,
	}

	stats := review.PlayerMatchStats{
		PlayerID:   1,
		PlayerName: "山田",
		Hits:       3,
		RBI:        2,
		Runs:       1,
	}

	got, err := service.Generate(context.Background(), stats, false)
	require.NoError(t, err)

	assert.True(t, fake.called)
	assert.Equal(t, stats, fake.gotStats)
	assert.Equal(t, review.PlayerReview{
		PlayerID:   1,
		PlayerName: "山田",
		Stats:      stats,
		Title:      review.TitleHitKing,
		Comment:    "今日はバットが火を吹いた。完全に主役。",
		IsMVP:      false,
	}, got)
}

func TestPlayerReviewServiceGenerateMarksMVP(t *testing.T) {
	t.Parallel()

	fake := &fakePlayerCommentGenerator{
		comment: "勝負どころで仕事した。本日の主役。",
	}

	service := PlayerReviewService{
		generator: fake,
	}

	stats := review.PlayerMatchStats{
		PlayerID:   7,
		PlayerName: "佐藤",
		Hits:       2,
		RBI:        2,
	}

	got, err := service.Generate(context.Background(), stats, true)
	require.NoError(t, err)

	assert.Equal(t, review.PlayerReview{
		PlayerID:   7,
		PlayerName: "佐藤",
		Stats:      stats,
		Title:      review.TitleMainCharacter,
		Comment:    "勝負どころで仕事した。本日の主役。",
		IsMVP:      true,
	}, got)
}

 func TestPlayerReviewServiceGenerateReturnsErrorWhenCommentGenerationFails(t *testing.T) {
  	t.Parallel()

  	fake := &fakePlayerCommentGenerator{
  		err: assert.AnError,
  	}

  	service := PlayerReviewService{
  		generator: fake,
  	}

  	stats := review.PlayerMatchStats{
  		PlayerID:   1,
  		PlayerName: "山田",
  	}

  	got, err := service.Generate(context.Background(), stats, false)
  	require.Error(t, err)
  	assert.ErrorIs(t, err, assert.AnError)
  	assert.Equal(t, review.PlayerReview{}, got)
  }