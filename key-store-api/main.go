package main

import (
	"example/cloud-app/store/controller"
	"example/cloud-app/store/router"
	"example/cloud-app/store/usecase/interactor"
	repository "example/cloud-app/store/usecase/repository/local"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// func keyValuePutHandler(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	key := vars["key"]

// 	value, err := ioutil.ReadAll(r.Body)
// 	defer r.Body.Close()

// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	err = Put(key, string(value))
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusCreated)

// }

// func keyValueDeleteHandler(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	key := vars["key"]

// 	err := Delete(key)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// }

var PORT = "PORT"

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	portEnv := os.Getenv(PORT)

	datastore := repository.NewKVStoreLocal()
	r := router.Start(controller.New(interactor.NewKVStoreInteractor(datastore)))

	//r.HandleFunc("/v1/{key}", keyValuePutHandler).Methods("PUT")
	//r.HandleFunc("/v1/{key}", keyValueDeleteHandler).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":"+portEnv, r))
}
