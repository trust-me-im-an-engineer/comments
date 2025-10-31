package post

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/trust-me-im-an-engineer/comments/internal/model"
	"github.com/trust-me-im-an-engineer/comments/internal/storage"
)

type Service struct {
	storage storage.Storage
}

func NewService(storage storage.Storage) *Service {
	return &Service{storage}
}

func (s *Service) GetPost(ctx context.Context, id int) (*model.Post, error) {
	post, err := s.storage.GetPost(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("storage failed to get post: %w", err)
	}
	return post, nil
}

func (s *Service) CreatePost(ctx context.Context, createPostInput *model.CreatePostInput) (*model.Post, error) {
	post, err := s.storage.CreatePost(ctx, createPostInput)
	if err != nil {
		return nil, fmt.Errorf("storage failed to create post: %w", err)
	}

	slog.Debug("post created", "postID", post.ID, "authorID", post.AuthorID)
	return post, nil
}

func (s *Service) UpdatePost(ctx context.Context, updatePostInput *model.UpdatePostInput) (*model.Post, error) {
	post, err := s.storage.UpdatePost(ctx, updatePostInput)
	if errors.Is(err, storage.PostNotFound) {
		return nil, fmt.Errorf("storage failed to update post: %w", err)
	}

	slog.Debug("post updated", "postID", post.ID)
	return post, nil
}
