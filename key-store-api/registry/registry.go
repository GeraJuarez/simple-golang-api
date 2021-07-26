package registry

import (
	"example/cloud-app/store/controller"
	"example/cloud-app/store/usecase/repository"
)

type registry struct {
	kvstore_repo repository.KeyValStoreRepository
}

type Registry interface {
	NewAppController() controller.AppController
}

func NewRegistry(kvstore_repo repository.KeyValStoreRepository) Registry {
	return &registry{kvstore_repo}
}

func (r *registry) NewAppController() controller.AppController {
	return controller.AppController{
		KVStore: r.NewKVStoreController(),
	}
}
