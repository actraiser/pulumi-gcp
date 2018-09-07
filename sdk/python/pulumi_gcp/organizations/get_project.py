# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class GetProjectResult(object):
    """
    A collection of values returned by getProject.
    """
    def __init__(__self__, app_engines=None, auto_create_network=None, billing_account=None, folder_id=None, labels=None, name=None, number=None, org_id=None, policy_data=None, policy_etag=None, skip_delete=None, id=None):
        if app_engines and not isinstance(app_engines, list):
            raise TypeError('Expected argument app_engines to be a list')
        __self__.app_engines = app_engines
        if auto_create_network and not isinstance(auto_create_network, bool):
            raise TypeError('Expected argument auto_create_network to be a bool')
        __self__.auto_create_network = auto_create_network
        if billing_account and not isinstance(billing_account, basestring):
            raise TypeError('Expected argument billing_account to be a basestring')
        __self__.billing_account = billing_account
        if folder_id and not isinstance(folder_id, basestring):
            raise TypeError('Expected argument folder_id to be a basestring')
        __self__.folder_id = folder_id
        if labels and not isinstance(labels, dict):
            raise TypeError('Expected argument labels to be a dict')
        __self__.labels = labels
        if name and not isinstance(name, basestring):
            raise TypeError('Expected argument name to be a basestring')
        __self__.name = name
        if number and not isinstance(number, basestring):
            raise TypeError('Expected argument number to be a basestring')
        __self__.number = number
        if org_id and not isinstance(org_id, basestring):
            raise TypeError('Expected argument org_id to be a basestring')
        __self__.org_id = org_id
        if policy_data and not isinstance(policy_data, basestring):
            raise TypeError('Expected argument policy_data to be a basestring')
        __self__.policy_data = policy_data
        if policy_etag and not isinstance(policy_etag, basestring):
            raise TypeError('Expected argument policy_etag to be a basestring')
        __self__.policy_etag = policy_etag
        if skip_delete and not isinstance(skip_delete, bool):
            raise TypeError('Expected argument skip_delete to be a bool')
        __self__.skip_delete = skip_delete
        if id and not isinstance(id, basestring):
            raise TypeError('Expected argument id to be a basestring')
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

def get_project(project_id=None):
    """
    Use this data source to get project details.
    For more information see 
    [API](https://cloud.google.com/resource-manager/reference/rest/v1/projects#Project)
    """
    __args__ = dict()

    __args__['projectId'] = project_id
    __ret__ = pulumi.runtime.invoke('gcp:organizations/getProject:getProject', __args__)

    return GetProjectResult(
        app_engines=__ret__.get('appEngines'),
        auto_create_network=__ret__.get('autoCreateNetwork'),
        billing_account=__ret__.get('billingAccount'),
        folder_id=__ret__.get('folderId'),
        labels=__ret__.get('labels'),
        name=__ret__.get('name'),
        number=__ret__.get('number'),
        org_id=__ret__.get('orgId'),
        policy_data=__ret__.get('policyData'),
        policy_etag=__ret__.get('policyEtag'),
        skip_delete=__ret__.get('skipDelete'),
        id=__ret__.get('id'))
