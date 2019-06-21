package main

import (
	"./AnalyticsService"
	"./BillingService"
	"./DeliveryService"
	"./RoutineObserverService"
	"./RoutingService"
	"./UserManagementService"
)

func main() {
	c := make(chan string)
	go UserManagementService.StartServer()
	go BillingService.StartServer()
	go AnalyticsService.StartServer()
	go DeliveryService.StartServer()
	go RoutingService.StartServer()
	go RoutineObserverService.StartServer()

	// this is to make it wait indefinitely (only for debugging purposes, they need to be separate processes,
	// on separate machines)
	<-c
}
