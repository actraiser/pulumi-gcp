// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package spanner

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// An isolated set of Cloud Spanner resources on which databases can be
// hosted.
//
//
// To get more information about Instance, see:
//
// * [API documentation](https://cloud.google.com/spanner/docs/reference/rest/v1/projects.instances)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/spanner/)
type Instance struct {
	pulumi.CustomResourceState

	// The name of the instance's configuration (similar but not
	// quite the same as a region) which defines defines the geographic placement and
	// replication of your databases in this instance. It determines where your data
	// is stored. Values are typically of the form `regional-europe-west1` , `us-central` etc.
	// In order to obtain a valid list please consult the
	// [Configuration section of the docs](https://cloud.google.com/spanner/docs/instances).
	Config pulumi.StringOutput `pulumi:"config"`
	// The descriptive name for this instance as it appears in UIs. Must be
	// unique per project and between 4 and 30 characters in length.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// An object containing a list of "key": value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// A unique identifier for the instance, which cannot be changed after
	// the instance is created. The name must be between 6 and 30 characters
	// in length.
	Name pulumi.StringOutput `pulumi:"name"`
	// The number of nodes allocated to this instance.
	NumNodes pulumi.IntPtrOutput `pulumi:"numNodes"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Instance status: 'CREATING' or 'READY'.
	State pulumi.StringOutput `pulumi:"state"`
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil || args.Config == nil {
		return nil, errors.New("missing required argument 'Config'")
	}
	if args == nil || args.DisplayName == nil {
		return nil, errors.New("missing required argument 'DisplayName'")
	}
	if args == nil {
		args = &InstanceArgs{}
	}
	var resource Instance
	err := ctx.RegisterResource("gcp:spanner/instance:Instance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstance gets an existing Instance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceState, opts ...pulumi.ResourceOption) (*Instance, error) {
	var resource Instance
	err := ctx.ReadResource("gcp:spanner/instance:Instance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Instance resources.
type instanceState struct {
	// The name of the instance's configuration (similar but not
	// quite the same as a region) which defines defines the geographic placement and
	// replication of your databases in this instance. It determines where your data
	// is stored. Values are typically of the form `regional-europe-west1` , `us-central` etc.
	// In order to obtain a valid list please consult the
	// [Configuration section of the docs](https://cloud.google.com/spanner/docs/instances).
	Config *string `pulumi:"config"`
	// The descriptive name for this instance as it appears in UIs. Must be
	// unique per project and between 4 and 30 characters in length.
	DisplayName *string `pulumi:"displayName"`
	// An object containing a list of "key": value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	Labels map[string]string `pulumi:"labels"`
	// A unique identifier for the instance, which cannot be changed after
	// the instance is created. The name must be between 6 and 30 characters
	// in length.
	Name *string `pulumi:"name"`
	// The number of nodes allocated to this instance.
	NumNodes *int `pulumi:"numNodes"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Instance status: 'CREATING' or 'READY'.
	State *string `pulumi:"state"`
}

type InstanceState struct {
	// The name of the instance's configuration (similar but not
	// quite the same as a region) which defines defines the geographic placement and
	// replication of your databases in this instance. It determines where your data
	// is stored. Values are typically of the form `regional-europe-west1` , `us-central` etc.
	// In order to obtain a valid list please consult the
	// [Configuration section of the docs](https://cloud.google.com/spanner/docs/instances).
	Config pulumi.StringPtrInput
	// The descriptive name for this instance as it appears in UIs. Must be
	// unique per project and between 4 and 30 characters in length.
	DisplayName pulumi.StringPtrInput
	// An object containing a list of "key": value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	Labels pulumi.StringMapInput
	// A unique identifier for the instance, which cannot be changed after
	// the instance is created. The name must be between 6 and 30 characters
	// in length.
	Name pulumi.StringPtrInput
	// The number of nodes allocated to this instance.
	NumNodes pulumi.IntPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Instance status: 'CREATING' or 'READY'.
	State pulumi.StringPtrInput
}

func (InstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceState)(nil)).Elem()
}

type instanceArgs struct {
	// The name of the instance's configuration (similar but not
	// quite the same as a region) which defines defines the geographic placement and
	// replication of your databases in this instance. It determines where your data
	// is stored. Values are typically of the form `regional-europe-west1` , `us-central` etc.
	// In order to obtain a valid list please consult the
	// [Configuration section of the docs](https://cloud.google.com/spanner/docs/instances).
	Config string `pulumi:"config"`
	// The descriptive name for this instance as it appears in UIs. Must be
	// unique per project and between 4 and 30 characters in length.
	DisplayName string `pulumi:"displayName"`
	// An object containing a list of "key": value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	Labels map[string]string `pulumi:"labels"`
	// A unique identifier for the instance, which cannot be changed after
	// the instance is created. The name must be between 6 and 30 characters
	// in length.
	Name *string `pulumi:"name"`
	// The number of nodes allocated to this instance.
	NumNodes *int `pulumi:"numNodes"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// The name of the instance's configuration (similar but not
	// quite the same as a region) which defines defines the geographic placement and
	// replication of your databases in this instance. It determines where your data
	// is stored. Values are typically of the form `regional-europe-west1` , `us-central` etc.
	// In order to obtain a valid list please consult the
	// [Configuration section of the docs](https://cloud.google.com/spanner/docs/instances).
	Config pulumi.StringInput
	// The descriptive name for this instance as it appears in UIs. Must be
	// unique per project and between 4 and 30 characters in length.
	DisplayName pulumi.StringInput
	// An object containing a list of "key": value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	Labels pulumi.StringMapInput
	// A unique identifier for the instance, which cannot be changed after
	// the instance is created. The name must be between 6 and 30 characters
	// in length.
	Name pulumi.StringPtrInput
	// The number of nodes allocated to this instance.
	NumNodes pulumi.IntPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (InstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceArgs)(nil)).Elem()
}
