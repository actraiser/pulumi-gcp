# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Job(pulumi.CustomResource):
    labels: pulumi.Output[dict]
    """
    User labels to be specified for the job. Keys and values should follow the restrictions specified in the [labeling restrictions](https://cloud.google.com/compute/docs/labeling-resources#restrictions) page.
    """
    machine_type: pulumi.Output[str]
    """
    The machine type to use for the job.
    """
    max_workers: pulumi.Output[float]
    """
    The number of workers permitted to work on the job.  More workers may improve processing speed at additional cost.
    """
    name: pulumi.Output[str]
    """
    A unique name for the resource, required by Dataflow.
    """
    network: pulumi.Output[str]
    """
    The network to which VMs will be assigned. If it is not provided, "default" will be used.
    """
    on_delete: pulumi.Output[str]
    """
    One of "drain" or "cancel".  Specifies behavior of deletion during a destroy.  See above note.
    """
    parameters: pulumi.Output[dict]
    """
    Key/Value pairs to be passed to the Dataflow job (as used in the template).
    """
    project: pulumi.Output[str]
    """
    The project in which the resource belongs. If it is not provided, the provider project is used.
    """
    region: pulumi.Output[str]
    service_account_email: pulumi.Output[str]
    """
    The Service Account email used to create the job.
    """
    state: pulumi.Output[str]
    """
    The current state of the resource, selected from the [JobState enum](https://cloud.google.com/dataflow/docs/reference/rest/v1b3/projects.jobs#Job.JobState)
    """
    subnetwork: pulumi.Output[str]
    """
    The subnetwork to which VMs will be assigned. Should be of the form "regions/REGION/subnetworks/SUBNETWORK".
    """
    temp_gcs_location: pulumi.Output[str]
    """
    A writeable location on GCS for the Dataflow job to dump its temporary data.
    """
    template_gcs_path: pulumi.Output[str]
    """
    The GCS path to the Dataflow job template.
    """
    zone: pulumi.Output[str]
    """
    The zone in which the created job should run. If it is not provided, the provider zone is used.
    """
    def __init__(__self__, resource_name, opts=None, labels=None, machine_type=None, max_workers=None, name=None, network=None, on_delete=None, parameters=None, project=None, region=None, service_account_email=None, subnetwork=None, temp_gcs_location=None, template_gcs_path=None, zone=None, __name__=None, __opts__=None):
        """
        Creates a job on Dataflow, which is an implementation of Apache Beam running on Google Compute Engine. For more information see
        the official documentation for
        [Beam](https://beam.apache.org) and [Dataflow](https://cloud.google.com/dataflow/).
        
        
        ## Note on "destroy" / "apply"
        
        There are many types of Dataflow jobs.  Some Dataflow jobs run constantly, getting new data from (e.g.) a GCS bucket, and outputting data continuously.  Some jobs process a set amount of data then terminate.  All jobs can fail while running due to programming errors or other issues.  In this way, Dataflow jobs are different from most other resources.
        
        The Dataflow resource is considered 'existing' while it is in a nonterminal state.  If it reaches a terminal state (e.g. 'FAILED', 'COMPLETE', 'CANCELLED'), it will be recreated on the next 'apply'.  This is as expected for jobs which run continously, but may surprise users who use this resource for other kinds of Dataflow jobs.
        
        A Dataflow job which is 'destroyed' may be "cancelled" or "drained".  If "cancelled", the job terminates - any data written remains where it is, but no new data will be processed.  If "drained", no new data will enter the pipeline, but any data currently in the pipeline will finish being processed.  The default is "cancelled", but if a user sets `on_delete` to `"drain"` in the configuration, you may experience a long wait for your destroy to complete.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] labels: User labels to be specified for the job. Keys and values should follow the restrictions specified in the [labeling restrictions](https://cloud.google.com/compute/docs/labeling-resources#restrictions) page.
        :param pulumi.Input[str] machine_type: The machine type to use for the job.
        :param pulumi.Input[float] max_workers: The number of workers permitted to work on the job.  More workers may improve processing speed at additional cost.
        :param pulumi.Input[str] name: A unique name for the resource, required by Dataflow.
        :param pulumi.Input[str] network: The network to which VMs will be assigned. If it is not provided, "default" will be used.
        :param pulumi.Input[str] on_delete: One of "drain" or "cancel".  Specifies behavior of deletion during a destroy.  See above note.
        :param pulumi.Input[dict] parameters: Key/Value pairs to be passed to the Dataflow job (as used in the template).
        :param pulumi.Input[str] project: The project in which the resource belongs. If it is not provided, the provider project is used.
        :param pulumi.Input[str] service_account_email: The Service Account email used to create the job.
        :param pulumi.Input[str] subnetwork: The subnetwork to which VMs will be assigned. Should be of the form "regions/REGION/subnetworks/SUBNETWORK".
        :param pulumi.Input[str] temp_gcs_location: A writeable location on GCS for the Dataflow job to dump its temporary data.
        :param pulumi.Input[str] template_gcs_path: The GCS path to the Dataflow job template.
        :param pulumi.Input[str] zone: The zone in which the created job should run. If it is not provided, the provider zone is used.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/dataflow_job.html.markdown.
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

        __props__['labels'] = labels

        __props__['machine_type'] = machine_type

        __props__['max_workers'] = max_workers

        __props__['name'] = name

        __props__['network'] = network

        __props__['on_delete'] = on_delete

        __props__['parameters'] = parameters

        __props__['project'] = project

        __props__['region'] = region

        __props__['service_account_email'] = service_account_email

        __props__['subnetwork'] = subnetwork

        if temp_gcs_location is None:
            raise TypeError("Missing required property 'temp_gcs_location'")
        __props__['temp_gcs_location'] = temp_gcs_location

        if template_gcs_path is None:
            raise TypeError("Missing required property 'template_gcs_path'")
        __props__['template_gcs_path'] = template_gcs_path

        __props__['zone'] = zone

        __props__['state'] = None

        super(Job, __self__).__init__(
            'gcp:dataflow/job:Job',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

