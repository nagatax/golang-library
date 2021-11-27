// Package utils validation
package utils

import (
	"math"
)

// IsWrapAround ラップアラウンド対策。ラップアラウンドが発生するか検査する
func IsWrapAround(x, y uint32) bool {
	if (math.MaxUint32 - x) < y {
		return true
	}

	return false
}
