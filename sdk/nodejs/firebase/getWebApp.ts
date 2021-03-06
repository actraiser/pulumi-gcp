// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * A Google Cloud Firebase web application instance
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/datasource_firebase_web_app.html.markdown.
 */
export function getWebApp(args: GetWebAppArgs, opts?: pulumi.InvokeOptions): Promise<GetWebAppResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("gcp:firebase/getWebApp:getWebApp", {
        "appId": args.appId,
    }, opts);
}

/**
 * A collection of arguments for invoking getWebApp.
 */
export interface GetWebAppArgs {
    /**
     * The appIp of name of the Firebase webApp.
     */
    readonly appId: string;
}

/**
 * A collection of values returned by getWebApp.
 */
export interface GetWebAppResult {
    readonly appId: string;
    readonly displayName: string;
    readonly name: string;
    readonly project: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
