// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute
{
    /// <summary>
    /// Manages a VM instance resource within GCE. For more information see
    /// [the official documentation](https://cloud.google.com/compute/docs/instances)
    /// and
    /// [API](https://cloud.google.com/compute/docs/reference/latest/instances).
    /// 
    /// This resource is specifically to create a compute instance from a given
    /// `source_instance_template`. To create an instance without a template, use the
    /// `gcp.compute.Instance` resource.
    /// </summary>
    public partial class InstanceFromTemplate : Pulumi.CustomResource
    {
        [Output("allowStoppingForUpdate")]
        public Output<bool> AllowStoppingForUpdate { get; private set; } = null!;

        [Output("attachedDisks")]
        public Output<ImmutableArray<Outputs.InstanceFromTemplateAttachedDisk>> AttachedDisks { get; private set; } = null!;

        [Output("bootDisk")]
        public Output<Outputs.InstanceFromTemplateBootDisk> BootDisk { get; private set; } = null!;

        [Output("canIpForward")]
        public Output<bool> CanIpForward { get; private set; } = null!;

        [Output("cpuPlatform")]
        public Output<string> CpuPlatform { get; private set; } = null!;

        [Output("currentStatus")]
        public Output<string> CurrentStatus { get; private set; } = null!;

        [Output("deletionProtection")]
        public Output<bool> DeletionProtection { get; private set; } = null!;

        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        [Output("desiredStatus")]
        public Output<string> DesiredStatus { get; private set; } = null!;

        [Output("enableDisplay")]
        public Output<bool> EnableDisplay { get; private set; } = null!;

        [Output("guestAccelerators")]
        public Output<ImmutableArray<Outputs.InstanceFromTemplateGuestAccelerator>> GuestAccelerators { get; private set; } = null!;

        [Output("hostname")]
        public Output<string> Hostname { get; private set; } = null!;

        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        [Output("labelFingerprint")]
        public Output<string> LabelFingerprint { get; private set; } = null!;

        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        [Output("machineType")]
        public Output<string> MachineType { get; private set; } = null!;

        [Output("metadata")]
        public Output<ImmutableDictionary<string, string>> Metadata { get; private set; } = null!;

        [Output("metadataFingerprint")]
        public Output<string> MetadataFingerprint { get; private set; } = null!;

        [Output("metadataStartupScript")]
        public Output<string> MetadataStartupScript { get; private set; } = null!;

        [Output("minCpuPlatform")]
        public Output<string> MinCpuPlatform { get; private set; } = null!;

        /// <summary>
        /// A unique name for the resource, required by GCE.
        /// Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("networkInterfaces")]
        public Output<ImmutableArray<Outputs.InstanceFromTemplateNetworkInterface>> NetworkInterfaces { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        [Output("resourcePolicies")]
        public Output<string> ResourcePolicies { get; private set; } = null!;

        [Output("scheduling")]
        public Output<Outputs.InstanceFromTemplateScheduling> Scheduling { get; private set; } = null!;

        [Output("scratchDisks")]
        public Output<ImmutableArray<Outputs.InstanceFromTemplateScratchDisk>> ScratchDisks { get; private set; } = null!;

        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        [Output("serviceAccount")]
        public Output<Outputs.InstanceFromTemplateServiceAccount> ServiceAccount { get; private set; } = null!;

        [Output("shieldedInstanceConfig")]
        public Output<Outputs.InstanceFromTemplateShieldedInstanceConfig> ShieldedInstanceConfig { get; private set; } = null!;

        /// <summary>
        /// Name or self link of an instance
        /// template to create the instance based on.
        /// </summary>
        [Output("sourceInstanceTemplate")]
        public Output<string> SourceInstanceTemplate { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        [Output("tagsFingerprint")]
        public Output<string> TagsFingerprint { get; private set; } = null!;

        /// <summary>
        /// The zone that the machine should be created in. If not
        /// set, the provider zone is used.
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a InstanceFromTemplate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public InstanceFromTemplate(string name, InstanceFromTemplateArgs args, CustomResourceOptions? options = null)
            : base("gcp:compute/instanceFromTemplate:InstanceFromTemplate", name, args ?? new InstanceFromTemplateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private InstanceFromTemplate(string name, Input<string> id, InstanceFromTemplateState? state = null, CustomResourceOptions? options = null)
            : base("gcp:compute/instanceFromTemplate:InstanceFromTemplate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing InstanceFromTemplate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static InstanceFromTemplate Get(string name, Input<string> id, InstanceFromTemplateState? state = null, CustomResourceOptions? options = null)
        {
            return new InstanceFromTemplate(name, id, state, options);
        }
    }

    public sealed class InstanceFromTemplateArgs : Pulumi.ResourceArgs
    {
        [Input("allowStoppingForUpdate")]
        public Input<bool>? AllowStoppingForUpdate { get; set; }

        [Input("attachedDisks")]
        private InputList<Inputs.InstanceFromTemplateAttachedDiskArgs>? _attachedDisks;
        public InputList<Inputs.InstanceFromTemplateAttachedDiskArgs> AttachedDisks
        {
            get => _attachedDisks ?? (_attachedDisks = new InputList<Inputs.InstanceFromTemplateAttachedDiskArgs>());
            set => _attachedDisks = value;
        }

        [Input("bootDisk")]
        public Input<Inputs.InstanceFromTemplateBootDiskArgs>? BootDisk { get; set; }

        [Input("canIpForward")]
        public Input<bool>? CanIpForward { get; set; }

        [Input("deletionProtection")]
        public Input<bool>? DeletionProtection { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("desiredStatus")]
        public Input<string>? DesiredStatus { get; set; }

        [Input("enableDisplay")]
        public Input<bool>? EnableDisplay { get; set; }

        [Input("guestAccelerators")]
        private InputList<Inputs.InstanceFromTemplateGuestAcceleratorArgs>? _guestAccelerators;
        public InputList<Inputs.InstanceFromTemplateGuestAcceleratorArgs> GuestAccelerators
        {
            get => _guestAccelerators ?? (_guestAccelerators = new InputList<Inputs.InstanceFromTemplateGuestAcceleratorArgs>());
            set => _guestAccelerators = value;
        }

        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("machineType")]
        public Input<string>? MachineType { get; set; }

        [Input("metadata")]
        private InputMap<string>? _metadata;
        public InputMap<string> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<string>());
            set => _metadata = value;
        }

        [Input("metadataStartupScript")]
        public Input<string>? MetadataStartupScript { get; set; }

        [Input("minCpuPlatform")]
        public Input<string>? MinCpuPlatform { get; set; }

        /// <summary>
        /// A unique name for the resource, required by GCE.
        /// Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("networkInterfaces")]
        private InputList<Inputs.InstanceFromTemplateNetworkInterfaceArgs>? _networkInterfaces;
        public InputList<Inputs.InstanceFromTemplateNetworkInterfaceArgs> NetworkInterfaces
        {
            get => _networkInterfaces ?? (_networkInterfaces = new InputList<Inputs.InstanceFromTemplateNetworkInterfaceArgs>());
            set => _networkInterfaces = value;
        }

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("resourcePolicies")]
        public Input<string>? ResourcePolicies { get; set; }

        [Input("scheduling")]
        public Input<Inputs.InstanceFromTemplateSchedulingArgs>? Scheduling { get; set; }

        [Input("scratchDisks")]
        private InputList<Inputs.InstanceFromTemplateScratchDiskArgs>? _scratchDisks;
        public InputList<Inputs.InstanceFromTemplateScratchDiskArgs> ScratchDisks
        {
            get => _scratchDisks ?? (_scratchDisks = new InputList<Inputs.InstanceFromTemplateScratchDiskArgs>());
            set => _scratchDisks = value;
        }

        [Input("serviceAccount")]
        public Input<Inputs.InstanceFromTemplateServiceAccountArgs>? ServiceAccount { get; set; }

        [Input("shieldedInstanceConfig")]
        public Input<Inputs.InstanceFromTemplateShieldedInstanceConfigArgs>? ShieldedInstanceConfig { get; set; }

        /// <summary>
        /// Name or self link of an instance
        /// template to create the instance based on.
        /// </summary>
        [Input("sourceInstanceTemplate", required: true)]
        public Input<string> SourceInstanceTemplate { get; set; } = null!;

        [Input("tags")]
        private InputList<string>? _tags;
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The zone that the machine should be created in. If not
        /// set, the provider zone is used.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public InstanceFromTemplateArgs()
        {
        }
    }

    public sealed class InstanceFromTemplateState : Pulumi.ResourceArgs
    {
        [Input("allowStoppingForUpdate")]
        public Input<bool>? AllowStoppingForUpdate { get; set; }

        [Input("attachedDisks")]
        private InputList<Inputs.InstanceFromTemplateAttachedDiskGetArgs>? _attachedDisks;
        public InputList<Inputs.InstanceFromTemplateAttachedDiskGetArgs> AttachedDisks
        {
            get => _attachedDisks ?? (_attachedDisks = new InputList<Inputs.InstanceFromTemplateAttachedDiskGetArgs>());
            set => _attachedDisks = value;
        }

        [Input("bootDisk")]
        public Input<Inputs.InstanceFromTemplateBootDiskGetArgs>? BootDisk { get; set; }

        [Input("canIpForward")]
        public Input<bool>? CanIpForward { get; set; }

        [Input("cpuPlatform")]
        public Input<string>? CpuPlatform { get; set; }

        [Input("currentStatus")]
        public Input<string>? CurrentStatus { get; set; }

        [Input("deletionProtection")]
        public Input<bool>? DeletionProtection { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("desiredStatus")]
        public Input<string>? DesiredStatus { get; set; }

        [Input("enableDisplay")]
        public Input<bool>? EnableDisplay { get; set; }

        [Input("guestAccelerators")]
        private InputList<Inputs.InstanceFromTemplateGuestAcceleratorGetArgs>? _guestAccelerators;
        public InputList<Inputs.InstanceFromTemplateGuestAcceleratorGetArgs> GuestAccelerators
        {
            get => _guestAccelerators ?? (_guestAccelerators = new InputList<Inputs.InstanceFromTemplateGuestAcceleratorGetArgs>());
            set => _guestAccelerators = value;
        }

        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        [Input("labelFingerprint")]
        public Input<string>? LabelFingerprint { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("machineType")]
        public Input<string>? MachineType { get; set; }

        [Input("metadata")]
        private InputMap<string>? _metadata;
        public InputMap<string> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<string>());
            set => _metadata = value;
        }

        [Input("metadataFingerprint")]
        public Input<string>? MetadataFingerprint { get; set; }

        [Input("metadataStartupScript")]
        public Input<string>? MetadataStartupScript { get; set; }

        [Input("minCpuPlatform")]
        public Input<string>? MinCpuPlatform { get; set; }

        /// <summary>
        /// A unique name for the resource, required by GCE.
        /// Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("networkInterfaces")]
        private InputList<Inputs.InstanceFromTemplateNetworkInterfaceGetArgs>? _networkInterfaces;
        public InputList<Inputs.InstanceFromTemplateNetworkInterfaceGetArgs> NetworkInterfaces
        {
            get => _networkInterfaces ?? (_networkInterfaces = new InputList<Inputs.InstanceFromTemplateNetworkInterfaceGetArgs>());
            set => _networkInterfaces = value;
        }

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("resourcePolicies")]
        public Input<string>? ResourcePolicies { get; set; }

        [Input("scheduling")]
        public Input<Inputs.InstanceFromTemplateSchedulingGetArgs>? Scheduling { get; set; }

        [Input("scratchDisks")]
        private InputList<Inputs.InstanceFromTemplateScratchDiskGetArgs>? _scratchDisks;
        public InputList<Inputs.InstanceFromTemplateScratchDiskGetArgs> ScratchDisks
        {
            get => _scratchDisks ?? (_scratchDisks = new InputList<Inputs.InstanceFromTemplateScratchDiskGetArgs>());
            set => _scratchDisks = value;
        }

        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        [Input("serviceAccount")]
        public Input<Inputs.InstanceFromTemplateServiceAccountGetArgs>? ServiceAccount { get; set; }

        [Input("shieldedInstanceConfig")]
        public Input<Inputs.InstanceFromTemplateShieldedInstanceConfigGetArgs>? ShieldedInstanceConfig { get; set; }

        /// <summary>
        /// Name or self link of an instance
        /// template to create the instance based on.
        /// </summary>
        [Input("sourceInstanceTemplate")]
        public Input<string>? SourceInstanceTemplate { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        [Input("tagsFingerprint")]
        public Input<string>? TagsFingerprint { get; set; }

        /// <summary>
        /// The zone that the machine should be created in. If not
        /// set, the provider zone is used.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public InstanceFromTemplateState()
        {
        }
    }
}
