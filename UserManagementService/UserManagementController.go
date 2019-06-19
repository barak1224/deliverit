package UserManagementService

import (
	"log"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	log.Println("This function will make a login")
}

func register(w http.ResponseWriter, r *http.Request) {
	log.Println("This function will make a registration")
}

func changePassword(w http.ResponseWriter, r *http.Request) {
	log.Println("This function will change the password")
}

func changeUserName(w http.ResponseWriter, r *http.Request) {
	log.Println("This function will change user name")
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	log.Println("This function will delete user!")
}
