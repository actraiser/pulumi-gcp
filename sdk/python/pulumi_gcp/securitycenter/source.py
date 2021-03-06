# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Source(pulumi.CustomResource):
    description: pulumi.Output[str]
    """
    The description of the source (max of 1024 characters).
    """
    display_name: pulumi.Output[str]
    """
    The source’s display name. A source’s display name must be unique
    amongst its siblings, for example, two sources with the same parent
    can't share the same display name. The display name must start and end
    with a letter or digit, may contain letters, digits, spaces, hyphens,
    and underscores, and can be no longer than 32 characters.
    """
    name: pulumi.Output[str]
    """
    The resource name of this source, in the format 'organizations/{{organization}}/sources/{{source}}'.
    """
    organization: pulumi.Output[str]
    """
    The organization whose Cloud Security Command Center the Source
    lives in.
    """
    def __init__(__self__, resource_name, opts=None, description=None, display_name=None, organization=None, __props__=None, __name__=None, __opts__=None):
        """
        A Cloud Security Command Center's (Cloud SCC) finding source. A finding
        source is an entity or a mechanism that can produce a finding. A source is
        like a container of findings that come from the same scanner, logger,
        monitor, etc.


        To get more information about Source, see:

        * [API documentation](https://cloud.google.com/security-command-center/docs/reference/rest/v1beta1/organizations.sources)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/binary-authorization/)

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the source (max of 1024 characters).
        :param pulumi.Input[str] display_name: The source’s display name. A source’s display name must be unique
               amongst its siblings, for example, two sources with the same parent
               can't share the same display name. The display name must start and end
               with a letter or digit, may contain letters, digits, spaces, hyphens,
               and underscores, and can be no longer than 32 characters.
        :param pulumi.Input[str] organization: The organization whose Cloud Security Command Center the Source
               lives in.
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

            __props__['description'] = description
            if display_name is None:
                raise TypeError("Missing required property 'display_name'")
            __props__['display_name'] = display_name
            if organization is None:
                raise TypeError("Missing required property 'organization'")
            __props__['organization'] = organization
            __props__['name'] = None
        super(Source, __self__).__init__(
            'gcp:securitycenter/source:Source',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, description=None, display_name=None, name=None, organization=None):
        """
        Get an existing Source resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the source (max of 1024 characters).
        :param pulumi.Input[str] display_name: The source’s display name. A source’s display name must be unique
               amongst its siblings, for example, two sources with the same parent
               can't share the same display name. The display name must start and end
               with a letter or digit, may contain letters, digits, spaces, hyphens,
               and underscores, and can be no longer than 32 characters.
        :param pulumi.Input[str] name: The resource name of this source, in the format 'organizations/{{organization}}/sources/{{source}}'.
        :param pulumi.Input[str] organization: The organization whose Cloud Security Command Center the Source
               lives in.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["description"] = description
        __props__["display_name"] = display_name
        __props__["name"] = name
        __props__["organization"] = organization
        return Source(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

