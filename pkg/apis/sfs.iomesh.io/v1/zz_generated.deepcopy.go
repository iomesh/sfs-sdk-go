//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023 The IOMesh Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessConfig) DeepCopyInto(out *AccessConfig) {
	*out = *in
	if in.SpecialAccessRules != nil {
		in, out := &in.SpecialAccessRules, &out.SpecialAccessRules
		*out = make([]AccessRule, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessConfig.
func (in *AccessConfig) DeepCopy() *AccessConfig {
	if in == nil {
		return nil
	}
	out := new(AccessConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessRule) DeepCopyInto(out *AccessRule) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessRule.
func (in *AccessRule) DeepCopy() *AccessRule {
	if in == nil {
		return nil
	}
	out := new(AccessRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnonymousUserInfo) DeepCopyInto(out *AnonymousUserInfo) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnonymousUserInfo.
func (in *AnonymousUserInfo) DeepCopy() *AnonymousUserInfo {
	if in == nil {
		return nil
	}
	out := new(AnonymousUserInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudProvider) DeepCopyInto(out *CloudProvider) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudProvider.
func (in *CloudProvider) DeepCopy() *CloudProvider {
	if in == nil {
		return nil
	}
	out := new(CloudProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudProvider) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudProviderList) DeepCopyInto(out *CloudProviderList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CloudProvider, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudProviderList.
func (in *CloudProviderList) DeepCopy() *CloudProviderList {
	if in == nil {
		return nil
	}
	out := new(CloudProviderList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudProviderList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudProviderSpec) DeepCopyInto(out *CloudProviderSpec) {
	*out = *in
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudProviderSpec.
func (in *CloudProviderSpec) DeepCopy() *CloudProviderSpec {
	if in == nil {
		return nil
	}
	out := new(CloudProviderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cluster) DeepCopyInto(out *Cluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cluster.
func (in *Cluster) DeepCopy() *Cluster {
	if in == nil {
		return nil
	}
	out := new(Cluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Cluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterList) DeepCopyInto(out *ClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Cluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterList.
func (in *ClusterList) DeepCopy() *ClusterList {
	if in == nil {
		return nil
	}
	out := new(ClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpec) DeepCopyInto(out *ClusterSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpec.
func (in *ClusterSpec) DeepCopy() *ClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IOMetric) DeepCopyInto(out *IOMetric) {
	*out = *in
	out.Read = in.Read
	out.Write = in.Write
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IOMetric.
func (in *IOMetric) DeepCopy() *IOMetric {
	if in == nil {
		return nil
	}
	out := new(IOMetric)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Metric) DeepCopyInto(out *Metric) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Metric.
func (in *Metric) DeepCopy() *Metric {
	if in == nil {
		return nil
	}
	out := new(Metric)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Metric) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricList) DeepCopyInto(out *MetricList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Metric, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricList.
func (in *MetricList) DeepCopy() *MetricList {
	if in == nil {
		return nil
	}
	out := new(MetricList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MetricList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricSpec) DeepCopyInto(out *MetricSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricSpec.
func (in *MetricSpec) DeepCopy() *MetricSpec {
	if in == nil {
		return nil
	}
	out := new(MetricSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricStatus) DeepCopyInto(out *MetricStatus) {
	*out = *in
	out.ClusterIOMetric = in.ClusterIOMetric
	out.ClusterStatMetric = in.ClusterStatMetric
	if in.NamespaceIOMetrics != nil {
		in, out := &in.NamespaceIOMetrics, &out.NamespaceIOMetrics
		*out = make(map[string]IOMetric, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.NamespaceStatMetrics != nil {
		in, out := &in.NamespaceStatMetrics, &out.NamespaceStatMetrics
		*out = make(map[string]StatMetric, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricStatus.
func (in *MetricStatus) DeepCopy() *MetricStatus {
	if in == nil {
		return nil
	}
	out := new(MetricStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NFSExportConfig) DeepCopyInto(out *NFSExportConfig) {
	*out = *in
	out.AnonymousUserInfo = in.AnonymousUserInfo
	in.AccessConfig.DeepCopyInto(&out.AccessConfig)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NFSExportConfig.
func (in *NFSExportConfig) DeepCopy() *NFSExportConfig {
	if in == nil {
		return nil
	}
	out := new(NFSExportConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Namespace) DeepCopyInto(out *Namespace) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Namespace.
func (in *Namespace) DeepCopy() *Namespace {
	if in == nil {
		return nil
	}
	out := new(Namespace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Namespace) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceList) DeepCopyInto(out *NamespaceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Namespace, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceList.
func (in *NamespaceList) DeepCopy() *NamespaceList {
	if in == nil {
		return nil
	}
	out := new(NamespaceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NamespaceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceSpec) DeepCopyInto(out *NamespaceSpec) {
	*out = *in
	in.StoragePolicy.DeepCopyInto(&out.StoragePolicy)
	out.Protocols = in.Protocols
	if in.NFSExportConfig != nil {
		in, out := &in.NFSExportConfig, &out.NFSExportConfig
		*out = new(NFSExportConfig)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceSpec.
func (in *NamespaceSpec) DeepCopy() *NamespaceSpec {
	if in == nil {
		return nil
	}
	out := new(NamespaceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceStatus) DeepCopyInto(out *NamespaceStatus) {
	*out = *in
	if in.DshardIDS != nil {
		in, out := &in.DshardIDS, &out.DshardIDS
		*out = make([]int64, len(*in))
		copy(*out, *in)
	}
	if in.MshardIDS != nil {
		in, out := &in.MshardIDS, &out.MshardIDS
		*out = make([]int64, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceStatus.
func (in *NamespaceStatus) DeepCopy() *NamespaceStatus {
	if in == nil {
		return nil
	}
	out := new(NamespaceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Node) DeepCopyInto(out *Node) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Node.
func (in *Node) DeepCopy() *Node {
	if in == nil {
		return nil
	}
	out := new(Node)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Node) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeHealth) DeepCopyInto(out *NodeHealth) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeHealth.
func (in *NodeHealth) DeepCopy() *NodeHealth {
	if in == nil {
		return nil
	}
	out := new(NodeHealth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeHealth) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeHealthList) DeepCopyInto(out *NodeHealthList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NodeHealth, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeHealthList.
func (in *NodeHealthList) DeepCopy() *NodeHealthList {
	if in == nil {
		return nil
	}
	out := new(NodeHealthList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeHealthList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeHealthSpec) DeepCopyInto(out *NodeHealthSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeHealthSpec.
func (in *NodeHealthSpec) DeepCopy() *NodeHealthSpec {
	if in == nil {
		return nil
	}
	out := new(NodeHealthSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeHealthStatus) DeepCopyInto(out *NodeHealthStatus) {
	*out = *in
	if in.DataService != nil {
		in, out := &in.DataService, &out.DataService
		*out = new(bool)
		**out = **in
	}
	if in.MetaService != nil {
		in, out := &in.MetaService, &out.MetaService
		*out = new(bool)
		**out = **in
	}
	if in.RenewTime != nil {
		in, out := &in.RenewTime, &out.RenewTime
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeHealthStatus.
func (in *NodeHealthStatus) DeepCopy() *NodeHealthStatus {
	if in == nil {
		return nil
	}
	out := new(NodeHealthStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeList) DeepCopyInto(out *NodeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Node, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeList.
func (in *NodeList) DeepCopy() *NodeList {
	if in == nil {
		return nil
	}
	out := new(NodeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeSpec) DeepCopyInto(out *NodeSpec) {
	*out = *in
	if in.DataServer != nil {
		in, out := &in.DataServer, &out.DataServer
		*out = new(string)
		**out = **in
	}
	if in.DataShards != nil {
		in, out := &in.DataShards, &out.DataShards
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.MetaServer != nil {
		in, out := &in.MetaServer, &out.MetaServer
		*out = new(string)
		**out = **in
	}
	if in.MetaShards != nil {
		in, out := &in.MetaShards, &out.MetaShards
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.NodeUUID != nil {
		in, out := &in.NodeUUID, &out.NodeUUID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeSpec.
func (in *NodeSpec) DeepCopy() *NodeSpec {
	if in == nil {
		return nil
	}
	out := new(NodeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeStatus) DeepCopyInto(out *NodeStatus) {
	*out = *in
	if in.DataShards != nil {
		in, out := &in.DataShards, &out.DataShards
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.DataStatus != nil {
		in, out := &in.DataStatus, &out.DataStatus
		*out = new(bool)
		**out = **in
	}
	if in.LastActive != nil {
		in, out := &in.LastActive, &out.LastActive
		*out = new(string)
		**out = **in
	}
	if in.MetaShards != nil {
		in, out := &in.MetaShards, &out.MetaShards
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.MetaStatus != nil {
		in, out := &in.MetaStatus, &out.MetaStatus
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeStatus.
func (in *NodeStatus) DeepCopy() *NodeStatus {
	if in == nil {
		return nil
	}
	out := new(NodeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NsReference) DeepCopyInto(out *NsReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NsReference.
func (in *NsReference) DeepCopy() *NsReference {
	if in == nil {
		return nil
	}
	out := new(NsReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerfMetric) DeepCopyInto(out *PerfMetric) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerfMetric.
func (in *PerfMetric) DeepCopy() *PerfMetric {
	if in == nil {
		return nil
	}
	out := new(PerfMetric)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Protocols) DeepCopyInto(out *Protocols) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Protocols.
func (in *Protocols) DeepCopy() *Protocols {
	if in == nil {
		return nil
	}
	out := new(Protocols)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecoveryDatabase) DeepCopyInto(out *RecoveryDatabase) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecoveryDatabase.
func (in *RecoveryDatabase) DeepCopy() *RecoveryDatabase {
	if in == nil {
		return nil
	}
	out := new(RecoveryDatabase)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RecoveryDatabase) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecoveryDatabaseList) DeepCopyInto(out *RecoveryDatabaseList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RecoveryDatabase, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecoveryDatabaseList.
func (in *RecoveryDatabaseList) DeepCopy() *RecoveryDatabaseList {
	if in == nil {
		return nil
	}
	out := new(RecoveryDatabaseList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RecoveryDatabaseList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecoveryDatabaseSpec) DeepCopyInto(out *RecoveryDatabaseSpec) {
	*out = *in
	if in.Clids != nil {
		in, out := &in.Clids, &out.Clids
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ReclaimingClids != nil {
		in, out := &in.ReclaimingClids, &out.ReclaimingClids
		*out = make([]int64, len(*in))
		copy(*out, *in)
	}
	if in.Rfhs != nil {
		in, out := &in.Rfhs, &out.Rfhs
		*out = make(map[string][]string, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecoveryDatabaseSpec.
func (in *RecoveryDatabaseSpec) DeepCopy() *RecoveryDatabaseSpec {
	if in == nil {
		return nil
	}
	out := new(RecoveryDatabaseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Route) DeepCopyInto(out *Route) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Route.
func (in *Route) DeepCopy() *Route {
	if in == nil {
		return nil
	}
	out := new(Route)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Route) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteInfo) DeepCopyInto(out *RouteInfo) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteInfo.
func (in *RouteInfo) DeepCopy() *RouteInfo {
	if in == nil {
		return nil
	}
	out := new(RouteInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteList) DeepCopyInto(out *RouteList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Route, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteList.
func (in *RouteList) DeepCopy() *RouteList {
	if in == nil {
		return nil
	}
	out := new(RouteList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RouteList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteSpec) DeepCopyInto(out *RouteSpec) {
	*out = *in
	if in.Routes != nil {
		in, out := &in.Routes, &out.Routes
		*out = make([]RouteInfo, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteSpec.
func (in *RouteSpec) DeepCopy() *RouteSpec {
	if in == nil {
		return nil
	}
	out := new(RouteSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Session) DeepCopyInto(out *Session) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Session.
func (in *Session) DeepCopy() *Session {
	if in == nil {
		return nil
	}
	out := new(Session)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Session) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SessionList) DeepCopyInto(out *SessionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Session, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SessionList.
func (in *SessionList) DeepCopy() *SessionList {
	if in == nil {
		return nil
	}
	out := new(SessionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SessionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SessionSpec) DeepCopyInto(out *SessionSpec) {
	*out = *in
	if in.Vips != nil {
		in, out := &in.Vips, &out.Vips
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.MACAddr != nil {
		in, out := &in.MACAddr, &out.MACAddr
		*out = make([]int64, len(*in))
		copy(*out, *in)
	}
	if in.Netmask != nil {
		in, out := &in.Netmask, &out.Netmask
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SessionSpec.
func (in *SessionSpec) DeepCopy() *SessionSpec {
	if in == nil {
		return nil
	}
	out := new(SessionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SessionStatus) DeepCopyInto(out *SessionStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SessionStatus.
func (in *SessionStatus) DeepCopy() *SessionStatus {
	if in == nil {
		return nil
	}
	out := new(SessionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Shard) DeepCopyInto(out *Shard) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Shard.
func (in *Shard) DeepCopy() *Shard {
	if in == nil {
		return nil
	}
	out := new(Shard)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Shard) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShardList) DeepCopyInto(out *ShardList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Shard, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShardList.
func (in *ShardList) DeepCopy() *ShardList {
	if in == nil {
		return nil
	}
	out := new(ShardList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ShardList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShardSpec) DeepCopyInto(out *ShardSpec) {
	*out = *in
	out.NsReference = in.NsReference
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShardSpec.
func (in *ShardSpec) DeepCopy() *ShardSpec {
	if in == nil {
		return nil
	}
	out := new(ShardSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShardStatus) DeepCopyInto(out *ShardStatus) {
	*out = *in
	if in.VolumeName != nil {
		in, out := &in.VolumeName, &out.VolumeName
		*out = new(string)
		**out = **in
	}
	if in.HAMembers != nil {
		in, out := &in.HAMembers, &out.HAMembers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NodeName != nil {
		in, out := &in.NodeName, &out.NodeName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShardStatus.
func (in *ShardStatus) DeepCopy() *ShardStatus {
	if in == nil {
		return nil
	}
	out := new(ShardStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatMetric) DeepCopyInto(out *StatMetric) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatMetric.
func (in *StatMetric) DeepCopy() *StatMetric {
	if in == nil {
		return nil
	}
	out := new(StatMetric)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoragePolicy) DeepCopyInto(out *StoragePolicy) {
	*out = *in
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoragePolicy.
func (in *StoragePolicy) DeepCopy() *StoragePolicy {
	if in == nil {
		return nil
	}
	out := new(StoragePolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Vip) DeepCopyInto(out *Vip) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Vip.
func (in *Vip) DeepCopy() *Vip {
	if in == nil {
		return nil
	}
	out := new(Vip)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Vip) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VipList) DeepCopyInto(out *VipList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Vip, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VipList.
func (in *VipList) DeepCopy() *VipList {
	if in == nil {
		return nil
	}
	out := new(VipList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VipList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VipSpec) DeepCopyInto(out *VipSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VipSpec.
func (in *VipSpec) DeepCopy() *VipSpec {
	if in == nil {
		return nil
	}
	out := new(VipSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VipStatus) DeepCopyInto(out *VipStatus) {
	*out = *in
	if in.SessionName != nil {
		in, out := &in.SessionName, &out.SessionName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VipStatus.
func (in *VipStatus) DeepCopy() *VipStatus {
	if in == nil {
		return nil
	}
	out := new(VipStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Volume) DeepCopyInto(out *Volume) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Volume.
func (in *Volume) DeepCopy() *Volume {
	if in == nil {
		return nil
	}
	out := new(Volume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Volume) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeList) DeepCopyInto(out *VolumeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeList.
func (in *VolumeList) DeepCopy() *VolumeList {
	if in == nil {
		return nil
	}
	out := new(VolumeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VolumeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeSpec) DeepCopyInto(out *VolumeSpec) {
	*out = *in
	out.NsReference = in.NsReference
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeSpec.
func (in *VolumeSpec) DeepCopy() *VolumeSpec {
	if in == nil {
		return nil
	}
	out := new(VolumeSpec)
	in.DeepCopyInto(out)
	return out
}
