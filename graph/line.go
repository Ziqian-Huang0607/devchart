package graph

import (
	"io"
	"math"
	"os"
	"strings"
	"unicode/utf8"
)

// LineGraph represents a line graph structure.
type LineGraph struct {
	data    []float64
	labels  []string // Optional labels for X-axis
	options LineGraphOptions
}

// NewLineGraph creates a new LineGraph instance with default options.
func NewLineGraph() *LineGraph {
	return &LineGraph{
		options: DefaultLineGraphOptions(),
	}
}

// SetData sets the line graph data.
func (lg *LineGraph) SetData(data []float64) *LineGraph {
	lg.data = data
	return lg
}

// SetLabels sets the line graph labels (optional X-axis labels).
func (lg *LineGraph) SetLabels(labels []string) *LineGraph {
	lg.labels = labels
	return lg
}

// SetOptions sets the line graph options.
func (lg *LineGraph) SetOptions(options LineGraphOptions) *LineGraph {
	lg.options = options
	return lg
}

// Render renders the line graph to the given io.Writer.
func (lg *LineGraph) Render(w io.Writer) error {
	if len(lg.data) == 0 {
		return nil // Nothing to render for empty data
	}

	maxValue := lg.data[0]
	minValue := lg.data[0]
	for _, val := range lg.data {
		if val > maxValue {
			maxValue = val
		}
		if val < minValue {
			minValue = val
		}
	}

	if maxValue == minValue {
		maxValue = minValue + 1 // Avoid division by zero if all data is the same, set a min range
	}

	graphWidth := lg.options.Width
	graphHeight := lg.options.Height

	yScale := float64(graphHeight-1) / (maxValue - minValue) // -1 for axis line

	// Initialize graph grid (empty spaces)
	graphGrid := make([][]rune, graphHeight)
	for i := range graphGrid {
		graphGrid[i] = make([]rune, graphWidth)
		for j := range graphGrid[i] {
			graphGrid[i][j] = ' '
		}
	}

	// Plot data points and connect them (very basic line)
	prevX, prevY := -1, -1 // Initialize previous point to invalid values

	for xIndex, value := range lg.data {
		yValue := int(math.Round((maxValue - value) * yScale)) // Invert Y-axis for terminal (top is 0)
		if yValue < 0 {
			yValue = 0
		}
		if yValue >= graphHeight {
			yValue = graphHeight - 1
		}

		xPos := int(float64(xIndex) / float64(len(lg.data)-1) * float64(graphWidth-1)) // Scale X to width

		graphGrid[yValue][xPos] = []rune(lg.options.PointCharacter)[0] // Plot point

		if prevX != -1 { // Connect to previous point (very basic line)
			// Basic horizontal line for simplicity. More complex line drawing is much harder in text.
			for x := min(prevX, xPos) + 1; x < max(prevX, xPos); x++ {
				graphGrid[prevY][x] = []rune(lg.options.AxisCharacter)[0] // Use axis char for line
			}
			// Basic vertical line (if needed for very steep slopes) - optional to add, but kept horizontal for now
		}
		prevX, prevY = xPos, yValue
	}

	// Draw X and Y axis (basic) - Y axis at the start, X axis at the bottom
	for i := 0; i < graphHeight; i++ {
		graphGrid[i][0] = []rune("│")[0] // Y-axis
	}
	for j := 0; j < graphWidth; j++ {
		graphGrid[graphHeight-1][j] = []rune(lg.options.AxisCharacter)[0] // X-axis
	}
	graphGrid[graphHeight-1][0] = []rune("┼")[0] // Origin

	// Render graph grid
	for _, row := range graphGrid {
		_, err := w.Write([]byte(string(row) + "\n"))
		if err != nil {
			return err
		}
	}

	// Render X-axis labels (optional) - basic label rendering below the graph
	if len(lg.labels) == len(lg.data) {
		labelLine := ""
		labelSpacing := 0
		var err error
		if len(lg.labels) > 1 { // Avoid division by zero
			labelSpacing = graphWidth / (len(lg.labels) - 1)
		}
		for i, label := range lg.labels {
			labelLine += label
			if i < len(lg.labels)-1 {
				labelLine += strings.Repeat(" ", labelSpacing-utf8.RuneCountInString(label)) // Basic spacing attempt
			}
		}
		_, err = w.Write([]byte(labelLine + "\n"))
		if err != nil {
			return err
		}
	}

	return nil
}

// Print renders the line graph to os.Stdout.
func (lg *LineGraph) Print() error {
	return lg.Render(os.Stdout)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
