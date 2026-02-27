package zinc

import (
	"encoding/json"
	"html/template"
	"sort"
	"strconv"
)

type DataTable struct {
	Rows []Row `json:"rows"`

	PerPage int64 `json:"perPage"`
	Total   int64 `json:"total"`
	Page    int64 `json:"page"`
}

type Row struct {
	Id      string   `json:"id"`
	Uri     string   `json:"uri,omitempty"`
	Target  string   `json:"target,omitempty"`
	Cells   []Cell   `json:"cells,omitempty"`
	Actions []Action `json:"actions,omitempty"`
}

type Cell struct {
	Text           string `json:"text"`
	Column         string `json:"column"`
	Color          string `json:"color,omitempty"`
	Style          string `json:"style,omitempty"`
	IconSrc        string `json:"iconSrc,omitempty"`
	IconColor      string `json:"iconColor,omitempty"`
	HoverContent   string `json:"hoverContent,omitempty"`
	HoverPlacement string `json:"hoverPlacement,omitempty"`
	ChipColor      string `json:"chipColor,omitempty"`
	SortValue      string `json:"sortValue,omitempty"`
	Gaid           string `json:"gaid,omitempty"`
	Uri            string `json:"uri,omitempty"`
	Target         string `json:"target,omitempty"`
	Copyable       bool   `json:"copyable,omitempty"`
}

type Action struct {
	Text           string `json:"text"`
	Uri            string `json:"uri"`
	Target         string `json:"target,omitempty"`
	Gaid           string `json:"gaid,omitempty"`
	Color          string `json:"color,omitempty"`
	IconSrc        string `json:"iconSrc,omitempty"`
	ConfirmType    string `json:"confirmType,omitempty"`
	ConfirmTitle   string `json:"confirmTitle,omitempty"`
	ConfirmContent string `json:"confirmContent,omitempty"`
}

type DataRequest struct {
	Page          int64  `json:"page,omitempty"`
	PerPage       int64  `json:"perPage,omitempty"`
	SortColumn    string `json:"sortColumn,omitempty"`
	SortDirection string `json:"sortDirection,omitempty"`
	Filter        string `json:"filter,omitempty"`
	Search        string `json:"search,omitempty"`
}

type HeaderConfigs []HeaderConfig

func (h HeaderConfigs) String() string {
	b, _ := json.Marshal(h)
	return string(b)
}

func (h HeaderConfigs) Attr() template.HTMLAttr {
	return template.HTMLAttr(h.String())
}

type HeaderConfig struct {
	Key   string `json:"key"`
	Label string `json:"label"`

	Required   *bool `json:"required,omitempty"`
	Default    *bool `json:"default,omitempty"`
	Sortable   *bool `json:"sortable,omitempty"`
	Filterable *bool `json:"filterable,omitempty"`
	HideHeader *bool `json:"hideHeader,omitempty"`
	HideColumn *bool `json:"hideColumn,omitempty"`
	Secondary  *bool `json:"secondary,omitempty"`
}

// SortDataTable sorts dataTable rows in place based on the request's sort column and direction.
func (dataTable DataTable) SortDataTable(req DataRequest) {
	if req.SortColumn == "" {
		return
	}

	sort.Slice(dataTable.Rows, func(i, j int) bool {
		var a, b string
		for _, cell := range dataTable.Rows[i].Cells {
			if cell.Column == req.SortColumn {
				a = cell.SortValue
				if a == "" {
					a = cell.Text
				}
				break
			}
		}
		for _, cell := range dataTable.Rows[j].Cells {
			if cell.Column == req.SortColumn {
				b = cell.SortValue
				if b == "" {
					b = cell.Text
				}
				break
			}
		}

		aNum, aErr := strconv.ParseFloat(a, 64)
		bNum, bErr := strconv.ParseFloat(b, 64)
		if aErr == nil && bErr == nil {
			if req.SortDirection == "desc" {
				return aNum > bNum
			}
			return aNum < bNum
		}

		if req.SortDirection == "desc" {
			return a > b
		}
		return a < b
	})
}
