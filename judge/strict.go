package judge

import (
	"strings"
)

func RunWhitespaceCheck(in string, out string, sout string) string {
	newOut := strings.Replace(out, "\n", " ", -1);
	newSout := strings.Replace(sout, "\n", " ", -1);

	newOut = strings.Replace(newOut, "\t", " ", -1);
	newSout = strings.Replace(newSout, "\t", " ", -1);

	return RunStrictCheck(in, newOut, newSout);
}

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