package manifest_test

import (
	"runwayclub.dev/aquila/manifest"
	"testing"
)

func TestParseDeployment(t *testing.T) {
	yamlStr :=
		`id: my-deployment
name: My Deployment
images:
- alpine
- hello-world`
	deployment, err := manifest.ParseDeployment(yamlStr)
	if err != nil {
		t.Fatalf("Failed to parse deployment: %v", err)
	}

	if deployment.Id != "my-deployment" {
		t.Errorf("Expected deployment ID to be 'my-deployment', got '%s'", deployment.Id)
	}

	if deployment.Name != "My Deployment" {
		t.Errorf("Expected deployment name to be 'My Deployment', got '%s'", deployment.Name)
	}

	if len(deployment.Images) != 2 {
		t.Errorf("Expected 2 images, got %d", len(deployment.Images))
	}

	if deployment.Images[0] != "alpine" {
		t.Errorf("Expected first image to be 'alpine', got '%s'", deployment.Images[0])
	}

	if deployment.Images[1] != "hello-world" {
		t.Errorf("Expected second image to be 'hello-world', got '%s'", deployment.Images[1])
	}

}
