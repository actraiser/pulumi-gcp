// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";

/**
 * Manages a health check within GCE. This is used to monitor instances
 * behind load balancers. Timeouts or HTTP errors cause the instance to be
 * removed from the pool. For more information, see [the official
 * documentation](https://cloud.google.com/compute/docs/load-balancing/health-checks)
 * and
 * [API](https://cloud.google.com/compute/docs/reference/latest/healthChecks).
 */
export class HealthCheck extends pulumi.CustomResource {
    /**
     * Get an existing HealthCheck resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: HealthCheckState): HealthCheck {
        return new HealthCheck(name, <any>state, { id });
    }

    /**
     * The number of seconds between each poll of
     * the instance instance (default 5).
     */
    public readonly checkIntervalSec: pulumi.Output<number | undefined>;
    /**
     * Textual description field.
     */
    public readonly description: pulumi.Output<string | undefined>;
    /**
     * Consecutive successes required (default 2).
     */
    public readonly healthyThreshold: pulumi.Output<number | undefined>;
    /**
     * An HTTP Health Check. Only one kind of Health Check can be added.
     * Structure is documented below.
     */
    public readonly httpHealthCheck: pulumi.Output<{ host?: string, port?: number, proxyHeader?: string, requestPath?: string } | undefined>;
    /**
     * An HTTPS Health Check. Only one kind of Health Check can be added.
     * Structure is documented below.
     */
    public readonly httpsHealthCheck: pulumi.Output<{ host?: string, port?: number, proxyHeader?: string, requestPath?: string } | undefined>;
    /**
     * A unique name for the resource, required by GCE.
     * Changing this forces a new resource to be created.
     */
    public readonly name: pulumi.Output<string>;
    /**
     * The project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    public readonly project: pulumi.Output<string>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink: pulumi.Output<string>;
    /**
     * An SSL Health Check. Only one kind of Health Check can be added.
     * Structure is documented below.
     */
    public readonly sslHealthCheck: pulumi.Output<{ port?: number, proxyHeader?: string, request?: string, response?: string } | undefined>;
    /**
     * A TCP Health Check. Only one kind of Health Check can be added.
     * Structure is documented below.
     */
    public readonly tcpHealthCheck: pulumi.Output<{ port?: number, proxyHeader?: string, request?: string, response?: string } | undefined>;
    /**
     * The number of seconds to wait before declaring
     * failure (default 5).
     */
    public readonly timeoutSec: pulumi.Output<number | undefined>;
    /**
     * Consecutive failures required (default 2).
     */
    public readonly unhealthyThreshold: pulumi.Output<number | undefined>;

