package RoutineObserverService

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func setupRouter(router *mux.Router) {
	router.
		Methods("POST").
		Path("/routine/saveLocation").
		HandlerFunc(saveLocation)
	router.
		Methods("PUT").
		Path("/routine/setEnabled").
		HandlerFunc(setEnabled)
}

func StartServer() {
	log.Println("In main")
	router := mux.NewRouter().StrictSlash(true)

	setupRouter(router)

	log.Fatal(http.ListenAndServe(":8080", router))
}
