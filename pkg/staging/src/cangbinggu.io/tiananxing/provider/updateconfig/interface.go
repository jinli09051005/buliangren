package updateconfig

import (
	tiankuixingv1client "cangbinggu.io/buliangren/pkg/generated/clientset/versioned/typed/tiankuixing/v1"
	"k8s.io/apiserver/pkg/server/mux"
)

type APIProvider interface {
	RegisterHandler(mux *mux.PathRecorderMux)
}

type Provider interface {
	Name() string

	APIProvider
}

var _ Provider = &DelegateProvider{}

type DelegateProvider struct {
	ProviderName string

	PlatformClient tiankuixingv1client.TiankuixingV1Interface
}

func (p *DelegateProvider) Name() string {
	if p.ProviderName == "" {
		return "unknown"
	}
	return p.ProviderName
}

func (p *DelegateProvider) RegisterHandler(mux *mux.PathRecorderMux) {
}
