package fieldformat

// Alignment defines the alignment value of a field.
type Alignment uint8

// Alignments.
const (
	DefaultAlignment Alignment = iota

	// Left indicates that a field should be left-aligned.
	Left

	// Right indicates that a field should be right-aligned.
	Right
)

// Apply will set the field format alignment for ref.
func (a Alignment) Apply(ref *Options) {
	ref.Alignment = a
}

// Options will return a set of options with the field format alignment
// applied to them.
func (a Alignment) Options() (opts Options) {
	a.Apply(&opts)
	return
}
