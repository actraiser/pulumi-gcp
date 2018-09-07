# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class Repository(pulumi.CustomResource):
    """
    For more information, see [the official
    documentation](https://cloud.google.com/source-repositories/) and
    [API](https://cloud.google.com/source-repositories/docs/reference/rest/v1/projects.repos)
    """
    def __init__(__self__, __name__, __opts__=None, name=None, project=None):
        """Create a Repository resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, basestring):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if name and not isinstance(name, basestring):
            raise TypeError('Expected property name to be a basestring')
        __self__.name = name
        """
        The name of the repository that will be created.
        """
        __props__['name'] = name

        if project and not isinstance(project, basestring):
            raise TypeError('Expected property project to be a basestring')
        __self__.project = project
        """
        The ID of the project in which the resource belongs. If it
        is not provided, the provider project is used.
        """
        __props__['project'] = project

        __self__.size = pulumi.runtime.UNKNOWN
        """
        The size of the repository.
        """
        __self__.url = pulumi.runtime.UNKNOWN
        """
        The url to clone the repository.
        """

        super(Repository, __self__).__init__(
            'gcp:sourcerepo/repository:Repository',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'name' in outs:
            self.name = outs['name']
        if 'project' in outs:
            self.project = outs['project']
        if 'size' in outs:
            self.size = outs['size']
        if 'url' in outs:
            self.url = outs['url']
