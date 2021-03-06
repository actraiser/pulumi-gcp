# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetNetworkResult:
    """
    A collection of values returned by getNetwork.
    """
    def __init__(__self__, description=None, gateway_ipv4=None, id=None, name=None, project=None, self_link=None, subnetworks_self_links=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        __self__.description = description
        """
        Description of this network.
        """
        if gateway_ipv4 and not isinstance(gateway_ipv4, str):
            raise TypeError("Expected argument 'gateway_ipv4' to be a str")
        __self__.gateway_ipv4 = gateway_ipv4
        """
        The IP address of the gateway.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        __self__.project = project
        if self_link and not isinstance(self_link, str):
            raise TypeError("Expected argument 'self_link' to be a str")
        __self__.self_link = self_link
        """
        The URI of the resource.
        """
        if subnetworks_self_links and not isinstance(subnetworks_self_links, list):
            raise TypeError("Expected argument 'subnetworks_self_links' to be a list")
        __self__.subnetworks_self_links = subnetworks_self_links
        """
        the list of subnetworks which belong to the network
        """
class AwaitableGetNetworkResult(GetNetworkResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNetworkResult(
            description=self.description,
            gateway_ipv4=self.gateway_ipv4,
            id=self.id,
            name=self.name,
            project=self.project,
            self_link=self.self_link,
            subnetworks_self_links=self.subnetworks_self_links)

def get_network(name=None,project=None,opts=None):
    """
    Get a network within GCE from its name.




    :param str name: The name of the network.
    :param str project: The ID of the project in which the resource belongs. If it
           is not provided, the provider project is used.
    """
    __args__ = dict()


    __args__['name'] = name
    __args__['project'] = project
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:compute/getNetwork:getNetwork', __args__, opts=opts).value

    return AwaitableGetNetworkResult(
        description=__ret__.get('description'),
        gateway_ipv4=__ret__.get('gatewayIpv4'),
        id=__ret__.get('id'),
        name=__ret__.get('name'),
        project=__ret__.get('project'),
        self_link=__ret__.get('selfLink'),
        subnetworks_self_links=__ret__.get('subnetworksSelfLinks'))
