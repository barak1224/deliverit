package main

import (
	"./AnalyticsService"
	"./RoutineObserverService"
	"./RoutingService"
	"./UserManagementService"
)

func main() {
	UserManagementService.StartServer()
	AnalyticsService.StartServer()
	RoutingService.StartServer()
	RoutineObserverService.StartServer()
}
