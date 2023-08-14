// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1 "github.com/openshift/api/machineconfiguration/v1"
	machineconfigurationv1 "github.com/openshift/client-go/machineconfiguration/applyconfigurations/machineconfiguration/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeKubeletConfigs implements KubeletConfigInterface
type FakeKubeletConfigs struct {
	Fake *FakeMachineconfigurationV1
}

var kubeletconfigsResource = v1.SchemeGroupVersion.WithResource("kubeletconfigs")

var kubeletconfigsKind = v1.SchemeGroupVersion.WithKind("KubeletConfig")

// Get takes name of the kubeletConfig, and returns the corresponding kubeletConfig object, and an error if there is any.
func (c *FakeKubeletConfigs) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.KubeletConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(kubeletconfigsResource, name), &v1.KubeletConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.KubeletConfig), err
}

// List takes label and field selectors, and returns the list of KubeletConfigs that match those selectors.
func (c *FakeKubeletConfigs) List(ctx context.Context, opts metav1.ListOptions) (result *v1.KubeletConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(kubeletconfigsResource, kubeletconfigsKind, opts), &v1.KubeletConfigList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.KubeletConfigList{ListMeta: obj.(*v1.KubeletConfigList).ListMeta}
	for _, item := range obj.(*v1.KubeletConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested kubeletConfigs.
func (c *FakeKubeletConfigs) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(kubeletconfigsResource, opts))
}

// Create takes the representation of a kubeletConfig and creates it.  Returns the server's representation of the kubeletConfig, and an error, if there is any.
func (c *FakeKubeletConfigs) Create(ctx context.Context, kubeletConfig *v1.KubeletConfig, opts metav1.CreateOptions) (result *v1.KubeletConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(kubeletconfigsResource, kubeletConfig), &v1.KubeletConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.KubeletConfig), err
}

// Update takes the representation of a kubeletConfig and updates it. Returns the server's representation of the kubeletConfig, and an error, if there is any.
func (c *FakeKubeletConfigs) Update(ctx context.Context, kubeletConfig *v1.KubeletConfig, opts metav1.UpdateOptions) (result *v1.KubeletConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(kubeletconfigsResource, kubeletConfig), &v1.KubeletConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.KubeletConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeKubeletConfigs) UpdateStatus(ctx context.Context, kubeletConfig *v1.KubeletConfig, opts metav1.UpdateOptions) (*v1.KubeletConfig, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(kubeletconfigsResource, "status", kubeletConfig), &v1.KubeletConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.KubeletConfig), err
}

// Delete takes name of the kubeletConfig and deletes it. Returns an error if one occurs.
func (c *FakeKubeletConfigs) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(kubeletconfigsResource, name, opts), &v1.KubeletConfig{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKubeletConfigs) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(kubeletconfigsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1.KubeletConfigList{})
	return err
}

// Patch applies the patch and returns the patched kubeletConfig.
func (c *FakeKubeletConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.KubeletConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(kubeletconfigsResource, name, pt, data, subresources...), &v1.KubeletConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.KubeletConfig), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied kubeletConfig.
func (c *FakeKubeletConfigs) Apply(ctx context.Context, kubeletConfig *machineconfigurationv1.KubeletConfigApplyConfiguration, opts metav1.ApplyOptions) (result *v1.KubeletConfig, err error) {
	if kubeletConfig == nil {
		return nil, fmt.Errorf("kubeletConfig provided to Apply must not be nil")
	}
	data, err := json.Marshal(kubeletConfig)
	if err != nil {
		return nil, err
	}
	name := kubeletConfig.Name
	if name == nil {
		return nil, fmt.Errorf("kubeletConfig.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(kubeletconfigsResource, *name, types.ApplyPatchType, data), &v1.KubeletConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.KubeletConfig), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeKubeletConfigs) ApplyStatus(ctx context.Context, kubeletConfig *machineconfigurationv1.KubeletConfigApplyConfiguration, opts metav1.ApplyOptions) (result *v1.KubeletConfig, err error) {
	if kubeletConfig == nil {
		return nil, fmt.Errorf("kubeletConfig provided to Apply must not be nil")
	}
	data, err := json.Marshal(kubeletConfig)
	if err != nil {
		return nil, err
	}
	name := kubeletConfig.Name
	if name == nil {
		return nil, fmt.Errorf("kubeletConfig.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(kubeletconfigsResource, *name, types.ApplyPatchType, data, "status"), &v1.KubeletConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.KubeletConfig), err
}
