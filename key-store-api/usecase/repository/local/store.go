package repository

import (
	"example/cloud-app/store/usecase/repository"
	"sync"
)

type kvStoreLocal struct {
	sync.RWMutex
	m map[string]string
}

func NewKVStoreLocal() repository.KeyValStoreRepository {
	return &kvStoreLocal{m: make(map[string]string)}
}

func (store *kvStoreLocal) Put(key string, value string) error {
	store.Lock()
	store.m[key] = value
	store.Unlock()

	return nil
}

func (store *kvStoreLocal) Get(key string) (string, error) {
	store.RLock()
	value, ok := store.m[key]
	store.RUnlock()

	if !ok {
		return "", repository.ErrorNoSuchKey
	}

	return value, nil
}

func (store *kvStoreLocal) Delete(key string) error {
	store.Lock()
	delete(store.m, key)
	store.Unlock()

	return nil
}
