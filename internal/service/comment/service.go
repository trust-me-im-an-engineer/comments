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
