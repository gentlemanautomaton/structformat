package structformat

import "github.com/gentlemanautomaton/structformat/fieldformat"

// Rules hold a set of rules for formatting structs.
type Rules struct {
	// RequireInclusion means that only fields with the Include option set
	// will be included.
	RequireInclusion bool

	// FieldSeparator is appended after fields and identifiers.
	// If empty, a default field separator (a space) will be used.
	FieldSeparator string
}

// ShouldInclude returns true if the rules dictate that a field with the given
// options should be included.
func (rules Rules) ShouldInclude(field fieldformat.Options) bool {
	if field.Exclude {
		return false
	}
	if rules.RequireInclusion {
		return field.Include
	}
	return true
}

// InferRules infers a set of struct formatting rules from the given set
// of field formatting rules.
func InferRules(opts ...fieldformat.Options) Rules {
	var rules Rules
	for _, opt := range opts {
		if opt.Include {
			rules.RequireInclusion = true
			break
		}
	}
	return rules
}
