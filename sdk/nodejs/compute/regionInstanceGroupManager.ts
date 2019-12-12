// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * The Google Compute Engine Regional Instance Group Manager API creates and manages pools
 * of homogeneous Compute Engine virtual machine instances from a common instance
 * template. For more information, see [the official documentation](https://cloud.google.com/compute/docs/instance-groups/distributing-instances-with-regional-instance-groups)
 * and [API](https://cloud.google.com/compute/docs/reference/latest/regionInstanceGroupManagers)
 * 
 * > **Note:** Use [gcp.compute.InstanceGroupManager](https://www.terraform.io/docs/providers/google/r/compute_instance_group_manager.html) to create a single-zone instance group manager.
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_region_instance_group_manager.html.markdown.
 */
export class RegionInstanceGroupManager extends pulumi.CustomResource {
    /**
     * Get an existing RegionInstanceGroupManager resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RegionInstanceGroupManagerState, opts?: pulumi.CustomResourceOptions): RegionInstanceGroupManager {
        return new RegionInstanceGroupManager(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/regionInstanceGroupManager:RegionInstanceGroupManager';

    /**
     * Returns true if the given object is an instance of RegionInstanceGroupManager.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RegionInstanceGroupManager {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RegionInstanceGroupManager.__pulumiType;
    }

    /**
     * The autohealing policies for this managed instance
     * group. You can specify only one value. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/creating-groups-of-managed-instances#monitoring_groups).
     */
    public readonly autoHealingPolicies!: pulumi.Output<outputs.compute.RegionInstanceGroupManagerAutoHealingPolicies | undefined>;
    /**
     * The base instance name to use for
     * instances in this group. The value must be a valid
     * [RFC1035](https://www.ietf.org/rfc/rfc1035.txt) name. Supported characters
     * are lowercase letters, numbers, and hyphens (-). Instances are named by
     * appending a hyphen and a random four-character string to the base instance
     * name.
     */
    public readonly baseInstanceName!: pulumi.Output<string>;
    /**
     * An optional textual description of the instance
     * group manager.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The distribution policy for this managed instance
     * group. You can specify one or more values. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/distributing-instances-with-regional-instance-groups#selectingzones).
     * - - -
     */
    public readonly distributionPolicyZones!: pulumi.Output<string[]>;
    /**
     * The fingerprint of the instance group manager.
     */
    public /*out*/ readonly fingerprint!: pulumi.Output<string>;
    /**
     * The full URL of the instance group created by the manager.
     */
    public /*out*/ readonly instanceGroup!: pulumi.Output<string>;
    /**
     * The name of the instance group manager. Must be 1-63
     * characters long and comply with
     * [RFC1035](https://www.ietf.org/rfc/rfc1035.txt). Supported characters
     * include lowercase letters, numbers, and hyphens.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The named port configuration. See the section below
     * for details on configuration.
     */
    public readonly namedPorts!: pulumi.Output<outputs.compute.RegionInstanceGroupManagerNamedPort[] | undefined>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The region where the managed instance group resides.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The URL of the created resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * The full URL of all target pools to which new
     * instances in the group are added. Updating the target pools attribute does
     * not affect existing instances.
     */
    public readonly targetPools!: pulumi.Output<string[] | undefined>;
    /**
     * The target number of running instances for this managed
     * instance group. This value should always be explicitly set unless this resource is attached to
     * an autoscaler, in which case it should never be set. Defaults to `0`.
     */
    public readonly targetSize!: pulumi.Output<number>;
    /**
     * The update policy for this managed instance group. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/updating-managed-instance-groups) and [API](https://cloud.google.com/compute/docs/reference/rest/beta/regionInstanceGroupManagers/patch)
     */
    public readonly updatePolicy!: pulumi.Output<outputs.compute.RegionInstanceGroupManagerUpdatePolicy>;
    /**
     * Application versions managed by this instance group. Each
     * version deals with a specific instance template, allowing canary release scenarios.
     * Structure is documented below.
     */
    public readonly versions!: pulumi.Output<outputs.compute.RegionInstanceGroupManagerVersion[]>;
    /**
     * Whether to wait for all instances to be created/updated before
     * returning. Note that if this is set to true and the operation does not succeed, this provider will
     * continue trying until it times out.
     */
    public readonly waitForInstances!: pulumi.Output<boolean | undefined>;

    /**
     * Create a RegionInstanceGroupManager resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RegionInstanceGroupManagerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RegionInstanceGroupManagerArgs | RegionInstanceGroupManagerState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as RegionInstanceGroupManagerState | undefined;
            inputs["autoHealingPolicies"] = state ? state.autoHealingPolicies : undefined;
            inputs["baseInstanceName"] = state ? state.baseInstanceName : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["distributionPolicyZones"] = state ? state.distributionPolicyZones : undefined;
            inputs["fingerprint"] = state ? state.fingerprint : undefined;
            inputs["instanceGroup"] = state ? state.instanceGroup : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["namedPorts"] = state ? state.namedPorts : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
            inputs["targetPools"] = state ? state.targetPools : undefined;
            inputs["targetSize"] = state ? state.targetSize : undefined;
            inputs["updatePolicy"] = state ? state.updatePolicy : undefined;
            inputs["versions"] = state ? state.versions : undefined;
            inputs["waitForInstances"] = state ? state.waitForInstances : undefined;
        } else {
            const args = argsOrState as RegionInstanceGroupManagerArgs | undefined;
            if (!args || args.baseInstanceName === undefined) {
                throw new Error("Missing required property 'baseInstanceName'");
            }
            if (!args || args.region === undefined) {
                throw new Error("Missing required property 'region'");
            }
            if (!args || args.versions === undefined) {
                throw new Error("Missing required property 'versions'");
            }
            inputs["autoHealingPolicies"] = args ? args.autoHealingPolicies : undefined;
            inputs["baseInstanceName"] = args ? args.baseInstanceName : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["distributionPolicyZones"] = args ? args.distributionPolicyZones : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["namedPorts"] = args ? args.namedPorts : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["targetPools"] = args ? args.targetPools : undefined;
            inputs["targetSize"] = args ? args.targetSize : undefined;
            inputs["updatePolicy"] = args ? args.updatePolicy : undefined;
            inputs["versions"] = args ? args.versions : undefined;
            inputs["waitForInstances"] = args ? args.waitForInstances : undefined;
            inputs["fingerprint"] = undefined /*out*/;
            inputs["instanceGroup"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(RegionInstanceGroupManager.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RegionInstanceGroupManager resources.
 */
export interface RegionInstanceGroupManagerState {
    /**
     * The autohealing policies for this managed instance
     * group. You can specify only one value. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/creating-groups-of-managed-instances#monitoring_groups).
     */
    readonly autoHealingPolicies?: pulumi.Input<inputs.compute.RegionInstanceGroupManagerAutoHealingPolicies>;
    /**
     * The base instance name to use for
     * instances in this group. The value must be a valid
     * [RFC1035](https://www.ietf.org/rfc/rfc1035.txt) name. Supported characters
     * are lowercase letters, numbers, and hyphens (-). Instances are named by
     * appending a hyphen and a random four-character string to the base instance
     * name.
     */
    readonly baseInstanceName?: pulumi.Input<string>;
    /**
     * An optional textual description of the instance
     * group manager.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The distribution policy for this managed instance
     * group. You can specify one or more values. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/distributing-instances-with-regional-instance-groups#selectingzones).
     * - - -
     */
    readonly distributionPolicyZones?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The fingerprint of the instance group manager.
     */
    readonly fingerprint?: pulumi.Input<string>;
    /**
     * The full URL of the instance group created by the manager.
     */
    readonly instanceGroup?: pulumi.Input<string>;
    /**
     * The name of the instance group manager. Must be 1-63
     * characters long and comply with
     * [RFC1035](https://www.ietf.org/rfc/rfc1035.txt). Supported characters
     * include lowercase letters, numbers, and hyphens.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The named port configuration. See the section below
     * for details on configuration.
     */
    readonly namedPorts?: pulumi.Input<pulumi.Input<inputs.compute.RegionInstanceGroupManagerNamedPort>[]>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The region where the managed instance group resides.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The URL of the created resource.
     */
    readonly selfLink?: pulumi.Input<string>;
    /**
     * The full URL of all target pools to which new
     * instances in the group are added. Updating the target pools attribute does
     * not affect existing instances.
     */
    readonly targetPools?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The target number of running instances for this managed
     * instance group. This value should always be explicitly set unless this resource is attached to
     * an autoscaler, in which case it should never be set. Defaults to `0`.
     */
    readonly targetSize?: pulumi.Input<number>;
    /**
     * The update policy for this managed instance group. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/updating-managed-instance-groups) and [API](https://cloud.google.com/compute/docs/reference/rest/beta/regionInstanceGroupManagers/patch)
     */
    readonly updatePolicy?: pulumi.Input<inputs.compute.RegionInstanceGroupManagerUpdatePolicy>;
    /**
     * Application versions managed by this instance group. Each
     * version deals with a specific instance template, allowing canary release scenarios.
     * Structure is documented below.
     */
    readonly versions?: pulumi.Input<pulumi.Input<inputs.compute.RegionInstanceGroupManagerVersion>[]>;
    /**
     * Whether to wait for all instances to be created/updated before
     * returning. Note that if this is set to true and the operation does not succeed, this provider will
     * continue trying until it times out.
     */
    readonly waitForInstances?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a RegionInstanceGroupManager resource.
 */
export interface RegionInstanceGroupManagerArgs {
    /**
     * The autohealing policies for this managed instance
     * group. You can specify only one value. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/creating-groups-of-managed-instances#monitoring_groups).
     */
    readonly autoHealingPolicies?: pulumi.Input<inputs.compute.RegionInstanceGroupManagerAutoHealingPolicies>;
    /**
     * The base instance name to use for
     * instances in this group. The value must be a valid
     * [RFC1035](https://www.ietf.org/rfc/rfc1035.txt) name. Supported characters
     * are lowercase letters, numbers, and hyphens (-). Instances are named by
     * appending a hyphen and a random four-character string to the base instance
     * name.
     */
    readonly baseInstanceName: pulumi.Input<string>;
    /**
     * An optional textual description of the instance
     * group manager.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The distribution policy for this managed instance
     * group. You can specify one or more values. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/distributing-instances-with-regional-instance-groups#selectingzones).
     * - - -
     */
    readonly distributionPolicyZones?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the instance group manager. Must be 1-63
     * characters long and comply with
     * [RFC1035](https://www.ietf.org/rfc/rfc1035.txt). Supported characters
     * include lowercase letters, numbers, and hyphens.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The named port configuration. See the section below
     * for details on configuration.
     */
    readonly namedPorts?: pulumi.Input<pulumi.Input<inputs.compute.RegionInstanceGroupManagerNamedPort>[]>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The region where the managed instance group resides.
     */
    readonly region: pulumi.Input<string>;
    /**
     * The full URL of all target pools to which new
     * instances in the group are added. Updating the target pools attribute does
     * not affect existing instances.
     */
    readonly targetPools?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The target number of running instances for this managed
     * instance group. This value should always be explicitly set unless this resource is attached to
     * an autoscaler, in which case it should never be set. Defaults to `0`.
     */
    readonly targetSize?: pulumi.Input<number>;
    /**
     * The update policy for this managed instance group. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/updating-managed-instance-groups) and [API](https://cloud.google.com/compute/docs/reference/rest/beta/regionInstanceGroupManagers/patch)
     */
    readonly updatePolicy?: pulumi.Input<inputs.compute.RegionInstanceGroupManagerUpdatePolicy>;
    /**
     * Application versions managed by this instance group. Each
     * version deals with a specific instance template, allowing canary release scenarios.
     * Structure is documented below.
     */
    readonly versions: pulumi.Input<pulumi.Input<inputs.compute.RegionInstanceGroupManagerVersion>[]>;
    /**
     * Whether to wait for all instances to be created/updated before
     * returning. Note that if this is set to true and the operation does not succeed, this provider will
     * continue trying until it times out.
     */
    readonly waitForInstances?: pulumi.Input<boolean>;
}
