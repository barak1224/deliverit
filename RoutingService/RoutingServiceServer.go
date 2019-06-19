package RoutingService

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func setupRouter(router *mux.Router) {
	router.
		Methods("POST").
		Path("/routing/getPossibleRoutes").
		HandlerFunc(getPossibleRoutes)
}

func StartServer() {
	log.Println("In main")
	router := mux.NewRouter().StrictSlash(true)

	setupRouter(router)

	log.Fatal(http.ListenAndServe(":8080", router))
}
