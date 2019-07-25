package contest

import (
	"log"
	"time"
	"encoding/json"
	"io/ioutil"

	"gojudge/db"
)

const (
	CONTEST_MODE_TIME = "time"
	CONTEST_MODE_TESTS = "tests"
)

type ContestProblem struct {
	Problem db.Problem `json:"problem"`
	Points int `json:"points"`
}

type Contest struct {
	Name string `json:"name"`
	StartTime int64 `json:"start_time"`
	EndTime int64 `json:"end_time"`
	Mode string `json:"mode"`
	Problemset []ContestProblem `json:"problemset"`
}

var currentContest Contest
var isActive bool

func GetProblemset() []ContestProblem {
	return currentContest.Problemset;
}

func IsRunning() bool {
	now := time.Now().Unix()*1000;
	start := currentContest.StartTime;
	end := currentContest.EndTime;

	return now > start && now < end;
}

func SetContest(c Contest) {
	currentContest = c;
}

func Activate() {
	isActive = true;
}

func Deactivate() {
	isActive = false;
}

func GetContest() Contest {
	return currentContest;
}

func IsContestActive() bool {
	return isActive;
}

func SaveContest() {
	file, err := json.MarshalIndent(currentContest, "", " ");

	if err != nil {
		log.Printf("[Contest] Failed to marshal contest data: %s\n", err.Error());
		return;
	}

	err = ioutil.WriteFile("contest.json", file, 0644)

	if err != nil {
		log.Printf("[Contest] Failed to write contest data: %s\n", err.Error());
	}
}

func LoadContest() {
	file, err := ioutil.ReadFile("contest.json");
 
	if err != nil {
		log.Printf("[Contest] Didn't load contest data: %s\n", err.Error());
		return;
	}
 
	err = json.Unmarshal([]byte(file), &currentContest);

	if err != nil {
		log.Printf("[Contest] Didn't load JSON contest data: %s\n", err.Error());
		return;
	}
}