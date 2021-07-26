package router

import (
	"example/cloud-app/store/controller"

	"github.com/gorilla/mux"
)

func Start(c controller.AppController) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	api := router.PathPrefix("/api").Subrouter()
	kvstore_v1 := api.PathPrefix("/v1/kvstore").Subrouter()

	kvstore_v1.HandleFunc("/{key}", c.KVStore.GetValue).Methods("GET")

	return router
}
