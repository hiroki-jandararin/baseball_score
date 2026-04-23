package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"baseball-score-app/backend/internal/platform/db"
)

func main() {
	database, err := db.NewSQLiteConnection()
	if err != nil {
		log.Fatalf("database connection failed: %v", err)
	}
	defer func(database *sql.DB) {
		if err := database.Close(); err != nil {
			log.Printf("failed to close database: %v", err)
		}
	}(database)

	if err := applySchema(database); err != nil {
		log.Fatalf("schema setup failed: %v", err)
	}

	if err := seed(database); err != nil {
		log.Fatalf("seed failed: %v", err)
	}

	fmt.Println("seed completed")
}

func applySchema(database *sql.DB) error {
	schemaPath := filepath.Join("internal", "platform", "db", "schema.sql")
	schema, err := os.ReadFile(schemaPath)
	if err != nil {
		return fmt.Errorf("read schema: %w", err)
	}

	if _, err := database.Exec(string(schema)); err != nil {
		return fmt.Errorf("apply schema: %w", err)
	}

	return nil
}

func seed(database *sql.DB) error {
	tx, err := database.Begin()
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}

	committed := false
	defer func() {
		if !committed {
			_ = tx.Rollback()
		}
	}()

	const now = "2026-04-18T00:00:00Z"

	statements := []string{
		`INSERT OR REPLACE INTO teams (id, name, created_at, updated_at)
		 VALUES (1, '青空シャークス', '` + now + `', '` + now + `')`,
		`INSERT OR REPLACE INTO players (id, team_id, name, created_at, updated_at)
		 VALUES
		 	(1, 1, '山田', '` + now + `', '` + now + `'),
		 	(2, 1, '佐藤', '` + now + `', '` + now + `'),
		 	(3, 1, '鈴木', '` + now + `', '` + now + `')`,
		`INSERT OR REPLACE INTO matches (
			id, team_id, opponent_name, match_date, location, is_win,
			team_score, opponent_score, note, created_at, updated_at
		 )
		 VALUES (
		 	1, 1, '東町ライバルズ', '2026-04-18T10:00:00Z', '河川敷グラウンド', 1,
		 	5, 3, '終盤の集中打で接戦を制した', '` + now + `', '` + now + `'
		 )`,
		`DELETE FROM player_match_stats WHERE match_id = 1`,
		`INSERT INTO player_match_stats (
			match_id, player_id, batting_order, position,
			hits, at_bats, rbi, runs, walks, strikeouts, errors,
			good_play, highlight_moment, memo, created_at, updated_at
		 )
		 VALUES
		 	(1, 1, 1, '中堅手', 1, 4, 0, 1, 1, 1, 0, 1, 0, '先頭打者として出塁し攻撃の流れを作った', '` + now + `', '` + now + `'),
		 	(1, 2, 4, '一塁手', 2, 4, 2, 1, 0, 0, 0, 0, 1, '勝ち越しにつながる一打でチームを勢いづけた', '` + now + `', '` + now + `'),
		 	(1, 3, 6, '遊撃手', 1, 3, 1, 0, 0, 1, 0, 1, 0, '守備で落ち着いたプレーを見せた', '` + now + `', '` + now + `')`,
	}

	for _, statement := range statements {
		if _, err := tx.Exec(statement); err != nil {
			return fmt.Errorf("execute seed statement: %w", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit transaction: %w", err)
	}
	committed = true

	return nil
}
