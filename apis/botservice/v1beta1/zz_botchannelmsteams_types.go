/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type BotChannelMSTeamsObservation struct {

	// The ID of the Microsoft Teams Integration for a Bot Channel.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type BotChannelMSTeamsParameters struct {

	// The name of the Bot Resource this channel will be associated with. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/botservice/v1beta1.BotChannelsRegistration
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("name",false)
	// +kubebuilder:validation:Optional
	BotName *string `json:"botName,omitempty" tf:"bot_name,omitempty"`

	// Reference to a BotChannelsRegistration in botservice to populate botName.
	// +kubebuilder:validation:Optional
	BotNameRef *v1.Reference `json:"botNameRef,omitempty" tf:"-"`

	// Selector for a BotChannelsRegistration in botservice to populate botName.
	// +kubebuilder:validation:Optional
	BotNameSelector *v1.Selector `json:"botNameSelector,omitempty" tf:"-"`

	// Specifies the webhook for Microsoft Teams channel calls.
	// +kubebuilder:validation:Optional
	CallingWebHook *string `json:"callingWebHook,omitempty" tf:"calling_web_hook,omitempty"`

	// Specifies whether to enable Microsoft Teams channel calls. This defaults to false.
	// +kubebuilder:validation:Optional
	EnableCalling *bool `json:"enableCalling,omitempty" tf:"enable_calling,omitempty"`

	// The supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// The name of the resource group in which to create the Bot Channel. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`
}

// BotChannelMSTeamsSpec defines the desired state of BotChannelMSTeams
type BotChannelMSTeamsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BotChannelMSTeamsParameters `json:"forProvider"`
}

// BotChannelMSTeamsStatus defines the observed state of BotChannelMSTeams.
type BotChannelMSTeamsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BotChannelMSTeamsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BotChannelMSTeams is the Schema for the BotChannelMSTeamss API. Manages an MS Teams integration for a Bot Channel
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type BotChannelMSTeams struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BotChannelMSTeamsSpec   `json:"spec"`
	Status            BotChannelMSTeamsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BotChannelMSTeamsList contains a list of BotChannelMSTeamss
type BotChannelMSTeamsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BotChannelMSTeams `json:"items"`
}

// Repository type metadata.
var (
	BotChannelMSTeams_Kind             = "BotChannelMSTeams"
	BotChannelMSTeams_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BotChannelMSTeams_Kind}.String()
	BotChannelMSTeams_KindAPIVersion   = BotChannelMSTeams_Kind + "." + CRDGroupVersion.String()
	BotChannelMSTeams_GroupVersionKind = CRDGroupVersion.WithKind(BotChannelMSTeams_Kind)
)

func init() {
	SchemeBuilder.Register(&BotChannelMSTeams{}, &BotChannelMSTeamsList{})
}