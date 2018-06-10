# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime

class GetRegistryRepositoryResult(object):
    """
    A collection of values returned by getRegistryRepository.
    """
    def __init__(__self__, project=None, repository_url=None):
        if not project:
            raise TypeError('Missing required argument project')
        elif not isinstance(project, basestring):
            raise TypeError('Expected argument project to be a basestring')
        __self__.project = project
        if not repository_url:
            raise TypeError('Missing required argument repository_url')
        elif not isinstance(repository_url, basestring):
            raise TypeError('Expected argument repository_url to be a basestring')
        __self__.repository_url = repository_url

def get_registry_repository(project=None, region=None):
    """
    This data source fetches the project name, and provides the appropriate URLs to use for container registry for this project.
    
    The URLs are computed entirely offline - as long as the project exists, they will be valid, but this data source does not contact Google Container Registry (GCR) at any point.
    """
    __args__ = dict()

    __args__['project'] = project
    __args__['region'] = region
    __ret__ = pulumi.runtime.invoke('gcp:container/getRegistryRepository:getRegistryRepository', __args__)

    return GetRegistryRepositoryResult(
        project=__ret__['project'],
        repository_url=__ret__['repositoryUrl'])
