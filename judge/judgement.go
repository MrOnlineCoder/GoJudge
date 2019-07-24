package judge

import (
  "fmt"
  "os"
  "os/exec"
  "path/filepath"
  "crypto/sha256"
  "encoding/hex"
  "runtime"
  "log"

  "gojudge/db"
)

const SANDBOX_DIR = "./sandbox";

type Judgement struct {
	Problem *db.Problem
	Submission *db.Submission
	Tests []db.Test
}

func hash256(raw string) string {
	hash := sha256.New()
  hash.Write([]byte(raw))
  
  return hex.EncodeToString(hash.Sum(nil))
}

func (jd Judgement) GetSandboxId() string {
	return fmt.Sprintf("sub_%d_%s", jd.Submission.Id, hash256(jd.Submission.Sourcecode));
}

func (jd Judgement) GetSourcePath() string {
	return filepath.Join(SANDBOX_DIR, fmt.Sprintf("%s.%s", jd.GetSandboxId(), jd.Submission.Language));
}

func (jd Judgement) GetExePath() string {
	if runtime.GOOS == "windows" {
		return filepath.Join(SANDBOX_DIR, fmt.Sprintf("%s.exe", jd.GetSandboxId()));
	} else {
		return filepath.Join(SANDBOX_DIR, fmt.Sprintf("%s", jd.GetSandboxId()));
	}
}

func (jd Judgement) SaveSourcecode() error {
	path := jd.GetSourcePath();
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

func (jd Judgement) Compile() error {
	sourcepath := jd.GetSourcePath();
	exepath := jd.GetExePath();

	cmd := exec.Command("g++", "-o", exepath, sourcepath);

	err := cmd.Run();

	if err != nil {
		return err
	}

	return nil
}

func runSingleTest(jd *Judgement, test *db.Test) string {
	output, err := RunInSandbox(jd.GetExePath(), test.Input, jd.Problem.Timelimit, jd.Problem.Memlimit);

	if err != nil {
		//Sandbox runner puts verdict into error object
		sandboxVerdict := err.Error();
		return sandboxVerdict;
	}

	var verdict string;


	switch test.CheckMethod {
		case CHECK_STRICT:
			verdict = RunStrictCheck(test.Input, test.Output, output);
		case CHECK_WHITESPACE_STRICT:
			verdict = RunWhitespaceCheck(test.Input, test.Output, output);
		case CHECK_BY_CHECKER:
			verdict = VERDICT_ERROR_FAIL;
		default:
			log.Printf("[Judge] Test #%d has invalid checking method: %d", test.Id, test.CheckMethod);
			verdict = VERDICT_ERROR_FAIL;
	}

	return verdict;
}

func (jd Judgement) RunTests() JudgeResult {
	result := JudgeResult{
		SubmissionId: jd.Submission.Id,
		Verdict: VERDICT_ERROR_FAIL,
		PassedTests: 0,
	};

	for _, test := range jd.Tests {
		v := runSingleTest(&jd, &test);

		result.Verdict = v;

		if v != VERDICT_OK {
			break;
		} 

		result.PassedTests++;
	}

	return result
}

func (jd Judgement) Cleanup() {
	os.Remove(jd.GetSourcePath());
	os.Remove(jd.GetExePath());
}