// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.AppEngine.Inputs
{

    public sealed class FlexibleAppVersionApiConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Action to take when users access resources that require authentication.
        /// </summary>
        [Input("authFailAction")]
        public Input<string>? AuthFailAction { get; set; }

        /// <summary>
        /// Level of login required to access this resource.
        /// </summary>
        [Input("login")]
        public Input<string>? Login { get; set; }

        /// <summary>
        /// Path to the script from the application root directory.
        /// </summary>
        [Input("script", required: true)]
        public Input<string> Script { get; set; } = null!;

        /// <summary>
        /// Security (HTTPS) enforcement for this URL.
        /// </summary>
        [Input("securityLevel")]
        public Input<string>? SecurityLevel { get; set; }

        /// <summary>
        /// URL to serve the endpoint at.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public FlexibleAppVersionApiConfigArgs()
        {
        }
    }
}