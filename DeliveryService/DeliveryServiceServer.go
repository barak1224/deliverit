package DeliveryService

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func setupRouter(router *mux.Router) {
	router.
		Methods("POST").
		Path("/delivery/create").
		HandlerFunc(createDelivery)
	router.
		Methods("GET").
		Path("/delivery/trackDelivery/").
		HandlerFunc(trackDelivery)
	router.
		Methods("POST").
		Path("/delivery/updateDeliveryDestination").
		HandlerFunc(updateDeliveryDestination)
}

func StartServer() {
	log.Println("In main")
	router := mux.NewRouter().StrictSlash(true)

	setupRouter(router)

	log.Fatal(http.ListenAndServe(":8080", router))
}