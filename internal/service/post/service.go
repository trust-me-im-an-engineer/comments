package post

import "github.com/trust-me-im-an-engineer/comments/internal/storage"

type Service struct {
	storage *storage.Storage
}

func NewService(storage *storage.Storage) *Service {
	return &Service{storage}
}
