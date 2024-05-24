package querypsr

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	var dm map[string]any
	var excepted any

	dm = Parse("a", "value")
	excepted = map[string]any{"a": "value"}
	if !reflect.DeepEqual(dm, excepted) {
		t.Errorf("Expected '%v', got '%v'", excepted, dm)
	}

	dm = Parse("a[]", "value")
	excepted = map[string]any{"a": []any{"value"}}
	if !reflect.DeepEqual(dm, excepted) {
		t.Errorf("Expected '%v', got '%v'", excepted, dm)
	}

	dm = Parse("a[b]", "value")
	excepted = map[string]any{"a": map[string]any{"b": "value"}}
	if !reflect.DeepEqual(dm, excepted) {
		t.Errorf("Expected '%v', got '%v'", excepted, dm)
	}

	dm = Parse("a[][b]", "value")
	excepted = map[string]any{"a": []any{map[string]any{"b": "value"}}}
	if !reflect.DeepEqual(dm, excepted) {
		t.Errorf("Expected '%v', got '%v'", excepted, dm)
	}

	dm = Parse("a[b][]", "value")
	excepted = map[string]any{"a": map[string]any{"b": []any{"value"}}}
	if !reflect.DeepEqual(dm, excepted) {
		t.Errorf("Expected '%v', got '%v'", excepted, dm)
	}

	dm = Parse("a[b][c]", "value")
	excepted = map[string]any{"a": map[string]any{"b": map[string]any{"c": "value"}}}
	if !reflect.DeepEqual(dm, excepted) {
		t.Errorf("Expected '%v', got '%v'", excepted, dm)
	}

	dm = Parse("[a][b][c]", "value")
	excepted = map[string]any{"[a][b][c]": "value"}
	if !reflect.DeepEqual(dm, excepted) {
		t.Errorf("Expected '%v', got '%v'", excepted, dm)
	}
}

func TestParseString(t *testing.T) {
	var dm map[string]any
	var excepted any

	queryString := "a=1&b[]=1&b[]=2&b[][c]=3&b[][]=4&d[e]=1&d[f]=2&d[g][]=3&d[g][]=4&d[h][i]=5&d[h][j]=6&[k]=1"
	dm = ParseString(queryString)
	excepted = map[string]any{
		"a": "1",
		"b": []any{"1", "2", map[string]any{"c": "3"}, []any{"4"}},
		"d": map[string]any{
			"e": "1",
			"f": "2",
			"g": []any{"3", "4"},
			"h": map[string]any{"i": "5", "j": "6"},
		},
		"[k]": "1",
	}
	if !reflect.DeepEqual(dm, excepted) {
		t.Errorf("Expected '%v', got '%v'", excepted, dm)
	}
}

func TestParseToExistingMap(t *testing.T) {
	dm := map[string]any{}
	var excepted any

	dm = ParseToExistingMap("a", "value", dm)
	excepted = map[string]any{"a": "value"}
	if !reflect.DeepEqual(dm, excepted) {
		t.Errorf("Expected '%v', got '%v'", excepted, dm)
	}

	dm = ParseToExistingMap("b", "value", dm)
	excepted = map[string]any{"a": "value", "b": "value"}
	if !reflect.DeepEqual(dm, excepted) {
		t.Errorf("Expected '%v', got '%v'", excepted, dm)
	}

	dm = ParseToExistingMap("c[]", "value", dm)
	excepted = map[string]any{"a": "value", "b": "value", "c": []any{"value"}}
	if !reflect.DeepEqual(dm, excepted) {
		t.Errorf("Expected '%v', got '%v'", excepted, dm)
	}

	dm = ParseToExistingMap("d[e]", "value", dm)
	excepted = map[string]any{"a": "value", "b": "value", "c": []any{"value"}, "d": map[string]any{"e": "value"}}
	if !reflect.DeepEqual(dm, excepted) {
		t.Errorf("Expected '%v', got '%v'", excepted, dm)
	}

	dm = ParseToExistingMap("c[]", "value", dm)
	excepted = map[string]any{"a": "value", "b": "value", "c": []any{"value", "value"}, "d": map[string]any{"e": "value"}}
	if !reflect.DeepEqual(dm, excepted) {
		t.Errorf("Expected '%v', got '%v'", excepted, dm)
	}

	dm = ParseToExistingMap("d[f]", "value", dm)
	excepted = map[string]any{"a": "value", "b": "value", "c": []any{"value", "value"}, "d": map[string]any{"e": "value", "f": "value"}}
	if !reflect.DeepEqual(dm, excepted) {
		t.Errorf("Expected '%v', got '%v'", excepted, dm)
	}

	dm = ParseToExistingMap("c[][b]", "value", dm)
	excepted = map[string]any{"a": "value", "b": "value", "c": []any{"value", "value", map[string]any{"b": "value"}}, "d": map[string]any{"e": "value", "f": "value"}}
	if !reflect.DeepEqual(dm, excepted) {
		t.Errorf("Expected '%v', got '%v'", excepted, dm)
	}

	dm = ParseToExistingMap("h[a][b]", "value", dm)
	excepted = map[string]any{"a": "value", "b": "value", "c": []any{"value", "value", map[string]any{"b": "value"}}, "d": map[string]any{"e": "value", "f": "value"}, "h": map[string]any{"a": map[string]any{"b": "value"}}}
	if !reflect.DeepEqual(dm, excepted) {
		t.Errorf("Expected '%v', got '%v'", excepted, dm)
	}

	dm = ParseToExistingMap("h[a][c]", "value", dm)
	excepted = map[string]any{"a": "value", "b": "value", "c": []any{"value", "value", map[string]any{"b": "value"}}, "d": map[string]any{"e": "value", "f": "value"}, "h": map[string]any{"a": map[string]any{"b": "value", "c": "value"}}}
	if !reflect.DeepEqual(dm, excepted) {
		t.Errorf("Expected '%v', got '%v'", excepted, dm)
	}

	dm = ParseToExistingMap("[a][b][c]", "value", dm)
	excepted = map[string]any{"a": "value", "b": "value", "c": []any{"value", "value", map[string]any{"b": "value"}}, "d": map[string]any{"e": "value", "f": "value"}, "h": map[string]any{"a": map[string]any{"b": "value", "c": "value"}}, "[a][b][c]": "value"}
	if !reflect.DeepEqual(dm, excepted) {
		t.Errorf("Expected '%v', got '%v'", excepted, dm)
	}
}

func TestParseStringToExistingMap(t *testing.T) {
	dm := map[string]any{
		"a": "1",
		"b": []any{"1", "2", map[string]any{"c": "3"}, []any{"4"}},
		"d": map[string]any{
			"e": "1",
			"f": "2",
			"g": []any{"3", "4"},
			"h": map[string]any{"i": "5", "j": "6"},
		},
	}

	queryString := "a=2&b[]=5&d[e]=2&d[g][]=5&d[h][i]=6&d[h][k]=7&[k]=1"
	dm = ParseStringToExistingMap(queryString, dm)
	excepted := map[string]any{
		"a": "2",
		"b": []any{"1", "2", map[string]any{"c": "3"}, []any{"4"}, "5"},
		"d": map[string]any{
			"e": "2",
			"f": "2",
			"g": []any{"3", "4", "5"},
			"h": map[string]any{"i": "6", "j": "6", "k": "7"},
		},
		"[k]": "1",
	}
	if !reflect.DeepEqual(dm, excepted) {
		t.Errorf("Expected '%v', got '%v'", excepted, dm)
	}
}
