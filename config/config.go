package config

import (
	"os" 
	"encoding/json"
	"io/ioutil"
)

type LimitsConfig struct {
	Sourcecode uint `json:"sourcecode"`
}

type ServerConfig struct {
	Port int `json:"port"`
	MaxWorkers int `json:"max_workers"`
}

type Config struct {
	Server ServerConfig `json:"server"` 
	Limits LimitsConfig `json:"limits"`
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

func SetDefault() {
	cfg.Server.Port = 1337;
	cfg.Server.MaxWorkers = 0;
	cfg.Limits.Sourcecode = 131072;
}

func Load() {
	data, err := ioutil.ReadFile("config.json")

	if err != nil {
		panic(err)
		return;
	}

	err = json.Unmarshal(data, &cfg);

	if err != nil {
		panic(err)
		return;
	}	
}

func Save() {
	bytes, err := json.MarshalIndent(cfg, "", "	");

	if err != nil {
		panic(err)
		return;
	}

	err = ioutil.WriteFile("config.json", bytes, 0644)	

	if err != nil {
		panic(err)
		return;
	}	
}

func Set(newCfg Config) {
	cfg = newCfg;
}

func Get() *Config {
	return &cfg;
}