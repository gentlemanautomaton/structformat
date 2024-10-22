package fieldformat

// Group returns a field format option with the given group.
func Group(group string) Options {
	return Options{Group: group}
}

// Label returns a field format option with the given label.
func Label(label string) Options {
	return Options{Label: label}
}

// Width returns a field format option with the given width.
func Width(width int) Options {
	return Options{Width: width}
}

// Indent returns a field format option with the given block indent.
func Indent(indent int) Options {
	return Options{Indent: indent}
}
