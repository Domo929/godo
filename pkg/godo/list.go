package godo

import (
	"fmt"
)

type List struct {
	CurrentTopic string            `json:"current_topic"`
	Topics       map[string][]Item `json:"topics,omitempty"`
}

type Item struct {
	Completion float64 `json:"completion"`
	Name       string  `json:"name"`
}

func NewList() *List {
	l := new(List)
	l.Topics = make(map[string][]Item)
	return l
}

func (l *List) isValid() error {
	if _, ok := l.Topics[l.CurrentTopic]; !ok {
		return fmt.Errorf("%w: current topic (%s) not in list of topics", ErrInvalidConfiguration, l.CurrentTopic)
	}

	return nil
}
