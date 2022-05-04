/*
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
// Code generated by informer-gen. DO NOT EDIT.

package v1alpha2

import (
	internalinterfaces "github.com/openservicemesh/osm/pkg/gen/client/config/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// MeshConfigs returns a MeshConfigInformer.
	MeshConfigs() MeshConfigInformer
	// MeshRootCertificates returns a MeshRootCertificateInformer.
	MeshRootCertificates() MeshRootCertificateInformer
	// MultiClusterServices returns a MultiClusterServiceInformer.
	MultiClusterServices() MultiClusterServiceInformer
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

// MeshConfigs returns a MeshConfigInformer.
func (v *version) MeshConfigs() MeshConfigInformer {
	return &meshConfigInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// MeshRootCertificates returns a MeshRootCertificateInformer.
func (v *version) MeshRootCertificates() MeshRootCertificateInformer {
	return &meshRootCertificateInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// MultiClusterServices returns a MultiClusterServiceInformer.
func (v *version) MultiClusterServices() MultiClusterServiceInformer {
	return &multiClusterServiceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
