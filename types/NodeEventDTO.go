package types

type NodeEventDTO struct {
	Timestamp string `json:"timestamp"`
	Category  string `json:"category"`
	Message   string `json:"message"`
}
