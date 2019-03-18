package types

type RemoteProcessGroupStatusSnapshotEntity struct {
	Id                               string                              `json:"id"`
	RemoteProcessGroupStatusSnapshot RemoteProcessGroupStatusSnapshotDTO `json:"remoteprocessgroupstatussnapshot"`
	CanRead                          bool                                `json:"canread"`
}
