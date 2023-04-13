/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this AccessPoint.
func (mg *AccessPoint) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this AccessPoint.
func (mg *AccessPoint) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicy of this AccessPoint.
func (mg *AccessPoint) GetManagementPolicy() xpv1.ManagementPolicy {
	return mg.Spec.ManagementPolicy
}

// GetProviderConfigReference of this AccessPoint.
func (mg *AccessPoint) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this AccessPoint.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *AccessPoint) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this AccessPoint.
func (mg *AccessPoint) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this AccessPoint.
func (mg *AccessPoint) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this AccessPoint.
func (mg *AccessPoint) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this AccessPoint.
func (mg *AccessPoint) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicy of this AccessPoint.
func (mg *AccessPoint) SetManagementPolicy(r xpv1.ManagementPolicy) {
	mg.Spec.ManagementPolicy = r
}

// SetProviderConfigReference of this AccessPoint.
func (mg *AccessPoint) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this AccessPoint.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *AccessPoint) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this AccessPoint.
func (mg *AccessPoint) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this AccessPoint.
func (mg *AccessPoint) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this AccessPointPolicy.
func (mg *AccessPointPolicy) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this AccessPointPolicy.
func (mg *AccessPointPolicy) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicy of this AccessPointPolicy.
func (mg *AccessPointPolicy) GetManagementPolicy() xpv1.ManagementPolicy {
	return mg.Spec.ManagementPolicy
}

// GetProviderConfigReference of this AccessPointPolicy.
func (mg *AccessPointPolicy) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this AccessPointPolicy.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *AccessPointPolicy) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this AccessPointPolicy.
func (mg *AccessPointPolicy) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this AccessPointPolicy.
func (mg *AccessPointPolicy) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this AccessPointPolicy.
func (mg *AccessPointPolicy) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this AccessPointPolicy.
func (mg *AccessPointPolicy) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicy of this AccessPointPolicy.
func (mg *AccessPointPolicy) SetManagementPolicy(r xpv1.ManagementPolicy) {
	mg.Spec.ManagementPolicy = r
}

// SetProviderConfigReference of this AccessPointPolicy.
func (mg *AccessPointPolicy) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this AccessPointPolicy.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *AccessPointPolicy) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this AccessPointPolicy.
func (mg *AccessPointPolicy) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this AccessPointPolicy.
func (mg *AccessPointPolicy) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this AccountPublicAccessBlock.
func (mg *AccountPublicAccessBlock) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this AccountPublicAccessBlock.
func (mg *AccountPublicAccessBlock) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicy of this AccountPublicAccessBlock.
func (mg *AccountPublicAccessBlock) GetManagementPolicy() xpv1.ManagementPolicy {
	return mg.Spec.ManagementPolicy
}

// GetProviderConfigReference of this AccountPublicAccessBlock.
func (mg *AccountPublicAccessBlock) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this AccountPublicAccessBlock.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *AccountPublicAccessBlock) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this AccountPublicAccessBlock.
func (mg *AccountPublicAccessBlock) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this AccountPublicAccessBlock.
func (mg *AccountPublicAccessBlock) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this AccountPublicAccessBlock.
func (mg *AccountPublicAccessBlock) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this AccountPublicAccessBlock.
func (mg *AccountPublicAccessBlock) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicy of this AccountPublicAccessBlock.
func (mg *AccountPublicAccessBlock) SetManagementPolicy(r xpv1.ManagementPolicy) {
	mg.Spec.ManagementPolicy = r
}

// SetProviderConfigReference of this AccountPublicAccessBlock.
func (mg *AccountPublicAccessBlock) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this AccountPublicAccessBlock.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *AccountPublicAccessBlock) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this AccountPublicAccessBlock.
func (mg *AccountPublicAccessBlock) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this AccountPublicAccessBlock.
func (mg *AccountPublicAccessBlock) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this MultiRegionAccessPoint.
func (mg *MultiRegionAccessPoint) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MultiRegionAccessPoint.
func (mg *MultiRegionAccessPoint) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicy of this MultiRegionAccessPoint.
func (mg *MultiRegionAccessPoint) GetManagementPolicy() xpv1.ManagementPolicy {
	return mg.Spec.ManagementPolicy
}

