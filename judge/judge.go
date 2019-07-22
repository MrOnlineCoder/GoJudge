package judge

import (
	"gojudge/db"
	"runtime"
)

const (
	VERDICT_PENDING = "PENDING"
	VERDICT_QUEUED = "QUEUED"
	VERDICT_COMPILING = "COMPILING"
	VERDICT_CHECKING = "CHECKING"
	VERDICT_ERROR_WA = "WRONG_ANSWER"
	VERDICT_ERROR_PE = "PRESENTATION_ERROR"
	VERDICT_ERROR_FAIL = "FAIL"
	VERDICT_OK = "OK" 
)

type Judgement struct {
	Id string // internal id for storing files
	Problem *db.Problem
	Submission *db.Submission
	Tests []db.Test
}

type JudgeResult struct {
	submission_id int
	verdict string
}

var resultsChan chan JudgeResult
var judgementsChan chan *Judgement

func StartWorkers() {
	resultsChan = make(chan JudgeResult)
	judgementsChan = make(chan *Judgement)

	numWorkers := runtime.NumCPU();

	for w := 1; w <= numWorkers; w++ {
    go JudgeWorker(w, judgementsChan, resultsChan)
  } 

  go ResultsWatch();
}

func ResultsWatch() {
	for r := range resultsChan {
		
	}
}

func ProcessSubmission(problem *db.Problem, sub *db.Submission, tests []db.Test) {
	judgement := &Judgement{
		Problem: problem,
		Submission: sub,
		Tests: tests,
	};

	judgementsChan <- judgement;
}