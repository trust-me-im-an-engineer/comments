package storage

import (
	"context"
	"errors"

	"github.com/trust-me-im-an-engineer/comments/internal/model"
)

var (
	PostNotFound = errors.New("post not found")
)

type Storage interface {
	CreatePost(ctx context.Context, input *model.CreatePostInput) (*model.Post, error)
	GetPost(ctx context.Context, id int) (*model.Post, error)
	UpdatePost(ctx context.Context, input *model.UpdatePostInput) (*model.Post, error)
}
