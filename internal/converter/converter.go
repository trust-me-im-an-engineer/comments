package converter

import (
	"strconv"

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

func PostToGQL(internal *internalmodel.Post) *gqlmodel.Post {

	return &gqlmodel.Post{
		ID:                 strconv.Itoa(internal.ID),
		AuthorID:           internal.AuthorID,
		Title:              internal.Title,
		Content:            internal.Content,
		CreatedAt:          internal.CreatedAt,
		Rating:             internal.Rating,
		CommentsCount:      internal.CommentsCount,
		CommentsRestricted: internal.CommentsRestricted,
	}
}
