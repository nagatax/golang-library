package utils

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOS名が取得できること(t *testing.T) {
	expect := runtime.GOOS
	actual := ""

	sut := GetOSName

	actual = sut()
	assert.Equal(t, expect, actual, "")
}
