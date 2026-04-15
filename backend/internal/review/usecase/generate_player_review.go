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

func (s PlayerReviewService) Generate(ctx context.Context, stats review.PlayerMatchStats, isMVP bool) (review.PlayerReview, error) {
	comment, err := s.generator.GeneratePlayerComment(ctx, stats)
	if err != nil {
		return review.PlayerReview{}, err
	}

	title := review.AssignTitle(stats)

	return review.PlayerReview{
		PlayerID:   stats.PlayerID,
		PlayerName: stats.PlayerName,
		Stats:      stats,
		Title:      title,
		Comment:    comment,
		IsMVP:      isMVP,
	}, nil
}
