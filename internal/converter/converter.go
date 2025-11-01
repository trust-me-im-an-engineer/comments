package converter

import (
	"strconv"

	"github.com/trust-me-im-an-engineer/comments/graph/model"
	"github.com/trust-me-im-an-engineer/comments/internal/domain"
)

func CreatePostInput_ModelToDomain(m *model.CreatePostInput) *domain.CreatePostInput {
	return &domain.CreatePostInput{
		AuthorID: m.AuthorID,
		Title:    m.Title,
		Content:  m.Content,
	}
}

func Post_DomainToModel(d *domain.Post) *model.Post {

	return &model.Post{
		ID:                 strconv.Itoa(d.ID),
		AuthorID:           d.AuthorID,
		Title:              d.Title,
		Content:            d.Content,
		CreatedAt:          d.CreatedAt,
		Rating:             d.Rating,
		CommentsCount:      d.CommentsCount,
		CommentsRestricted: d.CommentsRestricted,
	}
}

func UpdatePost_ModelToDomain(m *model.UpdatePostInput) *domain.UpdatePostInput {
	id, _ := strconv.Atoi(m.ID) // id already validated
	return &domain.UpdatePostInput{
		ID:      id,
		Title:   m.Title,
		Content: m.Content,
	}
}

func ModelVoteInputToDomainPostVote(m *model.VoteInput, value int8) *domain.PostVote {
	id, _ := strconv.Atoi(m.ID) // id already validated
	return &domain.PostVote{
		PostID: id,
		Vote: domain.Vote{
			VoterID: m.VoterID,
			Value:   value,
		},
	}
}
