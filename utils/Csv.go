// Package utils csv
package utils

// Csv CSVファイルを取り扱う
type Csv struct {
	header string
	body   string
}

// SetHeader ヘッダ部を設定する
func (obj *Csv) SetHeader(header string) {
	obj.header = header
}

// SetBody ボディ部を設定する
func (obj *Csv) SetBody(body string) {
	obj.body = body
}