// GetProviderConfigReference of this MultiRegionAccessPoint.
func (mg *MultiRegionAccessPoint) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this MultiRegionAccessPoint.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *MultiRegionAccessPoint) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this MultiRegionAccessPoint.
func (mg *MultiRegionAccessPoint) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this MultiRegionAccessPoint.
func (mg *MultiRegionAccessPoint) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MultiRegionAccessPoint.
func (mg *MultiRegionAccessPoint) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MultiRegionAccessPoint.
func (mg *MultiRegionAccessPoint) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicy of this MultiRegionAccessPoint.
func (mg *MultiRegionAccessPoint) SetManagementPolicy(r xpv1.ManagementPolicy) {
	mg.Spec.ManagementPolicy = r
}

// SetProviderConfigReference of this MultiRegionAccessPoint.
func (mg *MultiRegionAccessPoint) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this MultiRegionAccessPoint.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *MultiRegionAccessPoint) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this MultiRegionAccessPoint.
func (mg *MultiRegionAccessPoint) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this MultiRegionAccessPoint.
func (mg *MultiRegionAccessPoint) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this MultiRegionAccessPointPolicy.
func (mg *MultiRegionAccessPointPolicy) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MultiRegionAccessPointPolicy.
func (mg *MultiRegionAccessPointPolicy) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicy of this MultiRegionAccessPointPolicy.
func (mg *MultiRegionAccessPointPolicy) GetManagementPolicy() xpv1.ManagementPolicy {
	return mg.Spec.ManagementPolicy
}

// GetProviderConfigReference of this MultiRegionAccessPointPolicy.
func (mg *MultiRegionAccessPointPolicy) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this MultiRegionAccessPointPolicy.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *MultiRegionAccessPointPolicy) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this MultiRegionAccessPointPolicy.
func (mg *MultiRegionAccessPointPolicy) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this MultiRegionAccessPointPolicy.
func (mg *MultiRegionAccessPointPolicy) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MultiRegionAccessPointPolicy.
func (mg *MultiRegionAccessPointPolicy) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MultiRegionAccessPointPolicy.
func (mg *MultiRegionAccessPointPolicy) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicy of this MultiRegionAccessPointPolicy.
func (mg *MultiRegionAccessPointPolicy) SetManagementPolicy(r xpv1.ManagementPolicy) {
	mg.Spec.ManagementPolicy = r
}

// SetProviderConfigReference of this MultiRegionAccessPointPolicy.
func (mg *MultiRegionAccessPointPolicy) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this MultiRegionAccessPointPolicy.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *MultiRegionAccessPointPolicy) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this MultiRegionAccessPointPolicy.
func (mg *MultiRegionAccessPointPolicy) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this MultiRegionAccessPointPolicy.
func (mg *MultiRegionAccessPointPolicy) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this ObjectLambdaAccessPoint.
func (mg *ObjectLambdaAccessPoint) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this ObjectLambdaAccessPoint.
func (mg *ObjectLambdaAccessPoint) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicy of this ObjectLambdaAccessPoint.
func (mg *ObjectLambdaAccessPoint) GetManagementPolicy() xpv1.ManagementPolicy {
	return mg.Spec.ManagementPolicy
}

// GetProviderConfigReference of this ObjectLambdaAccessPoint.
func (mg *ObjectLambdaAccessPoint) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this ObjectLambdaAccessPoint.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *ObjectLambdaAccessPoint) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this ObjectLambdaAccessPoint.
func (mg *ObjectLambdaAccessPoint) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this ObjectLambdaAccessPoint.
func (mg *ObjectLambdaAccessPoint) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this ObjectLambdaAccessPoint.
func (mg *ObjectLambdaAccessPoint) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this ObjectLambdaAccessPoint.
func (mg *ObjectLambdaAccessPoint) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicy of this ObjectLambdaAccessPoint.
func (mg *ObjectLambdaAccessPoint) SetManagementPolicy(r xpv1.ManagementPolicy) {
	mg.Spec.ManagementPolicy = r
}

