// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_forwarding_rule.html.markdown.
type ForwardingRule struct {
	s *pulumi.ResourceState
}

// NewForwardingRule registers a new resource with the given unique name, arguments, and options.
func NewForwardingRule(ctx *pulumi.Context,
	name string, args *ForwardingRuleArgs, opts ...pulumi.ResourceOpt) (*ForwardingRule, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["allPorts"] = nil
		inputs["backendService"] = nil
		inputs["description"] = nil
		inputs["ipAddress"] = nil
		inputs["ipProtocol"] = nil
		inputs["labels"] = nil
		inputs["loadBalancingScheme"] = nil
		inputs["name"] = nil
		inputs["network"] = nil
		inputs["networkTier"] = nil
		inputs["portRange"] = nil
		inputs["ports"] = nil
		inputs["project"] = nil
		inputs["region"] = nil
		inputs["serviceLabel"] = nil
		inputs["subnetwork"] = nil
		inputs["target"] = nil
	} else {
		inputs["allPorts"] = args.AllPorts
		inputs["backendService"] = args.BackendService
		inputs["description"] = args.Description
		inputs["ipAddress"] = args.IpAddress
		inputs["ipProtocol"] = args.IpProtocol
		inputs["labels"] = args.Labels
		inputs["loadBalancingScheme"] = args.LoadBalancingScheme
		inputs["name"] = args.Name
		inputs["network"] = args.Network
		inputs["networkTier"] = args.NetworkTier
		inputs["portRange"] = args.PortRange
		inputs["ports"] = args.Ports
		inputs["project"] = args.Project
		inputs["region"] = args.Region
		inputs["serviceLabel"] = args.ServiceLabel
		inputs["subnetwork"] = args.Subnetwork
		inputs["target"] = args.Target
	}
	inputs["creationTimestamp"] = nil
	inputs["labelFingerprint"] = nil
	inputs["selfLink"] = nil
	inputs["serviceName"] = nil
	s, err := ctx.RegisterResource("gcp:compute/forwardingRule:ForwardingRule", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ForwardingRule{s: s}, nil
}

// GetForwardingRule gets an existing ForwardingRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetForwardingRule(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ForwardingRuleState, opts ...pulumi.ResourceOpt) (*ForwardingRule, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["allPorts"] = state.AllPorts
		inputs["backendService"] = state.BackendService
		inputs["creationTimestamp"] = state.CreationTimestamp
		inputs["description"] = state.Description
		inputs["ipAddress"] = state.IpAddress
		inputs["ipProtocol"] = state.IpProtocol
		inputs["labelFingerprint"] = state.LabelFingerprint
		inputs["labels"] = state.Labels
		inputs["loadBalancingScheme"] = state.LoadBalancingScheme
		inputs["name"] = state.Name
		inputs["network"] = state.Network
		inputs["networkTier"] = state.NetworkTier
		inputs["portRange"] = state.PortRange
		inputs["ports"] = state.Ports
		inputs["project"] = state.Project
		inputs["region"] = state.Region
		inputs["selfLink"] = state.SelfLink
		inputs["serviceLabel"] = state.ServiceLabel
		inputs["serviceName"] = state.ServiceName
		inputs["subnetwork"] = state.Subnetwork
		inputs["target"] = state.Target
	}
	s, err := ctx.ReadResource("gcp:compute/forwardingRule:ForwardingRule", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ForwardingRule{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *ForwardingRule) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *ForwardingRule) ID() pulumi.IDOutput {
	return r.s.ID()
}

// For internal TCP/UDP load balancing (i.e. load balancing scheme is INTERNAL and protocol is TCP/UDP), set this to true
// to allow packets addressed to any ports to be forwarded to the backends configured with this forwarding rule. Used with
// backend service. Cannot be set if port or portRange are set.
func (r *ForwardingRule) AllPorts() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["allPorts"])
}

// A BackendService to receive the matched traffic. This is used only for INTERNAL load balancing.
func (r *ForwardingRule) BackendService() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["backendService"])
}

