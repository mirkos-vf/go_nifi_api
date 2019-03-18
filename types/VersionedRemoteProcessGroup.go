package main

type VersionedRemoteProcessGroup struct {
	Identifier            string        `json:"identifier"`
	Name                  string        `json:"name"`
	Comments              string        `json:"comments"`
	Position              Position      `json:"position"`
	TargetUri             string        `json:"targeturi"`
	TargetUris            string        `json:"targeturis"`
	CommunicationsTimeout string        `json:"communicationstimeout"`
	YieldDuration         string        `json:"yieldduration"`
	TransportProtocol     string        `json:"transportprotocol"`
	LocalNetworkInterface string        `json:"localnetworkinterface"`
	ProxyHost             string        `json:"proxyhost"`
	ProxyPort             int32         `json:"proxyport"`
	ProxyUser             string        `json:"proxyuser"`
	InputPorts            []InputPorts  `json:"inputports"`
	OutputPorts           []OutputPorts `json:"outputports"`
	ComponentType         string        `json:"componenttype"`
	GroupIdentifier       string        `json:"groupidentifier"`
}
