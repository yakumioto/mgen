package mgen

import (
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	"gopkg.in/urfave/cli.v2"
	"gopkg.in/yaml.v2"
)

func InterfaceAction(_ *cli.Context) error {
	t, err := template.New("interface").Parse(string(MustAsset("template/interface.tmpl")))
	if err != nil {
		log.Fatalf("[ERROR] parse template files error: %s\n", err)
	}

	fp, err := os.Create("model.mg.go")
	if err != nil {
		log.Fatalf("[ERROR] create model.go error: %s\n", err)
	}
	defer fp.Close()

	if err := t.Execute(fp, nil); err != nil {
		log.Fatalf("[ERROR] execute template error: %s\n", err)
	}

	return nil
}

func MgoAction(context *cli.Context) error {
	configPath := context.String("config-file")
	output := context.Bool("output")

	bytes, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatalf("[ERROR] read config file error: %s\n", err)
	}

	mg := new(ModelGenerator)
	mg.ConfigName = path.Base(configPath)
	if err := yaml.Unmarshal(bytes, mg); err != nil {
		log.Fatalf("[ERROR] unmarshal yaml error: %s\n", err)
	}

	t, err := template.New("interface").Funcs(template.FuncMap{
		"ToLower":     strings.ToLower,
		"SnakeString": SnakeString,
	}).Parse(string(MustAsset("template/mgo.tmpl")))
	if err != nil {
		log.Fatalf("[ERROR] parse template files error: %s\n", err)
	}

	if output {
		filename := strings.Replace(configPath, path.Ext(configPath), ".mg.go", 1)
		fp, err := os.Create(filename)
		if err != nil {
			log.Fatalf("[ERROR] create %s error: %s\n", filename, err)
		}
		defer fp.Close()

		if err := t.Execute(fp, mg); err != nil {
			log.Fatalf("[ERROR] execute template error: %s\n", err)
		}

		return nil
	}

	if err := t.Execute(os.Stdout, mg); err != nil {
		log.Fatalf("[ERROR] execute template error: %s\n", err)
	}

	return nil
}
