package main

type PortStatusDTO struct {
	Id                 string                      `json:"id"`
	GroupId            string                      `json:"group_id"`
	Name               string                      `json:"name"`
	Transmitting       bool                        `json:"transmitting"`
	RunStatus          string                      `json:"run_status"`
	StatsLastRefreshed string                      `json:"stats_last_refreshed"`
	AggregateSnapshot  PortStatusSnapshotDTO       `json:"aggregate_snapshot"`
	NodeSnapshots      []NodePortStatusSnapshotDTO `json:"node_snapshots"`
}
