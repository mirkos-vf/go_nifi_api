package main

type versionedFlowSnapshot struct {
	SnapshotMetadata VersionedFlowSnapshotMetadata `json:"snapshotmetadata"`
	FlowContents     VersionedProcessGroup         `json:"flowcontents"`
	Flow             VersionedFlow                 `json:"flow"`
	Bucket           Bucket                        `json:"bucket"`
	Latest           bool                          `json:"latest"`
}
