// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package runtimeconfig

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/runtimeconfig_config_iam_policy.html.markdown.
type ConfigIamPolicy struct {
	s *pulumi.ResourceState
}

// NewConfigIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewConfigIamPolicy(ctx *pulumi.Context,
	name string, args *ConfigIamPolicyArgs, opts ...pulumi.ResourceOpt) (*ConfigIamPolicy, error) {
	if args == nil || args.Config == nil {
		return nil, errors.New("missing required argument 'Config'")
	}
	if args == nil || args.PolicyData == nil {
		return nil, errors.New("missing required argument 'PolicyData'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["config"] = nil
		inputs["policyData"] = nil
		inputs["project"] = nil
	} else {
		inputs["config"] = args.Config
		inputs["policyData"] = args.PolicyData
		inputs["project"] = args.Project
	}
	inputs["etag"] = nil
	s, err := ctx.RegisterResource("gcp:runtimeconfig/configIamPolicy:ConfigIamPolicy", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ConfigIamPolicy{s: s}, nil
}

// GetConfigIamPolicy gets an existing ConfigIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfigIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ConfigIamPolicyState, opts ...pulumi.ResourceOpt) (*ConfigIamPolicy, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["config"] = state.Config
		inputs["etag"] = state.Etag
		inputs["policyData"] = state.PolicyData
		inputs["project"] = state.Project
	}
	s, err := ctx.ReadResource("gcp:runtimeconfig/configIamPolicy:ConfigIamPolicy", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ConfigIamPolicy{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *ConfigIamPolicy) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *ConfigIamPolicy) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// Used to find the parent resource to bind the IAM policy to
func (r *ConfigIamPolicy) Config() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["config"])
}

// (Computed) The etag of the IAM policy.
func (r *ConfigIamPolicy) Etag() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["etag"])
}

// The policy data generated by
// a `organizations.getIAMPolicy` data source.
func (r *ConfigIamPolicy) PolicyData() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["policyData"])
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
func (r *ConfigIamPolicy) Project() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["project"])
}

// Input properties used for looking up and filtering ConfigIamPolicy resources.
type ConfigIamPolicyState struct {
	// Used to find the parent resource to bind the IAM policy to
	Config interface{}
	// (Computed) The etag of the IAM policy.
	Etag interface{}
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project interface{}
}

// The set of arguments for constructing a ConfigIamPolicy resource.
type ConfigIamPolicyArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	Config interface{}
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project interface{}
}