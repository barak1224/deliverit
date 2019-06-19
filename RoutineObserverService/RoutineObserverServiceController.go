package RoutineObserverService

import (
	"log"
	"net/http"
)

func saveLocation(w http.ResponseWriter, r *http.Request) {
	log.Println("Location saved")
}

func setEnabled(w http.ResponseWriter, r *http.Request) {
	log.Println("Setting has been changed")
}
