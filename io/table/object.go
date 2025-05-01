package table

import (
	"encoding/json"
	"fmt"
	"sort"
)

type FlatObject = map[string]any

func FromObject(obj any, options ...Option) *Builder[FlatObject] {

	// convert to map
	raw, err := json.Marshal(obj)
	if err != nil {
		return createErrorTable(err)
	}
	data := map[string]interface{}{}
	if err := json.Unmarshal(raw, &data); err != nil {
		return createErrorTable(err)
	}

	// extract data
	fields := make([]string, 0, len(data))
	for k := range data {
		fields = append(fields, k)
	}
	cf := func(record FlatObject, header string) (string, string) {
		return "%s", fmt.Sprintf("%v", record[header])
	}
	sort.Slice(fields, func(i, j int) bool {
		return fields[i] < fields[j]
	})

	// build table
	builder := NewBuilder[FlatObject](options...).
		AddHeaders(fields...).
		AddRow(data).
		AddCellFormatter(cf)
	return builder
}

func createErrorTable(err error) *Builder[FlatObject] {
	cf := func(record FlatObject, header string) (string, string) {
		return "%s", fmt.Sprintf("%v", record[header])
	}
	return NewBuilder[FlatObject]().
		AddHeaders("error").
		AddRow(map[string]any{"error": err}).
		AddCellFormatter(cf)
}
