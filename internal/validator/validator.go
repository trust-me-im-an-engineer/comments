package validator

import (
	"errors"
	"strconv"

	"github.com/trust-me-im-an-engineer/comments/graph/model"
)

const (
	MaxTitleLen   = 200
	MaxContentLen = 20000
)

var (
	EmptyTitleErr      = errors.New("post title cannot be empty")
	EmptyContentErr    = errors.New("post content cannot be empty")
	TooLongTitleErr    = errors.New("post title cannot be longer than " + strconv.Itoa(MaxTitleLen) + " characters")
	TooLongContentErr  = errors.New("post content cannot be longer than " + strconv.Itoa(MaxContentLen) + " characters")
	NothingToUpdateErr = errors.New("at least one field needed to update")
	InvalidPostID      = errors.New("post id must be valid integer")
)

func ValidateCreatePostInput(in model.CreatePostInput) error {
	if err := validateTitle(in.Title); err != nil {
		return err
	}
	return validateContent(in.Content)
}

func validateTitle(title string) error {
	if title == "" {
		return EmptyTitleErr
	}
	if len(title) > MaxTitleLen {
		return TooLongTitleErr
	}
	return nil
}

func validateContent(content string) error {
	if content == "" {
		return EmptyContentErr
	}
	if len(content) > MaxContentLen {
		return TooLongContentErr
	}
	return nil
}

func ValidateUpdatePostInput(in model.UpdatePostInput) error {
	if _, err := strconv.Atoi(in.ID); err != nil {
		return InvalidPostID
	}

	if in.Title == nil && in.Content == nil {
		return NothingToUpdateErr
	}

	if in.Title != nil {
		if err := validateTitle(*in.Title); err != nil {
			return err
		}
	}
	if in.Content != nil {
		if err := validateContent(*in.Content); err != nil {
			return err
		}
	}

	return nil
}
