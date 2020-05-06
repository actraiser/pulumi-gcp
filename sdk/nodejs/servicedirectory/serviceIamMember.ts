// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Three different resources help you manage your IAM policy for Service Directory Service. Each of these resources serves a different use case:
 * 
 * * `gcp.servicedirectory.ServiceIamPolicy`: Authoritative. Sets the IAM policy for the service and replaces any existing policy already attached.
 * * `gcp.servicedirectory.ServiceIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the service are preserved.
 * * `gcp.servicedirectory.ServiceIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the service are preserved.
 * 
 * > **Note:** `gcp.servicedirectory.ServiceIamPolicy` **cannot** be used in conjunction with `gcp.servicedirectory.ServiceIamBinding` and `gcp.servicedirectory.ServiceIamMember` or they will fight over what your policy should be.
 * 
 * > **Note:** `gcp.servicedirectory.ServiceIamBinding` resources **can be** used in conjunction with `gcp.servicedirectory.ServiceIamMember` resources **only if** they do not grant privilege to the same role.
 * 
 * ## google\_service\_directory\_service\_iam\_policy
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const admin = gcp.organizations.getIAMPolicy({
 *     binding: [{
 *         role: "roles/viewer",
 *         members: ["user:jane@example.com"],
 *     }],
 * });
 * const policy = new gcp.servicedirectory.ServiceIamPolicy("policy", {policyData: admin.then(admin => admin.policyData)});
 * ```
 * 
 * ## google\_service\_directory\_service\_iam\_binding
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const binding = new gcp.servicedirectory.ServiceIamBinding("binding", {
 *     role: "roles/viewer",
 *     members: ["user:jane@example.com"],
 * });
 * ```
 * 
 * ## google\_service\_directory\_service\_iam\_member
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const member = new gcp.servicedirectory.ServiceIamMember("member", {
 *     role: "roles/viewer",
 *     member: "user:jane@example.com",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/service_directory_service_iam.html.markdown.
 */
export class ServiceIamMember extends pulumi.CustomResource {
    /**
     * Get an existing ServiceIamMember resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceIamMemberState, opts?: pulumi.CustomResourceOptions): ServiceIamMember {
        return new ServiceIamMember(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:servicedirectory/serviceIamMember:ServiceIamMember';

    /**
     * Returns true if the given object is an instance of ServiceIamMember.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServiceIamMember {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServiceIamMember.__pulumiType;
    }

    public readonly condition!: pulumi.Output<outputs.servicedirectory.ServiceIamMemberCondition | undefined>;
    /**
     * (Computed) The etag of the IAM policy.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    public readonly member!: pulumi.Output<string>;
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.servicedirectory.ServiceIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    public readonly role!: pulumi.Output<string>;

    /**
     * Create a ServiceIamMember resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceIamMemberArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceIamMemberArgs | ServiceIamMemberState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ServiceIamMemberState | undefined;
            inputs["condition"] = state ? state.condition : undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["member"] = state ? state.member : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["role"] = state ? state.role : undefined;
        } else {
            const args = argsOrState as ServiceIamMemberArgs | undefined;
            if (!args || args.member === undefined) {
                throw new Error("Missing required property 'member'");
            }
            if (!args || args.role === undefined) {
                throw new Error("Missing required property 'role'");
            }
            inputs["condition"] = args ? args.condition : undefined;
            inputs["member"] = args ? args.member : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["role"] = args ? args.role : undefined;
            inputs["etag"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ServiceIamMember.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServiceIamMember resources.
 */
export interface ServiceIamMemberState {
    readonly condition?: pulumi.Input<inputs.servicedirectory.ServiceIamMemberCondition>;
    /**
     * (Computed) The etag of the IAM policy.
     */
    readonly etag?: pulumi.Input<string>;
    readonly member?: pulumi.Input<string>;
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.servicedirectory.ServiceIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    readonly role?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ServiceIamMember resource.
 */
export interface ServiceIamMemberArgs {
    readonly condition?: pulumi.Input<inputs.servicedirectory.ServiceIamMemberCondition>;
    readonly member: pulumi.Input<string>;
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.servicedirectory.ServiceIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    readonly role: pulumi.Input<string>;
}