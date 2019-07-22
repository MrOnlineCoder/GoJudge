package db

import (
	"errors"
	"log"
)

type Test struct {
	Id int `json:"id"`
	ProblemId int `json:"problem_id"`
	Index int `json:"test_index"`
	CheckMethod int `json:"check_method"`
	CheckerId int `json:"checker_id"`
	IsSample bool `json:"is_sample"`
	Input string `json:"input"`
	Output string `json:"output"`
}

func CreateTest(t Test) bool {
	const createSql = `
		INSERT INTO "tests" (problem_id, test_index, check_method, checker_id, is_sample, input, output) VALUES
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

	_, err := Maindb.Exec(createSql, t.ProblemId, t.Index, t.CheckMethod, t.CheckerId, t.IsSample, t.Input, t.Output);

	if err != nil {
		log.Println(err);
		return false
	}

	return true
}

func GetTestsForProblem(problem_id int) ([]Test, error) {
	const getAllSql = `
		SELECT * FROM "tests" WHERE "problem_id" = $1
	`;

	list := []Test{}

	rows, err := Maindb.Query(getAllSql, problem_id);

	if err != nil {
		return list, errors.New("Database query failed.");
	}

	for rows.Next() {
    t := Test{}
    err := rows.Scan(&t.Id, 
    	&t.ProblemId, 
    	&t.Index, 
    	&t.CheckMethod, 
    	&t.CheckerId, 
    	&t.IsSample, 
    	&t.Input,
    	&t.Output)

    if err != nil {
      continue
    }
    list = append(list, t)
  }

	rows.Close()

	return list, nil
}

/*
func GetProblem(id int) (Problem, error) {
	const getProblemSql = `
		SELECT * FROM "problems" WHERE "id" = $1
	`;

	row := Maindb.QueryRow(getProblemSql, id);

	problem := Problem{};

	err := row.Scan(&problem.Id, &problem.Name, &problem.Timelimit, &problem.Memlimit, &problem.Text);

	if err != nil {
		return problem, errors.New("Database query failed.");
 }

 return problem, nil
}
*/
func UpdateTest(t Test) bool {
	const updateSql = `
		UPDATE "tests" SET
		"problem_id" = $1,
		"test_index" = $2,
		"check_method" = $3,
		"checker_id" = $4,
		"is_sample" = $5,
		"input" = $6,
		"output" = $7
		WHERE
		"id" = $8
	`;

	_, err := Maindb.Exec(updateSql, 
		t.ProblemId,
		t.Index,
		t.CheckMethod,
		t.CheckerId,
		t.IsSample,
		t.Input,
		t.Output,
		t.Id,
	);

	if err != nil {
		log.Println(err)
		return false
	}

	return true
}


func DeleteTest(tid int) bool {
	const deleteSql = `
		DELETE FROM "tests" WHERE "id" = $1
	`;

	_, err := Maindb.Exec(deleteSql, tid);

	if err != nil {
		return false
	}

	return true
}