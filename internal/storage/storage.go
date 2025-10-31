package storage

import (
	"context"

	"github.com/trust-me-im-an-engineer/comments/internal/model"
)

type Storage interface {
	CreatePost(ctx context.Context, input *model.CreatePostInput) (*model.Post, error)
}
