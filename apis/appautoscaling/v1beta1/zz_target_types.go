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

type TargetObservation struct {

	// The ARN of the scalable target.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Max capacity of the scalable target.
	MaxCapacity *float64 `json:"maxCapacity,omitempty" tf:"max_capacity,omitempty"`

	// Min capacity of the scalable target.
	MinCapacity *float64 `json:"minCapacity,omitempty" tf:"min_capacity,omitempty"`

	// Resource type and unique identifier string for the resource associated with the scaling policy. Documentation can be found in the ResourceId parameter at: AWS Application Auto Scaling API Reference
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// ARN of the IAM role that allows Application AutoScaling to modify your scalable target on your behalf. This defaults to an IAM Service-Linked Role for most services and custom IAM Roles are ignored by the API for those namespaces. See the AWS Application Auto Scaling documentation for more information about how this service interacts with IAM.
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// Scalable dimension of the scalable target. Documentation can be found in the ScalableDimension parameter at: AWS Application Auto Scaling API Reference
	ScalableDimension *string `json:"scalableDimension,omitempty" tf:"scalable_dimension,omitempty"`

	// AWS service namespace of the scalable target. Documentation can be found in the ServiceNamespace parameter at: AWS Application Auto Scaling API Reference
	ServiceNamespace *string `json:"serviceNamespace,omitempty" tf:"service_namespace,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type TargetParameters struct {

	// Max capacity of the scalable target.
	// +kubebuilder:validation:Optional
	MaxCapacity *float64 `json:"maxCapacity,omitempty" tf:"max_capacity,omitempty"`

	// Min capacity of the scalable target.
	// +kubebuilder:validation:Optional
	MinCapacity *float64 `json:"minCapacity,omitempty" tf:"min_capacity,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Resource type and unique identifier string for the resource associated with the scaling policy. Documentation can be found in the ResourceId parameter at: AWS Application Auto Scaling API Reference
	// +kubebuilder:validation:Required
	ResourceID *string `json:"resourceId" tf:"resource_id,omitempty"`

	// ARN of the IAM role that allows Application AutoScaling to modify your scalable target on your behalf. This defaults to an IAM Service-Linked Role for most services and custom IAM Roles are ignored by the API for those namespaces. See the AWS Application Auto Scaling documentation for more information about how this service interacts with IAM.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// Reference to a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnRef *v1.Reference `json:"roleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnSelector *v1.Selector `json:"roleArnSelector,omitempty" tf:"-"`

	// Scalable dimension of the scalable target. Documentation can be found in the ScalableDimension parameter at: AWS Application Auto Scaling API Reference
	// +kubebuilder:validation:Required
	ScalableDimension *string `json:"scalableDimension" tf:"scalable_dimension,omitempty"`

	// AWS service namespace of the scalable target. Documentation can be found in the ServiceNamespace parameter at: AWS Application Auto Scaling API Reference
	// +kubebuilder:validation:Required
	ServiceNamespace *string `json:"serviceNamespace" tf:"service_namespace,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// TargetSpec defines the desired state of Target
type TargetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TargetParameters `json:"forProvider"`
}

// TargetStatus defines the observed state of Target.
type TargetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TargetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Target is the Schema for the Targets API. Provides an Application AutoScaling ScalableTarget resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Target struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.maxCapacity)",message="maxCapacity is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.minCapacity)",message="minCapacity is a required parameter"
	Spec   TargetSpec   `json:"spec"`
	Status TargetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TargetList contains a list of Targets
type TargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Target `json:"items"`
}

// Repository type metadata.
var (
	Target_Kind             = "Target"
	Target_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Target_Kind}.String()
	Target_KindAPIVersion   = Target_Kind + "." + CRDGroupVersion.String()
	Target_GroupVersionKind = CRDGroupVersion.WithKind(Target_Kind)
)

func init() {
	SchemeBuilder.Register(&Target{}, &TargetList{})
}
