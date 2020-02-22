package main

import (
	"os"

	"github.com/urfave/cli"
	"github.com/username/qi/cmd"
)

func main() {
	app := cli.NewApp()
	app.Name = "atcoder abc make files"
	app.Version = "0.1.0"

	app.Commands = []cli.Command{
		{
			Name:  "abc",
			Usage: "atcoder + mine : you get cpp files for atcoder abc",
			Action: func(c *cli.Context) error {
				cmd.MakeFiles()
				return nil
			},
		},
	}

	app.Run(os.Args)
}
