package controller

import "example/cloud-app/store/interactor"

type AppController struct {
	KVStore KVStoreController
}

func New(kvsInteractor interactor.KVStoreInteractor) AppController {
	return AppController{
		KVStore: NewKVStoreController(kvsInteractor),
	}
}