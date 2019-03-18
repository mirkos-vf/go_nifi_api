package main

type AccessStatusDTO struct {
	Identity string `json:"identity"`
	Status   string `json:"status"`
	Message  string `json:"message"`
}
