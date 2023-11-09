package generator

import (
	"fmt"
	"regexp"

	"github.com/iancoleman/strcase"
)

type Label struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
}

func (l Label) GoVarName() string {
	return strcase.ToCamel(l.Name)
}

func (l Label) GoType() string {
	return l.Type
}

func (l Label) NodeVarName() string {
	return strcase.ToCamel(l.Name)
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

func (l Label) PythonVarName() string {
	return strcase.ToScreamingSnake(l.Name)
}

// Currently not used, kept for concistency
func (l Label) PythonType() string {
	return l.Type
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
	for team, teamLabels := range labels.Labels {
		if len(teamLabels) == 0 {
			return fmt.Errorf("labels are empty for team %s", team)
		}
		for _, label := range teamLabels {
			if label.Name == "" {
				return fmt.Errorf("label name is empty for team %s", team)
			}
			if !validLabelName.MatchString(label.Name) {
				return fmt.Errorf("label name %s is invalid for team %s", label.Name, team)
			}
			if label.Type == "" {
				return fmt.Errorf("label type is empty for team %s", team)
			}
			if _, ok := allLabels[label.Name]; ok {
				return fmt.Errorf("label name %s is already used in team %s", label.Name, team)
			}
			allLabels[label.Name] = team
		}
	}
	return nil
}
