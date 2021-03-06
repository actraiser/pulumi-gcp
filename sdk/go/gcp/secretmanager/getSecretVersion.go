// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package secretmanager

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Get a Secret Manager secret's version. For more information see the [official documentation](https://cloud.google.com/secret-manager/docs/) and [API](https://cloud.google.com/secret-manager/docs/reference/rest/v1beta1/projects.secrets.versions).
func LookupSecretVersion(ctx *pulumi.Context, args *LookupSecretVersionArgs, opts ...pulumi.InvokeOption) (*LookupSecretVersionResult, error) {
	var rv LookupSecretVersionResult
	err := ctx.Invoke("gcp:secretmanager/getSecretVersion:getSecretVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSecretVersion.
type LookupSecretVersionArgs struct {
	// The project to get the secret version for. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The secret to get the secret version for.
	Secret string `pulumi:"secret"`
	// The version of the secret to get. If it
	// is not provided, the latest version is retrieved.
	Version *string `pulumi:"version"`
}

// A collection of values returned by getSecretVersion.
type LookupSecretVersionResult struct {
	// The time at which the Secret was created.
	CreateTime string `pulumi:"createTime"`
	// The time at which the Secret was destroyed. Only present if state is DESTROYED.
	DestroyTime string `pulumi:"destroyTime"`
	// True if the current state of the SecretVersion is enabled.
	Enabled bool `pulumi:"enabled"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The resource name of the SecretVersion. Format:
	// `projects/{{project}}/secrets/{{secret_id}}/versions/{{version}}`
	Name    string `pulumi:"name"`
	Project string `pulumi:"project"`
	Secret  string `pulumi:"secret"`
	// The secret data. No larger than 64KiB.
	SecretData string `pulumi:"secretData"`
	Version    string `pulumi:"version"`
}
