package errs

import (
	"errors"
)

var (
	PostNotFound    = errors.New("post not found")
	CommentNotFound = errors.New("comment not found")
	InvalidIDErr    = errors.New("id must be valid integer")
	CommentDeleted  = errors.New("comment is deleted")
)
