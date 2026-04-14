package review

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type fakePlayerCommentGenerator struct {
	comment  string
	err      error
	called   bool
	gotStats PlayerMatchStats
}

func (f *fakePlayerCommentGenerator) GeneratePlayerComment(ctx context.Context, stats PlayerMatchStats) (string, error) {
	f.called = true
	f.gotStats = stats
	if f.err != nil {
		return "", f.err
	}
	return f.comment, nil
}

func TestPlayerReviewServiceGenerateReturnsGeneratedComment(t *testing.T) {
	t.Parallel()

	fake := &fakePlayerCommentGenerator{
		comment: "今日はバットが火を吹いた。完全に主役。",
	}

	service := PlayerReviewService{
		generator: fake,
	}

	stats := PlayerMatchStats{
		PlayerID:   1,
		PlayerName: "山田",
		Hits:       3,
		RBI:        2,
		Runs:       1,
	}

	got, err := service.Generate(context.Background(), stats)
	require.NoError(t, err)

	assert.True(t, fake.called)
	assert.Equal(t, stats, fake.gotStats)
	assert.Equal(t, PlayerAIReview{
		PlayerID:   1,
		PlayerName: "山田",
		Comment:    "今日はバットが火を吹いた。完全に主役。",
	}, got)
}
