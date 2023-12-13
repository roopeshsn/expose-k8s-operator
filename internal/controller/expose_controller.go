/*
Copyright 2023 roopeshsn.

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

package controller

import (
	"context"
	"fmt"
	"time"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	apiv1alpha1 "github.com/roopeshsn/expose-k8s-operator/api/v1alpha1"
)

// ExposeReconciler reconciles a Expose object
type ExposeReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=api.core.expose-k8s-operator.io,resources=exposes,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=api.core.expose-k8s-operator.io,resources=exposes/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=api.core.expose-k8s-operator.io,resources=exposes/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Expose object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.16.3/pkg/reconcile
func (r *ExposeReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	expose := &apiv1alpha1.Expose{}
	err := r.Get(ctx, req.NamespacedName, expose)
	if err != nil {
		return ctrl.Result{}, nil
	}

	fmt.Println(expose)

	// deploymentName := expose.Spec.Deployment.Name
	// replicas := expose.Spec.Deployment.Replicas

	// fmt.Printf("%s, %s", deploymentName, replicas)

	// // Getting container details to create a deployment
	// for _, container := range expose.Spec.Deployment.Containers {
	// 	fmt.Printf("%s, %s", container.name, container.image)
	// }

	return ctrl.Result{RequeueAfter: time.Duration(30 * time.Second)}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ExposeReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&apiv1alpha1.Expose{}).
		Complete(r)
}
