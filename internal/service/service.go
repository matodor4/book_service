package service

import (
	"context"
	"test_1/internal/domain"
)

type BookRepository interface {
	GetBooks(ctx context.Context) ([]domain.Book, error)
	DeleteBookByID(ctx context.Context, id string) error
}

type Service struct {
	Repo BookRepository
}

func NewService(repo BookRepository) *Service {
	return &Service{Repo: repo}
}
