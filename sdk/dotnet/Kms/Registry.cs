// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Kms
{
    /// <summary>
    /// 
    /// Deprecated: gcp.kms.Registry has been deprecated in favour of gcp.iot.Registry
    /// </summary>
    [Obsolete(@"gcp.kms.Registry has been deprecated in favour of gcp.iot.Registry")]
    public partial class Registry : Pulumi.CustomResource
    {
        [Output("credentials")]
        public Output<ImmutableArray<Outputs.RegistryCredential>> Credentials { get; private set; } = null!;

        [Output("eventNotificationConfigs")]
        public Output<ImmutableArray<Outputs.RegistryEventNotificationConfigItem>> EventNotificationConfigs { get; private set; } = null!;

        [Output("httpConfig")]
        public Output<Outputs.RegistryHttpConfig> HttpConfig { get; private set; } = null!;

        [Output("logLevel")]
        public Output<string?> LogLevel { get; private set; } = null!;

        [Output("mqttConfig")]
        public Output<Outputs.RegistryMqttConfig> MqttConfig { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        [Output("stateNotificationConfig")]
        public Output<Outputs.RegistryStateNotificationConfig?> StateNotificationConfig { get; private set; } = null!;


        /// <summary>
        /// Create a Registry resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Registry(string name, RegistryArgs? args = null, CustomResourceOptions? options = null)
            : base("gcp:kms/registry:Registry", name, args ?? new RegistryArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Registry(string name, Input<string> id, RegistryState? state = null, CustomResourceOptions? options = null)
            : base("gcp:kms/registry:Registry", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Registry resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Registry Get(string name, Input<string> id, RegistryState? state = null, CustomResourceOptions? options = null)
        {
            return new Registry(name, id, state, options);
        }
    }

    public sealed class RegistryArgs : Pulumi.ResourceArgs
    {
        [Input("credentials")]
        private InputList<Inputs.RegistryCredentialArgs>? _credentials;
        public InputList<Inputs.RegistryCredentialArgs> Credentials
        {
            get => _credentials ?? (_credentials = new InputList<Inputs.RegistryCredentialArgs>());
            set => _credentials = value;
        }

        [Input("eventNotificationConfigs")]
        private InputList<Inputs.RegistryEventNotificationConfigItemArgs>? _eventNotificationConfigs;
        public InputList<Inputs.RegistryEventNotificationConfigItemArgs> EventNotificationConfigs
        {
            get => _eventNotificationConfigs ?? (_eventNotificationConfigs = new InputList<Inputs.RegistryEventNotificationConfigItemArgs>());
            set => _eventNotificationConfigs = value;
        }

        [Input("httpConfig")]
        public Input<Inputs.RegistryHttpConfigArgs>? HttpConfig { get; set; }

        [Input("logLevel")]
        public Input<string>? LogLevel { get; set; }

        [Input("mqttConfig")]
        public Input<Inputs.RegistryMqttConfigArgs>? MqttConfig { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("stateNotificationConfig")]
        public Input<Inputs.RegistryStateNotificationConfigArgs>? StateNotificationConfig { get; set; }

        public RegistryArgs()
        {
        }
    }

    public sealed class RegistryState : Pulumi.ResourceArgs
    {
        [Input("credentials")]
        private InputList<Inputs.RegistryCredentialGetArgs>? _credentials;
        public InputList<Inputs.RegistryCredentialGetArgs> Credentials
        {
            get => _credentials ?? (_credentials = new InputList<Inputs.RegistryCredentialGetArgs>());
            set => _credentials = value;
        }

        [Input("eventNotificationConfigs")]
        private InputList<Inputs.RegistryEventNotificationConfigItemGetArgs>? _eventNotificationConfigs;
        public InputList<Inputs.RegistryEventNotificationConfigItemGetArgs> EventNotificationConfigs
        {
            get => _eventNotificationConfigs ?? (_eventNotificationConfigs = new InputList<Inputs.RegistryEventNotificationConfigItemGetArgs>());
            set => _eventNotificationConfigs = value;
        }

        [Input("httpConfig")]
        public Input<Inputs.RegistryHttpConfigGetArgs>? HttpConfig { get; set; }

        [Input("logLevel")]
        public Input<string>? LogLevel { get; set; }

        [Input("mqttConfig")]
        public Input<Inputs.RegistryMqttConfigGetArgs>? MqttConfig { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("stateNotificationConfig")]
        public Input<Inputs.RegistryStateNotificationConfigGetArgs>? StateNotificationConfig { get; set; }

        public RegistryState()
        {
        }
    }
}
