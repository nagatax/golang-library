package utils

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
 * ラップアラウンドの判定チェック
 */
func Testラップアラウンドが検知されること(t *testing.T) {
	// 実測値
	actual := IsWrapAround(math.MaxUint32, 1)
	assert.True(t, actual, "")

	// 実測値
	actual = IsWrapAround(1, math.MaxUint32)
	assert.True(t, actual, "")
}

/**
 * ラップアラウンドの判定チェック
 */
func Testラップアラウンドが検知されないこと(t *testing.T) {
	// 実測値
	actual := IsWrapAround(math.MaxUint32, 0)
	assert.False(t, actual, "")

	// 実測値
	actual = IsWrapAround(0, math.MaxUint32)
	assert.False(t, actual, "")
}
