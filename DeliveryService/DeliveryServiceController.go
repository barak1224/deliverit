package DeliveryService

import (
    "log"
    "net/http"
)

func createDelivery(w http.ResponseWriter, r *http.Request) {
    log.Println("Delivery created!")
}

func updateDeliveryDestination(w http.ResponseWriter, r *http.Request) {
    log.Println("Delivery destination has been changed!")
}

func trackDelivery(w http.ResponseWriter, r *http.Request) {
    log.Println("Tracking shipment...")
}
