package port

import (
	review "baseball-score-app/backend/internal/review/domain"
	"context"
)

type PlayerCommentGenerator interface {
	GeneratePlayerComment(ctx context.Context, stats review.PlayerMatchStats) (string, error)
}