// Creation timestamp in RFC3339 text format.
func (r *ForwardingRule) CreationTimestamp() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["creationTimestamp"])
}

// An optional description of this resource. Provide this property when you create the resource.
func (r *ForwardingRule) Description() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["description"])
}

// The IP address that this forwarding rule is serving on behalf of. Addresses are restricted based on the forwarding
// rule's load balancing scheme (EXTERNAL or INTERNAL) and scope (global or regional). When the load balancing scheme is
// EXTERNAL, for global forwarding rules, the address must be a global IP, and for regional forwarding rules, the address
// must live in the same region as the forwarding rule. If this field is empty, an ephemeral IPv4 address from the same
// scope (global or regional) will be assigned. A regional forwarding rule supports IPv4 only. A global forwarding rule
// supports either IPv4 or IPv6. When the load balancing scheme is INTERNAL, this can only be an RFC 1918 IP address
// belonging to the network/subnet configured for the forwarding rule. By default, if this field is empty, an ephemeral
// internal IP address will be automatically allocated from the IP range of the subnet or network configured for this
// forwarding rule. An address must be specified by a literal IP address. ~> **NOTE**: While the API allows you to specify
// various resource paths for an address resource instead, Terraform requires this to specifically be an IP address to
// avoid needing to fetching the IP address from resource paths on refresh or unnecessary diffs.
func (r *ForwardingRule) IpAddress() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["ipAddress"])
}

// The IP protocol to which this rule applies. Valid options are TCP, UDP, ESP, AH, SCTP or ICMP. When the load balancing
// scheme is INTERNAL, only TCP and UDP are valid.
func (r *ForwardingRule) IpProtocol() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["ipProtocol"])
}

// The fingerprint used for optimistic locking of this resource. Used internally during updates.
func (r *ForwardingRule) LabelFingerprint() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["labelFingerprint"])
}

// Labels to apply to this forwarding rule. A list of key->value pairs.
func (r *ForwardingRule) Labels() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["labels"])
}

// This signifies what the ForwardingRule will be used for and can be EXTERNAL, INTERNAL, or INTERNAL_MANAGED. EXTERNAL is
// used for Classic Cloud VPN gateways, protocol forwarding to VMs from an external IP address, and HTTP(S), SSL Proxy, TCP
// Proxy, and Network TCP/UDP load balancers. INTERNAL is used for protocol forwarding to VMs from an internal IP address,
// and internal TCP/UDP load balancers. INTERNAL_MANAGED is used for internal HTTP(S) load balancers.
func (r *ForwardingRule) LoadBalancingScheme() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["loadBalancingScheme"])
}

// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and
// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (r *ForwardingRule) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// For internal load balancing, this field identifies the network that the load balanced IP should belong to for this
// Forwarding Rule. If this field is not specified, the default network will be used. This field is only used for INTERNAL
// load balancing.
func (r *ForwardingRule) Network() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["network"])
}

// The networking tier used for configuring this address. This field can take the following values: PREMIUM or STANDARD. If
// this field is not specified, it is assumed to be PREMIUM.
func (r *ForwardingRule) NetworkTier() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["networkTier"])
}

// This field is used along with the target field for TargetHttpProxy, TargetHttpsProxy, TargetSslProxy, TargetTcpProxy,
// TargetVpnGateway, TargetPool, TargetInstance. Applicable only when IPProtocol is TCP, UDP, or SCTP, only packets
// addressed to ports in the specified range will be forwarded to target. Forwarding rules with the same [IPAddress,
// IPProtocol] pair must have disjoint port ranges. Some types of forwarding target have constraints on the acceptable
// ports: * TargetHttpProxy: 80, 8080 * TargetHttpsProxy: 443 * TargetTcpProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700,
// 993, 995, 1883, 5222 * TargetSslProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995, 1883, 5222 *
// TargetVpnGateway: 500, 4500
func (r *ForwardingRule) PortRange() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["portRange"])
}

// This field is used along with the backend_service field for internal load balancing. When the load balancing scheme is
// INTERNAL, a single port or a comma separated list of ports can be configured. Only packets addressed to these ports will
// be forwarded to the backends configured with this forwarding rule. You may specify a maximum of up to 5 ports.
func (r *ForwardingRule) Ports() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["ports"])
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (r *ForwardingRule) Project() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["project"])
}

