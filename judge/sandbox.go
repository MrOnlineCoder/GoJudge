package judge

import (
	"os/exec"
	"strings"
	"bytes"
	"errors"
	"strconv"
)

const SANDBOX_RUNNER_EXE = "sandbox_runner";

func RunInSandbox(exePath string, input string, timelimit int, memlimit int) (string, error) {
	cmd := exec.Command(SANDBOX_RUNNER_EXE, 
		exePath, 
		strconv.Itoa(timelimit), 
		strconv.Itoa(memlimit),
	);

	finput := strings.TrimSpace(input);

	cmd.Stdin = strings.NewReader(finput);

	var exeOut bytes.Buffer;
	var runnerErr bytes.Buffer;

	cmd.Stdout = &exeOut;
	cmd.Stderr = &runnerErr;

	err := cmd.Run();

	if err != nil {
		return "", errors.New(runnerErr.String());
	}

	return strings.TrimSpace(exeOut.String()), nil
}