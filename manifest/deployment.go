package manifest

import (
	"gopkg.in/yaml.v3"
)

type Deployment struct {
	Id     string   `yaml:"id"`
	Name   string   `yaml:"name"`
	Images []string `yaml:"images"`
}

func ParseDeployment(yamlStr string) (*Deployment, error) {
	deployment := &Deployment{}
	err := yaml.Unmarshal([]byte(yamlStr), deployment)
	if err != nil {
		return nil, err
	}
	return deployment, nil
}
