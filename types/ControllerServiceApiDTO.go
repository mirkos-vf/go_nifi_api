package main

type ControllerServiceApiDTO struct {
	Type   string    `json:"type"`
	Bundle BundleDTO `json:"bundle"`
}
