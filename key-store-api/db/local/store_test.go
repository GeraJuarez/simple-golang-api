package db

import (
	"example/cloud-app/store/db"
	"testing"

	"errors"
)

func TestPut(t *testing.T) {
	store := NewKVStoreLocal()
	const key = "create-key"
	const value = "create-value"

	var val interface{}
	var contains bool

	defer delete(store.m, key)

	// Sanity check
	_, contains = store.m[key]
	if contains {
		t.Error("key/value already exists")
	}

	// err should be nil
	err := store.Put(key, value)
	if err != nil {
		t.Error(err)
	}

	val, contains = store.m[key]
	if !contains {
		t.Error("create failed")
	}

	if val != value {
		t.Error("val/value mismatch")
	}
}

func TestGet(t *testing.T) {
	store := NewKVStoreLocal()
	const key = "read-key"
	const value = "read-value"

	var val interface{}
	var err error

	defer delete(store.m, key)

	// Read a non-thing
	val, err = store.Get(key)
	if err == nil {
		t.Error("expected an error")
	}
	if !errors.Is(err, db.ErrorNoSuchKey) {
		t.Error("unexpected error:", err)
	}

	store.m[key] = value

	val, err = store.Get(key)
	if err != nil {
		t.Error("unexpected error:", err)
	}

	if val != value {
		t.Error("val/value mismatch")
	}
}

func TestDelete(t *testing.T) {
	store := NewKVStoreLocal()
	const key = "delete-key"
	const value = "delete-value"

	var contains bool

	defer delete(store.m, key)

	store.m[key] = value

	_, contains = store.m[key]
	if !contains {
		t.Error("key/value doesn't exist")
	}

	store.Delete(key)

	_, contains = store.m[key]
	if contains {
		t.Error("Delete failed")
	}
}
