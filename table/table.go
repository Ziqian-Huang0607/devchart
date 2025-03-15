package table

import (
	"fmt"
	"io"
	"os" // Add this import
	"strings"
	"unicode/utf8"
)

// Table structure
type Table struct {
	headers []string
	data    [][]string
	options Options
}

// NewTable(default)
func NewTable() *Table {
	return &Table{
		options: DefaultOptions(),
	}
}

// SetHeaders
func (t *Table) SetHeaders(headers []string) *Table {
	t.headers = headers
	return t
}

// SetData
func (t *Table) SetData(data [][]string) *Table {
	t.data = data
	return t
}

// SetOptions
func (t *Table) SetOptions(options Options) *Table {
	t.options = options
	return t
}

// calculateColumnWidths
func (t *Table) calculateColumnWidths() []int {
	numCols := 0
	if t.headers != nil {
		numCols = len(t.headers)
	} else if len(t.data) > 0 && len(t.data[0]) > 0 { // Check for data & row length
		numCols = len(t.data[0])
	}

	if numCols == 0 {
		return []int{} // No columns
	}

	widths := make([]int, numCols)

	if t.headers != nil {
		for i, header := range t.headers {
			widths[i] = utf8.RuneCountInString(header) + (t.options.Padding * 2) // Add padding to width
		}
	}

	for _, row := range t.data {
		for i, cell := range row {
			if i >= numCols { // Handle rows with more columns than headers (or first row)
				continue
			}
			cellWidth := utf8.RuneCountInString(cell) + (t.options.Padding * 2) // Add padding to width
			if cellWidth > widths[i] {
				widths[i] = cellWidth
			}
		}
	}
	return widths
}

// getBorderChars, returns border characters based on BorderStyle
func (t *Table) getBorderChars() (hBorder, vBorder, junction string) {
	switch t.options.BorderStyle {
	case BorderDouble:
		return "═", "║", "╬"
	case BorderNone:
		return " ", " ", " "
	default: // BorderSingle and default
		return "─", "│", "┼"
	}
}

// getHorizontalBorderLine, generates a horizontal border line.
func (t *Table) getHorizontalBorderLine(columnWidths []int, junctionChar string) string {
	hBorder, _, _ := t.getBorderChars()
	line := junctionChar
	for _, width := range columnWidths {
		line += strings.Repeat(hBorder, width) + junctionChar // Corrected: No +2 padding here
	}
	return line + "\n"
}

// renderCell formats a cell with alignment and padding.
func (t *Table) renderCell(cell string, width int) string {
	padding := strings.Repeat(" ", t.options.Padding)
	contentWidth := width - (t.options.Padding * 2) // Calculate content width correctly

	if contentWidth < 0 { // Handle case where padding is larger than width
		contentWidth = 0
	}

	formattedCell := ""
	cellRuneCount := utf8.RuneCountInString(cell)

	switch t.options.Alignment {
	case AlignCenter:
		paddingSize := (contentWidth - cellRuneCount) / 2
		if paddingSize < 0 {
			paddingSize = 0
		}
		cellPadding := strings.Repeat(" ", paddingSize)
		formattedCell = cellPadding + cell + cellPadding
		// Ensure width is exactly 'contentWidth' - important for even centering with odd widths
		formattedCell = fmt.Sprintf("%-*s", contentWidth, formattedCell)
	case AlignRight:
		formattedCell = fmt.Sprintf("%*s", contentWidth, cell)
	default: // AlignLeft and default
		formattedCell = fmt.Sprintf("%-*s", contentWidth, cell)
	}

	return padding + formattedCell + padding
}

// Render renders the table to the given io.Writer with formatting.
func (t *Table) Render(w io.Writer) error {
	columnWidths := t.calculateColumnWidths()
	_, vBorder, junction := t.getBorderChars()

	// Top border (if borders are enabled and there are columns)
	if t.options.BorderStyle != BorderNone && len(columnWidths) > 0 {
		_, err := w.Write([]byte(t.getHorizontalBorderLine(columnWidths, junction)))
		if err != nil {
			return err
		}
	}

	if t.headers != nil {
		if t.options.BorderStyle != BorderNone && len(columnWidths) > 0 {
			_, err := w.Write([]byte(vBorder)) // Initial vertical border
			if err != nil {
				return err
			}
		}
		for i, header := range t.headers {
			cellContent := t.renderCell(header, columnWidths[i])
			_, err := w.Write([]byte(cellContent + vBorder)) // Vertical border after each cell
			if err != nil {
				return err
			}
		}
		_, err := w.Write([]byte("\n")) // Newline after headers
		if err != nil {
			return err
		}
		if t.options.BorderStyle != BorderNone && len(columnWidths) > 0 {
			_, err = w.Write([]byte(t.getHorizontalBorderLine(columnWidths, junction))) // Border after headers
			if err != nil {
				return err
			}
		}
	}

	for _, row := range t.data {
		if t.options.BorderStyle != BorderNone && len(columnWidths) > 0 {
			_, err := w.Write([]byte(vBorder)) // Initial vertical border for each row
			if err != nil {
				return err
			}
		}
		for i, cell := range row {
			if i >= len(columnWidths) { // Handle rows with more columns than calculated widths
				cellContent := t.renderCell(cell, 0) // Render with 0 width if column widths are not available
				_, err := w.Write([]byte(cellContent + vBorder))
				if err != nil {
					return err
				}
			} else {
				cellContent := t.renderCell(cell, columnWidths[i])
				_, err := w.Write([]byte(cellContent + vBorder))
				if err != nil {
					return err
				}
			}
		}
		_, err := w.Write([]byte("\n")) // Newline after each row
		if err != nil {
			return err
		}
		if t.options.BorderStyle != BorderNone && len(columnWidths) > 0 {
			_, err = w.Write([]byte(t.getHorizontalBorderLine(columnWidths, junction))) // Border after each row
			if err != nil {
				return err
			}
		}
	}

	// Bottom border (if borders are enabled and there are columns) - only if there's data or headers
	if t.options.BorderStyle != BorderNone && len(columnWidths) > 0 && (len(t.data) > 0 || t.headers != nil) {
		_, err := w.Write([]byte(t.getHorizontalBorderLine(columnWidths, junction))) // Bottom border
		if err != nil {
			return err
		}
	}

	return nil
}

// Print renders the table to os.Stdout.
func (t *Table) Print() error {
	return t.Render(os.Stdout) // Use os.Stdout instead of io.Stdout
}
