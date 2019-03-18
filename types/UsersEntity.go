package types

type UsersEntity struct {
	Generated string       `json:"generated"`
	Users     []UserEntity `json:"users"`
}
