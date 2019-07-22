package judge

import (
  "fmt"
  "os"
  "os/exec"
  "log"
  "bytes"
  "path/filepath"
  "crypto/sha256"
  "encoding/hex"
)

const SANDBOX_DIR = "./sandbox";

func hash256(raw string) string {
	hash := sha256.New()
  hash.Write([]byte(raw))
  
  return hex.EncodeToString(hash.Sum(nil))
}

func makeJudgementId(jd *Judgement) string {
	return fmt.Sprintf("%d_%s", jd.Submission.Time, hash256(jd.Submission.Sourcecode)[0:16]);
}

func makeSourcePath(jd *Judgement) string {
	return filepath.Join(SANDBOX_DIR, fmt.Sprintf("%s.%s", jd.Id, jd.Submission.Language));
}

func makeExePath(jd *Judgement) string {
	return filepath.Join(SANDBOX_DIR, fmt.Sprintf("%s", jd.Id));
}

func writeSourcecode(jd *Judgement) error {
	path := makeSourcePath(jd);

	sourcefile, err := os.Create(path)

	if err != nil {
		return err
	}

	_, err = sourcefile.WriteString(jd.Submission.Sourcecode);

	if err != nil {
		return err
	}

	sourcefile.Sync();
	sourcefile.Close();

	return nil
}

func compile(jd *Judgement) error {
	sourcepath := makeSourcePath(jd);
	exepath := makeExePath(jd);

	//args := fmt.Sprintf("-o \"%s\" \"%s\"", exepath, sourcepath);

	cmd := exec.Command("g++", "-o", exepath, sourcepath);

	var out bytes.Buffer;
	var eb bytes.Buffer;
	cmd.Stdout = &out;
	cmd.Stderr = &eb;

	err := cmd.Run();

	if err != nil {
		return err
	}

	return nil
}

func singleTest() {

}

func runTests(jd *Judgement) error {
	
}

func JudgeWorker(worker_id int, judgements chan *Judgement, results chan JudgeResult) {
	log.Printf("[Judge] Started worker %d", worker_id);

	for j := range judgements {
		j.Id = makeJudgementId(j);

    log.Printf("[Judge] Worker %d starts judging %s", worker_id, j.Id);

    err := writeSourcecode(j);

    if err != nil {
    	log.Printf("[Judge] ERROR: Couldn't write sourcecode: %s", err.Error())
    	continue;
    }

    err = compile(j);

    log.Printf("[Judge] Worker %d finished judging %s", worker_id, j.Id);
  }
}