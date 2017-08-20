package main

import (
	"os"

	"github.com/uncompiled/steganosaurus/modules"
	"gopkg.in/urfave/cli.v1"
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
			Usage:   "Tools for the whitespace programming language",
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
		{
			Name:    "zero-width",
			Aliases: []string{"zw"},
			Usage:   "Zero-width codepoint steganography",
			Subcommands: []cli.Command{
				{
					Name:  "encode",
					Usage: "Encodes plaintext as zero-width codepoints from input stream",
					Action: func(c *cli.Context) error {
						modules.ZeroWidthEncode(os.Stdin, os.Stdout)
						return nil
					},
				},
				{
					Name:  "decode",
					Usage: "Decodes hidden zero-width codepoints from input stream",
					Action: func(c *cli.Context) error {
						modules.ZeroWidthDecode(os.Stdin, os.Stdout)
						return nil
					},
				},
			},
		},
	}

	app.Run(os.Args)
}
