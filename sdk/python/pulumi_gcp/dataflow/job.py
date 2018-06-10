# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime

class Job(pulumi.CustomResource):
    """
    Creates a job on Dataflow, which is an implementation of Apache Beam running on Google Compute Engine. For more information see
    the official documentation for
    [Beam](https://beam.apache.org) and [Dataflow](https://cloud.google.com/dataflow/).
    
    """
    def __init__(__self__, __name__, __opts__=None, max_workers=None, name=None, on_delete=None, parameters=None, project=None, temp_gcs_location=None, template_gcs_path=None, zone=None):
        """Create a Job resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, basestring):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if max_workers and not isinstance(max_workers, int):
            raise TypeError('Expected property max_workers to be a int')
        __self__.max_workers = max_workers
        """
        The number of workers permitted to work on the job.  More workers may improve processing speed at additional cost.
        """
        __props__['maxWorkers'] = max_workers

        if name and not isinstance(name, basestring):
            raise TypeError('Expected property name to be a basestring')
        __self__.name = name
        """
        A unique name for the resource, required by Dataflow.
        """
        __props__['name'] = name

        if on_delete and not isinstance(on_delete, basestring):
            raise TypeError('Expected property on_delete to be a basestring')
        __self__.on_delete = on_delete
        """
        One of "drain" or "cancel".  Specifies behavior of deletion during `terraform destroy`.  See above note.
        """
        __props__['onDelete'] = on_delete

        if parameters and not isinstance(parameters, dict):
            raise TypeError('Expected property parameters to be a dict')
        __self__.parameters = parameters
        """
        Key/Value pairs to be passed to the Dataflow job (as used in the template).
        """
        __props__['parameters'] = parameters

        if project and not isinstance(project, basestring):
            raise TypeError('Expected property project to be a basestring')
        __self__.project = project
        """
        The project in which the resource belongs. If it is not provided, the provider project is used.
        """
        __props__['project'] = project

        if not temp_gcs_location:
            raise TypeError('Missing required property temp_gcs_location')
        elif not isinstance(temp_gcs_location, basestring):
            raise TypeError('Expected property temp_gcs_location to be a basestring')
        __self__.temp_gcs_location = temp_gcs_location
        """
        A writeable location on GCS for the Dataflow job to dump its temporary data.
        """
        __props__['tempGcsLocation'] = temp_gcs_location

        if not template_gcs_path:
            raise TypeError('Missing required property template_gcs_path')
        elif not isinstance(template_gcs_path, basestring):
            raise TypeError('Expected property template_gcs_path to be a basestring')
        __self__.template_gcs_path = template_gcs_path
        """
        The GCS path to the Dataflow job template.
        """
        __props__['templateGcsPath'] = template_gcs_path

        if zone and not isinstance(zone, basestring):
            raise TypeError('Expected property zone to be a basestring')
        __self__.zone = zone
        """
        The zone in which the created job should run. If it is not provided, the provider zone is used.
        """
        __props__['zone'] = zone

        __self__.state = pulumi.runtime.UNKNOWN
        """
        The current state of the resource, selected from the [JobState enum](https://cloud.google.com/dataflow/docs/reference/rest/v1b3/projects.jobs#Job.JobState)
        """

        super(Job, __self__).__init__(
            'gcp:dataflow/job:Job',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'maxWorkers' in outs:
            self.max_workers = outs['maxWorkers']
        if 'name' in outs:
            self.name = outs['name']
        if 'onDelete' in outs:
            self.on_delete = outs['onDelete']
        if 'parameters' in outs:
            self.parameters = outs['parameters']
        if 'project' in outs:
            self.project = outs['project']
        if 'state' in outs:
            self.state = outs['state']
        if 'tempGcsLocation' in outs:
            self.temp_gcs_location = outs['tempGcsLocation']
        if 'templateGcsPath' in outs:
            self.template_gcs_path = outs['templateGcsPath']
        if 'zone' in outs:
            self.zone = outs['zone']
