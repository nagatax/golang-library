//go:build uint
// +build uint

package utils

import (
	"path/filepath"
	"runtime"
	"testing"
	"time"

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

func Test乱数が取得できること(t *testing.T) {
	var expect int64
	actual := GetRandomSeed()
	assert.IsType(t, expect, actual, "")

	actual1 := GetRandomSeed()
	time.Sleep(time.Millisecond)
	actual2 := GetRandomSeed()
	time.Sleep(time.Millisecond)
	actual3 := GetRandomSeed()
	assert.NotEqual(t, actual1, actual2)
	assert.NotEqual(t, actual1, actual3)
	assert.NotEqual(t, actual2, actual3)
}

func Test結合したファイルパスが取得できること(t *testing.T) {
	separator := string(filepath.Separator)
	actual := JoinFilePath([]string{"a", "b", "c"})
	assert.Equal(t, actual, "a"+separator+"b"+separator+"c")
}
