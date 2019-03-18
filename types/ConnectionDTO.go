package main

type ConnectionDTO struct {
	Id                            string         `json:"id"`
	VersionedComponentId          string         `json:"versionedcomponentid"`
	ParentGroupId                 string         `json:"parentgroupid"`
	Position                      PositionDTO    `json:"position"`
	Source                        ConnectableDTO `json:"source"`
	Destination                   ConnectableDTO `json:"destination"`
	Name                          string         `json:"name"`
	LabelIndex                    int32          `json:"labelindex"`
	GetzIndex                     int32          `json:"getzindex"`
	SelectedRelationships         []string       `json:"selectedrelationships"`
	AvailableRelationships        []string       `json:"availablerelationships"`
	BackPressureObjectThreshold   int32          `json:"backpressureobjectthreshold"`
	BackPressureDataSizeThreshold string         `json:"backpressuredatasizethreshold"`
	FlowFileExpiration            string         `json:"flowfileexpiration"`
	Prioritizers                  []string       `json:"prioritizers"`
	Bends                         []PositionDTO  `json:"bends"`
	LoadBalanceStrategy           string         `json:"loadbalancestrategy"`
	LoadBalancePartitionAttribute string         `json:"loadbalancepartitionattribute"`
	LoadBalanceCompression        string         `json:"loadbalancecompression"`
	LoadBalanceStatus             string         `json:"loadbalancestatus"`
}