    /**
     * Create a HealthCheck resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: HealthCheckArgs, opts?: pulumi.ResourceOptions)
    constructor(name: string, argsOrState?: HealthCheckArgs | HealthCheckState, opts?: pulumi.ResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: HealthCheckState = argsOrState as HealthCheckState | undefined;
            inputs["checkIntervalSec"] = state ? state.checkIntervalSec : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["healthyThreshold"] = state ? state.healthyThreshold : undefined;
            inputs["httpHealthCheck"] = state ? state.httpHealthCheck : undefined;
            inputs["httpsHealthCheck"] = state ? state.httpsHealthCheck : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
            inputs["sslHealthCheck"] = state ? state.sslHealthCheck : undefined;
            inputs["tcpHealthCheck"] = state ? state.tcpHealthCheck : undefined;
            inputs["timeoutSec"] = state ? state.timeoutSec : undefined;
            inputs["unhealthyThreshold"] = state ? state.unhealthyThreshold : undefined;
        } else {
            const args = argsOrState as HealthCheckArgs | undefined;
            inputs["checkIntervalSec"] = args ? args.checkIntervalSec : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["healthyThreshold"] = args ? args.healthyThreshold : undefined;
            inputs["httpHealthCheck"] = args ? args.httpHealthCheck : undefined;
            inputs["httpsHealthCheck"] = args ? args.httpsHealthCheck : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["sslHealthCheck"] = args ? args.sslHealthCheck : undefined;
            inputs["tcpHealthCheck"] = args ? args.tcpHealthCheck : undefined;
            inputs["timeoutSec"] = args ? args.timeoutSec : undefined;
            inputs["unhealthyThreshold"] = args ? args.unhealthyThreshold : undefined;
            inputs["selfLink"] = undefined /*out*/;
        }
        super("gcp:compute/healthCheck:HealthCheck", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering HealthCheck resources.
 */
export interface HealthCheckState {
    /**
     * The number of seconds between each poll of
     * the instance instance (default 5).
     */
    readonly checkIntervalSec?: pulumi.Input<number>;
    /**
     * Textual description field.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Consecutive successes required (default 2).
     */
    readonly healthyThreshold?: pulumi.Input<number>;
    /**
     * An HTTP Health Check. Only one kind of Health Check can be added.
     * Structure is documented below.
     */
    readonly httpHealthCheck?: pulumi.Input<{ host?: pulumi.Input<string>, port?: pulumi.Input<number>, proxyHeader?: pulumi.Input<string>, requestPath?: pulumi.Input<string> }>;
    /**
     * An HTTPS Health Check. Only one kind of Health Check can be added.
     * Structure is documented below.
     */
    readonly httpsHealthCheck?: pulumi.Input<{ host?: pulumi.Input<string>, port?: pulumi.Input<number>, proxyHeader?: pulumi.Input<string>, requestPath?: pulumi.Input<string> }>;
    /**
     * A unique name for the resource, required by GCE.
     * Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The URI of the created resource.
     */
    readonly selfLink?: pulumi.Input<string>;
    /**
     * An SSL Health Check. Only one kind of Health Check can be added.
     * Structure is documented below.
     */
    readonly sslHealthCheck?: pulumi.Input<{ port?: pulumi.Input<number>, proxyHeader?: pulumi.Input<string>, request?: pulumi.Input<string>, response?: pulumi.Input<string> }>;
    /**
     * A TCP Health Check. Only one kind of Health Check can be added.
     * Structure is documented below.
     */
    readonly tcpHealthCheck?: pulumi.Input<{ port?: pulumi.Input<number>, proxyHeader?: pulumi.Input<string>, request?: pulumi.Input<string>, response?: pulumi.Input<string> }>;
    /**
     * The number of seconds to wait before declaring
     * failure (default 5).
     */
    readonly timeoutSec?: pulumi.Input<number>;
    /**
     * Consecutive failures required (default 2).
     */
    readonly unhealthyThreshold?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a HealthCheck resource.
 */
export interface HealthCheckArgs {
    /**
     * The number of seconds between each poll of
     * the instance instance (default 5).
     */
    readonly checkIntervalSec?: pulumi.Input<number>;
    /**
     * Textual description field.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Consecutive successes required (default 2).
     */
    readonly healthyThreshold?: pulumi.Input<number>;
    /**
     * An HTTP Health Check. Only one kind of Health Check can be added.
     * Structure is documented below.
     */
    readonly httpHealthCheck?: pulumi.Input<{ host?: pulumi.Input<string>, port?: pulumi.Input<number>, proxyHeader?: pulumi.Input<string>, requestPath?: pulumi.Input<string> }>;
    /**
     * An HTTPS Health Check. Only one kind of Health Check can be added.
     * Structure is documented below.
     */
    readonly httpsHealthCheck?: pulumi.Input<{ host?: pulumi.Input<string>, port?: pulumi.Input<number>, proxyHeader?: pulumi.Input<string>, requestPath?: pulumi.Input<string> }>;
    /**
     * A unique name for the resource, required by GCE.
     * Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * An SSL Health Check. Only one kind of Health Check can be added.
     * Structure is documented below.
     */
    readonly sslHealthCheck?: pulumi.Input<{ port?: pulumi.Input<number>, proxyHeader?: pulumi.Input<string>, request?: pulumi.Input<string>, response?: pulumi.Input<string> }>;
    /**
     * A TCP Health Check. Only one kind of Health Check can be added.
     * Structure is documented below.
     */
    readonly tcpHealthCheck?: pulumi.Input<{ port?: pulumi.Input<number>, proxyHeader?: pulumi.Input<string>, request?: pulumi.Input<string>, response?: pulumi.Input<string> }>;
    /**
     * The number of seconds to wait before declaring
     * failure (default 5).
     */
    readonly timeoutSec?: pulumi.Input<number>;
    /**
     * Consecutive failures required (default 2).
     */
    readonly unhealthyThreshold?: pulumi.Input<number>;
}
