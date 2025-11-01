package storage

import (
	"context"

	"github.com/trust-me-im-an-engineer/comments/internal/domain"
)

type Storage interface {
	CreatePost(ctx context.Context, input *domain.CreatePostInput) (*domain.Post, error)
	GetPost(ctx context.Context, id int) (*domain.Post, error)
	UpdatePost(ctx context.Context, input *domain.UpdatePostInput) (*domain.Post, error)
	DeletePost(ctx context.Context, id int) error
	SetCommentsRestricted(ctx context.Context, id int, restricted bool) (*domain.Post, error)
	VotePost(ctx context.Context, vote *domain.PostVote) (*domain.Post, error)
	CreateComment(ctx context.Context, input *domain.CreateCommentInput) (*domain.Comment, error)
	UpdateComment(ctx context.Context, input *domain.UpdateCommentInput) (*domain.Comment, error)
}
