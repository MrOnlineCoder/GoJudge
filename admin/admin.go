package admin

import (
 	"github.com/gorilla/mux"

	"gojudge/admin/users"
	"gojudge/admin/problems"
	"gojudge/admin/tests"
	"gojudge/admin/contest"
)

func InitAdminAPI(router *mux.Router) {
	router.HandleFunc("/users", users.ListUsersHandler).Methods("GET");
	router.HandleFunc("/createUser", users.CreateUserHandler).Methods("POST");
	router.HandleFunc("/deleteUser", users.DeleteUserHandler).Methods("POST");

	router.HandleFunc("/problems", problems.ListProblemsHandler).Methods("GET");
	router.HandleFunc("/createProblem", problems.CreateProblemHandler).Methods("POST");
	router.HandleFunc("/editProblem", problems.UpdateProblemHandler).Methods("POST");
	router.HandleFunc("/deleteProblem", problems.DeleteProblemHandler).Methods("POST");
	
	router.HandleFunc("/tests/{problem_id}", tests.GetProblemTestsHandler).Methods("GET");
	router.HandleFunc("/tests/createTest", tests.CreateTestHandler).Methods("POST");
	router.HandleFunc("/tests/editTest", tests.EditTestHandler).Methods("POST");
	
	router.HandleFunc("/contest/setActive", contest.ActivateContestHandler).Methods("POST");
	router.HandleFunc("/contest/load", contest.LoadContestHandler).Methods("GET");
	router.HandleFunc("/contest/save", contest.SaveContestHandler).Methods("POST");
}