//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstallPatchesObservation) DeepCopyInto(out *InstallPatchesObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstallPatchesObservation.
func (in *InstallPatchesObservation) DeepCopy() *InstallPatchesObservation {
	if in == nil {
		return nil
	}
	out := new(InstallPatchesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstallPatchesParameters) DeepCopyInto(out *InstallPatchesParameters) {
	*out = *in
	if in.Linux != nil {
		in, out := &in.Linux, &out.Linux
		*out = make([]LinuxParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Reboot != nil {
		in, out := &in.Reboot, &out.Reboot
		*out = new(string)
		**out = **in
	}
	if in.Windows != nil {
		in, out := &in.Windows, &out.Windows
		*out = make([]WindowsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstallPatchesParameters.
func (in *InstallPatchesParameters) DeepCopy() *InstallPatchesParameters {
	if in == nil {
		return nil
	}
	out := new(InstallPatchesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LinuxObservation) DeepCopyInto(out *LinuxObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinuxObservation.
func (in *LinuxObservation) DeepCopy() *LinuxObservation {
	if in == nil {
		return nil
	}
	out := new(LinuxObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LinuxParameters) DeepCopyInto(out *LinuxParameters) {
	*out = *in
	if in.ClassificationsToInclude != nil {
		in, out := &in.ClassificationsToInclude, &out.ClassificationsToInclude
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.PackageNamesMaskToExclude != nil {
		in, out := &in.PackageNamesMaskToExclude, &out.PackageNamesMaskToExclude
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.PackageNamesMaskToInclude != nil {
		in, out := &in.PackageNamesMaskToInclude, &out.PackageNamesMaskToInclude
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinuxParameters.
func (in *LinuxParameters) DeepCopy() *LinuxParameters {
	if in == nil {
		return nil
	}
	out := new(LinuxParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaintenanceAssignmentDedicatedHost) DeepCopyInto(out *MaintenanceAssignmentDedicatedHost) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaintenanceAssignmentDedicatedHost.
func (in *MaintenanceAssignmentDedicatedHost) DeepCopy() *MaintenanceAssignmentDedicatedHost {
	if in == nil {
		return nil
	}
	out := new(MaintenanceAssignmentDedicatedHost)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MaintenanceAssignmentDedicatedHost) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaintenanceAssignmentDedicatedHostList) DeepCopyInto(out *MaintenanceAssignmentDedicatedHostList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MaintenanceAssignmentDedicatedHost, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaintenanceAssignmentDedicatedHostList.
func (in *MaintenanceAssignmentDedicatedHostList) DeepCopy() *MaintenanceAssignmentDedicatedHostList {
	if in == nil {
		return nil
	}
	out := new(MaintenanceAssignmentDedicatedHostList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MaintenanceAssignmentDedicatedHostList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaintenanceAssignmentDedicatedHostObservation) DeepCopyInto(out *MaintenanceAssignmentDedicatedHostObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaintenanceAssignmentDedicatedHostObservation.
func (in *MaintenanceAssignmentDedicatedHostObservation) DeepCopy() *MaintenanceAssignmentDedicatedHostObservation {
	if in == nil {
		return nil
	}
	out := new(MaintenanceAssignmentDedicatedHostObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaintenanceAssignmentDedicatedHostParameters) DeepCopyInto(out *MaintenanceAssignmentDedicatedHostParameters) {
	*out = *in
	if in.DedicatedHostID != nil {
		in, out := &in.DedicatedHostID, &out.DedicatedHostID
		*out = new(string)
		**out = **in
	}
	if in.DedicatedHostIDRef != nil {
		in, out := &in.DedicatedHostIDRef, &out.DedicatedHostIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.DedicatedHostIDSelector != nil {
		in, out := &in.DedicatedHostIDSelector, &out.DedicatedHostIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.MaintenanceConfigurationID != nil {
		in, out := &in.MaintenanceConfigurationID, &out.MaintenanceConfigurationID
		*out = new(string)
		**out = **in
	}
	if in.MaintenanceConfigurationIDRef != nil {
		in, out := &in.MaintenanceConfigurationIDRef, &out.MaintenanceConfigurationIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.MaintenanceConfigurationIDSelector != nil {
		in, out := &in.MaintenanceConfigurationIDSelector, &out.MaintenanceConfigurationIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaintenanceAssignmentDedicatedHostParameters.
func (in *MaintenanceAssignmentDedicatedHostParameters) DeepCopy() *MaintenanceAssignmentDedicatedHostParameters {
	if in == nil {
		return nil
	}
	out := new(MaintenanceAssignmentDedicatedHostParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaintenanceAssignmentDedicatedHostSpec) DeepCopyInto(out *MaintenanceAssignmentDedicatedHostSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaintenanceAssignmentDedicatedHostSpec.
func (in *MaintenanceAssignmentDedicatedHostSpec) DeepCopy() *MaintenanceAssignmentDedicatedHostSpec {
	if in == nil {
		return nil
	}
	out := new(MaintenanceAssignmentDedicatedHostSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaintenanceAssignmentDedicatedHostStatus) DeepCopyInto(out *MaintenanceAssignmentDedicatedHostStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaintenanceAssignmentDedicatedHostStatus.
func (in *MaintenanceAssignmentDedicatedHostStatus) DeepCopy() *MaintenanceAssignmentDedicatedHostStatus {
	if in == nil {
		return nil
	}
	out := new(MaintenanceAssignmentDedicatedHostStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaintenanceAssignmentVirtualMachine) DeepCopyInto(out *MaintenanceAssignmentVirtualMachine) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaintenanceAssignmentVirtualMachine.
func (in *MaintenanceAssignmentVirtualMachine) DeepCopy() *MaintenanceAssignmentVirtualMachine {
	if in == nil {
		return nil
	}
	out := new(MaintenanceAssignmentVirtualMachine)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MaintenanceAssignmentVirtualMachine) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaintenanceAssignmentVirtualMachineList) DeepCopyInto(out *MaintenanceAssignmentVirtualMachineList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MaintenanceAssignmentVirtualMachine, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaintenanceAssignmentVirtualMachineList.
func (in *MaintenanceAssignmentVirtualMachineList) DeepCopy() *MaintenanceAssignmentVirtualMachineList {
	if in == nil {
		return nil
	}
	out := new(MaintenanceAssignmentVirtualMachineList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MaintenanceAssignmentVirtualMachineList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaintenanceAssignmentVirtualMachineObservation) DeepCopyInto(out *MaintenanceAssignmentVirtualMachineObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaintenanceAssignmentVirtualMachineObservation.
func (in *MaintenanceAssignmentVirtualMachineObservation) DeepCopy() *MaintenanceAssignmentVirtualMachineObservation {
	if in == nil {
		return nil
	}
	out := new(MaintenanceAssignmentVirtualMachineObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaintenanceAssignmentVirtualMachineParameters) DeepCopyInto(out *MaintenanceAssignmentVirtualMachineParameters) {
	*out = *in
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.MaintenanceConfigurationID != nil {
		in, out := &in.MaintenanceConfigurationID, &out.MaintenanceConfigurationID
		*out = new(string)
		**out = **in
	}
	if in.MaintenanceConfigurationIDRef != nil {
		in, out := &in.MaintenanceConfigurationIDRef, &out.MaintenanceConfigurationIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.MaintenanceConfigurationIDSelector != nil {
		in, out := &in.MaintenanceConfigurationIDSelector, &out.MaintenanceConfigurationIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.VirtualMachineID != nil {
		in, out := &in.VirtualMachineID, &out.VirtualMachineID
		*out = new(string)
		**out = **in
	}
	if in.VirtualMachineIDRef != nil {
		in, out := &in.VirtualMachineIDRef, &out.VirtualMachineIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.VirtualMachineIDSelector != nil {
		in, out := &in.VirtualMachineIDSelector, &out.VirtualMachineIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaintenanceAssignmentVirtualMachineParameters.
func (in *MaintenanceAssignmentVirtualMachineParameters) DeepCopy() *MaintenanceAssignmentVirtualMachineParameters {
	if in == nil {
		return nil
	}
	out := new(MaintenanceAssignmentVirtualMachineParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaintenanceAssignmentVirtualMachineSpec) DeepCopyInto(out *MaintenanceAssignmentVirtualMachineSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaintenanceAssignmentVirtualMachineSpec.
func (in *MaintenanceAssignmentVirtualMachineSpec) DeepCopy() *MaintenanceAssignmentVirtualMachineSpec {
	if in == nil {
		return nil
	}
	out := new(MaintenanceAssignmentVirtualMachineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaintenanceAssignmentVirtualMachineStatus) DeepCopyInto(out *MaintenanceAssignmentVirtualMachineStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaintenanceAssignmentVirtualMachineStatus.
func (in *MaintenanceAssignmentVirtualMachineStatus) DeepCopy() *MaintenanceAssignmentVirtualMachineStatus {
	if in == nil {
		return nil
	}
	out := new(MaintenanceAssignmentVirtualMachineStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaintenanceConfiguration) DeepCopyInto(out *MaintenanceConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaintenanceConfiguration.
func (in *MaintenanceConfiguration) DeepCopy() *MaintenanceConfiguration {
	if in == nil {
		return nil
	}
	out := new(MaintenanceConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MaintenanceConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaintenanceConfigurationList) DeepCopyInto(out *MaintenanceConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MaintenanceConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaintenanceConfigurationList.
func (in *MaintenanceConfigurationList) DeepCopy() *MaintenanceConfigurationList {
	if in == nil {
		return nil
	}
	out := new(MaintenanceConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MaintenanceConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaintenanceConfigurationObservation) DeepCopyInto(out *MaintenanceConfigurationObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaintenanceConfigurationObservation.
func (in *MaintenanceConfigurationObservation) DeepCopy() *MaintenanceConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(MaintenanceConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaintenanceConfigurationParameters) DeepCopyInto(out *MaintenanceConfigurationParameters) {
	*out = *in
	if in.InGuestUserPatchMode != nil {
		in, out := &in.InGuestUserPatchMode, &out.InGuestUserPatchMode
		*out = new(string)
		**out = **in
	}
	if in.InstallPatches != nil {
		in, out := &in.InstallPatches, &out.InstallPatches
		*out = make([]InstallPatchesParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupNameRef != nil {
		in, out := &in.ResourceGroupNameRef, &out.ResourceGroupNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ResourceGroupNameSelector != nil {
		in, out := &in.ResourceGroupNameSelector, &out.ResourceGroupNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Scope != nil {
		in, out := &in.Scope, &out.Scope
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Visibility != nil {
		in, out := &in.Visibility, &out.Visibility
		*out = new(string)
		**out = **in
	}
	if in.Window != nil {
		in, out := &in.Window, &out.Window
		*out = make([]WindowParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaintenanceConfigurationParameters.
func (in *MaintenanceConfigurationParameters) DeepCopy() *MaintenanceConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(MaintenanceConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaintenanceConfigurationSpec) DeepCopyInto(out *MaintenanceConfigurationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaintenanceConfigurationSpec.
func (in *MaintenanceConfigurationSpec) DeepCopy() *MaintenanceConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(MaintenanceConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaintenanceConfigurationStatus) DeepCopyInto(out *MaintenanceConfigurationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaintenanceConfigurationStatus.
func (in *MaintenanceConfigurationStatus) DeepCopy() *MaintenanceConfigurationStatus {
	if in == nil {
		return nil
	}
	out := new(MaintenanceConfigurationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WindowObservation) DeepCopyInto(out *WindowObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WindowObservation.
func (in *WindowObservation) DeepCopy() *WindowObservation {
	if in == nil {
		return nil
	}
	out := new(WindowObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WindowParameters) DeepCopyInto(out *WindowParameters) {
	*out = *in
	if in.Duration != nil {
		in, out := &in.Duration, &out.Duration
		*out = new(string)
		**out = **in
	}
	if in.ExpirationDateTime != nil {
		in, out := &in.ExpirationDateTime, &out.ExpirationDateTime
		*out = new(string)
		**out = **in
	}
	if in.RecurEvery != nil {
		in, out := &in.RecurEvery, &out.RecurEvery
		*out = new(string)
		**out = **in
	}
	if in.StartDateTime != nil {
		in, out := &in.StartDateTime, &out.StartDateTime
		*out = new(string)
		**out = **in
	}
	if in.TimeZone != nil {
		in, out := &in.TimeZone, &out.TimeZone
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WindowParameters.
func (in *WindowParameters) DeepCopy() *WindowParameters {
	if in == nil {
		return nil
	}
	out := new(WindowParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WindowsObservation) DeepCopyInto(out *WindowsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WindowsObservation.
func (in *WindowsObservation) DeepCopy() *WindowsObservation {
	if in == nil {
		return nil
	}
	out := new(WindowsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WindowsParameters) DeepCopyInto(out *WindowsParameters) {
	*out = *in
	if in.ClassificationsToInclude != nil {
		in, out := &in.ClassificationsToInclude, &out.ClassificationsToInclude
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.KbNumbersToExclude != nil {
		in, out := &in.KbNumbersToExclude, &out.KbNumbersToExclude
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.KbNumbersToInclude != nil {
		in, out := &in.KbNumbersToInclude, &out.KbNumbersToInclude
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WindowsParameters.
func (in *WindowsParameters) DeepCopy() *WindowsParameters {
	if in == nil {
		return nil
	}
	out := new(WindowsParameters)
	in.DeepCopyInto(out)
	return out
}