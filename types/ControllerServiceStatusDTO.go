package types

type ControllerServiceStatusDTO struct {
	RunStatus         string `json:"run_status"`
	ValidationStatus  string `json:"validation_status"`
	ActiveThreadCount int32  `json:"active_thread_count"`
}
