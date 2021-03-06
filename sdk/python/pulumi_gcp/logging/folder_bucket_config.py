# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class FolderBucketConfig(pulumi.CustomResource):
    bucket_id: pulumi.Output[str]
    """
    The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
    """
    description: pulumi.Output[str]
    """
    Describes this bucket.
    """
    folder: pulumi.Output[str]
    """
    The parent resource that contains the logging bucket.
    """
    lifecycle_state: pulumi.Output[str]
    """
    The bucket's lifecycle such as active or deleted. See [LifecycleState](https://cloud.google.com/logging/docs/reference/v2/rest/v2/billingAccounts.buckets#LogBucket.LifecycleState).
    """
    location: pulumi.Output[str]
    """
    The location of the bucket. The supported locations are: "global" "us-central1"
    """
    name: pulumi.Output[str]
    """
    The resource name of the bucket. For example: "folders/my-folder-id/locations/my-location/buckets/my-bucket-id"
    """
    retention_days: pulumi.Output[float]
    """
    Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
    """
    def __init__(__self__, resource_name, opts=None, bucket_id=None, description=None, folder=None, location=None, retention_days=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a FolderBucketConfig resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bucket_id: The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
        :param pulumi.Input[str] description: Describes this bucket.
        :param pulumi.Input[str] folder: The parent resource that contains the logging bucket.
        :param pulumi.Input[str] location: The location of the bucket. The supported locations are: "global" "us-central1"
        :param pulumi.Input[float] retention_days: Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
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

            if bucket_id is None:
                raise TypeError("Missing required property 'bucket_id'")
            __props__['bucket_id'] = bucket_id
            __props__['description'] = description
            if folder is None:
                raise TypeError("Missing required property 'folder'")
            __props__['folder'] = folder
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            __props__['retention_days'] = retention_days
            __props__['lifecycle_state'] = None
            __props__['name'] = None
        super(FolderBucketConfig, __self__).__init__(
            'gcp:logging/folderBucketConfig:FolderBucketConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, bucket_id=None, description=None, folder=None, lifecycle_state=None, location=None, name=None, retention_days=None):
        """
        Get an existing FolderBucketConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bucket_id: The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
        :param pulumi.Input[str] description: Describes this bucket.
        :param pulumi.Input[str] folder: The parent resource that contains the logging bucket.
        :param pulumi.Input[str] lifecycle_state: The bucket's lifecycle such as active or deleted. See [LifecycleState](https://cloud.google.com/logging/docs/reference/v2/rest/v2/billingAccounts.buckets#LogBucket.LifecycleState).
        :param pulumi.Input[str] location: The location of the bucket. The supported locations are: "global" "us-central1"
        :param pulumi.Input[str] name: The resource name of the bucket. For example: "folders/my-folder-id/locations/my-location/buckets/my-bucket-id"
        :param pulumi.Input[float] retention_days: Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["bucket_id"] = bucket_id
        __props__["description"] = description
        __props__["folder"] = folder
        __props__["lifecycle_state"] = lifecycle_state
        __props__["location"] = location
        __props__["name"] = name
        __props__["retention_days"] = retention_days
        return FolderBucketConfig(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

