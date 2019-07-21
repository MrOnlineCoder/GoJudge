package db

import (
	"database/sql"
	"log"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

var Maindb *sql.DB

func runSimpleStmt(raw string) bool {
	stmt, err := Maindb.Prepare(raw);

	if err != nil {
		log.Fatal(err)
		return false
	}

	stmt.Exec()

	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

func checkTables() {
	const usersSql = `
		CREATE TABLE IF NOT EXISTS "users" (
		  "id" INTEGER PRIMARY KEY AUTOINCREMENT,
		  "username" VARCHAR(255),
		  "fullname" VARCHAR(255),
		  "password" VARCHAR(255),
		  "access" INTEGER
	);`;

	const problemsSql = `
		CREATE TABLE IF NOT EXISTS "problems" (
		  "id" INTEGER PRIMARY KEY AUTOINCREMENT,
		  "name" VARCHAR(255),
		  "timelimit" INTEGER,
		  "memlimit" INTEGER,
		  "text" TEXT
	);`;

	const submissionsSql = `
		CREATE TABLE IF NOT EXISTS "submissions" (
		  "id" INTEGER PRIMARY KEY AUTOINCREMENT,
		  "user_id" INTEGER,
		  "time" INTEGER,
		  "lang" VARCHAR(255),
		  "sourcecode" BLOB,
		  "problem_id" INTEGER,
		  "verdict" VARCHAR(64),
		  "passed_tests" INTEGER
	);`;

	runSimpleStmt(usersSql);
	runSimpleStmt(problemsSql);
	runSimpleStmt(submissionsSql);
}

func insertDefaults() {
	CreateUser(User{
		Username: "admin",
		Fullname: "Administrator",
		Password: "c7ad44cbad762a5da0a452f9e854fdc1e0e7a52a38015f23f3eab1d80b931dd472634dfac71cd34ebc35d16ab7fb8a90c81f975113d6c7538dc69dd8de9077ec",
		Access: 2,
	});
}

func exists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil { return true, nil }
    if os.IsNotExist(err) { return false, nil }
    return true, err
}

func isDBPresent() bool {
	ex, err := exists("./gojudge.db")

	if err != nil {
		panic(err)
		return false
	}

	return ex
}

func Initialize() bool {
	var err error;

	shouldInsertValues := !isDBPresent();

	Maindb, err = sql.Open("sqlite3", "./gojudge.db");

	if err != nil {
		log.Fatal(err)
		return false
	}

	checkTables();

	if shouldInsertValues {
		insertDefaults();
	}

	return true
}

func Close() bool {
	Maindb.Close()
	return true
}