package service

import "errors"

var (
	ErrOperations = errors.New("failed db operation")
	ErrCommitDB   = errors.New("failed when committing")
)
