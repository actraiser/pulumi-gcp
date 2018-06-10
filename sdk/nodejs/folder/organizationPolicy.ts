// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";

/**
 * Allows management of Organization policies for a Google Folder. For more information see
 * [the official
 * documentation](https://cloud.google.com/resource-manager/docs/organization-policy/overview) and
 * [API](https://cloud.google.com/resource-manager/reference/rest/v1/folders/setOrgPolicy).
 */
export class OrganizationPolicy extends pulumi.CustomResource {
    /**
     * Get an existing OrganizationPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OrganizationPolicyState): OrganizationPolicy {
        return new OrganizationPolicy(name, <any>state, { id });
    }

    /**
     * A boolean policy is a constraint that is either enforced or not. Structure is documented below. 
     */
    public readonly booleanPolicy: pulumi.Output<{ enforced: boolean } | undefined>;
    /**
     * The name of the Constraint the Policy is configuring, for example, `serviceuser.services`. Check out the [complete list of available constraints](https://cloud.google.com/resource-manager/docs/organization-policy/understanding-constraints#available_constraints).
     */
    public readonly constraint: pulumi.Output<string>;
    /**
     * (Computed) The etag of the organization policy. `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. 
     */
    public /*out*/ readonly etag: pulumi.Output<string>;
    /**
     * The resource name of the folder to set the policy for. Its format is folders/{folder_id}.
     */
    public readonly folder: pulumi.Output<string>;
    /**
     * A policy that can define specific values that are allowed or denied for the given constraint. It can also be used to allow or deny all values. Structure is documented below.
     */
    public readonly listPolicy: pulumi.Output<{ allow?: { all?: boolean, values?: string[] }, deny?: { all?: boolean, values?: string[] }, suggestedValue: string } | undefined>;
    /**
     * (Computed) The timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds, representing when the variable was last updated. Example: "2016-10-09T12:33:37.578138407Z".
     */
    public /*out*/ readonly updateTime: pulumi.Output<string>;
    /**
     * Version of the Policy. Default version is 0.
     */
    public readonly version: pulumi.Output<number>;

    /**
     * Create a OrganizationPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OrganizationPolicyArgs, opts?: pulumi.ResourceOptions)
    constructor(name: string, argsOrState?: OrganizationPolicyArgs | OrganizationPolicyState, opts?: pulumi.ResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: OrganizationPolicyState = argsOrState as OrganizationPolicyState | undefined;
            inputs["booleanPolicy"] = state ? state.booleanPolicy : undefined;
            inputs["constraint"] = state ? state.constraint : undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["folder"] = state ? state.folder : undefined;
            inputs["listPolicy"] = state ? state.listPolicy : undefined;
            inputs["updateTime"] = state ? state.updateTime : undefined;
            inputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as OrganizationPolicyArgs | undefined;
            if (!args || args.constraint === undefined) {
                throw new Error("Missing required property 'constraint'");
            }
            if (!args || args.folder === undefined) {
                throw new Error("Missing required property 'folder'");
            }
            inputs["booleanPolicy"] = args ? args.booleanPolicy : undefined;
            inputs["constraint"] = args ? args.constraint : undefined;
            inputs["folder"] = args ? args.folder : undefined;
            inputs["listPolicy"] = args ? args.listPolicy : undefined;
            inputs["version"] = args ? args.version : undefined;
            inputs["etag"] = undefined /*out*/;
            inputs["updateTime"] = undefined /*out*/;
        }
        super("gcp:folder/organizationPolicy:OrganizationPolicy", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering OrganizationPolicy resources.
 */
export interface OrganizationPolicyState {
    /**
     * A boolean policy is a constraint that is either enforced or not. Structure is documented below. 
     */
    readonly booleanPolicy?: pulumi.Input<{ enforced: pulumi.Input<boolean> }>;
    /**
     * The name of the Constraint the Policy is configuring, for example, `serviceuser.services`. Check out the [complete list of available constraints](https://cloud.google.com/resource-manager/docs/organization-policy/understanding-constraints#available_constraints).
     */
    readonly constraint?: pulumi.Input<string>;
    /**
     * (Computed) The etag of the organization policy. `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. 
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * The resource name of the folder to set the policy for. Its format is folders/{folder_id}.
     */
    readonly folder?: pulumi.Input<string>;
    /**
     * A policy that can define specific values that are allowed or denied for the given constraint. It can also be used to allow or deny all values. Structure is documented below.
     */
    readonly listPolicy?: pulumi.Input<{ allow?: pulumi.Input<{ all?: pulumi.Input<boolean>, values?: pulumi.Input<pulumi.Input<string>[]> }>, deny?: pulumi.Input<{ all?: pulumi.Input<boolean>, values?: pulumi.Input<pulumi.Input<string>[]> }>, suggestedValue?: pulumi.Input<string> }>;
    /**
     * (Computed) The timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds, representing when the variable was last updated. Example: "2016-10-09T12:33:37.578138407Z".
     */
    readonly updateTime?: pulumi.Input<string>;
    /**
     * Version of the Policy. Default version is 0.
     */
    readonly version?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a OrganizationPolicy resource.
 */
export interface OrganizationPolicyArgs {
    /**
     * A boolean policy is a constraint that is either enforced or not. Structure is documented below. 
     */
    readonly booleanPolicy?: pulumi.Input<{ enforced: pulumi.Input<boolean> }>;
    /**
     * The name of the Constraint the Policy is configuring, for example, `serviceuser.services`. Check out the [complete list of available constraints](https://cloud.google.com/resource-manager/docs/organization-policy/understanding-constraints#available_constraints).
     */
    readonly constraint: pulumi.Input<string>;
    /**
     * The resource name of the folder to set the policy for. Its format is folders/{folder_id}.
     */
    readonly folder: pulumi.Input<string>;
    /**
     * A policy that can define specific values that are allowed or denied for the given constraint. It can also be used to allow or deny all values. Structure is documented below.
     */
    readonly listPolicy?: pulumi.Input<{ allow?: pulumi.Input<{ all?: pulumi.Input<boolean>, values?: pulumi.Input<pulumi.Input<string>[]> }>, deny?: pulumi.Input<{ all?: pulumi.Input<boolean>, values?: pulumi.Input<pulumi.Input<string>[]> }>, suggestedValue?: pulumi.Input<string> }>;
    /**
     * Version of the Policy. Default version is 0.
     */
    readonly version?: pulumi.Input<number>;
}
