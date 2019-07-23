package judge

import (
  "log"
)

func JudgeWorker(worker_id int, judgements chan *Judgement, results chan JudgeResult) {
	log.Printf("[Judge] Started worker %d", worker_id);

	for j := range judgements {
		results <- MakeResult(j, VERDICT_QUEUED);

    log.Printf("[Judge] Worker %d starts judging #%d", worker_id, j.Submission.Id);

    err := j.SaveSourcecode();

    if err != nil {
    	results <- MakeResult(j, VERDICT_ERROR_FAIL);
    	log.Printf("[Judge] ERROR: Couldn't write sourcecode: %s", err.Error())
    	continue;
    }

    results <- MakeResult(j, VERDICT_COMPILING);

    err = j.Compile();

    if err != nil {
    	j.Cleanup();
    	results <- MakeResult(j, VERDICT_ERROR_COMPILE);
    	continue;
    }

    results <- MakeResult(j, VERDICT_CHECKING);

    testResult := j.RunTests();

    results <- testResult;

    j.Cleanup();

    log.Printf("[Judge] Worker %d finished judging #%d, verdict = %s", worker_id, j.Submission.Id, testResult.Verdict);
  }
}