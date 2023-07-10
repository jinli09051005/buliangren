package tiankuixing

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
)

var (
	Scheme         = runtime.NewScheme()
	Codecs         = serializer.NewCodecFactory(Scheme)
	ParameterCodec = runtime.NewParameterCodec(Scheme)
)

const Group = "tiankuixing.cangbinggu.io"

var SchemeGroupVersion = schema.GroupVersion{
	Group:   Group,
	Version: runtime.APIVersionInternal,
}

func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

var (
	SchemeBuilder     runtime.SchemeBuilder
	thisSchemeBuilder = &SchemeBuilder
	AddToScheme       = thisSchemeBuilder.AddToScheme
)

func init() {
	thisSchemeBuilder.Register(addKnownTypes)
}

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&UpdateConfig{},
		&UpdateConfigList{},
	)
	return nil
}