// SetProviderConfigReference of this ObjectLambdaAccessPoint.
func (mg *ObjectLambdaAccessPoint) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this ObjectLambdaAccessPoint.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *ObjectLambdaAccessPoint) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this ObjectLambdaAccessPoint.
func (mg *ObjectLambdaAccessPoint) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this ObjectLambdaAccessPoint.
func (mg *ObjectLambdaAccessPoint) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this ObjectLambdaAccessPointPolicy.
func (mg *ObjectLambdaAccessPointPolicy) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this ObjectLambdaAccessPointPolicy.
func (mg *ObjectLambdaAccessPointPolicy) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicy of this ObjectLambdaAccessPointPolicy.
func (mg *ObjectLambdaAccessPointPolicy) GetManagementPolicy() xpv1.ManagementPolicy {
	return mg.Spec.ManagementPolicy
}

// GetProviderConfigReference of this ObjectLambdaAccessPointPolicy.
func (mg *ObjectLambdaAccessPointPolicy) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this ObjectLambdaAccessPointPolicy.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *ObjectLambdaAccessPointPolicy) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this ObjectLambdaAccessPointPolicy.
func (mg *ObjectLambdaAccessPointPolicy) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this ObjectLambdaAccessPointPolicy.
func (mg *ObjectLambdaAccessPointPolicy) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this ObjectLambdaAccessPointPolicy.
func (mg *ObjectLambdaAccessPointPolicy) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this ObjectLambdaAccessPointPolicy.
func (mg *ObjectLambdaAccessPointPolicy) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicy of this ObjectLambdaAccessPointPolicy.
func (mg *ObjectLambdaAccessPointPolicy) SetManagementPolicy(r xpv1.ManagementPolicy) {
	mg.Spec.ManagementPolicy = r
}

// SetProviderConfigReference of this ObjectLambdaAccessPointPolicy.
func (mg *ObjectLambdaAccessPointPolicy) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this ObjectLambdaAccessPointPolicy.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *ObjectLambdaAccessPointPolicy) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this ObjectLambdaAccessPointPolicy.
func (mg *ObjectLambdaAccessPointPolicy) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this ObjectLambdaAccessPointPolicy.
func (mg *ObjectLambdaAccessPointPolicy) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this StorageLensConfiguration.
func (mg *StorageLensConfiguration) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this StorageLensConfiguration.
func (mg *StorageLensConfiguration) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicy of this StorageLensConfiguration.
func (mg *StorageLensConfiguration) GetManagementPolicy() xpv1.ManagementPolicy {
	return mg.Spec.ManagementPolicy
}

// GetProviderConfigReference of this StorageLensConfiguration.
func (mg *StorageLensConfiguration) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this StorageLensConfiguration.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *StorageLensConfiguration) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this StorageLensConfiguration.
func (mg *StorageLensConfiguration) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this StorageLensConfiguration.
func (mg *StorageLensConfiguration) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this StorageLensConfiguration.
func (mg *StorageLensConfiguration) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this StorageLensConfiguration.
func (mg *StorageLensConfiguration) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicy of this StorageLensConfiguration.
func (mg *StorageLensConfiguration) SetManagementPolicy(r xpv1.ManagementPolicy) {
	mg.Spec.ManagementPolicy = r
}

// SetProviderConfigReference of this StorageLensConfiguration.
func (mg *StorageLensConfiguration) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this StorageLensConfiguration.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *StorageLensConfiguration) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this StorageLensConfiguration.
func (mg *StorageLensConfiguration) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this StorageLensConfiguration.
func (mg *StorageLensConfiguration) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}
