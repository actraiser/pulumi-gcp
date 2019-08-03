# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class InstanceIamBinding(pulumi.CustomResource):
    etag: pulumi.Output[str]
    """
    (Computed) The etag of the instances's IAM policy.
    """
    instance: pulumi.Output[str]
    """
    The name or relative resource id of the instance to manage IAM policies for.
    """
    members: pulumi.Output[list]
    project: pulumi.Output[str]
    """
    The project in which the instance belongs. If it
    is not provided, this provider will use the provider default.
    """
    role: pulumi.Output[str]
    """
    The role that should be applied. Only one
    `google_bigtable_instance_iam_binding` can be used per role. Note that custom roles must be of the format
    `[projects|organizations]/{parent-name}/roles/{role-name}`.
    """
    def __init__(__self__, resource_name, opts=None, instance=None, members=None, project=None, role=None, __name__=None, __opts__=None):
        """
        Three different resources help you manage IAM policies on bigtable instances. Each of these resources serves a different use case:
        
        * `google_bigtable_instance_iam_policy`: Authoritative. Sets the IAM policy for the instance and replaces any existing policy already attached.
        * `google_bigtable_instance_iam_binding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the instance are preserved.
        * `google_bigtable_instance_iam_member`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the instance are preserved.
        
        > **Note:** `google_bigtable_instance_iam_policy` **cannot** be used in conjunction with `google_bigtable_instance_iam_binding` and `google_bigtable_instance_iam_member` or they will fight over what your policy should be. In addition, be careful not to accidentaly unset ownership of the instance as `google_bigtable_instance_iam_policy` replaces the entire policy.
        
        > **Note:** `google_bigtable_instance_iam_binding` resources **can be** used in conjunction with `google_bigtable_instance_iam_member` resources **only if** they do not grant privilege to the same role.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance: The name or relative resource id of the instance to manage IAM policies for.
        :param pulumi.Input[str] project: The project in which the instance belongs. If it
               is not provided, this provider will use the provider default.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `google_bigtable_instance_iam_binding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/bigtable_instance_iam_binding.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if instance is None:
            raise TypeError("Missing required property 'instance'")
        __props__['instance'] = instance

        if members is None:
            raise TypeError("Missing required property 'members'")
        __props__['members'] = members

        __props__['project'] = project

        if role is None:
            raise TypeError("Missing required property 'role'")
        __props__['role'] = role

        __props__['etag'] = None

        super(InstanceIamBinding, __self__).__init__(
            'gcp:bigtable/instanceIamBinding:InstanceIamBinding',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

