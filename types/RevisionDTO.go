package types

type RevisionDTO struct {
	ClientId     string `json:"clientid"`
	Version      int    `json:"version"`
	LastModifier string `json:"lastmodifier"`
}
