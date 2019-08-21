package utils

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOS名が取得できること(t *testing.T) {
	expect := runtime.GOOS
	actual := GetOSName()
	assert.Equal(t, expect, actual, "")
}

func Testバージョン番号が表示されないこと(t *testing.T) {
	expect := false
	actual := IsPrintVersion()
	assert.Equal(t, expect, actual, "")
}

func Testバージョン番号が表示されること(t *testing.T) {
	expect := true
	actual := IsPrintVersion("-v")
	assert.Equal(t, expect, actual, "")

	expect = true
	actual = IsPrintVersion("-version")
	assert.Equal(t, expect, actual, "")
}
