package AnalyticsService

import (
	"log"
	"net/http"
)

func getUserRate(w http.ResponseWriter, r *http.Request) {
	log.Println("This function will get the user rate.")
}

func getDeliveryTimeStatistics(w http.ResponseWriter, r *http.Request) {
	log.Println("This function will get the specific delivery time statistics.")
}

func getTariffDetermination(writer http.ResponseWriter, request *http.Request) {
	log.Println("This function will get the tariff fee for a delivery.")
}

func insertNewFeedback(writer http.ResponseWriter, request *http.Request) {
	log.Println("This function will insert new feedback for user.")
}

func addShockData(writer http.ResponseWriter, request *http.Request) {
	log.Println("This function will add shock data about the package.")
}
