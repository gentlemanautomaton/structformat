package structformat

import (
	"strings"

	"github.com/gentlemanautomaton/structformat/fieldformat"
	"github.com/gentlemanautomaton/structformat/internal/fieldpadding"
)

// A Builder is used to efficiently produce a formatted string for a struct
// using [Builder.Write] methods.
//
// The zero value is ready to use. Do not copy a non-zero Builder.
type Builder struct {
	builder strings.Builder
	rules   Rules

	lastWritten fieldformat.Type
	padding     fieldpadding.Spec

	lastSkipped fieldformat.Type
	skipped     int

	divided bool
}

// ApplyRules applies the given rules to the builder.
func (b *Builder) ApplyRules(rules Rules) {
	b.rules = rules
}

// String returns a completed string from the builder.
func (b *Builder) String() string {
	b.finish()
	return b.builder.String()
}

// Divide divides any subsequent fields from any preceeding fields by a colon.
func (b *Builder) Divide() {
	b.divided = true
}

// WriteField writes a field to the builder.
func (b *Builder) WriteField(value string, opts ...fieldformat.Option) {
	// Combine field format options.
	field := fieldformat.Combine(opts...)

	// Apply defaults.
	if field.Type == fieldformat.DefaultType {
		field.Type = fieldformat.Standard
	}

	// If the field should not be included, ignore it.
	if !b.rules.ShouldInclude(field) {
		return
	}

	// Check whether this field is missing.
	if value == "" {
		// If the value is missing and the it doesn't have a fixed width,
		// omit it entirely.
		if field.Width < 1 {
			return
		}

		// If the field is missing but it has a fixed width, keep track
		// of the number of characters we need to skip, then move on.
		b.skipped += field.Width

		if field.Label != "" {
			b.skipped += len(field.Label) + 2
		}

		switch field.Type {
		case fieldformat.Primary:
			b.skipped += 1 + len(b.fieldSeparator())
		case fieldformat.Standard:
			b.skipped += len(b.fieldSeparator())
		case fieldformat.Note:
			if b.lastSkipped == fieldformat.Note {
				b.skipped += 2
			} else {
				b.skipped += 1
				b.skipped += len(b.fieldSeparator())
			}
		}

		b.lastSkipped = field.Type

		return
	}

	// Complete any tasks related to the last field that was written.
	b.prepareFor(field.Type)

	// Calculate the padding needed for the value, if any.
	padding := fieldpadding.New(field.Width-len(value), field.Padding)

	// Record information about the value that is about to be written.
	b.lastWritten = field.Type

	// Clear out old skip values.
	b.lastSkipped = fieldformat.DefaultType
	b.skipped = 0

	// Write the field label, if present.
	if field.Label != "" {
		b.builder.WriteString(field.Label)
		b.builder.WriteString(": ")
	}

	// If no padding is needed, simply write the value.
	if padding.Length < 1 {
		b.padding = fieldpadding.Spec{}
		b.builder.WriteString(value)
		return
	}

	// Write the value and apply or record the padding, depending on its
	// alignment.
	switch field.Alignment {
	default:
		// The field is either left-aligned or has a default alignment.
		// This means the padding goes on the right, and won't be written
		// until the next value.
		b.padding = padding
		b.builder.WriteString(value)
	case fieldformat.Right:
		// The field is right-aligned. This means the padding goes on
		// the left.
		b.padding = fieldpadding.Spec{}
		b.builder.WriteString(padding.String())
		b.builder.WriteString(value)
	}
}

func (b *Builder) prepareFor(next fieldformat.Type) {
	switch b.lastWritten {
	case fieldformat.Primary:
		b.builder.WriteString(":")
		b.finishPadding()
		b.builder.WriteString(b.fieldSeparator())
	case fieldformat.Standard:
		if b.divided {
			b.builder.WriteString(":")
		}
		b.finishPadding()
		b.builder.WriteString(b.fieldSeparator())
	case fieldformat.Note:
		if next == fieldformat.Note {
			if b.divided {
				b.builder.WriteString("):")
				b.finishPadding()
			} else {
				b.builder.WriteString(",")
				b.finishPadding()
				b.builder.WriteString(" ")
			}
		} else {
			b.builder.WriteString(")")
			if b.divided {
				b.builder.WriteString(":")
			}
			b.finishPadding()
			b.builder.WriteString(b.fieldSeparator())
		}
	}

	if b.skipped > 0 {
		if b.lastSkipped == fieldformat.Note {
			b.skipped++
		}
		b.builder.WriteString(strings.Repeat(" ", b.skipped))
	}

	if next == fieldformat.Note {
		if b.lastWritten != fieldformat.Note || b.divided {
			b.builder.WriteString("(")
		}
	}
}

func (b *Builder) finishPadding() {
	s := b.padding.String()
	if s != "" {
		b.builder.WriteString(s)
	}
}

func (b *Builder) finish() {
	if b.lastWritten == fieldformat.Note {
		b.builder.WriteString(")")
		b.lastWritten = fieldformat.Standard
	}
}

func (b *Builder) fieldSeparator() string {
	if b.rules.FieldSeparator != "" {
		return b.rules.FieldSeparator
	}
	return " "
}
