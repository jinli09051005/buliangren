package app

import "cangbinggu.io/buliangren/cmd/apiserver/app/config"

func Run(cfg *config.Config, stopCh <-chan struct{}) error {
	server, err := CreateServerChain(cfg)
	if err != nil {
		return err
	}

	return server.PrepareRun().Run(stopCh)
}
