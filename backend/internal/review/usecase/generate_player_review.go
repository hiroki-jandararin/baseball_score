package usecase

import (
	review "baseball-score-app/backend/internal/review/domain"
	"baseball-score-app/backend/internal/review/port"
	"context"
)

type PlayerAIReview struct {
	PlayerID   int
	PlayerName string
	Comment    string
}

type PlayerReviewService struct {
	generator port.PlayerCommentGenerator
}

func (s PlayerReviewService) Generate(ctx context.Context, stats review.PlayerMatchStats) (PlayerAIReview, error) {
	comment, err := s.generator.GeneratePlayerComment(ctx, stats)
	if err != nil {
		return PlayerAIReview{}, err
	}

	return PlayerAIReview{
		PlayerID:   stats.PlayerID,
		PlayerName: stats.PlayerName,
		Comment:    comment,
	}, nil
}