// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package xyz

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-xyz/sdk/go/xyz/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetDataSource(ctx *pulumi.Context, args *GetDataSourceArgs, opts ...pulumi.InvokeOption) (*GetDataSourceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetDataSourceResult
	err := ctx.Invoke("xyz:index/getDataSource:getDataSource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDataSource.
type GetDataSourceArgs struct {
	SampleAttribute string `pulumi:"sampleAttribute"`
}

// A collection of values returned by getDataSource.
type GetDataSourceResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id              string `pulumi:"id"`
	SampleAttribute string `pulumi:"sampleAttribute"`
}

func GetDataSourceOutput(ctx *pulumi.Context, args GetDataSourceOutputArgs, opts ...pulumi.InvokeOption) GetDataSourceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDataSourceResultOutput, error) {
			args := v.(GetDataSourceArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv GetDataSourceResult
			secret, err := ctx.InvokePackageRaw("xyz:index/getDataSource:getDataSource", args, &rv, "", opts...)
			if err != nil {
				return GetDataSourceResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(GetDataSourceResultOutput)
			if secret {
				return pulumi.ToSecret(output).(GetDataSourceResultOutput), nil
			}
			return output, nil
		}).(GetDataSourceResultOutput)
}

// A collection of arguments for invoking getDataSource.
type GetDataSourceOutputArgs struct {
	SampleAttribute pulumi.StringInput `pulumi:"sampleAttribute"`
}

func (GetDataSourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDataSourceArgs)(nil)).Elem()
}

// A collection of values returned by getDataSource.
type GetDataSourceResultOutput struct{ *pulumi.OutputState }

func (GetDataSourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDataSourceResult)(nil)).Elem()
}

func (o GetDataSourceResultOutput) ToGetDataSourceResultOutput() GetDataSourceResultOutput {
	return o
}

func (o GetDataSourceResultOutput) ToGetDataSourceResultOutputWithContext(ctx context.Context) GetDataSourceResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetDataSourceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDataSourceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetDataSourceResultOutput) SampleAttribute() pulumi.StringOutput {
	return o.ApplyT(func(v GetDataSourceResult) string { return v.SampleAttribute }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDataSourceResultOutput{})
}