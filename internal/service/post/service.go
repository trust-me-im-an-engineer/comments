package post

import (
	"context"
	"fmt"

	"github.com/trust-me-im-an-engineer/comments/internal/model"
	"github.com/trust-me-im-an-engineer/comments/internal/storage"
)

type Service struct {
	storage storage.Storage
}

func NewService(storage storage.Storage) *Service {
	return &Service{storage}
}

func (s *Service) CreatePost(ctx context.Context, createPostInput *model.CreatePostInput) (*model.Post, error) {
	post, err := s.storage.CreatePost(ctx, createPostInput)
	if err != nil {
		return nil, fmt.Errorf("storage failed to create post: %w", err)
	}
	return post, nil
}
