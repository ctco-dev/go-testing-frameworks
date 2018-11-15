package ginkgo

type Book struct {
	Title  string
	Author string
	Pages  uint
}

type BookType string

const (
	Novel      BookType = "NOVEL"
	ShortStory BookType = "SHORT STORY"
)

func (book Book) CategoryByLength() BookType {
	switch p := book.Pages; {
	case p > 20:
		return Novel
	default:
		return ShortStory
	}
}
