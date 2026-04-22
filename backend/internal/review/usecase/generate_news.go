package usecase

import (
	"context"

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

	article := NewsArticle{
		Title:   "SharksがRivalsに5-3で勝利",
		Source:  "Team Sports",
		Time:    match.MatchDate.Format("2006.01.02 15:04"),
		Summary: "SharksはRivalsとの試合を5-3で制し、佐藤の活躍が光った。",
		Lead:    "1番山田が攻撃の起点を作り、MVPの佐藤が勝負どころで流れを引き寄せた。",
		Paragraphs: []string{
			"SharksはRivalsを相手に5-3で勝利した。",
			"佐藤は2安打2打点の内容で、チームの勝利に大きく貢献した。",
			"山田も出塁と得点で攻撃を支え、打線全体に流れを作った。",
		},
	}
	return article, nil
}
