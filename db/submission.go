package db

import (
	"errors"
	"log"
)

type Submission struct {
	Id int `json:"id"`
	UserId int `json:"user_id"`
	Time int64 `json:"time"`
	Language string `json:"lang"`
	Sourcecode string `json:"sourcecode"`
	ProblemId int `json:"problem_id"`
	Verdict string `json:"verdict"`
	PassedTests int `json:"passed_tests"`
}

func CreateSubmission(s Submission) (int, error) {
	const createSql = `
		INSERT INTO "submissions" (user_id, time, lang, sourcecode, problem_id, verdict, passed_tests) VALUES
		(
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7
		);
	`;

	res, err := Maindb.Exec(createSql,
		s.UserId,
		s.Time,
		s.Language,
		s.Sourcecode,
		s.ProblemId,
		s.Verdict,
		s.PassedTests,
	);

	if err != nil {
		log.Printf("[DB] ERROR: Failed to create submission: %s\n", err.Error());
		return -1, err
	}

	id, err := res.LastInsertId();

	if err != nil {
		return -1, err
	}

	return int(id), nil
}

func GetUserSubmissions(user_id int) ([]Submission, error) {
	const getForUserSql = `
		SELECT * FROM "submissions" WHERE "user_id" = $1 ORDER BY "time" DESC
	`;

	list := []Submission{}

	rows, err := Maindb.Query(getForUserSql, user_id);

	if err != nil {
		log.Printf("[DB] ERROR: Failed to get user (%d) submissions: %s\n", user_id, err.Error());
		return list, errors.New("Database query failed.");
	}

	for rows.Next() {
    s := Submission{}
    err := rows.Scan(&s.Id,
    	&s.UserId,
    	&s.Time,
    	&s.Language,
    	&s.Sourcecode,
    	&s.ProblemId,
    	&s.Verdict,
    	&s.PassedTests,
    )
    if err != nil {
      continue
    }
    list = append(list, s)
  }

	rows.Close()

	return list, nil
}

func SetSubmissionVerdict(id int, v string, tests int) bool {
	const updateSql = `
		UPDATE "submissions" SET
		"verdict" = $1,
		"passed_tests" = $2
		WHERE
		"id" = $3
	`;

	_, err := Maindb.Exec(updateSql, 
		v,
		tests,
		id,
	);

	if err != nil {
		log.Printf("[DB] ERROR: Failed to set submission (#%d) verdict to %s: %s\n", id, v, err.Error());
		return false
	}

	return true
}