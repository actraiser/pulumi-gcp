# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class EngineModel(pulumi.CustomResource):
    default_version: pulumi.Output[dict]
    """
    The default version of the model. This version will be used to handle
    prediction requests that do not specify a version.  Structure is documented below.

      * `name` (`str`) - The name specified for the version when it was created.
    """
    description: pulumi.Output[str]
    """
    The description specified for the model when it was created.
    """
    labels: pulumi.Output[dict]
    """
    One or more labels that you can add, to organize your models.
    """
    name: pulumi.Output[str]
    """
    The name specified for the version when it was created.
    """
    online_prediction_console_logging: pulumi.Output[bool]
    """
    If true, online prediction nodes send stderr and stdout streams to Stackdriver Logging
    """
    online_prediction_logging: pulumi.Output[bool]
    """
    If true, online prediction access logs are sent to StackDriver Logging.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    regions: pulumi.Output[str]
    """
    The list of regions where the model is going to be deployed.
    Currently only one region per model is supported
    """
    def __init__(__self__, resource_name, opts=None, default_version=None, description=None, labels=None, name=None, online_prediction_console_logging=None, online_prediction_logging=None, project=None, regions=None, __props__=None, __name__=None, __opts__=None):
        """
        Represents a machine learning solution.

        A model can have multiple versions, each of which is a deployed, trained model
        ready to receive prediction requests. The model itself is just a container.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] default_version: The default version of the model. This version will be used to handle
               prediction requests that do not specify a version.  Structure is documented below.
        :param pulumi.Input[str] description: The description specified for the model when it was created.
        :param pulumi.Input[dict] labels: One or more labels that you can add, to organize your models.
        :param pulumi.Input[str] name: The name specified for the version when it was created.
        :param pulumi.Input[bool] online_prediction_console_logging: If true, online prediction nodes send stderr and stdout streams to Stackdriver Logging
        :param pulumi.Input[bool] online_prediction_logging: If true, online prediction access logs are sent to StackDriver Logging.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] regions: The list of regions where the model is going to be deployed.
               Currently only one region per model is supported

        The **default_version** object supports the following:

          * `name` (`pulumi.Input[str]`) - The name specified for the version when it was created.
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

            __props__['default_version'] = default_version
            __props__['description'] = description
            __props__['labels'] = labels
            __props__['name'] = name
            __props__['online_prediction_console_logging'] = online_prediction_console_logging
            __props__['online_prediction_logging'] = online_prediction_logging
            __props__['project'] = project
            __props__['regions'] = regions
        super(EngineModel, __self__).__init__(
            'gcp:ml/engineModel:EngineModel',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, default_version=None, description=None, labels=None, name=None, online_prediction_console_logging=None, online_prediction_logging=None, project=None, regions=None):
        """
        Get an existing EngineModel resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] default_version: The default version of the model. This version will be used to handle
               prediction requests that do not specify a version.  Structure is documented below.
        :param pulumi.Input[str] description: The description specified for the model when it was created.
        :param pulumi.Input[dict] labels: One or more labels that you can add, to organize your models.
        :param pulumi.Input[str] name: The name specified for the version when it was created.
        :param pulumi.Input[bool] online_prediction_console_logging: If true, online prediction nodes send stderr and stdout streams to Stackdriver Logging
        :param pulumi.Input[bool] online_prediction_logging: If true, online prediction access logs are sent to StackDriver Logging.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] regions: The list of regions where the model is going to be deployed.
               Currently only one region per model is supported

        The **default_version** object supports the following:

          * `name` (`pulumi.Input[str]`) - The name specified for the version when it was created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["default_version"] = default_version
        __props__["description"] = description
        __props__["labels"] = labels
        __props__["name"] = name
        __props__["online_prediction_console_logging"] = online_prediction_console_logging
        __props__["online_prediction_logging"] = online_prediction_logging
        __props__["project"] = project
        __props__["regions"] = regions
        return EngineModel(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

