// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * App profile is a configuration object describing how Cloud Bigtable should treat traffic from a particular end user application.
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/bigtable_app_profile.html.markdown.
 */
export class AppProfile extends pulumi.CustomResource {
    /**
     * Get an existing AppProfile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AppProfileState, opts?: pulumi.CustomResourceOptions): AppProfile {
        return new AppProfile(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:bigquery/appProfile:AppProfile';

    /**
     * Returns true if the given object is an instance of AppProfile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AppProfile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AppProfile.__pulumiType;
    }

    /**
     * The unique name of the app profile in the form `[_a-zA-Z0-9][-_.a-zA-Z0-9]*`.
     */
    public readonly appProfileId!: pulumi.Output<string>;
    /**
     * Long form description of the use case for this app profile.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * If true, ignore safety checks when deleting/updating the app profile.
     */
    public readonly ignoreWarnings!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the instance to create the app profile within.
     */
    public readonly instance!: pulumi.Output<string | undefined>;
    /**
     * If true, read/write requests are routed to the nearest cluster in the instance, and will fail over to the nearest cluster that is available
     * in the event of transient errors or delays. Clusters in a region are considered equidistant. Choosing this option sacrifices read-your-writes
     * consistency to improve availability.
     */
    public readonly multiClusterRoutingUseAny!: pulumi.Output<boolean | undefined>;
    /**
     * The unique name of the requested app profile. Values are of the form
     * 'projects/<project>/instances/<instance>/appProfiles/<appProfileId>'.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Use a single-cluster routing policy.  Structure is documented below.
     */
    public readonly singleClusterRouting!: pulumi.Output<outputs.bigquery.AppProfileSingleClusterRouting | undefined>;

    /**
     * Create a AppProfile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AppProfileArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AppProfileArgs | AppProfileState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as AppProfileState | undefined;
            inputs["appProfileId"] = state ? state.appProfileId : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["ignoreWarnings"] = state ? state.ignoreWarnings : undefined;
            inputs["instance"] = state ? state.instance : undefined;
            inputs["multiClusterRoutingUseAny"] = state ? state.multiClusterRoutingUseAny : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["singleClusterRouting"] = state ? state.singleClusterRouting : undefined;
        } else {
            const args = argsOrState as AppProfileArgs | undefined;
            if (!args || args.appProfileId === undefined) {
                throw new Error("Missing required property 'appProfileId'");
            }
            inputs["appProfileId"] = args ? args.appProfileId : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["ignoreWarnings"] = args ? args.ignoreWarnings : undefined;
            inputs["instance"] = args ? args.instance : undefined;
            inputs["multiClusterRoutingUseAny"] = args ? args.multiClusterRoutingUseAny : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["singleClusterRouting"] = args ? args.singleClusterRouting : undefined;
            inputs["name"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(AppProfile.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AppProfile resources.
 */
export interface AppProfileState {
    /**
     * The unique name of the app profile in the form `[_a-zA-Z0-9][-_.a-zA-Z0-9]*`.
     */
    readonly appProfileId?: pulumi.Input<string>;
    /**
     * Long form description of the use case for this app profile.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * If true, ignore safety checks when deleting/updating the app profile.
     */
    readonly ignoreWarnings?: pulumi.Input<boolean>;
    /**
     * The name of the instance to create the app profile within.
     */
    readonly instance?: pulumi.Input<string>;
    /**
     * If true, read/write requests are routed to the nearest cluster in the instance, and will fail over to the nearest cluster that is available
     * in the event of transient errors or delays. Clusters in a region are considered equidistant. Choosing this option sacrifices read-your-writes
     * consistency to improve availability.
     */
    readonly multiClusterRoutingUseAny?: pulumi.Input<boolean>;
    /**
     * The unique name of the requested app profile. Values are of the form
     * 'projects/<project>/instances/<instance>/appProfiles/<appProfileId>'.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Use a single-cluster routing policy.  Structure is documented below.
     */
    readonly singleClusterRouting?: pulumi.Input<inputs.bigquery.AppProfileSingleClusterRouting>;
}

/**
 * The set of arguments for constructing a AppProfile resource.
 */
export interface AppProfileArgs {
    /**
     * The unique name of the app profile in the form `[_a-zA-Z0-9][-_.a-zA-Z0-9]*`.
     */
    readonly appProfileId: pulumi.Input<string>;
    /**
     * Long form description of the use case for this app profile.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * If true, ignore safety checks when deleting/updating the app profile.
     */
    readonly ignoreWarnings?: pulumi.Input<boolean>;
    /**
     * The name of the instance to create the app profile within.
     */
    readonly instance?: pulumi.Input<string>;
    /**
     * If true, read/write requests are routed to the nearest cluster in the instance, and will fail over to the nearest cluster that is available
     * in the event of transient errors or delays. Clusters in a region are considered equidistant. Choosing this option sacrifices read-your-writes
     * consistency to improve availability.
     */
    readonly multiClusterRoutingUseAny?: pulumi.Input<boolean>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Use a single-cluster routing policy.  Structure is documented below.
     */
    readonly singleClusterRouting?: pulumi.Input<inputs.bigquery.AppProfileSingleClusterRouting>;
}
