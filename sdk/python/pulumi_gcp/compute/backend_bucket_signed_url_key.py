# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class BackendBucketSignedUrlKey(pulumi.CustomResource):
    backend_bucket: pulumi.Output[str]
    key_value: pulumi.Output[str]
    name: pulumi.Output[str]
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    def __init__(__self__, resource_name, opts=None, backend_bucket=None, key_value=None, name=None, project=None, __name__=None, __opts__=None):
        """
        A key for signing Cloud CDN signed URLs for BackendBuckets.
        
        
        To get more information about BackendBucketSignedUrlKey, see:
        
        * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/backendBuckets)
        * How-to Guides
            * [Using Signed URLs](https://cloud.google.com/cdn/docs/using-signed-urls/)
        
        > **Warning:** All arguments including the key's value will be stored in the raw
        state as plain-text. [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
        Because the API does not return the sensitive key value,
        we cannot confirm or reverse changes to a key outside of this provider.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_backend_bucket_signed_url_key.html.markdown.
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

        if backend_bucket is None:
            raise TypeError("Missing required property 'backend_bucket'")
        __props__['backend_bucket'] = backend_bucket

        if key_value is None:
            raise TypeError("Missing required property 'key_value'")
        __props__['key_value'] = key_value

        __props__['name'] = name

        __props__['project'] = project

        super(BackendBucketSignedUrlKey, __self__).__init__(
            'gcp:compute/backendBucketSignedUrlKey:BackendBucketSignedUrlKey',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

