package types

type TemplateDTO struct {
	Uri             string         `json:"uri"`
	Id              string         `json:"id"`
	GroupId         string         `json:"groupId"`
	Name            string         `json:"name"`
	Description     string         `json:"description"`
	Timestamp       string         `json:"timestamp"`
	EncodingVersion string         `json:"encodingVersion"`
	Snippet         FlowSnippetDTO `json:"snippet"`
}
