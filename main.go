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
	"gojudge/config"
)

func main() {
	log.Println("=== [GoJudge] ===");	
	log.Println("[Main] Loading config...");

	if !config.IsPresent() {
		config.SetDefault();
		config.Save();

		log.Println("[Main] ------------------------------------------------------------------- ");
		log.Println("[Main] Default config has been set.");
		log.Println("[Main] Please, configure your GoJudge installation via admin user-interface");
		log.Println("[Main] ------------------------------------------------------------------- ");
	} else {
		config.Load();
	}

	log.Println("[Main] Loading database...");

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