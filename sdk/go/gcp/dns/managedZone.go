// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dns

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A zone is a subtree of the DNS namespace under one administrative
// responsibility. A ManagedZone is a resource that represents a DNS zone
// hosted by the Cloud DNS service.
//
//
// To get more information about ManagedZone, see:
//
// * [API documentation](https://cloud.google.com/dns/api/v1/managedZones)
// * How-to Guides
//     * [Managing Zones](https://cloud.google.com/dns/zones/)
type ManagedZone struct {
	pulumi.CustomResourceState

	// A textual description field. Defaults to 'Managed by Pulumi'.
	Description pulumi.StringOutput `pulumi:"description"`
	// The DNS name of this managed zone, for instance "example.com.".
	DnsName pulumi.StringOutput `pulumi:"dnsName"`
	// DNSSEC configuration  Structure is documented below.
	DnssecConfig ManagedZoneDnssecConfigPtrOutput `pulumi:"dnssecConfig"`
	// The presence for this field indicates that outbound forwarding is enabled
	// for this zone. The value of this field contains the set of destinations
	// to forward to.  Structure is documented below.
	ForwardingConfig ManagedZoneForwardingConfigPtrOutput `pulumi:"forwardingConfig"`
	// A set of key/value label pairs to assign to this ManagedZone.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// User assigned name for this resource.
	// Must be unique within the project.
	Name pulumi.StringOutput `pulumi:"name"`
	// Delegate your managed_zone to these virtual name servers; defined by the server
	NameServers pulumi.StringArrayOutput `pulumi:"nameServers"`
	// The presence of this field indicates that DNS Peering is enabled for this
	// zone. The value of this field contains the network to peer with.  Structure is documented below.
	PeeringConfig ManagedZonePeeringConfigPtrOutput `pulumi:"peeringConfig"`
	// For privately visible zones, the set of Virtual Private Cloud
	// resources that the zone is visible from.  Structure is documented below.
	PrivateVisibilityConfig ManagedZonePrivateVisibilityConfigPtrOutput `pulumi:"privateVisibilityConfig"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Specifies if this is a managed reverse lookup zone. If true, Cloud DNS will resolve reverse
	// lookup queries using automatically configured records for VPC resources. This only applies
	// to networks listed under `privateVisibilityConfig`.
	ReverseLookup pulumi.BoolPtrOutput `pulumi:"reverseLookup"`
	// The presence of this field indicates that this zone is backed by Service Directory. The value of this field contains
	// information related to the namespace associated with the zone.
	ServiceDirectoryConfig ManagedZoneServiceDirectoryConfigPtrOutput `pulumi:"serviceDirectoryConfig"`
	// The zone's visibility: public zones are exposed to the Internet,
	// while private zones are visible only to Virtual Private Cloud resources.
	Visibility pulumi.StringPtrOutput `pulumi:"visibility"`
}

// NewManagedZone registers a new resource with the given unique name, arguments, and options.
func NewManagedZone(ctx *pulumi.Context,
	name string, args *ManagedZoneArgs, opts ...pulumi.ResourceOption) (*ManagedZone, error) {
	if args == nil || args.DnsName == nil {
		return nil, errors.New("missing required argument 'DnsName'")
	}
	if args == nil {
		args = &ManagedZoneArgs{}
	}
	if args.Description == nil {
		args.Description = pulumi.StringPtr("Managed by Pulumi")
	}
	var resource ManagedZone
	err := ctx.RegisterResource("gcp:dns/managedZone:ManagedZone", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagedZone gets an existing ManagedZone resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagedZone(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedZoneState, opts ...pulumi.ResourceOption) (*ManagedZone, error) {
	var resource ManagedZone
	err := ctx.ReadResource("gcp:dns/managedZone:ManagedZone", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagedZone resources.
type managedZoneState struct {
	// A textual description field. Defaults to 'Managed by Pulumi'.
	Description *string `pulumi:"description"`
	// The DNS name of this managed zone, for instance "example.com.".
	DnsName *string `pulumi:"dnsName"`
	// DNSSEC configuration  Structure is documented below.
	DnssecConfig *ManagedZoneDnssecConfig `pulumi:"dnssecConfig"`
	// The presence for this field indicates that outbound forwarding is enabled
	// for this zone. The value of this field contains the set of destinations
	// to forward to.  Structure is documented below.
	ForwardingConfig *ManagedZoneForwardingConfig `pulumi:"forwardingConfig"`
	// A set of key/value label pairs to assign to this ManagedZone.
	Labels map[string]string `pulumi:"labels"`
	// User assigned name for this resource.
	// Must be unique within the project.
	Name *string `pulumi:"name"`
	// Delegate your managed_zone to these virtual name servers; defined by the server
	NameServers []string `pulumi:"nameServers"`
	// The presence of this field indicates that DNS Peering is enabled for this
	// zone. The value of this field contains the network to peer with.  Structure is documented below.
	PeeringConfig *ManagedZonePeeringConfig `pulumi:"peeringConfig"`
	// For privately visible zones, the set of Virtual Private Cloud
	// resources that the zone is visible from.  Structure is documented below.
	PrivateVisibilityConfig *ManagedZonePrivateVisibilityConfig `pulumi:"privateVisibilityConfig"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Specifies if this is a managed reverse lookup zone. If true, Cloud DNS will resolve reverse
	// lookup queries using automatically configured records for VPC resources. This only applies
	// to networks listed under `privateVisibilityConfig`.
	ReverseLookup *bool `pulumi:"reverseLookup"`
	// The presence of this field indicates that this zone is backed by Service Directory. The value of this field contains
	// information related to the namespace associated with the zone.
	ServiceDirectoryConfig *ManagedZoneServiceDirectoryConfig `pulumi:"serviceDirectoryConfig"`
	// The zone's visibility: public zones are exposed to the Internet,
	// while private zones are visible only to Virtual Private Cloud resources.
	Visibility *string `pulumi:"visibility"`
}

