package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetTimeboosterProjectId(t *testing.T) {
	assert.NotEqual(t, GetTimeboosterProjectId(), "") // 何らかの値が入っている
}
