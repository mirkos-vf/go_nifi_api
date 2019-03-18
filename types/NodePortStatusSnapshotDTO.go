package types

type NodePortStatusSnapshotDTO struct {
	NodeId         string                `json:"node_id"`
	Address        string                `json:"address"`
	ApiPort        int32                 `json:"api_port"`
	StatusSnapshot PortStatusSnapshotDTO `json:"status_snapshot"`
}
