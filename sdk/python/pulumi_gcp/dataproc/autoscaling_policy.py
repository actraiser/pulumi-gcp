# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class AutoscalingPolicy(pulumi.CustomResource):
    basic_algorithm: pulumi.Output[dict]
    """
    Basic algorithm for autoscaling.  Structure is documented below.

      * `cooldownPeriod` (`str`) - Duration between scaling events. A scaling period starts after the
        update operation from the previous event has completed.
        Bounds: [2m, 1d]. Default: 2m.
      * `yarnConfig` (`dict`) - YARN autoscaling configuration.  Structure is documented below.
        * `gracefulDecommissionTimeout` (`str`) - Timeout for YARN graceful decommissioning of Node Managers. Specifies the
          duration to wait for jobs to complete before forcefully removing workers
          (and potentially interrupting jobs). Only applicable to downscaling operations.
          Bounds: [0s, 1d].
        * `scaleDownFactor` (`float`) - Fraction of average pending memory in the last cooldown period for which to
          remove workers. A scale-down factor of 1 will result in scaling down so that there
          is no available memory remaining after the update (more aggressive scaling).
          A scale-down factor of 0 disables removing workers, which can be beneficial for
          autoscaling a single job.
          Bounds: [0.0, 1.0].
        * `scaleDownMinWorkerFraction` (`float`) - Minimum scale-down threshold as a fraction of total cluster size before scaling occurs.
          For example, in a 20-worker cluster, a threshold of 0.1 means the autoscaler must
          recommend at least a 2 worker scale-down for the cluster to scale. A threshold of 0
          means the autoscaler will scale down on any recommended change.
          Bounds: [0.0, 1.0]. Default: 0.0.
        * `scaleUpFactor` (`float`) - Fraction of average pending memory in the last cooldown period for which to
          add workers. A scale-up factor of 1.0 will result in scaling up so that there
          is no pending memory remaining after the update (more aggressive scaling).
          A scale-up factor closer to 0 will result in a smaller magnitude of scaling up
          (less aggressive scaling).
          Bounds: [0.0, 1.0].
        * `scaleUpMinWorkerFraction` (`float`) - Minimum scale-up threshold as a fraction of total cluster size before scaling
          occurs. For example, in a 20-worker cluster, a threshold of 0.1 means the autoscaler
          must recommend at least a 2-worker scale-up for the cluster to scale. A threshold of
          0 means the autoscaler will scale up on any recommended change.
          Bounds: [0.0, 1.0]. Default: 0.0.
    """
    location: pulumi.Output[str]
    """
    The  location where the autoscaling poicy should reside.
    The default value is `global`.
    """
    name: pulumi.Output[str]
    """
    The "resource name" of the autoscaling policy.
    """
    policy_id: pulumi.Output[str]
    """
    The policy id. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_),
    and hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between
    3 and 50 characters.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    secondary_worker_config: pulumi.Output[dict]
    """
    Describes how the autoscaler will operate for secondary workers.  Structure is documented below.

      * `max_instances` (`float`) - Maximum number of instances for this group. Note that by default, clusters will not use
        secondary workers. Required for secondary workers if the minimum secondary instances is set.
        Bounds: [minInstances, ). Defaults to 0.
      * `minInstances` (`float`) - Minimum number of instances for this group. Bounds: [0, maxInstances]. Defaults to 0.
      * `weight` (`float`) - Weight for the instance group, which is used to determine the fraction of total workers
        in the cluster from this instance group. For example, if primary workers have weight 2,
        and secondary workers have weight 1, the cluster will have approximately 2 primary workers
        for each secondary worker.
        The cluster may not reach the specified balance if constrained by min/max bounds or other
        autoscaling settings. For example, if maxInstances for secondary workers is 0, then only
        primary workers will be added. The cluster can also be out of balance when created.
        If weight is not set on any instance group, the cluster will default to equal weight for
        all groups: the cluster will attempt to maintain an equal number of workers in each group
        within the configured size bounds for each group. If weight is set for one group only,
        the cluster will default to zero weight on the unset group. For example if weight is set
        only on primary workers, the cluster will use primary workers only and no secondary workers.
    """
    worker_config: pulumi.Output[dict]
    """
    Describes how the autoscaler will operate for primary workers.  Structure is documented below.

      * `max_instances` (`float`) - Maximum number of instances for this group. Note that by default, clusters will not use
        secondary workers. Required for secondary workers if the minimum secondary instances is set.
        Bounds: [minInstances, ). Defaults to 0.
      * `minInstances` (`float`) - Minimum number of instances for this group. Bounds: [0, maxInstances]. Defaults to 0.
      * `weight` (`float`) - Weight for the instance group, which is used to determine the fraction of total workers
        in the cluster from this instance group. For example, if primary workers have weight 2,
        and secondary workers have weight 1, the cluster will have approximately 2 primary workers
        for each secondary worker.
        The cluster may not reach the specified balance if constrained by min/max bounds or other
        autoscaling settings. For example, if maxInstances for secondary workers is 0, then only
        primary workers will be added. The cluster can also be out of balance when created.
        If weight is not set on any instance group, the cluster will default to equal weight for
        all groups: the cluster will attempt to maintain an equal number of workers in each group
        within the configured size bounds for each group. If weight is set for one group only,
        the cluster will default to zero weight on the unset group. For example if weight is set
        only on primary workers, the cluster will use primary workers only and no secondary workers.
    """
    def __init__(__self__, resource_name, opts=None, basic_algorithm=None, location=None, policy_id=None, project=None, secondary_worker_config=None, worker_config=None, __props__=None, __name__=None, __opts__=None):
        """
        Describes an autoscaling policy for Dataproc cluster autoscaler.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] basic_algorithm: Basic algorithm for autoscaling.  Structure is documented below.
        :param pulumi.Input[str] location: The  location where the autoscaling poicy should reside.
               The default value is `global`.
        :param pulumi.Input[str] policy_id: The policy id. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_),
               and hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between
               3 and 50 characters.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[dict] secondary_worker_config: Describes how the autoscaler will operate for secondary workers.  Structure is documented below.
        :param pulumi.Input[dict] worker_config: Describes how the autoscaler will operate for primary workers.  Structure is documented below.

        The **basic_algorithm** object supports the following:

          * `cooldownPeriod` (`pulumi.Input[str]`) - Duration between scaling events. A scaling period starts after the
            update operation from the previous event has completed.
            Bounds: [2m, 1d]. Default: 2m.
          * `yarnConfig` (`pulumi.Input[dict]`) - YARN autoscaling configuration.  Structure is documented below.
            * `gracefulDecommissionTimeout` (`pulumi.Input[str]`) - Timeout for YARN graceful decommissioning of Node Managers. Specifies the
              duration to wait for jobs to complete before forcefully removing workers
              (and potentially interrupting jobs). Only applicable to downscaling operations.
              Bounds: [0s, 1d].
            * `scaleDownFactor` (`pulumi.Input[float]`) - Fraction of average pending memory in the last cooldown period for which to
              remove workers. A scale-down factor of 1 will result in scaling down so that there
              is no available memory remaining after the update (more aggressive scaling).
              A scale-down factor of 0 disables removing workers, which can be beneficial for
              autoscaling a single job.
              Bounds: [0.0, 1.0].
            * `scaleDownMinWorkerFraction` (`pulumi.Input[float]`) - Minimum scale-down threshold as a fraction of total cluster size before scaling occurs.
              For example, in a 20-worker cluster, a threshold of 0.1 means the autoscaler must
              recommend at least a 2 worker scale-down for the cluster to scale. A threshold of 0
              means the autoscaler will scale down on any recommended change.
              Bounds: [0.0, 1.0]. Default: 0.0.
            * `scaleUpFactor` (`pulumi.Input[float]`) - Fraction of average pending memory in the last cooldown period for which to
              add workers. A scale-up factor of 1.0 will result in scaling up so that there
              is no pending memory remaining after the update (more aggressive scaling).
              A scale-up factor closer to 0 will result in a smaller magnitude of scaling up
              (less aggressive scaling).
              Bounds: [0.0, 1.0].
            * `scaleUpMinWorkerFraction` (`pulumi.Input[float]`) - Minimum scale-up threshold as a fraction of total cluster size before scaling
              occurs. For example, in a 20-worker cluster, a threshold of 0.1 means the autoscaler
              must recommend at least a 2-worker scale-up for the cluster to scale. A threshold of
              0 means the autoscaler will scale up on any recommended change.
              Bounds: [0.0, 1.0]. Default: 0.0.

        The **secondary_worker_config** object supports the following:

          * `max_instances` (`pulumi.Input[float]`) - Maximum number of instances for this group. Note that by default, clusters will not use
            secondary workers. Required for secondary workers if the minimum secondary instances is set.
            Bounds: [minInstances, ). Defaults to 0.
          * `minInstances` (`pulumi.Input[float]`) - Minimum number of instances for this group. Bounds: [0, maxInstances]. Defaults to 0.
          * `weight` (`pulumi.Input[float]`) - Weight for the instance group, which is used to determine the fraction of total workers
            in the cluster from this instance group. For example, if primary workers have weight 2,
            and secondary workers have weight 1, the cluster will have approximately 2 primary workers
            for each secondary worker.
            The cluster may not reach the specified balance if constrained by min/max bounds or other
            autoscaling settings. For example, if maxInstances for secondary workers is 0, then only
            primary workers will be added. The cluster can also be out of balance when created.
            If weight is not set on any instance group, the cluster will default to equal weight for
            all groups: the cluster will attempt to maintain an equal number of workers in each group
            within the configured size bounds for each group. If weight is set for one group only,
            the cluster will default to zero weight on the unset group. For example if weight is set
            only on primary workers, the cluster will use primary workers only and no secondary workers.

        The **worker_config** object supports the following:

          * `max_instances` (`pulumi.Input[float]`) - Maximum number of instances for this group. Note that by default, clusters will not use
            secondary workers. Required for secondary workers if the minimum secondary instances is set.
            Bounds: [minInstances, ). Defaults to 0.
          * `minInstances` (`pulumi.Input[float]`) - Minimum number of instances for this group. Bounds: [0, maxInstances]. Defaults to 0.
          * `weight` (`pulumi.Input[float]`) - Weight for the instance group, which is used to determine the fraction of total workers
            in the cluster from this instance group. For example, if primary workers have weight 2,
            and secondary workers have weight 1, the cluster will have approximately 2 primary workers
            for each secondary worker.
            The cluster may not reach the specified balance if constrained by min/max bounds or other
            autoscaling settings. For example, if maxInstances for secondary workers is 0, then only
            primary workers will be added. The cluster can also be out of balance when created.
            If weight is not set on any instance group, the cluster will default to equal weight for
            all groups: the cluster will attempt to maintain an equal number of workers in each group
            within the configured size bounds for each group. If weight is set for one group only,
            the cluster will default to zero weight on the unset group. For example if weight is set
            only on primary workers, the cluster will use primary workers only and no secondary workers.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['basic_algorithm'] = basic_algorithm
            __props__['location'] = location
            if policy_id is None:
                raise TypeError("Missing required property 'policy_id'")
            __props__['policy_id'] = policy_id
            __props__['project'] = project
            __props__['secondary_worker_config'] = secondary_worker_config
            __props__['worker_config'] = worker_config
            __props__['name'] = None
        super(AutoscalingPolicy, __self__).__init__(
            'gcp:dataproc/autoscalingPolicy:AutoscalingPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, basic_algorithm=None, location=None, name=None, policy_id=None, project=None, secondary_worker_config=None, worker_config=None):
        """
        Get an existing AutoscalingPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] basic_algorithm: Basic algorithm for autoscaling.  Structure is documented below.
        :param pulumi.Input[str] location: The  location where the autoscaling poicy should reside.
               The default value is `global`.
        :param pulumi.Input[str] name: The "resource name" of the autoscaling policy.
        :param pulumi.Input[str] policy_id: The policy id. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_),
               and hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between
               3 and 50 characters.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[dict] secondary_worker_config: Describes how the autoscaler will operate for secondary workers.  Structure is documented below.
        :param pulumi.Input[dict] worker_config: Describes how the autoscaler will operate for primary workers.  Structure is documented below.

        The **basic_algorithm** object supports the following:

          * `cooldownPeriod` (`pulumi.Input[str]`) - Duration between scaling events. A scaling period starts after the
            update operation from the previous event has completed.
            Bounds: [2m, 1d]. Default: 2m.
          * `yarnConfig` (`pulumi.Input[dict]`) - YARN autoscaling configuration.  Structure is documented below.
            * `gracefulDecommissionTimeout` (`pulumi.Input[str]`) - Timeout for YARN graceful decommissioning of Node Managers. Specifies the
              duration to wait for jobs to complete before forcefully removing workers
              (and potentially interrupting jobs). Only applicable to downscaling operations.
              Bounds: [0s, 1d].
            * `scaleDownFactor` (`pulumi.Input[float]`) - Fraction of average pending memory in the last cooldown period for which to
              remove workers. A scale-down factor of 1 will result in scaling down so that there
              is no available memory remaining after the update (more aggressive scaling).
              A scale-down factor of 0 disables removing workers, which can be beneficial for
              autoscaling a single job.
              Bounds: [0.0, 1.0].
            * `scaleDownMinWorkerFraction` (`pulumi.Input[float]`) - Minimum scale-down threshold as a fraction of total cluster size before scaling occurs.
              For example, in a 20-worker cluster, a threshold of 0.1 means the autoscaler must
              recommend at least a 2 worker scale-down for the cluster to scale. A threshold of 0
              means the autoscaler will scale down on any recommended change.
              Bounds: [0.0, 1.0]. Default: 0.0.
            * `scaleUpFactor` (`pulumi.Input[float]`) - Fraction of average pending memory in the last cooldown period for which to
              add workers. A scale-up factor of 1.0 will result in scaling up so that there
              is no pending memory remaining after the update (more aggressive scaling).
              A scale-up factor closer to 0 will result in a smaller magnitude of scaling up
              (less aggressive scaling).
              Bounds: [0.0, 1.0].
            * `scaleUpMinWorkerFraction` (`pulumi.Input[float]`) - Minimum scale-up threshold as a fraction of total cluster size before scaling
              occurs. For example, in a 20-worker cluster, a threshold of 0.1 means the autoscaler
              must recommend at least a 2-worker scale-up for the cluster to scale. A threshold of
              0 means the autoscaler will scale up on any recommended change.
              Bounds: [0.0, 1.0]. Default: 0.0.

        The **secondary_worker_config** object supports the following:

          * `max_instances` (`pulumi.Input[float]`) - Maximum number of instances for this group. Note that by default, clusters will not use
            secondary workers. Required for secondary workers if the minimum secondary instances is set.
            Bounds: [minInstances, ). Defaults to 0.
          * `minInstances` (`pulumi.Input[float]`) - Minimum number of instances for this group. Bounds: [0, maxInstances]. Defaults to 0.
          * `weight` (`pulumi.Input[float]`) - Weight for the instance group, which is used to determine the fraction of total workers
            in the cluster from this instance group. For example, if primary workers have weight 2,
            and secondary workers have weight 1, the cluster will have approximately 2 primary workers
            for each secondary worker.
            The cluster may not reach the specified balance if constrained by min/max bounds or other
            autoscaling settings. For example, if maxInstances for secondary workers is 0, then only
            primary workers will be added. The cluster can also be out of balance when created.
            If weight is not set on any instance group, the cluster will default to equal weight for
            all groups: the cluster will attempt to maintain an equal number of workers in each group
            within the configured size bounds for each group. If weight is set for one group only,
            the cluster will default to zero weight on the unset group. For example if weight is set
            only on primary workers, the cluster will use primary workers only and no secondary workers.

        The **worker_config** object supports the following:

          * `max_instances` (`pulumi.Input[float]`) - Maximum number of instances for this group. Note that by default, clusters will not use
            secondary workers. Required for secondary workers if the minimum secondary instances is set.
            Bounds: [minInstances, ). Defaults to 0.
          * `minInstances` (`pulumi.Input[float]`) - Minimum number of instances for this group. Bounds: [0, maxInstances]. Defaults to 0.
          * `weight` (`pulumi.Input[float]`) - Weight for the instance group, which is used to determine the fraction of total workers
            in the cluster from this instance group. For example, if primary workers have weight 2,
            and secondary workers have weight 1, the cluster will have approximately 2 primary workers
            for each secondary worker.
            The cluster may not reach the specified balance if constrained by min/max bounds or other
            autoscaling settings. For example, if maxInstances for secondary workers is 0, then only
            primary workers will be added. The cluster can also be out of balance when created.
            If weight is not set on any instance group, the cluster will default to equal weight for
            all groups: the cluster will attempt to maintain an equal number of workers in each group
            within the configured size bounds for each group. If weight is set for one group only,
            the cluster will default to zero weight on the unset group. For example if weight is set
            only on primary workers, the cluster will use primary workers only and no secondary workers.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["basic_algorithm"] = basic_algorithm
        __props__["location"] = location
        __props__["name"] = name
        __props__["policy_id"] = policy_id
        __props__["project"] = project
        __props__["secondary_worker_config"] = secondary_worker_config
        __props__["worker_config"] = worker_config
        return AutoscalingPolicy(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

