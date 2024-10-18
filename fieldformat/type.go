package fieldformat

// Type represents a type of field formatting.
type Type uint8

// Field types.
const (
	DefaultType Type = iota

	// Primary fields values are followed by a colon.
	Primary

	// Standard fields don't have any special adornment.
	Standard

	// Notes are wrapped in parentheses. Successive notes are separated
	// by commas.
	Note
)

// Apply will set the field format type for ref.
func (t Type) Apply(ref *Options) {
	ref.Type = t
}

// Options will return a set of options with the field format type
// applied to them.
func (t Type) Options() (opts Options) {
	t.Apply(&opts)
	return
}
