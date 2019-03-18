package types

type VersionedFlow struct {
	Link              Link        `json:"link"`
	Identifier        string      `json:"identifier"`
	Name              string      `json:"name"`
	Description       string      `json:"description"`
	BucketIdentifier  string      `json:"bucketidentifier"`
	BucketName        string      `json:"bucketname"`
	CreatedTimestamp  int64       `json:"createdtimestamp"`
	ModifiedTimestamp int64       `json:"modifiedtimestamp"`
	Type              string      `json:"type"`
	Permissions       Permissions `json:"permissions"`
	VersionCount      int64       `json:"versioncount"`
}
