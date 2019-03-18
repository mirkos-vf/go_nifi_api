package types

type ControllerServiceApiDTO struct {
	Type   string    `json:"type"`
	Bundle BundleDTO `json:"bundle"`
}
