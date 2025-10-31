package converter

import (
	"strconv"

	"github.com/trust-me-im-an-engineer/comments/graph/model"
	"github.com/trust-me-im-an-engineer/comments/internal/domain"
)

func CreatePostInputToInternal(gql *model.CreatePostInput) *domain.CreatePostInput {
	return &domain.CreatePostInput{
		AuthorID: gql.AuthorID,
		Title:    gql.Title,
		Content:  gql.Content,
	}
}

func PostToGQL(internal *domain.Post) *model.Post {

	return &model.Post{
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

func UpdatePostToInternal(gql *model.UpdatePostInput) *domain.UpdatePostInput {
	id, _ := strconv.Atoi(gql.ID) // id already validated
	return &domain.UpdatePostInput{
		ID:      id,
		Title:   gql.Title,
		Content: gql.Content,
	}
}
