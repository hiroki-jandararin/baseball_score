package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	review "baseball-score-app/backend/internal/review/domain"
)

type MatchRepository struct {
	db *sql.DB
}

func NewMatchRepository(db *sql.DB) *MatchRepository {
	return &MatchRepository{db: db}
}

func (r *MatchRepository) SaveMatch(ctx context.Context, match review.Match) (int, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, fmt.Errorf("begin transaction: %w", err)
	}

	committed := false
	defer func() {
		if !committed {
			_ = tx.Rollback()
		}
	}()

	now := time.Now().UTC().Format(time.RFC3339)

	result, err := tx.ExecContext(ctx, `
		INSERT INTO matches (
			team_id, opponent_name, match_date, location, is_win, team_score, opponent_score, note, created_at, updated_at
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`,
		match.TeamID,
		match.OpponentName,
		match.MatchDate.Format(time.RFC3339),
		match.Location,
		match.IsWin,
		match.TeamScore,
		match.OpponentScore,
		match.Note,
		now,
		now,
	)
	if err != nil {
		return 0, fmt.Errorf("insert match: %w", err)
	}

	matchID, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("read match id: %w", err)
	}

	for _, stats := range match.PlayerStats {
		_, err := tx.ExecContext(ctx, `
			INSERT INTO player_match_stats (
				match_id, player_id, batting_order, position, hits, at_bats, rbi, runs, walks, strikeouts, errors, good_play, highlight_moment, memo, created_at, updated_at
			) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		`,
			matchID,
			stats.PlayerID,
			stats.BattingOrder,
			stats.Position,
			stats.Hits,
			stats.AtBats,
			stats.RBI,
			stats.Runs,
			stats.Walks,
			stats.Strikeouts,
			stats.Errors,
			stats.GoodPlay,
			stats.HighlightMoment,
			stats.Memo,
			now,
			now,
		)
		if err != nil {
			return 0, fmt.Errorf("insert player stats: %w", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return 0, fmt.Errorf("commit transaction: %w", err)
	}
	committed = true

	return int(matchID), nil
}
