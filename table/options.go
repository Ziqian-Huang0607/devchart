package table

type Alignment int

const (
	AlignLeft Alignment = iota
	AlignCenter
	AlignRight
)

// BorderStyle
type BorderStyle int

const (
	BorderSingle BorderStyle = iota
	BorderDouble
	BorderNone
)

// Customization options
type Options struct {
	BorderStyle BorderStyle
	Alignment   Alignment
	Padding     int // Padding around cell content (spaces)
}

// DefaultOptions
func DefaultOptions() Options {
	return Options{
		BorderStyle: BorderSingle,
		Alignment:   AlignLeft,
		Padding:     1,
	}
}

// SetBorderStyle
func (o *Options) SetBorderStyle(style BorderStyle) *Options {
	o.BorderStyle = style
	return o
}

// SetAlignment
func (o *Options) SetAlignment(alignment Alignment) *Options {
	o.Alignment = alignment
	return o
}

// SetPadding
func (o *Options) SetPadding(padding int) *Options {
	o.Padding = padding
	return o
}
