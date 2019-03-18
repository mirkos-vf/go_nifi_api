package types

type UserEntity struct {
	Revision                     RevisionDTO      `json:"revision"`
	Id                           string           `json:"id"`
	Uri                          string           `json:"uri"`
	Position                     PositionDTO      `json:"position"`
	Permissions                  PermissionsDTO   `json:"permissions"`
	Bulletins                    []BulletinEntity `json:"bulletins"`
	DisconnectedNodeAcknowledged bool             `json:"disconnectednodeacknowledged"`
	Component                    UserDTO          `json:"component"`
}
