package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Domo929/godo/pkg/command"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "godo",
		Usage: "track your various todo lists!",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "location",
				Aliases: []string{"l"},
				Usage:   "`LOCATION` is where the godo.json file is stored, defaults to ~/.godo/godo.json",
				EnvVars: []string{"GODO_LOCATION"},
			},
		},
		Commands: []*cli.Command{
			{
				Name:   "init",
				Usage:  "create a new godo list with the given topic",
				Action: command.InitAction,
			},
			{
				Name:   "switch",
				Usage:  "switch the current active godo list. Must be an already existing list",
				Action: command.SwitchAction,
			},
		},

		Action: func(c *cli.Context) error {
			fmt.Println("started app")
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
