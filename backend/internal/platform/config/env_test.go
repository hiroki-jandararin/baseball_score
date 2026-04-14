package config

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadOpenAIReturnsZeroValueWhenDisabled(t *testing.T) {
	t.Setenv("OPENAI_API_KEY", "")
	t.Setenv("OPENAI_MODEL", "")
	t.Setenv("OPENAI_PROJECT", "")
	t.Setenv("OPENAI_BASE_URL", "")
	t.Setenv("OPENAI_TIMEOUT_SECONDS", "")

	cfg, err := LoadOpenAI()
	require.NoError(t, err)
	assert.Equal(t, OpenAI{}, cfg)
}

func TestLoadOpenAIUsesDefaults(t *testing.T) {
	t.Setenv("OPENAI_API_KEY", "test-key")
	t.Setenv("OPENAI_MODEL", "")
	t.Setenv("OPENAI_PROJECT", "")
	t.Setenv("OPENAI_BASE_URL", "")
	t.Setenv("OPENAI_TIMEOUT_SECONDS", "")

	cfg, err := LoadOpenAI()
	require.NoError(t, err)
	assert.Equal(t, "test-key", cfg.APIKey)
	assert.Equal(t, defaultOpenAIModel, cfg.Model)
	assert.Equal(t, defaultOpenAIBaseURL, cfg.BaseURL)
	assert.Equal(t, defaultOpenAITimeout, cfg.Timeout)
}

func TestLoadOpenAIReturnsErrorWhenKeyMissing(t *testing.T) {
	t.Setenv("OPENAI_API_KEY", "")
	t.Setenv("OPENAI_MODEL", "gpt-4.1")

	_, err := LoadOpenAI()
	require.Error(t, err)
	assert.Contains(t, err.Error(), "OPENAI_API_KEY")
}

func TestLoadOpenAIReadsCustomValues(t *testing.T) {
	t.Setenv("OPENAI_API_KEY", "test-key")
	t.Setenv("OPENAI_MODEL", "gpt-5-mini")
	t.Setenv("OPENAI_PROJECT", "proj_123")
	t.Setenv("OPENAI_BASE_URL", "https://example.test/v1/")
	t.Setenv("OPENAI_TIMEOUT_SECONDS", "45")

	cfg, err := LoadOpenAI()
	require.NoError(t, err)
	assert.Equal(t, "test-key", cfg.APIKey)
	assert.Equal(t, "gpt-5-mini", cfg.Model)
	assert.Equal(t, "proj_123", cfg.Project)
	assert.Equal(t, "https://example.test/v1", cfg.BaseURL)
	assert.Equal(t, 45*time.Second, cfg.Timeout)
}

func TestLoadOpenAIRejectsInvalidTimeout(t *testing.T) {
	t.Setenv("OPENAI_API_KEY", "test-key")
	t.Setenv("OPENAI_TIMEOUT_SECONDS", "abc")

	_, err := LoadOpenAI()
	require.Error(t, err)
	assert.Contains(t, err.Error(), "OPENAI_TIMEOUT_SECONDS")
}
