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

package v1

import (
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var demokindlog = logf.Log.WithName("demokind-resource")

func (r *DemoKind) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

//+kubebuilder:webhook:path=/mutate-demo-my-domain-v1-demokind,mutating=true,failurePolicy=fail,sideEffects=None,groups=demo.my.domain,resources=demokinds,verbs=create;update,versions=v1,name=mdemokind.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &DemoKind{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *DemoKind) Default() {
	demokindlog.Info("default", "name", r.Name)

	// TODO(user): fill in your defaulting logic.
	if r.Spec.Image == "" {
		r.Spec.Image = "This id a default Image"
		demokindlog.Info("webhook set the default image!")
	} else {
		demokindlog.Info("webhook did not set the default image!")
	}
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-demo-my-domain-v1-demokind,mutating=false,failurePolicy=fail,sideEffects=None,groups=demo.my.domain,resources=demokinds,verbs=create;update,versions=v1,name=vdemokind.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &DemoKind{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *DemoKind) ValidateCreate() error {
	demokindlog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	return r.validateDemoKined()
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *DemoKind) ValidateUpdate(old runtime.Object) error {
	demokindlog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return r.validateDemoKined()
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *DemoKind) ValidateDelete() error {
	demokindlog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}

func (r *DemoKind) validateDemoKined() error {
	var err field.ErrorList
	if len(r.Spec.Image) < 3 {
		err = append(err, field.Invalid(field.NewPath("spec").Child("image"), r.Spec.Image, "image name must be at least 3 characters"))
		return apierrors.NewInvalid(schema.GroupKind{Group: "demo.my.domain", Kind: "DemoKind"}, r.Name, err)
	} else {
		demokindlog.Info("demokind is valid!")
		return nil
	}
}
