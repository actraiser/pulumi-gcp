// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package compute

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Get a subnetwork within GCE from its name and region.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/datasource_compute_subnetwork.html.markdown.
func LookupSubnetwork(ctx *pulumi.Context, args *LookupSubnetworkArgs, opts ...pulumi.InvokeOption) (*LookupSubnetworkResult, error) {
	var rv LookupSubnetworkResult
	err := ctx.Invoke("gcp:compute/getSubnetwork:getSubnetwork", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSubnetwork.
type LookupSubnetworkArgs struct {
	// The name of the subnetwork. One of `name` or `selfLink`
	// must be specified.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The region this subnetwork has been created in. If
	// unspecified, this defaults to the region configured in the provider.
	Region *string `pulumi:"region"`
	// The self link of the subnetwork. If `selfLink` is
	// specified, `name`, `project`, and `region` are ignored.
	SelfLink *string `pulumi:"selfLink"`
}


// A collection of values returned by getSubnetwork.
type LookupSubnetworkResult struct {
	// Description of this subnetwork.
	Description string `pulumi:"description"`
	// The IP address of the gateway.
	GatewayAddress string `pulumi:"gatewayAddress"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The range of IP addresses belonging to this subnetwork
	// secondary range.
	IpCidrRange string `pulumi:"ipCidrRange"`
	Name *string `pulumi:"name"`
	// The network name or resource link to the parent
	// network of this subnetwork.
	Network string `pulumi:"network"`
	// Whether the VMs in this subnet
	// can access Google services without assigned external IP
	// addresses.
	PrivateIpGoogleAccess bool `pulumi:"privateIpGoogleAccess"`
	Project string `pulumi:"project"`
	Region string `pulumi:"region"`
	// An array of configurations for secondary IP ranges for
	// VM instances contained in this subnetwork. Structure is documented below.
	SecondaryIpRanges []GetSubnetworkSecondaryIpRange `pulumi:"secondaryIpRanges"`
	SelfLink string `pulumi:"selfLink"`
}

