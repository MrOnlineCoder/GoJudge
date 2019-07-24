package contest

import (
	"strconv"
	"time"
	"net/http"
 	"encoding/json"
	
	"gojudge/api/utils"
	"gojudge/db"
	"gojudge/auth"
	"gojudge/judge"

 	"github.com/gorilla/mux"
)

type ContestSubmitBody struct {
	ProblemIndex int `json:"problem_index"`
	Language string `json:"language"`
	Sourcecode string `json:"sourcecode"`
}

type ContestExample struct {
	Input string `json:"input"`
	Output string `json:"output"`
}

type ContestProblemNamesBody struct {
	Ids []string `json:"problem_ids"`
}

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

	if !IsRunning() {
		utils.SendSuccess(w, map[string]interface{} {
			"active": true,
			"not_started": true,
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

	if !IsRunning() {
		utils.SendSuccess(w, map[string]interface{} {
			"active": true,
			"not_started": true,
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

func ContestProblemExamplesHandler(w http.ResponseWriter, r *http.Request) {
	if !IsContestActive() {
		utils.SendSuccess(w, map[string]interface{} {
			"active": false,
		});
		return;
	}

	if !IsRunning() {
		utils.SendSuccess(w, map[string]interface{} {
			"active": true,
			"not_started": true,
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

	tests, err := db.GetSamplesForProblem(problem.Id);

	if err != nil {
		utils.SendError(w, "Database query error.");
		return;
	}

	examples := []ContestExample{};

	for _, test := range tests {
		ex := ContestExample{};

		ex.Input = test.Input;
		ex.Output = test.Output;

		examples = append(examples, ex);
	}


	utils.SendSuccess(w, map[string]interface{} {
		"examples": examples,
	});
}

func ContestSubmitHandler(w http.ResponseWriter, r *http.Request) {
	if !IsContestActive() {
		utils.SendError(w, "Contest not active.");
		return;
	}

	if !IsRunning() {
		utils.SendError(w, "Contest not started.");
		return;
	}

	user, err := auth.ValidateAccess(r, auth.ACCESS_PARTICIPANT);

	if err != nil {
		utils.SendError(w, err.Error());
		return;
	}

	body := &ContestSubmitBody{};

	problemset := GetProblemset();

	err = json.NewDecoder(r.Body).Decode(body);

	if err != nil {
		utils.SendError(w, err.Error());
		return;
	}

	problem := problemset[body.ProblemIndex];

	tests, err := db.GetTestsForProblem(problem.Id);

	if err != nil {
		utils.SendError(w, err.Error());
		return;
	}

	sub := db.Submission{};

	sub.UserId = user.Id;
	sub.Time = time.Now().Unix() * 1000;
	sub.Language = body.Language;
	sub.Sourcecode = body.Sourcecode;
	sub.ProblemId = problem.Id;
	sub.Verdict = judge.VERDICT_PENDING;
	sub.PassedTests = 0;

	subId, err := db.CreateSubmission(sub);

	if err != nil {
		utils.SendError(w, "Database write error.");
		return;
	}

	sub.Id = subId;

	judge.ProcessSubmission(&problem, &sub, tests);

	utils.SendSuccess(w, map[string]interface{} {});
}

func ContestSubmissionsHandler(w http.ResponseWriter, r *http.Request) {
	user, err := auth.ValidateAccess(r, auth.ACCESS_PARTICIPANT);

	if err != nil {
		utils.SendError(w, err.Error());
		return;
	}

	submissions, err := db.GetUserSubmissions(user.Id);	

	if err != nil {
		utils.SendError(w, err.Error());
		return;
	}

	utils.SendSuccess(w, map[string]interface{} {
		"submissions": submissions,
	});
}

func ContestProblemNamesHandler(w http.ResponseWriter, r *http.Request) {
	_, err := auth.ValidateAccess(r, auth.ACCESS_PARTICIPANT);

	if err != nil {
		utils.SendError(w, err.Error());
		return;
	}

	parsedBody := &ContestProblemNamesBody{};

	err = json.NewDecoder(r.Body).Decode(parsedBody);

	if err != nil {
		utils.SendError(w, err.Error());
		return;
	}

	names, err := db.GetProblemNames(parsedBody.Ids);

	if err != nil {
		utils.SendError(w, err.Error());
		return;
	}

	utils.SendSuccess(w, map[string]interface{} {
		"problems": names,
	});
}

func InitContestAPI(router *mux.Router) {
	router.HandleFunc("/status", ContestStatusHandler).Methods("GET");
	router.HandleFunc("/problemset", ContestProblemsetHandler).Methods("GET");
	router.HandleFunc("/problemset/{index}", ContestProblemHandler).Methods("GET");
	router.HandleFunc("/problemset/{index}/examples", ContestProblemExamplesHandler).Methods("GET");
	
	router.HandleFunc("/problemNames", ContestProblemNamesHandler).Methods("POST");
	
	router.HandleFunc("/submit", ContestSubmitHandler).Methods("POST");
	router.HandleFunc("/submissions", ContestSubmissionsHandler).Methods("GET");
}