package godo

import (
	"fmt"
	"strings"
	"unicode/utf8"
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

func (l *List) AddTopic(topic string) error {
	if topic == "" {
		return fmt.Errorf("%w: new topic must not be empty", ErrBadTopicName)
	}
	if !utf8.ValidString(topic) {
		return fmt.Errorf("%w: new topic must be valid utf8 string", ErrBadTopicName)
	}

	loweredTopic := strings.ToLower(topic)
	if _, ok := l.Topics[loweredTopic]; !ok {
		return fmt.Errorf("%w: new topic (%s) not in list of topics", ErrInvalidConfiguration, loweredTopic)
	}

	l.CurrentTopic = loweredTopic
	l.Topics[loweredTopic] = make([]Item, 0)

	return nil
}

func (l *List) SwitchTopic(newTopic string) error {
	loweredTopic := strings.ToLower(newTopic)
	if _, ok := l.Topics[loweredTopic]; !ok {
		return fmt.Errorf("%w: provided topic %s not in list of topics", ErrBadTopicName, loweredTopic)
	}

	l.CurrentTopic = loweredTopic

	return nil
}
