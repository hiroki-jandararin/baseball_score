package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"baseball-score-app/backend/internal/review/usecase"
)

type NewsArticleGenerator interface {
	GenerateArticle(ctx context.Context, matchID int) (usecase.NewsArticle, error)
}

func newNewsHandler(service NewsArticleGenerator) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		matchID, err := parseNewsMatchID(r.URL.Path)
		if err != nil {
			http.Error(w, "invalid match id", http.StatusBadRequest)
			return
		}

		article, err := service.GenerateArticle(r.Context(), matchID)
		if err != nil {
			if isNotFoundError(err) {
				http.Error(w, "match not found", http.StatusNotFound)
				return
			}
			http.Error(w, "failed to generate news", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(article); err != nil {
			http.Error(w, "failed to encode response", http.StatusInternalServerError)
		}
	})
}

func isNotFoundError(err error) bool {
	return errors.Is(err, sql.ErrNoRows) || strings.Contains(strings.ToLower(err.Error()), "not found")
}

func parseNewsMatchID(path string) (int, error) {
	path = strings.Trim(path, "/")
	parts := strings.Split(path, "/")
	if len(parts) != 3 || parts[0] != "matches" || parts[2] != "news" {
		return 0, strconv.ErrSyntax
	}

	return strconv.Atoi(parts[1])
}
