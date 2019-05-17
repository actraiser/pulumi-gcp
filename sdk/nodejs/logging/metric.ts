// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Logs-based metric can also be used to extract values from logs and create a a distribution
 * of the values. The distribution records the statistics of the extracted values along with
 * an optional histogram of the values as specified by the bucket options.
 * 
 * 
 * To get more information about Metric, see:
 * 
 * * [API documentation](https://cloud.google.com/logging/docs/reference/v2/rest/v2/projects.metrics/create)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/logging/docs/apis)
 * 
 * ## Example Usage - Logging Metric Basic
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const loggingMetric = new gcp.logging.Metric("logging_metric", {
 *     bucketOptions: {
 *         linearBuckets: {
 *             numFiniteBuckets: 3,
 *             offset: 1,
 *             width: 1,
 *         },
 *     },
 *     filter: "resource.type=gae_app AND severity>=ERROR",
 *     labelExtractors: {
 *         "EXTRACT(jsonPayload.request)": [{}],
 *         mass: [{}],
 *     },
 *     metricDescriptor: {
 *         labels: [{
 *             description: "amount of matter",
 *             key: "mass",
 *             valueType: "STRING",
 *         }],
 *         metricKind: "DELTA",
 *         valueType: "DISTRIBUTION",
 *     },
 *     valueExtractor: "EXTRACT(jsonPayload.request)",
 * });
 * ```
 */
export class Metric extends pulumi.CustomResource {
    /**
     * Get an existing Metric resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MetricState, opts?: pulumi.CustomResourceOptions): Metric {
        return new Metric(name, <any>state, { ...opts, id: id });
    }

    public readonly bucketOptions: pulumi.Output<{ explicit?: { bounds?: string[] }, exponentialBuckets?: { growthFactor?: number, numFiniteBuckets?: number, scale?: number }, linearBuckets?: { numFiniteBuckets?: number, offset?: number, width?: number } } | undefined>;
    public readonly description: pulumi.Output<string | undefined>;
    public readonly filter: pulumi.Output<string>;
    public readonly labelExtractors: pulumi.Output<{[key: string]: string} | undefined>;
    public readonly metricDescriptor: pulumi.Output<{ labels?: { description?: string, key: string, valueType?: string }[], metricKind: string, valueType: string }>;
    public readonly name: pulumi.Output<string>;
    public readonly project: pulumi.Output<string>;
    public readonly valueExtractor: pulumi.Output<string | undefined>;

    /**
     * Create a Metric resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MetricArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MetricArgs | MetricState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: MetricState = argsOrState as MetricState | undefined;
            inputs["bucketOptions"] = state ? state.bucketOptions : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["filter"] = state ? state.filter : undefined;
            inputs["labelExtractors"] = state ? state.labelExtractors : undefined;
            inputs["metricDescriptor"] = state ? state.metricDescriptor : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["valueExtractor"] = state ? state.valueExtractor : undefined;
        } else {
            const args = argsOrState as MetricArgs | undefined;
            if (!args || args.filter === undefined) {
                throw new Error("Missing required property 'filter'");
            }
            if (!args || args.metricDescriptor === undefined) {
                throw new Error("Missing required property 'metricDescriptor'");
            }
            inputs["bucketOptions"] = args ? args.bucketOptions : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["filter"] = args ? args.filter : undefined;
            inputs["labelExtractors"] = args ? args.labelExtractors : undefined;
            inputs["metricDescriptor"] = args ? args.metricDescriptor : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["valueExtractor"] = args ? args.valueExtractor : undefined;
        }
        super("gcp:logging/metric:Metric", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Metric resources.
 */
export interface MetricState {
    readonly bucketOptions?: pulumi.Input<{ explicit?: pulumi.Input<{ bounds?: pulumi.Input<pulumi.Input<string>[]> }>, exponentialBuckets?: pulumi.Input<{ growthFactor?: pulumi.Input<number>, numFiniteBuckets?: pulumi.Input<number>, scale?: pulumi.Input<number> }>, linearBuckets?: pulumi.Input<{ numFiniteBuckets?: pulumi.Input<number>, offset?: pulumi.Input<number>, width?: pulumi.Input<number> }> }>;
    readonly description?: pulumi.Input<string>;
    readonly filter?: pulumi.Input<string>;
    readonly labelExtractors?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly metricDescriptor?: pulumi.Input<{ labels?: pulumi.Input<pulumi.Input<{ description?: pulumi.Input<string>, key: pulumi.Input<string>, valueType?: pulumi.Input<string> }>[]>, metricKind: pulumi.Input<string>, valueType: pulumi.Input<string> }>;
    readonly name?: pulumi.Input<string>;
    readonly project?: pulumi.Input<string>;
    readonly valueExtractor?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Metric resource.
 */
export interface MetricArgs {
    readonly bucketOptions?: pulumi.Input<{ explicit?: pulumi.Input<{ bounds?: pulumi.Input<pulumi.Input<string>[]> }>, exponentialBuckets?: pulumi.Input<{ growthFactor?: pulumi.Input<number>, numFiniteBuckets?: pulumi.Input<number>, scale?: pulumi.Input<number> }>, linearBuckets?: pulumi.Input<{ numFiniteBuckets?: pulumi.Input<number>, offset?: pulumi.Input<number>, width?: pulumi.Input<number> }> }>;
    readonly description?: pulumi.Input<string>;
    readonly filter: pulumi.Input<string>;
    readonly labelExtractors?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly metricDescriptor: pulumi.Input<{ labels?: pulumi.Input<pulumi.Input<{ description?: pulumi.Input<string>, key: pulumi.Input<string>, valueType?: pulumi.Input<string> }>[]>, metricKind: pulumi.Input<string>, valueType: pulumi.Input<string> }>;
    readonly name?: pulumi.Input<string>;
    readonly project?: pulumi.Input<string>;
    readonly valueExtractor?: pulumi.Input<string>;
}