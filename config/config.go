package config

import (
	"os" 
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Port int
}

var cfg Config

func exists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil { return true, nil }
    if os.IsNotExist(err) { return false, nil }
    return true, err
}

func IsPresent() bool {
	ex, err := exists("./config.json")

	if err != nil {
		panic(err)
		return false
	}

	return ex
}