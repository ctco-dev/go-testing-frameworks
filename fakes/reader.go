package fakes

type reader interface {
	Read() string
}

type myReader struct{}

func NewMyReader() myReader {
	return myReader{}
}

func (r myReader) Read() string {
	return "Some text"
}
