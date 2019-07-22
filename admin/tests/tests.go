package tests

import (
	"net/http"
 	"encoding/json"
 	"strconv"

	"gojudge/auth"
	"gojudge/api/utils"
	"gojudge/db"

	"github.com/gorilla/mux"
)

type CreateTestBody struct {
	Test db.Test `json:"test"`
}


func GetProblemTestsHandler(w http.ResponseWriter, r *http.Request) {
	_, err := auth.ValidateAccess(r, auth.ACCESS_ADMIN);

	if err != nil {
		utils.SendError(w, err.Error());
		return;
	}

	params := mux.Vars(r);

	problem_id, err := strconv.Atoi(params["problem_id"]);

	if err != nil {
		utils.SendError(w, "Invalid problem_id.");
		return;
	}

	tests, err := db.GetTestsForProblem(problem_id);

	if err != nil {
		utils.SendError(w, err.Error());
		return;
	}

	utils.SendSuccess(w, map[string]interface{} {
		"tests": tests,
	});
}

func CreateTestHandler(w http.ResponseWriter, r *http.Request) {
	_, err := auth.ValidateAccess(r, auth.ACCESS_ADMIN);
	
	if err != nil {
		utils.SendError(w, err.Error());
		return;
	}

	parsedBody := &CreateTestBody{};

	err = json.NewDecoder(r.Body).Decode(parsedBody);

	if err != nil {
		utils.SendError(w, err.Error());
		return;
	}

	ok := db.CreateTest(parsedBody.Test);

	if !ok {
		utils.SendError(w, "Database write error.");
		return;
	}

	utils.SendSuccess(w, map[string]interface{}{});
}