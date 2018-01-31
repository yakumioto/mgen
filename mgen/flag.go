package main

import "gopkg.in/urfave/cli.v2"

func defaultInterfaceFlag() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    "package",
			Aliases: []string{"p"},
			Usage:   "set interface file package name",
		},
	}
}

func defaultModelFlag() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    "config-file",
			Aliases: []string{"c"},
			Usage:   "set the config file path",
		},
	}
}
