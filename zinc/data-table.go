package zinc

import (
	"encoding/base64"
	"encoding/json"
	"html/template"
	"sort"
	"strconv"
	"strings"
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
	IconSize       int    `json:"iconSize,omitempty"`
	IconStyle      string `json:"iconStyle,omitempty"` // round,tile,depth,padded,blink,squared | Accept multiple values separated by comma
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

// DataTableConfig configures the <zn-data-table> web component. Each field
// maps to an HTML attribute on the rendered tag
//
// See https://zinc.style/components/data-table/ for visual examples and the
// full request/response data schemas.
// example:
//
//	<zn-data-table {{$.DataTableConfig.Attrs}}></zn-data-table>
type DataTableConfig struct {
	// DataUri is the URL the table fetches rows from.
	DataUri string `json:"data-uri,omitempty"`

	// Data is an inline JSON-encoded list of rows.
	Data string `json:"data,omitempty"`

	// Method is the HTTP verb for data requests: "GET" or "POST".
	// Component default is "POST" when empty.
	Method string `json:"method,omitempty"`

	// Headers defines the columns.
	Headers HeaderConfigs `json:"headers,omitempty"`

	// HideHeaders is a JSON array of column keys whose header text should
	// be hidden while the column content stays visible (e.g. action
	// columns): `["actions"]`.
	HideHeaders []string `json:"hide-headers,omitempty"`

	// UnsortableHeaders is a JSON array of column keys that must not be
	// sortable, overriding sortable:true on their HeaderConfig.
	UnsortableHeaders string `json:"unsortable-headers,omitempty"`

	// WideColumn is the column key that should expand to fill remaining
	// horizontal space. Only one column should be designated wide.
	WideColumn string `json:"wide-column,omitempty'"`

	// HideColumns is a JSON array of column keys to remove from the table
	// entirely (unlike HideHeaders, which only hides the label).
	HideColumns []string `json:"hide-columns,omitempty"`

	// LocalSort enables client-side sorting. The full result set is sorted
	// in the browser without re-fetching. Best for small datasets.
	LocalSort *bool `json:"local-sort,omitempty"`

	// SortColumn is the initial column key to sort by.
	SortColumn string `json:"sort-column,omitempty"`

	// SortDirection is the initial sort direction: "asc" or "desc".
	// Component default is "asc".
	SortDirection string `json:"sort-direction,omitempty"`

	// GroupBy is a column key. When set, rows are split into separate
	// sub-tables, one per distinct value of that column.
	GroupBy string `json:"group-by,omitempty"`

	// Groups is a comma-separated list of group values to display when
	// GroupBy is set.
	Groups string `json:"groups,omitempty"`

	// Search is the initial search query passed to the data endpoint.
	// Changes to the zn-data-table-search slot component update this at runtime.
	Search string `json:"search,omitempty"`

	// Filter is the initial filter string passed to the data endpoint -
	// typically a serialized query produced by zn-data-table-filter.
	Filter string `json:"filter,omitempty"`

	// Captions is the table caption / heading text.
	Captions *string `json:"captions,omitempty"`

	// EmptyStateCaption is the text shown when the result set is empty
	// and no custom empty-state slot is supplied.
	EmptyStateCaption *string `json:"empty-state-caption,omitempty"`

	// EmptyStateIcon is the icon rendered alongside EmptyStateCaption.
	// Component default is "data_alert".
	EmptyStateIcon *string `json:"empty-state-icon,omitempty"`

	// Unsortable disables column sorting on the entire table.
	Unsortable *bool `json:"unsortable,omitempty"`

	// PerPageSize is the initial number of rows requested per page.
	// Default is 10 when unset.
	PerPageSize int `json:"per-page-size,omitempty"`

	// HidePagination hides the pagination footer even when total > perPage.
	HidePagination *bool `json:"hide-pagination,omitempty"`

	// HideCheckboxes disables row selection - both the per-row checkboxes
	// and the select-all action are hidden.
	HideCheckboxes *bool `json:"hide-checkboxes,omitempty"`

	// Standalone renders the table without its container wrapper. Use when
	// embedding inside another zinc component that already provides padding/borders.
	Standalone *bool `json:"standalone,omitempty"`

	// NoInitialLoad suppresses the data fetch that normally fires when
	// the component connects.
	NoInitialLoad *bool `json:"no-initial-load,omitempty"`
}

// Attrs renders this config as a space-separated HTML attribute for insertion
// inside a tag in a template.
func (c DataTableConfig) Attrs() template.HTMLAttr {
	return HTMLAttrs(c)
}

type HeaderConfigs []HeaderConfig

func (h HeaderConfigs) String() string {
	b, _ := json.Marshal(h)
	return string(b)
}

func (h HeaderConfigs) AttrValue() string {
	return h.String()
}

func (h HeaderConfigs) Attr() template.HTMLAttr {
	return template.HTMLAttr(h.String())
}

type HeaderConfig struct {
	Key   string `json:"key"`
	Label string `json:"label"`

	Type           string `json:"type,omitempty"`
	RenderTemplate string `json:"renderTemplate,omitempty"`
	CellTemplate   *Cell  `json:"cellTemplate,omitempty"`
	IfEmpty        *Cell  `json:"ifEmpty,omitempty"`
	Required       *bool  `json:"required,omitempty"`
	Default        *bool  `json:"default,omitempty"`
	Sortable       *bool  `json:"sortable,omitempty"`
	Filterable     *bool  `json:"filterable,omitempty"`
	HideHeader     *bool  `json:"hideHeader,omitempty"`
	HideColumn     *bool  `json:"hideColumn,omitempty"`
	Secondary      *bool  `json:"secondary,omitempty"`
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

type DataTableFilterType string

const (
	DataTableFilterTypeBoolean  DataTableFilterType = "boolean"
	DataTableFilterTypeDate     DataTableFilterType = "date"
	DataTableFilterTypeDateTime DataTableFilterType = "dateTime"
	DataTableFilterTypeNumber   DataTableFilterType = "number"
	DataTableFilterTypeSelect   DataTableFilterType = ""
	DataTableFilterTypeString   DataTableFilterType = ""
)

type DataTableFilterOptions map[string]any

type DataTableFilterOperator string

const (
	DataTableFilterOperatorEq              DataTableFilterOperator = "eq"              // Equals
	DataTableFilterOperatorNeq             DataTableFilterOperator = "neq"             // Not Equals
	DataTableFilterOperatorEqi             DataTableFilterOperator = "eqi"             // Equals (Insensitive)
	DataTableFilterOperatorNeqi            DataTableFilterOperator = "neqi"            // Not Equals (Insensitive)
	DataTableFilterOperatorBefore          DataTableFilterOperator = "before"          // Was Before
	DataTableFilterOperatorAfter           DataTableFilterOperator = "after"           // Was After
	DataTableFilterOperatorIn              DataTableFilterOperator = "in"              // In
	DataTableFilterOperatorNin             DataTableFilterOperator = "nin"             // Not In
	DataTableFilterOperatorMatchPhrasePre  DataTableFilterOperator = "matchphrasepre"  // Match Phrase Prefix
	DataTableFilterOperatorNMatchPhrasePre DataTableFilterOperator = "nmatchphrasepre" // Does Not Match Phrase Prefix
	DataTableFilterOperatorMatchPhrase     DataTableFilterOperator = "matchphrase"     // Match Phrase
	DataTableFilterOperatorNMatchPhrase    DataTableFilterOperator = "nmatchphrase"    // Does Not Match Phrase
	DataTableFilterOperatorMatch           DataTableFilterOperator = "match"           // Match
	DataTableFilterOperatorNMatch          DataTableFilterOperator = "nmatch"          // Does Not Match
	DataTableFilterOperatorContains        DataTableFilterOperator = "contains"        // Contains
	DataTableFilterOperatorDoesNotContain  DataTableFilterOperator = "doesnotcontain"  // Does Not Contain
	DataTableFilterOperatorStarts          DataTableFilterOperator = "starts"          // Starts With
	DataTableFilterOperatorNStarts         DataTableFilterOperator = "nstarts"         // Does Not Start With
	DataTableFilterOperatorEnds            DataTableFilterOperator = "ends"            // Ends With
	DataTableFilterOperatorNEnds           DataTableFilterOperator = "nends"           // Does Not End With
	DataTableFilterOperatorWild            DataTableFilterOperator = "wild"            // Wildcard Match
	DataTableFilterOperatorNWild           DataTableFilterOperator = "nwild"           // Does Not Match Wildcard
	DataTableFilterOperatorLike            DataTableFilterOperator = "like"            // Like Match With
	DataTableFilterOperatorNLike           DataTableFilterOperator = "nlike"           // Does Not Like Match With
	DataTableFilterOperatorFuzzy           DataTableFilterOperator = "fuzzy"           // Fuzzy Match With
	DataTableFilterOperatorNFuzzy          DataTableFilterOperator = "nfuzzy"          // Does Not Match Fuzzy With
	DataTableFilterOperatorGte             DataTableFilterOperator = "gte"             // Greater Than or Equals
	DataTableFilterOperatorGt              DataTableFilterOperator = "gt"              // Greater Than
	DataTableFilterOperatorLt              DataTableFilterOperator = "lt"              // Less Than
	DataTableFilterOperatorLte             DataTableFilterOperator = "lte"             // Less Than or Equals
)

// DataTableFilterDateSubmitFormat controls how date and dateTime filter
// values are serialized when the query is submitted.
type DataTableFilterDateSubmitFormat string

const (
	// DataTableFilterDateSubmitFormatIso emits RFC 3339 / ISO 8601 strings
	// (e.g. "2026-06-09T16:05:00Z").
	DataTableFilterDateSubmitFormatIso DataTableFilterDateSubmitFormat = "iso"

	// DataTableFilterDateSubmitFormatTimestamp emits a Unix timestamp in seconds since epoch.
	DataTableFilterDateSubmitFormatTimestamp DataTableFilterDateSubmitFormat = "timestamp"

	// DataTableFilterDateSubmitFormatLegacy emits whatever format the current
	// system produces. Default.
	DataTableFilterDateSubmitFormatLegacy DataTableFilterDateSubmitFormat = "legacy"
)

type DataTableFilterItem struct {
	Id                string                          `json:"id"`
	Title             string                          `json:"name"`
	Type              DataTableFilterType             `json:"type,omitempty"`
	Options           DataTableFilterOptions          `json:"options,omitempty"`
	Operators         []DataTableFilterOperator       `json:"operators"`
	DateSubmitFormat  DataTableFilterDateSubmitFormat `json:"dateSubmitFormat,omitempty"`
	MaxOptionsVisible uint                            `json:"maxOptionsVisible,omitempty"`
}

type DataTableFilterConfig struct {
	Name    string                `json:"name,omitempty"`
	Filters []DataTableFilterItem `json:"filters,omitempty"`
	Value   string                `json:"value,omitempty"`
	Values  []string              `json:"values,omitempty"`
}

func (f DataTableFilterConfig) String() string {
	b, _ := json.Marshal(f)
	return string(b)
}

func (f DataTableFilterConfig) Attrs() template.HTMLAttr {
	return HTMLAttrs(f)
}

type DataTableFilterRequest []DataTableFilterRequestItem

type DataTableFilterRequestItem struct {
	Key        string                  `json:"key"`
	Comparator DataTableFilterOperator `json:"comparator"`
	Value      any                     `json:"value"`
}

func DecodeFilterRequest(s string) (DataTableFilterRequest, error) {
	var v DataTableFilterRequest
	dec := json.NewDecoder(base64.NewDecoder(base64.StdEncoding, strings.NewReader(s)))
	return v, dec.Decode(&v)
}
