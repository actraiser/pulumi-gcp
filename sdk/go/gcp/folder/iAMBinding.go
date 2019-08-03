// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package folder

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Allows creation and management of a single binding within IAM policy for
// an existing Google Cloud Platform folder.
// 
// > **Note:** This resource _must not_ be used in conjunction with
//    `google_folder_iam_policy` or they will fight over what your policy
//    should be.
// 
// > **Note:** On create, this resource will overwrite members of any existing roles.
//     Use `import` and inspect the preview output to ensure
//     your existing members are preserved.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/folder_iam_binding.html.markdown.
type IAMBinding struct {
	s *pulumi.ResourceState
}

// NewIAMBinding registers a new resource with the given unique name, arguments, and options.
func NewIAMBinding(ctx *pulumi.Context,
	name string, args *IAMBindingArgs, opts ...pulumi.ResourceOpt) (*IAMBinding, error) {
	if args == nil || args.Folder == nil {
		return nil, errors.New("missing required argument 'Folder'")
	}
	if args == nil || args.Members == nil {
		return nil, errors.New("missing required argument 'Members'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["folder"] = nil
		inputs["members"] = nil
		inputs["role"] = nil
	} else {
		inputs["folder"] = args.Folder
		inputs["members"] = args.Members
		inputs["role"] = args.Role
	}
	inputs["etag"] = nil
	s, err := ctx.RegisterResource("gcp:folder/iAMBinding:IAMBinding", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &IAMBinding{s: s}, nil
}

// GetIAMBinding gets an existing IAMBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIAMBinding(ctx *pulumi.Context,
	name string, id pulumi.ID, state *IAMBindingState, opts ...pulumi.ResourceOpt) (*IAMBinding, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["etag"] = state.Etag
		inputs["folder"] = state.Folder
		inputs["members"] = state.Members
		inputs["role"] = state.Role
	}
	s, err := ctx.ReadResource("gcp:folder/iAMBinding:IAMBinding", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &IAMBinding{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *IAMBinding) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *IAMBinding) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// (Computed) The etag of the folder's IAM policy.
func (r *IAMBinding) Etag() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["etag"])
}

// The resource name of the folder the policy is attached to. Its format is folders/{folder_id}.
func (r *IAMBinding) Folder() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["folder"])
}

// An array of identites that will be granted the privilege in the `role`.
// Each entry can have one of the following values:
// * **user:{emailid}**: An email address that is associated with a specific Google account. For example, alice@gmail.com.
// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
// * For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding
func (r *IAMBinding) Members() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["members"])
}

// The role that should be applied. Only one
// `google_folder_iam_binding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (r *IAMBinding) Role() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["role"])
}

// Input properties used for looking up and filtering IAMBinding resources.
type IAMBindingState struct {
	// (Computed) The etag of the folder's IAM policy.
	Etag interface{}
	// The resource name of the folder the policy is attached to. Its format is folders/{folder_id}.
	Folder interface{}
	// An array of identites that will be granted the privilege in the `role`.
	// Each entry can have one of the following values:
	// * **user:{emailid}**: An email address that is associated with a specific Google account. For example, alice@gmail.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	// * For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding
	Members interface{}
	// The role that should be applied. Only one
	// `google_folder_iam_binding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role interface{}
}

// The set of arguments for constructing a IAMBinding resource.
type IAMBindingArgs struct {
	// The resource name of the folder the policy is attached to. Its format is folders/{folder_id}.
	Folder interface{}
	// An array of identites that will be granted the privilege in the `role`.
	// Each entry can have one of the following values:
	// * **user:{emailid}**: An email address that is associated with a specific Google account. For example, alice@gmail.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	// * For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding
	Members interface{}
	// The role that should be applied. Only one
	// `google_folder_iam_binding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role interface{}
}
