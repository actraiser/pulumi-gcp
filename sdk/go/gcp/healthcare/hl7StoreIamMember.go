// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package healthcare

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Three different resources help you manage your IAM policy for Healthcare HL7v2 store. Each of these resources serves a different use case:
//
// * `healthcare.Hl7StoreIamPolicy`: Authoritative. Sets the IAM policy for the HL7v2 store and replaces any existing policy already attached.
// * `healthcare.Hl7StoreIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the HL7v2 store are preserved.
// * `healthcare.Hl7StoreIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the HL7v2 store are preserved.
//
// > **Note:** `healthcare.Hl7StoreIamPolicy` **cannot** be used in conjunction with `healthcare.Hl7StoreIamBinding` and `healthcare.Hl7StoreIamMember` or they will fight over what your policy should be.
//
// > **Note:** `healthcare.Hl7StoreIamBinding` resources **can be** used in conjunction with `healthcare.Hl7StoreIamMember` resources **only if** they do not grant privilege to the same role.
type Hl7StoreIamMember struct {
	pulumi.CustomResourceState

	Condition Hl7StoreIamMemberConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the HL7v2 store's IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The HL7v2 store ID, in the form
	// `{project_id}/{location_name}/{dataset_name}/{hl7_v2_store_name}` or
	// `{location_name}/{dataset_name}/{hl7_v2_store_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	Hl7V2StoreId pulumi.StringOutput `pulumi:"hl7V2StoreId"`
	Member       pulumi.StringOutput `pulumi:"member"`
	// The role that should be applied. Only one
	// `healthcare.Hl7StoreIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewHl7StoreIamMember registers a new resource with the given unique name, arguments, and options.
func NewHl7StoreIamMember(ctx *pulumi.Context,
	name string, args *Hl7StoreIamMemberArgs, opts ...pulumi.ResourceOption) (*Hl7StoreIamMember, error) {
	if args == nil || args.Hl7V2StoreId == nil {
		return nil, errors.New("missing required argument 'Hl7V2StoreId'")
	}
	if args == nil || args.Member == nil {
		return nil, errors.New("missing required argument 'Member'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil {
		args = &Hl7StoreIamMemberArgs{}
	}
	var resource Hl7StoreIamMember
	err := ctx.RegisterResource("gcp:healthcare/hl7StoreIamMember:Hl7StoreIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHl7StoreIamMember gets an existing Hl7StoreIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHl7StoreIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Hl7StoreIamMemberState, opts ...pulumi.ResourceOption) (*Hl7StoreIamMember, error) {
	var resource Hl7StoreIamMember
	err := ctx.ReadResource("gcp:healthcare/hl7StoreIamMember:Hl7StoreIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Hl7StoreIamMember resources.
type hl7StoreIamMemberState struct {
	Condition *Hl7StoreIamMemberCondition `pulumi:"condition"`
	// (Computed) The etag of the HL7v2 store's IAM policy.
	Etag *string `pulumi:"etag"`
	// The HL7v2 store ID, in the form
	// `{project_id}/{location_name}/{dataset_name}/{hl7_v2_store_name}` or
	// `{location_name}/{dataset_name}/{hl7_v2_store_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	Hl7V2StoreId *string `pulumi:"hl7V2StoreId"`
	Member       *string `pulumi:"member"`
	// The role that should be applied. Only one
	// `healthcare.Hl7StoreIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type Hl7StoreIamMemberState struct {
	Condition Hl7StoreIamMemberConditionPtrInput
	// (Computed) The etag of the HL7v2 store's IAM policy.
	Etag pulumi.StringPtrInput
	// The HL7v2 store ID, in the form
	// `{project_id}/{location_name}/{dataset_name}/{hl7_v2_store_name}` or
	// `{location_name}/{dataset_name}/{hl7_v2_store_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	Hl7V2StoreId pulumi.StringPtrInput
	Member       pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `healthcare.Hl7StoreIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (Hl7StoreIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*hl7StoreIamMemberState)(nil)).Elem()
}

type hl7StoreIamMemberArgs struct {
	Condition *Hl7StoreIamMemberCondition `pulumi:"condition"`
	// The HL7v2 store ID, in the form
	// `{project_id}/{location_name}/{dataset_name}/{hl7_v2_store_name}` or
	// `{location_name}/{dataset_name}/{hl7_v2_store_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	Hl7V2StoreId string `pulumi:"hl7V2StoreId"`
	Member       string `pulumi:"member"`
	// The role that should be applied. Only one
	// `healthcare.Hl7StoreIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a Hl7StoreIamMember resource.
type Hl7StoreIamMemberArgs struct {
	Condition Hl7StoreIamMemberConditionPtrInput
	// The HL7v2 store ID, in the form
	// `{project_id}/{location_name}/{dataset_name}/{hl7_v2_store_name}` or
	// `{location_name}/{dataset_name}/{hl7_v2_store_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	Hl7V2StoreId pulumi.StringInput
	Member       pulumi.StringInput
	// The role that should be applied. Only one
	// `healthcare.Hl7StoreIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (Hl7StoreIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hl7StoreIamMemberArgs)(nil)).Elem()
}
