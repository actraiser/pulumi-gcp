// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";

/**
 * Allows creation and management of a single binding within IAM policy for
 * an existing Google Cloud Platform Organization.
 * 
 * ~> **Note:** This resource __must not__ be used in conjunction with
 *    `google_organization_iam_member` for the __same role__ or they will fight over
 *    what your policy should be.
 */
export class IAMBinding extends pulumi.CustomResource {
    /**
     * Get an existing IAMBinding resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IAMBindingState): IAMBinding {
        return new IAMBinding(name, <any>state, { id });
    }

    /**
     * (Computed) The etag of the organization's IAM policy.
     */
    public /*out*/ readonly etag: pulumi.Output<string>;
    /**
     * A list of users that the role should apply to.
     */
    public readonly members: pulumi.Output<string[]>;
    /**
     * The numeric ID of the organization in which you want to create a custom role.
     */
    public readonly orgId: pulumi.Output<string>;
    /**
     * The role that should be applied. Only one
     * `google_organization_iam_binding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    public readonly role: pulumi.Output<string>;

    /**
     * Create a IAMBinding resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IAMBindingArgs, opts?: pulumi.ResourceOptions)
    constructor(name: string, argsOrState?: IAMBindingArgs | IAMBindingState, opts?: pulumi.ResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: IAMBindingState = argsOrState as IAMBindingState | undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["members"] = state ? state.members : undefined;
            inputs["orgId"] = state ? state.orgId : undefined;
            inputs["role"] = state ? state.role : undefined;
        } else {
            const args = argsOrState as IAMBindingArgs | undefined;
            if (!args || args.members === undefined) {
                throw new Error("Missing required property 'members'");
            }
            if (!args || args.orgId === undefined) {
                throw new Error("Missing required property 'orgId'");
            }
            if (!args || args.role === undefined) {
                throw new Error("Missing required property 'role'");
            }
            inputs["members"] = args ? args.members : undefined;
            inputs["orgId"] = args ? args.orgId : undefined;
            inputs["role"] = args ? args.role : undefined;
            inputs["etag"] = undefined /*out*/;
        }
        super("gcp:organizations/iAMBinding:IAMBinding", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IAMBinding resources.
 */
export interface IAMBindingState {
    /**
     * (Computed) The etag of the organization's IAM policy.
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * A list of users that the role should apply to.
     */
    readonly members?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The numeric ID of the organization in which you want to create a custom role.
     */
    readonly orgId?: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `google_organization_iam_binding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    readonly role?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IAMBinding resource.
 */
export interface IAMBindingArgs {
    /**
     * A list of users that the role should apply to.
     */
    readonly members: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The numeric ID of the organization in which you want to create a custom role.
     */
    readonly orgId: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `google_organization_iam_binding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    readonly role: pulumi.Input<string>;
}
