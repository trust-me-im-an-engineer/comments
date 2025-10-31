package post

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/trust-me-im-an-engineer/comments/internal/model"
	"github.com/trust-me-im-an-engineer/comments/internal/storage"
)

type Service struct {
	storage storage.Storage
}

func (s *Service) GetPost(ctx context.Context, id int) (*model.Post, error) {
	panic("unimplemented")
}

func NewService(storage storage.Storage) *Service {
	return &Service{storage}
}

func (s *Service) CreatePost(ctx context.Context, createPostInput *model.CreatePostInput) (*model.Post, error) {
	post, err := s.storage.CreatePost(ctx, createPostInput)
	if err != nil {
		return nil, fmt.Errorf("storage failed to create post: %w", err)
	}

	slog.Debug("post created", "postID", post.ID, "authorID", post.AuthorID)
	return post, nil
}
