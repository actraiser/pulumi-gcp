// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package spanner

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Three different resources help you manage your IAM policy for a Spanner database. Each of these resources serves a different use case:
//
// * `spanner.DatabaseIAMPolicy`: Authoritative. Sets the IAM policy for the database and replaces any existing policy already attached.
//
// > **Warning:** It's entirely possibly to lock yourself out of your database using `spanner.DatabaseIAMPolicy`. Any permissions granted by default will be removed unless you include them in your config.
//
// * `spanner.DatabaseIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the database are preserved.
// * `spanner.DatabaseIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the database are preserved.
//
// > **Note:** `spanner.DatabaseIAMPolicy` **cannot** be used in conjunction with `spanner.DatabaseIAMBinding` and `spanner.DatabaseIAMMember` or they will fight over what your policy should be.
//
// > **Note:** `spanner.DatabaseIAMBinding` resources **can be** used in conjunction with `spanner.DatabaseIAMMember` resources **only if** they do not grant privilege to the same role.
type DatabaseIAMPolicy struct {
	pulumi.CustomResourceState

	// The name of the Spanner database.
	Database pulumi.StringOutput `pulumi:"database"`
	// (Computed) The etag of the database's IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The name of the Spanner instance the database belongs to.
	Instance pulumi.StringOutput `pulumi:"instance"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewDatabaseIAMPolicy registers a new resource with the given unique name, arguments, and options.
func NewDatabaseIAMPolicy(ctx *pulumi.Context,
	name string, args *DatabaseIAMPolicyArgs, opts ...pulumi.ResourceOption) (*DatabaseIAMPolicy, error) {
	if args == nil || args.Database == nil {
		return nil, errors.New("missing required argument 'Database'")
	}
	if args == nil || args.Instance == nil {
		return nil, errors.New("missing required argument 'Instance'")
	}
	if args == nil || args.PolicyData == nil {
		return nil, errors.New("missing required argument 'PolicyData'")
	}
	if args == nil {
		args = &DatabaseIAMPolicyArgs{}
	}
	var resource DatabaseIAMPolicy
	err := ctx.RegisterResource("gcp:spanner/databaseIAMPolicy:DatabaseIAMPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabaseIAMPolicy gets an existing DatabaseIAMPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabaseIAMPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseIAMPolicyState, opts ...pulumi.ResourceOption) (*DatabaseIAMPolicy, error) {
	var resource DatabaseIAMPolicy
	err := ctx.ReadResource("gcp:spanner/databaseIAMPolicy:DatabaseIAMPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatabaseIAMPolicy resources.
type databaseIAMPolicyState struct {
	// The name of the Spanner database.
	Database *string `pulumi:"database"`
	// (Computed) The etag of the database's IAM policy.
	Etag *string `pulumi:"etag"`
	// The name of the Spanner instance the database belongs to.
	Instance *string `pulumi:"instance"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

type DatabaseIAMPolicyState struct {
	// The name of the Spanner database.
	Database pulumi.StringPtrInput
	// (Computed) The etag of the database's IAM policy.
	Etag pulumi.StringPtrInput
	// The name of the Spanner instance the database belongs to.
	Instance pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (DatabaseIAMPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseIAMPolicyState)(nil)).Elem()
}

type databaseIAMPolicyArgs struct {
	// The name of the Spanner database.
	Database string `pulumi:"database"`
	// The name of the Spanner instance the database belongs to.
	Instance string `pulumi:"instance"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a DatabaseIAMPolicy resource.
type DatabaseIAMPolicyArgs struct {
	// The name of the Spanner database.
	Database pulumi.StringInput
	// The name of the Spanner instance the database belongs to.
	Instance pulumi.StringInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (DatabaseIAMPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseIAMPolicyArgs)(nil)).Elem()
}
