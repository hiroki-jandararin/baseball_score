package usecase

import (
	"context"

	review "baseball-score-app/backend/internal/review/domain"
	"baseball-score-app/backend/internal/review/port"
)

type SaveMatchResultService struct {
	repo port.MatchRepository
}

func (s SaveMatchResultService) SaveMatchResult(ctx context.Context, match review.Match) (int, error) {
	return s.repo.SaveMatch(ctx, match)
}
