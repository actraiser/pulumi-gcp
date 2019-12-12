// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages a RuntimeConfig variable in Google Cloud. For more information, see the
 * [official documentation](https://cloud.google.com/deployment-manager/runtime-configurator/),
 * or the
 * [JSON API](https://cloud.google.com/deployment-manager/runtime-configurator/reference/rest/).
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/runtimeconfig_variable.html.markdown.
 */
export class Variable extends pulumi.CustomResource {
    /**
     * Get an existing Variable resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VariableState, opts?: pulumi.CustomResourceOptions): Variable {
        return new Variable(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:runtimeconfig/variable:Variable';

    /**
     * Returns true if the given object is an instance of Variable.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Variable {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Variable.__pulumiType;
    }

    /**
     * The name of the variable to manage. Note that variable
     * names can be hierarchical using slashes (e.g. "prod-variables/hostname").
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the RuntimeConfig resource containing this
     * variable.
     */
    public readonly parent!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    public readonly text!: pulumi.Output<string | undefined>;
    /**
     * (Computed) The timestamp in RFC3339 UTC "Zulu" format,
     * accurate to nanoseconds, representing when the variable was last updated.
     * Example: "2016-10-09T12:33:37.578138407Z".
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;
    public readonly value!: pulumi.Output<string | undefined>;

    /**
     * Create a Variable resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VariableArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VariableArgs | VariableState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as VariableState | undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["parent"] = state ? state.parent : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["text"] = state ? state.text : undefined;
            inputs["updateTime"] = state ? state.updateTime : undefined;
            inputs["value"] = state ? state.value : undefined;
        } else {
            const args = argsOrState as VariableArgs | undefined;
            if (!args || args.parent === undefined) {
                throw new Error("Missing required property 'parent'");
            }
            inputs["name"] = args ? args.name : undefined;
            inputs["parent"] = args ? args.parent : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["text"] = args ? args.text : undefined;
            inputs["value"] = args ? args.value : undefined;
            inputs["updateTime"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Variable.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Variable resources.
 */
export interface VariableState {
    /**
     * The name of the variable to manage. Note that variable
     * names can be hierarchical using slashes (e.g. "prod-variables/hostname").
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the RuntimeConfig resource containing this
     * variable.
     */
    readonly parent?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    readonly text?: pulumi.Input<string>;
    /**
     * (Computed) The timestamp in RFC3339 UTC "Zulu" format,
     * accurate to nanoseconds, representing when the variable was last updated.
     * Example: "2016-10-09T12:33:37.578138407Z".
     */
    readonly updateTime?: pulumi.Input<string>;
    readonly value?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Variable resource.
 */
export interface VariableArgs {
    /**
     * The name of the variable to manage. Note that variable
     * names can be hierarchical using slashes (e.g. "prod-variables/hostname").
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the RuntimeConfig resource containing this
     * variable.
     */
    readonly parent: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    readonly text?: pulumi.Input<string>;
    readonly value?: pulumi.Input<string>;
}
