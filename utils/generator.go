// Package utils generator
package utils

// GenerateInteger 整数を取得する
func GenerateInteger() func() int {
	i := 0

	return func() int {
		i++
		return i
	}

}
