package types

type Permissions struct {
	CanRead   bool `json:"canread"`
	CanWrite  bool `json:"canwrite"`
	CanDelete bool `json:"candelete"`
}
