package fieldformat

// Label returns a field format option with the given label.
func Label(label string) Options {
	return Options{Label: label}
}

// Width returns a field format option with the given width.
func Width(width int) Options {
	return Options{Width: width}
}
