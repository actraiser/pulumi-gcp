// *** WARNING: this file was generated by the Lumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as lumi from "@lumi/lumi";
import * as lumirt from "@lumi/lumirt";

export class TargetHttpProxy extends lumi.NamedResource implements TargetHttpProxyArgs {
    public readonly description?: string;
    public /*out*/ readonly proxyId: string;
    public readonly targetHttpProxyName?: string;
    public readonly project?: string;
    public /*out*/ readonly selfLink: string;
    public readonly urlMap: string;

    constructor(name: string, args: TargetHttpProxyArgs) {
        super(name);
        this.description = args.description;
        this.targetHttpProxyName = args.targetHttpProxyName;
        this.project = args.project;
        if (lumirt.defaultIfComputed(args.urlMap, "") === undefined) {
            throw new Error("Property argument 'urlMap' is required, but was missing");
        }
        this.urlMap = args.urlMap;
    }
}

export interface TargetHttpProxyArgs {
    readonly description?: string;
    readonly targetHttpProxyName?: string;
    readonly project?: string;
    readonly urlMap: string;
}
