package updateconfig

import (
	"k8s.io/apiserver/pkg/server/mux"
	"sync"
)

var (
	providersMu sync.RWMutex
	providers   = make(map[string]Provider)
)

func RegisterHandler(mux *mux.PathRecorderMux) {
	for _, p := range providers {
		p.RegisterHandler(mux)
	}
}
