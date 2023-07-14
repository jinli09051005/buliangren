package app

import (
	"cangbinggu.io/buliangren/cmd/apiserver/app/config"
	"cangbinggu.io/buliangren/cmd/apiserver/app/options"
	"cangbinggu.io/buliangren/pkg/staging/src/cangbinggu.io/tiananxing/app"
	k8sapiserver "k8s.io/apiserver/pkg/server"
)

const commandDesc = "Debug"

func NewApp(basename string) *app.App {
	opts := options.NewOptions(basename)
	tiankuixing := app.NewApp("CangBingGu Buliangren Tiankuixing API Server", basename,
		app.WithOptions(opts),
		app.WithDescription(commandDesc),
		app.WithRunFunc(run(opts)),
	)
	return tiankuixing
}

func run(opts *options.Options) app.RunFunc {
	return func(basename string) error {
		if err := opts.Complete(); err != nil {
			return err
		}

		cfg, err := config.CreateConfigFromOptions(basename, opts)
		if err != nil {
			return err
		}

		stopCh := k8sapiserver.SetupSignalHandler()
		return Run(cfg, stopCh)
	}
}