type ManagedZoneState struct {
	// A textual description field. Defaults to 'Managed by Pulumi'.
	Description pulumi.StringPtrInput
	// The DNS name of this managed zone, for instance "example.com.".
	DnsName pulumi.StringPtrInput
	// DNSSEC configuration  Structure is documented below.
	DnssecConfig ManagedZoneDnssecConfigPtrInput
	// The presence for this field indicates that outbound forwarding is enabled
	// for this zone. The value of this field contains the set of destinations
	// to forward to.  Structure is documented below.
	ForwardingConfig ManagedZoneForwardingConfigPtrInput
	// A set of key/value label pairs to assign to this ManagedZone.
	Labels pulumi.StringMapInput
	// User assigned name for this resource.
	// Must be unique within the project.
	Name pulumi.StringPtrInput
	// Delegate your managed_zone to these virtual name servers; defined by the server
	NameServers pulumi.StringArrayInput
	// The presence of this field indicates that DNS Peering is enabled for this
	// zone. The value of this field contains the network to peer with.  Structure is documented below.
	PeeringConfig ManagedZonePeeringConfigPtrInput
	// For privately visible zones, the set of Virtual Private Cloud
	// resources that the zone is visible from.  Structure is documented below.
	PrivateVisibilityConfig ManagedZonePrivateVisibilityConfigPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Specifies if this is a managed reverse lookup zone. If true, Cloud DNS will resolve reverse
	// lookup queries using automatically configured records for VPC resources. This only applies
	// to networks listed under `privateVisibilityConfig`.
	ReverseLookup pulumi.BoolPtrInput
	// The presence of this field indicates that this zone is backed by Service Directory. The value of this field contains
	// information related to the namespace associated with the zone.
	ServiceDirectoryConfig ManagedZoneServiceDirectoryConfigPtrInput
	// The zone's visibility: public zones are exposed to the Internet,
	// while private zones are visible only to Virtual Private Cloud resources.
	Visibility pulumi.StringPtrInput
}

func (ManagedZoneState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedZoneState)(nil)).Elem()
}

