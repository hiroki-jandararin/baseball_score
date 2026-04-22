package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"baseball-score-app/backend/internal/platform/config"
	"baseball-score-app/backend/internal/platform/db"
	sqlitereview "baseball-score-app/backend/internal/review/adapter/sqlite"
	"baseball-score-app/backend/internal/review/usecase"
)

type healthResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	database, err := db.NewSQLiteConnection()
	if err != nil {
		log.Fatalf("database connection failed: %v", err)
	}
	defer func(database *sql.DB) {
		if err := database.Close(); err != nil {
			log.Printf("failed to close database: %v", err)
		}
	}(database)

	openAIConfig, err := config.LoadOpenAI()
	if err != nil {
		log.Fatalf("openai config invalid: %v", err)
	}
	if openAIConfig.APIKey != "" {
		log.Printf("openai integration enabled with model %s", openAIConfig.Model)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)

	matchRepository := sqlitereview.NewMatchRepository(database)
	newsService := usecase.NewNewsService(matchRepository)
	mux.Handle("/matches/", newNewsHandler(newsService))

	handler := corsMiddleware(mux)

	log.Printf("server starting on :%s", port)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	response := healthResponse{
		Status:  "ok",
		Message: "backend is running",
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}

func corsMiddleware(next http.Handler) http.Handler {
	allowOrigin := os.Getenv("CORS_ALLOW_ORIGIN")
	if allowOrigin == "" {
		allowOrigin = "http://localhost:5173"
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", allowOrigin)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
