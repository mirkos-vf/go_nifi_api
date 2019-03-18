package types

type VersionedFlowCoordinates struct {
	RegistryUrl string `json:"registryurl"`
	BucketId    string `json:"bucketid"`
	FlowId      string `json:"flowid"`
	Version     int32  `json:"version"`
	Latest      bool   `json:"latest"`
}
