package main

type ProcessGroupStatusSnapshotEntity struct {
	Id                         string                        `json:"id"`
	ProcessGroupStatusSnapshot ProcessGroupStatusSnapshotDTO `json:"processgroupstatussnapshot"`
	CanRead                    bool                          `json:"canread"`
}
