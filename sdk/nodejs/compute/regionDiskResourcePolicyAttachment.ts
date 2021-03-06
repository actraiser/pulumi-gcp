// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Adds existing resource policies to a disk. You can only add one policy
 * which will be applied to this disk for scheduling snapshot creation.
 * 
 * > **Note:** This resource does not support zonal disks (`gcp.compute.Disk`).
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_region_disk_resource_policy_attachment.html.markdown.
 */
export class RegionDiskResourcePolicyAttachment extends pulumi.CustomResource {
    /**
     * Get an existing RegionDiskResourcePolicyAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RegionDiskResourcePolicyAttachmentState, opts?: pulumi.CustomResourceOptions): RegionDiskResourcePolicyAttachment {
        return new RegionDiskResourcePolicyAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/regionDiskResourcePolicyAttachment:RegionDiskResourcePolicyAttachment';

    /**
     * Returns true if the given object is an instance of RegionDiskResourcePolicyAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RegionDiskResourcePolicyAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RegionDiskResourcePolicyAttachment.__pulumiType;
    }

    /**
     * The name of the regional disk in which the resource policies are attached to.
     */
    public readonly disk!: pulumi.Output<string>;
    /**
     * The resource policy to be attached to the disk for scheduling snapshot
     * creation. Do not specify the self link.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * A reference to the region where the disk resides.
     */
    public readonly region!: pulumi.Output<string>;

    /**
     * Create a RegionDiskResourcePolicyAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RegionDiskResourcePolicyAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RegionDiskResourcePolicyAttachmentArgs | RegionDiskResourcePolicyAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as RegionDiskResourcePolicyAttachmentState | undefined;
            inputs["disk"] = state ? state.disk : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["region"] = state ? state.region : undefined;
        } else {
            const args = argsOrState as RegionDiskResourcePolicyAttachmentArgs | undefined;
            if (!args || args.disk === undefined) {
                throw new Error("Missing required property 'disk'");
            }
            inputs["disk"] = args ? args.disk : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["region"] = args ? args.region : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(RegionDiskResourcePolicyAttachment.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RegionDiskResourcePolicyAttachment resources.
 */
export interface RegionDiskResourcePolicyAttachmentState {
    /**
     * The name of the regional disk in which the resource policies are attached to.
     */
    readonly disk?: pulumi.Input<string>;
    /**
     * The resource policy to be attached to the disk for scheduling snapshot
     * creation. Do not specify the self link.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * A reference to the region where the disk resides.
     */
    readonly region?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RegionDiskResourcePolicyAttachment resource.
 */
export interface RegionDiskResourcePolicyAttachmentArgs {
    /**
     * The name of the regional disk in which the resource policies are attached to.
     */
    readonly disk: pulumi.Input<string>;
    /**
     * The resource policy to be attached to the disk for scheduling snapshot
     * creation. Do not specify the self link.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * A reference to the region where the disk resides.
     */
    readonly region?: pulumi.Input<string>;
}
