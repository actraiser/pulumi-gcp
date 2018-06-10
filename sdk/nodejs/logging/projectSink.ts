// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";

/**
 * Manages a project-level logging sink. For more information see
 * [the official documentation](https://cloud.google.com/logging/docs/),
 * [Exporting Logs in the API](https://cloud.google.com/logging/docs/api/tasks/exporting-logs)
 * and
 * [API](https://cloud.google.com/logging/docs/reference/v2/rest/).
 * 
 * Note that you must have the "Logs Configuration Writer" IAM role (`roles/logging.configWriter`)
 * granted to the credentials used with terraform.
 */
export class ProjectSink extends pulumi.CustomResource {
    /**
     * Get an existing ProjectSink resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProjectSinkState): ProjectSink {
        return new ProjectSink(name, <any>state, { id });
    }

    /**
     * The destination of the sink (or, in other words, where logs are written to). Can be a
     * Cloud Storage bucket, a PubSub topic, or a BigQuery dataset. Examples:
     * ```
     * "storage.googleapis.com/[GCS_BUCKET]"
     * "bigquery.googleapis.com/projects/[PROJECT_ID]/datasets/[DATASET]"
     * "pubsub.googleapis.com/projects/[PROJECT_ID]/topics/[TOPIC_ID]"
     * ```
     * The writer associated with the sink must have access to write to the above resource.
     */
    public readonly destination: pulumi.Output<string>;
    /**
     * The filter to apply when exporting logs. Only log entries that match the filter are exported.
     * See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
     * write a filter.
     */
    public readonly filter: pulumi.Output<string | undefined>;
    /**
     * The name of the logging sink.
     */
    public readonly name: pulumi.Output<string>;
    /**
     * The ID of the project to create the sink in. If omitted, the project associated with the provider is
     * used.
     */
    public readonly project: pulumi.Output<string>;
    /**
     * Whether or not to create a unique identity associated with this sink. If `false`
     * (the default), then the `writer_identity` used is `serviceAccount:cloud-logs@system.gserviceaccount.com`. If `true`,
     * then a unique service account is created and used for this sink. If you wish to publish logs across projects, you
     * must set `unique_writer_identity` to true.
     */
    public readonly uniqueWriterIdentity: pulumi.Output<boolean | undefined>;
    /**
     * The identity associated with this sink. This identity must be granted write access to the
     * configured `destination`.
     */
    public /*out*/ readonly writerIdentity: pulumi.Output<string>;

    /**
     * Create a ProjectSink resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProjectSinkArgs, opts?: pulumi.ResourceOptions)
    constructor(name: string, argsOrState?: ProjectSinkArgs | ProjectSinkState, opts?: pulumi.ResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: ProjectSinkState = argsOrState as ProjectSinkState | undefined;
            inputs["destination"] = state ? state.destination : undefined;
            inputs["filter"] = state ? state.filter : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["uniqueWriterIdentity"] = state ? state.uniqueWriterIdentity : undefined;
            inputs["writerIdentity"] = state ? state.writerIdentity : undefined;
        } else {
            const args = argsOrState as ProjectSinkArgs | undefined;
            if (!args || args.destination === undefined) {
                throw new Error("Missing required property 'destination'");
            }
            inputs["destination"] = args ? args.destination : undefined;
            inputs["filter"] = args ? args.filter : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["uniqueWriterIdentity"] = args ? args.uniqueWriterIdentity : undefined;
            inputs["writerIdentity"] = undefined /*out*/;
        }
        super("gcp:logging/projectSink:ProjectSink", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProjectSink resources.
 */
export interface ProjectSinkState {
    /**
     * The destination of the sink (or, in other words, where logs are written to). Can be a
     * Cloud Storage bucket, a PubSub topic, or a BigQuery dataset. Examples:
     * ```
     * "storage.googleapis.com/[GCS_BUCKET]"
     * "bigquery.googleapis.com/projects/[PROJECT_ID]/datasets/[DATASET]"
     * "pubsub.googleapis.com/projects/[PROJECT_ID]/topics/[TOPIC_ID]"
     * ```
     * The writer associated with the sink must have access to write to the above resource.
     */
    readonly destination?: pulumi.Input<string>;
    /**
     * The filter to apply when exporting logs. Only log entries that match the filter are exported.
     * See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
     * write a filter.
     */
    readonly filter?: pulumi.Input<string>;
    /**
     * The name of the logging sink.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project to create the sink in. If omitted, the project associated with the provider is
     * used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Whether or not to create a unique identity associated with this sink. If `false`
     * (the default), then the `writer_identity` used is `serviceAccount:cloud-logs@system.gserviceaccount.com`. If `true`,
     * then a unique service account is created and used for this sink. If you wish to publish logs across projects, you
     * must set `unique_writer_identity` to true.
     */
    readonly uniqueWriterIdentity?: pulumi.Input<boolean>;
    /**
     * The identity associated with this sink. This identity must be granted write access to the
     * configured `destination`.
     */
    readonly writerIdentity?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ProjectSink resource.
 */
export interface ProjectSinkArgs {
    /**
     * The destination of the sink (or, in other words, where logs are written to). Can be a
     * Cloud Storage bucket, a PubSub topic, or a BigQuery dataset. Examples:
     * ```
     * "storage.googleapis.com/[GCS_BUCKET]"
     * "bigquery.googleapis.com/projects/[PROJECT_ID]/datasets/[DATASET]"
     * "pubsub.googleapis.com/projects/[PROJECT_ID]/topics/[TOPIC_ID]"
     * ```
     * The writer associated with the sink must have access to write to the above resource.
     */
    readonly destination: pulumi.Input<string>;
    /**
     * The filter to apply when exporting logs. Only log entries that match the filter are exported.
     * See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
     * write a filter.
     */
    readonly filter?: pulumi.Input<string>;
    /**
     * The name of the logging sink.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project to create the sink in. If omitted, the project associated with the provider is
     * used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Whether or not to create a unique identity associated with this sink. If `false`
     * (the default), then the `writer_identity` used is `serviceAccount:cloud-logs@system.gserviceaccount.com`. If `true`,
     * then a unique service account is created and used for this sink. If you wish to publish logs across projects, you
     * must set `unique_writer_identity` to true.
     */
    readonly uniqueWriterIdentity?: pulumi.Input<boolean>;
}
