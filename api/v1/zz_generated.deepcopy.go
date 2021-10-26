//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CUDAValidatorSpec) DeepCopyInto(out *CUDAValidatorSpec) {
	*out = *in
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]corev1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CUDAValidatorSpec.
func (in *CUDAValidatorSpec) DeepCopy() *CUDAValidatorSpec {
	if in == nil {
		return nil
	}
	out := new(CUDAValidatorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterPolicy) DeepCopyInto(out *ClusterPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterPolicy.
func (in *ClusterPolicy) DeepCopy() *ClusterPolicy {
	if in == nil {
		return nil
	}
	out := new(ClusterPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterPolicyList) DeepCopyInto(out *ClusterPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterPolicyList.
func (in *ClusterPolicyList) DeepCopy() *ClusterPolicyList {
	if in == nil {
		return nil
	}
	out := new(ClusterPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterPolicySpec) DeepCopyInto(out *ClusterPolicySpec) {
	*out = *in
	in.Operator.DeepCopyInto(&out.Operator)
	in.Daemonsets.DeepCopyInto(&out.Daemonsets)
	in.Driver.DeepCopyInto(&out.Driver)
	in.Toolkit.DeepCopyInto(&out.Toolkit)
	in.DevicePlugin.DeepCopyInto(&out.DevicePlugin)
	in.DCGMExporter.DeepCopyInto(&out.DCGMExporter)
	in.DCGM.DeepCopyInto(&out.DCGM)
	in.NodeStatusExporter.DeepCopyInto(&out.NodeStatusExporter)
	in.GPUFeatureDiscovery.DeepCopyInto(&out.GPUFeatureDiscovery)
	out.MIG = in.MIG
	in.MIGManager.DeepCopyInto(&out.MIGManager)
	in.PSP.DeepCopyInto(&out.PSP)
	in.Validator.DeepCopyInto(&out.Validator)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterPolicySpec.
func (in *ClusterPolicySpec) DeepCopy() *ClusterPolicySpec {
	if in == nil {
		return nil
	}
	out := new(ClusterPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterPolicyStatus) DeepCopyInto(out *ClusterPolicyStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterPolicyStatus.
func (in *ClusterPolicyStatus) DeepCopy() *ClusterPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DCGMExporterMetricsConfig) DeepCopyInto(out *DCGMExporterMetricsConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DCGMExporterMetricsConfig.
func (in *DCGMExporterMetricsConfig) DeepCopy() *DCGMExporterMetricsConfig {
	if in == nil {
		return nil
	}
	out := new(DCGMExporterMetricsConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DCGMExporterSpec) DeepCopyInto(out *DCGMExporterSpec) {
	*out = *in
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(corev1.SecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]corev1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MetricsConfig != nil {
		in, out := &in.MetricsConfig, &out.MetricsConfig
		*out = new(DCGMExporterMetricsConfig)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DCGMExporterSpec.
func (in *DCGMExporterSpec) DeepCopy() *DCGMExporterSpec {
	if in == nil {
		return nil
	}
	out := new(DCGMExporterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DCGMSpec) DeepCopyInto(out *DCGMSpec) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(corev1.SecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]corev1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DCGMSpec.
func (in *DCGMSpec) DeepCopy() *DCGMSpec {
	if in == nil {
		return nil
	}
	out := new(DCGMSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DaemonsetsSpec) DeepCopyInto(out *DaemonsetsSpec) {
	*out = *in
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]corev1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DaemonsetsSpec.
func (in *DaemonsetsSpec) DeepCopy() *DaemonsetsSpec {
	if in == nil {
		return nil
	}
	out := new(DaemonsetsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevicePluginSpec) DeepCopyInto(out *DevicePluginSpec) {
	*out = *in
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(corev1.SecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]corev1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevicePluginSpec.
func (in *DevicePluginSpec) DeepCopy() *DevicePluginSpec {
	if in == nil {
		return nil
	}
	out := new(DevicePluginSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DriverCertConfigSpec) DeepCopyInto(out *DriverCertConfigSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DriverCertConfigSpec.
func (in *DriverCertConfigSpec) DeepCopy() *DriverCertConfigSpec {
	if in == nil {
		return nil
	}
	out := new(DriverCertConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DriverLicensingConfigSpec) DeepCopyInto(out *DriverLicensingConfigSpec) {
	*out = *in
	if in.NLSEnabled != nil {
		in, out := &in.NLSEnabled, &out.NLSEnabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DriverLicensingConfigSpec.
func (in *DriverLicensingConfigSpec) DeepCopy() *DriverLicensingConfigSpec {
	if in == nil {
		return nil
	}
	out := new(DriverLicensingConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DriverManagerSpec) DeepCopyInto(out *DriverManagerSpec) {
	*out = *in
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]corev1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DriverManagerSpec.
func (in *DriverManagerSpec) DeepCopy() *DriverManagerSpec {
	if in == nil {
		return nil
	}
	out := new(DriverManagerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DriverRepoConfigSpec) DeepCopyInto(out *DriverRepoConfigSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DriverRepoConfigSpec.
func (in *DriverRepoConfigSpec) DeepCopy() *DriverRepoConfigSpec {
	if in == nil {
		return nil
	}
	out := new(DriverRepoConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DriverSpec) DeepCopyInto(out *DriverSpec) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.GPUDirectRDMA != nil {
		in, out := &in.GPUDirectRDMA, &out.GPUDirectRDMA
		*out = new(GPUDirectRDMASpec)
		(*in).DeepCopyInto(*out)
	}
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.Manager.DeepCopyInto(&out.Manager)
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(corev1.SecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]corev1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RepoConfig != nil {
		in, out := &in.RepoConfig, &out.RepoConfig
		*out = new(DriverRepoConfigSpec)
		**out = **in
	}
	if in.CertConfig != nil {
		in, out := &in.CertConfig, &out.CertConfig
		*out = new(DriverCertConfigSpec)
		**out = **in
	}
	if in.LicensingConfig != nil {
		in, out := &in.LicensingConfig, &out.LicensingConfig
		*out = new(DriverLicensingConfigSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.VirtualTopology != nil {
		in, out := &in.VirtualTopology, &out.VirtualTopology
		*out = new(VirtualTopologyConfigSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DriverSpec.
func (in *DriverSpec) DeepCopy() *DriverSpec {
	if in == nil {
		return nil
	}
	out := new(DriverSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DriverValidatorSpec) DeepCopyInto(out *DriverValidatorSpec) {
	*out = *in
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]corev1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DriverValidatorSpec.
func (in *DriverValidatorSpec) DeepCopy() *DriverValidatorSpec {
	if in == nil {
		return nil
	}
	out := new(DriverValidatorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GPUDirectRDMASpec) DeepCopyInto(out *GPUDirectRDMASpec) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.UseHostMOFED != nil {
		in, out := &in.UseHostMOFED, &out.UseHostMOFED
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GPUDirectRDMASpec.
func (in *GPUDirectRDMASpec) DeepCopy() *GPUDirectRDMASpec {
	if in == nil {
		return nil
	}
	out := new(GPUDirectRDMASpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GPUFeatureDiscoverySpec) DeepCopyInto(out *GPUFeatureDiscoverySpec) {
	*out = *in
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(corev1.SecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]corev1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GPUFeatureDiscoverySpec.
func (in *GPUFeatureDiscoverySpec) DeepCopy() *GPUFeatureDiscoverySpec {
	if in == nil {
		return nil
	}
	out := new(GPUFeatureDiscoverySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InitContainerSpec) DeepCopyInto(out *InitContainerSpec) {
	*out = *in
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InitContainerSpec.
func (in *InitContainerSpec) DeepCopy() *InitContainerSpec {
	if in == nil {
		return nil
	}
	out := new(InitContainerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MIGGPUClientsConfigSpec) DeepCopyInto(out *MIGGPUClientsConfigSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MIGGPUClientsConfigSpec.
func (in *MIGGPUClientsConfigSpec) DeepCopy() *MIGGPUClientsConfigSpec {
	if in == nil {
		return nil
	}
	out := new(MIGGPUClientsConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MIGManagerSpec) DeepCopyInto(out *MIGManagerSpec) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(corev1.SecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]corev1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = new(MIGPartedConfigSpec)
		**out = **in
	}
	if in.GPUClientsConfig != nil {
		in, out := &in.GPUClientsConfig, &out.GPUClientsConfig
		*out = new(MIGGPUClientsConfigSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MIGManagerSpec.
func (in *MIGManagerSpec) DeepCopy() *MIGManagerSpec {
	if in == nil {
		return nil
	}
	out := new(MIGManagerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MIGPartedConfigSpec) DeepCopyInto(out *MIGPartedConfigSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MIGPartedConfigSpec.
func (in *MIGPartedConfigSpec) DeepCopy() *MIGPartedConfigSpec {
	if in == nil {
		return nil
	}
	out := new(MIGPartedConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MIGSpec) DeepCopyInto(out *MIGSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MIGSpec.
func (in *MIGSpec) DeepCopy() *MIGSpec {
	if in == nil {
		return nil
	}
	out := new(MIGSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeStatusExporterSpec) DeepCopyInto(out *NodeStatusExporterSpec) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(corev1.SecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]corev1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeStatusExporterSpec.
func (in *NodeStatusExporterSpec) DeepCopy() *NodeStatusExporterSpec {
	if in == nil {
		return nil
	}
	out := new(NodeStatusExporterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperatorSpec) DeepCopyInto(out *OperatorSpec) {
	*out = *in
	in.InitContainer.DeepCopyInto(&out.InitContainer)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperatorSpec.
func (in *OperatorSpec) DeepCopy() *OperatorSpec {
	if in == nil {
		return nil
	}
	out := new(OperatorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PSPSpec) DeepCopyInto(out *PSPSpec) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PSPSpec.
func (in *PSPSpec) DeepCopy() *PSPSpec {
	if in == nil {
		return nil
	}
	out := new(PSPSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PluginValidatorSpec) DeepCopyInto(out *PluginValidatorSpec) {
	*out = *in
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]corev1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PluginValidatorSpec.
func (in *PluginValidatorSpec) DeepCopy() *PluginValidatorSpec {
	if in == nil {
		return nil
	}
	out := new(PluginValidatorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ToolkitSpec) DeepCopyInto(out *ToolkitSpec) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(corev1.SecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]corev1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ToolkitSpec.
func (in *ToolkitSpec) DeepCopy() *ToolkitSpec {
	if in == nil {
		return nil
	}
	out := new(ToolkitSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ToolkitValidatorSpec) DeepCopyInto(out *ToolkitValidatorSpec) {
	*out = *in
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]corev1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ToolkitValidatorSpec.
func (in *ToolkitValidatorSpec) DeepCopy() *ToolkitValidatorSpec {
	if in == nil {
		return nil
	}
	out := new(ToolkitValidatorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValidatorSpec) DeepCopyInto(out *ValidatorSpec) {
	*out = *in
	in.Plugin.DeepCopyInto(&out.Plugin)
	in.Toolkit.DeepCopyInto(&out.Toolkit)
	in.Driver.DeepCopyInto(&out.Driver)
	in.CUDA.DeepCopyInto(&out.CUDA)
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(corev1.SecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]corev1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValidatorSpec.
func (in *ValidatorSpec) DeepCopy() *ValidatorSpec {
	if in == nil {
		return nil
	}
	out := new(ValidatorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualTopologyConfigSpec) DeepCopyInto(out *VirtualTopologyConfigSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualTopologyConfigSpec.
func (in *VirtualTopologyConfigSpec) DeepCopy() *VirtualTopologyConfigSpec {
	if in == nil {
		return nil
	}
	out := new(VirtualTopologyConfigSpec)
	in.DeepCopyInto(out)
	return out
}
