package types

type RemoteProcessGroupContentsDTO struct {
	InputPorts  []RemoteProcessGroupPortDTO `json:"inputports"`
	OutputPorts []RemoteProcessGroupPortDTO `json:"outputports"`
}
