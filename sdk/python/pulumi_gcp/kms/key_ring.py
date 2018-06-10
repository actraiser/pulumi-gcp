# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime

class KeyRing(pulumi.CustomResource):
    """
    Allows creation of a Google Cloud Platform KMS KeyRing. For more information see
    [the official documentation](https://cloud.google.com/kms/docs/object-hierarchy#keyring)
    and 
    [API](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings).
    
    A KeyRing is a grouping of CryptoKeys for organizational purposes. A KeyRing belongs to a Google Cloud Platform Project
    and resides in a specific location.
    
    ~> Note: KeyRings cannot be deleted from Google Cloud Platform. Destroying a Terraform-managed KeyRing will remove it
    from state but **will not delete the resource on the server**.
    """
    def __init__(__self__, __name__, __opts__=None, location=None, name=None, project=None):
        """Create a KeyRing resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, basestring):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not location:
            raise TypeError('Missing required property location')
        elif not isinstance(location, basestring):
            raise TypeError('Expected property location to be a basestring')
        __self__.location = location
        """
        The Google Cloud Platform location for the KeyRing.
        A full list of valid locations can be found by running `gcloud kms locations list`.
        """
        __props__['location'] = location

        if name and not isinstance(name, basestring):
            raise TypeError('Expected property name to be a basestring')
        __self__.name = name
        """
        The KeyRing's name.
        A KeyRing’s name must be unique within a location and match the regular expression `[a-zA-Z0-9_-]{1,63}`
        """
        __props__['name'] = name

        if project and not isinstance(project, basestring):
            raise TypeError('Expected property project to be a basestring')
        __self__.project = project
        """
        The project in which the resource belongs. If it
        is not provided, the provider project is used.
        """
        __props__['project'] = project

        super(KeyRing, __self__).__init__(
            'gcp:kms/keyRing:KeyRing',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'location' in outs:
            self.location = outs['location']
        if 'name' in outs:
            self.name = outs['name']
        if 'project' in outs:
            self.project = outs['project']
