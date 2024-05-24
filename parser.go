package querypsr

import (
	"net/url"
	"reflect"
	"strings"
)

func Parse(k, v string) map[string]any {
	return ParseToExistingMap(k, v, dataMap{})
}

func ParseString(qs string) map[string]any {
	return ParseStringToExistingMap(qs, dataMap{})
}

func ParseToExistingMap(k, v string, m map[string]any) map[string]any {
	dm := dataMap(m)
	dm.add(k, v)
	return dm
}

func ParseStringToExistingMap(qs string, m map[string]any) map[string]any {
	dm := dataMap(m)
	if q, err := url.ParseQuery(qs); err == nil {
		for key, values := range q {
			for _, value := range values {
				dm.add(key, value)
			}
		}
	}

	return dm
}

type dataMap map[string]any

func (dm *dataMap) add(k, v string) {
	keys := strings.Split(k, "[")
	if len(keys) < 2 || len(keys[0]) == 0 {
		(*dm)[k] = v
		return
	}

	data := any(v)
	for i := len(keys[1:]) - 1; i >= 0; i-- {
		key := keys[1:][i]
		if len(key) == 0 || key[len(key)-1:] != "]" {
			(*dm)[k] = v
			return
		}

		key = key[:len(key)-1]
		if key == "" {
			data = []any{data}
		} else {
			data = map[string]any{key: data}
		}
	}

	if (*dm)[keys[0]] == nil {
		(*dm)[keys[0]] = data
	} else {
		(*dm)[keys[0]] = merge((*dm)[keys[0]], data)
	}
}

func merge(a, b any) (data any) {
	aType := reflect.TypeOf(a)
	bType := reflect.TypeOf(b)
	if aType.Kind() == reflect.Slice && bType.Kind() == reflect.Slice {
		data = append(a.([]any), b.([]any)...)
	} else if aType.Kind() == reflect.Map && bType.Kind() == reflect.Map {
		data = a.(map[string]any)
		for k, v := range b.(map[string]any) {
			if value, ok := data.(map[string]any)[k]; ok {
				data.(map[string]any)[k] = merge(value, v)
			} else {
				data.(map[string]any)[k] = v
			}
		}
	} else {
		data = b
	}
	return
}
