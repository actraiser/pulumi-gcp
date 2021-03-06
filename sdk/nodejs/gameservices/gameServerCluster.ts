// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * A game server cluster resource.
 * 
 * To get more information about GameServerCluster, see:
 * 
 * * [API documentation](https://cloud.google.com/game-servers/docs/reference/rest/v1beta/projects.locations.realms.gameServerClusters)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/game-servers/docs)
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/game_services_game_server_cluster.html.markdown.
 */
export class GameServerCluster extends pulumi.CustomResource {
    /**
     * Get an existing GameServerCluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GameServerClusterState, opts?: pulumi.CustomResourceOptions): GameServerCluster {
        return new GameServerCluster(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:gameservices/gameServerCluster:GameServerCluster';

    /**
     * Returns true if the given object is an instance of GameServerCluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GameServerCluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GameServerCluster.__pulumiType;
    }

    /**
     * Required. The resource name of the game server cluster
     */
    public readonly clusterId!: pulumi.Output<string>;
    /**
     * Game server cluster connection information. This information is used to
     * manage game server clusters.  Structure is documented below.
     */
    public readonly connectionInfo!: pulumi.Output<outputs.gameservices.GameServerClusterConnectionInfo>;
    /**
     * Human readable description of the cluster.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The labels associated with this game server cluster. Each label is a
     * key-value pair.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Location of the Cluster.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * The resource id of the game server cluster, eg:
     * 'projects/{project_id}/locations/{location}/realms/{realm_id}/gameServerClusters/{cluster_id}'. For example,
     * 'projects/my-project/locations/{location}/realms/zanzibar/gameServerClusters/my-onprem-cluster'.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The realm id of the game server realm.
     */
    public readonly realmId!: pulumi.Output<string>;

    /**
     * Create a GameServerCluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GameServerClusterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GameServerClusterArgs | GameServerClusterState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as GameServerClusterState | undefined;
            inputs["clusterId"] = state ? state.clusterId : undefined;
            inputs["connectionInfo"] = state ? state.connectionInfo : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["realmId"] = state ? state.realmId : undefined;
        } else {
            const args = argsOrState as GameServerClusterArgs | undefined;
            if (!args || args.clusterId === undefined) {
                throw new Error("Missing required property 'clusterId'");
            }
            if (!args || args.connectionInfo === undefined) {
                throw new Error("Missing required property 'connectionInfo'");
            }
            if (!args || args.realmId === undefined) {
                throw new Error("Missing required property 'realmId'");
            }
            inputs["clusterId"] = args ? args.clusterId : undefined;
            inputs["connectionInfo"] = args ? args.connectionInfo : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["realmId"] = args ? args.realmId : undefined;
            inputs["name"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(GameServerCluster.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GameServerCluster resources.
 */
export interface GameServerClusterState {
    /**
     * Required. The resource name of the game server cluster
     */
    readonly clusterId?: pulumi.Input<string>;
    /**
     * Game server cluster connection information. This information is used to
     * manage game server clusters.  Structure is documented below.
     */
    readonly connectionInfo?: pulumi.Input<inputs.gameservices.GameServerClusterConnectionInfo>;
    /**
     * Human readable description of the cluster.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The labels associated with this game server cluster. Each label is a
     * key-value pair.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Location of the Cluster.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The resource id of the game server cluster, eg:
     * 'projects/{project_id}/locations/{location}/realms/{realm_id}/gameServerClusters/{cluster_id}'. For example,
     * 'projects/my-project/locations/{location}/realms/zanzibar/gameServerClusters/my-onprem-cluster'.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The realm id of the game server realm.
     */
    readonly realmId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a GameServerCluster resource.
 */
export interface GameServerClusterArgs {
    /**
     * Required. The resource name of the game server cluster
     */
    readonly clusterId: pulumi.Input<string>;
    /**
     * Game server cluster connection information. This information is used to
     * manage game server clusters.  Structure is documented below.
     */
    readonly connectionInfo: pulumi.Input<inputs.gameservices.GameServerClusterConnectionInfo>;
    /**
     * Human readable description of the cluster.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The labels associated with this game server cluster. Each label is a
     * key-value pair.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Location of the Cluster.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The realm id of the game server realm.
     */
    readonly realmId: pulumi.Input<string>;
}
