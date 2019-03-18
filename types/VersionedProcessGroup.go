package main

type VersionedProcessGroup struct {
	Identifier               string                        `json:"identifier"`
	Name                     string                        `json:"name"`
	Comments                 string                        `json:"comments"`
	Position                 Position                      `json:"position"`
	ProcessGroups            []VersionedProcessGroup       `json:"processgroups"`
	RemoteProcessGroups      []VersionedRemoteProcessGroup `json:"remoteprocessgroups"`
	Processors               []VersionedProcessor          `json:"processors"`
	InputPorts               []VersionedPort               `json:"inputports"`
	OutputPorts              []VersionedPort               `json:"outputports"`
	Connections              []VersionedConnection         `json:"connections"`
	Labels                   []VersionedLabel              `json:"labels"`
	Funnels                  []VersionedFunnel             `json:"funnels"`
	ControllerServices       []VersionedControllerService  `json:"controllerservices"`
	VersionedFlowCoordinates VersionedFlowCoordinates      `json:"versionedflowcoordinates"`
	Variables                object                        `json:"variables"`
	ComponentType            string                        `json:"componenttype"`
	GroupIdentifier          string                        `json:"groupidentifier"`
}
