// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Represents a GlobalForwardingRule resource. Global forwarding rules are
// used to forward traffic to the correct load balancer for HTTP load
// balancing. Global forwarding rules can only be used for HTTP load
// balancing.
//
// For more information, see
// https://cloud.google.com/compute/docs/load-balancing/http/
type GlobalForwardingRule struct {
	pulumi.CustomResourceState

	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The IP address that this forwarding rule is serving on behalf of.
	// Addresses are restricted based on the forwarding rule's load balancing
	// scheme (EXTERNAL or INTERNAL) and scope (global or regional).
	// When the load balancing scheme is EXTERNAL, for global forwarding
	// rules, the address must be a global IP, and for regional forwarding
	// rules, the address must live in the same region as the forwarding
	// rule. If this field is empty, an ephemeral IPv4 address from the same
	// scope (global or regional) will be assigned. A regional forwarding
	// rule supports IPv4 only. A global forwarding rule supports either IPv4
	// or IPv6.
	// When the load balancing scheme is INTERNAL, this can only be an RFC
	// 1918 IP address belonging to the network/subnet configured for the
	// forwarding rule. By default, if this field is empty, an ephemeral
	// internal IP address will be automatically allocated from the IP range
	// of the subnet or network configured for this forwarding rule.
	// An address must be specified by a literal IP address. > **NOTE**: While
	// the API allows you to specify various resource paths for an address resource
	// instead, this provider requires this to specifically be an IP address to
	// avoid needing to fetching the IP address from resource paths on refresh
	// or unnecessary diffs.
	IpAddress pulumi.StringOutput `pulumi:"ipAddress"`
	// The IP protocol to which this rule applies. When the load balancing scheme is
	// INTERNAL_SELF_MANAGED, only TCP is valid.
	IpProtocol pulumi.StringOutput `pulumi:"ipProtocol"`
	// The IP Version that will be used by this global forwarding rule.
	IpVersion pulumi.StringPtrOutput `pulumi:"ipVersion"`
	// The fingerprint used for optimistic locking of this resource. Used internally during updates.
	LabelFingerprint pulumi.StringOutput `pulumi:"labelFingerprint"`
	// Labels to apply to this forwarding rule.  A list of key->value pairs.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// This signifies what the GlobalForwardingRule will be used for.
	// The value of INTERNAL_SELF_MANAGED means that this will be used for
	// Internal Global HTTP(S) LB. The value of EXTERNAL means that this
	// will be used for External Global Load Balancing (HTTP(S) LB,
	// External TCP/UDP LB, SSL Proxy)
	// NOTE: Currently global forwarding rules cannot be used for INTERNAL
	// load balancing.
	LoadBalancingScheme pulumi.StringPtrOutput `pulumi:"loadBalancingScheme"`
	// Opaque filter criteria used by Loadbalancer to restrict routing
	// configuration to a limited set xDS compliant clients. In their xDS
	// requests to Loadbalancer, xDS clients present node metadata. If a
	// match takes place, the relevant routing configuration is made available
	// to those proxies.
	// For each metadataFilter in this list, if its filterMatchCriteria is set
	// to MATCH_ANY, at least one of the filterLabels must match the
	// corresponding label provided in the metadata. If its filterMatchCriteria
	// is set to MATCH_ALL, then all of its filterLabels must match with
	// corresponding labels in the provided metadata.
	// metadataFilters specified here can be overridden by those specified in
	// the UrlMap that this ForwardingRule references.
	// metadataFilters only applies to Loadbalancers that have their
	// loadBalancingScheme set to INTERNAL_SELF_MANAGED.  Structure is documented below.
	MetadataFilters GlobalForwardingRuleMetadataFilterArrayOutput `pulumi:"metadataFilters"`
	// Name of the metadata label. The length must be between
	// 1 and 1024 characters, inclusive.
	Name pulumi.StringOutput `pulumi:"name"`
	// This field is not used for external load balancing.
	// For INTERNAL_SELF_MANAGED load balancing, this field
	// identifies the network that the load balanced IP should belong to
	// for this global forwarding rule. If this field is not specified,
	// the default network will be used.
	Network pulumi.StringOutput `pulumi:"network"`
	// This field is used along with the target field for TargetHttpProxy,
	// TargetHttpsProxy, TargetSslProxy, TargetTcpProxy, TargetVpnGateway,
	// TargetPool, TargetInstance.
	// Applicable only when IPProtocol is TCP, UDP, or SCTP, only packets
	// addressed to ports in the specified range will be forwarded to target.
	// Forwarding rules with the same [IPAddress, IPProtocol] pair must have
	// disjoint port ranges.
	// Some types of forwarding target have constraints on the acceptable
	// ports:
	// * TargetHttpProxy: 80, 8080
	// * TargetHttpsProxy: 443
	// * TargetTcpProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995,
	// 1883, 5222
	// * TargetSslProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995,
	// 1883, 5222
	// * TargetVpnGateway: 500, 4500
	PortRange pulumi.StringPtrOutput `pulumi:"portRange"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// The URL of the target resource to receive the matched traffic.
	// The forwarded traffic must be of a type appropriate to the target object.
	// For INTERNAL_SELF_MANAGED load balancing, only HTTP and HTTPS targets
	// are valid.
	Target pulumi.StringOutput `pulumi:"target"`
}

// NewGlobalForwardingRule registers a new resource with the given unique name, arguments, and options.
func NewGlobalForwardingRule(ctx *pulumi.Context,
	name string, args *GlobalForwardingRuleArgs, opts ...pulumi.ResourceOption) (*GlobalForwardingRule, error) {
	if args == nil || args.Target == nil {
		return nil, errors.New("missing required argument 'Target'")
	}
	if args == nil {
		args = &GlobalForwardingRuleArgs{}
	}
	var resource GlobalForwardingRule
	err := ctx.RegisterResource("gcp:compute/globalForwardingRule:GlobalForwardingRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGlobalForwardingRule gets an existing GlobalForwardingRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGlobalForwardingRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GlobalForwardingRuleState, opts ...pulumi.ResourceOption) (*GlobalForwardingRule, error) {
	var resource GlobalForwardingRule
	err := ctx.ReadResource("gcp:compute/globalForwardingRule:GlobalForwardingRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GlobalForwardingRule resources.
type globalForwardingRuleState struct {
	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description *string `pulumi:"description"`
	// The IP address that this forwarding rule is serving on behalf of.
	// Addresses are restricted based on the forwarding rule's load balancing
	// scheme (EXTERNAL or INTERNAL) and scope (global or regional).
	// When the load balancing scheme is EXTERNAL, for global forwarding
	// rules, the address must be a global IP, and for regional forwarding
	// rules, the address must live in the same region as the forwarding
	// rule. If this field is empty, an ephemeral IPv4 address from the same
	// scope (global or regional) will be assigned. A regional forwarding
	// rule supports IPv4 only. A global forwarding rule supports either IPv4
	// or IPv6.
	// When the load balancing scheme is INTERNAL, this can only be an RFC
	// 1918 IP address belonging to the network/subnet configured for the
	// forwarding rule. By default, if this field is empty, an ephemeral
	// internal IP address will be automatically allocated from the IP range
	// of the subnet or network configured for this forwarding rule.
	// An address must be specified by a literal IP address. > **NOTE**: While
	// the API allows you to specify various resource paths for an address resource
	// instead, this provider requires this to specifically be an IP address to
	// avoid needing to fetching the IP address from resource paths on refresh
	// or unnecessary diffs.
	IpAddress *string `pulumi:"ipAddress"`
	// The IP protocol to which this rule applies. When the load balancing scheme is
	// INTERNAL_SELF_MANAGED, only TCP is valid.
	IpProtocol *string `pulumi:"ipProtocol"`
	// The IP Version that will be used by this global forwarding rule.
	IpVersion *string `pulumi:"ipVersion"`
	// The fingerprint used for optimistic locking of this resource. Used internally during updates.
	LabelFingerprint *string `pulumi:"labelFingerprint"`
	// Labels to apply to this forwarding rule.  A list of key->value pairs.
	Labels map[string]string `pulumi:"labels"`
	// This signifies what the GlobalForwardingRule will be used for.
	// The value of INTERNAL_SELF_MANAGED means that this will be used for
	// Internal Global HTTP(S) LB. The value of EXTERNAL means that this
	// will be used for External Global Load Balancing (HTTP(S) LB,
	// External TCP/UDP LB, SSL Proxy)
	// NOTE: Currently global forwarding rules cannot be used for INTERNAL
	// load balancing.
	LoadBalancingScheme *string `pulumi:"loadBalancingScheme"`
	// Opaque filter criteria used by Loadbalancer to restrict routing
	// configuration to a limited set xDS compliant clients. In their xDS
	// requests to Loadbalancer, xDS clients present node metadata. If a
	// match takes place, the relevant routing configuration is made available
	// to those proxies.
	// For each metadataFilter in this list, if its filterMatchCriteria is set
	// to MATCH_ANY, at least one of the filterLabels must match the
	// corresponding label provided in the metadata. If its filterMatchCriteria
	// is set to MATCH_ALL, then all of its filterLabels must match with
	// corresponding labels in the provided metadata.
	// metadataFilters specified here can be overridden by those specified in
	// the UrlMap that this ForwardingRule references.
	// metadataFilters only applies to Loadbalancers that have their
	// loadBalancingScheme set to INTERNAL_SELF_MANAGED.  Structure is documented below.
	MetadataFilters []GlobalForwardingRuleMetadataFilter `pulumi:"metadataFilters"`
	// Name of the metadata label. The length must be between
	// 1 and 1024 characters, inclusive.
	Name *string `pulumi:"name"`
	// This field is not used for external load balancing.
	// For INTERNAL_SELF_MANAGED load balancing, this field
	// identifies the network that the load balanced IP should belong to
	// for this global forwarding rule. If this field is not specified,
	// the default network will be used.
	Network *string `pulumi:"network"`
	// This field is used along with the target field for TargetHttpProxy,
	// TargetHttpsProxy, TargetSslProxy, TargetTcpProxy, TargetVpnGateway,
	// TargetPool, TargetInstance.
	// Applicable only when IPProtocol is TCP, UDP, or SCTP, only packets
	// addressed to ports in the specified range will be forwarded to target.
	// Forwarding rules with the same [IPAddress, IPProtocol] pair must have
	// disjoint port ranges.
	// Some types of forwarding target have constraints on the acceptable
	// ports:
	// * TargetHttpProxy: 80, 8080
	// * TargetHttpsProxy: 443
	// * TargetTcpProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995,
	// 1883, 5222
	// * TargetSslProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995,
	// 1883, 5222
	// * TargetVpnGateway: 500, 4500
	PortRange *string `pulumi:"portRange"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
	// The URL of the target resource to receive the matched traffic.
	// The forwarded traffic must be of a type appropriate to the target object.
	// For INTERNAL_SELF_MANAGED load balancing, only HTTP and HTTPS targets
	// are valid.
	Target *string `pulumi:"target"`
}

