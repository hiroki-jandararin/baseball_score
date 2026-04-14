package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	defaultOpenAIModel   = "gpt-4.1-mini"
	defaultOpenAIBaseURL = "https://api.openai.com/v1"
	defaultOpenAITimeout = 30 * time.Second
)

type OpenAI struct {
	APIKey  string
	Model   string
	Project string
	BaseURL string
	Timeout time.Duration
}

func LoadOpenAI() (OpenAI, error) {
	enabled := strings.TrimSpace(os.Getenv("OPENAI_API_KEY")) != "" ||
		strings.TrimSpace(os.Getenv("OPENAI_MODEL")) != "" ||
		strings.TrimSpace(os.Getenv("OPENAI_PROJECT")) != ""

	if !enabled {
		return OpenAI{}, nil
	}

	apiKey := strings.TrimSpace(os.Getenv("OPENAI_API_KEY"))
	if apiKey == "" {
		return OpenAI{}, fmt.Errorf("OPENAI_API_KEY is required when OpenAI is enabled")
	}

	timeout, err := loadOpenAITimeout()
	if err != nil {
		return OpenAI{}, err
	}

	model := strings.TrimSpace(os.Getenv("OPENAI_MODEL"))
	if model == "" {
		model = defaultOpenAIModel
	}

	baseURL := strings.TrimRight(strings.TrimSpace(os.Getenv("OPENAI_BASE_URL")), "/")
	if baseURL == "" {
		baseURL = defaultOpenAIBaseURL
	}

	return OpenAI{
		APIKey:  apiKey,
		Model:   model,
		Project: strings.TrimSpace(os.Getenv("OPENAI_PROJECT")),
		BaseURL: baseURL,
		Timeout: timeout,
	}, nil
}

func loadOpenAITimeout() (time.Duration, error) {
	raw := strings.TrimSpace(os.Getenv("OPENAI_TIMEOUT_SECONDS"))
	if raw == "" {
		return defaultOpenAITimeout, nil
	}

	seconds, err := strconv.Atoi(raw)
	if err != nil || seconds <= 0 {
		return 0, fmt.Errorf("OPENAI_TIMEOUT_SECONDS must be a positive integer")
	}

	return time.Duration(seconds) * time.Second, nil
}
