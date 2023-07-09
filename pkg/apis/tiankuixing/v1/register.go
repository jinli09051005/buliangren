/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const (
	Group   = "tiankuixing.cangbinggu.io"
	Version = "v1"
)

var SchemeGroupVersion = schema.GroupVersion{
	Group:   Group,
	Version: Version,
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
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	scheme.AddKnownTypes(SchemeGroupVersion,
		&UpdateConfig{},
		&UpdateConfigList{},
	)
	return nil
}
