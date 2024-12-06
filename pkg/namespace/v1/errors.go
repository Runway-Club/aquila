package v1

import "errors"

var (
	ErrNamespaceAlreadyExists = errors.New("namespace already exists")
	ErrNamespaceNotFound      = errors.New("namespace not found")
	ErrCannotCreateNamespace  = errors.New("cannot create namespace")
	ErrCannotDeleteNamespace  = errors.New("cannot delete namespace")
)
