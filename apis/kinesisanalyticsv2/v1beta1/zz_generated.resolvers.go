/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta14 "github.com/upbound/provider-aws/apis/cloudwatchlogs/v1beta1"
	v1beta12 "github.com/upbound/provider-aws/apis/firehose/v1beta1"
	v1beta15 "github.com/upbound/provider-aws/apis/iam/v1beta1"
	v1beta11 "github.com/upbound/provider-aws/apis/kinesis/v1beta1"
	v1beta13 "github.com/upbound/provider-aws/apis/lambda/v1beta1"
	v1beta1 "github.com/upbound/provider-aws/apis/s3/v1beta1"
	common "github.com/upbound/provider-aws/config/common"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Application.
func (mg *Application) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.ApplicationConfiguration); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.ApplicationConfiguration[i3].ApplicationCodeConfiguration); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.ApplicationConfiguration[i3].ApplicationCodeConfiguration[i4].CodeContent); i5++ {
				for i6 := 0; i6 < len(mg.Spec.ForProvider.ApplicationConfiguration[i3].ApplicationCodeConfiguration[i4].CodeContent[i5].S3ContentLocation); i6++ {
					rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
						CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ApplicationConfiguration[i3].ApplicationCodeConfiguration[i4].CodeContent[i5].S3ContentLocation[i6].BucketArn),
						Extract:      common.ARNExtractor(),
						Reference:    mg.Spec.ForProvider.ApplicationConfiguration[i3].ApplicationCodeConfiguration[i4].CodeContent[i5].S3ContentLocation[i6].BucketArnRef,
						Selector:     mg.Spec.ForProvider.ApplicationConfiguration[i3].ApplicationCodeConfiguration[i4].CodeContent[i5].S3ContentLocation[i6].BucketArnSelector,
						To: reference.To{
							List:    &v1beta1.BucketList{},
							Managed: &v1beta1.Bucket{},
						},
					})
					if err != nil {
						return errors.Wrap(err, "mg.Spec.ForProvider.ApplicationConfiguration[i3].ApplicationCodeConfiguration[i4].CodeContent[i5].S3ContentLocation[i6].BucketArn")
					}
					mg.Spec.ForProvider.ApplicationConfiguration[i3].ApplicationCodeConfiguration[i4].CodeContent[i5].S3ContentLocation[i6].BucketArn = reference.ToPtrValue(rsp.ResolvedValue)
					mg.Spec.ForProvider.ApplicationConfiguration[i3].ApplicationCodeConfiguration[i4].CodeContent[i5].S3ContentLocation[i6].BucketArnRef = rsp.ResolvedReference

				}
			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.ApplicationConfiguration); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.ApplicationConfiguration[i3].ApplicationCodeConfiguration); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.ApplicationConfiguration[i3].ApplicationCodeConfiguration[i4].CodeContent); i5++ {
				for i6 := 0; i6 < len(mg.Spec.ForProvider.ApplicationConfiguration[i3].ApplicationCodeConfiguration[i4].CodeContent[i5].S3ContentLocation); i6++ {
					rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
						CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ApplicationConfiguration[i3].ApplicationCodeConfiguration[i4].CodeContent[i5].S3ContentLocation[i6].FileKey),
						Extract:      resource.ExtractParamPath("key", false),
						Reference:    mg.Spec.ForProvider.ApplicationConfiguration[i3].ApplicationCodeConfiguration[i4].CodeContent[i5].S3ContentLocation[i6].FileKeyRef,
						Selector:     mg.Spec.ForProvider.ApplicationConfiguration[i3].ApplicationCodeConfiguration[i4].CodeContent[i5].S3ContentLocation[i6].FileKeySelector,
						To: reference.To{
							List:    &v1beta1.ObjectList{},
							Managed: &v1beta1.Object{},
						},
					})
					if err != nil {
						return errors.Wrap(err, "mg.Spec.ForProvider.ApplicationConfiguration[i3].ApplicationCodeConfiguration[i4].CodeContent[i5].S3ContentLocation[i6].FileKey")
					}
					mg.Spec.ForProvider.ApplicationConfiguration[i3].ApplicationCodeConfiguration[i4].CodeContent[i5].S3ContentLocation[i6].FileKey = reference.ToPtrValue(rsp.ResolvedValue)
					mg.Spec.ForProvider.ApplicationConfiguration[i3].ApplicationCodeConfiguration[i4].CodeContent[i5].S3ContentLocation[i6].FileKeyRef = rsp.ResolvedReference

				}
			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.ApplicationConfiguration); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Input); i5++ {
				for i6 := 0; i6 < len(mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Input[i5].KinesisStreamsInput); i6++ {
					rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
						CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Input[i5].KinesisStreamsInput[i6].ResourceArn),
						Extract:      resource.ExtractParamPath("arn", false),
						Reference:    mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Input[i5].KinesisStreamsInput[i6].ResourceArnRef,
						Selector:     mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Input[i5].KinesisStreamsInput[i6].ResourceArnSelector,
						To: reference.To{
							List:    &v1beta11.StreamList{},
							Managed: &v1beta11.Stream{},
						},
					})
					if err != nil {
						return errors.Wrap(err, "mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Input[i5].KinesisStreamsInput[i6].ResourceArn")
					}
					mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Input[i5].KinesisStreamsInput[i6].ResourceArn = reference.ToPtrValue(rsp.ResolvedValue)
					mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Input[i5].KinesisStreamsInput[i6].ResourceArnRef = rsp.ResolvedReference

				}
			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.ApplicationConfiguration); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Output); i5++ {
				for i6 := 0; i6 < len(mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Output[i5].KinesisFirehoseOutput); i6++ {
					rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
						CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Output[i5].KinesisFirehoseOutput[i6].ResourceArn),
						Extract:      resource.ExtractParamPath("arn", false),
						Reference:    mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Output[i5].KinesisFirehoseOutput[i6].ResourceArnRef,
						Selector:     mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Output[i5].KinesisFirehoseOutput[i6].ResourceArnSelector,
						To: reference.To{
							List:    &v1beta12.DeliveryStreamList{},
							Managed: &v1beta12.DeliveryStream{},
						},
					})
					if err != nil {
						return errors.Wrap(err, "mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Output[i5].KinesisFirehoseOutput[i6].ResourceArn")
					}
					mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Output[i5].KinesisFirehoseOutput[i6].ResourceArn = reference.ToPtrValue(rsp.ResolvedValue)
					mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Output[i5].KinesisFirehoseOutput[i6].ResourceArnRef = rsp.ResolvedReference

				}
			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.ApplicationConfiguration); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Output); i5++ {
				for i6 := 0; i6 < len(mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Output[i5].LambdaOutput); i6++ {
					rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
						CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Output[i5].LambdaOutput[i6].ResourceArn),
						Extract:      resource.ExtractParamPath("arn", true),
						Reference:    mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Output[i5].LambdaOutput[i6].ResourceArnRef,
						Selector:     mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Output[i5].LambdaOutput[i6].ResourceArnSelector,
						To: reference.To{
							List:    &v1beta13.FunctionList{},
							Managed: &v1beta13.Function{},
						},
					})
					if err != nil {
						return errors.Wrap(err, "mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Output[i5].LambdaOutput[i6].ResourceArn")
					}
					mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Output[i5].LambdaOutput[i6].ResourceArn = reference.ToPtrValue(rsp.ResolvedValue)
					mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Output[i5].LambdaOutput[i6].ResourceArnRef = rsp.ResolvedReference

				}
			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.ApplicationConfiguration); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].ReferenceDataSource); i5++ {
				for i6 := 0; i6 < len(mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].ReferenceDataSource[i5].S3ReferenceDataSource); i6++ {
					rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
						CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].ReferenceDataSource[i5].S3ReferenceDataSource[i6].BucketArn),
						Extract:      resource.ExtractParamPath("arn", true),
						Reference:    mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].ReferenceDataSource[i5].S3ReferenceDataSource[i6].BucketArnRef,
						Selector:     mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].ReferenceDataSource[i5].S3ReferenceDataSource[i6].BucketArnSelector,
						To: reference.To{
							List:    &v1beta1.BucketList{},
							Managed: &v1beta1.Bucket{},
						},
					})
					if err != nil {
						return errors.Wrap(err, "mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].ReferenceDataSource[i5].S3ReferenceDataSource[i6].BucketArn")
					}
					mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].ReferenceDataSource[i5].S3ReferenceDataSource[i6].BucketArn = reference.ToPtrValue(rsp.ResolvedValue)
					mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].ReferenceDataSource[i5].S3ReferenceDataSource[i6].BucketArnRef = rsp.ResolvedReference

				}
			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.CloudwatchLoggingOptions); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CloudwatchLoggingOptions[i3].LogStreamArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.CloudwatchLoggingOptions[i3].LogStreamArnRef,
			Selector:     mg.Spec.ForProvider.CloudwatchLoggingOptions[i3].LogStreamArnSelector,
			To: reference.To{
				List:    &v1beta14.StreamList{},
				Managed: &v1beta14.Stream{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.CloudwatchLoggingOptions[i3].LogStreamArn")
		}
		mg.Spec.ForProvider.CloudwatchLoggingOptions[i3].LogStreamArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.CloudwatchLoggingOptions[i3].LogStreamArnRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServiceExecutionRole),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.ForProvider.ServiceExecutionRoleRef,
		Selector:     mg.Spec.ForProvider.ServiceExecutionRoleSelector,
		To: reference.To{
			List:    &v1beta15.RoleList{},
			Managed: &v1beta15.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ServiceExecutionRole")
	}
	mg.Spec.ForProvider.ServiceExecutionRole = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServiceExecutionRoleRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ApplicationSnapshot.
func (mg *ApplicationSnapshot) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ApplicationName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ApplicationNameRef,
		Selector:     mg.Spec.ForProvider.ApplicationNameSelector,
		To: reference.To{
			List:    &ApplicationList{},
			Managed: &Application{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ApplicationName")
	}
	mg.Spec.ForProvider.ApplicationName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ApplicationNameRef = rsp.ResolvedReference

	return nil
}
