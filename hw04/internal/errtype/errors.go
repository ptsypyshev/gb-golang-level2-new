package errtype

import "errors"

var (
	ErrBadRequest      = errors.New("bad request")
	ErrEmptyUserID     = errors.New("empty user ID")
	ErrBadUserID       = errors.New("bad user ID")
	ErrUserNotFound    = errors.New("user is not found")
	ErrFriendsNotFound = errors.New("friends are not found")
)
