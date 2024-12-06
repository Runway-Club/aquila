package store

import "context"

type StoreData struct {
	Actor string `json:"actor"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Controller interface {
	GetStore(ctx context.Context) map[string]*StoreData
	Create(ctx context.Context, actor string, key string, value string) error
	GetByActor(ctx context.Context, actor string) (map[string]*StoreData, error)
	GetByKey(ctx context.Context, actor string, key string) (string, error)
	Delete(ctx context.Context, actor string, key string) error
}
