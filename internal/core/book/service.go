package book

import (
	"fmt"
)

type Service struct {
	repository bookRepository
}

type bookRepository interface {
	Add(book Book) error
}

func NewService(repo bookRepository) *Service {
	return &Service{
		repository: repo,
	}
}

func (s *Service) AddBook(book Book) error {
	err := s.repository.Add(book)
	if err != nil {
		return fmt.Errorf("could not add book: %w", err)
	}
	return nil
}
