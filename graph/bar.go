package graph

import (
	"io"
	"math"
	"os"
	"strings"
)

// BarGraph structure
type BarGraph struct {
	data    []float64
	labels  []string
	options BarGraphOptions
}

// NewBarGraph (default)
func NewBarGraph() *BarGraph {
	return &BarGraph{
		options: DefaultBarGraphOptions(),
	}
}

// SetData
func (bg *BarGraph) SetData(data []float64) *BarGraph {
	bg.data = data
	return bg
}

// SetLabels
func (bg *BarGraph) SetLabels(labels []string) *BarGraph {
	bg.labels = labels
	return bg
}

// SetOptions
func (bg *BarGraph) SetOptions(options BarGraphOptions) *BarGraph {
	bg.options = options
	return bg
}

// Render renders the bar graph to the given io.Writer.
func (bg *BarGraph) Render(w io.Writer) error {
	if len(bg.data) == 0 {
		return nil // Nothing to render for empty data
	}

	maxValue := 0.0
	for _, val := range bg.data {
		if val > maxValue {
			maxValue = val
		}
	}

	if maxValue == 0 {
		maxValue = 1 // Avoid division by zero if all data is 0, set a min max for scaling
	}

	// numBars := len(bg.data)
	graphWidth := bg.options.Width

	barHeightScale := float64(10) / maxValue // Scale bars to max height of ~10 lines (adjust as needed)

	// Render bars from top to bottom
	for i := 10; i >= 0; i-- { // Assuming max height of 10 lines
		line := ""
		for barIndex, value := range bg.data {
			barHeight := int(math.Round(value * barHeightScale))
			if barHeight >= i {
				line += bg.options.BarCharacter
			} else {
				line += " "
			}
			if barIndex < len(bg.data)-1 {
				line += " " // Space between bars (fixed single space for now)
			}
		}
		_, err := w.Write([]byte(line + "\n"))
		if err != nil {
			return err
		}
	}

	// Render X-axis
	axisLine := strings.Repeat(bg.options.AxisCharacter, graphWidth)
	_, err := w.Write([]byte(axisLine + "\n"))
	if err != nil {
		return err
	}

	// Render labels (optional)
	if len(bg.labels) == len(bg.data) {
		labelLine := ""
		for i, label := range bg.labels {
			labelLine += label
			if i < len(bg.labels)-1 {
				labelLine += " " // Space between labels
			}
		}
		_, err = w.Write([]byte(labelLine + "\n"))
		if err != nil {
			return err
		}
	}

	return nil
}

// print
func (bg *BarGraph) Print() error {
	return bg.Render(os.Stdout)
}
