package types

type VersionedFunnel struct {
	Identifier      string   `json:"identifier"`
	Name            string   `json:"name"`
	Comments        string   `json:"comments"`
	Position        Position `json:"position"`
	ComponentType   string   `json:"componenttype"`
	GroupIdentifier string   `json:"groupidentifier"`
}
