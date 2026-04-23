package usecase

import (
	"context"
	"fmt"

	review "baseball-score-app/backend/internal/review/domain"
	"baseball-score-app/backend/internal/review/port"
)

type NewsArticle struct {
	Title      string   `json:"title"`
	Source     string   `json:"source"`
	Time       string   `json:"time"`
	Summary    string   `json:"summary"`
	Lead       string   `json:"lead"`
	Paragraphs []string `json:"paragraphs"`
}

type NewsService struct {
	repo port.MatchRepository
}

func NewNewsService(repo port.MatchRepository) *NewsService {
	return &NewsService{repo: repo}
}

func (s *NewsService) GenerateArticle(ctx context.Context, matchID int) (NewsArticle, error) {
	match, err := s.repo.FindMatchByID(ctx, matchID)
	if err != nil {
		return NewsArticle{}, err
	}

	mvp := review.SelectMVP(match)
	mvpStats := playerStatsByID(match, mvp.PlayerID)
	topPlayer := firstPlayer(match)
	scoreText := fmt.Sprintf("%d-%d", match.TeamScore, match.OpponentScore)

	article := NewsArticle{
		Title:   fmt.Sprintf("青空シャークスが%sに%sで勝利", match.OpponentName, scoreText),
		Source:  "Team Sports",
		Time:    match.MatchDate.Format("2006.01.02 15:04"),
		Summary: fmt.Sprintf("青空シャークスは%sとの試合を%sで制し、%sの活躍が光った。", match.OpponentName, scoreText, mvp.PlayerName),
		Lead:    fmt.Sprintf("%d番%sが攻撃の起点を作り、MVPの%sが勝負どころで流れを引き寄せた。", topPlayer.BattingOrder, topPlayer.PlayerName, mvp.PlayerName),
		Paragraphs: []string{
			fmt.Sprintf("青空シャークスは%sを相手に%sで勝利した。", match.OpponentName, scoreText),
			fmt.Sprintf("%sは%d安打%d打点の内容で、チームの勝利に大きく貢献した。", mvp.PlayerName, mvpStats.Hits, mvpStats.RBI),
			fmt.Sprintf("%sも出塁と得点で攻撃を支え、打線全体に流れを作った。", topPlayer.PlayerName),
		},
	}
	return article, nil
}

func playerStatsByID(match review.Match, playerID int) review.PlayerMatchStats {
	for _, stats := range match.PlayerStats {
		if stats.PlayerID == playerID {
			return stats
		}
	}

	return review.PlayerMatchStats{}
}

func firstPlayer(match review.Match) review.PlayerMatchStats {
	if len(match.PlayerStats) == 0 {
		return review.PlayerMatchStats{}
	}

	return match.PlayerStats[0]
}
