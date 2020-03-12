// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to access IP ranges in your firewall rules.
        /// 
        /// https://cloud.google.com/compute/docs/load-balancing/health-checks#health_check_source_ips_and_firewall_rules
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/datasource_compute_lb_ip_ranges.html.markdown.
        /// </summary>
        public static Task<GetLBIPRangesResult> GetLBIPRanges(InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetLBIPRangesResult>("gcp:compute/getLBIPRanges:getLBIPRanges", InvokeArgs.Empty, options.WithVersion());
    }

    [OutputType]
    public sealed class GetLBIPRangesResult
    {
        /// <summary>
        /// The IP ranges used for health checks when **HTTP(S), SSL proxy, TCP proxy, and Internal load balancing** is used
        /// </summary>
        public readonly ImmutableArray<string> HttpSslTcpInternals;
        /// <summary>
        /// The IP ranges used for health checks when **Network load balancing** is used
        /// </summary>
        public readonly ImmutableArray<string> Networks;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetLBIPRangesResult(
            ImmutableArray<string> httpSslTcpInternals,
            ImmutableArray<string> networks,
            string id)
        {
            HttpSslTcpInternals = httpSslTcpInternals;
            Networks = networks;
            Id = id;
        }
    }
}
