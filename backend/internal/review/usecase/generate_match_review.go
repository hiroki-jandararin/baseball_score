package usecase

import (
	review "baseball-score-app/backend/internal/review/domain"
	"context"
)

type PlayerReviewGenerator interface {
	GeneratePlayerReview(ctx context.Context, stats review.PlayerMatchStats, isMVP bool) (review.PlayerReview, error)
}

type MatchSummaryGenerator interface {
	GenerateMatchSummary(ctx context.Context, match review.Match, mvp review.MVPResult, players []review.PlayerReview) (review.MatchSummary, error)
}

type MatchReviewService struct {
	playerReviewGenerator PlayerReviewGenerator
	matchSummaryGenerator MatchSummaryGenerator
}

func (s MatchReviewService) GenerateMatchReview(ctx context.Context, match review.Match) (review.MatchReview, error) {
	mvp := review.SelectMVP(match)

	var playerReviews []review.PlayerReview
	for _, stats := range match.PlayerStats {
		isMVP := stats.PlayerID == mvp.PlayerID
		playerReview, err := s.playerReviewGenerator.GeneratePlayerReview(ctx, stats, isMVP)
		if err != nil {
			return review.MatchReview{}, err
		}
		playerReviews = append(playerReviews, playerReview)
	}

	matchSummary, err := s.matchSummaryGenerator.GenerateMatchSummary(ctx, match, mvp, playerReviews)
	if err != nil {
		return review.MatchReview{}, err
	}

	return review.MatchReview{
  		Match: review.MatchOverview{
  			ID:            match.ID,
  			MatchDate:     match.MatchDate,
  			OpponentName:  match.OpponentName,
  			TeamScore:     match.TeamScore,
  			OpponentScore: match.OpponentScore,
  			IsWin:         match.IsWin,
  		},
  		MVP:     mvp,
  		Summary: matchSummary,
  		Players: playerReviews,
  	}, nil

}


