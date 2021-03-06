// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Monitoring.Inputs
{

    public sealed class NotificationChannelSensitiveLabelsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An authorization token for a notification channel. Channel types that support this field include: slack  **Note**: This property is sensitive and will not be displayed in the plan.
        /// </summary>
        [Input("authToken")]
        public Input<string>? AuthToken { get; set; }

        /// <summary>
        /// An password for a notification channel. Channel types that support this field include: webhook_basicauth  **Note**: This property is sensitive and will not be displayed in the plan.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// An servicekey token for a notification channel. Channel types that support this field include: pagerduty  **Note**: This property is sensitive and will not be displayed in the plan.
        /// </summary>
        [Input("serviceKey")]
        public Input<string>? ServiceKey { get; set; }

        public NotificationChannelSensitiveLabelsArgs()
        {
        }
    }
}
