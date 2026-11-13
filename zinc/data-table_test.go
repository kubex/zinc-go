package zinc

import "testing"

func TestHeaderConfigsKeyedByKey(t *testing.T) {
	h := HeaderConfigs{{Key: "name", Label: "Name"}, {Key: "email", Label: "Email"}}
	want := `{"name":{"key":"name","label":"Name"},"email":{"key":"email","label":"Email"}}`
	if got := h.String(); got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
}

func TestFilterOptionsOrderedByValue(t *testing.T) {
	b, err := DataTableFilterOptions{"b": "Apple", "a": "Zebra", "c": "Apple"}.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	if want := `{"b":"Apple","c":"Apple","a":"Zebra"}`; string(b) != want {
		t.Fatalf("got %s, want %s", b, want)
	}
}
