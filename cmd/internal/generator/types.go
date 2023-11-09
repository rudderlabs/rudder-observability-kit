package generator

import (
	"fmt"
	"regexp"
)

type Label struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
}

func (l Label) GoType() string {
	switch l.Type {
	case "string":
		return "string"
	case "int", "int64", "int32":
		return "int64"
	case "bool":
		return "bool"
	case "float64", "float32":
		return "float64"
	case "time":
		return "time.Time"
	case "duration":
		return "time.Duration"
	case "error":
		return "error"
	default:
		panic(fmt.Sprintf("unsupported type %s", l.Type))
	}
}

func (l Label) NodeType() string {
	switch l.Type {
	case "string":
		return "string"
	case "int", "int64", "int32", "float64", "float32":
		return "number"
	case "bool":
		return "boolean"
	case "Time":
		return "Date"
	default:
		panic(fmt.Sprintf("unsupported type %s", l.Type))
	}
}

type Labels struct {
	Labels map[string][]Label `yaml:"labels"`
}

var validLabelName = regexp.MustCompile(`\b[a-zA-Z][a-zA-Z0-9]*\b`)

func (labels Labels) Validate() error {
	if len(labels.Labels) == 0 {
		return fmt.Errorf("labels should not be empty")
	}

	allLabels := map[string]string{}
	for domain, domainLabels := range labels.Labels {
		if len(domainLabels) == 0 {
			return fmt.Errorf("labels are empty for domain %s", domain)
		}
		for _, label := range domainLabels {
			if label.Name == "" {
				return fmt.Errorf("label name is empty for domain %s", domain)
			}
			if !validLabelName.MatchString(label.Name) {
				return fmt.Errorf("label name %s is invalid for domain %s", label.Name, domain)
			}
			if label.Type == "" {
				return fmt.Errorf("label type is empty for domain %s", domain)
			}
			if _, ok := allLabels[label.Name]; ok {
				return fmt.Errorf("label name %s is already used in domain %s", label.Name, domain)
			}
			allLabels[label.Name] = domain
		}
	}
	return nil
}
