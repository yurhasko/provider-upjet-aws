/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta1 "github.com/upbound/provider-aws/apis/cloudwatchlogs/v1beta1"
	v1beta13 "github.com/upbound/provider-aws/apis/firehose/v1beta1"
	v1beta11 "github.com/upbound/provider-aws/apis/iam/v1beta1"
	v1beta12 "github.com/upbound/provider-aws/apis/kinesis/v1beta1"
	common "github.com/upbound/provider-aws/config/common"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Application.
func (mg *Application) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.CloudwatchLoggingOptions); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CloudwatchLoggingOptions[i3].LogStreamArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.CloudwatchLoggingOptions[i3].LogStreamArnRef,
			Selector:     mg.Spec.ForProvider.CloudwatchLoggingOptions[i3].LogStreamArnSelector,
			To: reference.To{
				List:    &v1beta1.StreamList{},
				Managed: &v1beta1.Stream{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.CloudwatchLoggingOptions[i3].LogStreamArn")
		}
		mg.Spec.ForProvider.CloudwatchLoggingOptions[i3].LogStreamArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.CloudwatchLoggingOptions[i3].LogStreamArnRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.CloudwatchLoggingOptions); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CloudwatchLoggingOptions[i3].RoleArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.CloudwatchLoggingOptions[i3].RoleArnRef,
			Selector:     mg.Spec.ForProvider.CloudwatchLoggingOptions[i3].RoleArnSelector,
			To: reference.To{
				List:    &v1beta11.RoleList{},
				Managed: &v1beta11.Role{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.CloudwatchLoggingOptions[i3].RoleArn")
		}
		mg.Spec.ForProvider.CloudwatchLoggingOptions[i3].RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.CloudwatchLoggingOptions[i3].RoleArnRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Inputs); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Inputs[i3].KinesisStream); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Inputs[i3].KinesisStream[i4].ResourceArn),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.ForProvider.Inputs[i3].KinesisStream[i4].ResourceArnRef,
				Selector:     mg.Spec.ForProvider.Inputs[i3].KinesisStream[i4].ResourceArnSelector,
				To: reference.To{
					List:    &v1beta12.StreamList{},
					Managed: &v1beta12.Stream{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Inputs[i3].KinesisStream[i4].ResourceArn")
			}
			mg.Spec.ForProvider.Inputs[i3].KinesisStream[i4].ResourceArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Inputs[i3].KinesisStream[i4].ResourceArnRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Inputs); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Inputs[i3].KinesisStream); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Inputs[i3].KinesisStream[i4].RoleArn),
				Extract:      common.ARNExtractor(),
				Reference:    mg.Spec.ForProvider.Inputs[i3].KinesisStream[i4].RoleArnRef,
				Selector:     mg.Spec.ForProvider.Inputs[i3].KinesisStream[i4].RoleArnSelector,
				To: reference.To{
					List:    &v1beta11.RoleList{},
					Managed: &v1beta11.Role{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Inputs[i3].KinesisStream[i4].RoleArn")
			}
			mg.Spec.ForProvider.Inputs[i3].KinesisStream[i4].RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Inputs[i3].KinesisStream[i4].RoleArnRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Outputs); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Outputs[i3].KinesisFirehose); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Outputs[i3].KinesisFirehose[i4].ResourceArn),
				Extract:      resource.ExtractParamPath("arn", false),
				Reference:    mg.Spec.ForProvider.Outputs[i3].KinesisFirehose[i4].ResourceArnRef,
				Selector:     mg.Spec.ForProvider.Outputs[i3].KinesisFirehose[i4].ResourceArnSelector,
				To: reference.To{
					List:    &v1beta13.DeliveryStreamList{},
					Managed: &v1beta13.DeliveryStream{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Outputs[i3].KinesisFirehose[i4].ResourceArn")
			}
			mg.Spec.ForProvider.Outputs[i3].KinesisFirehose[i4].ResourceArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Outputs[i3].KinesisFirehose[i4].ResourceArnRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Outputs); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Outputs[i3].KinesisFirehose); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Outputs[i3].KinesisFirehose[i4].RoleArn),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.ForProvider.Outputs[i3].KinesisFirehose[i4].RoleArnRef,
				Selector:     mg.Spec.ForProvider.Outputs[i3].KinesisFirehose[i4].RoleArnSelector,
				To: reference.To{
					List:    &v1beta11.RoleList{},
					Managed: &v1beta11.Role{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Outputs[i3].KinesisFirehose[i4].RoleArn")
			}
			mg.Spec.ForProvider.Outputs[i3].KinesisFirehose[i4].RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Outputs[i3].KinesisFirehose[i4].RoleArnRef = rsp.ResolvedReference

		}
	}

	return nil
}
