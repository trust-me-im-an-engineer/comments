package converter

import (
	gqlmodel "github.com/trust-me-im-an-engineer/comments/graph/model"
	internalmodel "github.com/trust-me-im-an-engineer/comments/internal/model"
)

func CreatePostInputToInternal(gql *gqlmodel.CreatePostInput) *internalmodel.CreatePostInput {
	return &internalmodel.CreatePostInput{
		AuthorID: gql.AuthorID,
		Title:    gql.Title,
		Content:  gql.Content,
	}
}
