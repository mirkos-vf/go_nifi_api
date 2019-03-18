package types

type CreateTemplateRequestEntity struct {
	Name                         string `json:"name"`
	Description                  string `json:"description"`
	SnippetId                    string `json:"snippetid"`
	DisconnectedNodeAcknowledged bool   `json:"disconnectednodeacknowledged"`
}
