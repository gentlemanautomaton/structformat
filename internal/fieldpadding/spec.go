package fieldpadding

import "strings"

// Default is the default padding string.
const Default = " "

// Spec is a specification for some amount of padding.
type Spec struct {
	// Length is the number of padding characters needed.
	Length int

	// StringToRepeat is a string to be repeated until the necessary number
	// of characters have been produced. If empty, the default will be used
	// instead.
	StringToRepeat string
}

// New returns a new Spec with the given length and string.
func New(length int, stringToRepeat string) Spec {
	return Spec{
		Length:         length,
		StringToRepeat: stringToRepeat,
	}
}

// String produces a set of padding characters based on the specification and
// returns it.
func (spec Spec) String() string {
	// If no padding is needed, return an empty string.
	if spec.Length == 0 {
		return ""
	}

	// Apply the default if necessary.
	stringToRepeat := spec.StringToRepeat
	if stringToRepeat == "" {
		stringToRepeat = Default
	}

	// Determine how many times the string should be repeated.
	chunk := len(stringToRepeat)
	needed := spec.Length / chunk
	if spec.Length%chunk != 0 {
		needed++
	}

	// Build the repeated string.
	s := strings.Repeat(stringToRepeat, needed)

	// If spec.Length is a not a multiple of the chunk size, slice the padding
	// down to size.
	return s[:spec.Length]
}
