package ginkgo

type Book struct {
	Title  string
	Author string
	Pages  uint
}

func (book Book) CategoryByLength() string {
	switch p := book.Pages; {
	case p > 300:
		return "NOVEL"
	default:
		return "SHORT STORY"
	}
}
