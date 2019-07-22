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

func ActivateContestHandler(w http.ResponseWriter, r *http.Request) {
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
	contest.RebuildProblemsetCache();
	contest.Activate();

	utils.SendSuccess(w, map[string]interface{}{});
}