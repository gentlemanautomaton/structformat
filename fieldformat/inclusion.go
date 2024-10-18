package fieldformat

// Inclusion is an option that indicate whether a field should be included.
type Inclusion bool

// Include is a field format option for a field that should be included.
const Include = Inclusion(true)

// Apply will set the field format inclusion for ref.
func (i Inclusion) Apply(ref *Options) {
	ref.Include = bool(i)
}

// Options will return a set of options with the field format inclusion
// applied to them.
func (i Inclusion) Options() (opts Options) {
	i.Apply(&opts)
	return
}
