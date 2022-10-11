package token

import (
	"errors"
)

var (
	ErrInvalidClaim       = errors.New("invalid jwt claim")
	ErrSignAlgoNotMatched = errors.New("sign algorithm didn't match")
	ErrRefreshingToken    = errors.New("error when refreshing token")
)

type Token interface {
	Create(userID int64) (string, error)
	Validate(token string) (string, error)
}
