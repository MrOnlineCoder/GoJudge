/*
	Go Judge Software

	by MrOnlineCoder (github.com/MrOnlineCoder/gojudge)

	(c) 2019

	License: see LICENSE file
*/

package main

import (
	"log"
	"gojudge/server"
	"gojudge/db"
	"gojudge/judge"
	"gojudge/contest"
)

func main() {
	log.Println("=== [GoJudge] ===");	
	log.Println("[Main] Loading database....");

	if !db.Initialize() {
		log.Fatal("ERROR: couldn't initialize database.")
		return
	}

	contest.LoadContest();

	judge.StartWorkers();

	log.Println("[Main] Starting server...");

	server.Setup();
	server.Run();	

	db.Close();
}