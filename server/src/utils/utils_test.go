package utils

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestAtoi(t *testing.T) {
	assert.Equal(t, Atoi("123", 0), 123)
	assert.Equal(t, Atoi("いちにさん", 0), 0)
}

func TestEnvironment(t *testing.T) {
	assert.NotEqual(t, os.Getenv("GCP_PROJECT_ID"), "")
	assert.NotEqual(t, os.Getenv("TIMEBOOSTER_SERVICE_ACCOUNT"), "")
}
