package validator

import (
	"errors"
	"strconv"

	gqlmodel "github.com/trust-me-im-an-engineer/comments/graph/model"
)

const (
	MaxTitleLen   = 200
	MaxContentLen = 20000
)

var (
	EmptyTitleErr     = errors.New("post title cannot be empty")
	EmptyContentErr   = errors.New("post content cannot be empty")
	TooLongTitleErr   = errors.New("post title cannot be longer than " + strconv.Itoa(MaxTitleLen) + " characters")
	TooLongContentErr = errors.New("post content cannot be longer than " + strconv.Itoa(MaxContentLen) + " characters")
)

func ValidateCreatePostInput(in gqlmodel.CreatePostInput) error {
	if in.Title == "" {
		return EmptyTitleErr
	}
	if in.Content == "" {
		return EmptyContentErr
	}
	if len(in.Title) > MaxTitleLen {
		return TooLongTitleErr
	}
	if len(in.Content) > MaxContentLen {
		return TooLongContentErr
	}
	return nil
}
