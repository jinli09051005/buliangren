package updateconfig

import (
	"cangbinggu.io/buliangren/pkg/apis/tiankuixing"
	tiankuixinginternalclient "cangbinggu.io/buliangren/pkg/generated/clientset/internalversion/typed/tiankuixing/internalversion"
	utilnames "cangbinggu.io/buliangren/pkg/staging/src/cangbinggu.io/tiananxing/utils/names"
	"context"
	"fmt"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	genericregistry "k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/storage/names"
)

type Strategy struct {
	runtime.ObjectTyper
	names.NameGenerator
	tiankuixingClient tiankuixinginternalclient.TiankuixingInterface
}

func (s Strategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
}

func (s Strategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
}

func (s Strategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	return field.ErrorList{}
}

func (s Strategy) WarningsOnCreate(ctx context.Context, obj runtime.Object) []string {
	return nil
}

func (s Strategy) NamespaceScoped() bool {
	return true
}

func (s Strategy) AllowCreateOnUpdate() bool {
	return false
}

func (s Strategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	return field.ErrorList{}
}

func (s Strategy) WarningsOnUpdate(ctx context.Context, obj, old runtime.Object) []string {
	return nil
}

func (s Strategy) Canonicalize(obj runtime.Object) {
}

func (s Strategy) AllowUnconditionalUpdate() bool {
	return false
}

func NewStrategy(tiankuixingClient tiankuixinginternalclient.TiankuixingInterface) *Strategy {
	return &Strategy{tiankuixing.Scheme, utilnames.Generator, tiankuixingClient}
}

func GetAttrs(obj runtime.Object) (labels.Set, fields.Set, error) {
	updateconfigs, ok := obj.(*tiankuixing.UpdateConfig)
	if !ok {
		return nil, nil, fmt.Errorf("not a updateconfigs")
	}
	return updateconfigs.ObjectMeta.Labels, ToSelectableFields(updateconfigs), nil
}

func ToSelectableFields(updateconfigs *tiankuixing.UpdateConfig) fields.Set {
	objectMetaFieldsSet := genericregistry.ObjectMetaFieldsSet(&updateconfigs.ObjectMeta, false)
	specificFieldsSet := fields.Set{
		"spec.imageName":         updateconfigs.Spec.ImageName,
		"spec.configMapName":     updateconfigs.Spec.ConfigMapName,
		"spec.deploymentName":    updateconfigs.Spec.DeploymentName,
		"spec.counts":            string(updateconfigs.Spec.Counts),
		"status.reconcileCounts": string(updateconfigs.Status.ReconcileCounts),
	}
	return genericregistry.MergeFieldsSets(objectMetaFieldsSet, specificFieldsSet)
}
