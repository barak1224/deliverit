package Analytics

import (
	"log"
	"net/http"
)

func getUserRate(w http.ResponseWriter, r *http.Request) {
	log.Println("This function get the user rate")
}

func getDeliveryTimeStatistics(w http.ResponseWriter, r *http.Request) {
	log.Println("This function will change get the specific delivery time statistics")
}

func getTariffDetermination(writer http.ResponseWriter, request *http.Request) {
	log.Println("This function will get the tariff fee for a delivery")
}

func insertNewFeedback(writer http.ResponseWriter, request *http.Request) {
	log.Println("This function will insert new feedback for user!")
}