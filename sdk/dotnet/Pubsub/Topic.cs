// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.PubSub
{
    public partial class Topic : Pulumi.CustomResource
    {
        /// <summary>
        /// The resource name of the Cloud KMS CryptoKey to be used to protect access to messages published on this
        /// topic. Your project's PubSub service account
        /// ('service-{{PROJECT_NUMBER}}@gcp-sa-pubsub.iam.gserviceaccount.com') must have
        /// 'roles/cloudkms.cryptoKeyEncrypterDecrypter' to use this feature. The expected format is
        /// 'projects/*/locations/*/keyRings/*/cryptoKeys/*'
        /// </summary>
        [Output("kmsKeyName")]
        public Output<string?> KmsKeyName { get; private set; } = null!;

        /// <summary>
        /// A set of key/value label pairs to assign to this Topic.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// Policy constraining the set of Google Cloud Platform regions where messages published to the topic may be
        /// stored. If not present, then no constraints are in effect.
        /// </summary>
        [Output("messageStoragePolicy")]
        public Output<Outputs.TopicMessageStoragePolicy> MessageStoragePolicy { get; private set; } = null!;

        /// <summary>
        /// Name of the topic.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;


        /// <summary>
        /// Create a Topic resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Topic(string name, TopicArgs? args = null, CustomResourceOptions? options = null)
            : base("gcp:pubsub/topic:Topic", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Topic(string name, Input<string> id, TopicState? state = null, CustomResourceOptions? options = null)
            : base("gcp:pubsub/topic:Topic", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Topic resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Topic Get(string name, Input<string> id, TopicState? state = null, CustomResourceOptions? options = null)
        {
            return new Topic(name, id, state, options);
        }
    }

    public sealed class TopicArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The resource name of the Cloud KMS CryptoKey to be used to protect access to messages published on this
        /// topic. Your project's PubSub service account
        /// ('service-{{PROJECT_NUMBER}}@gcp-sa-pubsub.iam.gserviceaccount.com') must have
        /// 'roles/cloudkms.cryptoKeyEncrypterDecrypter' to use this feature. The expected format is
        /// 'projects/*/locations/*/keyRings/*/cryptoKeys/*'
        /// </summary>
        [Input("kmsKeyName")]
        public Input<string>? KmsKeyName { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// A set of key/value label pairs to assign to this Topic.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Policy constraining the set of Google Cloud Platform regions where messages published to the topic may be
        /// stored. If not present, then no constraints are in effect.
        /// </summary>
        [Input("messageStoragePolicy")]
        public Input<Inputs.TopicMessageStoragePolicyArgs>? MessageStoragePolicy { get; set; }

        /// <summary>
        /// Name of the topic.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public TopicArgs()
        {
        }
    }

    public sealed class TopicState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The resource name of the Cloud KMS CryptoKey to be used to protect access to messages published on this
        /// topic. Your project's PubSub service account
        /// ('service-{{PROJECT_NUMBER}}@gcp-sa-pubsub.iam.gserviceaccount.com') must have
        /// 'roles/cloudkms.cryptoKeyEncrypterDecrypter' to use this feature. The expected format is
        /// 'projects/*/locations/*/keyRings/*/cryptoKeys/*'
        /// </summary>
        [Input("kmsKeyName")]
        public Input<string>? KmsKeyName { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// A set of key/value label pairs to assign to this Topic.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Policy constraining the set of Google Cloud Platform regions where messages published to the topic may be
        /// stored. If not present, then no constraints are in effect.
        /// </summary>
        [Input("messageStoragePolicy")]
        public Input<Inputs.TopicMessageStoragePolicyGetArgs>? MessageStoragePolicy { get; set; }

        /// <summary>
        /// Name of the topic.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public TopicState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class TopicMessageStoragePolicyArgs : Pulumi.ResourceArgs
    {
        [Input("allowedPersistenceRegions", required: true)]
        private InputList<string>? _allowedPersistenceRegions;
        public InputList<string> AllowedPersistenceRegions
        {
            get => _allowedPersistenceRegions ?? (_allowedPersistenceRegions = new InputList<string>());
            set => _allowedPersistenceRegions = value;
        }

        public TopicMessageStoragePolicyArgs()
        {
        }
    }

    public sealed class TopicMessageStoragePolicyGetArgs : Pulumi.ResourceArgs
    {
        [Input("allowedPersistenceRegions", required: true)]
        private InputList<string>? _allowedPersistenceRegions;
        public InputList<string> AllowedPersistenceRegions
        {
            get => _allowedPersistenceRegions ?? (_allowedPersistenceRegions = new InputList<string>());
            set => _allowedPersistenceRegions = value;
        }

        public TopicMessageStoragePolicyGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class TopicMessageStoragePolicy
    {
        public readonly ImmutableArray<string> AllowedPersistenceRegions;

        [OutputConstructor]
        private TopicMessageStoragePolicy(ImmutableArray<string> allowedPersistenceRegions)
        {
            AllowedPersistenceRegions = allowedPersistenceRegions;
        }
    }
    }
}
