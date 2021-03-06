package nsdelete

import (
	"context"

	syncv1 "github.com/darkowlzz/operator-toolkit/controller/sync/v1"
	"github.com/go-logr/logr"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/label"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/storageos/api-manager/internal/pkg/storageos"
)

// Controller implements the SyncReconciler contoller interface, deleting
// namespaces in StorageOS when they have been detected as deleted in
// Kubernetes.
type Controller struct {
	api storageos.NamespaceDeleter
	log logr.Logger
}

var _ syncv1.Controller = &Controller{}

// NewController returns a Controller that implements namespace garbage
// collection in StorageOS.
func NewController(api storageos.NamespaceDeleter, log logr.Logger) (*Controller, error) {
	return &Controller{api: api, log: log}, nil
}

// Ensure is a no-op.  We only care about deletes.
func (c Controller) Ensure(ctx context.Context, obj client.Object) error {
	return nil
}

// Delete receives a k8s object that's been deleted and calls the StorageOS api
// to remove it from management.
func (c Controller) Delete(ctx context.Context, obj client.Object) error {
	tr := otel.Tracer("namespace-delete")
	ctx, span := tr.Start(ctx, "namespace delete")
	span.SetAttributes(label.String("name", obj.GetName()))
	defer span.End()

	ctx, cancel := context.WithTimeout(ctx, storageos.DefaultRequestTimeout)
	defer cancel()

	err := c.api.DeleteNamespace(ctx, client.ObjectKeyFromObject(obj))
	if err != nil && err != storageos.ErrNamespaceNotFound {
		span.RecordError(err)
		return err
	}
	span.SetStatus(codes.Ok, "namespace decommissioned in storageos")
	c.log.Info("namespace decommissioned in storageos", "name", obj.GetName())
	return nil
}

// List returns a list of namespaces known to StorageOS, as NamespacedNames.
// This is used for garbage collection and can be expensive. The garbage
// collector is run in a separate goroutine periodically, not affecting the main
// reconciliation control-loop.
func (c Controller) List(ctx context.Context) ([]types.NamespacedName, error) {
	tr := otel.Tracer("namespace-delete")
	ctx, span := tr.Start(ctx, "namespace list")
	defer span.End()

	ctx, cancel := context.WithTimeout(ctx, storageos.DefaultRequestTimeout)
	defer cancel()

	objects, err := c.api.ListNamespaces(ctx)
	if err != nil {
		span.RecordError(err)
		return nil, err
	}
	span.SetAttributes(label.Int("count", len(objects)))
	span.SetStatus(codes.Ok, "listed namespaces")
	return storageos.ObjectKeys(objects), nil
}
