package types

type PortEntity struct {
	Revision                     RevisionDTO      `json:"revision"`
	Id                           string           `json:"id"`
	Uri                          string           `json:"uri"`
	Position                     PositionDTO      `json:"position"`
	Permissions                  PermissionsDTO   `json:"permissions"`
	Bulletins                    []BulletinEntity `json:"bulletins"`
	DisconnectedNodeAcknowledged bool             `json:"disconnected_node_acknowledged"`
	Component                    PortDTO          `json:"component"`
	Status                       PortStatusDTO    `json:"status"`
	PortType                     string           `json:"port_type"`
	OperatePermissions           PermissionsDTO   `json:"operate_permissions"`
}
