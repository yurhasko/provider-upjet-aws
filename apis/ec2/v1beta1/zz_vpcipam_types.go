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

type OperatingRegionsObservation struct {

	// The name of the Region you want to add to the IPAM.
	RegionName *string `json:"regionName,omitempty" tf:"region_name,omitempty"`
}

type OperatingRegionsParameters struct {

	// The name of the Region you want to add to the IPAM.
	// +kubebuilder:validation:Required
	RegionName *string `json:"regionName" tf:"region_name,omitempty"`
}

type VPCIpamObservation struct {

	// Amazon Resource Name (ARN) of IPAM
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Enables you to quickly delete an IPAM, private scopes, pools in private scopes, and any allocations in the pools in private scopes.
	Cascade *bool `json:"cascade,omitempty" tf:"cascade,omitempty"`

	// The IPAM's default resource discovery association ID.
	DefaultResourceDiscoveryAssociationID *string `json:"defaultResourceDiscoveryAssociationId,omitempty" tf:"default_resource_discovery_association_id,omitempty"`

	// The IPAM's default resource discovery ID.
	DefaultResourceDiscoveryID *string `json:"defaultResourceDiscoveryId,omitempty" tf:"default_resource_discovery_id,omitempty"`

	// A description for the IPAM.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the IPAM
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Determines which locales can be chosen when you create pools. Locale is the Region where you want to make an IPAM pool available for allocations. You can only create pools with locales that match the operating Regions of the IPAM. You can only create VPCs from a pool whose locale matches the VPC's Region. You specify a region using the region_name parameter. You must set your provider block region as an operating_region.
	OperatingRegions []OperatingRegionsObservation `json:"operatingRegions,omitempty" tf:"operating_regions,omitempty"`

	// The ID of the IPAM's private scope. A scope is a top-level container in IPAM. Each scope represents an IP-independent network. Scopes enable you to represent networks where you have overlapping IP space. When you create an IPAM, IPAM automatically creates two scopes: public and private. The private scope is intended for private IP space. The public scope is intended for all internet-routable IP space.
	PrivateDefaultScopeID *string `json:"privateDefaultScopeId,omitempty" tf:"private_default_scope_id,omitempty"`

	// The ID of the IPAM's public scope. A scope is a top-level container in IPAM. Each scope represents an IP-independent network. Scopes enable you to represent networks where you have overlapping IP space. When you create an IPAM, IPAM automatically creates two scopes: public and private. The private scope is intended for private
	// IP space. The public scope is intended for all internet-routable IP space.
	PublicDefaultScopeID *string `json:"publicDefaultScopeId,omitempty" tf:"public_default_scope_id,omitempty"`

	// The number of scopes in the IPAM.
	ScopeCount *float64 `json:"scopeCount,omitempty" tf:"scope_count,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type VPCIpamParameters struct {

	// Enables you to quickly delete an IPAM, private scopes, pools in private scopes, and any allocations in the pools in private scopes.
	// +kubebuilder:validation:Optional
	Cascade *bool `json:"cascade,omitempty" tf:"cascade,omitempty"`

	// A description for the IPAM.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Determines which locales can be chosen when you create pools. Locale is the Region where you want to make an IPAM pool available for allocations. You can only create pools with locales that match the operating Regions of the IPAM. You can only create VPCs from a pool whose locale matches the VPC's Region. You specify a region using the region_name parameter. You must set your provider block region as an operating_region.
	// +kubebuilder:validation:Optional
	OperatingRegions []OperatingRegionsParameters `json:"operatingRegions,omitempty" tf:"operating_regions,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// VPCIpamSpec defines the desired state of VPCIpam
type VPCIpamSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VPCIpamParameters `json:"forProvider"`
}

// VPCIpamStatus defines the observed state of VPCIpam.
type VPCIpamStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VPCIpamObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VPCIpam is the Schema for the VPCIpams API. Provides an IPAM resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type VPCIpam struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.operatingRegions)",message="operatingRegions is a required parameter"
	Spec   VPCIpamSpec   `json:"spec"`
	Status VPCIpamStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VPCIpamList contains a list of VPCIpams
type VPCIpamList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPCIpam `json:"items"`
}

// Repository type metadata.
var (
	VPCIpam_Kind             = "VPCIpam"
	VPCIpam_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VPCIpam_Kind}.String()
	VPCIpam_KindAPIVersion   = VPCIpam_Kind + "." + CRDGroupVersion.String()
	VPCIpam_GroupVersionKind = CRDGroupVersion.WithKind(VPCIpam_Kind)
)

func init() {
	SchemeBuilder.Register(&VPCIpam{}, &VPCIpamList{})
}
