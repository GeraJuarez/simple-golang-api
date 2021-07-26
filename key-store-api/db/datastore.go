package db

import "errors"

var ErrorNoSuchKey = errors.New("no such key")

type KeyValStoreRepository interface {
	Get(key string) (string, error)
}
