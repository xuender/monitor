//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
)

func InitApp() *App {
	wire.Build(
		NewApp,
		NewProcess,
		NewPlot,
	)

	return &App{}
}
