// Code generated by injection-gen. DO NOT EDIT.

package policy

import (
	"context"

	factory "github.com/kyma-project/kyma/components/event-sources/client/generated/injection/istio/informers/factory"
	v1alpha1 "istio.io/client-go/pkg/informers/externalversions/authentication/v1alpha1"
	controller "knative.dev/pkg/controller"
	injection "knative.dev/pkg/injection"
	logging "knative.dev/pkg/logging"
)

func init() {
	injection.Default.RegisterInformer(withInformer)
}

// Key is used for associating the Informer inside the context.Context.
type Key struct{}

func withInformer(ctx context.Context) (context.Context, controller.Informer) {
	f := factory.Get(ctx)
	inf := f.Authentication().V1alpha1().Policies()
	return context.WithValue(ctx, Key{}, inf), inf.Informer()
}

// Get extracts the typed informer from the context.
func Get(ctx context.Context) v1alpha1.PolicyInformer {
	untyped := ctx.Value(Key{})
	if untyped == nil {
		logging.FromContext(ctx).Panic(
			"Unable to fetch istio.io/client-go/pkg/informers/externalversions/authentication/v1alpha1.PolicyInformer from context.")
	}
	return untyped.(v1alpha1.PolicyInformer)
}
