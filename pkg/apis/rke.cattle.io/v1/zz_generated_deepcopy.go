// +build !ignore_autogenerated

/*
Copyright 2021 Rancher Labs, Inc.

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

// Code generated by main. DO NOT EDIT.

package v1

import (
	genericcondition "github.com/rancher/wrangler/pkg/genericcondition"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterUpgradeStrategy) DeepCopyInto(out *ClusterUpgradeStrategy) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterUpgradeStrategy.
func (in *ClusterUpgradeStrategy) DeepCopy() *ClusterUpgradeStrategy {
	if in == nil {
		return nil
	}
	out := new(ClusterUpgradeStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Endpoint) DeepCopyInto(out *Endpoint) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Endpoint.
func (in *Endpoint) DeepCopy() *Endpoint {
	if in == nil {
		return nil
	}
	out := new(Endpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericMap.
func (in *GenericMap) DeepCopy() *GenericMap {
	if in == nil {
		return nil
	}
	out := new(GenericMap)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RKEBootstrap) DeepCopyInto(out *RKEBootstrap) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RKEBootstrap.
func (in *RKEBootstrap) DeepCopy() *RKEBootstrap {
	if in == nil {
		return nil
	}
	out := new(RKEBootstrap)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RKEBootstrap) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RKEBootstrapList) DeepCopyInto(out *RKEBootstrapList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RKEBootstrap, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RKEBootstrapList.
func (in *RKEBootstrapList) DeepCopy() *RKEBootstrapList {
	if in == nil {
		return nil
	}
	out := new(RKEBootstrapList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RKEBootstrapList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RKEBootstrapSpec) DeepCopyInto(out *RKEBootstrapSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RKEBootstrapSpec.
func (in *RKEBootstrapSpec) DeepCopy() *RKEBootstrapSpec {
	if in == nil {
		return nil
	}
	out := new(RKEBootstrapSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RKEBootstrapStatus) DeepCopyInto(out *RKEBootstrapStatus) {
	*out = *in
	if in.DataSecretName != nil {
		in, out := &in.DataSecretName, &out.DataSecretName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RKEBootstrapStatus.
func (in *RKEBootstrapStatus) DeepCopy() *RKEBootstrapStatus {
	if in == nil {
		return nil
	}
	out := new(RKEBootstrapStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RKEBootstrapTemplate) DeepCopyInto(out *RKEBootstrapTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RKEBootstrapTemplate.
func (in *RKEBootstrapTemplate) DeepCopy() *RKEBootstrapTemplate {
	if in == nil {
		return nil
	}
	out := new(RKEBootstrapTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RKEBootstrapTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RKEBootstrapTemplateList) DeepCopyInto(out *RKEBootstrapTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RKEBootstrapTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RKEBootstrapTemplateList.
func (in *RKEBootstrapTemplateList) DeepCopy() *RKEBootstrapTemplateList {
	if in == nil {
		return nil
	}
	out := new(RKEBootstrapTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RKEBootstrapTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RKEBootstrapTemplateSpec) DeepCopyInto(out *RKEBootstrapTemplateSpec) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RKEBootstrapTemplateSpec.
func (in *RKEBootstrapTemplateSpec) DeepCopy() *RKEBootstrapTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(RKEBootstrapTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RKECluster) DeepCopyInto(out *RKECluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RKECluster.
func (in *RKECluster) DeepCopy() *RKECluster {
	if in == nil {
		return nil
	}
	out := new(RKECluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RKECluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RKEClusterList) DeepCopyInto(out *RKEClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RKECluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RKEClusterList.
func (in *RKEClusterList) DeepCopy() *RKEClusterList {
	if in == nil {
		return nil
	}
	out := new(RKEClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RKEClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RKEClusterSpec) DeepCopyInto(out *RKEClusterSpec) {
	*out = *in
	if in.ControlPlaneEndpoint != nil {
		in, out := &in.ControlPlaneEndpoint, &out.ControlPlaneEndpoint
		*out = new(Endpoint)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RKEClusterSpec.
func (in *RKEClusterSpec) DeepCopy() *RKEClusterSpec {
	if in == nil {
		return nil
	}
	out := new(RKEClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RKEClusterSpecCommon) DeepCopyInto(out *RKEClusterSpecCommon) {
	*out = *in
	out.UpgradeStrategy = in.UpgradeStrategy
	in.CNIValues.DeepCopyInto(&out.CNIValues)
	in.ControlPlaneConfig.DeepCopyInto(&out.ControlPlaneConfig)
	if in.NodeConfig != nil {
		in, out := &in.NodeConfig, &out.NodeConfig
		*out = make([]RKESystemConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RKEClusterSpecCommon.
func (in *RKEClusterSpecCommon) DeepCopy() *RKEClusterSpecCommon {
	if in == nil {
		return nil
	}
	out := new(RKEClusterSpecCommon)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RKEClusterStatus) DeepCopyInto(out *RKEClusterStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]genericcondition.GenericCondition, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RKEClusterStatus.
func (in *RKEClusterStatus) DeepCopy() *RKEClusterStatus {
	if in == nil {
		return nil
	}
	out := new(RKEClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RKECommonNodeConfig) DeepCopyInto(out *RKECommonNodeConfig) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Taints != nil {
		in, out := &in.Taints, &out.Taints
		*out = make([]corev1.Taint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RKECommonNodeConfig.
func (in *RKECommonNodeConfig) DeepCopy() *RKECommonNodeConfig {
	if in == nil {
		return nil
	}
	out := new(RKECommonNodeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RKEControlPlane) DeepCopyInto(out *RKEControlPlane) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RKEControlPlane.
func (in *RKEControlPlane) DeepCopy() *RKEControlPlane {
	if in == nil {
		return nil
	}
	out := new(RKEControlPlane)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RKEControlPlane) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RKEControlPlaneList) DeepCopyInto(out *RKEControlPlaneList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RKEControlPlane, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RKEControlPlaneList.
func (in *RKEControlPlaneList) DeepCopy() *RKEControlPlaneList {
	if in == nil {
		return nil
	}
	out := new(RKEControlPlaneList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RKEControlPlaneList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RKEControlPlaneSpec) DeepCopyInto(out *RKEControlPlaneSpec) {
	*out = *in
	in.RKEClusterSpecCommon.DeepCopyInto(&out.RKEClusterSpecCommon)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RKEControlPlaneSpec.
func (in *RKEControlPlaneSpec) DeepCopy() *RKEControlPlaneSpec {
	if in == nil {
		return nil
	}
	out := new(RKEControlPlaneSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RKEControlPlaneStatus) DeepCopyInto(out *RKEControlPlaneStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]genericcondition.GenericCondition, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RKEControlPlaneStatus.
func (in *RKEControlPlaneStatus) DeepCopy() *RKEControlPlaneStatus {
	if in == nil {
		return nil
	}
	out := new(RKEControlPlaneStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RKEMachineStatus) DeepCopyInto(out *RKEMachineStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RKEMachineStatus.
func (in *RKEMachineStatus) DeepCopy() *RKEMachineStatus {
	if in == nil {
		return nil
	}
	out := new(RKEMachineStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RKESystemConfig) DeepCopyInto(out *RKESystemConfig) {
	*out = *in
	if in.MachineLabelSelector != nil {
		in, out := &in.MachineLabelSelector, &out.MachineLabelSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	in.Config.DeepCopyInto(&out.Config)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RKESystemConfig.
func (in *RKESystemConfig) DeepCopy() *RKESystemConfig {
	if in == nil {
		return nil
	}
	out := new(RKESystemConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UnmanagedMachine) DeepCopyInto(out *UnmanagedMachine) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UnmanagedMachine.
func (in *UnmanagedMachine) DeepCopy() *UnmanagedMachine {
	if in == nil {
		return nil
	}
	out := new(UnmanagedMachine)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UnmanagedMachine) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UnmanagedMachineList) DeepCopyInto(out *UnmanagedMachineList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]UnmanagedMachine, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UnmanagedMachineList.
func (in *UnmanagedMachineList) DeepCopy() *UnmanagedMachineList {
	if in == nil {
		return nil
	}
	out := new(UnmanagedMachineList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UnmanagedMachineList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UnmanagedMachineSpec) DeepCopyInto(out *UnmanagedMachineSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UnmanagedMachineSpec.
func (in *UnmanagedMachineSpec) DeepCopy() *UnmanagedMachineSpec {
	if in == nil {
		return nil
	}
	out := new(UnmanagedMachineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UnmanagedMachineStatus) DeepCopyInto(out *UnmanagedMachineStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UnmanagedMachineStatus.
func (in *UnmanagedMachineStatus) DeepCopy() *UnmanagedMachineStatus {
	if in == nil {
		return nil
	}
	out := new(UnmanagedMachineStatus)
	in.DeepCopyInto(out)
	return out
}
