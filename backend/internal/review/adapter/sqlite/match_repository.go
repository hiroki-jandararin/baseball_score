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

func (r *MatchRepository) FindMatchByID(ctx context.Context, matchID int) (review.Match, error) {
	var match review.Match
	var matchDate string
	var createdAt string
	var updatedAt string

	err := r.db.QueryRowContext(ctx, `
		SELECT id, team_id, opponent_name, match_date, location, is_win, team_score, opponent_score, note, created_at, updated_at
		FROM matches
		WHERE id = ?
	`, matchID).Scan(
		&match.ID,
		&match.TeamID,
		&match.OpponentName,
		&matchDate,
		&match.Location,
		&match.IsWin,
		&match.TeamScore,
		&match.OpponentScore,
		&match.Note,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		return review.Match{}, fmt.Errorf("find match: %w", err)
	}

	match.MatchDate, err = time.Parse(time.RFC3339, matchDate)
	if err != nil {
		return review.Match{}, fmt.Errorf("parse match date: %w", err)
	}
	match.CreatedAt, err = time.Parse(time.RFC3339, createdAt)
	if err != nil {
		return review.Match{}, fmt.Errorf("parse match created at: %w", err)
	}
	match.UpdatedAt, err = time.Parse(time.RFC3339, updatedAt)
	if err != nil {
		return review.Match{}, fmt.Errorf("parse match updated at: %w", err)
	}

	rows, err := r.db.QueryContext(ctx, `
		SELECT
			pms.id,
			pms.match_id,
			pms.player_id,
			p.name,
			pms.batting_order,
			pms.position,
			pms.hits,
			pms.at_bats,
			pms.rbi,
			pms.runs,
			pms.walks,
			pms.strikeouts,
			pms.errors,
			pms.good_play,
			pms.highlight_moment,
			pms.memo,
			pms.created_at,
			pms.updated_at
		FROM player_match_stats pms
		INNER JOIN players p ON p.id = pms.player_id
		WHERE pms.match_id = ?
		ORDER BY pms.batting_order, pms.id
	`, matchID)
	if err != nil {
		return review.Match{}, fmt.Errorf("find player stats: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var stats review.PlayerMatchStats
		var statsCreatedAt string
		var statsUpdatedAt string

		err := rows.Scan(
			&stats.ID,
			&stats.MatchID,
			&stats.PlayerID,
			&stats.PlayerName,
			&stats.BattingOrder,
			&stats.Position,
			&stats.Hits,
			&stats.AtBats,
			&stats.RBI,
			&stats.Runs,
			&stats.Walks,
			&stats.Strikeouts,
			&stats.Errors,
			&stats.GoodPlay,
			&stats.HighlightMoment,
			&stats.Memo,
			&statsCreatedAt,
			&statsUpdatedAt,
		)
		if err != nil {
			return review.Match{}, fmt.Errorf("scan player stats: %w", err)
		}

		stats.CreatedAt, err = time.Parse(time.RFC3339, statsCreatedAt)
		if err != nil {
			return review.Match{}, fmt.Errorf("parse player stats created at: %w", err)
		}
		stats.UpdatedAt, err = time.Parse(time.RFC3339, statsUpdatedAt)
		if err != nil {
			return review.Match{}, fmt.Errorf("parse player stats updated at: %w", err)
		}

		match.PlayerStats = append(match.PlayerStats, stats)
	}
	if err := rows.Err(); err != nil {
		return review.Match{}, fmt.Errorf("iterate player stats: %w", err)
	}

	return match, nil
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
