package controller

import (
	"errors"
	"example/cloud-app/store/usecase/interactor"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type kvStoreControllerV2 struct {
	kvStoreInteractor interactor.KVStoreInteractor
}

func NewKVStoreControllerV2(interactor interactor.KVStoreInteractor) KVStoreController {
	return &kvStoreControllerV2{interactor}
}

func (c *kvStoreControllerV2) PutValue(w http.ResponseWriter, r *http.Request) {
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

func (c *kvStoreControllerV2) GetValue(w http.ResponseWriter, r *http.Request) {
	// readjson
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

func (c *kvStoreControllerV2) DeleteValue(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	err := c.kvStoreInteractor.Delete(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
