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

func ModelVoteInputToDomainPostVote(m *model.VoteInput) *domain.PostVote {
	id, _ := strconv.Atoi(m.ID) // id already validated
	return &domain.PostVote{
		PostID: id,
		Vote: domain.Vote{
			VoterID: m.VoterID,
			Value:   int8(m.Value),
		},
	}
}

func CreateCommentInput_ModelToDomain(m *model.CreateCommentInput) *domain.CreateCommentInput {
	postID, _ := strconv.Atoi(m.PostID)
	d := &domain.CreateCommentInput{
		PostID:   postID,
		AuthorID: m.AuthorID,
		Text:     m.Text,
		ParentID: nil,
	}
	if m.ParentID != nil {
		parentID, _ := strconv.Atoi(*m.ParentID)
		d.ParentID = &parentID
	}
	return d
}

func Comment_DomainToModel(d *domain.Comment) *model.Comment {
	m := &model.Comment{
		ID:       strconv.Itoa(d.ID),
		PostID:   strconv.Itoa(d.PostID),
		AuthorID: d.AuthorID,
		// Text:      d.Text,
		CreatedAt: d.CreatedAt,
		Rating:    d.Rating,
		ParentID:  nil,
	}

	if d.Text != nil {
		m.Text = *d.Text
	}
	return m
}

func UpdateCommentInput_ModelToDomain(m *model.UpdateCommentInput) *domain.UpdateCommentInput {
	id, _ := strconv.Atoi(m.ID) // id already validated
	return &domain.UpdateCommentInput{
		ID:   id,
		Text: m.Text,
	}
}

func ModelVoteInputToDomainCommentVote(m *model.VoteInput) *domain.CommentVote {
	id, _ := strconv.Atoi(m.ID) // id already validated
	return &domain.CommentVote{
		CommentID: id,
		Vote: domain.Vote{
			VoterID: m.VoterID,
			Value:   int8(m.Value),
		},
	}
}
