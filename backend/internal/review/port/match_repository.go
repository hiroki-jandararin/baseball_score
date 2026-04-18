package port

import (
	"context"

	review "baseball-score-app/backend/internal/review/domain"
)

type MatchRepository interface {
	SaveMatch(ctx context.Context, match review.Match) (int, error)
}