// A reference to the region where the regional forwarding rule resides. This field is not applicable to global forwarding
// rules.
func (r *ForwardingRule) Region() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["region"])
}

// The URI of the created resource.
func (r *ForwardingRule) SelfLink() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["selfLink"])
}

// An optional prefix to the service name for this Forwarding Rule. If specified, will be the first label of the fully
// qualified service name. The label must be 1-63 characters long, and comply with RFC1035. Specifically, the label must be
// 1-63 characters long and match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must
// be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last
// character, which cannot be a dash. This field is only used for INTERNAL load balancing.
func (r *ForwardingRule) ServiceLabel() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["serviceLabel"])
}

// The internal fully qualified service name for this Forwarding Rule. This field is only used for INTERNAL load balancing.
func (r *ForwardingRule) ServiceName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["serviceName"])
}

// The subnetwork that the load balanced IP should belong to for this Forwarding Rule. This field is only used for INTERNAL
// load balancing. If the network specified is in auto subnet mode, this field is optional. However, if the network is in
// custom subnet mode, a subnetwork must be specified.
func (r *ForwardingRule) Subnetwork() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["subnetwork"])
}

// This field is only used for EXTERNAL load balancing. A reference to a TargetPool resource to receive the matched
// traffic. This target must live in the same region as the forwarding rule. The forwarded traffic must be of a type
// appropriate to the target object.
func (r *ForwardingRule) Target() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["target"])
}

