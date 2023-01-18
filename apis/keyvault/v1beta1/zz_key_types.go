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

type KeyObservation struct {

	// The RSA public exponent of this Key Vault Key.
	E *string `json:"e,omitempty" tf:"e,omitempty"`

	// The Key Vault Key ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The RSA modulus of this Key Vault Key.
	N *string `json:"n,omitempty" tf:"n,omitempty"`

	// The OpenSSH encoded public key of this Key Vault Key.
	PublicKeyOpenssh *string `json:"publicKeyOpenssh,omitempty" tf:"public_key_openssh,omitempty"`

	// The PEM encoded public key of this Key Vault Key.
	PublicKeyPem *string `json:"publicKeyPem,omitempty" tf:"public_key_pem,omitempty"`

	// The (Versioned) ID for this Key Vault Key. This property points to a specific version of a Key Vault Key, as such using this won't auto-rotate values if used in other Azure Services.
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// The Versionless ID of the Key Vault Key. This property allows other Azure Services (that support it) to auto-rotate their value when the Key Vault Key is updated.
	ResourceVersionlessID *string `json:"resourceVersionlessId,omitempty" tf:"resource_versionless_id,omitempty"`

	// The current version of the Key Vault Key.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`

	// The Base ID of the Key Vault Key.
	VersionlessID *string `json:"versionlessId,omitempty" tf:"versionless_id,omitempty"`

	// The EC X component of this Key Vault Key.
	X *string `json:"x,omitempty" tf:"x,omitempty"`

	// The EC Y component of this Key Vault Key.
	Y *string `json:"y,omitempty" tf:"y,omitempty"`
}

type KeyParameters struct {

	// Specifies the curve to use when creating an EC key. Possible values are P-256, P-256K, P-384, and P-521. This field will be required in a future release if key_type is EC or EC-HSM. The API will default to P-256 if nothing is specified. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Curve *string `json:"curve,omitempty" tf:"curve,omitempty"`

	// Expiration UTC datetime (Y-m-d'T'H:M:S'Z').
	// +kubebuilder:validation:Optional
	ExpirationDate *string `json:"expirationDate,omitempty" tf:"expiration_date,omitempty"`

	// A list of JSON web key operations. Possible values include: decrypt, encrypt, sign, unwrapKey, verify and wrapKey. Please note these values are case sensitive.
	// +kubebuilder:validation:Required
	KeyOpts []*string `json:"keyOpts" tf:"key_opts,omitempty"`

	// Specifies the Size of the RSA key to create in bytes. For example, 1024 or 2048. Note: This field is required if key_type is RSA or RSA-HSM. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	KeySize *float64 `json:"keySize,omitempty" tf:"key_size,omitempty"`

	// Specifies the Key Type to use for this Key Vault Key. Possible values are EC (Elliptic Curve), EC-HSM, RSA and RSA-HSM. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	KeyType *string `json:"keyType" tf:"key_type,omitempty"`

	// The ID of the Key Vault where the Key should be created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=Vault
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	KeyVaultID *string `json:"keyVaultId,omitempty" tf:"key_vault_id,omitempty"`

	// Reference to a Vault to populate keyVaultId.
	// +kubebuilder:validation:Optional
	KeyVaultIDRef *v1.Reference `json:"keyVaultIdRef,omitempty" tf:"-"`

	// Selector for a Vault to populate keyVaultId.
	// +kubebuilder:validation:Optional
	KeyVaultIDSelector *v1.Selector `json:"keyVaultIdSelector,omitempty" tf:"-"`

	// Key not usable before the provided UTC datetime (Y-m-d'T'H:M:S'Z').
	// +kubebuilder:validation:Optional
	NotBeforeDate *string `json:"notBeforeDate,omitempty" tf:"not_before_date,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// KeySpec defines the desired state of Key
type KeySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     KeyParameters `json:"forProvider"`
}

// KeyStatus defines the observed state of Key.
type KeyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        KeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Key is the Schema for the Keys API. Manages a Key Vault Key.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Key struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KeySpec   `json:"spec"`
	Status            KeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KeyList contains a list of Keys
type KeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Key `json:"items"`
}

// Repository type metadata.
var (
	Key_Kind             = "Key"
	Key_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Key_Kind}.String()
	Key_KindAPIVersion   = Key_Kind + "." + CRDGroupVersion.String()
	Key_GroupVersionKind = CRDGroupVersion.WithKind(Key_Kind)
)

func init() {
	SchemeBuilder.Register(&Key{}, &KeyList{})
}
