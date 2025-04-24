package zinc

import (
	"encoding/json"
	"html/template"
)

type Table struct {
	Header []string   `json:"header,omitempty"`
	Items  []TableRow `json:"items,omitempty"`
}

func (t Table) String() string {
	b, _ := json.Marshal(t)
	return string(b)
}

func (t Table) Attr() template.HTMLAttr {
	return template.HTMLAttr(t.String())
}

type TableRow struct {
	Actions          []MenuAction `json:"actions,omitempty"`
	ActionsPlacement string       `json:"actions-placement,omitempty"`
	URI              string       `json:"uri,omitempty"`
	Target           string       `json:"target,omitempty"`
	Caption          string       `json:"caption,omitempty"`
	Summary          string       `json:"summary,omitempty"`
	Icon             string       `json:"icon,omitempty"`
	ID               string       `json:"id,omitempty"`
	Data             []any        `json:"data,omitempty"`
}

type MenuAction struct {
	Confirm *Confirm `json:"confirm,omitempty"`
	Target  string   `json:"target,omitempty"`
	Path    string   `json:"path,omitempty"`
	Title   string   `json:"title,omitempty"`
	Style   string   `json:"style,omitempty"`
}

type Confirm struct {
	Type    string `json:"type,omitempty"`
	Trigger string `json:"trigger,omitempty"` // Unique identifier
	Caption string `json:"caption,omitempty"`
	Title   string `json:"title,omitempty"` // Trigger text
	Content string `json:"content,omitempty"`
	Action  string `json:"action,omitempty"` // URL
}

type TableAccordion struct {
	Title    string
	Table    Table
	Expanded bool
}
