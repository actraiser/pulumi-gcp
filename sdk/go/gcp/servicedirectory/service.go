// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package servicedirectory

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// An individual service. A service contains a name and optional metadata.
//
// To get more information about Service, see:
//
// * [API documentation](https://cloud.google.com/service-directory/docs/reference/rest/v1beta1/projects.locations.namespaces.services)
// * How-to Guides
//     * [Configuring a service](https://cloud.google.com/service-directory/docs/configuring-service-directory#configuring_a_service)
type Service struct {
	pulumi.CustomResourceState

	// Metadata for the service. This data can be consumed
	// by service clients. The entire metadata dictionary may contain
	// up to 2000 characters, spread across all key-value pairs.
	// Metadata that goes beyond any these limits will be rejected.
	Metadata pulumi.StringMapOutput `pulumi:"metadata"`
	// The resource name for the service in the format 'projects/*/locations/*/namespaces/*/services/*'.
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource name of the namespace this service will belong to.
	Namespace pulumi.StringOutput `pulumi:"namespace"`
	// The Resource ID must be 1-63 characters long, including digits,
	// lowercase letters or the hyphen character.
	ServiceId pulumi.StringOutput `pulumi:"serviceId"`
}

// NewService registers a new resource with the given unique name, arguments, and options.
func NewService(ctx *pulumi.Context,
	name string, args *ServiceArgs, opts ...pulumi.ResourceOption) (*Service, error) {
	if args == nil || args.Namespace == nil {
		return nil, errors.New("missing required argument 'Namespace'")
	}
	if args == nil || args.ServiceId == nil {
		return nil, errors.New("missing required argument 'ServiceId'")
	}
	if args == nil {
		args = &ServiceArgs{}
	}
	var resource Service
	err := ctx.RegisterResource("gcp:servicedirectory/service:Service", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetService gets an existing Service resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceState, opts ...pulumi.ResourceOption) (*Service, error) {
	var resource Service
	err := ctx.ReadResource("gcp:servicedirectory/service:Service", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Service resources.
type serviceState struct {
	// Metadata for the service. This data can be consumed
	// by service clients. The entire metadata dictionary may contain
	// up to 2000 characters, spread across all key-value pairs.
	// Metadata that goes beyond any these limits will be rejected.
	Metadata map[string]string `pulumi:"metadata"`
	// The resource name for the service in the format 'projects/*/locations/*/namespaces/*/services/*'.
	Name *string `pulumi:"name"`
	// The resource name of the namespace this service will belong to.
	Namespace *string `pulumi:"namespace"`
	// The Resource ID must be 1-63 characters long, including digits,
	// lowercase letters or the hyphen character.
	ServiceId *string `pulumi:"serviceId"`
}

type ServiceState struct {
	// Metadata for the service. This data can be consumed
	// by service clients. The entire metadata dictionary may contain
	// up to 2000 characters, spread across all key-value pairs.
	// Metadata that goes beyond any these limits will be rejected.
	Metadata pulumi.StringMapInput
	// The resource name for the service in the format 'projects/*/locations/*/namespaces/*/services/*'.
	Name pulumi.StringPtrInput
	// The resource name of the namespace this service will belong to.
	Namespace pulumi.StringPtrInput
	// The Resource ID must be 1-63 characters long, including digits,
	// lowercase letters or the hyphen character.
	ServiceId pulumi.StringPtrInput
}

func (ServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceState)(nil)).Elem()
}

type serviceArgs struct {
	// Metadata for the service. This data can be consumed
	// by service clients. The entire metadata dictionary may contain
	// up to 2000 characters, spread across all key-value pairs.
	// Metadata that goes beyond any these limits will be rejected.
	Metadata map[string]string `pulumi:"metadata"`
	// The resource name of the namespace this service will belong to.
	Namespace string `pulumi:"namespace"`
	// The Resource ID must be 1-63 characters long, including digits,
	// lowercase letters or the hyphen character.
	ServiceId string `pulumi:"serviceId"`
}

// The set of arguments for constructing a Service resource.
type ServiceArgs struct {
	// Metadata for the service. This data can be consumed
	// by service clients. The entire metadata dictionary may contain
	// up to 2000 characters, spread across all key-value pairs.
	// Metadata that goes beyond any these limits will be rejected.
	Metadata pulumi.StringMapInput
	// The resource name of the namespace this service will belong to.
	Namespace pulumi.StringInput
	// The Resource ID must be 1-63 characters long, including digits,
	// lowercase letters or the hyphen character.
	ServiceId pulumi.StringInput
}

func (ServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceArgs)(nil)).Elem()
}
