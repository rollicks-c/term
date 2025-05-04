package term

import (
	"fmt"
	"strings"
)

type CellFormatter func(value string, rowIndex, colIndex int) (string, string)
type RowFormatter func(row map[string]string, rowIndex int) string

type ColFormatter func(value string, rowIndex int) (string, string)

type TableViewBuilder struct {
	cellFormatter CellFormatter
	rowFormatter  RowFormatter
	colFormatters map[string]ColFormatter
	headers       []string
	rows          [][]string
}

func TableView() *TableViewBuilder {
	return &TableViewBuilder{
		headers: []string{},
		rows:    [][]string{},
		cellFormatter: func(value string, rowIndex, colIndex int) (string, string) {
			return "", value
		},
		rowFormatter: func(row map[string]string, rowIndex int) string {
			return "%s"
		},
		colFormatters: map[string]ColFormatter{},
	}
}

func (t *TableViewBuilder) AddHeaders(row ...string) *TableViewBuilder {
	t.headers = append(t.headers, row...)
	return t
}

func (t *TableViewBuilder) AddCellFormatter(cf CellFormatter) *TableViewBuilder {
	t.cellFormatter = cf
	return t
}

func (t *TableViewBuilder) AddColFormatter(col string, cf ColFormatter) *TableViewBuilder {
	t.colFormatters[col] = cf
	return t
}

func (t *TableViewBuilder) AddRowFormatter(rf RowFormatter) *TableViewBuilder {
	t.rowFormatter = rf
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
	for rowIndex := range t.rows {
		for colIndex := range t.rows[rowIndex] {

			// render
			cell := t.renderCell(rowIndex, colIndex, maxWidths)

			// append
			table += cell
			if colIndex < len(t.rows[rowIndex])-1 {
				table += "\t"
			}
		}
		table += "\n"
	}
	return table
}

func (t *TableViewBuilder) renderCell(rowIndex, colIndex int, widths []int) string {

	// gather data
	padLen := widths[colIndex]
	cell := t.rows[rowIndex][colIndex]

	// apply formatters
	rowStyle := t.rowFormatter(t.getRow(rowIndex), rowIndex)
	colStyle, colValue := "", ""
	if colFormatter, ok := t.colFormatters[t.headers[colIndex]]; ok {
		colStyle, colValue = colFormatter(cell, rowIndex)
	}
	cellStyle, cellValue := t.cellFormatter(cell, rowIndex, colIndex)
	style := firstNonEmpty(cellStyle, colStyle, rowStyle, "%s")
	value := firstNonEmpty(cellValue, colValue)

	// escape
	value = strings.ReplaceAll(value, "\n", "\\n")

	// render
	exp := fmt.Sprintf("%-*s", padLen, value)
	exp = fmt.Sprintf(style, exp)

	return exp
}

func (t *TableViewBuilder) getRow(rowIndex int) map[string]string {
	row := make(map[string]string)
	for colIndex, header := range t.headers {
		row[header] = t.rows[rowIndex][colIndex]
	}
	return row
}

func firstNonEmpty(s ...string) string {
	for _, s := range s {
		if s != "" {
			return s
		}
	}
	return ""
}
