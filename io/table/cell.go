package table

import (
	"fmt"
	"strings"
)

type Cell interface {
	Render(width int) string
	Len() int
}

type dataCell struct {
	value string
	style string
}
type separatorCell struct {
	char string
}

func (dc dataCell) Render(width int) string {
	valuePadded := fmt.Sprintf("%-*s", width, dc.value)
	content := fmt.Sprintf(dc.style, valuePadded)
	return content
}
func (dc dataCell) Len() int {
	return len(dc.value)
}

func (sc separatorCell) Render(width int) string {
	return strings.Repeat(sc.char, width)
}

func (sc separatorCell) Len() int {
	return 1
}
