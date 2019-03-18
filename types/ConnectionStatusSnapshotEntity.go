package types

type ConnectionStatusSnapshotEntity struct {
	Id                       string                      `json:"id"`
	ConnectionStatusSnapshot ConnectionStatusSnapshotDTO `json:"connectionstatussnapshot"`
	CanRead                  bool                        `json:"canread"`
}
