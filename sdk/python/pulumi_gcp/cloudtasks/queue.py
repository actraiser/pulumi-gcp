# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Queue(pulumi.CustomResource):
    app_engine_routing_override: pulumi.Output[dict]
    """
    Overrides for task-level appEngineRouting. These settings apply only
    to App Engine tasks in this queue  Structure is documented below.

      * `host` (`str`) - -
        The host that the task is sent to.
      * `instance` (`str`) - App instance.
        By default, the task is sent to an instance which is available when the task is attempted.
      * `service` (`str`) - App service.
        By default, the task is sent to the service which is the default service when the task is attempted.
      * `version` (`str`) - App version.
        By default, the task is sent to the version which is the default version when the task is attempted.
    """
    location: pulumi.Output[str]
    """
    The location of the queue
    """
    name: pulumi.Output[str]
    """
    The queue name.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    rate_limits: pulumi.Output[dict]
    """
    Rate limits for task dispatches.
    The queue's actual dispatch rate is the result of:
    * Number of tasks in the queue
    * User-specified throttling: rateLimits, retryConfig, and the queue's state.
    * System throttling due to 429 (Too Many Requests) or 503 (Service
    Unavailable) responses from the worker, high error rates, or to
    smooth sudden large traffic spikes.  Structure is documented below.

      * `maxBurstSize` (`float`) - -
        The max burst size.
        Max burst size limits how fast tasks in queue are processed when many tasks are
        in the queue and the rate is high. This field allows the queue to have a high
        rate so processing starts shortly after a task is enqueued, but still limits
        resource usage when many tasks are enqueued in a short period of time.
      * `maxConcurrentDispatches` (`float`) - The maximum number of concurrent tasks that Cloud Tasks allows to
        be dispatched for this queue. After this threshold has been
        reached, Cloud Tasks stops dispatching tasks until the number of
        concurrent requests decreases.
      * `maxDispatchesPerSecond` (`float`) - The maximum rate at which tasks are dispatched from this queue.
        If unspecified when the queue is created, Cloud Tasks will pick the default.
    """
    retry_config: pulumi.Output[dict]
    """
    Settings that determine the retry behavior.  Structure is documented below.

      * `maxAttempts` (`float`) - Number of attempts per task.
        Cloud Tasks will attempt the task maxAttempts times (that is, if
        the first attempt fails, then there will be maxAttempts - 1
        retries). Must be >= -1.
        If unspecified when the queue is created, Cloud Tasks will pick
        the default.
        -1 indicates unlimited attempts.
      * `maxBackoff` (`str`) - A task will be scheduled for retry between minBackoff and
        maxBackoff duration after it fails, if the queue's RetryConfig
        specifies that the task should be retried.
      * `maxDoublings` (`float`) - The time between retries will double maxDoublings times.
        A task's retry interval starts at minBackoff, then doubles maxDoublings times,
        then increases linearly, and finally retries retries at intervals of maxBackoff
        up to maxAttempts times.
      * `maxRetryDuration` (`str`) - If positive, maxRetryDuration specifies the time limit for
        retrying a failed task, measured from when the task was first
        attempted. Once maxRetryDuration time has passed and the task has
        been attempted maxAttempts times, no further attempts will be
        made and the task will be deleted.
        If zero, then the task age is unlimited.
      * `minBackoff` (`str`) - A task will be scheduled for retry between minBackoff and
        maxBackoff duration after it fails, if the queue's RetryConfig
        specifies that the task should be retried.
    """
    def __init__(__self__, resource_name, opts=None, app_engine_routing_override=None, location=None, name=None, project=None, rate_limits=None, retry_config=None, __props__=None, __name__=None, __opts__=None):
        """
        A named resource to which messages are sent by publishers.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] app_engine_routing_override: Overrides for task-level appEngineRouting. These settings apply only
               to App Engine tasks in this queue  Structure is documented below.
        :param pulumi.Input[str] location: The location of the queue
        :param pulumi.Input[str] name: The queue name.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[dict] rate_limits: Rate limits for task dispatches.
               The queue's actual dispatch rate is the result of:
               * Number of tasks in the queue
               * User-specified throttling: rateLimits, retryConfig, and the queue's state.
               * System throttling due to 429 (Too Many Requests) or 503 (Service
               Unavailable) responses from the worker, high error rates, or to
               smooth sudden large traffic spikes.  Structure is documented below.
        :param pulumi.Input[dict] retry_config: Settings that determine the retry behavior.  Structure is documented below.

        The **app_engine_routing_override** object supports the following:

          * `host` (`pulumi.Input[str]`) - -
            The host that the task is sent to.
          * `instance` (`pulumi.Input[str]`) - App instance.
            By default, the task is sent to an instance which is available when the task is attempted.
          * `service` (`pulumi.Input[str]`) - App service.
            By default, the task is sent to the service which is the default service when the task is attempted.
          * `version` (`pulumi.Input[str]`) - App version.
            By default, the task is sent to the version which is the default version when the task is attempted.

        The **rate_limits** object supports the following:

          * `maxBurstSize` (`pulumi.Input[float]`) - -
            The max burst size.
            Max burst size limits how fast tasks in queue are processed when many tasks are
            in the queue and the rate is high. This field allows the queue to have a high
            rate so processing starts shortly after a task is enqueued, but still limits
            resource usage when many tasks are enqueued in a short period of time.
          * `maxConcurrentDispatches` (`pulumi.Input[float]`) - The maximum number of concurrent tasks that Cloud Tasks allows to
            be dispatched for this queue. After this threshold has been
            reached, Cloud Tasks stops dispatching tasks until the number of
            concurrent requests decreases.
          * `maxDispatchesPerSecond` (`pulumi.Input[float]`) - The maximum rate at which tasks are dispatched from this queue.
            If unspecified when the queue is created, Cloud Tasks will pick the default.

        The **retry_config** object supports the following:

          * `maxAttempts` (`pulumi.Input[float]`) - Number of attempts per task.
            Cloud Tasks will attempt the task maxAttempts times (that is, if
            the first attempt fails, then there will be maxAttempts - 1
            retries). Must be >= -1.
            If unspecified when the queue is created, Cloud Tasks will pick
            the default.
            -1 indicates unlimited attempts.
          * `maxBackoff` (`pulumi.Input[str]`) - A task will be scheduled for retry between minBackoff and
            maxBackoff duration after it fails, if the queue's RetryConfig
            specifies that the task should be retried.
          * `maxDoublings` (`pulumi.Input[float]`) - The time between retries will double maxDoublings times.
            A task's retry interval starts at minBackoff, then doubles maxDoublings times,
            then increases linearly, and finally retries retries at intervals of maxBackoff
            up to maxAttempts times.
          * `maxRetryDuration` (`pulumi.Input[str]`) - If positive, maxRetryDuration specifies the time limit for
            retrying a failed task, measured from when the task was first
            attempted. Once maxRetryDuration time has passed and the task has
            been attempted maxAttempts times, no further attempts will be
            made and the task will be deleted.
            If zero, then the task age is unlimited.
          * `minBackoff` (`pulumi.Input[str]`) - A task will be scheduled for retry between minBackoff and
            maxBackoff duration after it fails, if the queue's RetryConfig
            specifies that the task should be retried.
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

            __props__['app_engine_routing_override'] = app_engine_routing_override
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            __props__['name'] = name
            __props__['project'] = project
            __props__['rate_limits'] = rate_limits
            __props__['retry_config'] = retry_config
        super(Queue, __self__).__init__(
            'gcp:cloudtasks/queue:Queue',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, app_engine_routing_override=None, location=None, name=None, project=None, rate_limits=None, retry_config=None):
        """
        Get an existing Queue resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] app_engine_routing_override: Overrides for task-level appEngineRouting. These settings apply only
               to App Engine tasks in this queue  Structure is documented below.
        :param pulumi.Input[str] location: The location of the queue
        :param pulumi.Input[str] name: The queue name.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[dict] rate_limits: Rate limits for task dispatches.
               The queue's actual dispatch rate is the result of:
               * Number of tasks in the queue
               * User-specified throttling: rateLimits, retryConfig, and the queue's state.
               * System throttling due to 429 (Too Many Requests) or 503 (Service
               Unavailable) responses from the worker, high error rates, or to
               smooth sudden large traffic spikes.  Structure is documented below.
        :param pulumi.Input[dict] retry_config: Settings that determine the retry behavior.  Structure is documented below.

        The **app_engine_routing_override** object supports the following:

          * `host` (`pulumi.Input[str]`) - -
            The host that the task is sent to.
          * `instance` (`pulumi.Input[str]`) - App instance.
            By default, the task is sent to an instance which is available when the task is attempted.
          * `service` (`pulumi.Input[str]`) - App service.
            By default, the task is sent to the service which is the default service when the task is attempted.
          * `version` (`pulumi.Input[str]`) - App version.
            By default, the task is sent to the version which is the default version when the task is attempted.

        The **rate_limits** object supports the following:

          * `maxBurstSize` (`pulumi.Input[float]`) - -
            The max burst size.
            Max burst size limits how fast tasks in queue are processed when many tasks are
            in the queue and the rate is high. This field allows the queue to have a high
            rate so processing starts shortly after a task is enqueued, but still limits
            resource usage when many tasks are enqueued in a short period of time.
          * `maxConcurrentDispatches` (`pulumi.Input[float]`) - The maximum number of concurrent tasks that Cloud Tasks allows to
            be dispatched for this queue. After this threshold has been
            reached, Cloud Tasks stops dispatching tasks until the number of
            concurrent requests decreases.
          * `maxDispatchesPerSecond` (`pulumi.Input[float]`) - The maximum rate at which tasks are dispatched from this queue.
            If unspecified when the queue is created, Cloud Tasks will pick the default.

        The **retry_config** object supports the following:

          * `maxAttempts` (`pulumi.Input[float]`) - Number of attempts per task.
            Cloud Tasks will attempt the task maxAttempts times (that is, if
            the first attempt fails, then there will be maxAttempts - 1
            retries). Must be >= -1.
            If unspecified when the queue is created, Cloud Tasks will pick
            the default.
            -1 indicates unlimited attempts.
          * `maxBackoff` (`pulumi.Input[str]`) - A task will be scheduled for retry between minBackoff and
            maxBackoff duration after it fails, if the queue's RetryConfig
            specifies that the task should be retried.
          * `maxDoublings` (`pulumi.Input[float]`) - The time between retries will double maxDoublings times.
            A task's retry interval starts at minBackoff, then doubles maxDoublings times,
            then increases linearly, and finally retries retries at intervals of maxBackoff
            up to maxAttempts times.
          * `maxRetryDuration` (`pulumi.Input[str]`) - If positive, maxRetryDuration specifies the time limit for
            retrying a failed task, measured from when the task was first
            attempted. Once maxRetryDuration time has passed and the task has
            been attempted maxAttempts times, no further attempts will be
            made and the task will be deleted.
            If zero, then the task age is unlimited.
          * `minBackoff` (`pulumi.Input[str]`) - A task will be scheduled for retry between minBackoff and
            maxBackoff duration after it fails, if the queue's RetryConfig
            specifies that the task should be retried.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["app_engine_routing_override"] = app_engine_routing_override
        __props__["location"] = location
        __props__["name"] = name
        __props__["project"] = project
        __props__["rate_limits"] = rate_limits
        __props__["retry_config"] = retry_config
        return Queue(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