type managedZoneArgs struct {
	// A textual description field. Defaults to 'Managed by Pulumi'.
	Description *string `pulumi:"description"`
	// The DNS name of this managed zone, for instance "example.com.".
	DnsName string `pulumi:"dnsName"`
	// DNSSEC configuration  Structure is documented below.
	DnssecConfig *ManagedZoneDnssecConfig `pulumi:"dnssecConfig"`
	// The presence for this field indicates that outbound forwarding is enabled
	// for this zone. The value of this field contains the set of destinations
	// to forward to.  Structure is documented below.
	ForwardingConfig *ManagedZoneForwardingConfig `pulumi:"forwardingConfig"`
	// A set of key/value label pairs to assign to this ManagedZone.
	Labels map[string]string `pulumi:"labels"`
	// User assigned name for this resource.
	// Must be unique within the project.
	Name *string `pulumi:"name"`
	// The presence of this field indicates that DNS Peering is enabled for this
	// zone. The value of this field contains the network to peer with.  Structure is documented below.
	PeeringConfig *ManagedZonePeeringConfig `pulumi:"peeringConfig"`
	// For privately visible zones, the set of Virtual Private Cloud
	// resources that the zone is visible from.  Structure is documented below.
	PrivateVisibilityConfig *ManagedZonePrivateVisibilityConfig `pulumi:"privateVisibilityConfig"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Specifies if this is a managed reverse lookup zone. If true, Cloud DNS will resolve reverse
	// lookup queries using automatically configured records for VPC resources. This only applies
	// to networks listed under `privateVisibilityConfig`.
	ReverseLookup *bool `pulumi:"reverseLookup"`
	// The presence of this field indicates that this zone is backed by Service Directory. The value of this field contains
	// information related to the namespace associated with the zone.
	ServiceDirectoryConfig *ManagedZoneServiceDirectoryConfig `pulumi:"serviceDirectoryConfig"`
	// The zone's visibility: public zones are exposed to the Internet,
	// while private zones are visible only to Virtual Private Cloud resources.
	Visibility *string `pulumi:"visibility"`
}

// The set of arguments for constructing a ManagedZone resource.
type ManagedZoneArgs struct {
	// A textual description field. Defaults to 'Managed by Pulumi'.
	Description pulumi.StringPtrInput
	// The DNS name of this managed zone, for instance "example.com.".
	DnsName pulumi.StringInput
	// DNSSEC configuration  Structure is documented below.
	DnssecConfig ManagedZoneDnssecConfigPtrInput
	// The presence for this field indicates that outbound forwarding is enabled
	// for this zone. The value of this field contains the set of destinations
	// to forward to.  Structure is documented below.
	ForwardingConfig ManagedZoneForwardingConfigPtrInput
	// A set of key/value label pairs to assign to this ManagedZone.
	Labels pulumi.StringMapInput
	// User assigned name for this resource.
	// Must be unique within the project.
	Name pulumi.StringPtrInput
	// The presence of this field indicates that DNS Peering is enabled for this
	// zone. The value of this field contains the network to peer with.  Structure is documented below.
	PeeringConfig ManagedZonePeeringConfigPtrInput
	// For privately visible zones, the set of Virtual Private Cloud
	// resources that the zone is visible from.  Structure is documented below.
	PrivateVisibilityConfig ManagedZonePrivateVisibilityConfigPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Specifies if this is a managed reverse lookup zone. If true, Cloud DNS will resolve reverse
	// lookup queries using automatically configured records for VPC resources. This only applies
	// to networks listed under `privateVisibilityConfig`.
	ReverseLookup pulumi.BoolPtrInput
	// The presence of this field indicates that this zone is backed by Service Directory. The value of this field contains
	// information related to the namespace associated with the zone.
	ServiceDirectoryConfig ManagedZoneServiceDirectoryConfigPtrInput
	// The zone's visibility: public zones are exposed to the Internet,
	// while private zones are visible only to Virtual Private Cloud resources.
	Visibility pulumi.StringPtrInput
}

func (ManagedZoneArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedZoneArgs)(nil)).Elem()
}
