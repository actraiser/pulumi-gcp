// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kms

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides access to Google Cloud Platform KMS KeyRing. For more information see
// [the official documentation](https://cloud.google.com/kms/docs/object-hierarchy#key_ring)
// and
// [API](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings).
//
// A KeyRing is a grouping of CryptoKeys for organizational purposes. A KeyRing belongs to a Google Cloud Platform Project
// and resides in a specific location.
func GetKMSKeyRing(ctx *pulumi.Context, args *GetKMSKeyRingArgs, opts ...pulumi.InvokeOption) (*GetKMSKeyRingResult, error) {
	var rv GetKMSKeyRingResult
	err := ctx.Invoke("gcp:kms/getKMSKeyRing:getKMSKeyRing", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getKMSKeyRing.
type GetKMSKeyRingArgs struct {
	// The Google Cloud Platform location for the KeyRing.
	// A full list of valid locations can be found by running `gcloud kms locations list`.
	Location string `pulumi:"location"`
	// The KeyRing's name.
	// A KeyRing name must exist within the provided location and match the regular expression `[a-zA-Z0-9_-]{1,63}`
	Name string `pulumi:"name"`
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// A collection of values returned by getKMSKeyRing.
type GetKMSKeyRingResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id       string  `pulumi:"id"`
	Location string  `pulumi:"location"`
	Name     string  `pulumi:"name"`
	Project  *string `pulumi:"project"`
	// The self link of the created KeyRing. Its format is `projects/{projectId}/locations/{location}/keyRings/{keyRingName}`.
	SelfLink string `pulumi:"selfLink"`
}
