package types

type Bucket struct {
	Link             Link        `json:"link"`
	Identifier       string      `json:"identifier"`
	Name             string      `json:"name"`
	CreatedTimestamp int64       `json:"createdtimestamp"`
	Description      string      `json:"description"`
	Permissions      Permissions `json:"permissions"`
}
