package fieldformat

// Exclusion is an option that indicate whether a field should be excluded.
type Exclusion bool

// Exclude is a field format option for a field that should be excluded.
const Exclude = Exclusion(true)

// Apply will set the field format exclusion for ref.
func (e Exclusion) Apply(ref *Options) {
	ref.Exclude = bool(e)
}

// Options returns a set of field format options with the exclusion.
func (e Exclusion) Options() (opts Options) {
	e.Apply(&opts)
	return
}
