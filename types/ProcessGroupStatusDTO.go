package types

type ProcessGroupStatusDTO struct {
	Id                 string                              `json:"id"`
	Name               string                              `json:"name"`
	StatsLastRefreshed string                              `json:"statslastrefreshed"`
	AggregateSnapshot  ProcessGroupStatusSnapshotDTO       `json:"aggregatesnapshot"`
	NodeSnapshots      []NodeProcessGroupStatusSnapshotDTO `json:"nodesnapshots"`
}
