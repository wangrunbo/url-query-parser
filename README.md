# url-query-parser

Parse url query into a map.

## Installation

1. Add `url-query-parser` as a dependency in your program.
```sh
go get -u github.com/wangrunbo/url-query-parser
```

2. Import it in your code:
```go
import querypsr "github.com/wangrunbo/url-query-parser"
```

## Usage

### Parse a query string into a map
```go
queryString := "a=1&b[]=2&b[]=3&c[d]=4&c[e]=5"
dataMap := querypsr.ParseString(queryString)

// dataMap = map[string]any{
//     "a": "1",
//     "b": []any{"2", "3"},
//     "c": map[string]any{"d": "4", "e": "5"},
// }
```

### Parse a key/value pair into a map
```go
dataMap := querypsr.Parse("a", "1")

// dataMap = map[string]any{
//     "a": "1",
// }
```
```go
dataMap := querypsr.Parse("a[]", "1")

// dataMap = map[string]any{
//     "a": []any{"1"},
// }
```
```go
dataMap := querypsr.Parse("a[b]", "1")

// dataMap = map[string]any{
//     "a": map[string]any{"b": "1"},
// }
```

### Parse a query string and append to an existing map
```go
dataMap := map[string]any{
    "a": "1",
    "b": []any{"2"},
    "c": map[string]any{"d": "3"},
}
queryString := "a=2&b[]=3&c[e]=4&f=1"
dataMap = querypsr.ParseStringToExistingMap(queryString, dataMap)

// dataMap = map[string]any{
//     "a": "2",
//     "b": []any{"2", "3"},
//     "c": map[string]any{"d": "3", "e": "4"},
//     "f": "1",
// }
```

### Parse a key/value pair and append to an existing map
```go
dataMap := map[string]any{
    "a": "1",
}
dataMap = querypsr.ParseToExistingMap("b", "2", dataMap)

// dataMap = map[string]any{
//     "a": "1",
//     "b": "2",
// }
```
```go
dataMap := map[string]any{
    "a": []any{"1"},
}
dataMap = querypsr.ParseToExistingMap("a[]", "2", dataMap)

// dataMap = map[string]any{
//     "a": []any{"1", "2"},
// }
```
```go
dataMap := map[string]any{
    "a": map[string]any{"b": "1"},
}
dataMap = querypsr.ParseToExistingMap("a[c]", "2", dataMap)

// dataMap = map[string]any{
//     "a": map[string]any{"b": "1", "c": "2"},
// }
```