# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class ManagedZone(pulumi.CustomResource):
    description: pulumi.Output[str]
    """
    A textual description field. Defaults to 'Managed by Pulumi'.
    """
    dns_name: pulumi.Output[str]
    """
    The DNS name of this managed zone, for instance "example.com.".
    """
    dnssec_config: pulumi.Output[dict]
    """
    DNSSEC configuration  Structure is documented below.

      * `defaultKeySpecs` (`list`) - Specifies parameters that will be used for generating initial DnsKeys
        for this ManagedZone. If you provide a spec for keySigning or zoneSigning,
        you must also provide one for the other.
        default_key_specs can only be updated when the state is `off`.  Structure is documented below.
        * `algorithm` (`str`) - String mnemonic specifying the DNSSEC algorithm of this key
        * `keyLength` (`float`) - Length of the keys in bits
        * `keyType` (`str`) - Specifies whether this is a key signing key (KSK) or a zone
          signing key (ZSK). Key signing keys have the Secure Entry
          Point flag set and, when active, will only be used to sign
          resource record sets of type DNSKEY. Zone signing keys do
          not have the Secure Entry Point flag set and will be used
          to sign all other types of resource record sets.
        * `kind` (`str`) - Identifies what kind of resource this is

      * `kind` (`str`) - Identifies what kind of resource this is
      * `nonExistence` (`str`) - Specifies the mechanism used to provide authenticated denial-of-existence responses.
        non_existence can only be updated when the state is `off`.
      * `state` (`str`) - Specifies whether DNSSEC is enabled, and what mode it is in
    """
    forwarding_config: pulumi.Output[dict]
    """
    The presence for this field indicates that outbound forwarding is enabled
    for this zone. The value of this field contains the set of destinations
    to forward to.  Structure is documented below.

      * `targetNameServers` (`list`) - List of target name servers to forward to. Cloud DNS will
        select the best available name server if more than
        one target is given.  Structure is documented below.
        * `forwardingPath` (`str`) - Forwarding path for this TargetNameServer. If unset or `default` Cloud DNS will make forwarding
          decision based on address ranges, i.e. RFC1918 addresses go to the VPC, Non-RFC1918 addresses go
          to the Internet. When set to `private`, Cloud DNS will always send queries through VPC for this target
        * `ipv4Address` (`str`) - IPv4 address of a target name server.
    """
    labels: pulumi.Output[dict]
    """
    A set of key/value label pairs to assign to this ManagedZone.
    """
    name: pulumi.Output[str]
    """
    User assigned name for this resource.
    Must be unique within the project.
    """
    name_servers: pulumi.Output[list]
    """
    Delegate your managed_zone to these virtual name servers; defined by the server
    """
    peering_config: pulumi.Output[dict]
    """
    The presence of this field indicates that DNS Peering is enabled for this
    zone. The value of this field contains the network to peer with.  Structure is documented below.

      * `targetNetwork` (`dict`) - The network with which to peer.  Structure is documented below.
        * `networkUrl` (`str`) - The fully qualified URL of the VPC network to forward queries to.
          This should be formatted like
          `https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}`
    """
    private_visibility_config: pulumi.Output[dict]
    """
    For privately visible zones, the set of Virtual Private Cloud
    resources that the zone is visible from.  Structure is documented below.

      * `networks` (`list`) - The list of VPC networks that can see this zone. Structure is documented below.
        * `networkUrl` (`str`) - The fully qualified URL of the VPC network to forward queries to.
          This should be formatted like
          `https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}`
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    reverse_lookup: pulumi.Output[bool]
    """
    Specifies if this is a managed reverse lookup zone. If true, Cloud DNS will resolve reverse
    lookup queries using automatically configured records for VPC resources. This only applies
    to networks listed under `private_visibility_config`.
    """
    service_directory_config: pulumi.Output[dict]
    """
    The presence of this field indicates that this zone is backed by Service Directory. The value of this field contains
    information related to the namespace associated with the zone.

      * `namespace` (`dict`) - The namespace associated with the zone.  Structure is documented below.
        * `namespaceUrl` (`str`) - The fully qualified or partial URL of the service directory namespace that should be
          associated with the zone. This should be formatted like
          `https://servicedirectory.googleapis.com/v1/projects/{project}/locations/{location}/namespaces/{namespace_id}`
          or simply `projects/{project}/locations/{location}/namespaces/{namespace_id}`
          Ignored for `public` visibility zones.
    """
    visibility: pulumi.Output[str]
    """
    The zone's visibility: public zones are exposed to the Internet,
    while private zones are visible only to Virtual Private Cloud resources.
    """
    def __init__(__self__, resource_name, opts=None, description=None, dns_name=None, dnssec_config=None, forwarding_config=None, labels=None, name=None, peering_config=None, private_visibility_config=None, project=None, reverse_lookup=None, service_directory_config=None, visibility=None, __props__=None, __name__=None, __opts__=None):
        """
        A zone is a subtree of the DNS namespace under one administrative
        responsibility. A ManagedZone is a resource that represents a DNS zone
        hosted by the Cloud DNS service.


        To get more information about ManagedZone, see:

        * [API documentation](https://cloud.google.com/dns/api/v1/managedZones)
        * How-to Guides
            * [Managing Zones](https://cloud.google.com/dns/zones/)

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A textual description field. Defaults to 'Managed by Pulumi'.
        :param pulumi.Input[str] dns_name: The DNS name of this managed zone, for instance "example.com.".
        :param pulumi.Input[dict] dnssec_config: DNSSEC configuration  Structure is documented below.
        :param pulumi.Input[dict] forwarding_config: The presence for this field indicates that outbound forwarding is enabled
               for this zone. The value of this field contains the set of destinations
               to forward to.  Structure is documented below.
        :param pulumi.Input[dict] labels: A set of key/value label pairs to assign to this ManagedZone.
        :param pulumi.Input[str] name: User assigned name for this resource.
               Must be unique within the project.
        :param pulumi.Input[dict] peering_config: The presence of this field indicates that DNS Peering is enabled for this
               zone. The value of this field contains the network to peer with.  Structure is documented below.
        :param pulumi.Input[dict] private_visibility_config: For privately visible zones, the set of Virtual Private Cloud
               resources that the zone is visible from.  Structure is documented below.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[bool] reverse_lookup: Specifies if this is a managed reverse lookup zone. If true, Cloud DNS will resolve reverse
               lookup queries using automatically configured records for VPC resources. This only applies
               to networks listed under `private_visibility_config`.
        :param pulumi.Input[dict] service_directory_config: The presence of this field indicates that this zone is backed by Service Directory. The value of this field contains
               information related to the namespace associated with the zone.
        :param pulumi.Input[str] visibility: The zone's visibility: public zones are exposed to the Internet,
               while private zones are visible only to Virtual Private Cloud resources.

        The **dnssec_config** object supports the following:

          * `defaultKeySpecs` (`pulumi.Input[list]`) - Specifies parameters that will be used for generating initial DnsKeys
            for this ManagedZone. If you provide a spec for keySigning or zoneSigning,
            you must also provide one for the other.
            default_key_specs can only be updated when the state is `off`.  Structure is documented below.
            * `algorithm` (`pulumi.Input[str]`) - String mnemonic specifying the DNSSEC algorithm of this key
            * `keyLength` (`pulumi.Input[float]`) - Length of the keys in bits
            * `keyType` (`pulumi.Input[str]`) - Specifies whether this is a key signing key (KSK) or a zone
              signing key (ZSK). Key signing keys have the Secure Entry
              Point flag set and, when active, will only be used to sign
              resource record sets of type DNSKEY. Zone signing keys do
              not have the Secure Entry Point flag set and will be used
              to sign all other types of resource record sets.
            * `kind` (`pulumi.Input[str]`) - Identifies what kind of resource this is

          * `kind` (`pulumi.Input[str]`) - Identifies what kind of resource this is
          * `nonExistence` (`pulumi.Input[str]`) - Specifies the mechanism used to provide authenticated denial-of-existence responses.
            non_existence can only be updated when the state is `off`.
          * `state` (`pulumi.Input[str]`) - Specifies whether DNSSEC is enabled, and what mode it is in

        The **forwarding_config** object supports the following:

          * `targetNameServers` (`pulumi.Input[list]`) - List of target name servers to forward to. Cloud DNS will
            select the best available name server if more than
            one target is given.  Structure is documented below.
            * `forwardingPath` (`pulumi.Input[str]`) - Forwarding path for this TargetNameServer. If unset or `default` Cloud DNS will make forwarding
              decision based on address ranges, i.e. RFC1918 addresses go to the VPC, Non-RFC1918 addresses go
              to the Internet. When set to `private`, Cloud DNS will always send queries through VPC for this target
            * `ipv4Address` (`pulumi.Input[str]`) - IPv4 address of a target name server.

        The **peering_config** object supports the following:

          * `targetNetwork` (`pulumi.Input[dict]`) - The network with which to peer.  Structure is documented below.
            * `networkUrl` (`pulumi.Input[str]`) - The fully qualified URL of the VPC network to forward queries to.
              This should be formatted like
              `https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}`

        The **private_visibility_config** object supports the following:

          * `networks` (`pulumi.Input[list]`) - The list of VPC networks that can see this zone. Structure is documented below.
            * `networkUrl` (`pulumi.Input[str]`) - The fully qualified URL of the VPC network to forward queries to.
              This should be formatted like
              `https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}`

        The **service_directory_config** object supports the following:

          * `namespace` (`pulumi.Input[dict]`) - The namespace associated with the zone.  Structure is documented below.
            * `namespaceUrl` (`pulumi.Input[str]`) - The fully qualified or partial URL of the service directory namespace that should be
              associated with the zone. This should be formatted like
              `https://servicedirectory.googleapis.com/v1/projects/{project}/locations/{location}/namespaces/{namespace_id}`
              or simply `projects/{project}/locations/{location}/namespaces/{namespace_id}`
              Ignored for `public` visibility zones.
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

            if description is None:
                description = 'Managed by Pulumi'
            __props__['description'] = description
            if dns_name is None:
                raise TypeError("Missing required property 'dns_name'")
            __props__['dns_name'] = dns_name
            __props__['dnssec_config'] = dnssec_config
            __props__['forwarding_config'] = forwarding_config
            __props__['labels'] = labels
            __props__['name'] = name
            __props__['peering_config'] = peering_config
            __props__['private_visibility_config'] = private_visibility_config
            __props__['project'] = project
            __props__['reverse_lookup'] = reverse_lookup
            __props__['service_directory_config'] = service_directory_config
            __props__['visibility'] = visibility
            __props__['name_servers'] = None
        super(ManagedZone, __self__).__init__(
            'gcp:dns/managedZone:ManagedZone',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, description=None, dns_name=None, dnssec_config=None, forwarding_config=None, labels=None, name=None, name_servers=None, peering_config=None, private_visibility_config=None, project=None, reverse_lookup=None, service_directory_config=None, visibility=None):
        """
        Get an existing ManagedZone resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A textual description field. Defaults to 'Managed by Pulumi'.
        :param pulumi.Input[str] dns_name: The DNS name of this managed zone, for instance "example.com.".
        :param pulumi.Input[dict] dnssec_config: DNSSEC configuration  Structure is documented below.
        :param pulumi.Input[dict] forwarding_config: The presence for this field indicates that outbound forwarding is enabled
               for this zone. The value of this field contains the set of destinations
               to forward to.  Structure is documented below.
        :param pulumi.Input[dict] labels: A set of key/value label pairs to assign to this ManagedZone.
        :param pulumi.Input[str] name: User assigned name for this resource.
               Must be unique within the project.
        :param pulumi.Input[list] name_servers: Delegate your managed_zone to these virtual name servers; defined by the server
        :param pulumi.Input[dict] peering_config: The presence of this field indicates that DNS Peering is enabled for this
               zone. The value of this field contains the network to peer with.  Structure is documented below.
        :param pulumi.Input[dict] private_visibility_config: For privately visible zones, the set of Virtual Private Cloud
               resources that the zone is visible from.  Structure is documented below.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[bool] reverse_lookup: Specifies if this is a managed reverse lookup zone. If true, Cloud DNS will resolve reverse
               lookup queries using automatically configured records for VPC resources. This only applies
               to networks listed under `private_visibility_config`.
        :param pulumi.Input[dict] service_directory_config: The presence of this field indicates that this zone is backed by Service Directory. The value of this field contains
               information related to the namespace associated with the zone.
        :param pulumi.Input[str] visibility: The zone's visibility: public zones are exposed to the Internet,
               while private zones are visible only to Virtual Private Cloud resources.

        The **dnssec_config** object supports the following:

          * `defaultKeySpecs` (`pulumi.Input[list]`) - Specifies parameters that will be used for generating initial DnsKeys
            for this ManagedZone. If you provide a spec for keySigning or zoneSigning,
            you must also provide one for the other.
            default_key_specs can only be updated when the state is `off`.  Structure is documented below.
            * `algorithm` (`pulumi.Input[str]`) - String mnemonic specifying the DNSSEC algorithm of this key
            * `keyLength` (`pulumi.Input[float]`) - Length of the keys in bits
            * `keyType` (`pulumi.Input[str]`) - Specifies whether this is a key signing key (KSK) or a zone
              signing key (ZSK). Key signing keys have the Secure Entry
              Point flag set and, when active, will only be used to sign
              resource record sets of type DNSKEY. Zone signing keys do
              not have the Secure Entry Point flag set and will be used
              to sign all other types of resource record sets.
            * `kind` (`pulumi.Input[str]`) - Identifies what kind of resource this is

          * `kind` (`pulumi.Input[str]`) - Identifies what kind of resource this is
          * `nonExistence` (`pulumi.Input[str]`) - Specifies the mechanism used to provide authenticated denial-of-existence responses.
            non_existence can only be updated when the state is `off`.
          * `state` (`pulumi.Input[str]`) - Specifies whether DNSSEC is enabled, and what mode it is in

        The **forwarding_config** object supports the following:

          * `targetNameServers` (`pulumi.Input[list]`) - List of target name servers to forward to. Cloud DNS will
            select the best available name server if more than
            one target is given.  Structure is documented below.
            * `forwardingPath` (`pulumi.Input[str]`) - Forwarding path for this TargetNameServer. If unset or `default` Cloud DNS will make forwarding
              decision based on address ranges, i.e. RFC1918 addresses go to the VPC, Non-RFC1918 addresses go
              to the Internet. When set to `private`, Cloud DNS will always send queries through VPC for this target
            * `ipv4Address` (`pulumi.Input[str]`) - IPv4 address of a target name server.

        The **peering_config** object supports the following:

          * `targetNetwork` (`pulumi.Input[dict]`) - The network with which to peer.  Structure is documented below.
            * `networkUrl` (`pulumi.Input[str]`) - The fully qualified URL of the VPC network to forward queries to.
              This should be formatted like
              `https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}`

        The **private_visibility_config** object supports the following:

          * `networks` (`pulumi.Input[list]`) - The list of VPC networks that can see this zone. Structure is documented below.
            * `networkUrl` (`pulumi.Input[str]`) - The fully qualified URL of the VPC network to forward queries to.
              This should be formatted like
              `https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}`

        The **service_directory_config** object supports the following:

          * `namespace` (`pulumi.Input[dict]`) - The namespace associated with the zone.  Structure is documented below.
            * `namespaceUrl` (`pulumi.Input[str]`) - The fully qualified or partial URL of the service directory namespace that should be
              associated with the zone. This should be formatted like
              `https://servicedirectory.googleapis.com/v1/projects/{project}/locations/{location}/namespaces/{namespace_id}`
              or simply `projects/{project}/locations/{location}/namespaces/{namespace_id}`
              Ignored for `public` visibility zones.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["description"] = description
        __props__["dns_name"] = dns_name
        __props__["dnssec_config"] = dnssec_config
        __props__["forwarding_config"] = forwarding_config
        __props__["labels"] = labels
        __props__["name"] = name
        __props__["name_servers"] = name_servers
        __props__["peering_config"] = peering_config
        __props__["private_visibility_config"] = private_visibility_config
        __props__["project"] = project
        __props__["reverse_lookup"] = reverse_lookup
        __props__["service_directory_config"] = service_directory_config
        __props__["visibility"] = visibility
        return ManagedZone(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

