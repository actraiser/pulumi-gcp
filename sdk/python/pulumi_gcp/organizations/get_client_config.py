# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetClientConfigResult:
    """
    A collection of values returned by getClientConfig.
    """
    def __init__(__self__, access_token=None, id=None, project=None, region=None, zone=None):
        if access_token and not isinstance(access_token, str):
            raise TypeError("Expected argument 'access_token' to be a str")
        __self__.access_token = access_token
        """
        The OAuth2 access token used by the client to authenticate against the Google Cloud API.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        __self__.project = project
        """
        The ID of the project to apply any resources to.
        """
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        __self__.region = region
        """
        The region to operate under.
        """
        if zone and not isinstance(zone, str):
            raise TypeError("Expected argument 'zone' to be a str")
        __self__.zone = zone
        """
        The zone to operate under.
        """
class AwaitableGetClientConfigResult(GetClientConfigResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetClientConfigResult(
            access_token=self.access_token,
            id=self.id,
            project=self.project,
            region=self.region,
            zone=self.zone)

def get_client_config(opts=None):
    """
    Use this data source to access the configuration of the Google Cloud provider.
    """
    __args__ = dict()


    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:organizations/getClientConfig:getClientConfig', __args__, opts=opts).value

    return AwaitableGetClientConfigResult(
        access_token=__ret__.get('accessToken'),
        id=__ret__.get('id'),
        project=__ret__.get('project'),
        region=__ret__.get('region'),
        zone=__ret__.get('zone'))
