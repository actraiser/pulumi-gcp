# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime

class Notification(pulumi.CustomResource):
    """
    Creates a new notification configuration on a specified bucket, establishing a flow of event notifications from GCS to a Cloud Pub/Sub topic.
     For more information see 
    [the official documentation](https://cloud.google.com/storage/docs/pubsub-notifications) 
    and 
    [API](https://cloud.google.com/storage/docs/json_api/v1/notifications).
    """
    def __init__(__self__, __name__, __opts__=None, bucket=None, custom_attributes=None, event_types=None, object_name_prefix=None, payload_format=None, topic=None):
        """Create a Notification resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, basestring):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not bucket:
            raise TypeError('Missing required property bucket')
        elif not isinstance(bucket, basestring):
            raise TypeError('Expected property bucket to be a basestring')
        __self__.bucket = bucket
        """
        The name of the bucket.
        """
        __props__['bucket'] = bucket

        if custom_attributes and not isinstance(custom_attributes, dict):
            raise TypeError('Expected property custom_attributes to be a dict')
        __self__.custom_attributes = custom_attributes
        """
        A set of key/value attribute pairs to attach to each Cloud PubSub message published for this notification subscription
        """
        __props__['customAttributes'] = custom_attributes

        if event_types and not isinstance(event_types, list):
            raise TypeError('Expected property event_types to be a list')
        __self__.event_types = event_types
        """
        List of event type filters for this notification config. If not specified, Cloud Storage will send notifications for all event types. The valid types are: `"OBJECT_FINALIZE"`, `"OBJECT_METADATA_UPDATE"`, `"OBJECT_DELETE"`, `"OBJECT_ARCHIVE"`
        """
        __props__['eventTypes'] = event_types

        if object_name_prefix and not isinstance(object_name_prefix, basestring):
            raise TypeError('Expected property object_name_prefix to be a basestring')
        __self__.object_name_prefix = object_name_prefix
        """
        Specifies a prefix path filter for this notification config. Cloud Storage will only send notifications for objects in this bucket whose names begin with the specified prefix.
        """
        __props__['objectNamePrefix'] = object_name_prefix

        if not payload_format:
            raise TypeError('Missing required property payload_format')
        elif not isinstance(payload_format, basestring):
            raise TypeError('Expected property payload_format to be a basestring')
        __self__.payload_format = payload_format
        """
        The desired content of the Payload. One of `"JSON_API_V1"` or `"NONE"`.
        """
        __props__['payloadFormat'] = payload_format

        if not topic:
            raise TypeError('Missing required property topic')
        elif not isinstance(topic, basestring):
            raise TypeError('Expected property topic to be a basestring')
        __self__.topic = topic
        """
        The Cloud PubSub topic to which this subscription publishes. Expects either the 
        topic name, assumed to belong to the default GCP provider project, or the project-level name,
        i.e. `projects/my-gcp-project/topics/my-topic` or `my-topic`.
        """
        __props__['topic'] = topic

        __self__.self_link = pulumi.runtime.UNKNOWN
        """
        The URI of the created resource.
        """

        super(Notification, __self__).__init__(
            'gcp:storage/notification:Notification',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'bucket' in outs:
            self.bucket = outs['bucket']
        if 'customAttributes' in outs:
            self.custom_attributes = outs['customAttributes']
        if 'eventTypes' in outs:
            self.event_types = outs['eventTypes']
        if 'objectNamePrefix' in outs:
            self.object_name_prefix = outs['objectNamePrefix']
        if 'payloadFormat' in outs:
            self.payload_format = outs['payloadFormat']
        if 'selfLink' in outs:
            self.self_link = outs['selfLink']
        if 'topic' in outs:
            self.topic = outs['topic']
