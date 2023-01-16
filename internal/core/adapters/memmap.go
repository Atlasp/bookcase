package adapters

import "github.com/Atlasp/bookcase/internal/core/domain/book"

type Memory struct {
	memMap map[int]book.Book
}

func NewMemMap() *Memory {
	memMap := make(map[int]book.Book)
	return &Memory{
		memMap: memMap,
	}
}

func (m *Memory) Add(book book.Book) error {
	m.memMap[1] = book
	return nil
}
