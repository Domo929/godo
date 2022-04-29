package godo

import "errors"

var (
	ErrBadPath            = errors.New("bad path")
	ErrBadTopicName       = errors.New("bad topic name")
	ErrTopicAlreadyExists = errors.New("topic already exists")
)
