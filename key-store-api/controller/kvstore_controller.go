package controller

import (
	"errors"
	"example/cloud-app/store/usecase/interactor"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type kvStoreController struct {
	kvStoreInteractor interactor.KVStoreInteractor
}

type KVStoreController interface {
	GetValue(w http.ResponseWriter, r *http.Request)
	PutValue(w http.ResponseWriter, r *http.Request)
	DeleteValue(w http.ResponseWriter, r *http.Request)
}

func NewKVStoreController(interactor interactor.KVStoreInteractor) KVStoreController {
	return &kvStoreController{interactor}
}

func (c *kvStoreController) PutValue(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	value, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = c.kvStoreInteractor.Put(key, string(value))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (c *kvStoreController) GetValue(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	value, err := c.kvStoreInteractor.Get(key)
	if errors.Is(err, interactor.ErrorKeyNotFound) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(value))
}

func (c *kvStoreController) DeleteValue(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	err := c.kvStoreInteractor.Delete(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
