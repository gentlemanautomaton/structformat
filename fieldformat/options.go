package fieldformat

// Option is an interface for type that can apply themselves to field
// format options.
type Option interface {
	Apply(*Options)
}

// Options holds a set of formatting options for a field.
type Options struct {
	// Label is a label for the field.
	Label string

	// Type is the field format type.
	Type Type

	// Include indicates that a field should be included.
	Include bool

	// Exclude indicates that a field should be excluded.
	Exclude bool

	// Alignment is the alignment of the field, which determines how padding
	// should be applied.
	Alignment Alignment

	// Width is the minimum width of the field.
	Width int

	// Padding is the padding used to fill field values that are less than
	// the minimum width. If an empty string is specified, the default padding
	// character (a space) will be used.
	Padding string
}

// AdjustWidth ensures that opt.Width is at least width.
func (opts *Options) AdjustWidth(width int) {
	if width > opts.Width {
		opts.Width = width
	}
}

// Apply copies all non-default options in opts to ref.
func (opts Options) Apply(ref *Options) {
	if opts.Label != "" {
		ref.Label = opts.Label
	}
	if opts.Type != DefaultType {
		ref.Type = opts.Type
	}
	if opts.Width > ref.Width {
		ref.Width = opts.Width
	}
	if opts.Padding != "" {
		ref.Padding = opts.Padding
	}
	if opts.Alignment != DefaultAlignment {
		ref.Alignment = opts.Alignment
	}
	if opts.Include {
		ref.Include = true
	}
	if opts.Exclude {
		ref.Exclude = true
	}
}

// Combine combines the given set of options.
//
// If there is a conflict for a particular option, the last definition for the
// option wins.
func Combine(opts ...Option) Options {
	var combined Options
	for _, opt := range opts {
		opt.Apply(&combined)
	}
	return combined
}
