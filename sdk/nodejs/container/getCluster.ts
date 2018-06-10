// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";

/**
 * Get info about a cluster within GKE from its name and zone.
 */
export function getCluster(args: GetClusterArgs): Promise<GetClusterResult> {
    return pulumi.runtime.invoke("gcp:container/getCluster:getCluster", {
        "name": args.name,
        "project": args.project,
        "region": args.region,
        "zone": args.zone,
    });
}

/**
 * A collection of arguments for invoking getCluster.
 */
export interface GetClusterArgs {
    /**
     * The name of the cluster.
     */
    readonly name: pulumi.Input<string>;
    /**
     * The project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    readonly region?: pulumi.Input<string>;
    readonly zone?: pulumi.Input<string>;
}

/**
 * A collection of values returned by getCluster.
 */
export interface GetClusterResult {
    readonly additionalZones: string[];
    readonly addonsConfigs: { horizontalPodAutoscalings: { disabled: boolean }[], httpLoadBalancings: { disabled: boolean }[], kubernetesDashboards: { disabled: boolean }[], networkPolicyConfigs: { disabled: boolean }[] }[];
    readonly clusterIpv4Cidr: string;
    readonly description: string;
    readonly enableKubernetesAlpha: boolean;
    readonly enableLegacyAbac: boolean;
    readonly endpoint: string;
    readonly initialNodeCount: number;
    readonly instanceGroupUrls: string[];
    readonly ipAllocationPolicies: { clusterSecondaryRangeName: string, servicesSecondaryRangeName: string }[];
    readonly loggingService: string;
    readonly maintenancePolicies: { dailyMaintenanceWindows: { duration: string, startTime: string }[] }[];
    readonly masterAuths: { clientCertificate: string, clientCertificateConfigs: { issueClientCertificate: boolean }[], clientKey: string, clusterCaCertificate: string, password: string, username: string }[];
    readonly masterAuthorizedNetworksConfigs: { cidrBlocks: { cidrBlock: string, displayName: string }[] }[];
    readonly masterIpv4CidrBlock: string;
    readonly masterVersion: string;
    readonly minMasterVersion: string;
    readonly monitoringService: string;
    readonly network: string;
    readonly networkPolicies: { enabled: boolean, provider: string }[];
    readonly nodeConfigs: { diskSizeGb: number, guestAccelerators: { count: number, type: string }[], imageType: string, labels: {[key: string]: string}, localSsdCount: number, machineType: string, metadata: {[key: string]: string}, minCpuPlatform: string, oauthScopes: string[], preemptible: boolean, serviceAccount: string, tags: string[], taints: { effect: string, key: string, value: string }[], workloadMetadataConfigs: { nodeMetadata: string }[] }[];
    readonly nodePools: { autoscalings: { maxNodeCount: number, minNodeCount: number }[], initialNodeCount: number, instanceGroupUrls: string[], managements: { autoRepair: boolean, autoUpgrade: boolean }[], name: string, namePrefix: string, nodeConfigs: { diskSizeGb: number, guestAccelerators: { count: number, type: string }[], imageType: string, labels: {[key: string]: string}, localSsdCount: number, machineType: string, metadata: {[key: string]: string}, minCpuPlatform: string, oauthScopes: string[], preemptible: boolean, serviceAccount: string, tags: string[], taints: { effect: string, key: string, value: string }[], workloadMetadataConfigs: { nodeMetadata: string }[] }[], nodeCount: number, version: string }[];
    readonly nodeVersion: string;
    readonly podSecurityPolicyConfigs: { enabled: boolean }[];
    readonly privateCluster: boolean;
    readonly removeDefaultNodePool: boolean;
    readonly subnetwork: string;
}
