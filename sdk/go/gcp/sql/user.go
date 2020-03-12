// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package sql

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Creates a new Google SQL User on a Google SQL User Instance. For more information, see the [official documentation](https://cloud.google.com/sql/), or the [JSON API](https://cloud.google.com/sql/docs/admin-api/v1beta4/users).
//
// > **Note:** All arguments including the username and password will be stored in the raw state as plain-text.
// [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html). Passwords will not be retrieved when running import.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/sql_user.html.markdown.
type User struct {
	pulumi.CustomResourceState

	// The host the user can connect from. This is only supported
	// for MySQL instances. Don't set this field for PostgreSQL instances.
	// Can be an IP address. Changing this forces a new resource to be created.
	Host pulumi.StringPtrOutput `pulumi:"host"`
	// The name of the Cloud SQL instance. Changing this
	// forces a new resource to be created.
	Instance pulumi.StringOutput `pulumi:"instance"`
	// The name of the user. Changing this forces a new resource
	// to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The password for the user. Can be updated. For Postgres
	// instances this is a Required field.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewUser registers a new resource with the given unique name, arguments, and options.
func NewUser(ctx *pulumi.Context,
	name string, args *UserArgs, opts ...pulumi.ResourceOption) (*User, error) {
	if args == nil || args.Instance == nil {
		return nil, errors.New("missing required argument 'Instance'")
	}
	if args == nil {
		args = &UserArgs{}
	}
	var resource User
	err := ctx.RegisterResource("gcp:sql/user:User", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUser gets an existing User resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserState, opts ...pulumi.ResourceOption) (*User, error) {
	var resource User
	err := ctx.ReadResource("gcp:sql/user:User", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering User resources.
type userState struct {
	// The host the user can connect from. This is only supported
	// for MySQL instances. Don't set this field for PostgreSQL instances.
	// Can be an IP address. Changing this forces a new resource to be created.
	Host *string `pulumi:"host"`
	// The name of the Cloud SQL instance. Changing this
	// forces a new resource to be created.
	Instance *string `pulumi:"instance"`
	// The name of the user. Changing this forces a new resource
	// to be created.
	Name *string `pulumi:"name"`
	// The password for the user. Can be updated. For Postgres
	// instances this is a Required field.
	Password *string `pulumi:"password"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

type UserState struct {
	// The host the user can connect from. This is only supported
	// for MySQL instances. Don't set this field for PostgreSQL instances.
	// Can be an IP address. Changing this forces a new resource to be created.
	Host pulumi.StringPtrInput
	// The name of the Cloud SQL instance. Changing this
	// forces a new resource to be created.
	Instance pulumi.StringPtrInput
	// The name of the user. Changing this forces a new resource
	// to be created.
	Name pulumi.StringPtrInput
	// The password for the user. Can be updated. For Postgres
	// instances this is a Required field.
	Password pulumi.StringPtrInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (UserState) ElementType() reflect.Type {
	return reflect.TypeOf((*userState)(nil)).Elem()
}

type userArgs struct {
	// The host the user can connect from. This is only supported
	// for MySQL instances. Don't set this field for PostgreSQL instances.
	// Can be an IP address. Changing this forces a new resource to be created.
	Host *string `pulumi:"host"`
	// The name of the Cloud SQL instance. Changing this
	// forces a new resource to be created.
	Instance string `pulumi:"instance"`
	// The name of the user. Changing this forces a new resource
	// to be created.
	Name *string `pulumi:"name"`
	// The password for the user. Can be updated. For Postgres
	// instances this is a Required field.
	Password *string `pulumi:"password"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a User resource.
type UserArgs struct {
	// The host the user can connect from. This is only supported
	// for MySQL instances. Don't set this field for PostgreSQL instances.
	// Can be an IP address. Changing this forces a new resource to be created.
	Host pulumi.StringPtrInput
	// The name of the Cloud SQL instance. Changing this
	// forces a new resource to be created.
	Instance pulumi.StringInput
	// The name of the user. Changing this forces a new resource
	// to be created.
	Name pulumi.StringPtrInput
	// The password for the user. Can be updated. For Postgres
	// instances this is a Required field.
	Password pulumi.StringPtrInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (UserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userArgs)(nil)).Elem()
}

