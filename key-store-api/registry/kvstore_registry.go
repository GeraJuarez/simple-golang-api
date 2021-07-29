package registry

import (
	"example/cloud-app/store/controller"
	"example/cloud-app/store/usecase/interactor"
	"example/cloud-app/store/usecase/repository"
)

func (r *registry) NewKVStoreController() controller.KVStoreController {
	return controller.NewKVStoreController(r.NewKVStoreInteractor())
}

func (r *registry) NewKVStoreInteractor() interactor.KVStoreInteractor {
	return interactor.NewKVStoreInteractor(r.NewKVStoreRepository())
}

func (r *registry) NewKVStoreRepository() repository.KeyValStoreRepository {
	return r.kvstore_repo
}

func (r *registry) NewKVStoreControllerV2() controller.KVStoreController {
	return controller.NewKVStoreControllerV2(r.NewKVStoreInteractor())
}
