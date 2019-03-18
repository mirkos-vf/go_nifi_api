package types

type SnippetEntity struct {
	Snippet                      SnippetDTO `json:"snippet"`
	DisconnectedNodeAcknowledged bool       `json:"disconnectednodeacknowledged"`
}
