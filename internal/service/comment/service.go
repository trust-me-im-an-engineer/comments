package comment

import (
	"context"
	"fmt"

	"github.com/trust-me-im-an-engineer/comments/internal/domain"
	"github.com/trust-me-im-an-engineer/comments/internal/storage"
)

type Service struct {
	storage storage.Storage
}

func (s *Service) DeleteComment(ctx context.Context, domainID int) error {
	err := s.storage.DeleteComment(ctx, domainID)
	if err != nil {
		return fmt.Errorf("storage failed to delete comment: %w", err)
	}
	return nil
}

func (s *Service) UpdateComment(ctx context.Context, domainInput *domain.UpdateCommentInput) (*domain.Comment, error) {
	comment, err := s.storage.UpdateCommentIfNotDeleted(ctx, domainInput)
	if err != nil {
		return nil, fmt.Errorf("storage failed to update comment: %w", err)
	}
	return comment, nil
}

func (s *Service) CreateComment(ctx context.Context, domainInput *domain.CreateCommentInput) (*domain.Comment, error) {
	comment, err := s.storage.CreateComment(ctx, domainInput)
	if err != nil {
		return nil, fmt.Errorf("storage failed to create comment: %w", err)
	}
	return comment, nil
}

func NewService(storage storage.Storage) *Service {
	return &Service{storage}
}
