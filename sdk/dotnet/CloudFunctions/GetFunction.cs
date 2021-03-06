// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CloudFunctions
{
    public static class GetFunction
    {
        /// <summary>
        /// Get information about a Google Cloud Function. For more information see
        /// the [official documentation](https://cloud.google.com/functions/docs/)
        /// and [API](https://cloud.google.com/functions/docs/apis).
        /// 
        /// {{% examples %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetFunctionResult> InvokeAsync(GetFunctionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetFunctionResult>("gcp:cloudfunctions/getFunction:getFunction", args ?? new GetFunctionArgs(), options.WithVersion());
    }


    public sealed class GetFunctionArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of a Cloud Function.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public string? Project { get; set; }

        /// <summary>
        /// The region in which the resource belongs. If it
        /// is not provided, the provider region is used.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        public GetFunctionArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetFunctionResult
    {
        /// <summary>
        /// Available memory (in MB) to the function.
        /// </summary>
        public readonly int AvailableMemoryMb;
        /// <summary>
        /// Description of the function.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Name of a JavaScript function that will be executed when the Google Cloud Function is triggered.
        /// </summary>
        public readonly string EntryPoint;
        public readonly ImmutableDictionary<string, object> EnvironmentVariables;
        /// <summary>
        /// A source that fires events in response to a condition in another service. Structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFunctionEventTriggerResult> EventTriggers;
        /// <summary>
        /// If function is triggered by HTTP, trigger URL is set here.
        /// </summary>
        public readonly string HttpsTriggerUrl;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Controls what traffic can reach the function.
        /// </summary>
        public readonly string IngressSettings;
        /// <summary>
        /// A map of labels applied to this function.
        /// </summary>
        public readonly ImmutableDictionary<string, object> Labels;
        public readonly int MaxInstances;
        /// <summary>
        /// The name of the Cloud Function.
        /// </summary>
        public readonly string Name;
        public readonly string? Project;
        public readonly string? Region;
        /// <summary>
        /// The runtime in which the function is running.
        /// </summary>
        public readonly string Runtime;
        /// <summary>
        /// The service account email to be assumed by the cloud function.
        /// </summary>
        public readonly string ServiceAccountEmail;
        /// <summary>
        /// The GCS bucket containing the zip archive which contains the function.
        /// </summary>
        public readonly string SourceArchiveBucket;
        /// <summary>
        /// The source archive object (file) in archive bucket.
        /// </summary>
        public readonly string SourceArchiveObject;
        public readonly ImmutableArray<Outputs.GetFunctionSourceRepositoryResult> SourceRepositories;
        /// <summary>
        /// Function execution timeout (in seconds).
        /// </summary>
        public readonly int Timeout;
        /// <summary>
        /// If function is triggered by HTTP, this boolean is set.
        /// </summary>
        public readonly bool TriggerHttp;
        /// <summary>
        /// The VPC Network Connector that this cloud function can connect to. 
        /// </summary>
        public readonly string VpcConnector;
        /// <summary>
        /// The egress settings for the connector, controlling what traffic is diverted through it.
        /// </summary>
        public readonly string VpcConnectorEgressSettings;

        [OutputConstructor]
        private GetFunctionResult(
            int availableMemoryMb,

            string description,

            string entryPoint,

            ImmutableDictionary<string, object> environmentVariables,

            ImmutableArray<Outputs.GetFunctionEventTriggerResult> eventTriggers,

            string httpsTriggerUrl,

            string id,

            string ingressSettings,

            ImmutableDictionary<string, object> labels,

            int maxInstances,

            string name,

            string? project,

            string? region,

            string runtime,

            string serviceAccountEmail,

            string sourceArchiveBucket,

            string sourceArchiveObject,

            ImmutableArray<Outputs.GetFunctionSourceRepositoryResult> sourceRepositories,

            int timeout,

            bool triggerHttp,

            string vpcConnector,

            string vpcConnectorEgressSettings)
        {
            AvailableMemoryMb = availableMemoryMb;
            Description = description;
            EntryPoint = entryPoint;
            EnvironmentVariables = environmentVariables;
            EventTriggers = eventTriggers;
            HttpsTriggerUrl = httpsTriggerUrl;
            Id = id;
            IngressSettings = ingressSettings;
            Labels = labels;
            MaxInstances = maxInstances;
            Name = name;
            Project = project;
            Region = region;
            Runtime = runtime;
            ServiceAccountEmail = serviceAccountEmail;
            SourceArchiveBucket = sourceArchiveBucket;
            SourceArchiveObject = sourceArchiveObject;
            SourceRepositories = sourceRepositories;
            Timeout = timeout;
            TriggerHttp = triggerHttp;
            VpcConnector = vpcConnector;
            VpcConnectorEgressSettings = vpcConnectorEgressSettings;
        }
    }
}
