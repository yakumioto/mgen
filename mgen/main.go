package main

import (
	"os"

	"github.com/yakumioto/mgen"
	"gopkg.in/urfave/cli.v2"
)

var version string

//go:generate go-bindata -o ../template.go -pkg mgen ../template
func main() {
	app := &cli.App{
		Name:  "mgen",
		Usage: "code generate for mgo",
		Commands: []*cli.Command{
			{Name: "interface", Usage: "craete model interface go file", Action: mgen.InterfaceAction},
			{Name: "mgo", Usage: "generate golang code", Flags: defaultModelFlag(), Action: mgen.MgoAction},
		},
		Version: version,
	}

	app.Run(os.Args)
}