// Input properties used for looking up and filtering ForwardingRule resources.
type ForwardingRuleState struct {
	// For internal TCP/UDP load balancing (i.e. load balancing scheme is INTERNAL and protocol is TCP/UDP), set this to true
	// to allow packets addressed to any ports to be forwarded to the backends configured with this forwarding rule. Used with
	// backend service. Cannot be set if port or portRange are set.
	AllPorts interface{}
	// A BackendService to receive the matched traffic. This is used only for INTERNAL load balancing.
	BackendService interface{}
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp interface{}
	// An optional description of this resource. Provide this property when you create the resource.
	Description interface{}
	// The IP address that this forwarding rule is serving on behalf of. Addresses are restricted based on the forwarding
	// rule's load balancing scheme (EXTERNAL or INTERNAL) and scope (global or regional). When the load balancing scheme is
	// EXTERNAL, for global forwarding rules, the address must be a global IP, and for regional forwarding rules, the address
	// must live in the same region as the forwarding rule. If this field is empty, an ephemeral IPv4 address from the same
	// scope (global or regional) will be assigned. A regional forwarding rule supports IPv4 only. A global forwarding rule
	// supports either IPv4 or IPv6. When the load balancing scheme is INTERNAL, this can only be an RFC 1918 IP address
	// belonging to the network/subnet configured for the forwarding rule. By default, if this field is empty, an ephemeral
	// internal IP address will be automatically allocated from the IP range of the subnet or network configured for this
	// forwarding rule. An address must be specified by a literal IP address. ~> **NOTE**: While the API allows you to specify
	// various resource paths for an address resource instead, Terraform requires this to specifically be an IP address to
	// avoid needing to fetching the IP address from resource paths on refresh or unnecessary diffs.
	IpAddress interface{}
	// The IP protocol to which this rule applies. Valid options are TCP, UDP, ESP, AH, SCTP or ICMP. When the load balancing
	// scheme is INTERNAL, only TCP and UDP are valid.
	IpProtocol interface{}
	// The fingerprint used for optimistic locking of this resource. Used internally during updates.
	LabelFingerprint interface{}
	// Labels to apply to this forwarding rule. A list of key->value pairs.
	Labels interface{}
	// This signifies what the ForwardingRule will be used for and can be EXTERNAL, INTERNAL, or INTERNAL_MANAGED. EXTERNAL is
	// used for Classic Cloud VPN gateways, protocol forwarding to VMs from an external IP address, and HTTP(S), SSL Proxy,
	// TCP Proxy, and Network TCP/UDP load balancers. INTERNAL is used for protocol forwarding to VMs from an internal IP
	// address, and internal TCP/UDP load balancers. INTERNAL_MANAGED is used for internal HTTP(S) load balancers.
	LoadBalancingScheme interface{}
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
	// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name interface{}
	// For internal load balancing, this field identifies the network that the load balanced IP should belong to for this
	// Forwarding Rule. If this field is not specified, the default network will be used. This field is only used for INTERNAL
	// load balancing.
	Network interface{}
	// The networking tier used for configuring this address. This field can take the following values: PREMIUM or STANDARD.
	// If this field is not specified, it is assumed to be PREMIUM.
	NetworkTier interface{}
	// This field is used along with the target field for TargetHttpProxy, TargetHttpsProxy, TargetSslProxy, TargetTcpProxy,
	// TargetVpnGateway, TargetPool, TargetInstance. Applicable only when IPProtocol is TCP, UDP, or SCTP, only packets
	// addressed to ports in the specified range will be forwarded to target. Forwarding rules with the same [IPAddress,
	// IPProtocol] pair must have disjoint port ranges. Some types of forwarding target have constraints on the acceptable
	// ports: * TargetHttpProxy: 80, 8080 * TargetHttpsProxy: 443 * TargetTcpProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700,
	// 993, 995, 1883, 5222 * TargetSslProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995, 1883, 5222 *
	// TargetVpnGateway: 500, 4500
	PortRange interface{}
	// This field is used along with the backend_service field for internal load balancing. When the load balancing scheme is
	// INTERNAL, a single port or a comma separated list of ports can be configured. Only packets addressed to these ports
	// will be forwarded to the backends configured with this forwarding rule. You may specify a maximum of up to 5 ports.
	Ports interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	// A reference to the region where the regional forwarding rule resides. This field is not applicable to global forwarding
	// rules.
	Region interface{}
	// The URI of the created resource.
	SelfLink interface{}
	// An optional prefix to the service name for this Forwarding Rule. If specified, will be the first label of the fully
	// qualified service name. The label must be 1-63 characters long, and comply with RFC1035. Specifically, the label must
	// be 1-63 characters long and match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character
	// must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash. This field is only used for INTERNAL load balancing.
	ServiceLabel interface{}
	// The internal fully qualified service name for this Forwarding Rule. This field is only used for INTERNAL load
	// balancing.
	ServiceName interface{}
	// The subnetwork that the load balanced IP should belong to for this Forwarding Rule. This field is only used for
	// INTERNAL load balancing. If the network specified is in auto subnet mode, this field is optional. However, if the
	// network is in custom subnet mode, a subnetwork must be specified.
	Subnetwork interface{}
	// This field is only used for EXTERNAL load balancing. A reference to a TargetPool resource to receive the matched
	// traffic. This target must live in the same region as the forwarding rule. The forwarded traffic must be of a type
	// appropriate to the target object.
	Target interface{}
}

