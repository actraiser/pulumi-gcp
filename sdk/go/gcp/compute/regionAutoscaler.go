// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type RegionAutoscaler struct {
	s *pulumi.ResourceState
}

// NewRegionAutoscaler registers a new resource with the given unique name, arguments, and options.
func NewRegionAutoscaler(ctx *pulumi.Context,
	name string, args *RegionAutoscalerArgs, opts ...pulumi.ResourceOpt) (*RegionAutoscaler, error) {
	if args == nil || args.AutoscalingPolicy == nil {
		return nil, errors.New("missing required argument 'AutoscalingPolicy'")
	}
	if args == nil || args.Target == nil {
		return nil, errors.New("missing required argument 'Target'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["autoscalingPolicy"] = nil
		inputs["description"] = nil
		inputs["name"] = nil
		inputs["project"] = nil
		inputs["region"] = nil
		inputs["target"] = nil
	} else {
		inputs["autoscalingPolicy"] = args.AutoscalingPolicy
		inputs["description"] = args.Description
		inputs["name"] = args.Name
		inputs["project"] = args.Project
		inputs["region"] = args.Region
		inputs["target"] = args.Target
	}
	inputs["creationTimestamp"] = nil
	inputs["selfLink"] = nil
	s, err := ctx.RegisterResource("gcp:compute/regionAutoscaler:RegionAutoscaler", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &RegionAutoscaler{s: s}, nil
}

// GetRegionAutoscaler gets an existing RegionAutoscaler resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegionAutoscaler(ctx *pulumi.Context,
	name string, id pulumi.ID, state *RegionAutoscalerState, opts ...pulumi.ResourceOpt) (*RegionAutoscaler, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["autoscalingPolicy"] = state.AutoscalingPolicy
		inputs["creationTimestamp"] = state.CreationTimestamp
		inputs["description"] = state.Description
		inputs["name"] = state.Name
		inputs["project"] = state.Project
		inputs["region"] = state.Region
		inputs["selfLink"] = state.SelfLink
		inputs["target"] = state.Target
	}
	s, err := ctx.ReadResource("gcp:compute/regionAutoscaler:RegionAutoscaler", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &RegionAutoscaler{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *RegionAutoscaler) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *RegionAutoscaler) ID() *pulumi.IDOutput {
	return r.s.ID
}

func (r *RegionAutoscaler) AutoscalingPolicy() *pulumi.Output {
	return r.s.State["autoscalingPolicy"]
}

func (r *RegionAutoscaler) CreationTimestamp() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["creationTimestamp"])
}

func (r *RegionAutoscaler) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

func (r *RegionAutoscaler) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

func (r *RegionAutoscaler) Project() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["project"])
}

func (r *RegionAutoscaler) Region() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["region"])
}

// The URI of the created resource.
func (r *RegionAutoscaler) SelfLink() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["selfLink"])
}

func (r *RegionAutoscaler) Target() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["target"])
}

// Input properties used for looking up and filtering RegionAutoscaler resources.
type RegionAutoscalerState struct {
	AutoscalingPolicy interface{}
	CreationTimestamp interface{}
	Description interface{}
	Name interface{}
	Project interface{}
	Region interface{}
	// The URI of the created resource.
	SelfLink interface{}
	Target interface{}
}

// The set of arguments for constructing a RegionAutoscaler resource.
type RegionAutoscalerArgs struct {
	AutoscalingPolicy interface{}
	Description interface{}
	Name interface{}
	Project interface{}
	Region interface{}
	Target interface{}
}
