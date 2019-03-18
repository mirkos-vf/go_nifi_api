package main

type FlowSnippetDTO struct {
	ProcessGroups       []ProcessGroupDTO       `json:"processgroups"`
	RemoteProcessGroups []RemoteProcessGroupDTO `json:"remoteprocessgroups"`
	Processors          []ProcessorDTO          `json:"processors"`
	InputPorts          []PortDTO               `json:"inputports"`
	OutputPorts         []PortDTO               `json:"outputports"`
	Connections         []ConnectionDTO         `json:"connections"`
	Labels              []LabelDTO              `json:"labels"`
	Funnels             []FunnelDTO             `json:"funnels"`
	ControllerServices  []ControllerServiceDTO  `json:"controllerservices"`
}
