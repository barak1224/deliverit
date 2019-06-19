package main

import (
	"../deliverit/UserManagement"
	"./BillingService"
	"deliverit/Analytics"
)

func main() {
	UserManagement.StartServer()
	Analytics.StartServer()
	BillingService.StartServer()
}
