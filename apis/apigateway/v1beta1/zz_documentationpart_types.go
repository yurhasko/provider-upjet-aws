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

type DocumentationPartObservation struct {

	// The unique ID of the Documentation Part
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DocumentationPartParameters struct {

	// The location of the targeted API entity of the to-be-created documentation part. See below.
	// +kubebuilder:validation:Required
	Location []LocationParameters `json:"location" tf:"location,omitempty"`

	// A content map of API-specific key-value pairs describing the targeted API entity. The map must be encoded as a JSON string, e.g., "{ "description": "The API does ..." }". Only Swagger-compliant key-value pairs can be exported and, hence, published.
	// +kubebuilder:validation:Required
	Properties *string `json:"properties" tf:"properties,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The ID of the associated Rest API
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/apigateway/v1beta1.RestAPI
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	RestAPIID *string `json:"restApiId,omitempty" tf:"rest_api_id,omitempty"`

	// Reference to a RestAPI in apigateway to populate restApiId.
	// +kubebuilder:validation:Optional
	RestAPIIDRef *v1.Reference `json:"restApiIdRef,omitempty" tf:"-"`

	// Selector for a RestAPI in apigateway to populate restApiId.
	// +kubebuilder:validation:Optional
	RestAPIIDSelector *v1.Selector `json:"restApiIdSelector,omitempty" tf:"-"`
}

type LocationObservation struct {
}

type LocationParameters struct {

	// The HTTP verb of a method. The default value is * for any method.
	// +kubebuilder:validation:Optional
	Method *string `json:"method,omitempty" tf:"method,omitempty"`

	// The name of the targeted API entity.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The URL path of the target. The default value is / for the root resource.
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// The HTTP status code of a response. The default value is * for any status code.
	// +kubebuilder:validation:Optional
	StatusCode *string `json:"statusCode,omitempty" tf:"status_code,omitempty"`

	// The type of API entity to which the documentation content appliesE.g., API, METHOD or REQUEST_BODY
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

// DocumentationPartSpec defines the desired state of DocumentationPart
type DocumentationPartSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DocumentationPartParameters `json:"forProvider"`
}

// DocumentationPartStatus defines the observed state of DocumentationPart.
type DocumentationPartStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DocumentationPartObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DocumentationPart is the Schema for the DocumentationParts API. Provides a settings of an API Gateway Documentation Part.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type DocumentationPart struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DocumentationPartSpec   `json:"spec"`
	Status            DocumentationPartStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DocumentationPartList contains a list of DocumentationParts
type DocumentationPartList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DocumentationPart `json:"items"`
}

// Repository type metadata.
var (
	DocumentationPart_Kind             = "DocumentationPart"
	DocumentationPart_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DocumentationPart_Kind}.String()
	DocumentationPart_KindAPIVersion   = DocumentationPart_Kind + "." + CRDGroupVersion.String()
	DocumentationPart_GroupVersionKind = CRDGroupVersion.WithKind(DocumentationPart_Kind)
)

func init() {
	SchemeBuilder.Register(&DocumentationPart{}, &DocumentationPartList{})
}