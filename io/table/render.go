package table

import (
	"fmt"
	"strings"
)

func (t *Builder[T]) getMaxWidths() []int {

	// determine max width of each column
	maxWidths := make([]int, len(t.headers))

	// header and footer
	for i, header := range t.headers {
		if t.config.HideHeaders {
			maxWidths[i] = 0
		} else {
			maxWidths[i] = len(header)
		}
		footer, ok := t.footer[header]
		if !ok {
			continue
		}
		if footer.Len() > maxWidths[i] {
			maxWidths[i] = footer.Len()
		}
	}

	// cells
	for _, row := range t.cells {
		for colIndex, cell := range row {
			if cell.Len() > maxWidths[colIndex] {
				maxWidths[colIndex] = cell.Len()
			}
		}
	}

	return maxWidths
}

func (t *Builder[T]) createCells() {
	t.cells = make([][]Cell, len(t.rows))
	for rowIndex := range t.rows {
		t.cells[rowIndex] = make([]Cell, len(t.headers))
		record := t.rows[rowIndex]
		for colIndex, header := range t.headers {
			cell := record.RenderCell(*t.renderContext, header)
			t.cells[rowIndex][colIndex] = cell
		}
	}
}

func (t *Builder[T]) renderHeaders(maxWidths []int) string {

	// no headers
	if t.config.HideHeaders {
		return ""
	}

	// render headers
	out := ""
	for i, header := range t.headers {
		padLen := maxWidths[i]
		exp := fmt.Sprintf("%-*s", padLen, header)
		out += exp
		if i < len(t.headers)-1 {
			out += "\t"
		}
	}

	// render separator
	out += fmt.Sprintf("\n%s", t.config.Indention)
	for i := range t.headers {
		padLen := maxWidths[i]
		exp := strings.Repeat("-", padLen)
		out += exp
		if i < len(t.headers)-1 {
			out += "\t"
		}
	}

	return out
}

func (t *Builder[T]) renderRows(maxWidths []int) string {

	out := ""

	// iterate all cells
	for rowIndex := range t.rows {
		for colIndex := range t.headers {

			// render
			padLen := maxWidths[colIndex]
			cell := t.cells[rowIndex][colIndex].Render(padLen)

			// append
			out += cell
			if colIndex < len(t.headers)-1 {
				out += "\t"
			}
		}
		out += fmt.Sprintf("\n%s", t.config.Indention)
	}

	return out
}

func (t *Builder[T]) renderFooter(maxWidths []int) string {

	// no footer
	if len(t.footer) == 0 {
		return ""
	}

	out := ""

	// render separator
	for i := range t.headers {
		padLen := maxWidths[i]
		_, ok := t.footer[t.headers[i]]
		char := " "
		if ok {
			char = "-"
		}
		out += strings.Repeat(char, padLen)
		if i < len(t.headers)-1 {
			out += "\t"
		}
	}

	// render footer
	out += fmt.Sprintf("\n%s", t.config.Indention)
	for i := range t.headers {
		padLen := maxWidths[i]
		cell, ok := t.footer[t.headers[i]]
		if !ok {
			exp := strings.Repeat(" ", padLen)
			out += exp
		} else {
			out += cell.Render(padLen)
		}
		if i < len(t.headers)-1 {
			out += "\t"
		}
	}

	return out
}
