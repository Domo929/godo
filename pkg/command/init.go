package command

import (
	"fmt"

	"github.com/Domo929/godo/pkg/godo"
	"github.com/urfave/cli/v2"
)

func InitAction(c *cli.Context) error {
	l, err := godo.Load(c)
	if err != nil {
		return fmt.Errorf("loading: %w", err)
	}

	topicName := c.Args().First()

	if topicName == "" {
		return fmt.Errorf("%w: must not be empty", godo.ErrBadTopicName)
	}

	if _, ok := l.Topics[topicName]; ok {
		return fmt.Errorf("%w, topic name: %s", godo.ErrTopicAlreadyExists, topicName)
	}

	l.Topics[topicName] = make([]godo.Item, 0)

	return godo.Save(c, l)
}
