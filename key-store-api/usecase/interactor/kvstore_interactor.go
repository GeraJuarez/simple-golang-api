package interactor

import (
	"errors"
	"example/cloud-app/store/usecase/repository"
)

var ErrorKeyNotFound = errors.New("key not found")

type kvStoreInteractor struct {
	repo repository.KeyValStoreRepository
}

type KVStoreInteractor interface {
	Put(key string, value string) error
	Get(key string) (string, error)
	Delete(key string) error
	GetAllVal() ([]string, error)
	// NOTE:
	// Methods named same as the controller
	// Business rules are appled in this layer
}

func NewKVStoreInteractor(r repository.KeyValStoreRepository) KVStoreInteractor {
	return &kvStoreInteractor{r}
}

func (kvstore *kvStoreInteractor) Put(key string, value string) error {
	err := kvstore.repo.Put(key, value)

	if err != nil {
		// Do more stuff
		return err
	}

	return nil
}

func (kvstore *kvStoreInteractor) Get(key string) (string, error) {
	val, err := kvstore.repo.Get(key)

	if errors.Is(err, repository.ErrorNoSuchKey) {
		return "", ErrorKeyNotFound
	}
	if err != nil {
		return "", err
	}

	return val, nil
}

func (kvstore *kvStoreInteractor) Delete(key string) error {
	err := kvstore.repo.Delete(key)

	return err
}

func (kvstore *kvStoreInteractor) GetAllVal() ([]string, error) {
	chanVal, err := kvstore.repo.FindAllValues()
	values := []string{}

	// maybe is better to apply here the Fan-in pattern here (Funnel function)
	for val := range chanVal {
		values = append(values, val)
	}

	if errors.Is(err, repository.ErrorNoSuchKey) {
		return nil, ErrorKeyNotFound
	}
	if err != nil {
		return nil, err
	}

	return values, nil
}