type GlobalForwardingRuleState struct {
	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description pulumi.StringPtrInput
	// The IP address that this forwarding rule is serving on behalf of.
	// Addresses are restricted based on the forwarding rule's load balancing
	// scheme (EXTERNAL or INTERNAL) and scope (global or regional).
	// When the load balancing scheme is EXTERNAL, for global forwarding
	// rules, the address must be a global IP, and for regional forwarding
	// rules, the address must live in the same region as the forwarding
	// rule. If this field is empty, an ephemeral IPv4 address from the same
	// scope (global or regional) will be assigned. A regional forwarding
	// rule supports IPv4 only. A global forwarding rule supports either IPv4
	// or IPv6.
	// When the load balancing scheme is INTERNAL, this can only be an RFC
	// 1918 IP address belonging to the network/subnet configured for the
	// forwarding rule. By default, if this field is empty, an ephemeral
	// internal IP address will be automatically allocated from the IP range
	// of the subnet or network configured for this forwarding rule.
	// An address must be specified by a literal IP address. > **NOTE**: While
	// the API allows you to specify various resource paths for an address resource
	// instead, this provider requires this to specifically be an IP address to
	// avoid needing to fetching the IP address from resource paths on refresh
	// or unnecessary diffs.
	IpAddress pulumi.StringPtrInput
	// The IP protocol to which this rule applies. When the load balancing scheme is
	// INTERNAL_SELF_MANAGED, only TCP is valid.
	IpProtocol pulumi.StringPtrInput
	// The IP Version that will be used by this global forwarding rule.
	IpVersion pulumi.StringPtrInput
	// The fingerprint used for optimistic locking of this resource. Used internally during updates.
	LabelFingerprint pulumi.StringPtrInput
	// Labels to apply to this forwarding rule.  A list of key->value pairs.
	Labels pulumi.StringMapInput
	// This signifies what the GlobalForwardingRule will be used for.
	// The value of INTERNAL_SELF_MANAGED means that this will be used for
	// Internal Global HTTP(S) LB. The value of EXTERNAL means that this
	// will be used for External Global Load Balancing (HTTP(S) LB,
	// External TCP/UDP LB, SSL Proxy)
	// NOTE: Currently global forwarding rules cannot be used for INTERNAL
	// load balancing.
	LoadBalancingScheme pulumi.StringPtrInput
	// Opaque filter criteria used by Loadbalancer to restrict routing
	// configuration to a limited set xDS compliant clients. In their xDS
	// requests to Loadbalancer, xDS clients present node metadata. If a
	// match takes place, the relevant routing configuration is made available
	// to those proxies.
	// For each metadataFilter in this list, if its filterMatchCriteria is set
	// to MATCH_ANY, at least one of the filterLabels must match the
	// corresponding label provided in the metadata. If its filterMatchCriteria
	// is set to MATCH_ALL, then all of its filterLabels must match with
	// corresponding labels in the provided metadata.
	// metadataFilters specified here can be overridden by those specified in
	// the UrlMap that this ForwardingRule references.
	// metadataFilters only applies to Loadbalancers that have their
	// loadBalancingScheme set to INTERNAL_SELF_MANAGED.  Structure is documented below.
	MetadataFilters GlobalForwardingRuleMetadataFilterArrayInput
	// Name of the metadata label. The length must be between
	// 1 and 1024 characters, inclusive.
	Name pulumi.StringPtrInput
	// This field is not used for external load balancing.
	// For INTERNAL_SELF_MANAGED load balancing, this field
	// identifies the network that the load balanced IP should belong to
	// for this global forwarding rule. If this field is not specified,
	// the default network will be used.
	Network pulumi.StringPtrInput
	// This field is used along with the target field for TargetHttpProxy,
	// TargetHttpsProxy, TargetSslProxy, TargetTcpProxy, TargetVpnGateway,
	// TargetPool, TargetInstance.
	// Applicable only when IPProtocol is TCP, UDP, or SCTP, only packets
	// addressed to ports in the specified range will be forwarded to target.
	// Forwarding rules with the same [IPAddress, IPProtocol] pair must have
	// disjoint port ranges.
	// Some types of forwarding target have constraints on the acceptable
	// ports:
	// * TargetHttpProxy: 80, 8080
	// * TargetHttpsProxy: 443
	// * TargetTcpProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995,
	// 1883, 5222
	// * TargetSslProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995,
	// 1883, 5222
	// * TargetVpnGateway: 500, 4500
	PortRange pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
	// The URL of the target resource to receive the matched traffic.
	// The forwarded traffic must be of a type appropriate to the target object.
	// For INTERNAL_SELF_MANAGED load balancing, only HTTP and HTTPS targets
	// are valid.
	Target pulumi.StringPtrInput
}

