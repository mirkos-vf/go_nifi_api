package types

type VersionedFlowSnapshotMetadata struct {
	Link             Link   `json:"link"`
	BucketIdentifier string `json:"bucketidentifier"`
	FlowIdentifier   string `json:"flowidentifier"`
	Version          int32  `json:"version"`
	Timestamp        int64  `json:"timestamp"`
	Author           string `json:"author"`
	Comments         string `json:"comments"`
}
