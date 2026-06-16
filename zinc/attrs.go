package zinc

import (
	"encoding/json"
	"html/template"
	"reflect"
	"strconv"
	"strings"
)

// AttrValuer is implemented by types that customize how they render as an
// HTML attribute value when passed through HTMLAttrs.
//
// The returned string is the raw attribute value - HTMLAttrs HTML-escapes it
// before inserting into name="value" position, so implementations should
// return the value pre-escaping (e.g. raw JSON, not HTML-entity-encoded JSON).
type AttrValuer interface {
	AttrValue() string
}

// HTMLAttrs renders a struct's exported fields as HTML attributes for use
// inside a tag in a template: <my-tag {{attrs .Config}}></my-tag>.
//
// The attribute name comes from the `json` tag (with ,omitempty stripped);
// falls back to the lowercased field name when no tag is set. Fields tagged
// `json:"-"` are skipped.
//
// Encoding precedence per field value:
//   - if the type implements AttrValuer, AttrValue() is called
//   - strings, numbers, and bools render as-is
//   - any other type is JSON-encoded
//
// All values are HTML-escaped before insertion. The return type is
// template.HTMLAttr so html/template inserts the result in attribute position
// without double-escaping.
//
// Example:
//
//	<data-table {{$.DataTableConfig.Attrs}}></data-table>
//	renders: <data-table method="POST" data-uri="/x" unsortable="false">
func HTMLAttrs(v any) template.HTMLAttr {
	rv := reflect.ValueOf(v)
	for rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return ""
		}
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Struct {
		return ""
	}

	var b strings.Builder
	rt := rv.Type()
	for i := range rv.NumField() {
		fld := rt.Field(i)
		if !fld.IsExported() {
			continue
		}

		tag := fld.Tag.Get("json")
		if tag == "-" {
			continue
		}
		name, _, _ := strings.Cut(tag, ",")
		if name == "" {
			name = strings.ToLower(fld.Name)
		}

		fv := rv.Field(i)
		if fv.Kind() == reflect.Ptr {
			if fv.IsNil() {
				continue
			}
			fv = fv.Elem()
		}
		if fv.IsZero() {
			continue
		}
		if (fv.Kind() == reflect.Slice || fv.Kind() == reflect.Map) && fv.Len() == 0 {
			continue
		}

		val, ok := formatAttrValue(fv)
		if !ok {
			continue
		}

		b.WriteByte(' ')
		b.WriteString(name)
		b.WriteString(`="`)
		b.WriteString(template.HTMLEscapeString(val))
		b.WriteString(`"`)
	}
	return template.HTMLAttr(b.String())
}

func formatAttrValue(v reflect.Value) (string, bool) {
	if av, ok := v.Interface().(AttrValuer); ok {
		return av.AttrValue(), true
	}
	if v.CanAddr() {
		if av, ok := v.Addr().Interface().(AttrValuer); ok {
			return av.AttrValue(), true
		}
	}

	switch v.Kind() {
	case reflect.String:
		return v.String(), true
	case reflect.Bool:
		return strconv.FormatBool(v.Bool()), true
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10), true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(v.Uint(), 10), true
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'f', -1, 64), true
	default:
		j, err := json.Marshal(v.Interface())
		if err != nil {
			return "", false
		}
		return string(j), true
	}
}