func (GlobalForwardingRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*globalForwardingRuleState)(nil)).Elem()
}

type globalForwardingRuleArgs struct {
	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description *string `pulumi:"description"`
	// The IP address that this forwarding rule is serving on behalf of.
	// Addresses are restricted based on the forwarding rule's load balancing
	// scheme (EXTERNAL or INTERNAL) and scope (global or regional).
	// When the load balancing scheme is EXTERNAL, for global forwarding
	// rules, the address must be a global IP, and for regional forwarding
	// rules, the address must live in the same region as the forwarding
	// rule. If this field is empty, an ephemeral IPv4 address from the same
	// scope (global or regional) will be assigned. A regional forwarding
	// rule supports IPv4 only. A global forwarding rule supports either IPv4
	// or IPv6.
	// When the load balancing scheme is INTERNAL, this can only be an RFC
	// 1918 IP address belonging to the network/subnet configured for the
	// forwarding rule. By default, if this field is empty, an ephemeral
	// internal IP address will be automatically allocated from the IP range
	// of the subnet or network configured for this forwarding rule.
	// An address must be specified by a literal IP address. > **NOTE**: While
	// the API allows you to specify various resource paths for an address resource
	// instead, this provider requires this to specifically be an IP address to
	// avoid needing to fetching the IP address from resource paths on refresh
	// or unnecessary diffs.
	IpAddress *string `pulumi:"ipAddress"`
	// The IP protocol to which this rule applies. When the load balancing scheme is
	// INTERNAL_SELF_MANAGED, only TCP is valid.
	IpProtocol *string `pulumi:"ipProtocol"`
	// The IP Version that will be used by this global forwarding rule.
	IpVersion *string `pulumi:"ipVersion"`
	// Labels to apply to this forwarding rule.  A list of key->value pairs.
	Labels map[string]string `pulumi:"labels"`
	// This signifies what the GlobalForwardingRule will be used for.
	// The value of INTERNAL_SELF_MANAGED means that this will be used for
	// Internal Global HTTP(S) LB. The value of EXTERNAL means that this
	// will be used for External Global Load Balancing (HTTP(S) LB,
	// External TCP/UDP LB, SSL Proxy)
	// NOTE: Currently global forwarding rules cannot be used for INTERNAL
	// load balancing.
	LoadBalancingScheme *string `pulumi:"loadBalancingScheme"`
	// Opaque filter criteria used by Loadbalancer to restrict routing
	// configuration to a limited set xDS compliant clients. In their xDS
	// requests to Loadbalancer, xDS clients present node metadata. If a
	// match takes place, the relevant routing configuration is made available
	// to those proxies.
	// For each metadataFilter in this list, if its filterMatchCriteria is set
	// to MATCH_ANY, at least one of the filterLabels must match the
	// corresponding label provided in the metadata. If its filterMatchCriteria
	// is set to MATCH_ALL, then all of its filterLabels must match with
	// corresponding labels in the provided metadata.
	// metadataFilters specified here can be overridden by those specified in
	// the UrlMap that this ForwardingRule references.
	// metadataFilters only applies to Loadbalancers that have their
	// loadBalancingScheme set to INTERNAL_SELF_MANAGED.  Structure is documented below.
	MetadataFilters []GlobalForwardingRuleMetadataFilter `pulumi:"metadataFilters"`
	// Name of the metadata label. The length must be between
	// 1 and 1024 characters, inclusive.
	Name *string `pulumi:"name"`
	// This field is not used for external load balancing.
	// For INTERNAL_SELF_MANAGED load balancing, this field
	// identifies the network that the load balanced IP should belong to
	// for this global forwarding rule. If this field is not specified,
	// the default network will be used.
	Network *string `pulumi:"network"`
	// This field is used along with the target field for TargetHttpProxy,
	// TargetHttpsProxy, TargetSslProxy, TargetTcpProxy, TargetVpnGateway,
	// TargetPool, TargetInstance.
	// Applicable only when IPProtocol is TCP, UDP, or SCTP, only packets
	// addressed to ports in the specified range will be forwarded to target.
	// Forwarding rules with the same [IPAddress, IPProtocol] pair must have
	// disjoint port ranges.
	// Some types of forwarding target have constraints on the acceptable
	// ports:
	// * TargetHttpProxy: 80, 8080
	// * TargetHttpsProxy: 443
	// * TargetTcpProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995,
	// 1883, 5222
	// * TargetSslProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995,
	// 1883, 5222
	// * TargetVpnGateway: 500, 4500
	PortRange *string `pulumi:"portRange"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The URL of the target resource to receive the matched traffic.
	// The forwarded traffic must be of a type appropriate to the target object.
	// For INTERNAL_SELF_MANAGED load balancing, only HTTP and HTTPS targets
	// are valid.
	Target string `pulumi:"target"`
}

// The set of arguments for constructing a GlobalForwardingRule resource.
type GlobalForwardingRuleArgs struct {
	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description pulumi.StringPtrInput
	// The IP address that this forwarding rule is serving on behalf of.
	// Addresses are restricted based on the forwarding rule's load balancing
	// scheme (EXTERNAL or INTERNAL) and scope (global or regional).
	// When the load balancing scheme is EXTERNAL, for global forwarding
	// rules, the address must be a global IP, and for regional forwarding
	// rules, the address must live in the same region as the forwarding
	// rule. If this field is empty, an ephemeral IPv4 address from the same
	// scope (global or regional) will be assigned. A regional forwarding
	// rule supports IPv4 only. A global forwarding rule supports either IPv4
	// or IPv6.
	// When the load balancing scheme is INTERNAL, this can only be an RFC
	// 1918 IP address belonging to the network/subnet configured for the
	// forwarding rule. By default, if this field is empty, an ephemeral
	// internal IP address will be automatically allocated from the IP range
	// of the subnet or network configured for this forwarding rule.
	// An address must be specified by a literal IP address. > **NOTE**: While
	// the API allows you to specify various resource paths for an address resource
	// instead, this provider requires this to specifically be an IP address to
	// avoid needing to fetching the IP address from resource paths on refresh
	// or unnecessary diffs.
	IpAddress pulumi.StringPtrInput
	// The IP protocol to which this rule applies. When the load balancing scheme is
	// INTERNAL_SELF_MANAGED, only TCP is valid.
	IpProtocol pulumi.StringPtrInput
	// The IP Version that will be used by this global forwarding rule.
	IpVersion pulumi.StringPtrInput
	// Labels to apply to this forwarding rule.  A list of key->value pairs.
	Labels pulumi.StringMapInput
	// This signifies what the GlobalForwardingRule will be used for.
	// The value of INTERNAL_SELF_MANAGED means that this will be used for
	// Internal Global HTTP(S) LB. The value of EXTERNAL means that this
	// will be used for External Global Load Balancing (HTTP(S) LB,
	// External TCP/UDP LB, SSL Proxy)
	// NOTE: Currently global forwarding rules cannot be used for INTERNAL
	// load balancing.
	LoadBalancingScheme pulumi.StringPtrInput
	// Opaque filter criteria used by Loadbalancer to restrict routing
	// configuration to a limited set xDS compliant clients. In their xDS
	// requests to Loadbalancer, xDS clients present node metadata. If a
	// match takes place, the relevant routing configuration is made available
	// to those proxies.
	// For each metadataFilter in this list, if its filterMatchCriteria is set
	// to MATCH_ANY, at least one of the filterLabels must match the
	// corresponding label provided in the metadata. If its filterMatchCriteria
	// is set to MATCH_ALL, then all of its filterLabels must match with
	// corresponding labels in the provided metadata.
	// metadataFilters specified here can be overridden by those specified in
	// the UrlMap that this ForwardingRule references.
	// metadataFilters only applies to Loadbalancers that have their
	// loadBalancingScheme set to INTERNAL_SELF_MANAGED.  Structure is documented below.
	MetadataFilters GlobalForwardingRuleMetadataFilterArrayInput
	// Name of the metadata label. The length must be between
	// 1 and 1024 characters, inclusive.
	Name pulumi.StringPtrInput
	// This field is not used for external load balancing.
	// For INTERNAL_SELF_MANAGED load balancing, this field
	// identifies the network that the load balanced IP should belong to
	// for this global forwarding rule. If this field is not specified,
	// the default network will be used.
	Network pulumi.StringPtrInput
	// This field is used along with the target field for TargetHttpProxy,
	// TargetHttpsProxy, TargetSslProxy, TargetTcpProxy, TargetVpnGateway,
	// TargetPool, TargetInstance.
	// Applicable only when IPProtocol is TCP, UDP, or SCTP, only packets
	// addressed to ports in the specified range will be forwarded to target.
	// Forwarding rules with the same [IPAddress, IPProtocol] pair must have
	// disjoint port ranges.
	// Some types of forwarding target have constraints on the acceptable
	// ports:
	// * TargetHttpProxy: 80, 8080
	// * TargetHttpsProxy: 443
	// * TargetTcpProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995,
	// 1883, 5222
	// * TargetSslProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995,
	// 1883, 5222
	// * TargetVpnGateway: 500, 4500
	PortRange pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The URL of the target resource to receive the matched traffic.
	// The forwarded traffic must be of a type appropriate to the target object.
	// For INTERNAL_SELF_MANAGED load balancing, only HTTP and HTTPS targets
	// are valid.
	Target pulumi.StringInput
}

func (GlobalForwardingRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*globalForwardingRuleArgs)(nil)).Elem()
}
