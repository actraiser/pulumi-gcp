// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Container.Inputs
{

    public sealed class ClusterNodePoolManagementGetArgs : Pulumi.ResourceArgs
    {
        [Input("autoRepair")]
        public Input<bool>? AutoRepair { get; set; }

        [Input("autoUpgrade")]
        public Input<bool>? AutoUpgrade { get; set; }

        public ClusterNodePoolManagementGetArgs()
        {
        }
    }
}
