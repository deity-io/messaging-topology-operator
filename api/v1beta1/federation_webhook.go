package v1beta1

import (
	"fmt"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	ctrl "sigs.k8s.io/controller-runtime"
)

func (f *Federation) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(f).
		Complete()
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// +kubebuilder:webhook:verbs=create;update,path=/validate-rabbitmq-com-v1beta1-federation,mutating=false,failurePolicy=fail,groups=rabbitmq.com,resources=federations,versions=v1beta1,name=vfederation.kb.io,sideEffects=none,admissionReviewVersions=v1

// no validation for create
func (f *Federation) ValidateCreate() error {
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (f *Federation) ValidateUpdate(old runtime.Object) error {
	oldFederation, ok := old.(*Federation)
	if !ok {
		return apierrors.NewBadRequest(fmt.Sprintf("expected a federation but got a %T", old))
	}

	detailMsg := "updates on name, vhost and rabbitmqClusterReference are all forbidden"
	if f.Spec.Name != oldFederation.Spec.Name {
		return apierrors.NewForbidden(f.GroupResource(), f.Name,
			field.Forbidden(field.NewPath("spec", "name"), detailMsg))
	}

	if f.Spec.Vhost != oldFederation.Spec.Vhost {
		return apierrors.NewForbidden(f.GroupResource(), f.Name,
			field.Forbidden(field.NewPath("spec", "vhost"), detailMsg))
	}

	if f.Spec.RabbitmqClusterReference != oldFederation.Spec.RabbitmqClusterReference {
		return apierrors.NewForbidden(f.GroupResource(), f.Name,
			field.Forbidden(field.NewPath("spec", "rabbitmqClusterReference"), detailMsg))
	}
	return nil
}

// no validation on delete
func (f *Federation) ValidateDelete() error {
	return nil
}
