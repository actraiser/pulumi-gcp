# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class Instance(pulumi.CustomResource):
    def __init__(__self__, __name__, __opts__=None, alternative_location_id=None, authorized_network=None, display_name=None, labels=None, location_id=None, memory_size_gb=None, name=None, project=None, redis_configs=None, redis_version=None, region=None, reserved_ip_range=None, tier=None):
        """Create a Instance resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, basestring):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if alternative_location_id and not isinstance(alternative_location_id, basestring):
            raise TypeError('Expected property alternative_location_id to be a basestring')
        __self__.alternative_location_id = alternative_location_id
        __props__['alternativeLocationId'] = alternative_location_id

        if authorized_network and not isinstance(authorized_network, basestring):
            raise TypeError('Expected property authorized_network to be a basestring')
        __self__.authorized_network = authorized_network
        __props__['authorizedNetwork'] = authorized_network

        if display_name and not isinstance(display_name, basestring):
            raise TypeError('Expected property display_name to be a basestring')
        __self__.display_name = display_name
        __props__['displayName'] = display_name

        if labels and not isinstance(labels, dict):
            raise TypeError('Expected property labels to be a dict')
        __self__.labels = labels
        __props__['labels'] = labels

        if location_id and not isinstance(location_id, basestring):
            raise TypeError('Expected property location_id to be a basestring')
        __self__.location_id = location_id
        __props__['locationId'] = location_id

        if not memory_size_gb:
            raise TypeError('Missing required property memory_size_gb')
        elif not isinstance(memory_size_gb, int):
            raise TypeError('Expected property memory_size_gb to be a int')
        __self__.memory_size_gb = memory_size_gb
        __props__['memorySizeGb'] = memory_size_gb

        if name and not isinstance(name, basestring):
            raise TypeError('Expected property name to be a basestring')
        __self__.name = name
        __props__['name'] = name

        if project and not isinstance(project, basestring):
            raise TypeError('Expected property project to be a basestring')
        __self__.project = project
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the provider project is used.
        """
        __props__['project'] = project

        if redis_configs and not isinstance(redis_configs, dict):
            raise TypeError('Expected property redis_configs to be a dict')
        __self__.redis_configs = redis_configs
        __props__['redisConfigs'] = redis_configs

        if redis_version and not isinstance(redis_version, basestring):
            raise TypeError('Expected property redis_version to be a basestring')
        __self__.redis_version = redis_version
        __props__['redisVersion'] = redis_version

        if region and not isinstance(region, basestring):
            raise TypeError('Expected property region to be a basestring')
        __self__.region = region
        __props__['region'] = region

        if reserved_ip_range and not isinstance(reserved_ip_range, basestring):
            raise TypeError('Expected property reserved_ip_range to be a basestring')
        __self__.reserved_ip_range = reserved_ip_range
        __props__['reservedIpRange'] = reserved_ip_range

        if tier and not isinstance(tier, basestring):
            raise TypeError('Expected property tier to be a basestring')
        __self__.tier = tier
        __props__['tier'] = tier

        __self__.create_time = pulumi.runtime.UNKNOWN
        __self__.current_location_id = pulumi.runtime.UNKNOWN
        __self__.host = pulumi.runtime.UNKNOWN
        __self__.port = pulumi.runtime.UNKNOWN

        super(Instance, __self__).__init__(
            'gcp:redis/instance:Instance',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'alternativeLocationId' in outs:
            self.alternative_location_id = outs['alternativeLocationId']
        if 'authorizedNetwork' in outs:
            self.authorized_network = outs['authorizedNetwork']
        if 'createTime' in outs:
            self.create_time = outs['createTime']
        if 'currentLocationId' in outs:
            self.current_location_id = outs['currentLocationId']
        if 'displayName' in outs:
            self.display_name = outs['displayName']
        if 'host' in outs:
            self.host = outs['host']
        if 'labels' in outs:
            self.labels = outs['labels']
        if 'locationId' in outs:
            self.location_id = outs['locationId']
        if 'memorySizeGb' in outs:
            self.memory_size_gb = outs['memorySizeGb']
        if 'name' in outs:
            self.name = outs['name']
        if 'port' in outs:
            self.port = outs['port']
        if 'project' in outs:
            self.project = outs['project']
        if 'redisConfigs' in outs:
            self.redis_configs = outs['redisConfigs']
        if 'redisVersion' in outs:
            self.redis_version = outs['redisVersion']
        if 'region' in outs:
            self.region = outs['region']
        if 'reservedIpRange' in outs:
            self.reserved_ip_range = outs['reservedIpRange']
        if 'tier' in outs:
            self.tier = outs['tier']
