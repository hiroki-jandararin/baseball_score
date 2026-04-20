package sqlite

import (
	"context"
	"database/sql"
	"os"
	"path/filepath"
	"testing"
	"time"

	review "baseball-score-app/backend/internal/review/domain"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMatchRepositorySaveMatch(t *testing.T) {
	t.Parallel()

	db := newTestDB(t)
	repo := NewMatchRepository(db)

	input := review.Match{
		TeamID:        1,
		OpponentName:  "Rivals",
		MatchDate:     time.Date(2026, 4, 18, 10, 0, 0, 0, time.UTC),
		Location:      "Tokyo",
		IsWin:         1,
		TeamScore:     5,
		OpponentScore: 3,
		Note:          "great finish",
		PlayerStats: []review.PlayerMatchStats{
			{
				PlayerID:        1,
				PlayerName:      "山田",
				BattingOrder:    1,
				Position:        "CF",
				Hits:            2,
				AtBats:          4,
				RBI:             1,
				Runs:            1,
				Walks:           1,
				Strikeouts:      0,
				Errors:          0,
				GoodPlay:        1,
				HighlightMoment: 1,
				Memo:            "lead-off spark",
			},
			{
				PlayerID:        2,
				PlayerName:      "佐藤",
				BattingOrder:    4,
				Position:        "1B",
				Hits:            1,
				AtBats:          3,
				RBI:             2,
				Runs:            1,
				Walks:           0,
				Strikeouts:      1,
				Errors:          0,
				GoodPlay:        0,
				HighlightMoment: 1,
				Memo:            "go-ahead hit",
			},
		},
	}

	got, err := repo.SaveMatch(context.Background(), input)
	require.NoError(t, err)

	require.NotZero(t, got)
	assertMatchRow(t, db, got, input)
	assertPlayerStatsRows(t, db, got, input)
}

func TestMatchRepositoryFindMatchByID(t *testing.T) {
	t.Parallel()

	db := newTestDB(t)
	repo := NewMatchRepository(db)

	input := review.Match{
		TeamID:        1,
		OpponentName:  "Rivals",
		MatchDate:     time.Date(2026, 4, 18, 10, 0, 0, 0, time.UTC),
		Location:      "Tokyo",
		IsWin:         1,
		TeamScore:     5,
		OpponentScore: 3,
		Note:          "great finish",
		PlayerStats: []review.PlayerMatchStats{
			{
				PlayerID:        1,
				PlayerName:      "山田",
				BattingOrder:    1,
				Position:        "CF",
				Hits:            2,
				AtBats:          4,
				RBI:             1,
				Runs:            1,
				Walks:           1,
				Strikeouts:      0,
				Errors:          0,
				GoodPlay:        1,
				HighlightMoment: 1,
				Memo:            "lead-off spark",
			},
			{
				PlayerID:        2,
				PlayerName:      "佐藤",
				BattingOrder:    4,
				Position:        "1B",
				Hits:            1,
				AtBats:          3,
				RBI:             2,
				Runs:            1,
				Walks:           0,
				Strikeouts:      1,
				Errors:          0,
				GoodPlay:        0,
				HighlightMoment: 1,
				Memo:            "go-ahead hit",
			},
		},
	}

	matchID, err := repo.SaveMatch(context.Background(), input)
	require.NoError(t, err)

	got, err := repo.FindMatchByID(context.Background(), matchID)
	require.NoError(t, err)

	assert.Equal(t, matchID, got.ID)
	assert.Equal(t, input.TeamID, got.TeamID)
	assert.Equal(t, input.OpponentName, got.OpponentName)
	assert.Equal(t, input.MatchDate, got.MatchDate)
	assert.Equal(t, input.Location, got.Location)
	assert.Equal(t, input.IsWin, got.IsWin)
	assert.Equal(t, input.TeamScore, got.TeamScore)
	assert.Equal(t, input.OpponentScore, got.OpponentScore)
	assert.Equal(t, input.Note, got.Note)
	assert.False(t, got.CreatedAt.IsZero())
	assert.False(t, got.UpdatedAt.IsZero())

	require.Len(t, got.PlayerStats, 2)
	assertPlayerStatsEqual(t, matchID, input.PlayerStats[0], got.PlayerStats[0])
	assertPlayerStatsEqual(t, matchID, input.PlayerStats[1], got.PlayerStats[1])
}

func TestMatchRepositoryFindMatchByIDReturnsErrorWhenMatchNotFound(t *testing.T) {
	t.Parallel()

	db := newTestDB(t)
	repo := NewMatchRepository(db)

	got, err := repo.FindMatchByID(context.Background(), 999)

	require.Error(t, err)
	assert.Equal(t, review.Match{}, got)
}

func newTestDB(t *testing.T) *sql.DB {
	t.Helper()

	db, err := sql.Open("sqlite3", ":memory:")
	require.NoError(t, err)
	t.Cleanup(func() {
		_ = db.Close()
	})

	schemaPath := filepath.Join("..", "..", "..", "platform", "db", "schema.sql")
	schema, err := os.ReadFile(schemaPath)
	require.NoError(t, err)

	_, err = db.Exec(string(schema))
	require.NoError(t, err)

	seedTestData(t, db)

	return db
}

func assertPlayerStatsEqual(t *testing.T, matchID int, expected review.PlayerMatchStats, actual review.PlayerMatchStats) {
	t.Helper()

	assert.NotZero(t, actual.ID)
	assert.Equal(t, matchID, actual.MatchID)
	assert.Equal(t, expected.PlayerID, actual.PlayerID)
	assert.Equal(t, expected.PlayerName, actual.PlayerName)
	assert.Equal(t, expected.BattingOrder, actual.BattingOrder)
	assert.Equal(t, expected.Position, actual.Position)
	assert.Equal(t, expected.Hits, actual.Hits)
	assert.Equal(t, expected.AtBats, actual.AtBats)
	assert.Equal(t, expected.RBI, actual.RBI)
	assert.Equal(t, expected.Runs, actual.Runs)
	assert.Equal(t, expected.Walks, actual.Walks)
	assert.Equal(t, expected.Strikeouts, actual.Strikeouts)
	assert.Equal(t, expected.Errors, actual.Errors)
	assert.Equal(t, expected.GoodPlay, actual.GoodPlay)
	assert.Equal(t, expected.HighlightMoment, actual.HighlightMoment)
	assert.Equal(t, expected.Memo, actual.Memo)
	assert.False(t, actual.CreatedAt.IsZero())
	assert.False(t, actual.UpdatedAt.IsZero())
}

func seedTestData(t *testing.T, db *sql.DB) {
	t.Helper()

	now := time.Date(2026, 4, 18, 0, 0, 0, 0, time.UTC).Format(time.RFC3339)

	_, err := db.Exec(`INSERT INTO teams (id, name, created_at, updated_at) VALUES (1, 'Sharks', ?, ?)`, now, now)
	require.NoError(t, err)
	_, err = db.Exec(`INSERT INTO players (id, team_id, name, created_at, updated_at) VALUES (1, 1, '山田', ?, ?)`, now, now)
	require.NoError(t, err)
	_, err = db.Exec(`INSERT INTO players (id, team_id, name, created_at, updated_at) VALUES (2, 1, '佐藤', ?, ?)`, now, now)
	require.NoError(t, err)
}

func assertMatchRow(t *testing.T, db *sql.DB, matchID int, input review.Match) {
	t.Helper()

	var opponentName string
	var teamScore, opponentScore, isWin int
	var location, note string

	err := db.QueryRow(`
		SELECT opponent_name, team_score, opponent_score, is_win, location, note
		FROM matches
		WHERE id = ?
	`, matchID).Scan(&opponentName, &teamScore, &opponentScore, &isWin, &location, &note)
	require.NoError(t, err)

	assert.Equal(t, input.OpponentName, opponentName)
	assert.Equal(t, input.TeamScore, teamScore)
	assert.Equal(t, input.OpponentScore, opponentScore)
	assert.Equal(t, input.IsWin, isWin)
	assert.Equal(t, input.Location, location)
	assert.Equal(t, input.Note, note)
}

func assertPlayerStatsRows(t *testing.T, db *sql.DB, matchID int, input review.Match) {
	t.Helper()

	rows, err := db.Query(`
		SELECT match_id, player_id, batting_order, position, hits, at_bats, rbi, runs, walks, strikeouts, errors, good_play, highlight_moment, memo
		FROM player_match_stats
		WHERE match_id = ?
		ORDER BY id
	`, matchID)
	require.NoError(t, err)
	defer rows.Close()

	var actual []review.PlayerMatchStats
	for rows.Next() {
		var stats review.PlayerMatchStats
		err := rows.Scan(
			&stats.MatchID,
			&stats.PlayerID,
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
		)
		require.NoError(t, err)
		actual = append(actual, stats)
	}
	require.NoError(t, rows.Err())

	require.Len(t, actual, 2)
	assert.Equal(t, matchID, actual[0].MatchID)
	assert.Equal(t, input.PlayerStats[0].PlayerID, actual[0].PlayerID)
	assert.Equal(t, input.PlayerStats[0].BattingOrder, actual[0].BattingOrder)
	assert.Equal(t, input.PlayerStats[0].Position, actual[0].Position)
	assert.Equal(t, input.PlayerStats[0].Hits, actual[0].Hits)
	assert.Equal(t, input.PlayerStats[0].AtBats, actual[0].AtBats)
	assert.Equal(t, input.PlayerStats[0].RBI, actual[0].RBI)
	assert.Equal(t, input.PlayerStats[0].Runs, actual[0].Runs)
	assert.Equal(t, input.PlayerStats[0].Walks, actual[0].Walks)
	assert.Equal(t, input.PlayerStats[0].Strikeouts, actual[0].Strikeouts)
	assert.Equal(t, input.PlayerStats[0].Errors, actual[0].Errors)
	assert.Equal(t, input.PlayerStats[0].GoodPlay, actual[0].GoodPlay)
	assert.Equal(t, input.PlayerStats[0].HighlightMoment, actual[0].HighlightMoment)
	assert.Equal(t, input.PlayerStats[0].Memo, actual[0].Memo)

	assert.Equal(t, matchID, actual[1].MatchID)
	assert.Equal(t, input.PlayerStats[1].PlayerID, actual[1].PlayerID)
	assert.Equal(t, input.PlayerStats[1].BattingOrder, actual[1].BattingOrder)
	assert.Equal(t, input.PlayerStats[1].Position, actual[1].Position)
	assert.Equal(t, input.PlayerStats[1].Hits, actual[1].Hits)
	assert.Equal(t, input.PlayerStats[1].AtBats, actual[1].AtBats)
	assert.Equal(t, input.PlayerStats[1].RBI, actual[1].RBI)
	assert.Equal(t, input.PlayerStats[1].Runs, actual[1].Runs)
	assert.Equal(t, input.PlayerStats[1].Walks, actual[1].Walks)
	assert.Equal(t, input.PlayerStats[1].Strikeouts, actual[1].Strikeouts)
	assert.Equal(t, input.PlayerStats[1].Errors, actual[1].Errors)
	assert.Equal(t, input.PlayerStats[1].GoodPlay, actual[1].GoodPlay)
	assert.Equal(t, input.PlayerStats[1].HighlightMoment, actual[1].HighlightMoment)
	assert.Equal(t, input.PlayerStats[1].Memo, actual[1].Memo)
}
