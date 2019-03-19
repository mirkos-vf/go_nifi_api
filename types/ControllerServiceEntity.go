package types

type ControllerServiceEntity struct {
	Revision                     RevisionDTO                `json:"revision"`
	Id                           string                     `json:"id"`
	Uri                          string                     `json:"uri"`
	Position                     PositionDTO                `json:"position"`
	Permissions                  PermissionsDTO             `json:"permissions"`
	Bulletins                    []BulletinEntity           `json:"bulletins"`
	DisconnectedNodeAcknowledged bool                       `json:"disconnectednodeacknowledged"`
	ParentGroupId                string                     `json:"parentgroupid"`
	Component                    ControllerServiceDTO       `json:"component"`
	OperatePermissions           PermissionsDTO             `json:"operatepermissions"`
	Status                       ControllerServiceStatusDTO `json:"status"`
}
