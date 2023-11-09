package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"path"
	"text/template"

	"github.com/rudderlabs/rudder-observability-kit/cmd/internal/generator"
	"gopkg.in/yaml.v3"
)

//go:embed labels.yaml
var labelsYAML []byte

func generateLabels(labels generator.Labels, templateFile string, outputDir string, extension string) {
	tmpl, err := template.New(path.Base(templateFile)).ParseFiles(templateFile)
	if err != nil {
		log.Fatal("error occurred while parsing template", err)
	}
	for team, teamLabels := range labels.Labels {
		fileName := fmt.Sprintf("%s/%s.%s", outputDir, team, extension)
		file, err := os.Create(fileName)
		if err != nil {
			log.Fatal("error occurred while creating file", err)
		}
		defer file.Close()
		if err := tmpl.Execute(file, struct {
			Team   string
			Labels []generator.Label
		}{
			Team:   team,
			Labels: teamLabels,
		}); err != nil {
			log.Fatal("error occurred while generating code", err)
		}
		log.Printf("generated %s", fileName)
	}
}

func main() {
	var labels generator.Labels
	if err := yaml.Unmarshal(labelsYAML, &labels); err != nil {
		log.Fatal("error occurred while parsing labels", err)
	}
	if err := labels.Validate(); err != nil {
		log.Fatal("error occurred while validating labels", err)
	}
	generateLabels(labels, "cmd/generate/templates/go.tmpl", "go/labels", "go")
	generateLabels(labels, "cmd/generate/templates/node.tmpl", "node/src/labels", "ts")
	generateLabels(labels, "cmd/generate/templates/python.tmpl", "python/labels", "py")
}


