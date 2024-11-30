package app

import (
	ah "github.com/takanoriyanagitani/go-avro-head"
	util "github.com/takanoriyanagitani/go-avro-head/util"

	head "github.com/takanoriyanagitani/go-avro-head/head"
)

type App struct {
	Input  util.Io[ah.Input]
	Count  util.Io[uint64]
	Output func(ah.Input) util.Io[util.Void]
}

func (a App) ToOutputAll() util.Io[util.Void] {
	var taken util.Io[ah.Input] = head.InputToHead(
		a.Count,
	)(a.Input)
	return util.Bind(
		taken,
		a.Output,
	)
}
