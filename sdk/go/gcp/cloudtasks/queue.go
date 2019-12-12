// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudtasks

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/cloud_tasks_queue.html.markdown.
type Queue struct {
	s *pulumi.ResourceState
}

// NewQueue registers a new resource with the given unique name, arguments, and options.
func NewQueue(ctx *pulumi.Context,
	name string, args *QueueArgs, opts ...pulumi.ResourceOpt) (*Queue, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["appEngineRoutingOverride"] = nil
		inputs["location"] = nil
		inputs["name"] = nil
		inputs["project"] = nil
		inputs["rateLimits"] = nil
		inputs["retryConfig"] = nil
	} else {
		inputs["appEngineRoutingOverride"] = args.AppEngineRoutingOverride
		inputs["location"] = args.Location
		inputs["name"] = args.Name
		inputs["project"] = args.Project
		inputs["rateLimits"] = args.RateLimits
		inputs["retryConfig"] = args.RetryConfig
	}
	s, err := ctx.RegisterResource("gcp:cloudtasks/queue:Queue", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Queue{s: s}, nil
}

// GetQueue gets an existing Queue resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetQueue(ctx *pulumi.Context,
	name string, id pulumi.ID, state *QueueState, opts ...pulumi.ResourceOpt) (*Queue, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["appEngineRoutingOverride"] = state.AppEngineRoutingOverride
		inputs["location"] = state.Location
		inputs["name"] = state.Name
		inputs["project"] = state.Project
		inputs["rateLimits"] = state.RateLimits
		inputs["retryConfig"] = state.RetryConfig
	}
	s, err := ctx.ReadResource("gcp:cloudtasks/queue:Queue", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Queue{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Queue) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Queue) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Overrides for task-level appEngineRouting. These settings apply only to App Engine tasks in this queue
func (r *Queue) AppEngineRoutingOverride() pulumi.Output {
	return r.s.State["appEngineRoutingOverride"]
}

// The location of the queue
func (r *Queue) Location() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["location"])
}

// The queue name.
func (r *Queue) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (r *Queue) Project() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["project"])
}

// Rate limits for task dispatches. The queue's actual dispatch rate is the result of: * Number of tasks in the queue *
// User-specified throttling: rateLimits, retryConfig, and the queue's state. * System throttling due to 429 (Too Many
// Requests) or 503 (Service Unavailable) responses from the worker, high error rates, or to smooth sudden large traffic
// spikes.
func (r *Queue) RateLimits() pulumi.Output {
	return r.s.State["rateLimits"]
}

// Settings that determine the retry behavior.
func (r *Queue) RetryConfig() pulumi.Output {
	return r.s.State["retryConfig"]
}

// Input properties used for looking up and filtering Queue resources.
type QueueState struct {
	// Overrides for task-level appEngineRouting. These settings apply only to App Engine tasks in this queue
	AppEngineRoutingOverride interface{}
	// The location of the queue
	Location interface{}
	// The queue name.
	Name interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	// Rate limits for task dispatches. The queue's actual dispatch rate is the result of: * Number of tasks in the queue *
	// User-specified throttling: rateLimits, retryConfig, and the queue's state. * System throttling due to 429 (Too Many
	// Requests) or 503 (Service Unavailable) responses from the worker, high error rates, or to smooth sudden large traffic
	// spikes.
	RateLimits interface{}
	// Settings that determine the retry behavior.
	RetryConfig interface{}
}

// The set of arguments for constructing a Queue resource.
type QueueArgs struct {
	// Overrides for task-level appEngineRouting. These settings apply only to App Engine tasks in this queue
	AppEngineRoutingOverride interface{}
	// The location of the queue
	Location interface{}
	// The queue name.
	Name interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	// Rate limits for task dispatches. The queue's actual dispatch rate is the result of: * Number of tasks in the queue *
	// User-specified throttling: rateLimits, retryConfig, and the queue's state. * System throttling due to 429 (Too Many
	// Requests) or 503 (Service Unavailable) responses from the worker, high error rates, or to smooth sudden large traffic
	// spikes.
	RateLimits interface{}
	// Settings that determine the retry behavior.
	RetryConfig interface{}
}
