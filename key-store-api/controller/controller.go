package controller

import (
	"encoding/json"
	"example/cloud-app/store/usecase/interactor"
	"mime"
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
	GetValues(w http.ResponseWriter, r *http.Request)
}

func (app *AppController) EnforceJSONHandler(next http.Handler) http.Handler {
	// TODO:
	// should this be in a separate package?
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get("Content-Type")

		if contentType != "" {
			mt, _, err := mime.ParseMediaType(contentType)
			if err != nil {
				http.Error(w, "Malformed Content-Type header", http.StatusBadRequest)
				return
			}

			if mt != "application/json" {
				http.Error(w, "Content-Type header must be application/json", http.StatusUnsupportedMediaType)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}

func JSONError(w http.ResponseWriter, message string, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	resp := make(map[string]string)
	resp["message"] = message
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}
