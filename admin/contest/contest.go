package contest

import (
	"net/http"
 	"encoding/json"

	"gojudge/auth"
	"gojudge/api/utils"
	"gojudge/contest"
)

type ContestBody struct {
	Contest contest.Contest `json:"contest"`
}

type ActivateContestBody struct {
	Active bool `json:"active"`
}

func SaveContestHandler(w http.ResponseWriter, r *http.Request) {
	_, err := auth.ValidateAccess(r, auth.ACCESS_ADMIN);

	if err != nil {
		utils.SendError(w, err.Error());
		return;
	}
	parsedBody := &ContestBody{};

	err = json.NewDecoder(r.Body).Decode(parsedBody);

	if err != nil {
		utils.SendError(w, err.Error());
		return;
	}

	contest.SetContest(parsedBody.Contest);
	contest.SaveContest();

	utils.SendSuccess(w, map[string]interface{}{});
}

func ActivateContestHandler(w http.ResponseWriter, r *http.Request) {
	_, err := auth.ValidateAccess(r, auth.ACCESS_ADMIN);

	if err != nil {
		utils.SendError(w, err.Error());
		return;
	}

	parsedBody := &ActivateContestBody{};

	err = json.NewDecoder(r.Body).Decode(parsedBody);

	if err != nil {
		utils.SendError(w, err.Error());
		return;
	}

	if parsedBody.Active {
		contest.Activate();
	} else {
		contest.Deactivate();
	}

	utils.SendSuccess(w, map[string]interface{}{});
}

func LoadContestHandler(w http.ResponseWriter, r *http.Request) {
	_, err := auth.ValidateAccess(r, auth.ACCESS_ADMIN);

	if err != nil {
		utils.SendError(w, err.Error());
		return;
	}

	utils.SendSuccess(w, map[string]interface{}{
		"contest": contest.GetContest(),
	});
}