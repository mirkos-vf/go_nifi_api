package types

type SnippetDTO struct {
	Id                  string      `json:"id"`
	Uri                 string      `json:"uri"`
	ParentGroupId       string      `json:"parentgroupid"`
	ProcessGroups       interface{} `json:"processgroups"`
	RemoteProcessGroups interface{} `json:"remoteprocessgroups"`
	Processors          interface{} `json:"processors"`
	InputPorts          interface{} `json:"inputports"`
	OutputPorts         interface{} `json:"outputports"`
	Connections         interface{} `json:"connections"`
	Labels              interface{} `json:"labels"`
	Funnels             interface{} `json:"funnels"`
}
