package UserManagementService

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func setupRouter(router *mux.Router) {
	router.
		Methods("POST").
		Path("/users/login").
		HandlerFunc(login)
	router.
		Methods("POST").
		Path("/users/register").
		HandlerFunc(register)
	router.
		Methods("PUT").
		Path("/users/changePassword").
		HandlerFunc(changePassword)
	router.Methods("PUT").
		Path("/users/changeUserName").
		HandlerFunc(changeUserName)
	router.Methods("DELETE").
		Path("/deleteUser").
		HandlerFunc(deleteUser)
}

func StartServer() {
	log.Println("In UserManagement main")
	router := mux.NewRouter().StrictSlash(true)

	setupRouter(router)

	log.Fatal(http.ListenAndServe(":8080", router))
}
