package head

import (
	"iter"
)

type AvroRow any

type AvroSchemaJson string

type Input struct {
	Rows iter.Seq2[AvroRow, error]
	AvroSchemaJson
}
