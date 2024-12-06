package v1

import (
	"context"
	"runwayclub.dev/aquila/pkg/store"
)

var storeData map[string]*store.StoreData = make(map[string]*store.StoreData)

var instance store.Controller

type controller struct{}

func GetController() store.Controller {
	if instance == nil {
		instance = &controller{}
	}
	return instance
}

func (i *controller) GetStore(ctx context.Context) map[string]*store.StoreData {
	return storeData
}

func (i *controller) Create(ctx context.Context, actor string, key string, value string) error {
	data := &store.StoreData{
		Actor: actor,
		Key:   key,
		Value: value,
	}
	// save to store as actor respective key
	storeData[actor+"/"+key] = data
	return nil
}

func (i *controller) List(ctx context.Context) (map[string]*store.StoreData, error) {
	return storeData, nil
}

func (i *controller) GetByActor(ctx context.Context, actor string) (map[string]*store.StoreData, error) {
	data := make(map[string]*store.StoreData)
	for k, v := range storeData {
		if v.Actor == actor {
			data[k] = v
		}
	}
	return data, nil
}

func (i *controller) GetByKey(ctx context.Context, actor string, key string) (string, error) {
	return storeData[actor+"/"+key].Value, nil
}

func (i *controller) Delete(ctx context.Context, actor string, key string) error {
	delete(storeData, actor+"/"+key)
	return nil
}
