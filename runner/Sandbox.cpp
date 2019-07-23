/*
	Go Judge Sandbox Runner

	by MrOnlineCoder (github.com/MrOnlineCoder/gojudge)

	(c) 2019

	License: see LICENSE file
*/

#include <Windows.h>
#include <iostream>
#include <string>
#include <sstream>

const static char* HELP_MSG = "Usage: sandbox_runner <target> <timelimit> <memlimit>\n"
	"Options:\n"
	"<target> - path to target executable\n"
	"<timelimit> - timelimit in milliseconds\n"
	"<memlimit> - memory limit in kilobytes\n"
	"\n"
	"Runner will exit with code 0 if no error happened.\n"
	"In that case, target's output will be piped and writen to runner stdout\n"
	"Otherwise, exit code is 1 and runner stderr contains the error verdict (TLE, MLE, FAIL)\n";


//This is minimum mem limit (4 MB), 1 page in x64
const static std::size_t MIN_MEM_LIMIT = 4 * 1024;

namespace Verdicts {
	const static std::string Fail = "FAIL";
	const static std::string TLE = "TIME_LIMIT_EXCEEDED";
	const static std::string MLE = "MEMORY_LIMIT_EXCEEDED";
	const static std::string RuntimeError = "RUNTIME_ERROR";
};

struct RunnerOptions {
	std::string exePath;
	long timelimit;
	long memlimit;
};

HANDLE createJob(RunnerOptions& options) {
	SECURITY_ATTRIBUTES jsa;
		
	jsa.nLength = sizeof(jsa);
	jsa.lpSecurityDescriptor = NULL;
	jsa.bInheritHandle = TRUE;

	HANDLE job = CreateJobObject(&jsa, "GoJudgeSandboxJob");

	if (job == NULL) {
		std::cerr << "CreateJobObject() error:" << GetLastError() << std::endl;
		return NULL;
	}

	JOBOBJECT_EXTENDED_LIMIT_INFORMATION jbeli;
	JOBOBJECT_BASIC_UI_RESTRICTIONS jbur;

	jbeli.BasicLimitInformation.LimitFlags = JOB_OBJECT_LIMIT_PROCESS_MEMORY;
	jbeli.ProcessMemoryLimit = options.memlimit * 1024;

	jbur.UIRestrictionsClass = JOB_OBJECT_UILIMIT_DESKTOP;
	jbur.UIRestrictionsClass |= JOB_OBJECT_UILIMIT_DISPLAYSETTINGS;
	jbur.UIRestrictionsClass |= JOB_OBJECT_UILIMIT_GLOBALATOMS;
	jbur.UIRestrictionsClass |= JOB_OBJECT_UILIMIT_READCLIPBOARD;
	jbur.UIRestrictionsClass |= JOB_OBJECT_UILIMIT_SYSTEMPARAMETERS;
	jbur.UIRestrictionsClass |= JOB_OBJECT_UILIMIT_WRITECLIPBOARD;

	if (SetInformationJobObject(job, JobObjectExtendedLimitInformation, &jbeli, sizeof(jbeli)) == 0) {
		std::cerr << "SetInformationJobObject(jbeli) error:" << GetLastError() << std::endl;
		CloseHandle(job);
		return NULL;
	}

	if (SetInformationJobObject(job, JobObjectBasicUIRestrictions, &jbur, sizeof(jbur)) == 0) {
		std::cerr << "SetInformationJobObject(jbur) error:" << GetLastError() << std::endl;
		CloseHandle(job);
		return NULL;
	}

	return job;
}

int runProcess(RunnerOptions& options) {
	STARTUPINFO si;
	PROCESS_INFORMATION pi;

	HANDLE jobObject = createJob(options);

	if (jobObject == NULL) {
		std::cerr << Verdicts::Fail;
		return 1;
	}

	ZeroMemory(&si, sizeof(si));
	si.cb = sizeof(si);
	ZeroMemory(&pi, sizeof(pi));

	//Redirect runner child std handles to runner handles
	//si.dwFlags = STARTF_USESTDHANDLES;
	si.hStdInput = GetStdHandle(STD_INPUT_HANDLE);
	si.hStdOutput = GetStdHandle(STD_OUTPUT_HANDLE);

	BOOL retCp = CreateProcess(
		NULL, //app name
		(LPSTR) options.exePath.c_str(), //command line
		NULL, //process attributes
		NULL, //thread attributes
		FALSE, //inherit handles
		CREATE_SUSPENDED, //creation flags
		NULL, //environment
		NULL, //current dir
		&si, //startup info
		&pi //process info
	);

	if (!retCp) {
		CloseHandle(jobObject);
		std::cerr << Verdicts::Fail;
		return 1;
	}

	if (AssignProcessToJobObject(jobObject, pi.hProcess) == 0) {
		std::cerr << "AssignProcessToJobObject() error:" << GetLastError() << std::endl;
		CloseHandle(jobObject);
		std::cerr << Verdicts::Fail;
		return 1;
	}

	ResumeThread(pi.hThread);

	DWORD processExitCode = 0;

	//This wait implemenets Time limit control
	DWORD waitResult = WaitForSingleObject(pi.hProcess, options.timelimit);

	int sandboxExitCode = 0;

	//If WaitForSingleObject timed out, return TLE
	if (waitResult == WAIT_TIMEOUT) {
		std::cerr << Verdicts::TLE;

		sandboxExitCode = 1;

		TerminateProcess(pi.hProcess, 1);

		CloseHandle(pi.hProcess);
		CloseHandle(pi.hThread);
		CloseHandle(jobObject);
	} else {
		GetExitCodeProcess(pi.hProcess, &processExitCode);
	}

	if (processExitCode != 0) {
		sandboxExitCode = 1;
		std::cerr << Verdicts::RuntimeError;
	}

	CloseHandle(pi.hProcess);
	CloseHandle(pi.hThread);
	CloseHandle(jobObject);

	return sandboxExitCode;
}

int main(int argc, char* argv[]) {
	if (argc < 4) {
		std::cout << HELP_MSG;
		return 1;
	}

	RunnerOptions options;

	options.exePath = argv[1];
	options.timelimit = std::stol(argv[2]);
	options.memlimit = std::stol(argv[3]);

	if (options.memlimit < MIN_MEM_LIMIT) {
		options.memlimit = MIN_MEM_LIMIT;
	}

	return runProcess(options);
}