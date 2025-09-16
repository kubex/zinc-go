package zinc

type DataTable struct {
	Rows []Row `json:"rows"`
}

type Row struct {
	Id      string   `json:"id"`
	Uri     string   `json:"uri,omitempty"`
	Target  string   `json:"target,omitempty"`
	Gaid    string   `json:"gaid,omitempty"`
	Cells   []Cell   `json:"cells,omitempty"`
	Actions []Action `json:"actions,omitempty"`
}

type Cell struct {
	Text           string `json:"text"`
	Heading        string `json:"heading"`
	Color          string `json:"color,omitempty"`
	Style          string `json:"style,omitempty"`
	IconSrc        string `json:"iconSrc,omitempty"`
	IconColor      string `json:"iconColor,omitempty"`
	HoverContent   string `json:"hoverContent,omitempty"`
	HoverPlacement string `json:"hoverPlacement,omitempty"`
	ChipColor      string `json:"chipColor,omitempty"`
	Gaid           string `json:"gaid,omitempty"`
	SortValue      string `json:"sortValue,omitempty"`
	Uri            string `json:"uri,omitempty"`
	Target         string `json:"target,omitempty"`
}

type Action struct {
	Text   string `json:"text"`
	Uri    string `json:"uri"`
	Target string `json:"target,omitempty"`
	Gaid   string `json:"gaid,omitempty"`
}
