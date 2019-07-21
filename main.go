/*
	Go Judge Software

	by MrOnlineCoder (github.com/MrOnlineCoder/gojudge)

	(c) 2019

	License: see LICENSE file
*/

package main

import (
	"fmt"
	"gojudge/server"
	"gojudge/db"
)

func main() {
	fmt.Println("= [GoJudge] =");	
	fmt.Println("Loading database....");

	if !db.Initialize() {
		fmt.Println("ERROR: couldn't initialize database.")
		return
	}

	fmt.Println("Starting server...");

	server.Setup();
	server.Run();	

	db.Close();
}