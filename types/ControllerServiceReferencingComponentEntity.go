package types

type ControllerServiceReferencingComponentEntity struct {
	Revision                     RevisionDTO                              `json:"revision"`
	Id                           string                                   `json:"id"`
	Uri                          string                                   `json:"uri"`
	Position                     PositionDTO                              `json:"position"`
	Permissions                  PermissionsDTO                           `json:"permissions"`
	Bulletins                    []BulletinEntity                         `json:"bulletins"`
	DisconnectedNodeAcknowledged bool                                     `json:"disconnectednodeacknowledged"`
	Component                    ControllerServiceReferencingComponentDTO `json:"component"`
	OperatePermissions           PermissionsDTO                           `json:"operatepermissions"`
}
