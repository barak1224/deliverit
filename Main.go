package main

import (
	"../deliverit/Analytics"
	"../deliverit/UserManagement"
)

func main() {
	UserManagement.StartServer()
	Analytics.StartServer()
}
