package types

type UserGroupDTO struct {
	Id                   string               `json:"id"`
	VersionedComponentId string               `json:"versionedcomponentid"`
	ParentGroupId        string               `json:"parentgroupid"`
	Position             PositionDTO          `json:"position"`
	Identity             string               `json:"identity"`
	Configurable         bool                 `json:"configurable"`
	Users                []TenantEntity       `json:"users"`
	AccessPolicies       []AccessPolicyEntity `json:"accesspolicies"`
}
