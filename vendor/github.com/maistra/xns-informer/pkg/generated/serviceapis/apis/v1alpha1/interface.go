/*
Copyright Red Hat, Inc.

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

// Code generated by xns-informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/maistra/xns-informer/pkg/generated/serviceapis/internalinterfaces"
	informers "github.com/maistra/xns-informer/pkg/informers"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// BackendPolicies returns a BackendPolicyInformer.
	BackendPolicies() BackendPolicyInformer
	// Gateways returns a GatewayInformer.
	Gateways() GatewayInformer
	// GatewayClasses returns a GatewayClassInformer.
	GatewayClasses() GatewayClassInformer
	// HTTPRoutes returns a HTTPRouteInformer.
	HTTPRoutes() HTTPRouteInformer
	// TCPRoutes returns a TCPRouteInformer.
	TCPRoutes() TCPRouteInformer
	// TLSRoutes returns a TLSRouteInformer.
	TLSRoutes() TLSRouteInformer
	// UDPRoutes returns a UDPRouteInformer.
	UDPRoutes() UDPRouteInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespaces       informers.NamespaceSet
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespaces informers.NamespaceSet, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespaces: namespaces, tweakListOptions: tweakListOptions}
}

// BackendPolicies returns a BackendPolicyInformer.
func (v *version) BackendPolicies() BackendPolicyInformer {
	return &backendPolicyInformer{factory: v.factory, namespaces: v.namespaces, tweakListOptions: v.tweakListOptions}
}

// Gateways returns a GatewayInformer.
func (v *version) Gateways() GatewayInformer {
	return &gatewayInformer{factory: v.factory, namespaces: v.namespaces, tweakListOptions: v.tweakListOptions}
}

// GatewayClasses returns a GatewayClassInformer.
func (v *version) GatewayClasses() GatewayClassInformer {
	return &gatewayClassInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// HTTPRoutes returns a HTTPRouteInformer.
func (v *version) HTTPRoutes() HTTPRouteInformer {
	return &hTTPRouteInformer{factory: v.factory, namespaces: v.namespaces, tweakListOptions: v.tweakListOptions}
}

// TCPRoutes returns a TCPRouteInformer.
func (v *version) TCPRoutes() TCPRouteInformer {
	return &tCPRouteInformer{factory: v.factory, namespaces: v.namespaces, tweakListOptions: v.tweakListOptions}
}

// TLSRoutes returns a TLSRouteInformer.
func (v *version) TLSRoutes() TLSRouteInformer {
	return &tLSRouteInformer{factory: v.factory, namespaces: v.namespaces, tweakListOptions: v.tweakListOptions}
}

// UDPRoutes returns a UDPRouteInformer.
func (v *version) UDPRoutes() UDPRouteInformer {
	return &uDPRouteInformer{factory: v.factory, namespaces: v.namespaces, tweakListOptions: v.tweakListOptions}
}