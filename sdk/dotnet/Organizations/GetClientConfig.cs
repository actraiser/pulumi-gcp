// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Organizations
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to access the configuration of the Google Cloud provider.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/client_config.html.markdown.
        /// </summary>
        public static Task<GetClientConfigResult> GetClientConfig(InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetClientConfigResult>("gcp:organizations/getClientConfig:getClientConfig", InvokeArgs.Empty, options.WithVersion());
    }

    [OutputType]
    public sealed class GetClientConfigResult
    {
        /// <summary>
        /// The OAuth2 access token used by the client to authenticate against the Google Cloud API.
        /// </summary>
        public readonly string AccessToken;
        /// <summary>
        /// The ID of the project to apply any resources to.
        /// </summary>
        public readonly string Project;
        /// <summary>
        /// The region to operate under.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// The zone to operate under.
        /// </summary>
        public readonly string Zone;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetClientConfigResult(
            string accessToken,
            string project,
            string region,
            string zone,
            string id)
        {
            AccessToken = accessToken;
            Project = project;
            Region = region;
            Zone = zone;
            Id = id;
        }
    }
}