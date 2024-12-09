package dec

import (
	"bufio"
	"context"
	"io"
	"iter"
	"os"

	ha "github.com/hamba/avro/v2"
	ho "github.com/hamba/avro/v2/ocf"

	ah "github.com/takanoriyanagitani/go-avro-head"
	util "github.com/takanoriyanagitani/go-avro-head/util"
)

func ConfigToOpts(c ah.InputConfig) []ho.DecoderFunc {
	var hcfg ha.Config = ha.Config{}
	hcfg.MaxByteSliceSize = c.BlobSizeMax()
	var hapi ha.API = hcfg.Freeze()
	return []ho.DecoderFunc{
		ho.WithDecoderConfig(hapi),
	}
}

func ReaderToRowsWithOptionsHamba(
	r io.Reader,
	opts ...ho.DecoderFunc,
) (ah.Input, error) {
	dec, e := ho.NewDecoder(r, opts...)
	if nil != e {
		return ah.Input{}, e
	}

	var s ha.Schema = dec.Schema()
	var sjson ah.AvroSchemaJson = ah.AvroSchemaJson(s.String())

	var it iter.Seq2[ah.AvroRow, error] = func(
		yield func(ah.AvroRow, error) bool,
	) {
		var buf any
		var err error
		for dec.HasNext() {
			err = dec.Decode(&buf)
			if !yield(ah.AvroRow(buf), err) {
				return
			}
		}
	}

	return ah.Input{
		Rows:           it,
		AvroSchemaJson: sjson,
	}, nil
}

func ReaderToRowsWithConfig(
	r io.Reader,
	cfg ah.InputConfig,
) (ah.Input, error) {
	var opts []ho.DecoderFunc = ConfigToOpts(cfg)
	return ReaderToRowsWithOptionsHamba(r, opts...)
}

func StdinToRowsWithConfig(
	cfg ah.InputConfig,
) (ah.Input, error) {
	return ReaderToRowsWithConfig(os.Stdin, cfg)
}

func ConfigToStdinToRows(cfg ah.InputConfig) util.Io[ah.Input] {
	return func(_ context.Context) (ah.Input, error) {
		return StdinToRowsWithConfig(cfg)
	}
}

func ReaderToRows(
	r io.Reader,
) (ah.Input, error) {
	return ReaderToRowsWithOptionsHamba(r)
}

func StdinToRows() (ah.Input, error) {
	var br io.Reader = bufio.NewReader(os.Stdin)
	return ReaderToRows(br)
}

var StdinToRecords util.Io[ah.Input] = func(
	_ context.Context,
) (ah.Input, error) {
	return StdinToRows()
}
