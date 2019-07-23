package judge

import (
	"strings"
)

func RunStrictCheck(testInput string, testOutput string, submissionOutput string) string {
	trimmedTest := strings.TrimSpace(testOutput);
	trimmedSubmission := strings.TrimSpace(submissionOutput);

	result := trimmedTest == trimmedSubmission;

	if result {
		return VERDICT_OK;
	} else {
		return VERDICT_ERROR_WA;
	}
}