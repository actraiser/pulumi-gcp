// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package monitoring

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Returns the list of IP addresses that checkers run from. For more information see
// the [official documentation](https://cloud.google.com/monitoring/uptime-checks#get-ips).
func GetUptimeCheckIPs(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetUptimeCheckIPsResult, error) {
	var rv GetUptimeCheckIPsResult
	err := ctx.Invoke("gcp:monitoring/getUptimeCheckIPs:getUptimeCheckIPs", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getUptimeCheckIPs.
type GetUptimeCheckIPsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of uptime check IPs used by Stackdriver Monitoring. Each `uptimeCheckIp` contains:
	UptimeCheckIps []GetUptimeCheckIPsUptimeCheckIp `pulumi:"uptimeCheckIps"`
}
