// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dataflow

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Creates a job on Dataflow, which is an implementation of Apache Beam running on Google Compute Engine. For more information see
// the official documentation for
// [Beam](https://beam.apache.org) and [Dataflow](https://cloud.google.com/dataflow/).
// 
// 
// ## Note on "destroy" / "apply"
// 
// There are many types of Dataflow jobs.  Some Dataflow jobs run constantly, getting new data from (e.g.) a GCS bucket, and outputting data continuously.  Some jobs process a set amount of data then terminate.  All jobs can fail while running due to programming errors or other issues.  In this way, Dataflow jobs are different from most other Terraform / Google resources.
// 
// The Dataflow resource is considered 'existing' while it is in a nonterminal state.  If it reaches a terminal state (e.g. 'FAILED', 'COMPLETE', 'CANCELLED'), it will be recreated on the next 'apply'.  This is as expected for jobs which run continously, but may surprise users who use this resource for other kinds of Dataflow jobs.
// 
// A Dataflow job which is 'destroyed' may be "cancelled" or "drained".  If "cancelled", the job terminates - any data written remains where it is, but no new data will be processed.  If "drained", no new data will enter the pipeline, but any data currently in the pipeline will finish being processed.  The default is "cancelled", but if a user sets `on_delete` to `"drain"` in the configuration, you may experience a long wait for your `terraform destroy` to complete.
type Job struct {
	s *pulumi.ResourceState
}

// NewJob registers a new resource with the given unique name, arguments, and options.
func NewJob(ctx *pulumi.Context,
	name string, args *JobArgs, opts ...pulumi.ResourceOpt) (*Job, error) {
	if args == nil || args.TempGcsLocation == nil {
		return nil, errors.New("missing required argument 'TempGcsLocation'")
	}
	if args == nil || args.TemplateGcsPath == nil {
		return nil, errors.New("missing required argument 'TemplateGcsPath'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["maxWorkers"] = nil
		inputs["name"] = nil
		inputs["network"] = nil
		inputs["onDelete"] = nil
		inputs["parameters"] = nil
		inputs["project"] = nil
		inputs["region"] = nil
		inputs["serviceAccountEmail"] = nil
		inputs["subnetwork"] = nil
		inputs["tempGcsLocation"] = nil
		inputs["templateGcsPath"] = nil
		inputs["zone"] = nil
	} else {
		inputs["maxWorkers"] = args.MaxWorkers
		inputs["name"] = args.Name
		inputs["network"] = args.Network
		inputs["onDelete"] = args.OnDelete
		inputs["parameters"] = args.Parameters
		inputs["project"] = args.Project
		inputs["region"] = args.Region
		inputs["serviceAccountEmail"] = args.ServiceAccountEmail
		inputs["subnetwork"] = args.Subnetwork
		inputs["tempGcsLocation"] = args.TempGcsLocation
		inputs["templateGcsPath"] = args.TemplateGcsPath
		inputs["zone"] = args.Zone
	}
	inputs["state"] = nil
	s, err := ctx.RegisterResource("gcp:dataflow/job:Job", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Job{s: s}, nil
}

// GetJob gets an existing Job resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJob(ctx *pulumi.Context,
	name string, id pulumi.ID, state *JobState, opts ...pulumi.ResourceOpt) (*Job, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["maxWorkers"] = state.MaxWorkers
		inputs["name"] = state.Name
		inputs["network"] = state.Network
		inputs["onDelete"] = state.OnDelete
		inputs["parameters"] = state.Parameters
		inputs["project"] = state.Project
		inputs["region"] = state.Region
		inputs["serviceAccountEmail"] = state.ServiceAccountEmail
		inputs["state"] = state.State
		inputs["subnetwork"] = state.Subnetwork
		inputs["tempGcsLocation"] = state.TempGcsLocation
		inputs["templateGcsPath"] = state.TemplateGcsPath
		inputs["zone"] = state.Zone
	}
	s, err := ctx.ReadResource("gcp:dataflow/job:Job", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Job{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Job) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Job) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The number of workers permitted to work on the job.  More workers may improve processing speed at additional cost.
func (r *Job) MaxWorkers() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["maxWorkers"])
}

