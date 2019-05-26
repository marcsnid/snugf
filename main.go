package main

import (
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "snugf"
	app.Usage = "protect, encrypt, and compress files"
	app.ArgsUsage = "[input file] [output file]"

	operator := operator{}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "keyfile, K",
			Usage: "Load key from `FILE`",
			Destination: &operator.fileKey,
		},
		cli.StringFlag{
			Name:  "key, k",
			Usage: "Use a `string` as a key",
			Destination: &operator.stringKey,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "write",
			Aliases: []string{"w"},
			Usage:   "write a file",
			Action:  operator.write,
		},
		{
			Name:    "read",
			Aliases: []string{"r"},
			Usage:   "complete a task on the list",
			Action:  operator.read,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}