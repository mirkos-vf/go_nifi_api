package types

type ClusterDTO struct {
	Nodes     []NodeDTO `json:"nodes"`
	Generated string    `json:"generated"`
}
