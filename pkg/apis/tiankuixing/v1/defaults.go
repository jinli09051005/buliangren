package v1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

func addDefaultingFuncs(scheme *runtime.Scheme) error {
	return RegisterDefaults(scheme)
}

func SetDefaults_UpdateConfigSpec(obj *UpdateConfigSpec) {
	if obj.Counts == 0 {
		obj.Counts = 1
	}
}
