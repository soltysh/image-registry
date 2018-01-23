// This file was automatically generated by informer-gen

package v1

import (
	internalinterfaces "github.com/openshift/client-go/build/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Builds returns a BuildInformer.
	Builds() BuildInformer
	// BuildConfigs returns a BuildConfigInformer.
	BuildConfigs() BuildConfigInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// Builds returns a BuildInformer.
func (v *version) Builds() BuildInformer {
	return &buildInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// BuildConfigs returns a BuildConfigInformer.
func (v *version) BuildConfigs() BuildConfigInformer {
	return &buildConfigInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
