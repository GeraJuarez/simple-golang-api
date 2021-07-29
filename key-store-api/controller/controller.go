package controller

import (
	"example/cloud-app/store/usecase/interactor"
	"net/http"
)

type AppController struct {
	KVStore   KVStoreController
	KVStoreV2 KVStoreController
}

func New(kvsInteractor interactor.KVStoreInteractor) AppController {
	return AppController{
		KVStore: NewKVStoreController(kvsInteractor),
	}
}

type KVStoreController interface {
	GetValue(w http.ResponseWriter, r *http.Request)
	PutValue(w http.ResponseWriter, r *http.Request)
	DeleteValue(w http.ResponseWriter, r *http.Request)
}
