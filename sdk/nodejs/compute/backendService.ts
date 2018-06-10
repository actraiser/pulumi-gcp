// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";

/**
 * A Backend Service defines a group of virtual machines that will serve traffic for load balancing. For more information
 * see [the official documentation](https://cloud.google.com/compute/docs/load-balancing/http/backend-service)
 * and the [API](https://cloud.google.com/compute/docs/reference/latest/backendServices).
 * 
 * For internal load balancing, use a [google_compute_region_backend_service](/docs/providers/google/r/compute_region_backend_service.html).
 */
export class BackendService extends pulumi.CustomResource {
    /**
     * Get an existing BackendService resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BackendServiceState): BackendService {
        return new BackendService(name, <any>state, { id });
    }

    /**
     * The list of backends that serve this BackendService. Structure is documented below.
     */
    public readonly backends: pulumi.Output<{ balancingMode?: string, capacityScaler?: number, description?: string, group?: string, maxConnections?: number, maxConnectionsPerInstance?: number, maxRate?: number, maxRatePerInstance?: number, maxUtilization?: number }[] | undefined>;
    /**
     * Cloud CDN configuration for this BackendService. Structure is documented below.
     */
    public readonly cdnPolicy: pulumi.Output<{ cacheKeyPolicy?: { includeHost?: boolean, includeProtocol?: boolean, includeQueryString?: boolean, queryStringBlacklists?: string[], queryStringWhitelists?: string[] } }>;
    /**
     * Time for which instance will be drained (not accept new connections,
     * but still work to finish started ones). Defaults to `300`.
     */
    public readonly connectionDrainingTimeoutSec: pulumi.Output<number | undefined>;
    /**
     * Textual description for the backend.
     */
    public readonly description: pulumi.Output<string | undefined>;
    /**
     * Whether or not to enable the Cloud CDN on the backend service.
     */
    public readonly enableCdn: pulumi.Output<boolean | undefined>;
    /**
     * The fingerprint of the backend service.
     */
    public /*out*/ readonly fingerprint: pulumi.Output<string>;
    /**
     * Specifies a list of HTTP/HTTPS health checks
     * for checking the health of the backend service. Currently at most one health
     * check can be specified, and a health check is required.
     */
    public readonly healthChecks: pulumi.Output<string>;
    /**
     * Specification for the Identity-Aware proxy. Disabled if not specified. Structure is documented below.
     */
    public readonly iap: pulumi.Output<{ oauth2ClientId: string, oauth2ClientSecret: string } | undefined>;
    /**
     * The name of the backend service.
     */
    public readonly name: pulumi.Output<string>;
    /**
     * The name of a service that has been added to an
     * instance group in this backend. See [related docs](https://cloud.google.com/compute/docs/instance-groups/#specifying_service_endpoints) for details. Defaults to http.
     */
    public readonly portName: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    public readonly project: pulumi.Output<string>;
    /**
     * The protocol for incoming requests. Defaults to
     * `HTTP`.
     */
    public readonly protocol: pulumi.Output<string>;
    /**
     * ) Name or URI of a
     * [security policy](https://cloud.google.com/armor/docs/security-policy-concepts) to add to the backend service.
     */
    public readonly securityPolicy: pulumi.Output<string | undefined>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink: pulumi.Output<string>;
    /**
     * How to distribute load. Options are `NONE` (no
     * affinity), `CLIENT_IP` (hash of the source/dest addresses / ports), and
     * `GENERATED_COOKIE` (distribute load using a generated session cookie).
     */
    public readonly sessionAffinity: pulumi.Output<string>;
    /**
     * The number of secs to wait for a backend to respond
     * to a request before considering the request failed. Defaults to `30`.
     */
    public readonly timeoutSec: pulumi.Output<number>;

