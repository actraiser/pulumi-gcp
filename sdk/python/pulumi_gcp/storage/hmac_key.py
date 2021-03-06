# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class HmacKey(pulumi.CustomResource):
    access_id: pulumi.Output[str]
    """
    The access ID of the HMAC Key.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    secret: pulumi.Output[str]
    """
    HMAC secret key material.
    """
    service_account_email: pulumi.Output[str]
    """
    The email address of the key's associated service account.
    """
    state: pulumi.Output[str]
    """
    The state of the key. Can be set to one of ACTIVE, INACTIVE.
    """
    time_created: pulumi.Output[str]
    """
    'The creation time of the HMAC key in RFC 3339 format. '
    """
    updated: pulumi.Output[str]
    """
    'The last modification time of the HMAC key metadata in RFC 3339 format.'
    """
    def __init__(__self__, resource_name, opts=None, project=None, service_account_email=None, state=None, __props__=None, __name__=None, __opts__=None):
        """
        The hmacKeys resource represents an HMAC key within Cloud Storage. The resource
        consists of a secret and HMAC key metadata. HMAC keys can be used as credentials
        for service accounts.


        To get more information about HmacKey, see:

        * [API documentation](https://cloud.google.com/storage/docs/json_api/v1/projects/hmacKeys)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/storage/docs/authentication/managing-hmackeys)

        > **Warning:** All arguments including the `secret` value will be stored in the raw
        state as plain-text. [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
        On import, the `secret` value will not be retrieved.

        > **Warning:** All arguments including `secret` will be stored in the raw
        state as plain-text. [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] service_account_email: The email address of the key's associated service account.
        :param pulumi.Input[str] state: The state of the key. Can be set to one of ACTIVE, INACTIVE.
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

            __props__['project'] = project
            if service_account_email is None:
                raise TypeError("Missing required property 'service_account_email'")
            __props__['service_account_email'] = service_account_email
            __props__['state'] = state
            __props__['access_id'] = None
            __props__['secret'] = None
            __props__['time_created'] = None
            __props__['updated'] = None
        super(HmacKey, __self__).__init__(
            'gcp:storage/hmacKey:HmacKey',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, access_id=None, project=None, secret=None, service_account_email=None, state=None, time_created=None, updated=None):
        """
        Get an existing HmacKey resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_id: The access ID of the HMAC Key.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] secret: HMAC secret key material.
        :param pulumi.Input[str] service_account_email: The email address of the key's associated service account.
        :param pulumi.Input[str] state: The state of the key. Can be set to one of ACTIVE, INACTIVE.
        :param pulumi.Input[str] time_created: 'The creation time of the HMAC key in RFC 3339 format. '
        :param pulumi.Input[str] updated: 'The last modification time of the HMAC key metadata in RFC 3339 format.'
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["access_id"] = access_id
        __props__["project"] = project
        __props__["secret"] = secret
        __props__["service_account_email"] = service_account_email
        __props__["state"] = state
        __props__["time_created"] = time_created
        __props__["updated"] = updated
        return HmacKey(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

