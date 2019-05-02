package utils

func Div(x, y int) (int, int) {
	// 商
	q := x / y
	// 余り
	r := x % y

	return q, r
}
