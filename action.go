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

//go:generate go-bindata -o ./template.go -pkg mgen ./template
func InterfaceAction(context *cli.Context) error {
	packageName := context.String("package")

	t, err := template.New("interface").Parse(string(MustAsset("template/interface.tmpl")))
	if err != nil {
		log.Fatalf("[ERROR] parse template files error: %s\n", err)
	}

	filename := "model.mg.go"
	fp, err := os.Create(filename)
	if err != nil {
		log.Fatalf("[ERROR] create model.go error: %s\n", err)
	}
	defer fp.Close()

	mg := new(ModelGenerator)
	mg.PackageName = packageName
	mg.FileName = filename
	if err := t.Execute(fp, mg); err != nil {
		log.Fatalf("[ERROR] execute template error: %s\n", err)
	}

	return nil
}

func MgoAction(context *cli.Context) error {
	configPath := context.String("config-file")

	bytes, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatalf("[ERROR] read config file error: %s\n", err)
	}

	mg := new(ModelGenerator)
	mg.ConfigName = path.Base(configPath)
	if err := yaml.Unmarshal(bytes, mg); err != nil {
		log.Fatalf("[ERROR] unmarshal yaml error: %s\n", err)
	}

	t, err := template.New("mgo").Funcs(template.FuncMap{
		"ToLower":     strings.ToLower,
		"SnakeString": SnakeString,
	}).Parse(string(MustAsset("template/mgo.tmpl")))
	if err != nil {
		log.Fatalf("[ERROR] parse template files error: %s\n", err)
	}

	filename := strings.Replace(configPath, path.Ext(configPath), ".mg.go", 1)
	mg.FileName = filename
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
