package v1

import (
	"context"
	"testing"
)

func TestBoostrap(t *testing.T) {
	controller := GetController()
	err := controller.Create(context.Background(), "actor", "key1", "value1")
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}
	data, err := controller.GetByActor(context.Background(), "actor")
	if err != nil {
		t.Errorf("GetByActor() failed: %v", err)
	}
	if len(data) != 1 {
		t.Errorf("GetByActor() returned %d items, expected 1", len(data))
	}
	if data["actor/key1"].Value != "value1" {
		t.Errorf("GetByActor() returned value %s, expected value1", data["key1"].Value)
	}

	// create another key
	err = controller.Create(context.Background(), "actor", "key2", "value2")
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	// get all keys
	data = controller.GetStore(context.Background())
	if err != nil {
		t.Errorf("List() failed: %v", err)
	}

	if len(data) != 2 {
		t.Errorf("List() returned %d items, expected 2", len(data))
	}

	// get by key
	value, err := controller.GetByKey(context.Background(), "actor", "key2")
	if err != nil {
		t.Errorf("GetByKey() failed: %v", err)
	}

	if value != "value2" {
		t.Errorf("GetByKey() returned value %s, expected value2", value)
	}

	// delete key
	err = controller.Delete(context.Background(), "actor", "key1")
	if err != nil {
		t.Errorf("Delete() failed: %v", err)
	}

}
