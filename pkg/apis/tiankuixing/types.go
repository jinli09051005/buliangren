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

package tiankuixing

import (
	"context"
	"fmt"
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"sigs.k8s.io/apiserver-runtime/pkg/builder/resource"
	"sigs.k8s.io/apiserver-runtime/pkg/builder/resource/resourcestrategy"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// UpdateConfig
// +k8s:openapi-gen=true
type UpdateConfig struct {
	metav1.TypeMeta
	// Standard object's metadata.
	// +optional
	metav1.ObjectMeta
	// +optional
	Spec UpdateConfigSpec
	// +optional
	Status UpdateConfigStatus
}

// UpdateConfigList
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type UpdateConfigList struct {
	metav1.TypeMeta
	metav1.ListMeta

	Items []UpdateConfig
}

// UpdateConfigSpec defines the desired state of UpdateConfig
type UpdateConfigSpec struct {
	// Image Name
	// +optional
	ImageName string
	// ConfigMap Name
	// +optional
	ConfigMapName string
	// Deployment Name
	// +optional
	DeploymentName string
	// Numbers
	// +optional
	Counts int32
}

var _ resource.Object = &UpdateConfig{}
var _ resourcestrategy.Validater = &UpdateConfig{}

func (in *UpdateConfig) GetObjectMeta() *metav1.ObjectMeta {
	return &in.ObjectMeta
}

func (in *UpdateConfig) NamespaceScoped() bool {
	return false
}

func (in *UpdateConfig) New() runtime.Object {
	return &UpdateConfig{}
}

func (in *UpdateConfig) NewList() runtime.Object {
	return &UpdateConfigList{}
}

func (in *UpdateConfig) GetGroupVersionResource() schema.GroupVersionResource {
	return schema.GroupVersionResource{
		Group:    "tiankuixing.cangbinggu.io",
		Version:  "v1",
		Resource: "updateconfigs",
	}
}

func (in *UpdateConfig) IsStorageVersion() bool {
	return true
}

func (in *UpdateConfig) Validate(ctx context.Context) field.ErrorList {
	// TODO(user): Modify it, adding your API validation here.
	errors := field.ErrorList{}

	if fmt.Sprintf("%v", reflect.TypeOf(in.Spec.ImageName)) != "string" {
		errors = append(errors, field.Invalid(field.NewPath("spec", "imageName"), in.Spec.ImageName, "must be string"))
	}

	if fmt.Sprintf("%v", reflect.TypeOf(in.Spec.ConfigMapName)) != "string" {
		errors = append(errors, field.Invalid(field.NewPath("spec", "configMapName"), in.Spec.ConfigMapName, "must be string"))
	}

	if fmt.Sprintf("%v", reflect.TypeOf(in.Spec.DeploymentName)) != "string" {
		errors = append(errors, field.Invalid(field.NewPath("spec", "deploymentName"), in.Spec.DeploymentName, "must be string"))
	}

	if in.Spec.Counts > 100 {
		errors = append(errors, field.Invalid(field.NewPath("spec", "counts"), in.Spec.Counts, "must be less than 100"))
	}
	return errors
}

var _ resource.ObjectList = &UpdateConfigList{}

func (in *UpdateConfigList) GetListMeta() *metav1.ListMeta {
	return &in.ListMeta
}

// UpdateConfigStatus defines the observed state of UpdateConfig
type UpdateConfigStatus struct {
	LastUpdate      metav1.Time
	ReconcileCounts int32
}

func (in UpdateConfigStatus) SubResourceName() string {
	return "status"
}

// UpdateConfig implements ObjectWithStatusSubResource interface.
var _ resource.ObjectWithStatusSubResource = &UpdateConfig{}

func (in *UpdateConfig) GetStatus() resource.StatusSubResource {
	return in.Status
}

// UpdateConfigStatus{} implements StatusSubResource interface.
var _ resource.StatusSubResource = &UpdateConfigStatus{}

func (in UpdateConfigStatus) CopyTo(parent resource.ObjectWithStatusSubResource) {
	parent.(*UpdateConfig).Status = in
}
