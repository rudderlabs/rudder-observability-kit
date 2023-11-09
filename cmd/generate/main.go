package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"path"
	"text/template"

	"gopkg.in/yaml.v3"

	"github.com/rudderlabs/rudder-observability-kit/cmd/internal/generator"
)

//go:embed labels.yaml
var labelsYAML []byte

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

func generateLabels(labels generator.Labels, templateFile, outputDir, extension string) {
	tmpl, err := template.New(path.Base(templateFile)).ParseFiles(templateFile)
	if err != nil {
		log.Fatal("error occurred while parsing template", err)
	}

	for domain, domainLabels := range labels.Labels {
		filename := fmt.Sprintf("%s/%s.%s", outputDir, domain, extension)
		if err := createFile(tmpl, domain, filename, domainLabels); err != nil {
			log.Fatalf("error occurred while generating file: %v", err)
		}
		log.Printf("generated %s", filename)
	}
}

func createFile(
	tmpl *template.Template, domain, fileName string, labels []generator.Label,
) error {
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("error occurred while creating file for %q: %v", domain, err)
	}

	defer func() { _ = file.Close() }()

	if err := tmpl.Execute(file, struct {
		Domain string
		Labels []generator.Label
	}{
		Domain: domain,
		Labels: labels,
	}); err != nil {
		return fmt.Errorf("error occurred while executing template for %q: %v", domain, err)
	}

	return nil
}
