package main

import (
	"os"

	"github.com/yakumioto/mgen"
	"gopkg.in/urfave/cli.v2"
)

func main() {
	app := &cli.App{
		Name:  "mgen",
		Usage: "code generate for mgo",
		Commands: []*cli.Command{
			{Name: "interface", Usage: "create model interface go file", Flags: defaultInterfaceFlag(), Action: mgen.InterfaceAction},
			{Name: "mgo", Usage: "generate golang code", Flags: defaultModelFlag(), Action: mgen.MgoAction},
		},
		Version: "0.1.0",
	}

	app.Run(os.Args)
}
