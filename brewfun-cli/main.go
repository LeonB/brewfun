package main

import (
	"fmt"
	"os"

	"github.com/leonb/brewfun/brewfun-cli/db"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Version = "0.0.1"
	app.Commands = []cli.Command{
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "add a task to the list",
			Action: func(c *cli.Context) error {
				fmt.Println("added task: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "database",
			Aliases: []string{"db"},
			Usage:   "db actions",
			Subcommands: []cli.Command{
				{
					Name:  "migrate",
					Usage: "run migrations",
					Action: func(c *cli.Context) error {
						return db.Migrate()
					},
				},
				{
					Name:  "rollback",
					Usage: "rollback migrations",
					Action: func(c *cli.Context) error {
						fmt.Println("new task template: ", c.Args().First())
						return nil
					},
				},
			},
		},
	}

	app.Run(os.Args)
}
