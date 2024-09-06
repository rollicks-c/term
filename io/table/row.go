package table

import "strings"

type renderContext[T any] struct {
	cellRenderer CellRenderer[T]
}

type row[T any] interface {
	RenderCell(ctx renderContext[T], header string) Cell
}

type dataRow[T any] struct {
	record T
}

func (d dataRow[T]) RenderCell(ctx renderContext[T], header string) Cell {

	// apply formatters
	style, valueRaw := ctx.cellRenderer(d.record, header)

	// escape
	valueRaw = strings.ReplaceAll(valueRaw, "\n", "\\n")

	// wrap
	return dataCell{
		value: valueRaw,
		style: style,
	}
}

type separatorRow[T any] struct {
	char string
}

func (s separatorRow[T]) RenderCell(ctx renderContext[T], header string) Cell {
	return separatorCell{
		char: s.char,
	}
}

type customRow[T any] struct {
	data map[string]dataCell
}

func (c customRow[T]) RenderCell(ctx renderContext[T], header string) Cell {
	data, ok := c.data[header]
	if !ok {
		return dataCell{
			value: "",
			style: "%s",
		}
	}
	return data
}
