# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class NotificationChannel(pulumi.CustomResource):
    description: pulumi.Output[str]
    """
    An optional human-readable description of this notification channel. This description may provide additional details, beyond the display name, for the channel. This may not exceed 1024 Unicode characters.
    """
    display_name: pulumi.Output[str]
    """
    An optional human-readable name for this notification channel. It is recommended that you specify a non-empty and unique name in order to make it easier to identify the channels in your project, though this is not enforced. The display name is limited to 512 Unicode characters.
    """
    enabled: pulumi.Output[bool]
    """
    Whether notifications are forwarded to the described channel. This makes it possible to disable delivery of notifications to a particular channel without removing the channel from all alerting policies that reference the channel. This is a more convenient approach when the change is temporary and you want to receive notifications from the same set of alerting policies on the channel at some point in the future.
    """
    labels: pulumi.Output[dict]
    """
    Configuration fields that define the channel and its behavior. The
    permissible and required labels are specified in the
    NotificationChannelDescriptor corresponding to the type field.
    Labels with sensitive data are obfuscated by the API and therefore the provider cannot
    determine if there are upstream changes to these fields. They can also be configured via
    the sensitive_labels block, but cannot be configured in both places.
    """
    name: pulumi.Output[str]
    """
    The full REST resource name for this channel. The syntax is: projects/[PROJECT_ID]/notificationChannels/[CHANNEL_ID] The
    [CHANNEL_ID] is automatically assigned by the server on creation.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    sensitive_labels: pulumi.Output[dict]
    """
    Different notification type behaviors are configured primarily using the the `labels` field on this
    resource. This block contains the labels which contain secrets or passwords so that they can be marked
    sensitive and hidden from plan output. The name of the field, eg: password, will be the key
    in the `labels` map in the api request.
    Credentials may not be specified in both locations and will cause an error. Changing from one location
    to a different credential configuration in the config will require an apply to update state.  Structure is documented below.

      * `authToken` (`str`) - An authorization token for a notification channel. Channel types that support this field include: slack  **Note**: This property is sensitive and will not be displayed in the plan.
      * `password` (`str`) - An password for a notification channel. Channel types that support this field include: webhook_basicauth  **Note**: This property is sensitive and will not be displayed in the plan.
      * `serviceKey` (`str`) - An servicekey token for a notification channel. Channel types that support this field include: pagerduty  **Note**: This property is sensitive and will not be displayed in the plan.
    """
    type: pulumi.Output[str]
    """
    The type of the notification channel. This field matches the value of the NotificationChannelDescriptor.type field. See https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.notificationChannelDescriptors/list to get the list of valid values such as "email", "slack", etc...
    """
    user_labels: pulumi.Output[dict]
    """
    User-supplied key/value data that does not need to conform to the corresponding NotificationChannelDescriptor's schema, unlike the labels field. This field is intended to be used for organizing and identifying the NotificationChannel objects.The field can contain up to 64 entries. Each key and value is limited to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values can contain only lowercase letters, numerals, underscores, and dashes. Keys must begin with a letter.
    """
    verification_status: pulumi.Output[str]
    """
    Indicates whether this channel has been verified or not. On a ListNotificationChannels or GetNotificationChannel
    operation, this field is expected to be populated.If the value is UNVERIFIED, then it indicates that the channel is
    non-functioning (it both requires verification and lacks verification); otherwise, it is assumed that the channel
    works.If the channel is neither VERIFIED nor UNVERIFIED, it implies that the channel is of a type that does not require
    verification or that this specific channel has been exempted from verification because it was created prior to
    verification being required for channels of this type.This field cannot be modified using a standard
    UpdateNotificationChannel operation. To change the value of this field, you must call VerifyNotificationChannel.
    """
    def __init__(__self__, resource_name, opts=None, description=None, display_name=None, enabled=None, labels=None, project=None, sensitive_labels=None, type=None, user_labels=None, __props__=None, __name__=None, __opts__=None):
        """
        A NotificationChannel is a medium through which an alert is delivered
        when a policy violation is detected. Examples of channels include email, SMS,
        and third-party messaging applications. Fields containing sensitive information
        like authentication tokens or contact info are only partially populated on retrieval.

        Notification Channels are designed to be flexible and are made up of a supported `type`
        and labels to configure that channel. Each `type` has specific labels that need to be
        present for that channel to be correctly configured. The labels that are required to be
        present for one channel `type` are often different than those required for another.
        Due to these loose constraints it's often best to set up a channel through the UI
        and import it to the provider when setting up a brand new channel type to determine which
        labels are required.

        A list of supported channels per project the `list` endpoint can be
        accessed programmatically or through the api explorer at  https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.notificationChannelDescriptors/list .
        This provides the channel type and all of the required labels that must be passed.


        To get more information about NotificationChannel, see:

        * [API documentation](https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.notificationChannels)
        * How-to Guides
            * [Notification Options](https://cloud.google.com/monitoring/support/notification-options)
            * [Monitoring API Documentation](https://cloud.google.com/monitoring/api/v3/)

        > **Warning:** All arguments including `sensitive_labels.auth_token`, `sensitive_labels.password`, and `sensitive_labels.service_key` will be stored in the raw
        state as plain-text. [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: An optional human-readable description of this notification channel. This description may provide additional details, beyond the display name, for the channel. This may not exceed 1024 Unicode characters.
        :param pulumi.Input[str] display_name: An optional human-readable name for this notification channel. It is recommended that you specify a non-empty and unique name in order to make it easier to identify the channels in your project, though this is not enforced. The display name is limited to 512 Unicode characters.
        :param pulumi.Input[bool] enabled: Whether notifications are forwarded to the described channel. This makes it possible to disable delivery of notifications to a particular channel without removing the channel from all alerting policies that reference the channel. This is a more convenient approach when the change is temporary and you want to receive notifications from the same set of alerting policies on the channel at some point in the future.
        :param pulumi.Input[dict] labels: Configuration fields that define the channel and its behavior. The
               permissible and required labels are specified in the
               NotificationChannelDescriptor corresponding to the type field.
               Labels with sensitive data are obfuscated by the API and therefore the provider cannot
               determine if there are upstream changes to these fields. They can also be configured via
               the sensitive_labels block, but cannot be configured in both places.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[dict] sensitive_labels: Different notification type behaviors are configured primarily using the the `labels` field on this
               resource. This block contains the labels which contain secrets or passwords so that they can be marked
               sensitive and hidden from plan output. The name of the field, eg: password, will be the key
               in the `labels` map in the api request.
               Credentials may not be specified in both locations and will cause an error. Changing from one location
               to a different credential configuration in the config will require an apply to update state.  Structure is documented below.
        :param pulumi.Input[str] type: The type of the notification channel. This field matches the value of the NotificationChannelDescriptor.type field. See https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.notificationChannelDescriptors/list to get the list of valid values such as "email", "slack", etc...
        :param pulumi.Input[dict] user_labels: User-supplied key/value data that does not need to conform to the corresponding NotificationChannelDescriptor's schema, unlike the labels field. This field is intended to be used for organizing and identifying the NotificationChannel objects.The field can contain up to 64 entries. Each key and value is limited to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values can contain only lowercase letters, numerals, underscores, and dashes. Keys must begin with a letter.

        The **sensitive_labels** object supports the following:

          * `authToken` (`pulumi.Input[str]`) - An authorization token for a notification channel. Channel types that support this field include: slack  **Note**: This property is sensitive and will not be displayed in the plan.
          * `password` (`pulumi.Input[str]`) - An password for a notification channel. Channel types that support this field include: webhook_basicauth  **Note**: This property is sensitive and will not be displayed in the plan.
          * `serviceKey` (`pulumi.Input[str]`) - An servicekey token for a notification channel. Channel types that support this field include: pagerduty  **Note**: This property is sensitive and will not be displayed in the plan.
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
            __props__['display_name'] = display_name
            __props__['enabled'] = enabled
            __props__['labels'] = labels
            __props__['project'] = project
            __props__['sensitive_labels'] = sensitive_labels
            if type is None:
                raise TypeError("Missing required property 'type'")
            __props__['type'] = type
            __props__['user_labels'] = user_labels
            __props__['name'] = None
            __props__['verification_status'] = None
        super(NotificationChannel, __self__).__init__(
            'gcp:monitoring/notificationChannel:NotificationChannel',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, description=None, display_name=None, enabled=None, labels=None, name=None, project=None, sensitive_labels=None, type=None, user_labels=None, verification_status=None):
        """
        Get an existing NotificationChannel resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: An optional human-readable description of this notification channel. This description may provide additional details, beyond the display name, for the channel. This may not exceed 1024 Unicode characters.
        :param pulumi.Input[str] display_name: An optional human-readable name for this notification channel. It is recommended that you specify a non-empty and unique name in order to make it easier to identify the channels in your project, though this is not enforced. The display name is limited to 512 Unicode characters.
        :param pulumi.Input[bool] enabled: Whether notifications are forwarded to the described channel. This makes it possible to disable delivery of notifications to a particular channel without removing the channel from all alerting policies that reference the channel. This is a more convenient approach when the change is temporary and you want to receive notifications from the same set of alerting policies on the channel at some point in the future.
        :param pulumi.Input[dict] labels: Configuration fields that define the channel and its behavior. The
               permissible and required labels are specified in the
               NotificationChannelDescriptor corresponding to the type field.
               Labels with sensitive data are obfuscated by the API and therefore the provider cannot
               determine if there are upstream changes to these fields. They can also be configured via
               the sensitive_labels block, but cannot be configured in both places.
        :param pulumi.Input[str] name: The full REST resource name for this channel. The syntax is: projects/[PROJECT_ID]/notificationChannels/[CHANNEL_ID] The
               [CHANNEL_ID] is automatically assigned by the server on creation.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[dict] sensitive_labels: Different notification type behaviors are configured primarily using the the `labels` field on this
               resource. This block contains the labels which contain secrets or passwords so that they can be marked
               sensitive and hidden from plan output. The name of the field, eg: password, will be the key
               in the `labels` map in the api request.
               Credentials may not be specified in both locations and will cause an error. Changing from one location
               to a different credential configuration in the config will require an apply to update state.  Structure is documented below.
        :param pulumi.Input[str] type: The type of the notification channel. This field matches the value of the NotificationChannelDescriptor.type field. See https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.notificationChannelDescriptors/list to get the list of valid values such as "email", "slack", etc...
        :param pulumi.Input[dict] user_labels: User-supplied key/value data that does not need to conform to the corresponding NotificationChannelDescriptor's schema, unlike the labels field. This field is intended to be used for organizing and identifying the NotificationChannel objects.The field can contain up to 64 entries. Each key and value is limited to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values can contain only lowercase letters, numerals, underscores, and dashes. Keys must begin with a letter.
        :param pulumi.Input[str] verification_status: Indicates whether this channel has been verified or not. On a ListNotificationChannels or GetNotificationChannel
               operation, this field is expected to be populated.If the value is UNVERIFIED, then it indicates that the channel is
               non-functioning (it both requires verification and lacks verification); otherwise, it is assumed that the channel
               works.If the channel is neither VERIFIED nor UNVERIFIED, it implies that the channel is of a type that does not require
               verification or that this specific channel has been exempted from verification because it was created prior to
               verification being required for channels of this type.This field cannot be modified using a standard
               UpdateNotificationChannel operation. To change the value of this field, you must call VerifyNotificationChannel.

        The **sensitive_labels** object supports the following:

          * `authToken` (`pulumi.Input[str]`) - An authorization token for a notification channel. Channel types that support this field include: slack  **Note**: This property is sensitive and will not be displayed in the plan.
          * `password` (`pulumi.Input[str]`) - An password for a notification channel. Channel types that support this field include: webhook_basicauth  **Note**: This property is sensitive and will not be displayed in the plan.
          * `serviceKey` (`pulumi.Input[str]`) - An servicekey token for a notification channel. Channel types that support this field include: pagerduty  **Note**: This property is sensitive and will not be displayed in the plan.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["description"] = description
        __props__["display_name"] = display_name
        __props__["enabled"] = enabled
        __props__["labels"] = labels
        __props__["name"] = name
        __props__["project"] = project
        __props__["sensitive_labels"] = sensitive_labels
        __props__["type"] = type
        __props__["user_labels"] = user_labels
        __props__["verification_status"] = verification_status
        return NotificationChannel(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

