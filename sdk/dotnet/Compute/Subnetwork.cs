// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute
{
    public partial class Subnetwork : Pulumi.CustomResource
    {
        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        [Output("creationTimestamp")]
        public Output<string> CreationTimestamp { get; private set; } = null!;

        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource. This field can
        /// be set only at resource creation time.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Fingerprint of this resource. This field is used internally during updates of this resource.
        /// </summary>
        [Output("fingerprint")]
        public Output<string> Fingerprint { get; private set; } = null!;

        /// <summary>
        /// The gateway address for default routes to reach destination addresses outside this subnetwork.
        /// </summary>
        [Output("gatewayAddress")]
        public Output<string> GatewayAddress { get; private set; } = null!;

        /// <summary>
        /// The range of internal addresses that are owned by this subnetwork. Provide this property when you create the
        /// subnetwork. For example, 10.0.0.0/8 or 192.168.0.0/16. Ranges must be unique and non-overlapping within a
        /// network. Only IPv4 is supported.
        /// </summary>
        [Output("ipCidrRange")]
        public Output<string> IpCidrRange { get; private set; } = null!;

        /// <summary>
        /// Denotes the logging options for the subnetwork flow logs. If logging is enabled logs will be exported to
        /// Stackdriver. This field cannot be set if the 'purpose' of this subnetwork is 'INTERNAL_HTTPS_LOAD_BALANCER'
        /// </summary>
        [Output("logConfig")]
        public Output<Outputs.SubnetworkLogConfig?> LogConfig { get; private set; } = null!;

        /// <summary>
        /// The name of the resource, provided by the client when initially creating the resource. The name must be 1-63
        /// characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the
        /// regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter,
        /// and all following characters must be a dash, lowercase letter, or digit, except the last character, which
        /// cannot be a dash.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The network this subnet belongs to. Only networks that are in the distributed mode can have subnetworks.
        /// </summary>
        [Output("network")]
        public Output<string> Network { get; private set; } = null!;

        /// <summary>
        /// When enabled, VMs in this subnetwork without external IP addresses can access Google APIs and services by
        /// using Private Google Access.
        /// </summary>
        [Output("privateIpGoogleAccess")]
        public Output<bool?> PrivateIpGoogleAccess { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The purpose of the resource. This field can be either PRIVATE or INTERNAL_HTTPS_LOAD_BALANCER. A subnetwork
        /// with purpose set to INTERNAL_HTTPS_LOAD_BALANCER is a user-created subnetwork that is reserved for Internal
        /// HTTP(S) Load Balancing. If unspecified, the purpose defaults to PRIVATE. If set to
        /// INTERNAL_HTTPS_LOAD_BALANCER you must also set the role.
        /// </summary>
        [Output("purpose")]
        public Output<string> Purpose { get; private set; } = null!;

        /// <summary>
        /// URL of the GCP region for this subnetwork.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The role of subnetwork. Currently, this field is only used when purpose = INTERNAL_HTTPS_LOAD_BALANCER. The
        /// value can be set to ACTIVE or BACKUP. An ACTIVE subnetwork is one that is currently being used for Internal
        /// HTTP(S) Load Balancing. A BACKUP subnetwork is one that is ready to be promoted to ACTIVE or is currently
        /// draining.
        /// </summary>
        [Output("role")]
        public Output<string?> Role { get; private set; } = null!;

        /// <summary>
        /// An array of configurations for secondary IP ranges for VM instances contained in this subnetwork. The
        /// primary IP of such VM must belong to the primary ipCidrRange of the subnetwork. The alias IPs may belong to
        /// either primary or secondary ranges. **Note**: This field uses [attr-as-block
        /// mode](https://www.terraform.io/docs/configuration/attr-as-blocks.html) to avoid breaking users during the
        /// 0.12 upgrade. To explicitly send a list of zero objects you must use the following syntax: 'example=[]' For
        /// more details about this behavior, see [this
        /// section](https://www.terraform.io/docs/configuration/attr-as-blocks.html#defining-a-fixed-object-collection-value).
        /// </summary>
        [Output("secondaryIpRanges")]
        public Output<ImmutableArray<Outputs.SubnetworkSecondaryIpRanges>> SecondaryIpRanges { get; private set; } = null!;

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;


        /// <summary>
        /// Create a Subnetwork resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Subnetwork(string name, SubnetworkArgs args, CustomResourceOptions? options = null)
            : base("gcp:compute/subnetwork:Subnetwork", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Subnetwork(string name, Input<string> id, SubnetworkState? state = null, CustomResourceOptions? options = null)
            : base("gcp:compute/subnetwork:Subnetwork", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Subnetwork resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Subnetwork Get(string name, Input<string> id, SubnetworkState? state = null, CustomResourceOptions? options = null)
        {
            return new Subnetwork(name, id, state, options);
        }
    }

    public sealed class SubnetworkArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource. This field can
        /// be set only at resource creation time.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The range of internal addresses that are owned by this subnetwork. Provide this property when you create the
        /// subnetwork. For example, 10.0.0.0/8 or 192.168.0.0/16. Ranges must be unique and non-overlapping within a
        /// network. Only IPv4 is supported.
        /// </summary>
        [Input("ipCidrRange", required: true)]
        public Input<string> IpCidrRange { get; set; } = null!;

        /// <summary>
        /// Denotes the logging options for the subnetwork flow logs. If logging is enabled logs will be exported to
        /// Stackdriver. This field cannot be set if the 'purpose' of this subnetwork is 'INTERNAL_HTTPS_LOAD_BALANCER'
        /// </summary>
        [Input("logConfig")]
        public Input<Inputs.SubnetworkLogConfigArgs>? LogConfig { get; set; }

        /// <summary>
        /// The name of the resource, provided by the client when initially creating the resource. The name must be 1-63
        /// characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the
        /// regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter,
        /// and all following characters must be a dash, lowercase letter, or digit, except the last character, which
        /// cannot be a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The network this subnet belongs to. Only networks that are in the distributed mode can have subnetworks.
        /// </summary>
        [Input("network", required: true)]
        public Input<string> Network { get; set; } = null!;

        /// <summary>
        /// When enabled, VMs in this subnetwork without external IP addresses can access Google APIs and services by
        /// using Private Google Access.
        /// </summary>
        [Input("privateIpGoogleAccess")]
        public Input<bool>? PrivateIpGoogleAccess { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The purpose of the resource. This field can be either PRIVATE or INTERNAL_HTTPS_LOAD_BALANCER. A subnetwork
        /// with purpose set to INTERNAL_HTTPS_LOAD_BALANCER is a user-created subnetwork that is reserved for Internal
        /// HTTP(S) Load Balancing. If unspecified, the purpose defaults to PRIVATE. If set to
        /// INTERNAL_HTTPS_LOAD_BALANCER you must also set the role.
        /// </summary>
        [Input("purpose")]
        public Input<string>? Purpose { get; set; }

        /// <summary>
        /// URL of the GCP region for this subnetwork.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The role of subnetwork. Currently, this field is only used when purpose = INTERNAL_HTTPS_LOAD_BALANCER. The
        /// value can be set to ACTIVE or BACKUP. An ACTIVE subnetwork is one that is currently being used for Internal
        /// HTTP(S) Load Balancing. A BACKUP subnetwork is one that is ready to be promoted to ACTIVE or is currently
        /// draining.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        [Input("secondaryIpRanges")]
        private InputList<Inputs.SubnetworkSecondaryIpRangesArgs>? _secondaryIpRanges;

        /// <summary>
        /// An array of configurations for secondary IP ranges for VM instances contained in this subnetwork. The
        /// primary IP of such VM must belong to the primary ipCidrRange of the subnetwork. The alias IPs may belong to
        /// either primary or secondary ranges. **Note**: This field uses [attr-as-block
        /// mode](https://www.terraform.io/docs/configuration/attr-as-blocks.html) to avoid breaking users during the
        /// 0.12 upgrade. To explicitly send a list of zero objects you must use the following syntax: 'example=[]' For
        /// more details about this behavior, see [this
        /// section](https://www.terraform.io/docs/configuration/attr-as-blocks.html#defining-a-fixed-object-collection-value).
        /// </summary>
        public InputList<Inputs.SubnetworkSecondaryIpRangesArgs> SecondaryIpRanges
        {
            get => _secondaryIpRanges ?? (_secondaryIpRanges = new InputList<Inputs.SubnetworkSecondaryIpRangesArgs>());
            set => _secondaryIpRanges = value;
        }

        public SubnetworkArgs()
        {
        }
    }

    public sealed class SubnetworkState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        [Input("creationTimestamp")]
        public Input<string>? CreationTimestamp { get; set; }

        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource. This field can
        /// be set only at resource creation time.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Fingerprint of this resource. This field is used internally during updates of this resource.
        /// </summary>
        [Input("fingerprint")]
        public Input<string>? Fingerprint { get; set; }

        /// <summary>
        /// The gateway address for default routes to reach destination addresses outside this subnetwork.
        /// </summary>
        [Input("gatewayAddress")]
        public Input<string>? GatewayAddress { get; set; }

        /// <summary>
        /// The range of internal addresses that are owned by this subnetwork. Provide this property when you create the
        /// subnetwork. For example, 10.0.0.0/8 or 192.168.0.0/16. Ranges must be unique and non-overlapping within a
        /// network. Only IPv4 is supported.
        /// </summary>
        [Input("ipCidrRange")]
        public Input<string>? IpCidrRange { get; set; }

        /// <summary>
        /// Denotes the logging options for the subnetwork flow logs. If logging is enabled logs will be exported to
        /// Stackdriver. This field cannot be set if the 'purpose' of this subnetwork is 'INTERNAL_HTTPS_LOAD_BALANCER'
        /// </summary>
        [Input("logConfig")]
        public Input<Inputs.SubnetworkLogConfigGetArgs>? LogConfig { get; set; }

        /// <summary>
        /// The name of the resource, provided by the client when initially creating the resource. The name must be 1-63
        /// characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the
        /// regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter,
        /// and all following characters must be a dash, lowercase letter, or digit, except the last character, which
        /// cannot be a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The network this subnet belongs to. Only networks that are in the distributed mode can have subnetworks.
        /// </summary>
        [Input("network")]
        public Input<string>? Network { get; set; }

        /// <summary>
        /// When enabled, VMs in this subnetwork without external IP addresses can access Google APIs and services by
        /// using Private Google Access.
        /// </summary>
        [Input("privateIpGoogleAccess")]
        public Input<bool>? PrivateIpGoogleAccess { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The purpose of the resource. This field can be either PRIVATE or INTERNAL_HTTPS_LOAD_BALANCER. A subnetwork
        /// with purpose set to INTERNAL_HTTPS_LOAD_BALANCER is a user-created subnetwork that is reserved for Internal
        /// HTTP(S) Load Balancing. If unspecified, the purpose defaults to PRIVATE. If set to
        /// INTERNAL_HTTPS_LOAD_BALANCER you must also set the role.
        /// </summary>
        [Input("purpose")]
        public Input<string>? Purpose { get; set; }

        /// <summary>
        /// URL of the GCP region for this subnetwork.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The role of subnetwork. Currently, this field is only used when purpose = INTERNAL_HTTPS_LOAD_BALANCER. The
        /// value can be set to ACTIVE or BACKUP. An ACTIVE subnetwork is one that is currently being used for Internal
        /// HTTP(S) Load Balancing. A BACKUP subnetwork is one that is ready to be promoted to ACTIVE or is currently
        /// draining.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        [Input("secondaryIpRanges")]
        private InputList<Inputs.SubnetworkSecondaryIpRangesGetArgs>? _secondaryIpRanges;

        /// <summary>
        /// An array of configurations for secondary IP ranges for VM instances contained in this subnetwork. The
        /// primary IP of such VM must belong to the primary ipCidrRange of the subnetwork. The alias IPs may belong to
        /// either primary or secondary ranges. **Note**: This field uses [attr-as-block
        /// mode](https://www.terraform.io/docs/configuration/attr-as-blocks.html) to avoid breaking users during the
        /// 0.12 upgrade. To explicitly send a list of zero objects you must use the following syntax: 'example=[]' For
        /// more details about this behavior, see [this
        /// section](https://www.terraform.io/docs/configuration/attr-as-blocks.html#defining-a-fixed-object-collection-value).
        /// </summary>
        public InputList<Inputs.SubnetworkSecondaryIpRangesGetArgs> SecondaryIpRanges
        {
            get => _secondaryIpRanges ?? (_secondaryIpRanges = new InputList<Inputs.SubnetworkSecondaryIpRangesGetArgs>());
            set => _secondaryIpRanges = value;
        }

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        public SubnetworkState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class SubnetworkLogConfigArgs : Pulumi.ResourceArgs
    {
        [Input("aggregationInterval")]
        public Input<string>? AggregationInterval { get; set; }

        [Input("flowSampling")]
        public Input<double>? FlowSampling { get; set; }

        [Input("metadata")]
        public Input<string>? Metadata { get; set; }

        public SubnetworkLogConfigArgs()
        {
        }
    }

    public sealed class SubnetworkLogConfigGetArgs : Pulumi.ResourceArgs
    {
        [Input("aggregationInterval")]
        public Input<string>? AggregationInterval { get; set; }

        [Input("flowSampling")]
        public Input<double>? FlowSampling { get; set; }

        [Input("metadata")]
        public Input<string>? Metadata { get; set; }

        public SubnetworkLogConfigGetArgs()
        {
        }
    }

    public sealed class SubnetworkSecondaryIpRangesArgs : Pulumi.ResourceArgs
    {
        [Input("ipCidrRange", required: true)]
        public Input<string> IpCidrRange { get; set; } = null!;

        [Input("rangeName", required: true)]
        public Input<string> RangeName { get; set; } = null!;

        public SubnetworkSecondaryIpRangesArgs()
        {
        }
    }

    public sealed class SubnetworkSecondaryIpRangesGetArgs : Pulumi.ResourceArgs
    {
        [Input("ipCidrRange", required: true)]
        public Input<string> IpCidrRange { get; set; } = null!;

        [Input("rangeName", required: true)]
        public Input<string> RangeName { get; set; } = null!;

        public SubnetworkSecondaryIpRangesGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class SubnetworkLogConfig
    {
        public readonly string? AggregationInterval;
        public readonly double? FlowSampling;
        public readonly string? Metadata;

        [OutputConstructor]
        private SubnetworkLogConfig(
            string? aggregationInterval,
            double? flowSampling,
            string? metadata)
        {
            AggregationInterval = aggregationInterval;
            FlowSampling = flowSampling;
            Metadata = metadata;
        }
    }

    [OutputType]
    public sealed class SubnetworkSecondaryIpRanges
    {
        public readonly string IpCidrRange;
        public readonly string RangeName;

        [OutputConstructor]
        private SubnetworkSecondaryIpRanges(
            string ipCidrRange,
            string rangeName)
        {
            IpCidrRange = ipCidrRange;
            RangeName = rangeName;
        }
    }
    }
}
