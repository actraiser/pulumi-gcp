// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Represents a Persistent Disk Snapshot resource.
 * 
 * Use snapshots to back up data from your persistent disks. Snapshots are
 * different from public images and custom images, which are used primarily
 * to create instances or configure instance templates. Snapshots are useful
 * for periodic backup of the data on your persistent disks. You can create
 * snapshots from persistent disks even while they are attached to running
 * instances.
 * 
 * Snapshots are incremental, so you can create regular snapshots on a
 * persistent disk faster and at a much lower cost than if you regularly
 * created a full image of the disk.
 * 
 * 
 * To get more information about Snapshot, see:
 * 
 * * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/snapshots)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/compute/docs/disks/create-snapshots)
 * 
 * > **Warning:** All arguments including `snapshot_encryption_key.raw_key` and `source_disk_encryption_key.raw_key` will be stored in the raw
 * state as plain-text. [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_snapshot.html.markdown.
 */
export class Snapshot extends pulumi.CustomResource {
    /**
     * Get an existing Snapshot resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SnapshotState, opts?: pulumi.CustomResourceOptions): Snapshot {
        return new Snapshot(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/snapshot:Snapshot';

    /**
     * Returns true if the given object is an instance of Snapshot.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Snapshot {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Snapshot.__pulumiType;
    }

    /**
     * Creation timestamp in RFC3339 text format.
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * An optional description of this resource.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Size of the snapshot, specified in GB.
     */
    public /*out*/ readonly diskSizeGb!: pulumi.Output<number>;
    /**
     * The fingerprint used for optimistic locking of this resource. Used internally during updates.
     */
    public /*out*/ readonly labelFingerprint!: pulumi.Output<string>;
    /**
     * Labels to apply to this Snapshot.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A list of public visible licenses that apply to this snapshot. This can be because the original image had licenses
     * attached (such as a Windows image). snapshotEncryptionKey nested object Encrypts the snapshot using a
     * customer-supplied encryption key.
     */
    public /*out*/ readonly licenses!: pulumi.Output<string[]>;
    /**
     * Name of the resource; provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035. Specifically, the name must be 1-63 characters long and match
     * the regular expression `a-z?` which means the
     * first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the last
     * character, which cannot be a dash.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * The customer-supplied encryption key of the snapshot. Required if the
     * source snapshot is protected by a customer-supplied encryption key.  Structure is documented below.
     */
    public readonly snapshotEncryptionKey!: pulumi.Output<outputs.compute.SnapshotSnapshotEncryptionKey | undefined>;
    /**
     * The unique identifier for the resource.
     */
    public /*out*/ readonly snapshotId!: pulumi.Output<number>;
    /**
     * A reference to the disk used to create this snapshot.
     */
    public readonly sourceDisk!: pulumi.Output<string>;
    /**
     * The customer-supplied encryption key of the source snapshot. Required
     * if the source snapshot is protected by a customer-supplied encryption
     * key.  Structure is documented below.
     */
    public readonly sourceDiskEncryptionKey!: pulumi.Output<outputs.compute.SnapshotSourceDiskEncryptionKey | undefined>;
    public /*out*/ readonly sourceDiskLink!: pulumi.Output<string>;
    /**
     * A size of the the storage used by the snapshot. As snapshots share storage, this number is expected to change with
     * snapshot creation/deletion.
     */
    public /*out*/ readonly storageBytes!: pulumi.Output<number>;
    /**
     * A reference to the zone where the disk is hosted.
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a Snapshot resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SnapshotArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SnapshotArgs | SnapshotState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as SnapshotState | undefined;
            inputs["creationTimestamp"] = state ? state.creationTimestamp : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["diskSizeGb"] = state ? state.diskSizeGb : undefined;
            inputs["labelFingerprint"] = state ? state.labelFingerprint : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["licenses"] = state ? state.licenses : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
            inputs["snapshotEncryptionKey"] = state ? state.snapshotEncryptionKey : undefined;
            inputs["snapshotId"] = state ? state.snapshotId : undefined;
            inputs["sourceDisk"] = state ? state.sourceDisk : undefined;
            inputs["sourceDiskEncryptionKey"] = state ? state.sourceDiskEncryptionKey : undefined;
            inputs["sourceDiskLink"] = state ? state.sourceDiskLink : undefined;
            inputs["storageBytes"] = state ? state.storageBytes : undefined;
            inputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as SnapshotArgs | undefined;
            if (!args || args.sourceDisk === undefined) {
                throw new Error("Missing required property 'sourceDisk'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["snapshotEncryptionKey"] = args ? args.snapshotEncryptionKey : undefined;
            inputs["sourceDisk"] = args ? args.sourceDisk : undefined;
            inputs["sourceDiskEncryptionKey"] = args ? args.sourceDiskEncryptionKey : undefined;
            inputs["zone"] = args ? args.zone : undefined;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["diskSizeGb"] = undefined /*out*/;
            inputs["labelFingerprint"] = undefined /*out*/;
            inputs["licenses"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
            inputs["snapshotId"] = undefined /*out*/;
            inputs["sourceDiskLink"] = undefined /*out*/;
            inputs["storageBytes"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Snapshot.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Snapshot resources.
 */
export interface SnapshotState {
    /**
     * Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp?: pulumi.Input<string>;
    /**
     * An optional description of this resource.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Size of the snapshot, specified in GB.
     */
    readonly diskSizeGb?: pulumi.Input<number>;
    /**
     * The fingerprint used for optimistic locking of this resource. Used internally during updates.
     */
    readonly labelFingerprint?: pulumi.Input<string>;
    /**
     * Labels to apply to this Snapshot.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A list of public visible licenses that apply to this snapshot. This can be because the original image had licenses
     * attached (such as a Windows image). snapshotEncryptionKey nested object Encrypts the snapshot using a
     * customer-supplied encryption key.
     */
    readonly licenses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the resource; provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035. Specifically, the name must be 1-63 characters long and match
     * the regular expression `a-z?` which means the
     * first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the last
     * character, which cannot be a dash.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The URI of the created resource.
     */
    readonly selfLink?: pulumi.Input<string>;
    /**
     * The customer-supplied encryption key of the snapshot. Required if the
     * source snapshot is protected by a customer-supplied encryption key.  Structure is documented below.
     */
    readonly snapshotEncryptionKey?: pulumi.Input<inputs.compute.SnapshotSnapshotEncryptionKey>;
    /**
     * The unique identifier for the resource.
     */
    readonly snapshotId?: pulumi.Input<number>;
    /**
     * A reference to the disk used to create this snapshot.
     */
    readonly sourceDisk?: pulumi.Input<string>;
    /**
     * The customer-supplied encryption key of the source snapshot. Required
     * if the source snapshot is protected by a customer-supplied encryption
     * key.  Structure is documented below.
     */
    readonly sourceDiskEncryptionKey?: pulumi.Input<inputs.compute.SnapshotSourceDiskEncryptionKey>;
    readonly sourceDiskLink?: pulumi.Input<string>;
    /**
     * A size of the the storage used by the snapshot. As snapshots share storage, this number is expected to change with
     * snapshot creation/deletion.
     */
    readonly storageBytes?: pulumi.Input<number>;
    /**
     * A reference to the zone where the disk is hosted.
     */
    readonly zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Snapshot resource.
 */
export interface SnapshotArgs {
    /**
     * An optional description of this resource.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Labels to apply to this Snapshot.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Name of the resource; provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035. Specifically, the name must be 1-63 characters long and match
     * the regular expression `a-z?` which means the
     * first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the last
     * character, which cannot be a dash.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The customer-supplied encryption key of the snapshot. Required if the
     * source snapshot is protected by a customer-supplied encryption key.  Structure is documented below.
     */
    readonly snapshotEncryptionKey?: pulumi.Input<inputs.compute.SnapshotSnapshotEncryptionKey>;
    /**
     * A reference to the disk used to create this snapshot.
     */
    readonly sourceDisk: pulumi.Input<string>;
    /**
     * The customer-supplied encryption key of the source snapshot. Required
     * if the source snapshot is protected by a customer-supplied encryption
     * key.  Structure is documented below.
     */
    readonly sourceDiskEncryptionKey?: pulumi.Input<inputs.compute.SnapshotSourceDiskEncryptionKey>;
    /**
     * A reference to the zone where the disk is hosted.
     */
    readonly zone?: pulumi.Input<string>;
}
