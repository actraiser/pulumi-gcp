// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package dataproc

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Three different resources help you manage IAM policies on dataproc clusters. Each of these resources serves a different use case:
//
// * `dataproc.ClusterIAMPolicy`: Authoritative. Sets the IAM policy for the cluster and replaces any existing policy already attached.
// * `dataproc.ClusterIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the cluster are preserved.
// * `dataproc.ClusterIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the cluster are preserved.
//
// > **Note:** `dataproc.ClusterIAMPolicy` **cannot** be used in conjunction with `dataproc.ClusterIAMBinding` and `dataproc.ClusterIAMMember` or they will fight over what your policy should be. In addition, be careful not to accidentally unset ownership of the cluster as `dataproc.ClusterIAMPolicy` replaces the entire policy.
//
// > **Note:** `dataproc.ClusterIAMBinding` resources **can be** used in conjunction with `dataproc.ClusterIAMMember` resources **only if** they do not grant privilege to the same role.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/dataproc_cluster_iam.html.markdown.
type ClusterIAMMember struct {
	pulumi.CustomResourceState

	// The name or relative resource id of the cluster to manage IAM policies for.
	Cluster pulumi.StringOutput `pulumi:"cluster"`
	Condition ClusterIAMMemberConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the clusters's IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	Member pulumi.StringOutput `pulumi:"member"`
	// The project in which the cluster belongs. If it
	// is not provided, this provider will use the provider default.
	Project pulumi.StringOutput `pulumi:"project"`
	// The region in which the cluster belongs. If it
	// is not provided, this provider will use the provider default.
	Region pulumi.StringOutput `pulumi:"region"`
	// The role that should be applied. Only one
	// `dataproc.ClusterIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewClusterIAMMember registers a new resource with the given unique name, arguments, and options.
func NewClusterIAMMember(ctx *pulumi.Context,
	name string, args *ClusterIAMMemberArgs, opts ...pulumi.ResourceOption) (*ClusterIAMMember, error) {
	if args == nil || args.Cluster == nil {
		return nil, errors.New("missing required argument 'Cluster'")
	}
	if args == nil || args.Member == nil {
		return nil, errors.New("missing required argument 'Member'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil {
		args = &ClusterIAMMemberArgs{}
	}
	var resource ClusterIAMMember
	err := ctx.RegisterResource("gcp:dataproc/clusterIAMMember:ClusterIAMMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClusterIAMMember gets an existing ClusterIAMMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterIAMMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterIAMMemberState, opts ...pulumi.ResourceOption) (*ClusterIAMMember, error) {
	var resource ClusterIAMMember
	err := ctx.ReadResource("gcp:dataproc/clusterIAMMember:ClusterIAMMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClusterIAMMember resources.
type clusterIAMMemberState struct {
	// The name or relative resource id of the cluster to manage IAM policies for.
	Cluster *string `pulumi:"cluster"`
	Condition *ClusterIAMMemberCondition `pulumi:"condition"`
	// (Computed) The etag of the clusters's IAM policy.
	Etag *string `pulumi:"etag"`
	Member *string `pulumi:"member"`
	// The project in which the cluster belongs. If it
	// is not provided, this provider will use the provider default.
	Project *string `pulumi:"project"`
	// The region in which the cluster belongs. If it
	// is not provided, this provider will use the provider default.
	Region *string `pulumi:"region"`
	// The role that should be applied. Only one
	// `dataproc.ClusterIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type ClusterIAMMemberState struct {
	// The name or relative resource id of the cluster to manage IAM policies for.
	Cluster pulumi.StringPtrInput
	Condition ClusterIAMMemberConditionPtrInput
	// (Computed) The etag of the clusters's IAM policy.
	Etag pulumi.StringPtrInput
	Member pulumi.StringPtrInput
	// The project in which the cluster belongs. If it
	// is not provided, this provider will use the provider default.
	Project pulumi.StringPtrInput
	// The region in which the cluster belongs. If it
	// is not provided, this provider will use the provider default.
	Region pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `dataproc.ClusterIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (ClusterIAMMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterIAMMemberState)(nil)).Elem()
}

type clusterIAMMemberArgs struct {
	// The name or relative resource id of the cluster to manage IAM policies for.
	Cluster string `pulumi:"cluster"`
	Condition *ClusterIAMMemberCondition `pulumi:"condition"`
	Member string `pulumi:"member"`
	// The project in which the cluster belongs. If it
	// is not provided, this provider will use the provider default.
	Project *string `pulumi:"project"`
	// The region in which the cluster belongs. If it
	// is not provided, this provider will use the provider default.
	Region *string `pulumi:"region"`
	// The role that should be applied. Only one
	// `dataproc.ClusterIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a ClusterIAMMember resource.
type ClusterIAMMemberArgs struct {
	// The name or relative resource id of the cluster to manage IAM policies for.
	Cluster pulumi.StringInput
	Condition ClusterIAMMemberConditionPtrInput
	Member pulumi.StringInput
	// The project in which the cluster belongs. If it
	// is not provided, this provider will use the provider default.
	Project pulumi.StringPtrInput
	// The region in which the cluster belongs. If it
	// is not provided, this provider will use the provider default.
	Region pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `dataproc.ClusterIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (ClusterIAMMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterIAMMemberArgs)(nil)).Elem()
}

