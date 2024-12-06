package pod

import (
	"context"
	"runwayclub.dev/aquila/manifest"
)

type Controller interface {
	Create(ctx context.Context, manifest *manifest.Pod) error
	List(ctx context.Context, filter *opts.) (map[string]*manifest.Pod, error)
	GetInfo(ctx context.Context, name string) (*manifest.Pod, error)
}
