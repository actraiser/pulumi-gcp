# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime

class Autoscalar(pulumi.CustomResource):
    """
    A Compute Engine Autoscaler automatically adds or removes virtual machines from
    a managed instance group based on increases or decreases in load. This allows
    your applications to gracefully handle increases in traffic and reduces cost
    when the need for resources is lower. You just define the autoscaling policy and
    the autoscaler performs automatic scaling based on the measured load. For more
    information, see [the official
    documentation](https://cloud.google.com/compute/docs/autoscaler/) and
    [API](https://cloud.google.com/compute/docs/reference/latest/autoscalers)
    
    """
    def __init__(__self__, __name__, __opts__=None, autoscaling_policy=None, description=None, name=None, project=None, target=None, zone=None):
        """Create a Autoscalar resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, basestring):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not autoscaling_policy:
            raise TypeError('Missing required property autoscaling_policy')
        elif not isinstance(autoscaling_policy, dict):
            raise TypeError('Expected property autoscaling_policy to be a dict')
        __self__.autoscaling_policy = autoscaling_policy
        """
        The parameters of the autoscaling
        algorithm. Structure is documented below.
        """
        __props__['autoscalingPolicy'] = autoscaling_policy

        if description and not isinstance(description, basestring):
            raise TypeError('Expected property description to be a basestring')
        __self__.description = description
        """
        An optional textual description of the instance
        group manager.
        """
        __props__['description'] = description

        if name and not isinstance(name, basestring):
            raise TypeError('Expected property name to be a basestring')
        __self__.name = name
        """
        The name of the Google Cloud Monitoring metric to follow, e.g.
        `compute.googleapis.com/instance/network/received_bytes_count`
        """
        __props__['name'] = name

        if project and not isinstance(project, basestring):
            raise TypeError('Expected property project to be a basestring')
        __self__.project = project
        """
        The ID of the project in which the resource belongs. If it
        is not provided, the provider project is used.
        """
        __props__['project'] = project

        if not target:
            raise TypeError('Missing required property target')
        elif not isinstance(target, basestring):
            raise TypeError('Expected property target to be a basestring')
        __self__.target = target
        """
        The floating point threshold where load balancing utilization
        should be. E.g. if the load balancer's `maxRatePerInstance` is 10 requests
        per second (RPS) then setting this to 0.5 would cause the group to be scaled
        such that each instance receives 5 RPS.
        """
        __props__['target'] = target

        if zone and not isinstance(zone, basestring):
            raise TypeError('Expected property zone to be a basestring')
        __self__.zone = zone
        """
        The zone of the target.
        """
        __props__['zone'] = zone

        __self__.self_link = pulumi.runtime.UNKNOWN
        """
        The URL of the created resource.
        """

        super(Autoscalar, __self__).__init__(
            'gcp:compute/autoscalar:Autoscalar',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'autoscalingPolicy' in outs:
            self.autoscaling_policy = outs['autoscalingPolicy']
        if 'description' in outs:
            self.description = outs['description']
        if 'name' in outs:
            self.name = outs['name']
        if 'project' in outs:
            self.project = outs['project']
        if 'selfLink' in outs:
            self.self_link = outs['selfLink']
        if 'target' in outs:
            self.target = outs['target']
        if 'zone' in outs:
            self.zone = outs['zone']
