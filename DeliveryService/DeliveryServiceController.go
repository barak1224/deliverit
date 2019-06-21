package DeliveryService

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

// these structs need to be in a common API so that they may be shared among multiple project (each microservice
// is a separate project).

type DeliveryMessage struct {
	Source      string
	Destination string
	Weight      float64
	// more fields need to be added
}

type RouteMessage struct {
	Source      string
	Destination string
}

func createDelivery(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var reqMsg DeliveryMessage
	err := decoder.Decode(&reqMsg)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	log.Println("DeliveryService got", reqMsg)

	var routingServiceMessage RouteMessage
	routingServiceMessage.Source = reqMsg.Source
	routingServiceMessage.Destination = reqMsg.Destination

	b := new(bytes.Buffer)
	err = json.NewEncoder(b).Encode(routingServiceMessage)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	_, err = http.Post("http://localhost:8081/routing/getPossibleRoutes", "application/json; charset=utf-8", b)
	log.Println("Delivery created!")
}

func updateDeliveryDestination(w http.ResponseWriter, r *http.Request) {
	log.Println("Delivery destination has been changed!")
}

func trackDelivery(w http.ResponseWriter, r *http.Request) {
	log.Println("Tracking shipment...")
}
