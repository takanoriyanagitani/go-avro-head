package head

import (
	"context"
	"errors"
	"iter"

	ah "github.com/takanoriyanagitani/go-avro-head"
	util "github.com/takanoriyanagitani/go-avro-head/util"

	it "github.com/takanoriyanagitani/go-avro-head/util/iter"
)

func Take(original ah.Input, count uint64) ah.Input {
	var rows iter.Seq2[ah.AvroRow, error] = original.Rows
	var schema ah.AvroSchemaJson = original.AvroSchemaJson

	var taken iter.Seq2[ah.AvroRow, error] = it.Take2(rows, count)
	return ah.Input{
		Rows:           taken,
		AvroSchemaJson: schema,
	}
}

func InputToHead(c util.Io[uint64]) func(util.Io[ah.Input]) util.Io[ah.Input] {
	return func(i util.Io[ah.Input]) util.Io[ah.Input] {
		return func(ctx context.Context) (ah.Input, error) {
			cnt, ce := c(ctx)
			inp, ie := i(ctx)
			var taken ah.Input = Take(inp, cnt)
			return taken, errors.Join(ce, ie)
		}
	}
}
