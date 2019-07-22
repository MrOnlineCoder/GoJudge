package contest

import (
	"log"
	"time"
	"gojudge/db"
)

type Contest struct {
	Name string `json:"name"`
	StartTime int64 `json:"start_time"`
	EndTime int64 `json:"end_time"`
	Problemset []int `json:"problemset"`
}

var currentContest Contest
var isActive bool
var problemsetCache []db.Problem

func RebuildProblemsetCache() {
	problemsetCache = []db.Problem{};

	problemsetInts := currentContest.Problemset;

	for _, id := range problemsetInts {
			problem, err := db.GetProblem(id);

			if err != nil {
				log.Println("ERROR: Couldn't build problemset cache entry:", err)
				continue;
			}

			problemsetCache = append(problemsetCache, problem)
	}
}

func GetProblemset() []db.Problem {
	return problemsetCache;
}

func IsStarted() bool {
	now := time.Now().Unix()*1000;
	start := currentContest.StartTime;

	return now > start;
}

func SetContest(c Contest) {
	currentContest = c;
}

func Activate() {
	isActive = true;
}

func GetContest() Contest {
	return currentContest;
}

func IsContestActive() bool {
	return isActive;
}