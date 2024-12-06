package deployment

import (
	"context"
	"runwayclub.dev/aquila/manifest"
	"runwayclub.dev/aquila/opts"
)

type Controller interface {
	Create(ctx context.Context, manifest *manifest.Deployment) error
	List(ctx context.Context, filters *opts.DeploymentFilter) ([]*manifest.Deployment, error)
}
