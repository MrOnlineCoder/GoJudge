package contest

import (
	"strconv"
	"net/http"
	
	"gojudge/api/utils"

 	"github.com/gorilla/mux"
)

func ContestStatusHandler(w http.ResponseWriter, r *http.Request) {
	if !IsContestActive() {
		utils.SendSuccess(w, map[string]interface{} {
			"active": false,
		});
		return;
	}

	contest := GetContest();

	utils.SendSuccess(w, map[string]interface{} {
		"active": true,
		"contest": contest,
	});
}

func ContestProblemsetHandler(w http.ResponseWriter, r *http.Request) {
	if !IsContestActive() {
		utils.SendSuccess(w, map[string]interface{} {
			"active": false,
		});
		return;
	}

	utils.SendSuccess(w, map[string]interface{} {
		"active": true,
		"problemset": GetProblemset(),
	});
}

func ContestProblemHandler(w http.ResponseWriter, r *http.Request) {
	if !IsContestActive() {
		utils.SendSuccess(w, map[string]interface{} {
			"active": false,
		});
		return;
	}

	params := mux.Vars(r);

	problemIndex, err := strconv.Atoi(params["index"]);

	if err != nil {
		utils.SendError(w, "Invalid problem index.");
		return;
	}

	if problemIndex >= len(GetProblemset()) {
		utils.SendError(w, "Problem index out of range.");
		return;
	}

	problem := GetProblemset()[problemIndex];

	utils.SendSuccess(w, map[string]interface{} {
		"active": true,
		"problem": problem,
	});
}

func InitContestAPI(router *mux.Router) {
	router.HandleFunc("/status", ContestStatusHandler).Methods("GET");
	router.HandleFunc("/problemset", ContestProblemsetHandler).Methods("GET");
	router.HandleFunc("/problemset/{index}", ContestProblemHandler).Methods("GET");
}