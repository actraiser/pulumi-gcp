# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetOrganizationPolicyResult:
    """
    A collection of values returned by getOrganizationPolicy.
    """
    def __init__(__self__, boolean_policies=None, constraint=None, etag=None, folder=None, id=None, list_policies=None, restore_policies=None, update_time=None, version=None):
        if boolean_policies and not isinstance(boolean_policies, list):
            raise TypeError("Expected argument 'boolean_policies' to be a list")
        __self__.boolean_policies = boolean_policies
        if constraint and not isinstance(constraint, str):
            raise TypeError("Expected argument 'constraint' to be a str")
        __self__.constraint = constraint
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        __self__.etag = etag
        if folder and not isinstance(folder, str):
            raise TypeError("Expected argument 'folder' to be a str")
        __self__.folder = folder
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if list_policies and not isinstance(list_policies, list):
            raise TypeError("Expected argument 'list_policies' to be a list")
        __self__.list_policies = list_policies
        if restore_policies and not isinstance(restore_policies, list):
            raise TypeError("Expected argument 'restore_policies' to be a list")
        __self__.restore_policies = restore_policies
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        __self__.update_time = update_time
        if version and not isinstance(version, float):
            raise TypeError("Expected argument 'version' to be a float")
        __self__.version = version
class AwaitableGetOrganizationPolicyResult(GetOrganizationPolicyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetOrganizationPolicyResult(
            boolean_policies=self.boolean_policies,
            constraint=self.constraint,
            etag=self.etag,
            folder=self.folder,
            id=self.id,
            list_policies=self.list_policies,
            restore_policies=self.restore_policies,
            update_time=self.update_time,
            version=self.version)

def get_organization_policy(constraint=None,folder=None,opts=None):
    """
    Allows management of Organization policies for a Google Folder. For more information see
    [the official
    documentation](https://cloud.google.com/resource-manager/docs/organization-policy/overview)




    :param str constraint: (Required) The name of the Constraint the Policy is configuring, for example, `serviceuser.services`. Check out the [complete list of available constraints](https://cloud.google.com/resource-manager/docs/organization-policy/understanding-constraints#available_constraints).
    :param str folder: The resource name of the folder to set the policy for. Its format is folders/{folder_id}.
    """
    __args__ = dict()


    __args__['constraint'] = constraint
    __args__['folder'] = folder
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:folder/getOrganizationPolicy:getOrganizationPolicy', __args__, opts=opts).value

    return AwaitableGetOrganizationPolicyResult(
        boolean_policies=__ret__.get('booleanPolicies'),
        constraint=__ret__.get('constraint'),
        etag=__ret__.get('etag'),
        folder=__ret__.get('folder'),
        id=__ret__.get('id'),
        list_policies=__ret__.get('listPolicies'),
        restore_policies=__ret__.get('restorePolicies'),
        update_time=__ret__.get('updateTime'),
        version=__ret__.get('version'))
