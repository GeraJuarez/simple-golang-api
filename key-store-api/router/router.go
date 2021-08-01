package router

import (
	"example/cloud-app/store/controller"

	"github.com/gorilla/mux"
)

func Start(c controller.AppController) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	api := router.PathPrefix("/api").Subrouter()

	addKVStoreRouter(c, api)
	addKVStoreRouterV2(c, api)

	return router
}

func addKVStoreRouter(c controller.AppController, api *mux.Router) {
	kvstore := api.PathPrefix("/v1/kvstore").Subrouter()

	kvstore.HandleFunc("/{key}", c.KVStore.PutValue).Methods("PUT")
	kvstore.HandleFunc("/{key}", c.KVStore.GetValue).Methods("GET")
	kvstore.HandleFunc("/{key}", c.KVStore.DeleteValue).Methods("DELETE")
}

func addKVStoreRouterV2(c controller.AppController, api *mux.Router) {
	kvstore_v2 := api.PathPrefix("/v2/kvstore").Subrouter()

	kvstore_v2.HandleFunc("", c.KVStoreV2.PutValue).Methods("PUT")
	kvstore_v2.HandleFunc("", c.KVStoreV2.GetValue).Methods("GET")
	kvstore_v2.HandleFunc("", c.KVStoreV2.DeleteValue).Methods("DELETE")
}

// Midleware example
// func HomeHandler(response http.ResponseWriter, request *http.Request) {

//     fmt.Fprintf(response, "Hello home")
// }

// func Middleware(h http.Handler) http.Handler {
//     return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//         log.Println("middleware", r.URL)
//         h.ServeHTTP(w, r)
//     })
// }
// func main() {
//     r := mux.NewRouter()
//     r.HandleFunc("/", HomeHandler)
//     http.Handle("/", Middleware(r))
// }
