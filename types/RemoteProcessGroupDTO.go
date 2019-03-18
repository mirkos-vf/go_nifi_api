package types

type RemoteProcessGroupDTO struct {
	Id                            string                        `json:"id"`
	VersionedComponentId          string                        `json:"versionedcomponentid"`
	ParentGroupId                 string                        `json:"parentgroupid"`
	Position                      PositionDTO                   `json:"position"`
	TargetUri                     string                        `json:"targeturi"`
	TargetUris                    string                        `json:"targeturis"`
	TargetSecure                  bool                          `json:"targetsecure"`
	Name                          string                        `json:"name"`
	Comments                      string                        `json:"comments"`
	CommunicationsTimeout         string                        `json:"communicationstimeout"`
	YieldDuration                 string                        `json:"yieldduration"`
	TransportProtocol             string                        `json:"transportprotocol"`
	LocalNetworkInterface         string                        `json:"localnetworkinterface"`
	ProxyHost                     string                        `json:"proxyhost"`
	ProxyPort                     int32                         `json:"proxyport"`
	ProxyUser                     string                        `json:"proxyuser"`
	ProxyPassword                 string                        `json:"proxypassword"`
	AuthorizationIssues           []string                      `json:"authorizationissues"`
	ValidationErrors              []string                      `json:"validationerrors"`
	Transmitting                  bool                          `json:"transmitting"`
	InputPortCount                int32                         `json:"inputportcount"`
	OutputPortCount               int32                         `json:"outputportcount"`
	ActiveRemoteInputPortCount    int32                         `json:"activeremoteinputportcount"`
	InactiveRemoteInputPortCount  int32                         `json:"inactiveremoteinputportcount"`
	ActiveRemoteOutputPortCount   int32                         `json:"activeremoteoutputportcount"`
	InactiveRemoteOutputPortCount int32                         `json:"inactiveremoteoutputportcount"`
	FlowRefreshed                 string                        `json:"flowrefreshed"`
	Contents                      RemoteProcessGroupContentsDTO `json:"contents"`
}