    /**
     * Create a BackendService resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BackendServiceArgs, opts?: pulumi.ResourceOptions)
    constructor(name: string, argsOrState?: BackendServiceArgs | BackendServiceState, opts?: pulumi.ResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: BackendServiceState = argsOrState as BackendServiceState | undefined;
            inputs["backends"] = state ? state.backends : undefined;
            inputs["cdnPolicy"] = state ? state.cdnPolicy : undefined;
            inputs["connectionDrainingTimeoutSec"] = state ? state.connectionDrainingTimeoutSec : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["enableCdn"] = state ? state.enableCdn : undefined;
            inputs["fingerprint"] = state ? state.fingerprint : undefined;
            inputs["healthChecks"] = state ? state.healthChecks : undefined;
            inputs["iap"] = state ? state.iap : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["portName"] = state ? state.portName : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["protocol"] = state ? state.protocol : undefined;
            inputs["securityPolicy"] = state ? state.securityPolicy : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
            inputs["sessionAffinity"] = state ? state.sessionAffinity : undefined;
            inputs["timeoutSec"] = state ? state.timeoutSec : undefined;
        } else {
            const args = argsOrState as BackendServiceArgs | undefined;
            if (!args || args.healthChecks === undefined) {
                throw new Error("Missing required property 'healthChecks'");
            }
            inputs["backends"] = args ? args.backends : undefined;
            inputs["cdnPolicy"] = args ? args.cdnPolicy : undefined;
            inputs["connectionDrainingTimeoutSec"] = args ? args.connectionDrainingTimeoutSec : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["enableCdn"] = args ? args.enableCdn : undefined;
            inputs["healthChecks"] = args ? args.healthChecks : undefined;
            inputs["iap"] = args ? args.iap : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["portName"] = args ? args.portName : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["protocol"] = args ? args.protocol : undefined;
            inputs["securityPolicy"] = args ? args.securityPolicy : undefined;
            inputs["sessionAffinity"] = args ? args.sessionAffinity : undefined;
            inputs["timeoutSec"] = args ? args.timeoutSec : undefined;
            inputs["fingerprint"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
        }
        super("gcp:compute/backendService:BackendService", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BackendService resources.
 */
export interface BackendServiceState {
    /**
     * The list of backends that serve this BackendService. Structure is documented below.
     */
    readonly backends?: pulumi.Input<{ balancingMode?: pulumi.Input<string>, capacityScaler?: pulumi.Input<number>, description?: pulumi.Input<string>, group?: pulumi.Input<string>, maxConnections?: pulumi.Input<number>, maxConnectionsPerInstance?: pulumi.Input<number>, maxRate?: pulumi.Input<number>, maxRatePerInstance?: pulumi.Input<number>, maxUtilization?: pulumi.Input<number> }[]>;
    /**
     * Cloud CDN configuration for this BackendService. Structure is documented below.
     */
    readonly cdnPolicy?: pulumi.Input<{ cacheKeyPolicy?: pulumi.Input<{ includeHost?: pulumi.Input<boolean>, includeProtocol?: pulumi.Input<boolean>, includeQueryString?: pulumi.Input<boolean>, queryStringBlacklists?: pulumi.Input<pulumi.Input<string>[]>, queryStringWhitelists?: pulumi.Input<pulumi.Input<string>[]> }> }>;
    /**
     * Time for which instance will be drained (not accept new connections,
     * but still work to finish started ones). Defaults to `300`.
     */
    readonly connectionDrainingTimeoutSec?: pulumi.Input<number>;
    /**
     * Textual description for the backend.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Whether or not to enable the Cloud CDN on the backend service.
     */
    readonly enableCdn?: pulumi.Input<boolean>;
    /**
     * The fingerprint of the backend service.
     */
    readonly fingerprint?: pulumi.Input<string>;
    /**
     * Specifies a list of HTTP/HTTPS health checks
     * for checking the health of the backend service. Currently at most one health
     * check can be specified, and a health check is required.
     */
    readonly healthChecks?: pulumi.Input<pulumi.Input<string>>;
    /**
     * Specification for the Identity-Aware proxy. Disabled if not specified. Structure is documented below.
     */
    readonly iap?: pulumi.Input<{ oauth2ClientId: pulumi.Input<string>, oauth2ClientSecret: pulumi.Input<string> }>;
    /**
     * The name of the backend service.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of a service that has been added to an
     * instance group in this backend. See [related docs](https://cloud.google.com/compute/docs/instance-groups/#specifying_service_endpoints) for details. Defaults to http.
     */
    readonly portName?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The protocol for incoming requests. Defaults to
     * `HTTP`.
     */
    readonly protocol?: pulumi.Input<string>;
    /**
     * ) Name or URI of a
     * [security policy](https://cloud.google.com/armor/docs/security-policy-concepts) to add to the backend service.
     */
    readonly securityPolicy?: pulumi.Input<string>;
    /**
     * The URI of the created resource.
     */
    readonly selfLink?: pulumi.Input<string>;
    /**
     * How to distribute load. Options are `NONE` (no
     * affinity), `CLIENT_IP` (hash of the source/dest addresses / ports), and
     * `GENERATED_COOKIE` (distribute load using a generated session cookie).
     */
    readonly sessionAffinity?: pulumi.Input<string>;
    /**
     * The number of secs to wait for a backend to respond
     * to a request before considering the request failed. Defaults to `30`.
     */
    readonly timeoutSec?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a BackendService resource.
 */
export interface BackendServiceArgs {
    /**
     * The list of backends that serve this BackendService. Structure is documented below.
     */
    readonly backends?: pulumi.Input<{ balancingMode?: pulumi.Input<string>, capacityScaler?: pulumi.Input<number>, description?: pulumi.Input<string>, group?: pulumi.Input<string>, maxConnections?: pulumi.Input<number>, maxConnectionsPerInstance?: pulumi.Input<number>, maxRate?: pulumi.Input<number>, maxRatePerInstance?: pulumi.Input<number>, maxUtilization?: pulumi.Input<number> }[]>;
    /**
     * Cloud CDN configuration for this BackendService. Structure is documented below.
     */
    readonly cdnPolicy?: pulumi.Input<{ cacheKeyPolicy?: pulumi.Input<{ includeHost?: pulumi.Input<boolean>, includeProtocol?: pulumi.Input<boolean>, includeQueryString?: pulumi.Input<boolean>, queryStringBlacklists?: pulumi.Input<pulumi.Input<string>[]>, queryStringWhitelists?: pulumi.Input<pulumi.Input<string>[]> }> }>;
    /**
     * Time for which instance will be drained (not accept new connections,
     * but still work to finish started ones). Defaults to `300`.
     */
    readonly connectionDrainingTimeoutSec?: pulumi.Input<number>;
    /**
     * Textual description for the backend.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Whether or not to enable the Cloud CDN on the backend service.
     */
    readonly enableCdn?: pulumi.Input<boolean>;
    /**
     * Specifies a list of HTTP/HTTPS health checks
     * for checking the health of the backend service. Currently at most one health
     * check can be specified, and a health check is required.
     */
    readonly healthChecks: pulumi.Input<pulumi.Input<string>>;
    /**
     * Specification for the Identity-Aware proxy. Disabled if not specified. Structure is documented below.
     */
    readonly iap?: pulumi.Input<{ oauth2ClientId: pulumi.Input<string>, oauth2ClientSecret: pulumi.Input<string> }>;
    /**
     * The name of the backend service.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of a service that has been added to an
     * instance group in this backend. See [related docs](https://cloud.google.com/compute/docs/instance-groups/#specifying_service_endpoints) for details. Defaults to http.
     */
    readonly portName?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The protocol for incoming requests. Defaults to
     * `HTTP`.
     */
    readonly protocol?: pulumi.Input<string>;
    /**
     * ) Name or URI of a
     * [security policy](https://cloud.google.com/armor/docs/security-policy-concepts) to add to the backend service.
     */
    readonly securityPolicy?: pulumi.Input<string>;
    /**
     * How to distribute load. Options are `NONE` (no
     * affinity), `CLIENT_IP` (hash of the source/dest addresses / ports), and
     * `GENERATED_COOKIE` (distribute load using a generated session cookie).
     */
    readonly sessionAffinity?: pulumi.Input<string>;
    /**
     * The number of secs to wait for a backend to respond
     * to a request before considering the request failed. Defaults to `30`.
     */
    readonly timeoutSec?: pulumi.Input<number>;
}
