package main

import (
	"github.com/urfave/cli"
	"log"
	"os"
	"sort"
)

func main() {
	app := cli.NewApp()
	app.Name = "snugf"
	app.Usage = "protect, encrypt, and compress files"
	app.ArgsUsage = "[input file] [output file]"

	o := operator{fileKey: new(string), stringKey: new(string)}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "keyfile, K",
			Usage: "Load key from `FILE`",
			Destination: o.fileKey,
		},
		cli.StringFlag{
			Name:  "key, k",
			Usage: "Use a `string` as a key",
			Destination: o.stringKey,
		},
	}


	app.Commands = []cli.Command{
		{
			Name:    "write",
			Aliases: []string{"w"},
			Usage:   "write a file",
			Action:  o.write,
			Flags: app.Flags,
		},
		{
			Name:    "read",
			Aliases: []string{"r"},
			Usage:   "read a file",
			Action:  o.read,
			Flags: app.Flags,
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}