package Analytics

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
}

func StartServer() {
	log.Println("In main")
	router := mux.NewRouter().StrictSlash(true)

	setupRouter(router)

	log.Fatal(http.ListenAndServe(":8080", router))
}