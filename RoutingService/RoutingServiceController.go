package RoutingService

import (
	"encoding/json"
	"log"
	"net/http"
)

type RouteMessage struct {
	Source      string
	Destination string
}

func getPossibleRoutes(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting possible routes...")

	decoder := json.NewDecoder(r.Body)

	var reqMsg RouteMessage
	err := decoder.Decode(&reqMsg)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	log.Println("RoutingService got", reqMsg)

	// next step is to initiate our search algorithm and return a response of the possible routes to the RoutingService
}
