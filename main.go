package main

import (
	"atc/cmd"
	"github.com/urfave/cli"
	"os"
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
				contestnum := c.Args().Get(0)
				alphabet := c.Args().Get(1)
				cmd.MakeABCFiles(contestnum, alphabet, "abc")
				return nil
			},
		},
	}

	app.Run(os.Args)
}
