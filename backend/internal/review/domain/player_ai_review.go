package review

import "context"

type PlayerCommentGenerator interface {
	GeneratePlayerComment(ctx context.Context, stats PlayerMatchStats) (string, error)
}

type PlayerAIReview struct {
	PlayerID   int
	PlayerName string
	Comment    string
}

type PlayerReviewService struct {
	generator PlayerCommentGenerator
}

func (s PlayerReviewService) Generate(ctx context.Context, stats PlayerMatchStats) (PlayerAIReview, error) {
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
