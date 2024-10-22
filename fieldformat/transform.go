package fieldformat

// Transform is a function that transforms a field value.
type Transform func(fieldValue string) string

// Apply will set the field format inclusion for ref.
func (t Transform) Apply(ref *Options) {
	ref.Transform = t
}

// Wrap returns a transformation that wraps non-empty field values with the
// given prefix and suffix.
func Wrap(prefix, suffix string) Transform {
	return func(fieldValue string) string {
		if fieldValue == "" {
			return ""
		}
		return prefix + fieldValue + suffix
	}
}