// The set of arguments for constructing a ForwardingRule resource.
type ForwardingRuleArgs struct {
	// For internal TCP/UDP load balancing (i.e. load balancing scheme is INTERNAL and protocol is TCP/UDP), set this to true
	// to allow packets addressed to any ports to be forwarded to the backends configured with this forwarding rule. Used with
	// backend service. Cannot be set if port or portRange are set.
	AllPorts interface{}
	// A BackendService to receive the matched traffic. This is used only for INTERNAL load balancing.
	BackendService interface{}
	// An optional description of this resource. Provide this property when you create the resource.
	Description interface{}
	// The IP address that this forwarding rule is serving on behalf of. Addresses are restricted based on the forwarding
	// rule's load balancing scheme (EXTERNAL or INTERNAL) and scope (global or regional). When the load balancing scheme is
	// EXTERNAL, for global forwarding rules, the address must be a global IP, and for regional forwarding rules, the address
	// must live in the same region as the forwarding rule. If this field is empty, an ephemeral IPv4 address from the same
	// scope (global or regional) will be assigned. A regional forwarding rule supports IPv4 only. A global forwarding rule
	// supports either IPv4 or IPv6. When the load balancing scheme is INTERNAL, this can only be an RFC 1918 IP address
	// belonging to the network/subnet configured for the forwarding rule. By default, if this field is empty, an ephemeral
	// internal IP address will be automatically allocated from the IP range of the subnet or network configured for this
	// forwarding rule. An address must be specified by a literal IP address. ~> **NOTE**: While the API allows you to specify
	// various resource paths for an address resource instead, Terraform requires this to specifically be an IP address to
	// avoid needing to fetching the IP address from resource paths on refresh or unnecessary diffs.
	IpAddress interface{}
	// The IP protocol to which this rule applies. Valid options are TCP, UDP, ESP, AH, SCTP or ICMP. When the load balancing
	// scheme is INTERNAL, only TCP and UDP are valid.
	IpProtocol interface{}
	// Labels to apply to this forwarding rule. A list of key->value pairs.
	Labels interface{}
	// This signifies what the ForwardingRule will be used for and can be EXTERNAL, INTERNAL, or INTERNAL_MANAGED. EXTERNAL is
	// used for Classic Cloud VPN gateways, protocol forwarding to VMs from an external IP address, and HTTP(S), SSL Proxy,
	// TCP Proxy, and Network TCP/UDP load balancers. INTERNAL is used for protocol forwarding to VMs from an internal IP
	// address, and internal TCP/UDP load balancers. INTERNAL_MANAGED is used for internal HTTP(S) load balancers.
	LoadBalancingScheme interface{}
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
	// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name interface{}
	// For internal load balancing, this field identifies the network that the load balanced IP should belong to for this
	// Forwarding Rule. If this field is not specified, the default network will be used. This field is only used for INTERNAL
	// load balancing.
	Network interface{}
	// The networking tier used for configuring this address. This field can take the following values: PREMIUM or STANDARD.
	// If this field is not specified, it is assumed to be PREMIUM.
	NetworkTier interface{}
	// This field is used along with the target field for TargetHttpProxy, TargetHttpsProxy, TargetSslProxy, TargetTcpProxy,
	// TargetVpnGateway, TargetPool, TargetInstance. Applicable only when IPProtocol is TCP, UDP, or SCTP, only packets
	// addressed to ports in the specified range will be forwarded to target. Forwarding rules with the same [IPAddress,
	// IPProtocol] pair must have disjoint port ranges. Some types of forwarding target have constraints on the acceptable
	// ports: * TargetHttpProxy: 80, 8080 * TargetHttpsProxy: 443 * TargetTcpProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700,
	// 993, 995, 1883, 5222 * TargetSslProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995, 1883, 5222 *
	// TargetVpnGateway: 500, 4500
	PortRange interface{}
	// This field is used along with the backend_service field for internal load balancing. When the load balancing scheme is
	// INTERNAL, a single port or a comma separated list of ports can be configured. Only packets addressed to these ports
	// will be forwarded to the backends configured with this forwarding rule. You may specify a maximum of up to 5 ports.
	Ports interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	// A reference to the region where the regional forwarding rule resides. This field is not applicable to global forwarding
	// rules.
	Region interface{}
	// An optional prefix to the service name for this Forwarding Rule. If specified, will be the first label of the fully
	// qualified service name. The label must be 1-63 characters long, and comply with RFC1035. Specifically, the label must
	// be 1-63 characters long and match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character
	// must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash. This field is only used for INTERNAL load balancing.
	ServiceLabel interface{}
	// The subnetwork that the load balanced IP should belong to for this Forwarding Rule. This field is only used for
	// INTERNAL load balancing. If the network specified is in auto subnet mode, this field is optional. However, if the
	// network is in custom subnet mode, a subnetwork must be specified.
	Subnetwork interface{}
	// This field is only used for EXTERNAL load balancing. A reference to a TargetPool resource to receive the matched
	// traffic. This target must live in the same region as the forwarding rule. The forwarded traffic must be of a type
	// appropriate to the target object.
	Target interface{}
}
