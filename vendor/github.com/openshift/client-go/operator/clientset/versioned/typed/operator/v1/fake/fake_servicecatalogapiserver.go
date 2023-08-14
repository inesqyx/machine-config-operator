// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1 "github.com/openshift/api/operator/v1"
	operatorv1 "github.com/openshift/client-go/operator/applyconfigurations/operator/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeServiceCatalogAPIServers implements ServiceCatalogAPIServerInterface
type FakeServiceCatalogAPIServers struct {
	Fake *FakeOperatorV1
}

var servicecatalogapiserversResource = v1.SchemeGroupVersion.WithResource("servicecatalogapiservers")

var servicecatalogapiserversKind = v1.SchemeGroupVersion.WithKind("ServiceCatalogAPIServer")

// Get takes name of the serviceCatalogAPIServer, and returns the corresponding serviceCatalogAPIServer object, and an error if there is any.
func (c *FakeServiceCatalogAPIServers) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ServiceCatalogAPIServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(servicecatalogapiserversResource, name), &v1.ServiceCatalogAPIServer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ServiceCatalogAPIServer), err
}

// List takes label and field selectors, and returns the list of ServiceCatalogAPIServers that match those selectors.
func (c *FakeServiceCatalogAPIServers) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ServiceCatalogAPIServerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(servicecatalogapiserversResource, servicecatalogapiserversKind, opts), &v1.ServiceCatalogAPIServerList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.ServiceCatalogAPIServerList{ListMeta: obj.(*v1.ServiceCatalogAPIServerList).ListMeta}
	for _, item := range obj.(*v1.ServiceCatalogAPIServerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested serviceCatalogAPIServers.
func (c *FakeServiceCatalogAPIServers) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(servicecatalogapiserversResource, opts))
}

// Create takes the representation of a serviceCatalogAPIServer and creates it.  Returns the server's representation of the serviceCatalogAPIServer, and an error, if there is any.
func (c *FakeServiceCatalogAPIServers) Create(ctx context.Context, serviceCatalogAPIServer *v1.ServiceCatalogAPIServer, opts metav1.CreateOptions) (result *v1.ServiceCatalogAPIServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(servicecatalogapiserversResource, serviceCatalogAPIServer), &v1.ServiceCatalogAPIServer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ServiceCatalogAPIServer), err
}

// Update takes the representation of a serviceCatalogAPIServer and updates it. Returns the server's representation of the serviceCatalogAPIServer, and an error, if there is any.
func (c *FakeServiceCatalogAPIServers) Update(ctx context.Context, serviceCatalogAPIServer *v1.ServiceCatalogAPIServer, opts metav1.UpdateOptions) (result *v1.ServiceCatalogAPIServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(servicecatalogapiserversResource, serviceCatalogAPIServer), &v1.ServiceCatalogAPIServer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ServiceCatalogAPIServer), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeServiceCatalogAPIServers) UpdateStatus(ctx context.Context, serviceCatalogAPIServer *v1.ServiceCatalogAPIServer, opts metav1.UpdateOptions) (*v1.ServiceCatalogAPIServer, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(servicecatalogapiserversResource, "status", serviceCatalogAPIServer), &v1.ServiceCatalogAPIServer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ServiceCatalogAPIServer), err
}

// Delete takes name of the serviceCatalogAPIServer and deletes it. Returns an error if one occurs.
func (c *FakeServiceCatalogAPIServers) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(servicecatalogapiserversResource, name, opts), &v1.ServiceCatalogAPIServer{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeServiceCatalogAPIServers) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(servicecatalogapiserversResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1.ServiceCatalogAPIServerList{})
	return err
}

// Patch applies the patch and returns the patched serviceCatalogAPIServer.
func (c *FakeServiceCatalogAPIServers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ServiceCatalogAPIServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(servicecatalogapiserversResource, name, pt, data, subresources...), &v1.ServiceCatalogAPIServer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ServiceCatalogAPIServer), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied serviceCatalogAPIServer.
func (c *FakeServiceCatalogAPIServers) Apply(ctx context.Context, serviceCatalogAPIServer *operatorv1.ServiceCatalogAPIServerApplyConfiguration, opts metav1.ApplyOptions) (result *v1.ServiceCatalogAPIServer, err error) {
	if serviceCatalogAPIServer == nil {
		return nil, fmt.Errorf("serviceCatalogAPIServer provided to Apply must not be nil")
	}
	data, err := json.Marshal(serviceCatalogAPIServer)
	if err != nil {
		return nil, err
	}
	name := serviceCatalogAPIServer.Name
	if name == nil {
		return nil, fmt.Errorf("serviceCatalogAPIServer.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(servicecatalogapiserversResource, *name, types.ApplyPatchType, data), &v1.ServiceCatalogAPIServer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ServiceCatalogAPIServer), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeServiceCatalogAPIServers) ApplyStatus(ctx context.Context, serviceCatalogAPIServer *operatorv1.ServiceCatalogAPIServerApplyConfiguration, opts metav1.ApplyOptions) (result *v1.ServiceCatalogAPIServer, err error) {
	if serviceCatalogAPIServer == nil {
		return nil, fmt.Errorf("serviceCatalogAPIServer provided to Apply must not be nil")
	}
	data, err := json.Marshal(serviceCatalogAPIServer)
	if err != nil {
		return nil, err
	}
	name := serviceCatalogAPIServer.Name
	if name == nil {
		return nil, fmt.Errorf("serviceCatalogAPIServer.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(servicecatalogapiserversResource, *name, types.ApplyPatchType, data, "status"), &v1.ServiceCatalogAPIServer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ServiceCatalogAPIServer), err
}
