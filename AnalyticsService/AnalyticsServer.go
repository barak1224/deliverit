package AnalyticsService

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func setupRouter(router *mux.Router) {
	router.
		Methods("GET").
		Path("/analytics/getUserRate").
		HandlerFunc(getUserRate)
	router.
		Methods("GET").
		Path("/analytics/getDeliveryTimeStatistics").
		HandlerFunc(getDeliveryTimeStatistics)
	router.
		Methods("GET").
		Path("/analytics/tariffDetermination").
		HandlerFunc(getTariffDetermination)
	router.
		Methods("PUT").
		Path("/analytics/newFeedback").
		HandlerFunc(insertNewFeedback)
	router.
		Methods("POST").
		Path("/analytics/addShockData").
		HandlerFunc(addShockData)
}

func StartServer() {
	log.Println("In AnalyticsService main")
	router := mux.NewRouter().StrictSlash(true)

	setupRouter(router)

	log.Fatal(http.ListenAndServe(":8085", router))
}
