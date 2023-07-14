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

package main

import (
	"cangbinggu.io/buliangren/cmd/apiserver/app"
	"math/rand"
	"os"
	"runtime"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	app.NewApp("tiankuixing").Run()
}

//import (
//	"k8s.io/klog"
//	"sigs.k8s.io/apiserver-runtime/pkg/builder"
//
//	// +kubebuilder:scaffold:resource-imports
//	externalapiv1 "cangbinggu.io/buliangren/pkg/apis/tiankuixing/v1"
//)
//
//func main() {
//	err := builder.APIServer.
//		// +kubebuilder:scaffold:resource-register
//		WithResource(&externalapiv1.updateconfig{}).
//		Execute()
//	if err != nil {
//		klog.Fatal(err)
//	}
//}
