package table

import (
	"fmt"
	"strings"
)

func (b *Builder[T]) getMaxWidths() []int {

	// determine max width of each column
	maxWidths := make([]int, len(b.headers))

	// header and footer
	for i, header := range b.headers {
		if b.config.HideHeaders {
			maxWidths[i] = 0
		} else {
			maxWidths[i] = len(header)
		}
		footer, ok := b.footer[header]
		if !ok {
			continue
		}
		if footer.Len() > maxWidths[i] {
			maxWidths[i] = footer.Len()
		}
	}

	// cells
	for _, row := range b.cells {
		for colIndex, cell := range row {
			if cell.Len() > maxWidths[colIndex] {
				maxWidths[colIndex] = cell.Len()
			}
		}
	}

	return maxWidths
}

func (b *Builder[T]) createCells() {
	b.cells = make([][]Cell, len(b.rows))
	for rowIndex := range b.rows {
		b.cells[rowIndex] = make([]Cell, len(b.headers))
		record := b.rows[rowIndex]
		for colIndex, header := range b.headers {
			cell := record.RenderCell(*b.renderContext, header)
			b.cells[rowIndex][colIndex] = cell
		}
	}
}

func (b *Builder[T]) renderHeaders(maxWidths []int) string {

	// no headers
	if b.config.HideHeaders {
		return ""
	}

	// render headers
	out := ""
	for i, header := range b.headers {
		padLen := maxWidths[i]
		exp := fmt.Sprintf("%-*s", padLen, header)
		out += exp
		if i < len(b.headers)-1 {
			out += "\t"
		}
	}

	// render separator
	out += fmt.Sprintf("\n%s", b.config.Indention)
	for i := range b.headers {
		padLen := maxWidths[i]
		exp := strings.Repeat("-", padLen)
		out += exp
		if i < len(b.headers)-1 {
			out += "\t"
		}
	}

	return out
}

func (b *Builder[T]) renderRows(maxWidths []int) string {

	out := ""

	// iterate all cells
	for rowIndex := range b.rows {
		for colIndex := range b.headers {

			// render
			padLen := maxWidths[colIndex]
			cell := b.cells[rowIndex][colIndex].Render(padLen)

			// append
			out += cell
			if colIndex < len(b.headers)-1 {
				out += "\t"
			}
		}
		out += fmt.Sprintf("\n%s", b.config.Indention)
	}

	return out
}

func (b *Builder[T]) renderFooter(maxWidths []int) string {

	// no footer
	if len(b.footer) == 0 {
		return ""
	}

	out := ""

	// render separator
	for i := range b.headers {
		padLen := maxWidths[i]
		_, ok := b.footer[b.headers[i]]
		char := " "
		if ok {
			char = "-"
		}
		out += strings.Repeat(char, padLen)
		if i < len(b.headers)-1 {
			out += "\t"
		}
	}

	// render footer
	out += fmt.Sprintf("\n%s", b.config.Indention)
	for i := range b.headers {
		padLen := maxWidths[i]
		cell, ok := b.footer[b.headers[i]]
		if !ok {
			exp := strings.Repeat(" ", padLen)
			out += exp
		} else {
			out += cell.Render(padLen)
		}
		if i < len(b.headers)-1 {
			out += "\t"
		}
	}

	return out
}
