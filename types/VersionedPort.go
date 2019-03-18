package types

type VersionedPort struct {
	Identifier                       string   `json:"identifier"`
	Name                             string   `json:"name"`
	Comments                         string   `json:"comments"`
	Position                         Position `json:"position"`
	Type                             string   `json:"type"`
	ConcurrentlySchedulableTaskCount int32    `json:"concurrentlyschedulabletaskcount"`
	ComponentType                    string   `json:"componenttype"`
	GroupIdentifier                  string   `json:"groupidentifier"`
}
