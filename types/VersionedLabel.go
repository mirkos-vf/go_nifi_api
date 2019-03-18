package types

type VersionedLabel struct {
	Identifier      string   `json:"identifier"`
	Name            string   `json:"name"`
	Comments        string   `json:"comments"`
	Position        Position `json:"position"`
	Label           string   `json:"label"`
	Width           float32  `json:"width"`
	Height          float32  `json:"height"`
	Style           object   `json:"style"`
	ComponentType   string   `json:"componenttype"`
	GroupIdentifier string   `json:"groupidentifier"`
}
