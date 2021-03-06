// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package firebase

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Sets the default Google Cloud Platform (GCP) resource location for the specified FirebaseProject.
//
// This method creates an App Engine application with a default Cloud Storage bucket, located in the specified
// locationId. This location must be one of the available GCP resource locations.
//
// After the default GCP resource location is finalized, or if it was already set, it cannot be changed.
// The default GCP resource location for the specified FirebaseProject might already be set because either the
// GCP Project already has an App Engine application or defaultLocation.finalize was previously called with a
// specified locationId. Any new calls to defaultLocation.finalize with a different specified locationId will
// return a 409 error.
//
// To get more information about ProjectLocation, see:
//
// * [API documentation](https://firebase.google.com/docs/projects/api/reference/rest/v1beta1/projects.defaultLocation/finalize)
// * How-to Guides
//     * [Official Documentation](https://firebase.google.com/)
type ProjectLocation struct {
	pulumi.CustomResourceState

	// The ID of the default GCP resource location for the Project. The location must be one of the available GCP
	// resource locations.
	LocationId pulumi.StringOutput `pulumi:"locationId"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewProjectLocation registers a new resource with the given unique name, arguments, and options.
func NewProjectLocation(ctx *pulumi.Context,
	name string, args *ProjectLocationArgs, opts ...pulumi.ResourceOption) (*ProjectLocation, error) {
	if args == nil || args.LocationId == nil {
		return nil, errors.New("missing required argument 'LocationId'")
	}
	if args == nil {
		args = &ProjectLocationArgs{}
	}
	var resource ProjectLocation
	err := ctx.RegisterResource("gcp:firebase/projectLocation:ProjectLocation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectLocation gets an existing ProjectLocation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectLocation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectLocationState, opts ...pulumi.ResourceOption) (*ProjectLocation, error) {
	var resource ProjectLocation
	err := ctx.ReadResource("gcp:firebase/projectLocation:ProjectLocation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectLocation resources.
type projectLocationState struct {
	// The ID of the default GCP resource location for the Project. The location must be one of the available GCP
	// resource locations.
	LocationId *string `pulumi:"locationId"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

type ProjectLocationState struct {
	// The ID of the default GCP resource location for the Project. The location must be one of the available GCP
	// resource locations.
	LocationId pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (ProjectLocationState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectLocationState)(nil)).Elem()
}

type projectLocationArgs struct {
	// The ID of the default GCP resource location for the Project. The location must be one of the available GCP
	// resource locations.
	LocationId string `pulumi:"locationId"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a ProjectLocation resource.
type ProjectLocationArgs struct {
	// The ID of the default GCP resource location for the Project. The location must be one of the available GCP
	// resource locations.
	LocationId pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (ProjectLocationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectLocationArgs)(nil)).Elem()
}
