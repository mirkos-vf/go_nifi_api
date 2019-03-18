package types

type BulletinEntity struct {
	Id          int64       `json:"id"`
	GroupId     string      `json:"groupid"`
	SourceId    string      `json:"sourceid"`
	Timestamp   string      `json:"timestamp"`
	NodeAddress string      `json:"nodeaddress"`
	CanRead     bool        `json:"canread"`
	Bulletin    BulletinDTO `json:"bulletin"`
}
