// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute
{
    /// <summary>
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_image.html.markdown.
    /// </summary>
    public partial class Image : Pulumi.CustomResource
    {
        /// <summary>
        /// Size of the image tar.gz archive stored in Google Cloud Storage (in bytes).
        /// </summary>
        [Output("archiveSizeBytes")]
        public Output<int> ArchiveSizeBytes { get; private set; } = null!;

        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        [Output("creationTimestamp")]
        public Output<string> CreationTimestamp { get; private set; } = null!;

        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Size of the image when restored onto a persistent disk (in GB).
        /// </summary>
        [Output("diskSizeGb")]
        public Output<int> DiskSizeGb { get; private set; } = null!;

        /// <summary>
        /// The name of the image family to which this image belongs. You can create disks by specifying an image family
        /// instead of a specific image name. The image family always returns its latest image that is not deprecated.
        /// The name of the image family must comply with RFC1035.
        /// </summary>
        [Output("family")]
        public Output<string?> Family { get; private set; } = null!;

        /// <summary>
        /// A list of features to enable on the guest operating system. Applicable only for bootable images.
        /// </summary>
        [Output("guestOsFeatures")]
        public Output<ImmutableArray<Outputs.ImageGuestOsFeatures>> GuestOsFeatures { get; private set; } = null!;

        /// <summary>
        /// The fingerprint used for optimistic locking of this resource. Used internally during updates.
        /// </summary>
        [Output("labelFingerprint")]
        public Output<string> LabelFingerprint { get; private set; } = null!;

        /// <summary>
        /// Labels to apply to this Image.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// Any applicable license URI.
        /// </summary>
        [Output("licenses")]
        public Output<ImmutableArray<string>> Licenses { get; private set; } = null!;

        /// <summary>
        /// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters
        /// long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular
        /// expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all
        /// following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be
        /// a dash.
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
        /// The parameters of the raw disk image.
        /// </summary>
        [Output("rawDisk")]
        public Output<Outputs.ImageRawDisk?> RawDisk { get; private set; } = null!;

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// The source disk to create this image based on. You must provide either this property or the rawDisk.source
        /// property but not both to create an image.
        /// </summary>
        [Output("sourceDisk")]
        public Output<string?> SourceDisk { get; private set; } = null!;


        /// <summary>
        /// Create a Image resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Image(string name, ImageArgs? args = null, CustomResourceOptions? options = null)
            : base("gcp:compute/image:Image", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Image(string name, Input<string> id, ImageState? state = null, CustomResourceOptions? options = null)
            : base("gcp:compute/image:Image", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Image resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Image Get(string name, Input<string> id, ImageState? state = null, CustomResourceOptions? options = null)
        {
            return new Image(name, id, state, options);
        }
    }

    public sealed class ImageArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Size of the image when restored onto a persistent disk (in GB).
        /// </summary>
        [Input("diskSizeGb")]
        public Input<int>? DiskSizeGb { get; set; }

        /// <summary>
        /// The name of the image family to which this image belongs. You can create disks by specifying an image family
        /// instead of a specific image name. The image family always returns its latest image that is not deprecated.
        /// The name of the image family must comply with RFC1035.
        /// </summary>
        [Input("family")]
        public Input<string>? Family { get; set; }

        [Input("guestOsFeatures")]
        private InputList<Inputs.ImageGuestOsFeaturesArgs>? _guestOsFeatures;

        /// <summary>
        /// A list of features to enable on the guest operating system. Applicable only for bootable images.
        /// </summary>
        public InputList<Inputs.ImageGuestOsFeaturesArgs> GuestOsFeatures
        {
            get => _guestOsFeatures ?? (_guestOsFeatures = new InputList<Inputs.ImageGuestOsFeaturesArgs>());
            set => _guestOsFeatures = value;
        }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Labels to apply to this Image.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("licenses")]
        private InputList<string>? _licenses;

        /// <summary>
        /// Any applicable license URI.
        /// </summary>
        public InputList<string> Licenses
        {
            get => _licenses ?? (_licenses = new InputList<string>());
            set => _licenses = value;
        }

        /// <summary>
        /// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters
        /// long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular
        /// expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all
        /// following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be
        /// a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The parameters of the raw disk image.
        /// </summary>
        [Input("rawDisk")]
        public Input<Inputs.ImageRawDiskArgs>? RawDisk { get; set; }

        /// <summary>
        /// The source disk to create this image based on. You must provide either this property or the rawDisk.source
        /// property but not both to create an image.
        /// </summary>
        [Input("sourceDisk")]
        public Input<string>? SourceDisk { get; set; }

        public ImageArgs()
        {
        }
    }

    public sealed class ImageState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Size of the image tar.gz archive stored in Google Cloud Storage (in bytes).
        /// </summary>
        [Input("archiveSizeBytes")]
        public Input<int>? ArchiveSizeBytes { get; set; }

        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        [Input("creationTimestamp")]
        public Input<string>? CreationTimestamp { get; set; }

        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Size of the image when restored onto a persistent disk (in GB).
        /// </summary>
        [Input("diskSizeGb")]
        public Input<int>? DiskSizeGb { get; set; }

        /// <summary>
        /// The name of the image family to which this image belongs. You can create disks by specifying an image family
        /// instead of a specific image name. The image family always returns its latest image that is not deprecated.
        /// The name of the image family must comply with RFC1035.
        /// </summary>
        [Input("family")]
        public Input<string>? Family { get; set; }

        [Input("guestOsFeatures")]
        private InputList<Inputs.ImageGuestOsFeaturesGetArgs>? _guestOsFeatures;

        /// <summary>
        /// A list of features to enable on the guest operating system. Applicable only for bootable images.
        /// </summary>
        public InputList<Inputs.ImageGuestOsFeaturesGetArgs> GuestOsFeatures
        {
            get => _guestOsFeatures ?? (_guestOsFeatures = new InputList<Inputs.ImageGuestOsFeaturesGetArgs>());
            set => _guestOsFeatures = value;
        }

        /// <summary>
        /// The fingerprint used for optimistic locking of this resource. Used internally during updates.
        /// </summary>
        [Input("labelFingerprint")]
        public Input<string>? LabelFingerprint { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Labels to apply to this Image.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("licenses")]
        private InputList<string>? _licenses;

        /// <summary>
        /// Any applicable license URI.
        /// </summary>
        public InputList<string> Licenses
        {
            get => _licenses ?? (_licenses = new InputList<string>());
            set => _licenses = value;
        }

        /// <summary>
        /// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters
        /// long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular
        /// expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all
        /// following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be
        /// a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The parameters of the raw disk image.
        /// </summary>
        [Input("rawDisk")]
        public Input<Inputs.ImageRawDiskGetArgs>? RawDisk { get; set; }

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        /// <summary>
        /// The source disk to create this image based on. You must provide either this property or the rawDisk.source
        /// property but not both to create an image.
        /// </summary>
        [Input("sourceDisk")]
        public Input<string>? SourceDisk { get; set; }

        public ImageState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class ImageGuestOsFeaturesArgs : Pulumi.ResourceArgs
    {
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public ImageGuestOsFeaturesArgs()
        {
        }
    }

    public sealed class ImageGuestOsFeaturesGetArgs : Pulumi.ResourceArgs
    {
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public ImageGuestOsFeaturesGetArgs()
        {
        }
    }

    public sealed class ImageRawDiskArgs : Pulumi.ResourceArgs
    {
        [Input("containerType")]
        public Input<string>? ContainerType { get; set; }

        [Input("sha1")]
        public Input<string>? Sha1 { get; set; }

        [Input("source", required: true)]
        public Input<string> Source { get; set; } = null!;

        public ImageRawDiskArgs()
        {
        }
    }

    public sealed class ImageRawDiskGetArgs : Pulumi.ResourceArgs
    {
        [Input("containerType")]
        public Input<string>? ContainerType { get; set; }

        [Input("sha1")]
        public Input<string>? Sha1 { get; set; }

        [Input("source", required: true)]
        public Input<string> Source { get; set; } = null!;

        public ImageRawDiskGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class ImageGuestOsFeatures
    {
        public readonly string Type;

        [OutputConstructor]
        private ImageGuestOsFeatures(string type)
        {
            Type = type;
        }
    }

    [OutputType]
    public sealed class ImageRawDisk
    {
        public readonly string? ContainerType;
        public readonly string? Sha1;
        public readonly string Source;

        [OutputConstructor]
        private ImageRawDisk(
            string? containerType,
            string? sha1,
            string source)
        {
            ContainerType = containerType;
            Sha1 = sha1;
            Source = source;
        }
    }
    }
}
