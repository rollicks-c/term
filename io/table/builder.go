package table

import (
	"fmt"
)

type Option func(config *Config)

type CellRenderer[T any] func(record T, header string) (string, string)

type Config struct {
	HideHeaders bool
	Indention   string
}

type Builder[T any] struct {
	renderContext *renderContext[T]
	headers       []string
	rows          []row[T]
	cells         [][]Cell
	footer        map[string]Cell
	config        *Config
}

func NewBuilder[T any](options ...Option) *Builder[T] {
	t := &Builder[T]{
		headers: []string{},
		rows:    []row[T]{},
		renderContext: &renderContext[T]{
			cellRenderer: func(value T, header string) (string, string) {
				return "%s", fmt.Sprintf("%v", value)
			},
		},

		footer: make(map[string]Cell),
		config: &Config{
			HideHeaders: false,
		},
	}
	for _, opt := range options {
		opt(t.config)
	}
	return t
}

func (t *Builder[T]) AddHeaders(row ...string) *Builder[T] {
	t.headers = append(t.headers, row...)
	return t
}

func (t *Builder[T]) AddCellFormatter(cf CellRenderer[T]) *Builder[T] {
	t.renderContext.cellRenderer = cf
	return t
}

func (t *Builder[T]) AddRow(rows ...T) *Builder[T] {
	for _, r := range rows {
		t.rows = append(t.rows, dataRow[T]{
			record: r,
		})
	}
	return t
}

func (t *Builder[T]) AddSeparator(char string) *Builder[T] {
	t.rows = append(t.rows, separatorRow[T]{
		char: char,
	})
	return t
}

func (t *Builder[T]) AddCustomCell(header, value, style string) *Builder[T] {
	cr := customRow[T]{
		data: make(map[string]dataCell),
	}
	t.rows = append(t.rows, cr)
	return t.AppendCustomCell(header, value, style)
}

func (t *Builder[T]) AppendCustomCell(header, value, style string) *Builder[T] {
	cr := t.ensureCustomRow()
	cell := dataCell{
		value: value,
		style: style,
	}
	cr.data[header] = cell
	t.rows[len(t.rows)-1] = cr
	return t
}

func (t *Builder[T]) ensureCustomRow() customRow[T] {

	// table is empty
	if len(t.rows) == 0 {
		cr := customRow[T]{
			data: make(map[string]dataCell),
		}
		t.rows = append(t.rows, cr)
		return cr
	}

	// last row is not custom row
	cr, ok := t.rows[len(t.rows)-1].(customRow[T])
	if !ok {
		cr = customRow[T]{
			data: make(map[string]dataCell),
		}
		t.rows = append(t.rows, cr)
	}

	return cr
}

func (t *Builder[T]) AddFooterCell(header, value, style string) *Builder[T] {
	cell := dataCell{
		value: value,
		style: style,
	}
	t.footer[header] = cell
	return t
}

func (t *Builder[T]) Build() string {

	// create cells
	t.createCells()

	// determine max width of each column
	maxWidths := t.getMaxWidths()

	// render
	var table string
	table += t.config.Indention

	// print headers
	table += t.renderHeaders(maxWidths)

	// print rows
	table += fmt.Sprintf("\n%s", t.config.Indention)
	table += t.renderRows(maxWidths)

	// print footer
	table += t.renderFooter(maxWidths)
	table += "\n"

	return table
}