// A unique name for the resource, required by Dataflow.
func (r *Job) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The network to which VMs will be assigned. If it is not provided, "default" will be used.
func (r *Job) Network() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["network"])
}

// One of "drain" or "cancel".  Specifies behavior of deletion during `terraform destroy`.  See above note.
func (r *Job) OnDelete() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["onDelete"])
}

// Key/Value pairs to be passed to the Dataflow job (as used in the template).
func (r *Job) Parameters() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["parameters"])
}

// The project in which the resource belongs. If it is not provided, the provider project is used.
func (r *Job) Project() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["project"])
}

func (r *Job) Region() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["region"])
}

// The Service Account email used to create the job.
func (r *Job) ServiceAccountEmail() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["serviceAccountEmail"])
}

// The current state of the resource, selected from the [JobState enum](https://cloud.google.com/dataflow/docs/reference/rest/v1b3/projects.jobs#Job.JobState)
func (r *Job) State() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["state"])
}

// The subnetwork to which VMs will be assigned. Should be of the form "regions/REGION/subnetworks/SUBNETWORK".
func (r *Job) Subnetwork() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["subnetwork"])
}

// A writeable location on GCS for the Dataflow job to dump its temporary data.
func (r *Job) TempGcsLocation() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["tempGcsLocation"])
}

// The GCS path to the Dataflow job template.
func (r *Job) TemplateGcsPath() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["templateGcsPath"])
}

// The zone in which the created job should run. If it is not provided, the provider zone is used.
func (r *Job) Zone() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["zone"])
}

// Input properties used for looking up and filtering Job resources.
type JobState struct {
	// The number of workers permitted to work on the job.  More workers may improve processing speed at additional cost.
	MaxWorkers interface{}
	// A unique name for the resource, required by Dataflow.
	Name interface{}
	// The network to which VMs will be assigned. If it is not provided, "default" will be used.
	Network interface{}
	// One of "drain" or "cancel".  Specifies behavior of deletion during `terraform destroy`.  See above note.
	OnDelete interface{}
	// Key/Value pairs to be passed to the Dataflow job (as used in the template).
	Parameters interface{}
	// The project in which the resource belongs. If it is not provided, the provider project is used.
	Project interface{}
	Region interface{}
	// The Service Account email used to create the job.
	ServiceAccountEmail interface{}
	// The current state of the resource, selected from the [JobState enum](https://cloud.google.com/dataflow/docs/reference/rest/v1b3/projects.jobs#Job.JobState)
	State interface{}
	// The subnetwork to which VMs will be assigned. Should be of the form "regions/REGION/subnetworks/SUBNETWORK".
	Subnetwork interface{}
	// A writeable location on GCS for the Dataflow job to dump its temporary data.
	TempGcsLocation interface{}
	// The GCS path to the Dataflow job template.
	TemplateGcsPath interface{}
	// The zone in which the created job should run. If it is not provided, the provider zone is used.
	Zone interface{}
}

// The set of arguments for constructing a Job resource.
type JobArgs struct {
	// The number of workers permitted to work on the job.  More workers may improve processing speed at additional cost.
	MaxWorkers interface{}
	// A unique name for the resource, required by Dataflow.
	Name interface{}
	// The network to which VMs will be assigned. If it is not provided, "default" will be used.
	Network interface{}
	// One of "drain" or "cancel".  Specifies behavior of deletion during `terraform destroy`.  See above note.
	OnDelete interface{}
	// Key/Value pairs to be passed to the Dataflow job (as used in the template).
	Parameters interface{}
	// The project in which the resource belongs. If it is not provided, the provider project is used.
	Project interface{}
	Region interface{}
	// The Service Account email used to create the job.
	ServiceAccountEmail interface{}
	// The subnetwork to which VMs will be assigned. Should be of the form "regions/REGION/subnetworks/SUBNETWORK".
	Subnetwork interface{}
	// A writeable location on GCS for the Dataflow job to dump its temporary data.
	TempGcsLocation interface{}
	// The GCS path to the Dataflow job template.
	TemplateGcsPath interface{}
	// The zone in which the created job should run. If it is not provided, the provider zone is used.
	Zone interface{}
}
