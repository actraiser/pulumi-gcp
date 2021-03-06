// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dataproc

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Three different resources help you manage IAM policies on dataproc jobs. Each of these resources serves a different use case:
//
// * `dataproc.JobIAMPolicy`: Authoritative. Sets the IAM policy for the job and replaces any existing policy already attached.
// * `dataproc.JobIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the job are preserved.
// * `dataproc.JobIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the job are preserved.
//
// > **Note:** `dataproc.JobIAMPolicy` **cannot** be used in conjunction with `dataproc.JobIAMBinding` and `dataproc.JobIAMMember` or they will fight over what your policy should be. In addition, be careful not to accidentally unset ownership of the job as `dataproc.JobIAMPolicy` replaces the entire policy.
//
// > **Note:** `dataproc.JobIAMBinding` resources **can be** used in conjunction with `dataproc.JobIAMMember` resources **only if** they do not grant privilege to the same role.
type JobIAMMember struct {
	pulumi.CustomResourceState

	Condition JobIAMMemberConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the jobs's IAM policy.
	Etag   pulumi.StringOutput `pulumi:"etag"`
	JobId  pulumi.StringOutput `pulumi:"jobId"`
	Member pulumi.StringOutput `pulumi:"member"`
	// The project in which the job belongs. If it
	// is not provided, the provider will use a default.
	Project pulumi.StringOutput `pulumi:"project"`
	// The region in which the job belongs. If it
	// is not provided, the provider will use a default.
	Region pulumi.StringOutput `pulumi:"region"`
	// The role that should be applied. Only one
	// `dataproc.JobIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewJobIAMMember registers a new resource with the given unique name, arguments, and options.
func NewJobIAMMember(ctx *pulumi.Context,
	name string, args *JobIAMMemberArgs, opts ...pulumi.ResourceOption) (*JobIAMMember, error) {
	if args == nil || args.JobId == nil {
		return nil, errors.New("missing required argument 'JobId'")
	}
	if args == nil || args.Member == nil {
		return nil, errors.New("missing required argument 'Member'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil {
		args = &JobIAMMemberArgs{}
	}
	var resource JobIAMMember
	err := ctx.RegisterResource("gcp:dataproc/jobIAMMember:JobIAMMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetJobIAMMember gets an existing JobIAMMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJobIAMMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobIAMMemberState, opts ...pulumi.ResourceOption) (*JobIAMMember, error) {
	var resource JobIAMMember
	err := ctx.ReadResource("gcp:dataproc/jobIAMMember:JobIAMMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering JobIAMMember resources.
type jobIAMMemberState struct {
	Condition *JobIAMMemberCondition `pulumi:"condition"`
	// (Computed) The etag of the jobs's IAM policy.
	Etag   *string `pulumi:"etag"`
	JobId  *string `pulumi:"jobId"`
	Member *string `pulumi:"member"`
	// The project in which the job belongs. If it
	// is not provided, the provider will use a default.
	Project *string `pulumi:"project"`
	// The region in which the job belongs. If it
	// is not provided, the provider will use a default.
	Region *string `pulumi:"region"`
	// The role that should be applied. Only one
	// `dataproc.JobIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type JobIAMMemberState struct {
	Condition JobIAMMemberConditionPtrInput
	// (Computed) The etag of the jobs's IAM policy.
	Etag   pulumi.StringPtrInput
	JobId  pulumi.StringPtrInput
	Member pulumi.StringPtrInput
	// The project in which the job belongs. If it
	// is not provided, the provider will use a default.
	Project pulumi.StringPtrInput
	// The region in which the job belongs. If it
	// is not provided, the provider will use a default.
	Region pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `dataproc.JobIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (JobIAMMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*jobIAMMemberState)(nil)).Elem()
}

type jobIAMMemberArgs struct {
	Condition *JobIAMMemberCondition `pulumi:"condition"`
	JobId     string                 `pulumi:"jobId"`
	Member    string                 `pulumi:"member"`
	// The project in which the job belongs. If it
	// is not provided, the provider will use a default.
	Project *string `pulumi:"project"`
	// The region in which the job belongs. If it
	// is not provided, the provider will use a default.
	Region *string `pulumi:"region"`
	// The role that should be applied. Only one
	// `dataproc.JobIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a JobIAMMember resource.
type JobIAMMemberArgs struct {
	Condition JobIAMMemberConditionPtrInput
	JobId     pulumi.StringInput
	Member    pulumi.StringInput
	// The project in which the job belongs. If it
	// is not provided, the provider will use a default.
	Project pulumi.StringPtrInput
	// The region in which the job belongs. If it
	// is not provided, the provider will use a default.
	Region pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `dataproc.JobIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (JobIAMMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jobIAMMemberArgs)(nil)).Elem()
}
