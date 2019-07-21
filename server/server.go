package server

import (
	"log"
	"net/http"
	"gojudge/api"
)

func Setup() {
	
}

func Run() bool {
	err := http.ListenAndServe(":1337", api.Create())

	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}