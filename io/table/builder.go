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

func (b *Builder[T]) AddHeaders(row ...string) *Builder[T] {
	b.headers = append(b.headers, row...)
	return b
}

func (b *Builder[T]) AddCellFormatter(cf CellRenderer[T]) *Builder[T] {
	b.renderContext.cellRenderer = cf
	return b
}

func (b *Builder[T]) AddRow(rows ...T) *Builder[T] {
	for _, r := range rows {
		b.rows = append(b.rows, dataRow[T]{
			record: r,
		})
	}
	return b
}

func (b *Builder[T]) AddSeparator(char string) *Builder[T] {
	b.rows = append(b.rows, separatorRow[T]{
		char: char,
	})
	return b
}

func (b *Builder[T]) AddCustomCell(header, value, style string) *Builder[T] {
	cr := customRow[T]{
		data: make(map[string]dataCell),
	}
	b.rows = append(b.rows, cr)
	return b.AppendCustomCell(header, value, style)
}

func (b *Builder[T]) AppendCustomCell(header, value, style string) *Builder[T] {
	cr := b.ensureCustomRow()
	cell := dataCell{
		value: value,
		style: style,
	}
	cr.data[header] = cell
	b.rows[len(b.rows)-1] = cr
	return b
}

func (b *Builder[T]) DefaultFormatter() *Builder[T] {

	cf := func(record T, header string) (string, string) {
		return "%s", fmt.Sprintf("%v", record)
	}
	b.AddCellFormatter(cf)
	return b

}

func (b *Builder[T]) ensureCustomRow() customRow[T] {

	// table is empty
	if len(b.rows) == 0 {
		cr := customRow[T]{
			data: make(map[string]dataCell),
		}
		b.rows = append(b.rows, cr)
		return cr
	}

	// last row is not custom row
	cr, ok := b.rows[len(b.rows)-1].(customRow[T])
	if !ok {
		cr = customRow[T]{
			data: make(map[string]dataCell),
		}
		b.rows = append(b.rows, cr)
	}

	return cr
}

func (b *Builder[T]) AddFooterCell(header, value, style string) *Builder[T] {
	cell := dataCell{
		value: value,
		style: style,
	}
	b.footer[header] = cell
	return b
}

func (b *Builder[T]) Build() string {

	// create cells
	b.createCells()

	// determine max width of each column
	maxWidths := b.getMaxWidths()

	// render
	var table string
	table += b.config.Indention

	// print headers
	table += b.renderHeaders(maxWidths)

	// print rows
	table += fmt.Sprintf("\n%s", b.config.Indention)
	table += b.renderRows(maxWidths)

	// print footer
	table += b.renderFooter(maxWidths)
	table += "\n"

	return table
}
