package errs

import (
	"errors"
)

var (
	PostNotFound    = errors.New("post not found")
	CommentNotFound = errors.New("comment not found")
)
