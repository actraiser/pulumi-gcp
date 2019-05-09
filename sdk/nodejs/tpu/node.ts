// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * A Cloud TPU instance.
 * 
 * 
 * To get more information about Node, see:
 * 
 * * [API documentation](https://cloud.google.com/tpu/docs/reference/rest/)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/tpu/docs/)
 * 
 * ## Example Usage - Tpu Node Basic
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const available = pulumi.output(gcp.tpu.getTensorflowVersions({}));
 * const tpu = new gcp.tpu.Node("tpu", {
 *     acceleratorType: "v3-8",
 *     cidrBlock: "10.2.0.0/29",
 *     tensorflowVersion: available.apply(available => available.versions[0]),
 *     zone: "us-central1-b",
 * });
 * ```
 * ## Example Usage - Tpu Node Full
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const available = pulumi.output(gcp.tpu.getTensorflowVersions({}));
 * const tpuNetwork = new gcp.compute.Network("tpu_network", {
 *     autoCreateSubnetworks: false,
 * });
 * const tpu = new gcp.tpu.Node("tpu", {
 *     acceleratorType: "v3-8",
 *     cidrBlock: "10.3.0.0/29",
 *     description: "Terraform Google Provider test TPU",
 *     labels: {
 *         foo: "bar",
 *     },
 *     network: tpuNetwork.name,
 *     schedulingConfig: {
 *         preemptible: true,
 *     },
 *     tensorflowVersion: available.apply(available => available.versions[0]),
 *     zone: "us-central1-b",
 * });
 * ```
 */
export class Node extends pulumi.CustomResource {
    /**
     * Get an existing Node resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NodeState, opts?: pulumi.CustomResourceOptions): Node {
        return new Node(name, <any>state, { ...opts, id: id });
    }

    public readonly acceleratorType: pulumi.Output<string>;
    public readonly cidrBlock: pulumi.Output<string>;
    public readonly description: pulumi.Output<string | undefined>;
    public readonly labels: pulumi.Output<{[key: string]: string} | undefined>;
    public readonly name: pulumi.Output<string>;
    public readonly network: pulumi.Output<string>;
    public /*out*/ readonly networkEndpoints: pulumi.Output<{ ipAddress: string, port: number }[]>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project: pulumi.Output<string>;
    public readonly schedulingConfig: pulumi.Output<{ preemptible?: boolean } | undefined>;
    public /*out*/ readonly serviceAccount: pulumi.Output<string>;
    public readonly tensorflowVersion: pulumi.Output<string>;
    public readonly zone: pulumi.Output<string>;

    /**
     * Create a Node resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NodeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NodeArgs | NodeState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: NodeState = argsOrState as NodeState | undefined;
            inputs["acceleratorType"] = state ? state.acceleratorType : undefined;
            inputs["cidrBlock"] = state ? state.cidrBlock : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["network"] = state ? state.network : undefined;
            inputs["networkEndpoints"] = state ? state.networkEndpoints : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["schedulingConfig"] = state ? state.schedulingConfig : undefined;
            inputs["serviceAccount"] = state ? state.serviceAccount : undefined;
            inputs["tensorflowVersion"] = state ? state.tensorflowVersion : undefined;
            inputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as NodeArgs | undefined;
            if (!args || args.acceleratorType === undefined) {
                throw new Error("Missing required property 'acceleratorType'");
            }
            if (!args || args.cidrBlock === undefined) {
                throw new Error("Missing required property 'cidrBlock'");
            }
            if (!args || args.tensorflowVersion === undefined) {
                throw new Error("Missing required property 'tensorflowVersion'");
            }
            if (!args || args.zone === undefined) {
                throw new Error("Missing required property 'zone'");
            }
            inputs["acceleratorType"] = args ? args.acceleratorType : undefined;
            inputs["cidrBlock"] = args ? args.cidrBlock : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["network"] = args ? args.network : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["schedulingConfig"] = args ? args.schedulingConfig : undefined;
            inputs["tensorflowVersion"] = args ? args.tensorflowVersion : undefined;
            inputs["zone"] = args ? args.zone : undefined;
            inputs["networkEndpoints"] = undefined /*out*/;
            inputs["serviceAccount"] = undefined /*out*/;
        }
        super("gcp:tpu/node:Node", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Node resources.
 */
export interface NodeState {
    readonly acceleratorType?: pulumi.Input<string>;
    readonly cidrBlock?: pulumi.Input<string>;
    readonly description?: pulumi.Input<string>;
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly name?: pulumi.Input<string>;
    readonly network?: pulumi.Input<string>;
    readonly networkEndpoints?: pulumi.Input<pulumi.Input<{ ipAddress?: pulumi.Input<string>, port?: pulumi.Input<number> }>[]>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    readonly schedulingConfig?: pulumi.Input<{ preemptible?: pulumi.Input<boolean> }>;
    readonly serviceAccount?: pulumi.Input<string>;
    readonly tensorflowVersion?: pulumi.Input<string>;
    readonly zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Node resource.
 */
export interface NodeArgs {
    readonly acceleratorType: pulumi.Input<string>;
    readonly cidrBlock: pulumi.Input<string>;
    readonly description?: pulumi.Input<string>;
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly name?: pulumi.Input<string>;
    readonly network?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    readonly schedulingConfig?: pulumi.Input<{ preemptible?: pulumi.Input<boolean> }>;
    readonly tensorflowVersion: pulumi.Input<string>;
    readonly zone: pulumi.Input<string>;
}
