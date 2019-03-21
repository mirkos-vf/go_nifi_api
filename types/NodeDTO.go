package types

type NodeDTO struct {
	NodeId              string       `json:"nodeId"`
	Address             string       `json:"address"`
	ApiPort             int32        `json:"api_port"`
	Status              string       `json:"status"`
	Heartbeat           string       `json:"heartbeat"`
	ConnectionRequested string       `json:"connectionRequested"`
	Roles               []string     `json:"roles"`
	ActiveThreadCount   int32        `json:"active_thread_count"`
	Queued              string       `json:"queued"`
	Events              NodeEventDTO `json:"events"`
	NodeStartTime       string       `json:"node_start_time"`
}
