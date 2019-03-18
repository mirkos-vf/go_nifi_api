package main

type ProcessorStatusSnapshotEntity struct {
	Id                      string                     `json:"id"`
	ProcessorStatusSnapshot ProcessorStatusSnapshotDTO `json:"processorstatussnapshot"`
	CanRead                 bool                       `json:"canread"`
}
