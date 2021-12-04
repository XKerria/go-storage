package storage

import "errors"

var (
	ErrFileNotFound     = errors.New("file not found")
	ErrFileNoPermission = errors.New("permission deny")
)
