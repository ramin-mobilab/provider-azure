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
func (in *AppIntegrationAccount) DeepCopyInto(out *AppIntegrationAccount) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppIntegrationAccount.
func (in *AppIntegrationAccount) DeepCopy() *AppIntegrationAccount {
	if in == nil {
		return nil
	}
	out := new(AppIntegrationAccount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppIntegrationAccount) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppIntegrationAccountBatchConfiguration) DeepCopyInto(out *AppIntegrationAccountBatchConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppIntegrationAccountBatchConfiguration.
func (in *AppIntegrationAccountBatchConfiguration) DeepCopy() *AppIntegrationAccountBatchConfiguration {
	if in == nil {
		return nil
	}
	out := new(AppIntegrationAccountBatchConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppIntegrationAccountBatchConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppIntegrationAccountBatchConfigurationList) DeepCopyInto(out *AppIntegrationAccountBatchConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AppIntegrationAccountBatchConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppIntegrationAccountBatchConfigurationList.
func (in *AppIntegrationAccountBatchConfigurationList) DeepCopy() *AppIntegrationAccountBatchConfigurationList {
	if in == nil {
		return nil
	}
	out := new(AppIntegrationAccountBatchConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppIntegrationAccountBatchConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppIntegrationAccountBatchConfigurationObservation) DeepCopyInto(out *AppIntegrationAccountBatchConfigurationObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppIntegrationAccountBatchConfigurationObservation.
func (in *AppIntegrationAccountBatchConfigurationObservation) DeepCopy() *AppIntegrationAccountBatchConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(AppIntegrationAccountBatchConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppIntegrationAccountBatchConfigurationParameters) DeepCopyInto(out *AppIntegrationAccountBatchConfigurationParameters) {
	*out = *in
	if in.BatchGroupName != nil {
		in, out := &in.BatchGroupName, &out.BatchGroupName
		*out = new(string)
		**out = **in
	}
	if in.IntegrationAccountName != nil {
		in, out := &in.IntegrationAccountName, &out.IntegrationAccountName
		*out = new(string)
		**out = **in
	}
	if in.IntegrationAccountNameRef != nil {
		in, out := &in.IntegrationAccountNameRef, &out.IntegrationAccountNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.IntegrationAccountNameSelector != nil {
		in, out := &in.IntegrationAccountNameSelector, &out.IntegrationAccountNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
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
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ReleaseCriteria != nil {
		in, out := &in.ReleaseCriteria, &out.ReleaseCriteria
		*out = make([]ReleaseCriteriaParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppIntegrationAccountBatchConfigurationParameters.
func (in *AppIntegrationAccountBatchConfigurationParameters) DeepCopy() *AppIntegrationAccountBatchConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(AppIntegrationAccountBatchConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppIntegrationAccountBatchConfigurationSpec) DeepCopyInto(out *AppIntegrationAccountBatchConfigurationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppIntegrationAccountBatchConfigurationSpec.
func (in *AppIntegrationAccountBatchConfigurationSpec) DeepCopy() *AppIntegrationAccountBatchConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(AppIntegrationAccountBatchConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppIntegrationAccountBatchConfigurationStatus) DeepCopyInto(out *AppIntegrationAccountBatchConfigurationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppIntegrationAccountBatchConfigurationStatus.
func (in *AppIntegrationAccountBatchConfigurationStatus) DeepCopy() *AppIntegrationAccountBatchConfigurationStatus {
	if in == nil {
		return nil
	}
	out := new(AppIntegrationAccountBatchConfigurationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppIntegrationAccountList) DeepCopyInto(out *AppIntegrationAccountList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AppIntegrationAccount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppIntegrationAccountList.
func (in *AppIntegrationAccountList) DeepCopy() *AppIntegrationAccountList {
	if in == nil {
		return nil
	}
	out := new(AppIntegrationAccountList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppIntegrationAccountList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppIntegrationAccountObservation) DeepCopyInto(out *AppIntegrationAccountObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppIntegrationAccountObservation.
func (in *AppIntegrationAccountObservation) DeepCopy() *AppIntegrationAccountObservation {
	if in == nil {
		return nil
	}
	out := new(AppIntegrationAccountObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppIntegrationAccountParameters) DeepCopyInto(out *AppIntegrationAccountParameters) {
	*out = *in
	if in.IntegrationServiceEnvironmentID != nil {
		in, out := &in.IntegrationServiceEnvironmentID, &out.IntegrationServiceEnvironmentID
		*out = new(string)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
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
	if in.SkuName != nil {
		in, out := &in.SkuName, &out.SkuName
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppIntegrationAccountParameters.
func (in *AppIntegrationAccountParameters) DeepCopy() *AppIntegrationAccountParameters {
	if in == nil {
		return nil
	}
	out := new(AppIntegrationAccountParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppIntegrationAccountSpec) DeepCopyInto(out *AppIntegrationAccountSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppIntegrationAccountSpec.
func (in *AppIntegrationAccountSpec) DeepCopy() *AppIntegrationAccountSpec {
	if in == nil {
		return nil
	}
	out := new(AppIntegrationAccountSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppIntegrationAccountStatus) DeepCopyInto(out *AppIntegrationAccountStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppIntegrationAccountStatus.
func (in *AppIntegrationAccountStatus) DeepCopy() *AppIntegrationAccountStatus {
	if in == nil {
		return nil
	}
	out := new(AppIntegrationAccountStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IntegrationServiceEnvironment) DeepCopyInto(out *IntegrationServiceEnvironment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IntegrationServiceEnvironment.
func (in *IntegrationServiceEnvironment) DeepCopy() *IntegrationServiceEnvironment {
	if in == nil {
		return nil
	}
	out := new(IntegrationServiceEnvironment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IntegrationServiceEnvironment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IntegrationServiceEnvironmentList) DeepCopyInto(out *IntegrationServiceEnvironmentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IntegrationServiceEnvironment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IntegrationServiceEnvironmentList.
func (in *IntegrationServiceEnvironmentList) DeepCopy() *IntegrationServiceEnvironmentList {
	if in == nil {
		return nil
	}
	out := new(IntegrationServiceEnvironmentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IntegrationServiceEnvironmentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IntegrationServiceEnvironmentObservation) DeepCopyInto(out *IntegrationServiceEnvironmentObservation) {
	*out = *in
	if in.ConnectorEndpointIPAddresses != nil {
		in, out := &in.ConnectorEndpointIPAddresses, &out.ConnectorEndpointIPAddresses
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ConnectorOutboundIPAddresses != nil {
		in, out := &in.ConnectorOutboundIPAddresses, &out.ConnectorOutboundIPAddresses
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.WorkflowEndpointIPAddresses != nil {
		in, out := &in.WorkflowEndpointIPAddresses, &out.WorkflowEndpointIPAddresses
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.WorkflowOutboundIPAddresses != nil {
		in, out := &in.WorkflowOutboundIPAddresses, &out.WorkflowOutboundIPAddresses
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IntegrationServiceEnvironmentObservation.
func (in *IntegrationServiceEnvironmentObservation) DeepCopy() *IntegrationServiceEnvironmentObservation {
	if in == nil {
		return nil
	}
	out := new(IntegrationServiceEnvironmentObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IntegrationServiceEnvironmentParameters) DeepCopyInto(out *IntegrationServiceEnvironmentParameters) {
	*out = *in
	if in.AccessEndpointType != nil {
		in, out := &in.AccessEndpointType, &out.AccessEndpointType
		*out = new(string)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
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
	if in.SkuName != nil {
		in, out := &in.SkuName, &out.SkuName
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
	if in.VirtualNetworkSubnetIds != nil {
		in, out := &in.VirtualNetworkSubnetIds, &out.VirtualNetworkSubnetIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.VirtualNetworkSubnetIdsRefs != nil {
		in, out := &in.VirtualNetworkSubnetIdsRefs, &out.VirtualNetworkSubnetIdsRefs
		*out = make([]v1.Reference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VirtualNetworkSubnetIdsSelector != nil {
		in, out := &in.VirtualNetworkSubnetIdsSelector, &out.VirtualNetworkSubnetIdsSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IntegrationServiceEnvironmentParameters.
func (in *IntegrationServiceEnvironmentParameters) DeepCopy() *IntegrationServiceEnvironmentParameters {
	if in == nil {
		return nil
	}
	out := new(IntegrationServiceEnvironmentParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IntegrationServiceEnvironmentSpec) DeepCopyInto(out *IntegrationServiceEnvironmentSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IntegrationServiceEnvironmentSpec.
func (in *IntegrationServiceEnvironmentSpec) DeepCopy() *IntegrationServiceEnvironmentSpec {
	if in == nil {
		return nil
	}
	out := new(IntegrationServiceEnvironmentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IntegrationServiceEnvironmentStatus) DeepCopyInto(out *IntegrationServiceEnvironmentStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IntegrationServiceEnvironmentStatus.
func (in *IntegrationServiceEnvironmentStatus) DeepCopy() *IntegrationServiceEnvironmentStatus {
	if in == nil {
		return nil
	}
	out := new(IntegrationServiceEnvironmentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonthlyObservation) DeepCopyInto(out *MonthlyObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonthlyObservation.
func (in *MonthlyObservation) DeepCopy() *MonthlyObservation {
	if in == nil {
		return nil
	}
	out := new(MonthlyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonthlyParameters) DeepCopyInto(out *MonthlyParameters) {
	*out = *in
	if in.Week != nil {
		in, out := &in.Week, &out.Week
		*out = new(float64)
		**out = **in
	}
	if in.Weekday != nil {
		in, out := &in.Weekday, &out.Weekday
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonthlyParameters.
func (in *MonthlyParameters) DeepCopy() *MonthlyParameters {
	if in == nil {
		return nil
	}
	out := new(MonthlyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecurrenceObservation) DeepCopyInto(out *RecurrenceObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecurrenceObservation.
func (in *RecurrenceObservation) DeepCopy() *RecurrenceObservation {
	if in == nil {
		return nil
	}
	out := new(RecurrenceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecurrenceParameters) DeepCopyInto(out *RecurrenceParameters) {
	*out = *in
	if in.EndTime != nil {
		in, out := &in.EndTime, &out.EndTime
		*out = new(string)
		**out = **in
	}
	if in.Frequency != nil {
		in, out := &in.Frequency, &out.Frequency
		*out = new(string)
		**out = **in
	}
	if in.Interval != nil {
		in, out := &in.Interval, &out.Interval
		*out = new(float64)
		**out = **in
	}
	if in.Schedule != nil {
		in, out := &in.Schedule, &out.Schedule
		*out = make([]ScheduleParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = new(string)
		**out = **in
	}
	if in.TimeZone != nil {
		in, out := &in.TimeZone, &out.TimeZone
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecurrenceParameters.
func (in *RecurrenceParameters) DeepCopy() *RecurrenceParameters {
	if in == nil {
		return nil
	}
	out := new(RecurrenceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseCriteriaObservation) DeepCopyInto(out *ReleaseCriteriaObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseCriteriaObservation.
func (in *ReleaseCriteriaObservation) DeepCopy() *ReleaseCriteriaObservation {
	if in == nil {
		return nil
	}
	out := new(ReleaseCriteriaObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseCriteriaParameters) DeepCopyInto(out *ReleaseCriteriaParameters) {
	*out = *in
	if in.BatchSize != nil {
		in, out := &in.BatchSize, &out.BatchSize
		*out = new(float64)
		**out = **in
	}
	if in.MessageCount != nil {
		in, out := &in.MessageCount, &out.MessageCount
		*out = new(float64)
		**out = **in
	}
	if in.Recurrence != nil {
		in, out := &in.Recurrence, &out.Recurrence
		*out = make([]RecurrenceParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseCriteriaParameters.
func (in *ReleaseCriteriaParameters) DeepCopy() *ReleaseCriteriaParameters {
	if in == nil {
		return nil
	}
	out := new(ReleaseCriteriaParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduleObservation) DeepCopyInto(out *ScheduleObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduleObservation.
func (in *ScheduleObservation) DeepCopy() *ScheduleObservation {
	if in == nil {
		return nil
	}
	out := new(ScheduleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduleParameters) DeepCopyInto(out *ScheduleParameters) {
	*out = *in
	if in.Hours != nil {
		in, out := &in.Hours, &out.Hours
		*out = make([]*float64, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(float64)
				**out = **in
			}
		}
	}
	if in.Minutes != nil {
		in, out := &in.Minutes, &out.Minutes
		*out = make([]*float64, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(float64)
				**out = **in
			}
		}
	}
	if in.MonthDays != nil {
		in, out := &in.MonthDays, &out.MonthDays
		*out = make([]*float64, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(float64)
				**out = **in
			}
		}
	}
	if in.Monthly != nil {
		in, out := &in.Monthly, &out.Monthly
		*out = make([]MonthlyParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.WeekDays != nil {
		in, out := &in.WeekDays, &out.WeekDays
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduleParameters.
func (in *ScheduleParameters) DeepCopy() *ScheduleParameters {
	if in == nil {
		return nil
	}
	out := new(ScheduleParameters)
	in.DeepCopyInto(out)
	return out
}
