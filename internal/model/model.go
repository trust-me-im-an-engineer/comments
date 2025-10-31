// package model contains internal model representation different from graph/model
package model

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID                 int       `db:"id"`
	AuthorID           uuid.UUID `db:"author_id"`
	Title              string    `db:"title"`
	Content            string    `db:"content"`
	CreatedAt          time.Time `db:"created_at"`
	Rating             int32     `db:"rating"`
	CommentsCount      int32     `db:"comments_count"`
	CommentsRestricted bool      `db:"comments_restricted"`
}

type Comment struct {
	ID        int       `db:"id"`
	PostID    int       `db:"post_id"`
	AuthorID  uuid.UUID `db:"author_id"`
	Text      string    `db:"text"`
	CreatedAt time.Time `db:"created_at"`
	Rating    int32     `db:"rating"`
	ParentID  *int      `db:"parent_id"`
}

type PostVote struct {
	PostID int `db:"post_id"`
	Vote
}

type CommentVote struct {
	CommentID int `db:"comment_id"`
	Vote
}

type Vote struct {
	VoterID uuid.UUID `db:"voter_id"`
	// +1 for upvote, -1 for downvote
	Value int8 `db:"value"`
}

type CreatePostInput struct {
	AuthorID uuid.UUID `json:"authorID"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
}

type UpdatePostInput struct {
	ID      int     `json:"id"`
	Title   *string `json:"title,omitempty"`
	Content *string `json:"content,omitempty"`
}

type CreateCommentInput struct {
	PostID   int       `json:"postID"`
	AuthorID uuid.UUID `json:"authorID"`
	Text     string    `json:"text"`
	ParentID *int      `json:"parentID,omitempty"`
}

type UpdateCommentInput struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}
