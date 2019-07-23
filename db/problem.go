package db

import (
	"fmt"
	"errors"
	"strings"
	"strconv"
	"log"
)

type Problem struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Timelimit int `json:"timelimit"`
	Memlimit int `json:"memlimit"`
	Text string `json:"text"`
}

func CreateProblem(p Problem) bool {
	const createSql = `
		INSERT INTO "problems" (name, timelimit, memlimit, text) VALUES
		(
			$1,
			$2,
			$3,
			$4
		);
	`;

	_, err := Maindb.Exec(createSql, p.Name, p.Timelimit, p.Memlimit, p.Text);

	if err != nil {
		log.Printf("[DB] ERROR: Failed to create problem: %s\n", err.Error());
		return false
	}

	return true
}

func GetAllProblems() ([]Problem, error) {
	const getAllSql = `
		SELECT * FROM "problems"
	`;

	list := []Problem{}

	rows, err := Maindb.Query(getAllSql);

	if err != nil {
		log.Printf("[DB] ERROR: Failed to get list of all problem: %s\n", err.Error());
		return list, errors.New("Database query failed.");
	}

	for rows.Next() {
    p := Problem{}
    err := rows.Scan(&p.Id, &p.Name, &p.Timelimit, &p.Memlimit, &p.Text)
    if err != nil {
      continue
    }
    list = append(list, p)
  }

	rows.Close()

	return list, nil
}
/*
func GetProblemset(arr []int) ([]Problem, error) {
	const getAllSql = `
		SELECT * FROM "problems" WHERE "id" IN (%s)
	`;

	list := []Problem{}

	//We receive array of ints
	//To build SQL query, we have to transform it to []string
	var ids []string
	for _, i := range arr {
	    ids = append(ids, strconv.Itoa(i))
	}

	builtQuery := fmt.Sprintf(getAllSql, strings.Join(ids, ", "));

	rows, err := Maindb.Query(builtQuery);

	if err != nil {
		log.Printf("[DB] Failed to create problem: %s\n", err.Error());
		return list, errors.New("Database query failed.");
	}

	for rows.Next() {
    p := Problem{}
    err := rows.Scan(&p.Id, &p.Name, &p.Timelimit, &p.Memlimit, &p.Text)
    if err != nil {
      continue
    }
    fmt.Println("id #", p.Id)
    list = append(list, p)
  }

	rows.Close()

	return list, nil
}*/

func GetProblem(id int) (Problem, error) {
	const getProblemSql = `
		SELECT * FROM "problems" WHERE "id" = $1
	`;

	row := Maindb.QueryRow(getProblemSql, id);

	problem := Problem{};

	err := row.Scan(&problem.Id, &problem.Name, &problem.Timelimit, &problem.Memlimit, &problem.Text);

	if err != nil {
		log.Printf("[DB] ERROR: Failed to get problem #%d: %s\n", id, err.Error());
		return problem, errors.New("Database query failed.");
 }

 return problem, nil
}

func UpdateProblem(p Problem) bool {
	const updateSql = `
		UPDATE "problems" SET
		"name" = $1,
		"timelimit" = $2,
		"memlimit" = $3,
		"text" = $4
		WHERE
		"id" = $5
	`;

	_, err := Maindb.Exec(updateSql, p.Name, p.Timelimit, p.Memlimit, p.Text, p.Id);

	if err != nil {
		log.Printf("[DB] ERROR: Failed to update problem: %s\n", err.Error());
		return false
	}

	return true
}

func DeleteProblem(problem_id int) bool {
	const deleteSql = `
		DELETE FROM "problems" WHERE "id" = $1
	`;

	_, err := Maindb.Exec(deleteSql, problem_id);

	if err != nil {
		log.Printf("[DB] ERROR: Failed to delete problem: #%d: %s\n", problem_id, err.Error());
		return false
	}

	return true
}