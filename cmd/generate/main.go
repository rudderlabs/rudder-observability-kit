package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"path"
	"regexp"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
	"gopkg.in/yaml.v3"

	"github.com/rudderlabs/rudder-observability-kit/cmd/internal/generator"
)

// https://github.com/golang/go/wiki/CodeReviewComments#initialisms
var (
	initialismsRegex *regexp.Regexp
	initialisms      = []string{
		"API", "ASCII", "CPU", "CSS", "DNS", "EOF", "GUID", "HTML", "HTTP", "HTTPS",
		"ID", "IP", "JSON", "QPS", "RAM", "RPC", "SSH", "TCP", "TLS", "TTL", "UDP",
		"UI", "UID", "UUID", "URI", "URL", "XML", "XSS",
	}
)

func init() {
	initialismsRegex = regexp.MustCompile(fmt.Sprintf(
		`(?i)\b(%s)\b`, strings.Join(initialisms, "|"),
	))
}

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("error occurred while getting current working directory: %v", err)
	}

	labelsPath := path.Join(cwd, "labels.yaml")
	labelsYAML, err := os.ReadFile(labelsPath)
	if err != nil {
		log.Fatalf("error occurred while reading %q: %v", labelsPath, err)
	}

	var labels generator.Labels
	if err := yaml.Unmarshal(labelsYAML, &labels); err != nil {
		log.Fatalf("error occurred while parsing labels: %v", err)
	}
	if err := labels.Validate(); err != nil {
		log.Fatalf("error occurred while validating labels: %v", err)
	}
	generateLabels(labels, "cmd/generate/templates/go.tmpl", "go/labels", "go")
	generateLabels(labels, "cmd/generate/templates/node.tmpl", "node/src/labels", "ts")
	generateLabels(labels, "cmd/generate/templates/python.tmpl", "python/labels", "py")
}

func generateLabels(labels generator.Labels, templateFile, outputDir, extension string) {
	funcMap := template.FuncMap{
		"ToCamel":                strcase.ToCamel,
		"ToCamelWithInitialisms": toCamelWithInitialisms,
		"ToScreamingSnake":       strcase.ToScreamingSnake,
	}
	tmpl, err := template.New(path.Base(templateFile)).Funcs(funcMap).ParseFiles(templateFile)
	if err != nil {
		log.Fatalf("error occurred while parsing template [%q]: %v", path.Base(templateFile), err)
	}

	for domain, domainLabels := range labels.Labels {
		filename := fmt.Sprintf("%s/%s.%s", outputDir, domain, extension)
		if err := createFile(tmpl, domain, filename, domainLabels); err != nil {
			log.Fatalf("error occurred while generating file %q: %v", filename, err)
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

// https://github.com/golang/go/wiki/CodeReviewComments#initialisms
func toCamelWithInitialisms(s string) string {
	s = strcase.ToDelimited(s, ' ')

	a := strings.Split(s, " ")
	for i, w := range a {
		if initialismsRegex.MatchString(w) {
			a[i] = strings.ToUpper(w)
		} else {
			a[i] = strcase.ToCamel(w)
		}
	}

	return strings.Join(a, "")
}
