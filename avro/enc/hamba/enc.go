package enc

import (
	"context"
	"io"
	"iter"
	"os"

	ha "github.com/hamba/avro/v2"
	ho "github.com/hamba/avro/v2/ocf"

	ah "github.com/takanoriyanagitani/go-avro-head"
	util "github.com/takanoriyanagitani/go-avro-head/util"
)

func RowsToWriter(
	ctx context.Context,
	rows iter.Seq2[ah.AvroRow, error],
	writer io.Writer,
	schema ha.Schema,
) error {
	enc, e := ho.NewEncoderWithSchema(schema, writer)
	if nil != e {
		return e
	}
	defer enc.Close()

	for row, e := range rows {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		if nil != e {
			return e
		}

		var ee error = enc.Encode(row)
		if nil != ee {
			return ee
		}
	}

	return enc.Flush()
}

func InputToWriter(
	ctx context.Context,
	input ah.Input,
	writer io.Writer,
) error {
	var s ah.AvroSchemaJson = input.AvroSchemaJson
	var rows iter.Seq2[ah.AvroRow, error] = input.Rows
	parsed, e := ha.Parse(string(s))
	if nil != e {
		return e
	}
	return RowsToWriter(
		ctx,
		rows,
		writer,
		parsed,
	)
}

func InputToStdout(
	ctx context.Context,
	input ah.Input,
) error {
	return InputToWriter(ctx, input, os.Stdout)
}

var InputToStandardOutput func(ah.Input) util.Io[util.Void] = func(
	i ah.Input,
) util.Io[util.Void] {
	return func(ctx context.Context) (util.Void, error) {
		return util.Empty, InputToStdout(ctx, i)
	}
}
