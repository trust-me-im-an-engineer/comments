package post

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/trust-me-im-an-engineer/comments/internal/domain"
	"github.com/trust-me-im-an-engineer/comments/internal/storage"
)

type Service struct {
	storage storage.Storage
}

func NewService(storage storage.Storage) *Service {
	return &Service{storage}
}

func (s *Service) GetPost(ctx context.Context, id int) (*domain.Post, error) {
	post, err := s.storage.GetPost(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("storage failed to get post: %w", err)
	}
	return post, nil
}

func (s *Service) CreatePost(ctx context.Context, createPostInput *domain.CreatePostInput) (*domain.Post, error) {
	post, err := s.storage.CreatePost(ctx, createPostInput)
	if err != nil {
		return nil, fmt.Errorf("storage failed to create post: %w", err)
	}

	slog.Debug("post created", "postID", post.ID, "authorID", post.AuthorID)
	return post, nil
}

func (s *Service) UpdatePost(ctx context.Context, updatePostInput *domain.UpdatePostInput) (*domain.Post, error) {
	post, err := s.storage.UpdatePost(ctx, updatePostInput)
	if errors.Is(err, storage.PostNotFound) {
		return nil, fmt.Errorf("storage failed to update post: %w", err)
	}

	slog.Debug("post updated", "postID", post.ID)
	return post, nil
}

func (s *Service) DeletePost(ctx context.Context, id int) error {
	err := s.storage.DeletePost(ctx, id)
	if errors.Is(err, storage.PostNotFound) {
		return fmt.Errorf("storage failed to delete post: %w", err)
	}

	if err != nil {
		return fmt.Errorf("storage failed to delete post: %w", err)
	}

	slog.Debug("post deleted", "postID", id)
	return nil
}

func (s *Service) SetCommentsRestricted(ctx context.Context, internalID int, restricted bool) (*domain.Post, error) {
	post, err := s.storage.SetCommentsRestricted(ctx, internalID, restricted)
	if err != nil {
		return nil, fmt.Errorf("storage failed to set comments restricted: %w", err)
	}

	slog.Debug("comments restriction changed", "postID", post.ID, "restricted", post.CommentsRestricted)
	return post, nil
}
