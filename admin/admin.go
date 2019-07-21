package admin

import (
 	"github.com/gorilla/mux"

	"gojudge/admin/users"
	"gojudge/admin/problems"
)

func InitAdminAPI(router *mux.Router) {
	router.HandleFunc("/users", users.ListUsersHandler).Methods("GET");
	router.HandleFunc("/createUser", users.CreateUserHandler).Methods("POST");
	router.HandleFunc("/deleteUser", users.DeleteUserHandler).Methods("POST");

	router.HandleFunc("/problems", problems.ListProblemsHandler).Methods("GET");
	router.HandleFunc("/createProblem", problems.CreateProblemHandler).Methods("POST");
	router.HandleFunc("/editProblem", problems.UpdateProblemHandler).Methods("POST");
	router.HandleFunc("/deleteProblem", problems.DeleteProblemHandler).Methods("POST");
}