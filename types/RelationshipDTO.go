package types

type RelationshipDTO struct {
	Name          string `json:"name"`
	Description   string `json:"description"`
	AutoTerminate bool   `json:"autoterminate"`
}
