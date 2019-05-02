package utils

import (
	"math"
)

/**
 * ラップアラウンド対策
 * ラップアラウンドが発生するか検査する
 */
func IsWrapAround(x, y uint32) bool {
	if (math.MaxUint32 - x) < y {
		return true
	} else {
		return false
	}
}
