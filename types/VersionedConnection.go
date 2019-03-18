package types

type VersionedConnection struct {
	Identifier                    string               `json:"identifier"`
	Name                          string               `json:"name"`
	Comments                      string               `json:"comments"`
	Position                      Position             `json:"position"`
	Source                        ConnectableComponent `json:"source"`
	Destination                   ConnectableComponent `json:"destination"`
	LabelIndex                    int32                `json:"labelindex"`
	ZIndex                        int64                `json:"zindex"`
	SelectedRelationships         []string             `json:"selectedrelationships"`
	BackPressureObjectThreshold   int64                `json:"backpressureobjectthreshold"`
	BackPressureDataSizeThreshold string               `json:"backpressuredatasizethreshold"`
	FlowFileExpiration            string               `json:"flowfileexpiration"`
	Prioritizers                  []string             `json:"prioritizers"`
	Bends                         []Position           `json:"bends"`
	LoadBalanceStrategy           string               `json:"loadbalancestrategy"`
	PartitioningAttribute         string               `json:"partitioningattribute"`
	LoadBalanceCompression        string               `json:"loadbalancecompression"`
	ComponentType                 string               `json:"componenttype"`
	GroupIdentifier               string               `json:"groupidentifier"`
}
