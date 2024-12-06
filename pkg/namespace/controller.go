package namespace

import "context"

type Controller interface {
	Create(ctx context.Context, namespace string) error
	Delete(ctx context.Context, namespace string) error
}
