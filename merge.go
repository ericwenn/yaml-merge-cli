package main // import "github.com/ericwenn/yaml-merge-cli"

import (
	"fmt"
	"github.com/imdario/mergo"
	"gopkg.in/alecthomas/kingpin.v2"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

func main() {
	app := kingpin.New("yaml-merge-cli", "merge yaml files using mergo")
	files := app.Arg("files", "Files to merge, from left to right").Required().ExistingFiles()
	_ = kingpin.MustParse(app.Parse(os.Args[1:]))

	merged := make(map[interface{}]interface{})
	for i, file := range *files {
		b, err := ioutil.ReadFile(file)
		if err != nil {
			app.Fatalf("unable to read file %d: %v", i, err)
		}
		m := make(map[interface{}]interface{})
		if err := yaml.Unmarshal(b, &m); err != nil {
			app.Fatalf("unable to unmarshal file %d: %v", i, err)
		}
		if err := mergo.Merge(&merged, m, mergo.WithOverride, mergo.WithAppendSlice); err != nil {
			app.Fatalf("unable to merge file %d: %v", i, err)
		}
	}
	out, err := yaml.Marshal(merged)
	if err != nil {
		app.Fatalf("unable to marshall output: %v", err)
	}
	fmt.Print(string(out))
}