package problems

import (
	"net/http"
 	"encoding/json"

	"gojudge/auth"
	"gojudge/api/utils"
	"gojudge/db"
)

type CreateProblemBody struct {
	Problem db.Problem `json:"problem"`
}

type UpdateProblemBody struct {
	Problem db.Problem `json:"problem"`
}

type DeleteProblemBody struct {
	ProblemID int `json:"problem_id"`
}

func ListProblemsHandler(w http.ResponseWriter, r *http.Request) {
	_, err := auth.ValidateAccess(r, auth.ACCESS_ADMIN);

	if err != nil {
		utils.SendError(w, err.Error());
		return;
	}

	problems, err := db.GetAllProblems();

	if err != nil {
		utils.SendError(w, err.Error());
		return;
	}

	utils.SendSuccess(w, map[string]interface{} {
		"problems": problems,
	});
}

func CreateProblemHandler(w http.ResponseWriter, r *http.Request) {
	_, err := auth.ValidateAccess(r, auth.ACCESS_ADMIN);

	if err != nil {
		utils.SendError(w, err.Error());
		return;
	}

	parsedBody := &CreateProblemBody{};

	err = json.NewDecoder(r.Body).Decode(parsedBody);

	if err != nil {
		utils.SendError(w, err.Error());
		return;
	}

	ok := db.CreateProblem(parsedBody.Problem);

	if !ok {
		utils.SendError(w, "Database write error.");
		return;
	}

	utils.SendSuccess(w, map[string]interface{}{});
}

func UpdateProblemHandler(w http.ResponseWriter, r *http.Request) {
	_, err := auth.ValidateAccess(r, auth.ACCESS_ADMIN);

	if err != nil {
		utils.SendError(w, err.Error());
		return;
	}

	parsedBody := &UpdateProblemBody{};

	err = json.NewDecoder(r.Body).Decode(parsedBody);

	if err != nil {
		utils.SendError(w, err.Error());
		return;
	}

	ok := db.UpdateProblem(parsedBody.Problem);

	if !ok {
		utils.SendError(w, "Database write error.");
		return;
	}

	utils.SendSuccess(w, map[string]interface{}{});
}

func DeleteProblemHandler(w http.ResponseWriter, r *http.Request) {
	_, err := auth.ValidateAccess(r, auth.ACCESS_ADMIN);

	if err != nil {
		utils.SendError(w, err.Error());
		return;
	}

	parsedBody := &DeleteProblemBody{};

	err = json.NewDecoder(r.Body).Decode(parsedBody);

	if err != nil {
		utils.SendError(w, err.Error());
		return;
	}

	ok := db.DeleteProblem(parsedBody.ProblemID);

	if !ok {
		utils.SendError(w, "Database write error.");
		return;
	}

	utils.SendSuccess(w, map[string]interface{}{});
}
