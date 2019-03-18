package main

type ConnectionStatusSnapshotEntity struct {
	Id                       string                      `json:"id"`
	ConnectionStatusSnapshot ConnectionStatusSnapshotDTO `json:"connectionstatussnapshot"`
	CanRead                  bool                        `json:"canread"`
}
