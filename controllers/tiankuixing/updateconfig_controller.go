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
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"time"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	tiankuixingv1 "cangbinggu.io/buliangren/pkg/apis/tiankuixing/v1"
	appsv1 "k8s.io/api/apps/v1"
)

// UpdateConfigReconciler reconciles a UpdateConfig object
type UpdateConfigReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=tiankuixing,resources=updateconfigs,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=tiankuixing,resources=updateconfigs/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=tiankuixing,resources=updateconfigs/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the UpdateConfig object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.11.0/pkg/reconcile
func (r *UpdateConfigReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	rlog := log.FromContext(ctx)
	rlog.Info("开始调谐" + req.Namespace + "/" + req.Name)

	// 获取自定义资源
	apiServerCR := &tiankuixingv1.UpdateConfig{}
	if err := r.Get(ctx, req.NamespacedName, apiServerCR); err != nil {
		if errors.IsNotFound(err) {
			rlog.Info("自定义资源已经删除了")
			return ctrl.Result{}, nil
		}
		rlog.Error(err, "无法获取自定义资源")
		return ctrl.Result{}, err
	}

	//获取调谐字段
	imageName := apiServerCR.Spec.ImageName
	configMapName := apiServerCR.Spec.ConfigMapName
	deploymentName := apiServerCR.Spec.DeploymentName
	maxCounts := apiServerCR.Spec.Counts
	reconcileCounts := apiServerCR.Status.ReconcileCounts

	if reconcileCounts >= maxCounts {
		rlog.Info("已经达到最大调谐次数，请考虑是否修改最大调谐次数")
		return ctrl.Result{}, fmt.Errorf("已经达到最大调谐次数")
	}

	//获取指定的deployment
	demoDeploy := &appsv1.Deployment{}
	if err := r.Get(ctx, types.NamespacedName{
		Namespace: req.Namespace,
		Name:      deploymentName,
	}, demoDeploy); err != nil {
		if errors.IsNotFound(err) {
			rlog.Info("deployment资源不存在")
			return ctrl.Result{}, err
		}
		rlog.Error(err, "无法获取指定deployment资源")
		return ctrl.Result{}, err
	}

	updateFlag := false
	if demoDeploy.Spec.Template.Spec.Containers[0].Image != imageName {
		demoDeploy.Spec.Template.Spec.Containers[0].Image = imageName
		updateFlag = true
	}
	if demoDeploy.Spec.Template.Spec.Volumes[0].ConfigMap.Name != configMapName {
		demoDeploy.Spec.Template.Spec.Volumes[0].ConfigMap.Name = configMapName
		updateFlag = true
	}

	if updateFlag {
		//更新deployment
		annotations := make(map[string]string)
		currentTime := time.Now()
		annotations["configmap-update-last-time"] = currentTime.Format("2006-01-02 15:04:05")
		demoDeploy.Spec.Template.ObjectMeta.Annotations = annotations

		// 绑定自定义资源和deployment
		if err := ctrl.SetControllerReference(apiServerCR, demoDeploy, r.Scheme); err != nil {
			rlog.Error(err, "unable to set deployment's owner reference")
			return ctrl.Result{}, err
		}
		if err := r.Update(ctx, demoDeploy); err != nil {
			rlog.Error(err, "无法更新指定的deployment资源")
			return ctrl.Result{}, err
		}
		// 更新自定义资源状态值
		apiServerCR.Status.LastUpdate = metav1.Time{
			Time: time.Now(),
		}
		apiServerCR.Status.ReconcileCounts = reconcileCounts + 1
		if err := r.Status().Update(ctx, apiServerCR); err != nil {
			rlog.Error(err, "无法更新自定义资源状态")
			return ctrl.Result{}, err
		}
	}

	fmt.Println("..........................")
	fmt.Println("本周期调谐结束，进入下一周期...")
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *UpdateConfigReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&tiankuixingv1.UpdateConfig{}).
		Owns(&appsv1.Deployment{}).
		Complete(r)
}
