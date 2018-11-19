package fakes

//Reader reads strings from some source
type Reader interface {
	Read() string
}

type myReader struct{}

//NewReader creates a new Reader
func NewReader() Reader {
	return myReader{}
}

func (r myReader) Read() string {
	return "Some text"
}
