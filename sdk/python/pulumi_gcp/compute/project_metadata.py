# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class ProjectMetadata(pulumi.CustomResource):
    """
    Manages metadata common to all instances for a project in GCE. For more information see
    [the official documentation](https://cloud.google.com/compute/docs/storing-retrieving-metadata)
    and
    [API](https://cloud.google.com/compute/docs/reference/latest/projects/setCommonInstanceMetadata).
    
    ~> **Note:**  If you want to manage only single key/value pairs within the project metadata
    rather than the entire set, then use
    google_compute_project_metadata_item.
    """
    def __init__(__self__, __name__, __opts__=None, metadata=None, project=None):
        """Create a ProjectMetadata resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, basestring):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not metadata:
            raise TypeError('Missing required property metadata')
        elif not isinstance(metadata, dict):
            raise TypeError('Expected property metadata to be a dict')
        __self__.metadata = metadata
        """
        A series of key value pairs. Changing this resource
        updates the GCE state.
        """
        __props__['metadata'] = metadata

        if project and not isinstance(project, basestring):
            raise TypeError('Expected property project to be a basestring')
        __self__.project = project
        """
        The ID of the project in which the resource belongs. If it
        is not provided, the provider project is used.
        """
        __props__['project'] = project

        super(ProjectMetadata, __self__).__init__(
            'gcp:compute/projectMetadata:ProjectMetadata',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'metadata' in outs:
            self.metadata = outs['metadata']
        if 'project' in outs:
            self.project = outs['project']
