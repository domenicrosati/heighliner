package k8sutils

import (
	"time"

	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
)

// CRDWatcher watches a CRD for the desired vs actual state.
type CRDWatcher struct {
	rc        *rest.RESTClient
	namespace string
	resource  *CustomResource
	handler   cache.ResourceEventHandlerFuncs
}

// NewCRDWatcher returns a new watcher that can be used to watch CRDs.
func NewCRDWatcher(rc *rest.RESTClient, namespace string, resource *CustomResource, handler cache.ResourceEventHandlerFuncs) *CRDWatcher {
	return &CRDWatcher{
		rc:        rc,
		namespace: namespace,
		resource:  resource,
		handler:   handler,
	}
}

// Watch starts watching the CRDs associated with the CRDWatcher.
func (w *CRDWatcher) Watch(done <-chan struct{}) error {
	source := cache.NewListWatchFromClient(
		w.rc,
		w.resource.Plural,
		w.namespace,
		fields.Everything(),
	)

	resyncPeriod := 10 * time.Second
	_, controller := cache.NewInformer(
		source,
		w.resource.Object,
		resyncPeriod,
		w.handler,
	)

	go controller.Run(done)
	return nil
}
