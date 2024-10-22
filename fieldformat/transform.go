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

// Prefix returns a transformation that adds a prefix to non-empty field
// values.
func Prefix(prefix string) Transform {
	return func(fieldValue string) string {
		if fieldValue == "" {
			return ""
		}
		return prefix + fieldValue
	}
}

// Suffix returns a transformation that adds a suffix to non-empty field
// values.
func Suffix(suffix string) Transform {
	return func(fieldValue string) string {
		if fieldValue == "" {
			return ""
		}
		return fieldValue + suffix
	}
}

// NumericSuffix returns a transformation that adds one of two suffixes to
// non-empty field values. A space is inserted between the value and the
// selected suffix.
//
// If the field value is "1" it uses the singular suffix, otherwise it uses
// the plural suffix.
func NumericSuffix(singular, plural string) Transform {
	return func(fieldValue string) string {
		switch fieldValue {
		case "":
			return ""
		case "1":
			return fieldValue + " " + singular
		default:
			return fieldValue + " " + plural
		}
	}
}
