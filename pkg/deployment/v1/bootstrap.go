package v1

import (
	"context"
	"runwayclub.dev/aquila/manifest"
	"runwayclub.dev/aquila/opts"
)

type ControllerV1 struct {
}

func (c ControllerV1) Create(ctx context.Context, manifest *manifest.Deployment) error {
	//TODO implement me
	panic("implement me")
}

func (c ControllerV1) List(ctx context.Context, filters *opts.DeploymentFilter) ([]*manifest.Deployment, error) {
	//TODO implement me
	panic("implement me")
}
