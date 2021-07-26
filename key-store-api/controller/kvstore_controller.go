package controller

import (
	"errors"
	"example/cloud-app/store/usecase/interactor"
	"example/cloud-app/store/usecase/repository"
	"net/http"

	"github.com/gorilla/mux"
)

type kvStoreController struct {
	kvStoreInteractor interactor.KVStoreInteractor
}

type KVStoreController interface {
	GetValue(w http.ResponseWriter, r *http.Request)
	//PutValue()
	//DeleteValue()
}

func NewKVStoreController(interactor interactor.KVStoreInteractor) KVStoreController {
	return &kvStoreController{interactor}
}

func (c *kvStoreController) GetValue(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	value, err := c.kvStoreInteractor.Get(key)
	if errors.Is(err, repository.ErrorNoSuchKey) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(value))
}
