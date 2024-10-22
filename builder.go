package structformat

import (
	"strings"

	"github.com/gentlemanautomaton/structformat/fieldformat"
	"github.com/gentlemanautomaton/structformat/internal/fieldpadding"
)

// A Builder is used to efficiently produce a formatted string for a struct.
//
// The zero value is ready to use. Do not copy a non-zero Builder.
type Builder struct {
	builder strings.Builder
	rules   Rules

	lastTypeWritten  fieldformat.Type
	lastGroupWritten string

	padding fieldpadding.Spec

	lastSkipped fieldformat.Type
	skipped     int

	divided bool

	blocks []pendingField
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

// WritePrimary writes a primary field to the builder.
func (b *Builder) WritePrimary(value string, opts ...fieldformat.Option) {
	b.WriteField(value, append([]fieldformat.Option{fieldformat.Primary}, opts...)...)
}

// WriteStandard writes a standard field to the builder.
func (b *Builder) WriteStandard(value string, opts ...fieldformat.Option) {
	b.WriteField(value, append([]fieldformat.Option{fieldformat.Standard}, opts...)...)
}

// WriteNote writes a note field to the builder.
func (b *Builder) WriteNote(value string, opts ...fieldformat.Option) {
	b.WriteField(value, append([]fieldformat.Option{fieldformat.Note}, opts...)...)
}

// WriteBlock writes a block field to the builder.
func (b *Builder) WriteBlock(value string, opts ...fieldformat.Option) {
	b.WriteField(value, append([]fieldformat.Option{fieldformat.Block}, opts...)...)
}

// WriteField writes a field to the builder.
func (b *Builder) WriteField(value string, opts ...fieldformat.Option) {
	// Combine field format options.
	field := fieldformat.Combine(opts...)

	// Apply value transformations.
	if field.Transform != nil {
		value = field.Transform(value)
	}

	// Apply defaults.
	if field.Type == fieldformat.DefaultType {
		field.Type = fieldformat.Standard
	}

	// If the field should not be included, ignore it.
	if !b.rules.ShouldInclude(field) {
		return
	}

	// If the field is a block, add it to the pending blocks list and exit.
	if field.Type == fieldformat.Block {
		// Only non-empty blocks needs to be processed.
		if value != "" {
			b.blocks = append(b.blocks, pendingField{
				Value:   value,
				Options: field,
			})
		}
		return
	}

	// Check whether this non-block field is missing.
	if value == "" {
		// If the value is missing and or it doesn't have a fixed width,
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
	b.prepareFor(field.Type, field.Group)

	// Update information and state after this field has been written.
	defer b.completeFor(field.Type, field.Group)

	// Calculate the padding needed for the value, if any.
	padding := fieldpadding.New(field.Width-len(value), field.Padding)

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

// Divide divides any subsequent fields from any preceeding fields by a colon.
func (b *Builder) Divide() {
	b.divided = true
}

// prepareFor is called before a field value is written.
func (b *Builder) prepareFor(nextType fieldformat.Type, nextGroup string) {
	if b.lastGroupWritten != nextGroup {
		b.Divide()
	}

	if nextType == fieldformat.Block {
		if b.lastTypeWritten == fieldformat.Note {
			b.builder.WriteString(")")
		}
		if b.lastTypeWritten != fieldformat.DefaultType {
			b.builder.WriteString("\n")
		}
		return
	}

	switch b.lastTypeWritten {
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
		if nextType == fieldformat.Note {
			if b.divided {
				b.builder.WriteString("):")
				b.finishPadding()
				b.builder.WriteString(b.fieldSeparator())
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

	if nextType == fieldformat.Note {
		if b.lastTypeWritten != fieldformat.Note || b.divided {
			b.builder.WriteString("(")
		}
	}
}

// completeFor is called after a field value has been written.
//
// It updates information about the last field written, skip state and
// division state.
func (b *Builder) completeFor(lastType fieldformat.Type, lastGroup string) {
	// Record information about the last field that has been written.
	b.lastTypeWritten = lastType
	b.lastGroupWritten = lastGroup

	// Clear out old skip values.
	b.lastSkipped = fieldformat.DefaultType
	b.skipped = 0

	// Reset the divided state.
	b.divided = false
}

// writeBlocks writes any blocks that have been deferred.
func (b *Builder) writeBlocks() {
	for _, field := range b.blocks {
		// Complete any tasks related to the last field that was written.
		// This will also add a newline before the impending block.
		b.prepareFor(fieldformat.Block, field.Options.Group)

		// If an indent has been specified for the block, add it.
		indent := strings.Repeat(" ", field.Options.Indent)
		if field.Options.Indent > 0 {
			b.builder.WriteString(indent)
		}

		// Write the field label, if present.
		if field.Options.Label != "" {
			b.builder.WriteString(field.Options.Label)
			b.builder.WriteString(": ")
		}

		// Write the field value.
		// TODO: Consider reformatting lines that exceed field.Options.Width
		if indent != "" {
			// Add the indent to the beginning of each line.
			for i, line := range strings.Split(field.Value, "\n") {
				if i > 0 {
					b.builder.WriteString("\n")
					b.builder.WriteString(indent)
				}
				b.builder.WriteString(line)
			}
			//b.builder.WriteString(strings.ReplaceAll(field.Value, "\n", strings.Repeat(" ", field.Options.Indent)+"\n"))
		} else {
			b.builder.WriteString(field.Value)
		}

		// Update information and state after this field has been written.
		b.completeFor(fieldformat.Block, field.Options.Group)
	}

	// Reset the pending blocks list after writing them.
	b.blocks = b.blocks[:0]
}

func (b *Builder) finishPadding() {
	s := b.padding.String()
	if s != "" {
		b.builder.WriteString(s)
	}
}

func (b *Builder) finish() {
	b.writeBlocks()
	if b.lastTypeWritten == fieldformat.Note {
		b.builder.WriteString(")")
		b.lastTypeWritten = fieldformat.Standard
	}
}

func (b *Builder) fieldSeparator() string {
	if b.rules.FieldSeparator != "" {
		return b.rules.FieldSeparator
	}
	return " "
}

type pendingField struct {
	Value   string
	Options fieldformat.Options
}
