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

const (
	BlobSizeMaxDefault int = 1048576
)

type InputConfig struct {
	maxBlobSize int
}

func (c InputConfig) WithBlobSizeMax(i int) InputConfig {
	c.maxBlobSize = i
	return c
}

func (c InputConfig) BlobSizeMax() int { return c.maxBlobSize }

var InputConfigDefault InputConfig = InputConfig{}.
	WithBlobSizeMax(BlobSizeMaxDefault)
