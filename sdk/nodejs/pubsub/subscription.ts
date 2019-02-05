// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Creates a subscription in Google's pubsub queueing system. For more information see
 * [the official documentation](https://cloud.google.com/pubsub/docs) and
 * [API](https://cloud.google.com/pubsub/docs/reference/rest/v1/projects.subscriptions).
 * 
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const google_pubsub_topic_default_topic = new gcp.pubsub.Topic("default-topic", {
 *     name: "default-topic",
 * });
 * const google_pubsub_subscription_default = new gcp.pubsub.Subscription("default", {
 *     ackDeadlineSeconds: 20,
 *     name: "default-subscription",
 *     pushConfig: {
 *         attributes: {
 *             "x-goog-version": "v1",
 *         },
 *         pushEndpoint: "https://example.com/push",
 *     },
 *     topic: google_pubsub_topic_default_topic.name,
 * });
 * ```
 * If the subscription has a topic in a different project:
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const google_pubsub_topic_topic_different_project = new gcp.pubsub.Topic("topic-different-project", {
 *     name: "topic-different-project",
 *     project: "another-project",
 * });
 * const google_pubsub_subscription_default = new gcp.pubsub.Subscription("default", {
 *     name: "default-subscription",
 *     topic: google_pubsub_topic_topic_different_project.id,
 * });
 * ```
 */
export class Subscription extends pulumi.CustomResource {
    /**
     * Get an existing Subscription resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SubscriptionState, opts?: pulumi.CustomResourceOptions): Subscription {
        return new Subscription(name, <any>state, { ...opts, id: id });
    }

    /**
     * The maximum number of seconds a
     * subscriber has to acknowledge a received message, otherwise the message is
     * redelivered. Changing this forces a new resource to be created.
     */
    public readonly ackDeadlineSeconds: pulumi.Output<number>;
    /**
     * A unique name for the resource, required by pubsub.
     * Changing this forces a new resource to be created.
     */
    public readonly name: pulumi.Output<string>;
    /**
     * Path of the subscription in the format `projects/{project}/subscriptions/{sub}`
     */
    public /*out*/ readonly path: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    public readonly project: pulumi.Output<string>;
    /**
     * Block configuration for push options. More
     * configuration options are detailed below.
     */
    public readonly pushConfig: pulumi.Output<{ attributes?: {[key: string]: string}, pushEndpoint: string } | undefined>;
    /**
     * The topic name or id to bind this subscription to, required by pubsub.
     * Changing this forces a new resource to be created.
     */
    public readonly topic: pulumi.Output<string>;

    /**
     * Create a Subscription resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SubscriptionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SubscriptionArgs | SubscriptionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: SubscriptionState = argsOrState as SubscriptionState | undefined;
            inputs["ackDeadlineSeconds"] = state ? state.ackDeadlineSeconds : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["path"] = state ? state.path : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["pushConfig"] = state ? state.pushConfig : undefined;
            inputs["topic"] = state ? state.topic : undefined;
        } else {
            const args = argsOrState as SubscriptionArgs | undefined;
            if (!args || args.topic === undefined) {
                throw new Error("Missing required property 'topic'");
            }
            inputs["ackDeadlineSeconds"] = args ? args.ackDeadlineSeconds : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["pushConfig"] = args ? args.pushConfig : undefined;
            inputs["topic"] = args ? args.topic : undefined;
            inputs["path"] = undefined /*out*/;
        }
        super("gcp:pubsub/subscription:Subscription", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Subscription resources.
 */
export interface SubscriptionState {
    /**
     * The maximum number of seconds a
     * subscriber has to acknowledge a received message, otherwise the message is
     * redelivered. Changing this forces a new resource to be created.
     */
    readonly ackDeadlineSeconds?: pulumi.Input<number>;
    /**
     * A unique name for the resource, required by pubsub.
     * Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Path of the subscription in the format `projects/{project}/subscriptions/{sub}`
     */
    readonly path?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Block configuration for push options. More
     * configuration options are detailed below.
     */
    readonly pushConfig?: pulumi.Input<{ attributes?: pulumi.Input<{[key: string]: pulumi.Input<string>}>, pushEndpoint: pulumi.Input<string> }>;
    /**
     * The topic name or id to bind this subscription to, required by pubsub.
     * Changing this forces a new resource to be created.
     */
    readonly topic?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Subscription resource.
 */
export interface SubscriptionArgs {
    /**
     * The maximum number of seconds a
     * subscriber has to acknowledge a received message, otherwise the message is
     * redelivered. Changing this forces a new resource to be created.
     */
    readonly ackDeadlineSeconds?: pulumi.Input<number>;
    /**
     * A unique name for the resource, required by pubsub.
     * Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Block configuration for push options. More
     * configuration options are detailed below.
     */
    readonly pushConfig?: pulumi.Input<{ attributes?: pulumi.Input<{[key: string]: pulumi.Input<string>}>, pushEndpoint: pulumi.Input<string> }>;
    /**
     * The topic name or id to bind this subscription to, required by pubsub.
     * Changing this forces a new resource to be created.
     */
    readonly topic: pulumi.Input<string>;
}
