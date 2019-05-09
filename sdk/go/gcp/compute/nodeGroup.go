// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Represents a NodeGroup resource to manage a group of sole-tenant nodes.
// 
// 
// To get more information about NodeGroup, see:
// 
// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/nodeGroups)
// * How-to Guides
//     * [Sole-Tenant Nodes](https://cloud.google.com/compute/docs/nodes/)
// 
// > **Warning:** Due to limitations of the API, Terraform cannot update the
// number of nodes in a node group and changes to node group size either
// through Terraform config or through external changes will cause
// Terraform to delete and recreate the node group.
type NodeGroup struct {
	s *pulumi.ResourceState
}

// NewNodeGroup registers a new resource with the given unique name, arguments, and options.
func NewNodeGroup(ctx *pulumi.Context,
	name string, args *NodeGroupArgs, opts ...pulumi.ResourceOpt) (*NodeGroup, error) {
	if args == nil || args.NodeTemplate == nil {
		return nil, errors.New("missing required argument 'NodeTemplate'")
	}
	if args == nil || args.Size == nil {
		return nil, errors.New("missing required argument 'Size'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["description"] = nil
		inputs["name"] = nil
		inputs["nodeTemplate"] = nil
		inputs["project"] = nil
		inputs["size"] = nil
		inputs["zone"] = nil
	} else {
		inputs["description"] = args.Description
		inputs["name"] = args.Name
		inputs["nodeTemplate"] = args.NodeTemplate
		inputs["project"] = args.Project
		inputs["size"] = args.Size
		inputs["zone"] = args.Zone
	}
	inputs["creationTimestamp"] = nil
	inputs["selfLink"] = nil
	s, err := ctx.RegisterResource("gcp:compute/nodeGroup:NodeGroup", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &NodeGroup{s: s}, nil
}

// GetNodeGroup gets an existing NodeGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNodeGroup(ctx *pulumi.Context,
	name string, id pulumi.ID, state *NodeGroupState, opts ...pulumi.ResourceOpt) (*NodeGroup, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["creationTimestamp"] = state.CreationTimestamp
		inputs["description"] = state.Description
		inputs["name"] = state.Name
		inputs["nodeTemplate"] = state.NodeTemplate
		inputs["project"] = state.Project
		inputs["selfLink"] = state.SelfLink
		inputs["size"] = state.Size
		inputs["zone"] = state.Zone
	}
	s, err := ctx.ReadResource("gcp:compute/nodeGroup:NodeGroup", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &NodeGroup{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *NodeGroup) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *NodeGroup) ID() *pulumi.IDOutput {
	return r.s.ID()
}

func (r *NodeGroup) CreationTimestamp() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["creationTimestamp"])
}

func (r *NodeGroup) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

func (r *NodeGroup) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

func (r *NodeGroup) NodeTemplate() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["nodeTemplate"])
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (r *NodeGroup) Project() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["project"])
}

// The URI of the created resource.
func (r *NodeGroup) SelfLink() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["selfLink"])
}

func (r *NodeGroup) Size() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["size"])
}

func (r *NodeGroup) Zone() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["zone"])
}

// Input properties used for looking up and filtering NodeGroup resources.
type NodeGroupState struct {
	CreationTimestamp interface{}
	Description interface{}
	Name interface{}
	NodeTemplate interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	// The URI of the created resource.
	SelfLink interface{}
	Size interface{}
	Zone interface{}
}

// The set of arguments for constructing a NodeGroup resource.
type NodeGroupArgs struct {
	Description interface{}
	Name interface{}
	NodeTemplate interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	Size interface{}
	Zone interface{}
}
