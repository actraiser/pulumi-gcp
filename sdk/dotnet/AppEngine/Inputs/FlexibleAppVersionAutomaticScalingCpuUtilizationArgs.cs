// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.AppEngine.Inputs
{

    public sealed class FlexibleAppVersionAutomaticScalingCpuUtilizationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Period of time over which CPU utilization is calculated.
        /// </summary>
        [Input("aggregationWindowLength")]
        public Input<string>? AggregationWindowLength { get; set; }

        /// <summary>
        /// Target CPU utilization ratio to maintain when scaling. Must be between 0 and 1.
        /// </summary>
        [Input("targetUtilization", required: true)]
        public Input<double> TargetUtilization { get; set; } = null!;

        public FlexibleAppVersionAutomaticScalingCpuUtilizationArgs()
        {
        }
    }
}