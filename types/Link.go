package types

type Link struct {
	Type       string     `json:"type"`
	Rels       []string   `json:"rels"`
	Rel        string     `json:"rel"`
	Uri        string     `json:"uri"`
	Params     object     `json:"params"`
	Title      string     `json:"title"`
	UriBuilder UriBuilder `json:"uribuilder"`
}
