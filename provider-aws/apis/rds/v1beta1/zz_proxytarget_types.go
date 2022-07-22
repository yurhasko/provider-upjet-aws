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

type ProxyTargetObservation struct {
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	RDSResourceID *string `json:"rdsResourceId,omitempty" tf:"rds_resource_id,omitempty"`

	TargetArn *string `json:"targetArn,omitempty" tf:"target_arn,omitempty"`

	TrackedClusterID *string `json:"trackedClusterId,omitempty" tf:"tracked_cluster_id,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ProxyTargetParameters struct {

	// +kubebuilder:validation:Optional
	DBClusterIdentifier *string `json:"dbClusterIdentifier,omitempty" tf:"db_cluster_identifier,omitempty"`

	// +kubebuilder:validation:Optional
	DBInstanceIdentifier *string `json:"dbInstanceIdentifier,omitempty" tf:"db_instance_identifier,omitempty"`

	// +kubebuilder:validation:Required
	DBProxyName *string `json:"dbProxyName" tf:"db_proxy_name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	TargetGroupName *string `json:"targetGroupName" tf:"target_group_name,omitempty"`
}

// ProxyTargetSpec defines the desired state of ProxyTarget
type ProxyTargetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProxyTargetParameters `json:"forProvider"`
}

// ProxyTargetStatus defines the observed state of ProxyTarget.
type ProxyTargetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProxyTargetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ProxyTarget is the Schema for the ProxyTargets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ProxyTarget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProxyTargetSpec   `json:"spec"`
	Status            ProxyTargetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProxyTargetList contains a list of ProxyTargets
type ProxyTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProxyTarget `json:"items"`
}

// Repository type metadata.
var (
	ProxyTarget_Kind             = "ProxyTarget"
	ProxyTarget_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ProxyTarget_Kind}.String()
	ProxyTarget_KindAPIVersion   = ProxyTarget_Kind + "." + CRDGroupVersion.String()
	ProxyTarget_GroupVersionKind = CRDGroupVersion.WithKind(ProxyTarget_Kind)
)

func init() {
	SchemeBuilder.Register(&ProxyTarget{}, &ProxyTargetList{})
}