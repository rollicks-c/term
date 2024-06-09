package term

import (
	"fmt"
	"strings"
)

type CellColorizer func(rowIndex, colIndex int) string

type TableViewBuilder struct {
	cellColorizer CellColorizer
	headers       []string
	rows          [][]string
}

func TableView() *TableViewBuilder {
	return &TableViewBuilder{
		cellColorizer: func(rowIndex, colIndex int) string {
			return "%s"
		},
		headers: []string{},
		rows:    [][]string{},
	}
}

func (t *TableViewBuilder) AddHeaders(row ...string) *TableViewBuilder {
	t.headers = append(t.headers, row...)
	return t
}

func (t *TableViewBuilder) AddCellFormatter(cf CellColorizer) *TableViewBuilder {
	t.cellColorizer = cf
	return t
}

func (t *TableViewBuilder) AddRow(row ...string) *TableViewBuilder {
	t.rows = append(t.rows, row)
	return t
}

func (t *TableViewBuilder) Build() string {

	// determine max width of each column
	maxWidths := make([]int, len(t.headers))
	for i, header := range t.headers {
		maxWidths[i] = len(header)
	}
	for _, row := range t.rows {
		for colIndex, cell := range row {
			if len(cell) > maxWidths[colIndex] {
				maxWidths[colIndex] = len(cell)
			}
		}
	}

	// print headers
	var table string
	for i, header := range t.headers {
		padLen := maxWidths[i]
		exp := fmt.Sprintf("%-*s", padLen, header)
		table += exp
		if i < len(t.headers)-1 {
			table += "\t"
		}
	}
	table += "\n"
	for i := range t.headers {
		padLen := maxWidths[i]
		exp := strings.Repeat("-", padLen)
		table += exp
		if i < len(t.headers)-1 {
			table += "\t"
		}
	}

	// print rows
	table += "\n"
	for rowIndex, row := range t.rows {
		for colIndex, cell := range row {
			padLen := maxWidths[colIndex]
			exp := fmt.Sprintf("%-*s", padLen, cell)
			color := t.cellColorizer(rowIndex, colIndex)
			exp = fmt.Sprintf(color, exp)
			table += exp
			if colIndex < len(row)-1 {
				table += "\t"
			}
		}
		table += "\n"
	}
	return table
}
