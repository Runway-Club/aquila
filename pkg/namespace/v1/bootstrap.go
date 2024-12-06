package v1

import (
	"context"
	"runwayclub.dev/aquila/pkg/namespace"
	"runwayclub.dev/aquila/pkg/store"
)

var storeController store.Controller

var instance namespace.Controller

type controller struct {
}

func GetController() namespace.Controller {
	if instance == nil {
		instance = &controller{}
	}
	return instance
}

func (c controller) Create(ctx context.Context, namespace string) error {
	err := storeController.Create(ctx, "namespace", namespace, "created")
	if err != nil {
		return ErrCannotCreateNamespace
	}
	return nil
}

func (c controller) Delete(ctx context.Context, namespace string) error {
	value, err := storeController.GetByKey(ctx, "namespace", namespace)
	if err != nil {
		return ErrNamespaceNotFound
	}
	if value != "created" {
		return ErrCannotDeleteNamespace
	}
	err = storeController.Delete(ctx, "namespace", namespace)
	if err != nil {
		return ErrCannotDeleteNamespace
	}
	return nil

}
