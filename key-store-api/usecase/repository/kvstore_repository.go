package repository

import "errors"

var ErrorNoSuchKey = errors.New("no such key")

type KeyValStoreRepository interface {
	Put(key string, value string) error
	Get(key string) (string, error)
	Delete(key string) error
	FindAllValues() (<-chan string, error)
	// NOTE:
	// These methods can be named diferent acording to how the database behave, ex:
	// FindAll
	// FindById
}
