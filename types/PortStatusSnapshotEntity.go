package main

type PortStatusSnapshotEntity struct {
	Id                 string                `json:"id"`
	PortStatusSnapshot PortStatusSnapshotDTO `json:"portStatusSnapshot"`
	CanRead            bool                  `json:"canread"`
}
