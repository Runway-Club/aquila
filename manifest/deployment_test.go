package manifest_test

import (
	"runwayclub.dev/aquila/manifest"
	"testing"
)

func TestParseDeployment(t *testing.T) {
	yamlStr := `
		id: my-deployment
		name: My Deployment
		images:
			- alpine
			- hello-world
		`
	deployment, err := manifest.ParseDeployment(yamlStr)
	if err != nil {
		t.Fatalf("Failed to parse deployment: %v", err)
	}

}
