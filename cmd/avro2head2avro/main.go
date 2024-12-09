package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	ah "github.com/takanoriyanagitani/go-avro-head"
	util "github.com/takanoriyanagitani/go-avro-head/util"

	dh "github.com/takanoriyanagitani/go-avro-head/avro/dec/hamba"
	eh "github.com/takanoriyanagitani/go-avro-head/avro/enc/hamba"

	ap "github.com/takanoriyanagitani/go-avro-head/app/avro2head2avro"
)

func GetEnvByKey(key string) util.Io[string] {
	return func(_ context.Context) (string, error) {
		val, found := os.LookupEnv(key)
		switch found {
		case true:
			return val, nil
		default:
			return "", fmt.Errorf("env var missing(%s)", key)
		}
	}
}

var blobSizeMaxString util.Io[string] = GetEnvByKey("ENV_BLOB_SIZE_MAX")
var blobSizeMax util.Io[int] = util.Bind(
	blobSizeMaxString,
	util.Lift(strconv.Atoi),
)

var inputConfig util.Io[ah.InputConfig] = util.Bind(
	blobSizeMax,
	util.Lift(func(i int) (ah.InputConfig, error) {
		return ah.InputConfigDefault.
			WithBlobSizeMax(i), nil
	}),
)

func StringToUint64(s string) (uint64, error) {
	return strconv.ParseUint(s, 10, 64)
}

var input util.Io[ah.Input] = util.Bind(
	inputConfig,
	func(c ah.InputConfig) util.Io[ah.Input] {
		return dh.ConfigToStdinToRows(c)
	},
)
var output func(ah.Input) util.Io[util.Void] = eh.InputToStandardOutput

var countString util.Io[string] = GetEnvByKey("ENV_COUNT")

var count util.Io[uint64] = util.Bind(
	countString,
	util.Lift(StringToUint64),
)

var app ap.App = ap.App{
	Input:  input,
	Count:  count,
	Output: output,
}

var outputAll util.Io[util.Void] = app.ToOutputAll()

func sub(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	_, e := outputAll(ctx)
	return e
}

func main() {
	e := sub(context.Background())
	if nil != e {
		log.Printf("%v\n", e)
	}
}
