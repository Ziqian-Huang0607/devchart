package graph

// BarGraphOptions
type BarGraphOptions struct {
	Width         int    // Maximum width of the graph in characters
	BarCharacter  string // Character used to draw bars
	AxisCharacter string // Character used for axis
}

// DefaultBarGraphOptions
func DefaultBarGraphOptions() BarGraphOptions {
	return BarGraphOptions{
		Width:         60,  // Default width
		BarCharacter:  "█", // Default bar character
		AxisCharacter: "─", // Default axis character
	}
}

// SetWidth
func (o *BarGraphOptions) SetWidth(width int) *BarGraphOptions {
	o.Width = width
	return o
}

// SetBarCharacter
func (o *BarGraphOptions) SetBarCharacter(char string) *BarGraphOptions {
	o.BarCharacter = char
	return o
}

// SetAxisCharacter
func (o *BarGraphOptions) SetAxisCharacter(char string) *BarGraphOptions {
	o.AxisCharacter = char
	return o
}

// LineGraphOptions
type LineGraphOptions struct {
	Width          int    // Maximum width of the graph in characters
	Height         int    // Maximum height of the graph in lines
	AxisCharacter  string // Character for axis
	PointCharacter string // Character for data points
}

// DefaultLineGraphOptions
func DefaultLineGraphOptions() LineGraphOptions {
	return LineGraphOptions{
		Width:          60,
		Height:         10,
		AxisCharacter:  "─",
		PointCharacter: "*",
	}
}

// SetWidth
func (o *LineGraphOptions) SetWidth(width int) *LineGraphOptions {
	o.Width = width
	return o
}

// SetHeight
func (o *LineGraphOptions) SetHeight(height int) *LineGraphOptions {
	o.Height = height
	return o
}

// SetAxisCharacter
func (o *LineGraphOptions) SetAxisCharacter(char string) *LineGraphOptions {
	o.AxisCharacter = char
	return o
}

// SetPointCharacter
func (o *LineGraphOptions) SetPointCharacter(char string) *LineGraphOptions {
	o.PointCharacter = char
	return o
}
