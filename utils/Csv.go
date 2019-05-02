package utils

type Csv struct {
	header string
	body   string
}

func (obj *Csv) SetHeader(header string) {
	obj.header = header
}

func (obj *Csv) SetBody(body string) {
	obj.body = body
}
