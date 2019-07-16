// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package serviceAccount

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Get the service account from a project. For more information see
// the official [API](https://cloud.google.com/compute/docs/access/service-accounts) documentation.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/service_account.html.markdown.
func LookupAccount(ctx *pulumi.Context, args *GetAccountArgs) (*GetAccountResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["accountId"] = args.AccountId
		inputs["project"] = args.Project
	}
	outputs, err := ctx.Invoke("gcp:serviceAccount/getAccount:getAccount", inputs)
	if err != nil {
		return nil, err
	}
	return &GetAccountResult{
		AccountId: outputs["accountId"],
		DisplayName: outputs["displayName"],
		Email: outputs["email"],
		Name: outputs["name"],
		Project: outputs["project"],
		UniqueId: outputs["uniqueId"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getAccount.
type GetAccountArgs struct {
	// The Service account id.  (This is the part of the service account's email field that comes before the @ symbol.)
	AccountId interface{}
	// The ID of the project that the service account is present in.
	// Defaults to the provider project configuration.
	Project interface{}
}

// A collection of values returned by getAccount.
type GetAccountResult struct {
	AccountId interface{}
	// The display name for the service account.
	DisplayName interface{}
	// The e-mail address of the service account. This value
	// should be referenced from any `google_iam_policy` data sources
	// that would grant the service account privileges.
	Email interface{}
	// The fully-qualified name of the service account.
	Name interface{}
	Project interface{}
	// The unique id of the service account.
	UniqueId interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
