package db

import (
	"errors"
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

func CreateSubmission(s Submission) bool {
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

	_, err := Maindb.Exec(createSql,
		s.UserId,
		s.Time,
		s.Language,
		s.Sourcecode,
		s.ProblemId,
		s.Verdict,
		s.PassedTests,
	);

	if err != nil {
		return false
	}

	return true
}

func GetUserSubmissions(user_id int) ([]Submission, error) {
	const getForUserSql = `
		SELECT * FROM "submissions" ORDER BY "time" DESC
	`;

	list := []Submission{}

	rows, err := Maindb.Query(getForUserSql);

	if err != nil {
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

func SetSubmissionVerdict(id int, v string) bool {
	const updateSql = `
		UPDATE "submissions" SET
		"verdict" = $1,
		WHERE
		"id" = $2
	`;

	_, err := Maindb.Exec(updateSql, 
		v,
		id
	);

	if err != nil {
		log.Println(err)
		return false
	}

	return true
}