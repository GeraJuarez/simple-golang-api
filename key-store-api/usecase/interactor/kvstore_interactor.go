package interactor

import "example/cloud-app/store/usecase/repository"

type kvStoreInteractor struct {
	data repository.KeyValStoreRepository
}

type KVStoreInteractor interface {
	Get(key string) (string, error)
}

func NewKVStoreInteractor(r repository.KeyValStoreRepository) KVStoreInteractor {
	return &kvStoreInteractor{r}
}

func (kvstore *kvStoreInteractor) Get(key string) (string, error) {
	val, err := kvstore.data.Get(key)
	if err != nil {
		return "", err
	}

	return val, nil
}
