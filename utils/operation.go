// Package utils operation
package utils

// Div 除算結果を取得する
func Div(x, y int) (int, int) {
	// 商
	q := x / y
	// 余り
	r := x % y

	return q, r
}
