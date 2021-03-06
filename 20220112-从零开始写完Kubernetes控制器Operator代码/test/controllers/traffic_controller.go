/*
Copyright 2022.

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

package controllers

import (
	"context"
	v1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/klog/v2"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	testiov1alpha1 "apiserver/api/v1alpha1"
)

// TrafficReconciler reconciles a Traffic object
type TrafficReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=test.io.apiserver,resources=traffics,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=test.io.apiserver,resources=traffics/finalizers,verbs=update
// +kubebuilder:rbac:groups=apps/v1,resources=deployments,verbs=get,list

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Traffic object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.10.0/pkg/reconcile
func (r *TrafficReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// Get deploy
	// TODO(user): your logic here
	// deploy := &v1.Deployment{}
	// err := r.Client.Get(ctx, types.NamespacedName{"xyl", "httpbin"}, deploy)
	// if err != nil {
	// 	klog.Error(err.Error())
	// 	return ctrl.Result{}, err
	// }
	// klog.Infof("deploy resource:",deploy.Name, deploy.Namespace)

	// Get CR
	traffic := &testiov1alpha1.Traffic{}
	err := r.Client.Get(ctx, types.NamespacedName{"xyl", "test-traffic"}, traffic)
	if err != nil {
		klog.Error(err.Error())
		return ctrl.Result{}, err
	}

	klog.Infof("traffic resource:",traffic.Namespace, traffic.Name)

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *TrafficReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&testiov1alpha1.Traffic{}).
		Complete(r)
}
