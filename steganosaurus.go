package main

import (
	"github.com/uncompiled/steganosaurus/modules"
	"gopkg.in/urfave/cli.v1"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "steganosaurus"
	app.Usage = "a tool to hide stuff inside other stuff"
	app.Version = "1.0.0"

	app.Commands = []cli.Command{
		{
			Name:    "whitespace",
			Aliases: []string{"ws"},
			Usage:   "Tools for whitespace steganography",
			Subcommands: []cli.Command{
				{
					Name:  "merge",
					Usage: "Merges visible code with whitespace code",
					Action: func(c *cli.Context) error {
						filename1 := c.Args().Get(0)
						filename2 := c.Args().Get(1)

						modules.WhitespaceMerge(filename1, filename2, os.Stdout)
						return nil
					},
				},
			},
		},
	}

	app.Run(os.Args)
}
