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

type DatabaseMigrationServiceObservation struct {

	// The ID of Database Migration Service.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DatabaseMigrationServiceParameters struct {

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// Specify the name of the database migration service. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Name of the resource group in which to create the database migration service. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The SKU name of the database migration service. Possible values are Premium_4vCores, Standard_1vCores, Standard_2vCores and Standard_4vCores. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	SkuName *string `json:"skuName" tf:"sku_name,omitempty"`

	// The ID of the virtual subnet resource to which the database migration service should be joined. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.Subnet
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in network to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in network to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// A mapping of tags to assigned to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// DatabaseMigrationServiceSpec defines the desired state of DatabaseMigrationService
type DatabaseMigrationServiceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DatabaseMigrationServiceParameters `json:"forProvider"`
}

// DatabaseMigrationServiceStatus defines the observed state of DatabaseMigrationService.
type DatabaseMigrationServiceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DatabaseMigrationServiceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DatabaseMigrationService is the Schema for the DatabaseMigrationServices API. Manage a Azure Database Migration Service.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type DatabaseMigrationService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseMigrationServiceSpec   `json:"spec"`
	Status            DatabaseMigrationServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DatabaseMigrationServiceList contains a list of DatabaseMigrationServices
type DatabaseMigrationServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DatabaseMigrationService `json:"items"`
}

// Repository type metadata.
var (
	DatabaseMigrationService_Kind             = "DatabaseMigrationService"
	DatabaseMigrationService_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DatabaseMigrationService_Kind}.String()
	DatabaseMigrationService_KindAPIVersion   = DatabaseMigrationService_Kind + "." + CRDGroupVersion.String()
	DatabaseMigrationService_GroupVersionKind = CRDGroupVersion.WithKind(DatabaseMigrationService_Kind)
)

func init() {
	SchemeBuilder.Register(&DatabaseMigrationService{}, &DatabaseMigrationServiceList{})
}
