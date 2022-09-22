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
	"fmt"
	demov1 "github.com/Gentleelephant/operator-test/api/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// DemoKindReconciler reconciles a DemoKind object
type DemoKindReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=demo.my.domain,resources=demokinds,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=demo.my.domain,resources=demokinds/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=demo.my.domain,resources=demokinds/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the DemoKind object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.12.2/pkg/reconcile
func (r *DemoKindReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// TODO(user): your logic here
	var demoKind demov1.DemoKind
	err := r.Get(ctx, req.NamespacedName, &demoKind)
	if err != nil {
		fmt.Println("get demoKind error", err)
		return ctrl.Result{}, err
	}
	// 创建deployment
	//var replicas int32 = 3
	//deployment := apps.Deployment{
	//	TypeMeta: metav1.TypeMeta{
	//		Kind:       "Deployment",
	//		APIVersion: "apps/v1beta2",
	//	},
	//	ObjectMeta: metav1.ObjectMeta{
	//		Name:                       "test-create-deployment",
	//		Namespace:                  "default",
	//		ResourceVersion:            "",
	//		Generation:                 0,
	//		CreationTimestamp:          metav1.Time{},
	//		DeletionTimestamp:          nil,
	//		DeletionGracePeriodSeconds: nil,
	//		Labels:                     nil,
	//		Annotations: map[string]string{
	//			"test": "test",
	//		},
	//		OwnerReferences:           nil,
	//		Finalizers:                nil,
	//		ZZZ_DeprecatedClusterName: "",
	//		ManagedFields:             nil,
	//	},
	//	Spec: apps.DeploymentSpec{
	//		Replicas: &replicas,
	//		Selector: nil,
	//		Template: v1.PodTemplateSpec{
	//			ObjectMeta: metav1.ObjectMeta{
	//				Name:                       "",
	//				GenerateName:               "",
	//				Namespace:                  "",
	//				SelfLink:                   "",
	//				UID:                        "",
	//				ResourceVersion:            "",
	//				Generation:                 0,
	//				CreationTimestamp:          metav1.Time{},
	//				DeletionTimestamp:          nil,
	//				DeletionGracePeriodSeconds: nil,
	//				Labels:                     nil,
	//				Annotations:                nil,
	//				OwnerReferences:            nil,
	//				Finalizers:                 nil,
	//				ZZZ_DeprecatedClusterName:  "",
	//				ManagedFields:              nil,
	//			},
	//			Spec: v1.PodSpec{
	//				Volumes:                       nil,
	//				InitContainers:                nil,
	//				Containers:                    nil,
	//				EphemeralContainers:           nil,
	//				RestartPolicy:                 "",
	//				TerminationGracePeriodSeconds: nil,
	//				ActiveDeadlineSeconds:         nil,
	//				DNSPolicy:                     "",
	//				NodeSelector:                  nil,
	//				ServiceAccountName:            "",
	//				DeprecatedServiceAccount:      "",
	//				AutomountServiceAccountToken:  nil,
	//				NodeName:                      "",
	//				HostNetwork:                   false,
	//				HostPID:                       false,
	//				HostIPC:                       false,
	//				ShareProcessNamespace:         nil,
	//				SecurityContext:               nil,
	//				ImagePullSecrets:              nil,
	//				Hostname:                      "",
	//				Subdomain:                     "",
	//				Affinity:                      nil,
	//				SchedulerName:                 "",
	//				Tolerations:                   nil,
	//				HostAliases:                   nil,
	//				PriorityClassName:             "",
	//				Priority:                      nil,
	//				DNSConfig:                     nil,
	//				ReadinessGates:                nil,
	//				RuntimeClassName:              nil,
	//				EnableServiceLinks:            nil,
	//				PreemptionPolicy:              nil,
	//				Overhead:                      nil,
	//				TopologySpreadConstraints:     nil,
	//				SetHostnameAsFQDN:             nil,
	//				OS:                            nil,
	//			},
	//		},
	//		Strategy:                apps.DeploymentStrategy{},
	//		MinReadySeconds:         0,
	//		RevisionHistoryLimit:    nil,
	//		Paused:                  false,
	//		ProgressDeadlineSeconds: nil,
	//	},
	//}
	r.Create(ctx, &demoKind)
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *DemoKindReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&demov1.DemoKind{}).
		Complete(r)
}
