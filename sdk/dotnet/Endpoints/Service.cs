// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Endpoints
{
    /// <summary>
    /// This resource creates and rolls out a Cloud Endpoints service using OpenAPI or gRPC.  View the relevant docs for [OpenAPI](https://cloud.google.com/endpoints/docs/openapi/) and [gRPC](https://cloud.google.com/endpoints/docs/grpc/).
    /// </summary>
    public partial class Service : Pulumi.CustomResource
    {
        [Output("apis")]
        public Output<ImmutableArray<Outputs.ServiceApi>> Apis { get; private set; } = null!;

        [Output("configId")]
        public Output<string> ConfigId { get; private set; } = null!;

        [Output("dnsAddress")]
        public Output<string> DnsAddress { get; private set; } = null!;

        [Output("endpoints")]
        public Output<ImmutableArray<Outputs.ServiceEndpoint>> Endpoints { get; private set; } = null!;

        [Output("grpcConfig")]
        public Output<string?> GrpcConfig { get; private set; } = null!;

        [Output("openapiConfig")]
        public Output<string?> OpenapiConfig { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        [Output("protocOutputBase64")]
        public Output<string?> ProtocOutputBase64 { get; private set; } = null!;

        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;


        /// <summary>
        /// Create a Service resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Service(string name, ServiceArgs args, CustomResourceOptions? options = null)
            : base("gcp:endpoints/service:Service", name, args ?? new ServiceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Service(string name, Input<string> id, ServiceState? state = null, CustomResourceOptions? options = null)
            : base("gcp:endpoints/service:Service", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Service resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Service Get(string name, Input<string> id, ServiceState? state = null, CustomResourceOptions? options = null)
        {
            return new Service(name, id, state, options);
        }
    }

    public sealed class ServiceArgs : Pulumi.ResourceArgs
    {
        [Input("grpcConfig")]
        public Input<string>? GrpcConfig { get; set; }

        [Input("openapiConfig")]
        public Input<string>? OpenapiConfig { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("protocOutputBase64")]
        public Input<string>? ProtocOutputBase64 { get; set; }

        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public ServiceArgs()
        {
        }
    }

    public sealed class ServiceState : Pulumi.ResourceArgs
    {
        [Input("apis")]
        private InputList<Inputs.ServiceApiGetArgs>? _apis;
        public InputList<Inputs.ServiceApiGetArgs> Apis
        {
            get => _apis ?? (_apis = new InputList<Inputs.ServiceApiGetArgs>());
            set => _apis = value;
        }

        [Input("configId")]
        public Input<string>? ConfigId { get; set; }

        [Input("dnsAddress")]
        public Input<string>? DnsAddress { get; set; }

        [Input("endpoints")]
        private InputList<Inputs.ServiceEndpointGetArgs>? _endpoints;
        public InputList<Inputs.ServiceEndpointGetArgs> Endpoints
        {
            get => _endpoints ?? (_endpoints = new InputList<Inputs.ServiceEndpointGetArgs>());
            set => _endpoints = value;
        }

        [Input("grpcConfig")]
        public Input<string>? GrpcConfig { get; set; }

        [Input("openapiConfig")]
        public Input<string>? OpenapiConfig { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("protocOutputBase64")]
        public Input<string>? ProtocOutputBase64 { get; set; }

        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        public ServiceState()
        {
        }
    }
}
