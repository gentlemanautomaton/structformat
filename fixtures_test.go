package structformat_test

import (
	"fmt"
	"math/rand/v2"

	"github.com/gentlemanautomaton/structformat"
	"github.com/gentlemanautomaton/structformat/fieldformat"
)

type Record struct {
	A string
	B string
	C string
	D string
	E string
	F string
}

func (r Record) Format(f Format) string {
	var builder structformat.Builder
	builder.ApplyRules(structformat.InferRules(
		f.A, f.B, f.C, f.D, f.E, f.F,
	))
	builder.WriteField(r.A, f.A)
	builder.WriteField(r.B, f.B)
	builder.WriteField(r.C, f.C)
	builder.WriteField(r.D, f.D)
	builder.WriteField(r.E, f.E)
	builder.WriteField(r.F, f.F)
	return builder.String()
}

func makeRecord(index int) Record {
	return Record{
		A: fmt.Sprintf("A%d", index),
		B: fmt.Sprintf("B%d", index),
		C: fmt.Sprintf("C%d", index),
		D: fmt.Sprintf("D%d", index),
		E: fmt.Sprintf("E%d", index),
		F: fmt.Sprintf("F%d", index),
	}
}

type Format struct {
	A fieldformat.Options
	B fieldformat.Options
	C fieldformat.Options
	D fieldformat.Options
	E fieldformat.Options
	F fieldformat.Options
}

func (f Format) String() string {
	return fmt.Sprintf("%s.%s.%s.%s.%s.%s", f.A, f.B, f.C, f.D, f.E, f.F)
}

type formatGenerator struct {
	rand *rand.Rand
}

func newFormatGenerator(seed1, seed2 uint64) formatGenerator {
	source := rand.NewPCG(seed1, seed2)
	return formatGenerator{rand: rand.New(source)}
}

func (gen formatGenerator) Generate() Format {
	return Format{
		A: gen.generateFieldOptions(),
		B: gen.generateFieldOptions(),
		C: gen.generateFieldOptions(),
		D: gen.generateFieldOptions(),
		E: gen.generateFieldOptions(),
		F: gen.generateFieldOptions(),
	}
}

func (gen formatGenerator) generateFieldOptions() fieldformat.Options {
	var opts []fieldformat.Option

	if roll := gen.rand.IntN(100); roll < 10 {
		opts = append(opts, fieldformat.Group("*"))
	}

	switch roll := gen.rand.IntN(5); roll {
	default:
		opts = append(opts, fieldformat.Standard)
	case 2, 3:
		opts = append(opts, fieldformat.Note)

		if roll := gen.rand.IntN(100); roll < 30 {
			opts = append(opts, fieldformat.Label("Label"))
		}
	case 4:
		opts = append(opts, fieldformat.Primary)
	}

	switch roll := gen.rand.IntN(100); {
	case roll < 10:
		opts = append(opts, fieldformat.Exclude)
	case roll < 30:
		opts = append(opts, fieldformat.Include)
	}

	return fieldformat.Combine(opts...)
}
