package openapi

import (
	"cangbinggu.io/buliangren/pkg/apis/tiankuixing"
	openapinamer "k8s.io/apiserver/pkg/endpoints/openapi"
	genericapiserver "k8s.io/apiserver/pkg/server"
	openapicommon "k8s.io/kube-openapi/pkg/common"
	"k8s.io/kube-openapi/pkg/validation/spec"
)

func SetupOpenAPI(genericAPIServerConfig *genericapiserver.Config, getDefinitions openapicommon.GetOpenAPIDefinitions, title string, license string) {
	genericAPIServerConfig.OpenAPIConfig = genericapiserver.DefaultOpenAPIConfig(getDefinitions, openapinamer.NewDefinitionNamer(tiankuixing.Scheme))
	genericAPIServerConfig.OpenAPIConfig.Info.Title = title
	genericAPIServerConfig.OpenAPIConfig.Info.License = &spec.License{Name: license}
}
