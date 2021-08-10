package controller

import (
	"encoding/json"
	"errors"
	"example/cloud-app/store/model"
	"example/cloud-app/store/usecase/interactor"
	"net/http"
)

type kvStoreControllerV2 struct {
	kvStoreInteractor interactor.KVStoreInteractor
	// TODO: add input validator
}

func NewKVStoreControllerV2(interactor interactor.KVStoreInteractor) KVStoreController {
	return &kvStoreControllerV2{interactor}
}

func (c *kvStoreControllerV2) PutValue(w http.ResponseWriter, r *http.Request) {
	var jsonKeyVal model.KeyVal
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	var unmarshalTypeErr *json.UnmarshalTypeError
	if err := decoder.Decode(&jsonKeyVal); err != nil {
		if errors.As(err, &unmarshalTypeErr) {
			http.Error(w, "Bad Request. Wrong Type provided for field "+unmarshalTypeErr.Field, http.StatusBadRequest)
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return
	}

	key := jsonKeyVal.Key
	value := jsonKeyVal.Value
	err := c.kvStoreInteractor.Put(key, string(value))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	respValKey := &model.KeyVal{Key: key, Value: value}
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	if err := enc.Encode(respValKey); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (c *kvStoreControllerV2) GetValue(w http.ResponseWriter, r *http.Request) {
	var jsonKey model.Key
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	var unmarshalTypeErr *json.UnmarshalTypeError
	if err := decoder.Decode(&jsonKey); err != nil {
		if errors.As(err, &unmarshalTypeErr) {
			http.Error(w, "Bad Request. Wrong Type provided for field "+unmarshalTypeErr.Field, http.StatusBadRequest)
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return
	}

	key := jsonKey.Key
	value, err := c.kvStoreInteractor.Get(key)
	if errors.Is(err, interactor.ErrorKeyNotFound) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respValKey := &model.KeyVal{Key: key, Value: value}
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	if err := enc.Encode(respValKey); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (c *kvStoreControllerV2) DeleteValue(w http.ResponseWriter, r *http.Request) {
	var jsonKey model.Key
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	var unmarshalTypeErr *json.UnmarshalTypeError
	if err := decoder.Decode(&jsonKey); err != nil {
		if errors.As(err, &unmarshalTypeErr) {
			http.Error(w, "Bad Request. Wrong Type provided for field "+unmarshalTypeErr.Field, http.StatusBadRequest)
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return
	}

	key := jsonKey.Key
	err := c.kvStoreInteractor.Delete(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (c *kvStoreControllerV2) GetValues(w http.ResponseWriter, r *http.Request) {
	values, err := c.kvStoreInteractor.GetAllVal()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respValKey := &model.Values{Array: values}
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	if err := enc.Encode(respValKey); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
