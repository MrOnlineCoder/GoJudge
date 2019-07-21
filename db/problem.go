package db

type Problem struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Timelimit int `json:"timelimit"`
	Memlimit int `json:"memlimit"`
	Text string
}