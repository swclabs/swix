package boot

import (
	"go.uber.org/fx"
)

func NewRestApp(adapterConstructors ...interface{}) *fx.App {
	return fx.New(
		FxRestModule,
		fx.Provide(
			adapterConstructors...,
		),
		fx.Provide(NewServer),
		fx.Invoke(StartServer),
	)
}

func NewWorkerApp() *fx.App {
	return fx.New(
		FxWorkerModule,
		fx.Provide(
			NewWorker,
		),
		fx.Invoke(StartWorker),
	)
}
