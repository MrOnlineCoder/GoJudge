package judge

import (
	"gojudge/db"
	"gojudge/realtime"
	"runtime"
)

const (
	VERDICT_PENDING = "PENDING"
	VERDICT_QUEUED = "QUEUED"
	VERDICT_COMPILING = "COMPILING"
	VERDICT_CHECKING = "CHECKING"
	VERDICT_ERROR_WA = "WRONG_ANSWER"
	VERDICT_ERROR_COMPILE = "COMPILATION_ERROR"
	VERDICT_ERROR_PE = "PRESENTATION_ERROR"
	VERDICT_ERROR_FAIL = "FAIL"
	VERDICT_ERROR_TIME = "TIME_LIMIT_EXCEEDED"
	VERDICT_ERROR_MEM = "MEMORY_LIMIT_EXCEEDED"
	VERDICT_ERROR_RUNTIME = "RUNTIME_ERROR"
	VERDICT_OK = "OK" 
)

const (
	CHECK_STRICT = 0
	CHECK_TOKEN = 1
	CHECK_BY_CHECKER = 2
)

type JudgeResult struct {
	SubmissionId int
	Verdict string
	PassedTests int
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

func MakeResult(jd *Judgement, verdict string) JudgeResult {
	return JudgeResult{
		SubmissionId: jd.Submission.Id,
		Verdict: verdict,
		PassedTests: 0,
	};
}

func ResultsWatch() {
	for r := range resultsChan {
		db.SetSubmissionVerdict(r.SubmissionId, r.Verdict, r.PassedTests);
		realtime.EmitSubmissionUpdate(r.SubmissionId, r.Verdict, r.PassedTests);
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