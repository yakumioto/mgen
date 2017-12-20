package main

import "gopkg.in/urfave/cli.v2"

func defaultModelFlag() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    "config-file",
			Aliases: []string{"c"},
			Usage:   "set the config file path",
		},
		&cli.BoolFlag{
			Name:    "output",
			Aliases: []string{"o"},
			Usage:   "set the output flag",
		},
	}
}
