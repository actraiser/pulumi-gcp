# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime

class RegionBackendService(pulumi.CustomResource):
    """
    A Region Backend Service defines a regionally-scoped group of virtual machines that will serve traffic for load balancing.
    For more information see [the official documentation](https://cloud.google.com/compute/docs/load-balancing/internal/)
    and [API](https://cloud.google.com/compute/docs/reference/latest/regionBackendServices).
    """
    def __init__(__self__, __name__, __opts__=None, backends=None, connection_draining_timeout_sec=None, description=None, health_checks=None, name=None, project=None, protocol=None, region=None, session_affinity=None, timeout_sec=None):
        """Create a RegionBackendService resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, basestring):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if backends and not isinstance(backends, list):
            raise TypeError('Expected property backends to be a list')
        __self__.backends = backends
        """
        The list of backends that serve this BackendService.
        Structure is documented below.
        """
        __props__['backends'] = backends

        if connection_draining_timeout_sec and not isinstance(connection_draining_timeout_sec, int):
            raise TypeError('Expected property connection_draining_timeout_sec to be a int')
        __self__.connection_draining_timeout_sec = connection_draining_timeout_sec
        """
        Time for which instance will be drained
        (not accept new connections, but still work to finish started ones). Defaults to `0`.
        """
        __props__['connectionDrainingTimeoutSec'] = connection_draining_timeout_sec

        if description and not isinstance(description, basestring):
            raise TypeError('Expected property description to be a basestring')
        __self__.description = description
        """
        Textual description for the backend.
        """
        __props__['description'] = description

        if not health_checks:
            raise TypeError('Missing required property health_checks')
        elif not isinstance(health_checks, basestring):
            raise TypeError('Expected property health_checks to be a basestring')
        __self__.health_checks = health_checks
        """
        Specifies a list of health checks
        for checking the health of the backend service. Currently at most
        one health check can be specified, and a health check is required.
        """
        __props__['healthChecks'] = health_checks

        if name and not isinstance(name, basestring):
            raise TypeError('Expected property name to be a basestring')
        __self__.name = name
        """
        The name of the backend service.
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

        if protocol and not isinstance(protocol, basestring):
            raise TypeError('Expected property protocol to be a basestring')
        __self__.protocol = protocol
        """
        The protocol for incoming requests. Defaults to
        `HTTP`.
        """
        __props__['protocol'] = protocol

        if region and not isinstance(region, basestring):
            raise TypeError('Expected property region to be a basestring')
        __self__.region = region
        """
        The Region in which the created address should reside.
        If it is not provided, the provider region is used.
        """
        __props__['region'] = region

        if session_affinity and not isinstance(session_affinity, basestring):
            raise TypeError('Expected property session_affinity to be a basestring')
        __self__.session_affinity = session_affinity
        """
        How to distribute load. Options are `NONE` (no
        affinity), `CLIENT_IP`, `CLIENT_IP_PROTO`, or `CLIENT_IP_PORT_PROTO`.
        Defaults to `NONE`.
        """
        __props__['sessionAffinity'] = session_affinity

        if timeout_sec and not isinstance(timeout_sec, int):
            raise TypeError('Expected property timeout_sec to be a int')
        __self__.timeout_sec = timeout_sec
        """
        The number of secs to wait for a backend to respond
        to a request before considering the request failed. Defaults to `30`.
        """
        __props__['timeoutSec'] = timeout_sec

        __self__.fingerprint = pulumi.runtime.UNKNOWN
        """
        The fingerprint of the backend service.
        """
        __self__.self_link = pulumi.runtime.UNKNOWN
        """
        The URI of the created resource.
        """

        super(RegionBackendService, __self__).__init__(
            'gcp:compute/regionBackendService:RegionBackendService',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'backends' in outs:
            self.backends = outs['backends']
        if 'connectionDrainingTimeoutSec' in outs:
            self.connection_draining_timeout_sec = outs['connectionDrainingTimeoutSec']
        if 'description' in outs:
            self.description = outs['description']
        if 'fingerprint' in outs:
            self.fingerprint = outs['fingerprint']
        if 'healthChecks' in outs:
            self.health_checks = outs['healthChecks']
        if 'name' in outs:
            self.name = outs['name']
        if 'project' in outs:
            self.project = outs['project']
        if 'protocol' in outs:
            self.protocol = outs['protocol']
        if 'region' in outs:
            self.region = outs['region']
        if 'selfLink' in outs:
            self.self_link = outs['selfLink']
        if 'sessionAffinity' in outs:
            self.session_affinity = outs['sessionAffinity']
        if 'timeoutSec' in outs:
            self.timeout_sec = outs['timeoutSec']
