package BillingService

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func setupRouter(router *mux.Router) {
	router.
		Methods("POST").
		Path("/users/deliveryFees").
		HandlerFunc(getDeliveryFees)
	router.
		Methods("POST").
		Path("/users/payment").
		HandlerFunc(setPayment)
}

func StartServer() {
	log.Println("In BillingService main")
	router := mux.NewRouter().StrictSlash(true)

	setupRouter(router)

	log.Fatal(http.ListenAndServe(":8084", router))
}
