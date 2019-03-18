package main

type BatchSettingsDTO struct {
	Count    int32  `json:"count"`
	Size     string `json:"size"`
	Duration string `json:"duration"`
}
