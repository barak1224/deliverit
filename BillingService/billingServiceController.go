package BillingService

import (
	"log"
	"net/http"
)

func getDeliveryFees(w http.ResponseWriter, r *http.Request) {
	log.Println("Calculate the delivery fees, by using the Analytics Service and package details (i.e weight etc.).")
}

func setPayment(w http.ResponseWriter, r *http.Request) {
	log.Println("Send payment from customer to courier.")
}
