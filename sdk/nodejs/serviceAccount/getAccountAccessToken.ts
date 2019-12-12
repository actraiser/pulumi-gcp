// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides a google `oauth2` `accessToken` for a different service account than the one initially running the script.
 * 
 * For more information see
 * [the official documentation](https://cloud.google.com/iam/docs/creating-short-lived-service-account-credentials) as well as [iamcredentials.generateAccessToken()](https://cloud.google.com/iam/credentials/reference/rest/v1/projects.serviceAccounts/generateAccessToken)
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/service_account_access_token.html.markdown.
 */
export function getAccountAccessToken(args: GetAccountAccessTokenArgs, opts?: pulumi.InvokeOptions): Promise<GetAccountAccessTokenResult> & GetAccountAccessTokenResult {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetAccountAccessTokenResult> = pulumi.runtime.invoke("gcp:serviceAccount/getAccountAccessToken:getAccountAccessToken", {
        "delegates": args.delegates,
        "lifetime": args.lifetime,
        "scopes": args.scopes,
        "targetServiceAccount": args.targetServiceAccount,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getAccountAccessToken.
 */
export interface GetAccountAccessTokenArgs {
    /**
     * Delegate chain of approvals needed to perform full impersonation. Specify the fully qualified service account name.  (e.g. `["projects/-/serviceAccounts/delegate-svc-account@project-id.iam.gserviceaccount.com"]`)
     */
    readonly delegates?: string[];
    /**
     * Lifetime of the impersonated token (defaults to its max: `3600s`).
     */
    readonly lifetime?: string;
    /**
     * The scopes the new credential should have (e.g. `["storage-ro", "cloud-platform"]`)
     */
    readonly scopes: string[];
    /**
     * The service account _to_ impersonate (e.g. `service_B@your-project-id.iam.gserviceaccount.com`)
     */
    readonly targetServiceAccount: string;
}

/**
 * A collection of values returned by getAccountAccessToken.
 */
export interface GetAccountAccessTokenResult {
    /**
     * The `accessToken` representing the new generated identity.
     */
    readonly accessToken: string;
    readonly delegates?: string[];
    readonly lifetime?: string;
    readonly scopes: string[];
    readonly targetServiceAccount: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
