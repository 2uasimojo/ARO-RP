// Code generated by client-gen. DO NOT EDIT.

package versioned

import (
	"fmt"

	hivev1 "github.com/openshift/hive/pkg/client/clientset/versioned/typed/hive/v1"
	hiveinternalv1alpha1 "github.com/openshift/hive/pkg/client/clientset/versioned/typed/hiveinternal/v1alpha1"
	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	HiveV1() hivev1.HiveV1Interface
	HiveinternalV1alpha1() hiveinternalv1alpha1.HiveinternalV1alpha1Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	hiveV1               *hivev1.HiveV1Client
	hiveinternalV1alpha1 *hiveinternalv1alpha1.HiveinternalV1alpha1Client
}

// HiveV1 retrieves the HiveV1Client
func (c *Clientset) HiveV1() hivev1.HiveV1Interface {
	return c.hiveV1
}

// HiveinternalV1alpha1 retrieves the HiveinternalV1alpha1Client
func (c *Clientset) HiveinternalV1alpha1() hiveinternalv1alpha1.HiveinternalV1alpha1Interface {
	return c.hiveinternalV1alpha1
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfig will generate a rate-limiter in configShallowCopy.
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		if configShallowCopy.Burst <= 0 {
			return nil, fmt.Errorf("burst is required to be greater than 0 when RateLimiter is not set and QPS is set to greater than 0")
		}
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error
	cs.hiveV1, err = hivev1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.hiveinternalV1alpha1, err = hiveinternalv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	var cs Clientset
	cs.hiveV1 = hivev1.NewForConfigOrDie(c)
	cs.hiveinternalV1alpha1 = hiveinternalv1alpha1.NewForConfigOrDie(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.hiveV1 = hivev1.New(c)
	cs.hiveinternalV1alpha1 = hiveinternalv1alpha1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
