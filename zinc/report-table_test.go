package zinc

import (
	"encoding/base64"
	"encoding/json"
	"reflect"
	"strings"
	"testing"
)

func TestReportTableJSON(t *testing.T) {
	dec, thr, sum := 2, 0.0, true
	rt := ReportTable{
		Dimension: &ReportDimension{Label: "Publisher"},
		Groups: []ReportGroup{{Label: "G", Columns: []ReportColumn{{
			Key: "rev", Label: "Revenue", Format: ReportFormatCurrency, Decimals: &dec,
			Style: ReportStyleBreakeven, Threshold: &thr, Summable: &sum,
			Currency: "USD", Value: "d1.rev", Formula: "{a}-{b}",
		}}}},
		Rows: []ReportRow{{Id: "1", Label: "A", Uri: "/u", Target: "_blank", Values: map[string]any{"d1": 1},
			Links: map[string]ReportLink{"rev": {Uri: "/u/rev", Target: "slide"}, "d1": {Uri: "/u/d1"}}}},
		Totals:   map[string]float64{"rev": 1},
		Currency: "USD",
		GroupBy:  "pub",
		GroupBys: []ReportOption{{Value: "pub", Label: "Publisher", Category: "P"}},
		FilterFields: []DataTableFilterItem{{Id: "c", Title: "Country", Type: DataTableFilterTypeDate,
			Options: DataTableFilterOptions{"GB": "UK"}, Operators: []DataTableFilterOperator{DataTableFilterOperatorEq},
			DateSubmitFormat: DataTableFilterDateSubmitFormatIso}},
		Error: &DataTableError{Text: "x", Level: "error"},
	}
	b, err := json.Marshal(rt)
	if err != nil {
		t.Fatal(err)
	}
	want := `{"dimension":{"label":"Publisher"},
"groups":[{"label":"G","columns":[{"key":"rev","label":"Revenue","format":"currency","decimals":2,"style":"breakeven","threshold":0,"summable":true,"currency":"USD","value":"d1.rev","formula":"{a}-{b}"}]}],
"rows":[{"id":"1","label":"A","uri":"/u","target":"_blank","values":{"d1":1},"links":{"rev":{"uri":"/u/rev","target":"slide"},"d1":{"uri":"/u/d1"}}}],
"totals":{"rev":1},"currency":"USD","groupBy":"pub",
"groupBys":[{"value":"pub","label":"Publisher","category":"P"}],
"filterFields":[{"id":"c","name":"Country","type":"date","options":{"GB":"UK"},"operators":["eq"],"dateSubmitFormat":"iso"}],
"error":{"text":"x","level":"error"}}`
	var got, exp any
	json.Unmarshal(b, &got)
	json.Unmarshal([]byte(want), &exp)
	if !reflect.DeepEqual(got, exp) {
		t.Fatalf("got %s", b)
	}
}

func TestReportRequestFilters(t *testing.T) {
	r := ReportRequest{Filter: base64.StdEncoding.EncodeToString([]byte(`[{"key":"country","comparator":"eq","value":"GB"}]`))}
	f, err := r.Filters()
	if err != nil || len(f) != 1 || f[0].Key != "country" || f[0].Comparator != DataTableFilterOperatorEq || f[0].Value != "GB" {
		t.Fatalf("got %v, %v", f, err)
	}
	if f, err := (ReportRequest{}).Filters(); f != nil || err != nil {
		t.Fatalf("empty: %v, %v", f, err)
	}
}

func TestReportTableConfigAttrs(t *testing.T) {
	a := string(ReportTableConfig{
		DataUri: "/r", GroupBy: "affiliateFid", Sharable: true, Filter: "W10=", Locale: "en-GB",
		SearchPlaceholder: "Find", DateNote: "Note", Presets: []ReportDatePreset{{Label: "Q3", From: "2026-07-01", To: "2026-09-30"}},
	}.Attrs())
	for _, w := range []string{`data-uri="/r"`, `group-by="affiliateFid"`, `sharable="true"`, `filter="W10="`,
		`locale="en-GB"`, `search-placeholder="Find"`, `date-note="Note"`,
		`presets="[{&#34;label&#34;:&#34;Q3&#34;,&#34;from&#34;:&#34;2026-07-01&#34;,&#34;to&#34;:&#34;2026-09-30&#34;}]"`} {
		if !strings.Contains(a, w) {
			t.Errorf("%q missing %s", a, w)
		}
	}
}
