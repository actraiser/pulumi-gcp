// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Container
{
    public static partial class Invokes
    {
        /// <summary>
        /// Get info about a GKE cluster from its name and location.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/container_cluster.html.markdown.
        /// </summary>
        public static Task<GetClusterResult> GetCluster(GetClusterArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetClusterResult>("gcp:container/getCluster:getCluster", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetClusterArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The location (zone or region) this cluster has been
        /// created in. One of `location`, `region`, `zone`, or a provider-level `zone` must
        /// be specified.
        /// </summary>
        [Input("location")]
        public string? Location { get; set; }

        /// <summary>
        /// The name of the cluster.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public string? Project { get; set; }

        /// <summary>
        /// The region this cluster has been created in. Deprecated
        /// in favour of `location`.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        /// <summary>
        /// The zone this cluster has been created in. Deprecated in
        /// favour of `location`.
        /// </summary>
        [Input("zone")]
        public string? Zone { get; set; }

        public GetClusterArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetClusterResult
    {
        public readonly ImmutableArray<string> AdditionalZones;
        public readonly ImmutableArray<Outputs.GetClusterAddonsConfigsResult> AddonsConfigs;
        public readonly ImmutableArray<Outputs.GetClusterAuthenticatorGroupsConfigsResult> AuthenticatorGroupsConfigs;
        public readonly ImmutableArray<Outputs.GetClusterClusterAutoscalingsResult> ClusterAutoscalings;
        public readonly string ClusterIpv4Cidr;
        public readonly ImmutableArray<Outputs.GetClusterDatabaseEncryptionsResult> DatabaseEncryptions;
        public readonly int DefaultMaxPodsPerNode;
        public readonly string Description;
        public readonly bool EnableBinaryAuthorization;
        public readonly bool EnableIntranodeVisibility;
        public readonly bool EnableKubernetesAlpha;
        public readonly bool EnableLegacyAbac;
        public readonly bool EnableShieldedNodes;
        public readonly bool EnableTpu;
        public readonly string Endpoint;
        public readonly int InitialNodeCount;
        public readonly ImmutableArray<string> InstanceGroupUrls;
        public readonly ImmutableArray<Outputs.GetClusterIpAllocationPoliciesResult> IpAllocationPolicies;
        public readonly string? Location;
        public readonly string LoggingService;
        public readonly ImmutableArray<Outputs.GetClusterMaintenancePoliciesResult> MaintenancePolicies;
        public readonly ImmutableArray<Outputs.GetClusterMasterAuthsResult> MasterAuths;
        public readonly ImmutableArray<Outputs.GetClusterMasterAuthorizedNetworksConfigsResult> MasterAuthorizedNetworksConfigs;
        public readonly string MasterVersion;
        public readonly string MinMasterVersion;
        public readonly string MonitoringService;
        public readonly string Name;
        public readonly string Network;
        public readonly ImmutableArray<Outputs.GetClusterNetworkPoliciesResult> NetworkPolicies;
        public readonly ImmutableArray<Outputs.GetClusterNodeConfigsResult> NodeConfigs;
        public readonly ImmutableArray<string> NodeLocations;
        public readonly ImmutableArray<Outputs.GetClusterNodePoolsResult> NodePools;
        public readonly string NodeVersion;
        public readonly ImmutableArray<Outputs.GetClusterPodSecurityPolicyConfigsResult> PodSecurityPolicyConfigs;
        public readonly ImmutableArray<Outputs.GetClusterPrivateClusterConfigsResult> PrivateClusterConfigs;
        public readonly string? Project;
        public readonly string? Region;
        public readonly ImmutableArray<Outputs.GetClusterReleaseChannelsResult> ReleaseChannels;
        public readonly bool RemoveDefaultNodePool;
        public readonly ImmutableDictionary<string, string> ResourceLabels;
        public readonly ImmutableArray<Outputs.GetClusterResourceUsageExportConfigsResult> ResourceUsageExportConfigs;
        public readonly string ServicesIpv4Cidr;
        public readonly string Subnetwork;
        public readonly string TpuIpv4CidrBlock;
        public readonly ImmutableArray<Outputs.GetClusterVerticalPodAutoscalingsResult> VerticalPodAutoscalings;
        public readonly ImmutableArray<Outputs.GetClusterWorkloadIdentityConfigsResult> WorkloadIdentityConfigs;
        public readonly string? Zone;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetClusterResult(
            ImmutableArray<string> additionalZones,
            ImmutableArray<Outputs.GetClusterAddonsConfigsResult> addonsConfigs,
            ImmutableArray<Outputs.GetClusterAuthenticatorGroupsConfigsResult> authenticatorGroupsConfigs,
            ImmutableArray<Outputs.GetClusterClusterAutoscalingsResult> clusterAutoscalings,
            string clusterIpv4Cidr,
            ImmutableArray<Outputs.GetClusterDatabaseEncryptionsResult> databaseEncryptions,
            int defaultMaxPodsPerNode,
            string description,
            bool enableBinaryAuthorization,
            bool enableIntranodeVisibility,
            bool enableKubernetesAlpha,
            bool enableLegacyAbac,
            bool enableShieldedNodes,
            bool enableTpu,
            string endpoint,
            int initialNodeCount,
            ImmutableArray<string> instanceGroupUrls,
            ImmutableArray<Outputs.GetClusterIpAllocationPoliciesResult> ipAllocationPolicies,
            string? location,
            string loggingService,
            ImmutableArray<Outputs.GetClusterMaintenancePoliciesResult> maintenancePolicies,
            ImmutableArray<Outputs.GetClusterMasterAuthsResult> masterAuths,
            ImmutableArray<Outputs.GetClusterMasterAuthorizedNetworksConfigsResult> masterAuthorizedNetworksConfigs,
            string masterVersion,
            string minMasterVersion,
            string monitoringService,
            string name,
            string network,
            ImmutableArray<Outputs.GetClusterNetworkPoliciesResult> networkPolicies,
            ImmutableArray<Outputs.GetClusterNodeConfigsResult> nodeConfigs,
            ImmutableArray<string> nodeLocations,
            ImmutableArray<Outputs.GetClusterNodePoolsResult> nodePools,
            string nodeVersion,
            ImmutableArray<Outputs.GetClusterPodSecurityPolicyConfigsResult> podSecurityPolicyConfigs,
            ImmutableArray<Outputs.GetClusterPrivateClusterConfigsResult> privateClusterConfigs,
            string? project,
            string? region,
            ImmutableArray<Outputs.GetClusterReleaseChannelsResult> releaseChannels,
            bool removeDefaultNodePool,
            ImmutableDictionary<string, string> resourceLabels,
            ImmutableArray<Outputs.GetClusterResourceUsageExportConfigsResult> resourceUsageExportConfigs,
            string servicesIpv4Cidr,
            string subnetwork,
            string tpuIpv4CidrBlock,
            ImmutableArray<Outputs.GetClusterVerticalPodAutoscalingsResult> verticalPodAutoscalings,
            ImmutableArray<Outputs.GetClusterWorkloadIdentityConfigsResult> workloadIdentityConfigs,
            string? zone,
            string id)
        {
            AdditionalZones = additionalZones;
            AddonsConfigs = addonsConfigs;
            AuthenticatorGroupsConfigs = authenticatorGroupsConfigs;
            ClusterAutoscalings = clusterAutoscalings;
            ClusterIpv4Cidr = clusterIpv4Cidr;
            DatabaseEncryptions = databaseEncryptions;
            DefaultMaxPodsPerNode = defaultMaxPodsPerNode;
            Description = description;
            EnableBinaryAuthorization = enableBinaryAuthorization;
            EnableIntranodeVisibility = enableIntranodeVisibility;
            EnableKubernetesAlpha = enableKubernetesAlpha;
            EnableLegacyAbac = enableLegacyAbac;
            EnableShieldedNodes = enableShieldedNodes;
            EnableTpu = enableTpu;
            Endpoint = endpoint;
            InitialNodeCount = initialNodeCount;
            InstanceGroupUrls = instanceGroupUrls;
            IpAllocationPolicies = ipAllocationPolicies;
            Location = location;
            LoggingService = loggingService;
            MaintenancePolicies = maintenancePolicies;
            MasterAuths = masterAuths;
            MasterAuthorizedNetworksConfigs = masterAuthorizedNetworksConfigs;
            MasterVersion = masterVersion;
            MinMasterVersion = minMasterVersion;
            MonitoringService = monitoringService;
            Name = name;
            Network = network;
            NetworkPolicies = networkPolicies;
            NodeConfigs = nodeConfigs;
            NodeLocations = nodeLocations;
            NodePools = nodePools;
            NodeVersion = nodeVersion;
            PodSecurityPolicyConfigs = podSecurityPolicyConfigs;
            PrivateClusterConfigs = privateClusterConfigs;
            Project = project;
            Region = region;
            ReleaseChannels = releaseChannels;
            RemoveDefaultNodePool = removeDefaultNodePool;
            ResourceLabels = resourceLabels;
            ResourceUsageExportConfigs = resourceUsageExportConfigs;
            ServicesIpv4Cidr = servicesIpv4Cidr;
            Subnetwork = subnetwork;
            TpuIpv4CidrBlock = tpuIpv4CidrBlock;
            VerticalPodAutoscalings = verticalPodAutoscalings;
            WorkloadIdentityConfigs = workloadIdentityConfigs;
            Zone = zone;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetClusterAddonsConfigsCloudrunConfigsResult
    {
        public readonly bool Disabled;

        [OutputConstructor]
        private GetClusterAddonsConfigsCloudrunConfigsResult(bool disabled)
        {
            Disabled = disabled;
        }
    }

    [OutputType]
    public sealed class GetClusterAddonsConfigsHorizontalPodAutoscalingsResult
    {
        public readonly bool Disabled;

        [OutputConstructor]
        private GetClusterAddonsConfigsHorizontalPodAutoscalingsResult(bool disabled)
        {
            Disabled = disabled;
        }
    }

    [OutputType]
    public sealed class GetClusterAddonsConfigsHttpLoadBalancingsResult
    {
        public readonly bool Disabled;

        [OutputConstructor]
        private GetClusterAddonsConfigsHttpLoadBalancingsResult(bool disabled)
        {
            Disabled = disabled;
        }
    }

    [OutputType]
    public sealed class GetClusterAddonsConfigsIstioConfigsResult
    {
        public readonly string Auth;
        public readonly bool Disabled;

        [OutputConstructor]
        private GetClusterAddonsConfigsIstioConfigsResult(
            string auth,
            bool disabled)
        {
            Auth = auth;
            Disabled = disabled;
        }
    }

    [OutputType]
    public sealed class GetClusterAddonsConfigsKubernetesDashboardsResult
    {
        public readonly bool Disabled;

        [OutputConstructor]
        private GetClusterAddonsConfigsKubernetesDashboardsResult(bool disabled)
        {
            Disabled = disabled;
        }
    }

    [OutputType]
    public sealed class GetClusterAddonsConfigsNetworkPolicyConfigsResult
    {
        public readonly bool Disabled;

        [OutputConstructor]
        private GetClusterAddonsConfigsNetworkPolicyConfigsResult(bool disabled)
        {
            Disabled = disabled;
        }
    }

    [OutputType]
    public sealed class GetClusterAddonsConfigsResult
    {
        public readonly ImmutableArray<GetClusterAddonsConfigsCloudrunConfigsResult> CloudrunConfigs;
        public readonly ImmutableArray<GetClusterAddonsConfigsHorizontalPodAutoscalingsResult> HorizontalPodAutoscalings;
        public readonly ImmutableArray<GetClusterAddonsConfigsHttpLoadBalancingsResult> HttpLoadBalancings;
        public readonly ImmutableArray<GetClusterAddonsConfigsIstioConfigsResult> IstioConfigs;
        public readonly ImmutableArray<GetClusterAddonsConfigsKubernetesDashboardsResult> KubernetesDashboards;
        public readonly ImmutableArray<GetClusterAddonsConfigsNetworkPolicyConfigsResult> NetworkPolicyConfigs;

        [OutputConstructor]
        private GetClusterAddonsConfigsResult(
            ImmutableArray<GetClusterAddonsConfigsCloudrunConfigsResult> cloudrunConfigs,
            ImmutableArray<GetClusterAddonsConfigsHorizontalPodAutoscalingsResult> horizontalPodAutoscalings,
            ImmutableArray<GetClusterAddonsConfigsHttpLoadBalancingsResult> httpLoadBalancings,
            ImmutableArray<GetClusterAddonsConfigsIstioConfigsResult> istioConfigs,
            ImmutableArray<GetClusterAddonsConfigsKubernetesDashboardsResult> kubernetesDashboards,
            ImmutableArray<GetClusterAddonsConfigsNetworkPolicyConfigsResult> networkPolicyConfigs)
        {
            CloudrunConfigs = cloudrunConfigs;
            HorizontalPodAutoscalings = horizontalPodAutoscalings;
            HttpLoadBalancings = httpLoadBalancings;
            IstioConfigs = istioConfigs;
            KubernetesDashboards = kubernetesDashboards;
            NetworkPolicyConfigs = networkPolicyConfigs;
        }
    }

    [OutputType]
    public sealed class GetClusterAuthenticatorGroupsConfigsResult
    {
        public readonly string SecurityGroup;

        [OutputConstructor]
        private GetClusterAuthenticatorGroupsConfigsResult(string securityGroup)
        {
            SecurityGroup = securityGroup;
        }
    }

    [OutputType]
    public sealed class GetClusterClusterAutoscalingsAutoProvisioningDefaultsResult
    {
        public readonly ImmutableArray<string> OauthScopes;
        public readonly string ServiceAccount;

        [OutputConstructor]
        private GetClusterClusterAutoscalingsAutoProvisioningDefaultsResult(
            ImmutableArray<string> oauthScopes,
            string serviceAccount)
        {
            OauthScopes = oauthScopes;
            ServiceAccount = serviceAccount;
        }
    }

    [OutputType]
    public sealed class GetClusterClusterAutoscalingsResourceLimitsResult
    {
        public readonly int Maximum;
        public readonly int Minimum;
        public readonly string ResourceType;

        [OutputConstructor]
        private GetClusterClusterAutoscalingsResourceLimitsResult(
            int maximum,
            int minimum,
            string resourceType)
        {
            Maximum = maximum;
            Minimum = minimum;
            ResourceType = resourceType;
        }
    }

    [OutputType]
    public sealed class GetClusterClusterAutoscalingsResult
    {
        public readonly ImmutableArray<GetClusterClusterAutoscalingsAutoProvisioningDefaultsResult> AutoProvisioningDefaults;
        public readonly bool Enabled;
        public readonly ImmutableArray<GetClusterClusterAutoscalingsResourceLimitsResult> ResourceLimits;

        [OutputConstructor]
        private GetClusterClusterAutoscalingsResult(
            ImmutableArray<GetClusterClusterAutoscalingsAutoProvisioningDefaultsResult> autoProvisioningDefaults,
            bool enabled,
            ImmutableArray<GetClusterClusterAutoscalingsResourceLimitsResult> resourceLimits)
        {
            AutoProvisioningDefaults = autoProvisioningDefaults;
            Enabled = enabled;
            ResourceLimits = resourceLimits;
        }
    }

    [OutputType]
    public sealed class GetClusterDatabaseEncryptionsResult
    {
        public readonly string KeyName;
        public readonly string State;

        [OutputConstructor]
        private GetClusterDatabaseEncryptionsResult(
            string keyName,
            string state)
        {
            KeyName = keyName;
            State = state;
        }
    }

    [OutputType]
    public sealed class GetClusterIpAllocationPoliciesResult
    {
        public readonly string ClusterIpv4CidrBlock;
        public readonly string ClusterSecondaryRangeName;
        public readonly bool CreateSubnetwork;
        public readonly string NodeIpv4CidrBlock;
        public readonly string ServicesIpv4CidrBlock;
        public readonly string ServicesSecondaryRangeName;
        public readonly string SubnetworkName;
        public readonly bool UseIpAliases;

        [OutputConstructor]
        private GetClusterIpAllocationPoliciesResult(
            string clusterIpv4CidrBlock,
            string clusterSecondaryRangeName,
            bool createSubnetwork,
            string nodeIpv4CidrBlock,
            string servicesIpv4CidrBlock,
            string servicesSecondaryRangeName,
            string subnetworkName,
            bool useIpAliases)
        {
            ClusterIpv4CidrBlock = clusterIpv4CidrBlock;
            ClusterSecondaryRangeName = clusterSecondaryRangeName;
            CreateSubnetwork = createSubnetwork;
            NodeIpv4CidrBlock = nodeIpv4CidrBlock;
            ServicesIpv4CidrBlock = servicesIpv4CidrBlock;
            ServicesSecondaryRangeName = servicesSecondaryRangeName;
            SubnetworkName = subnetworkName;
            UseIpAliases = useIpAliases;
        }
    }

    [OutputType]
    public sealed class GetClusterMaintenancePoliciesDailyMaintenanceWindowsResult
    {
        public readonly string Duration;
        public readonly string StartTime;

        [OutputConstructor]
        private GetClusterMaintenancePoliciesDailyMaintenanceWindowsResult(
            string duration,
            string startTime)
        {
            Duration = duration;
            StartTime = startTime;
        }
    }

    [OutputType]
    public sealed class GetClusterMaintenancePoliciesRecurringWindowsResult
    {
        public readonly string EndTime;
        public readonly string Recurrence;
        public readonly string StartTime;

        [OutputConstructor]
        private GetClusterMaintenancePoliciesRecurringWindowsResult(
            string endTime,
            string recurrence,
            string startTime)
        {
            EndTime = endTime;
            Recurrence = recurrence;
            StartTime = startTime;
        }
    }

    [OutputType]
    public sealed class GetClusterMaintenancePoliciesResult
    {
        public readonly ImmutableArray<GetClusterMaintenancePoliciesDailyMaintenanceWindowsResult> DailyMaintenanceWindows;
        public readonly ImmutableArray<GetClusterMaintenancePoliciesRecurringWindowsResult> RecurringWindows;

        [OutputConstructor]
        private GetClusterMaintenancePoliciesResult(
            ImmutableArray<GetClusterMaintenancePoliciesDailyMaintenanceWindowsResult> dailyMaintenanceWindows,
            ImmutableArray<GetClusterMaintenancePoliciesRecurringWindowsResult> recurringWindows)
        {
            DailyMaintenanceWindows = dailyMaintenanceWindows;
            RecurringWindows = recurringWindows;
        }
    }

    [OutputType]
    public sealed class GetClusterMasterAuthorizedNetworksConfigsCidrBlocksResult
    {
        public readonly string CidrBlock;
        public readonly string DisplayName;

        [OutputConstructor]
        private GetClusterMasterAuthorizedNetworksConfigsCidrBlocksResult(
            string cidrBlock,
            string displayName)
        {
            CidrBlock = cidrBlock;
            DisplayName = displayName;
        }
    }

    [OutputType]
    public sealed class GetClusterMasterAuthorizedNetworksConfigsResult
    {
        public readonly ImmutableArray<GetClusterMasterAuthorizedNetworksConfigsCidrBlocksResult> CidrBlocks;

        [OutputConstructor]
        private GetClusterMasterAuthorizedNetworksConfigsResult(ImmutableArray<GetClusterMasterAuthorizedNetworksConfigsCidrBlocksResult> cidrBlocks)
        {
            CidrBlocks = cidrBlocks;
        }
    }

    [OutputType]
    public sealed class GetClusterMasterAuthsClientCertificateConfigsResult
    {
        public readonly bool IssueClientCertificate;

        [OutputConstructor]
        private GetClusterMasterAuthsClientCertificateConfigsResult(bool issueClientCertificate)
        {
            IssueClientCertificate = issueClientCertificate;
        }
    }

    [OutputType]
    public sealed class GetClusterMasterAuthsResult
    {
        public readonly string ClientCertificate;
        public readonly ImmutableArray<GetClusterMasterAuthsClientCertificateConfigsResult> ClientCertificateConfigs;
        public readonly string ClientKey;
        public readonly string ClusterCaCertificate;
        public readonly string Password;
        public readonly string Username;

        [OutputConstructor]
        private GetClusterMasterAuthsResult(
            string clientCertificate,
            ImmutableArray<GetClusterMasterAuthsClientCertificateConfigsResult> clientCertificateConfigs,
            string clientKey,
            string clusterCaCertificate,
            string password,
            string username)
        {
            ClientCertificate = clientCertificate;
            ClientCertificateConfigs = clientCertificateConfigs;
            ClientKey = clientKey;
            ClusterCaCertificate = clusterCaCertificate;
            Password = password;
            Username = username;
        }
    }

    [OutputType]
    public sealed class GetClusterNetworkPoliciesResult
    {
        public readonly bool Enabled;
        public readonly string Provider;

        [OutputConstructor]
        private GetClusterNetworkPoliciesResult(
            bool enabled,
            string provider)
        {
            Enabled = enabled;
            Provider = provider;
        }
    }

    [OutputType]
    public sealed class GetClusterNodeConfigsGuestAcceleratorsResult
    {
        public readonly int Count;
        public readonly string Type;

        [OutputConstructor]
        private GetClusterNodeConfigsGuestAcceleratorsResult(
            int count,
            string type)
        {
            Count = count;
            Type = type;
        }
    }

    [OutputType]
    public sealed class GetClusterNodeConfigsResult
    {
        public readonly int DiskSizeGb;
        public readonly string DiskType;
        public readonly ImmutableArray<GetClusterNodeConfigsGuestAcceleratorsResult> GuestAccelerators;
        public readonly string ImageType;
        public readonly ImmutableDictionary<string, string> Labels;
        public readonly int LocalSsdCount;
        public readonly string MachineType;
        public readonly ImmutableDictionary<string, string> Metadata;
        public readonly string MinCpuPlatform;
        public readonly ImmutableArray<string> OauthScopes;
        public readonly bool Preemptible;
        public readonly ImmutableArray<GetClusterNodeConfigsSandboxConfigsResult> SandboxConfigs;
        public readonly string ServiceAccount;
        public readonly ImmutableArray<GetClusterNodeConfigsShieldedInstanceConfigsResult> ShieldedInstanceConfigs;
        public readonly ImmutableArray<string> Tags;
        public readonly ImmutableArray<GetClusterNodeConfigsTaintsResult> Taints;
        public readonly ImmutableArray<GetClusterNodeConfigsWorkloadMetadataConfigsResult> WorkloadMetadataConfigs;

        [OutputConstructor]
        private GetClusterNodeConfigsResult(
            int diskSizeGb,
            string diskType,
            ImmutableArray<GetClusterNodeConfigsGuestAcceleratorsResult> guestAccelerators,
            string imageType,
            ImmutableDictionary<string, string> labels,
            int localSsdCount,
            string machineType,
            ImmutableDictionary<string, string> metadata,
            string minCpuPlatform,
            ImmutableArray<string> oauthScopes,
            bool preemptible,
            ImmutableArray<GetClusterNodeConfigsSandboxConfigsResult> sandboxConfigs,
            string serviceAccount,
            ImmutableArray<GetClusterNodeConfigsShieldedInstanceConfigsResult> shieldedInstanceConfigs,
            ImmutableArray<string> tags,
            ImmutableArray<GetClusterNodeConfigsTaintsResult> taints,
            ImmutableArray<GetClusterNodeConfigsWorkloadMetadataConfigsResult> workloadMetadataConfigs)
        {
            DiskSizeGb = diskSizeGb;
            DiskType = diskType;
            GuestAccelerators = guestAccelerators;
            ImageType = imageType;
            Labels = labels;
            LocalSsdCount = localSsdCount;
            MachineType = machineType;
            Metadata = metadata;
            MinCpuPlatform = minCpuPlatform;
            OauthScopes = oauthScopes;
            Preemptible = preemptible;
            SandboxConfigs = sandboxConfigs;
            ServiceAccount = serviceAccount;
            ShieldedInstanceConfigs = shieldedInstanceConfigs;
            Tags = tags;
            Taints = taints;
            WorkloadMetadataConfigs = workloadMetadataConfigs;
        }
    }

    [OutputType]
    public sealed class GetClusterNodeConfigsSandboxConfigsResult
    {
        public readonly string SandboxType;

        [OutputConstructor]
        private GetClusterNodeConfigsSandboxConfigsResult(string sandboxType)
        {
            SandboxType = sandboxType;
        }
    }

    [OutputType]
    public sealed class GetClusterNodeConfigsShieldedInstanceConfigsResult
    {
        public readonly bool EnableIntegrityMonitoring;
        public readonly bool EnableSecureBoot;

        [OutputConstructor]
        private GetClusterNodeConfigsShieldedInstanceConfigsResult(
            bool enableIntegrityMonitoring,
            bool enableSecureBoot)
        {
            EnableIntegrityMonitoring = enableIntegrityMonitoring;
            EnableSecureBoot = enableSecureBoot;
        }
    }

    [OutputType]
    public sealed class GetClusterNodeConfigsTaintsResult
    {
        public readonly string Effect;
        public readonly string Key;
        public readonly string Value;

        [OutputConstructor]
        private GetClusterNodeConfigsTaintsResult(
            string effect,
            string key,
            string value)
        {
            Effect = effect;
            Key = key;
            Value = value;
        }
    }

    [OutputType]
    public sealed class GetClusterNodeConfigsWorkloadMetadataConfigsResult
    {
        public readonly string NodeMetadata;

        [OutputConstructor]
        private GetClusterNodeConfigsWorkloadMetadataConfigsResult(string nodeMetadata)
        {
            NodeMetadata = nodeMetadata;
        }
    }

    [OutputType]
    public sealed class GetClusterNodePoolsAutoscalingsResult
    {
        public readonly int MaxNodeCount;
        public readonly int MinNodeCount;

        [OutputConstructor]
        private GetClusterNodePoolsAutoscalingsResult(
            int maxNodeCount,
            int minNodeCount)
        {
            MaxNodeCount = maxNodeCount;
            MinNodeCount = minNodeCount;
        }
    }

    [OutputType]
    public sealed class GetClusterNodePoolsManagementsResult
    {
        public readonly bool AutoRepair;
        public readonly bool AutoUpgrade;

        [OutputConstructor]
        private GetClusterNodePoolsManagementsResult(
            bool autoRepair,
            bool autoUpgrade)
        {
            AutoRepair = autoRepair;
            AutoUpgrade = autoUpgrade;
        }
    }

    [OutputType]
    public sealed class GetClusterNodePoolsNodeConfigsGuestAcceleratorsResult
    {
        public readonly int Count;
        public readonly string Type;

        [OutputConstructor]
        private GetClusterNodePoolsNodeConfigsGuestAcceleratorsResult(
            int count,
            string type)
        {
            Count = count;
            Type = type;
        }
    }

    [OutputType]
    public sealed class GetClusterNodePoolsNodeConfigsResult
    {
        public readonly int DiskSizeGb;
        public readonly string DiskType;
        public readonly ImmutableArray<GetClusterNodePoolsNodeConfigsGuestAcceleratorsResult> GuestAccelerators;
        public readonly string ImageType;
        public readonly ImmutableDictionary<string, string> Labels;
        public readonly int LocalSsdCount;
        public readonly string MachineType;
        public readonly ImmutableDictionary<string, string> Metadata;
        public readonly string MinCpuPlatform;
        public readonly ImmutableArray<string> OauthScopes;
        public readonly bool Preemptible;
        public readonly ImmutableArray<GetClusterNodePoolsNodeConfigsSandboxConfigsResult> SandboxConfigs;
        public readonly string ServiceAccount;
        public readonly ImmutableArray<GetClusterNodePoolsNodeConfigsShieldedInstanceConfigsResult> ShieldedInstanceConfigs;
        public readonly ImmutableArray<string> Tags;
        public readonly ImmutableArray<GetClusterNodePoolsNodeConfigsTaintsResult> Taints;
        public readonly ImmutableArray<GetClusterNodePoolsNodeConfigsWorkloadMetadataConfigsResult> WorkloadMetadataConfigs;

        [OutputConstructor]
        private GetClusterNodePoolsNodeConfigsResult(
            int diskSizeGb,
            string diskType,
            ImmutableArray<GetClusterNodePoolsNodeConfigsGuestAcceleratorsResult> guestAccelerators,
            string imageType,
            ImmutableDictionary<string, string> labels,
            int localSsdCount,
            string machineType,
            ImmutableDictionary<string, string> metadata,
            string minCpuPlatform,
            ImmutableArray<string> oauthScopes,
            bool preemptible,
            ImmutableArray<GetClusterNodePoolsNodeConfigsSandboxConfigsResult> sandboxConfigs,
            string serviceAccount,
            ImmutableArray<GetClusterNodePoolsNodeConfigsShieldedInstanceConfigsResult> shieldedInstanceConfigs,
            ImmutableArray<string> tags,
            ImmutableArray<GetClusterNodePoolsNodeConfigsTaintsResult> taints,
            ImmutableArray<GetClusterNodePoolsNodeConfigsWorkloadMetadataConfigsResult> workloadMetadataConfigs)
        {
            DiskSizeGb = diskSizeGb;
            DiskType = diskType;
            GuestAccelerators = guestAccelerators;
            ImageType = imageType;
            Labels = labels;
            LocalSsdCount = localSsdCount;
            MachineType = machineType;
            Metadata = metadata;
            MinCpuPlatform = minCpuPlatform;
            OauthScopes = oauthScopes;
            Preemptible = preemptible;
            SandboxConfigs = sandboxConfigs;
            ServiceAccount = serviceAccount;
            ShieldedInstanceConfigs = shieldedInstanceConfigs;
            Tags = tags;
            Taints = taints;
            WorkloadMetadataConfigs = workloadMetadataConfigs;
        }
    }

    [OutputType]
    public sealed class GetClusterNodePoolsNodeConfigsSandboxConfigsResult
    {
        public readonly string SandboxType;

        [OutputConstructor]
        private GetClusterNodePoolsNodeConfigsSandboxConfigsResult(string sandboxType)
        {
            SandboxType = sandboxType;
        }
    }

    [OutputType]
    public sealed class GetClusterNodePoolsNodeConfigsShieldedInstanceConfigsResult
    {
        public readonly bool EnableIntegrityMonitoring;
        public readonly bool EnableSecureBoot;

        [OutputConstructor]
        private GetClusterNodePoolsNodeConfigsShieldedInstanceConfigsResult(
            bool enableIntegrityMonitoring,
            bool enableSecureBoot)
        {
            EnableIntegrityMonitoring = enableIntegrityMonitoring;
            EnableSecureBoot = enableSecureBoot;
        }
    }

    [OutputType]
    public sealed class GetClusterNodePoolsNodeConfigsTaintsResult
    {
        public readonly string Effect;
        public readonly string Key;
        public readonly string Value;

        [OutputConstructor]
        private GetClusterNodePoolsNodeConfigsTaintsResult(
            string effect,
            string key,
            string value)
        {
            Effect = effect;
            Key = key;
            Value = value;
        }
    }

    [OutputType]
    public sealed class GetClusterNodePoolsNodeConfigsWorkloadMetadataConfigsResult
    {
        public readonly string NodeMetadata;

        [OutputConstructor]
        private GetClusterNodePoolsNodeConfigsWorkloadMetadataConfigsResult(string nodeMetadata)
        {
            NodeMetadata = nodeMetadata;
        }
    }

    [OutputType]
    public sealed class GetClusterNodePoolsResult
    {
        public readonly ImmutableArray<GetClusterNodePoolsAutoscalingsResult> Autoscalings;
        public readonly int InitialNodeCount;
        public readonly ImmutableArray<string> InstanceGroupUrls;
        public readonly ImmutableArray<GetClusterNodePoolsManagementsResult> Managements;
        public readonly int MaxPodsPerNode;
        /// <summary>
        /// The name of the cluster.
        /// </summary>
        public readonly string Name;
        public readonly string NamePrefix;
        public readonly ImmutableArray<GetClusterNodePoolsNodeConfigsResult> NodeConfigs;
        public readonly int NodeCount;
        public readonly ImmutableArray<string> NodeLocations;
        public readonly ImmutableArray<GetClusterNodePoolsUpgradeSettingsResult> UpgradeSettings;
        public readonly string Version;

        [OutputConstructor]
        private GetClusterNodePoolsResult(
            ImmutableArray<GetClusterNodePoolsAutoscalingsResult> autoscalings,
            int initialNodeCount,
            ImmutableArray<string> instanceGroupUrls,
            ImmutableArray<GetClusterNodePoolsManagementsResult> managements,
            int maxPodsPerNode,
            string name,
            string namePrefix,
            ImmutableArray<GetClusterNodePoolsNodeConfigsResult> nodeConfigs,
            int nodeCount,
            ImmutableArray<string> nodeLocations,
            ImmutableArray<GetClusterNodePoolsUpgradeSettingsResult> upgradeSettings,
            string version)
        {
            Autoscalings = autoscalings;
            InitialNodeCount = initialNodeCount;
            InstanceGroupUrls = instanceGroupUrls;
            Managements = managements;
            MaxPodsPerNode = maxPodsPerNode;
            Name = name;
            NamePrefix = namePrefix;
            NodeConfigs = nodeConfigs;
            NodeCount = nodeCount;
            NodeLocations = nodeLocations;
            UpgradeSettings = upgradeSettings;
            Version = version;
        }
    }

    [OutputType]
    public sealed class GetClusterNodePoolsUpgradeSettingsResult
    {
        public readonly int MaxSurge;
        public readonly int MaxUnavailable;

        [OutputConstructor]
        private GetClusterNodePoolsUpgradeSettingsResult(
            int maxSurge,
            int maxUnavailable)
        {
            MaxSurge = maxSurge;
            MaxUnavailable = maxUnavailable;
        }
    }

    [OutputType]
    public sealed class GetClusterPodSecurityPolicyConfigsResult
    {
        public readonly bool Enabled;

        [OutputConstructor]
        private GetClusterPodSecurityPolicyConfigsResult(bool enabled)
        {
            Enabled = enabled;
        }
    }

    [OutputType]
    public sealed class GetClusterPrivateClusterConfigsResult
    {
        public readonly bool EnablePrivateEndpoint;
        public readonly bool EnablePrivateNodes;
        public readonly string MasterIpv4CidrBlock;
        public readonly string PeeringName;
        public readonly string PrivateEndpoint;
        public readonly string PublicEndpoint;

        [OutputConstructor]
        private GetClusterPrivateClusterConfigsResult(
            bool enablePrivateEndpoint,
            bool enablePrivateNodes,
            string masterIpv4CidrBlock,
            string peeringName,
            string privateEndpoint,
            string publicEndpoint)
        {
            EnablePrivateEndpoint = enablePrivateEndpoint;
            EnablePrivateNodes = enablePrivateNodes;
            MasterIpv4CidrBlock = masterIpv4CidrBlock;
            PeeringName = peeringName;
            PrivateEndpoint = privateEndpoint;
            PublicEndpoint = publicEndpoint;
        }
    }

    [OutputType]
    public sealed class GetClusterReleaseChannelsResult
    {
        public readonly string Channel;

        [OutputConstructor]
        private GetClusterReleaseChannelsResult(string channel)
        {
            Channel = channel;
        }
    }

    [OutputType]
    public sealed class GetClusterResourceUsageExportConfigsBigqueryDestinationsResult
    {
        public readonly string DatasetId;

        [OutputConstructor]
        private GetClusterResourceUsageExportConfigsBigqueryDestinationsResult(string datasetId)
        {
            DatasetId = datasetId;
        }
    }

    [OutputType]
    public sealed class GetClusterResourceUsageExportConfigsResult
    {
        public readonly ImmutableArray<GetClusterResourceUsageExportConfigsBigqueryDestinationsResult> BigqueryDestinations;
        public readonly bool EnableNetworkEgressMetering;

        [OutputConstructor]
        private GetClusterResourceUsageExportConfigsResult(
            ImmutableArray<GetClusterResourceUsageExportConfigsBigqueryDestinationsResult> bigqueryDestinations,
            bool enableNetworkEgressMetering)
        {
            BigqueryDestinations = bigqueryDestinations;
            EnableNetworkEgressMetering = enableNetworkEgressMetering;
        }
    }

    [OutputType]
    public sealed class GetClusterVerticalPodAutoscalingsResult
    {
        public readonly bool Enabled;

        [OutputConstructor]
        private GetClusterVerticalPodAutoscalingsResult(bool enabled)
        {
            Enabled = enabled;
        }
    }

    [OutputType]
    public sealed class GetClusterWorkloadIdentityConfigsResult
    {
        public readonly string IdentityNamespace;

        [OutputConstructor]
        private GetClusterWorkloadIdentityConfigsResult(string identityNamespace)
        {
            IdentityNamespace = identityNamespace;
        }
    }
    }
}
