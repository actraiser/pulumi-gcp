# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class CryptoKey(pulumi.CustomResource):
    key_ring: pulumi.Output[str]
    labels: pulumi.Output[dict]
    name: pulumi.Output[str]
    purpose: pulumi.Output[str]
    rotation_period: pulumi.Output[str]
    self_link: pulumi.Output[str]
    version_template: pulumi.Output[dict]
    def __init__(__self__, resource_name, opts=None, key_ring=None, labels=None, name=None, purpose=None, rotation_period=None, version_template=None, __name__=None, __opts__=None):
        """
        A `CryptoKey` represents a logical key that can be used for cryptographic operations.
        
        
        > **Note:** CryptoKeys cannot be deleted from Google Cloud Platform.
        Destroying a CryptoKey will remove it from state
        and delete all CryptoKeyVersions, rendering the key unusable, but *will
        not delete the resource on the server.* When this provider destroys these keys,
        any data previously encrypted with these keys will be irrecoverable.
        For this reason, it is strongly recommended that you add lifecycle hooks
        to the resource to prevent accidental destruction.
        
        
        To get more information about CryptoKey, see:
        
        * [API documentation](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys)
        * How-to Guides
            * [Creating a key](https://cloud.google.com/kms/docs/creating-keys#create_a_key)
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/kms_crypto_key.html.markdown.
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

        if key_ring is None:
            raise TypeError("Missing required property 'key_ring'")
        __props__['key_ring'] = key_ring

        __props__['labels'] = labels

        __props__['name'] = name

        __props__['purpose'] = purpose

        __props__['rotation_period'] = rotation_period

        __props__['version_template'] = version_template

        __props__['self_link'] = None

        super(CryptoKey, __self__).__init__(
            'gcp:kms/cryptoKey:CryptoKey',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

