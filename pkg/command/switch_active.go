package command

import (
	"fmt"

	"github.com/Domo929/godo/pkg/godo"
	"github.com/urfave/cli/v2"
)

func SwitchAction(c *cli.Context) error {
	l, err := godo.Load(c)
	if err != nil {
		return fmt.Errorf("loading: %w", err)
	}

	newTopicName := c.Args().First()

	if err := l.SwitchTopic(newTopicName); err != nil {
		return err
	}

	return godo.Save(c, l)
}
