package zinc

import (
	"html/template"
)

const (
	ReportFormatNumber    = "number"
	ReportFormatDecimal   = "decimal"
	ReportFormatPercent   = "percent"
	ReportFormatCurrency  = "currency"
	ReportFormatFinancial = "financial"
	ReportFormatText      = "text"

	ReportStyleNone       = "none"
	ReportStyleBreakeven  = "breakeven"
	ReportStyleIBreakeven = "ibreakeven"
	ReportStyleBold       = "bold"
)

// ReportTable is the response for <zn-report-table>.
type ReportTable struct {
	Dimension    *ReportDimension      `json:"dimension,omitempty"`
	Groups       []ReportGroup         `json:"groups"`
	Rows         []ReportRow           `json:"rows"`
	Totals       map[string]float64    `json:"totals,omitempty"`
	Currency     string                `json:"currency,omitempty"`
	GroupBy      string                `json:"groupBy,omitempty"`
	GroupBys     []ReportOption        `json:"groupBys,omitempty"`
	FilterFields []DataTableFilterItem `json:"filterFields,omitempty"`

	Error *DataTableError `json:"error,omitempty"`
}

type ReportDimension struct {
	Label string `json:"label"`
}

type ReportGroup struct {
	Label   string         `json:"label"`
	Columns []ReportColumn `json:"columns"`
}

// ReportColumn reads its value from Formula, else Value, else the dotted path Key.
type ReportColumn struct {
	Key       string   `json:"key"`
	Label     string   `json:"label"`
	Format    string   `json:"format,omitempty"` // ReportFormat*
	Decimals  *int     `json:"decimals,omitempty"`
	Style     string   `json:"style,omitempty"` // ReportStyle*
	Threshold *float64 `json:"threshold,omitempty"`
	Summable  *bool    `json:"summable,omitempty"`
	Currency  string   `json:"currency,omitempty"`
	Value     string   `json:"value,omitempty"`
	Formula   string   `json:"formula,omitempty"`
}

type ReportRow struct {
	Id     string         `json:"id"`
	Label  string         `json:"label"`
	Uri    string         `json:"uri,omitempty"`
	Target string         `json:"target,omitempty"`
	Values map[string]any `json:"values,omitempty"`
	// Links is keyed by column key.
	Links map[string]ReportLink `json:"links,omitempty"`
}

type ReportLink struct {
	Uri    string `json:"uri"`
	Target string `json:"target,omitempty"`
}

type ReportOption struct {
	Value    string `json:"value"`
	Label    string `json:"label"`
	Category string `json:"category,omitempty"`
}

type ReportDatePreset struct {
	Label string `json:"label"`
	From  string `json:"from"`
	To    string `json:"to"`
}

// ReportRequest is the body <zn-report-table> sends to its data-uri.
type ReportRequest struct {
	Type    string `json:"type,omitempty"`
	GroupBy string `json:"groupBy,omitempty"`
	From    string `json:"from,omitempty"`
	To      string `json:"to,omitempty"`
	Filter  string `json:"filter,omitempty"`
}

// Filters decodes Filter; it returns nil, nil when Filter is empty.
func (r ReportRequest) Filters() (DataTableFilterRequest, error) {
	if r.Filter == "" {
		return nil, nil
	}
	return DecodeFilterRequest(r.Filter)
}

// ReportTableConfig configures the <zn-report-table> web component. Each field
// maps to an HTML attribute on the rendered tag.
// example:
//
//	<zn-report-table {{$.ReportTableConfig.Attrs}}></zn-report-table>
type ReportTableConfig struct {
	DataUri           string                `json:"data-uri,omitempty"`
	Method            string                `json:"method,omitempty"`
	Caption           string                `json:"caption,omitempty"`
	Type              string                `json:"type,omitempty"`
	Types             []ReportOption        `json:"types,omitempty"`
	GroupBy           string                `json:"group-by,omitempty"`
	GroupBys          []ReportOption        `json:"group-bys,omitempty"`
	From              string                `json:"from,omitempty"`
	To                string                `json:"to,omitempty"`
	MaxDate           string                `json:"max-date,omitempty"`
	Currency          string                `json:"currency,omitempty"`
	FilterFields      []DataTableFilterItem `json:"filter-fields,omitempty"`
	Filter            string                `json:"filter,omitempty"`
	Presets           []ReportDatePreset    `json:"presets,omitempty"`
	Locale            string                `json:"locale,omitempty"`
	SearchPlaceholder string                `json:"search-placeholder,omitempty"`
	DateNote          string                `json:"date-note,omitempty"`
	Sharable          bool                  `json:"sharable,omitempty"`
}

func (c ReportTableConfig) Attrs() template.HTMLAttr {
	return HTMLAttrs(c)
}
