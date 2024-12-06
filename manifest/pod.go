package manifest

type Pod struct {
	Id           string   `yaml:"id"`
	DeploymentId string   `yaml:"deployment_id"`
	Name         string   `yaml:"name"`
	Images       []string `yaml:"images"`
}
