// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Outputs
{

    [OutputType]
    public sealed class URLMapPathMatcherPathRule
    {
        /// <summary>
        /// The list of path patterns to match. Each must start with / and the only place a
        /// * is allowed is at the end following a /. The string fed to the path matcher
        /// does not include any text after the first ? or #, and those chars are not
        /// allowed here.
        /// </summary>
        public readonly ImmutableArray<string> Paths;
        /// <summary>
        /// In response to a matching matchRule, the load balancer performs advanced routing
        /// actions like URL rewrites, header transformations, etc. prior to forwarding the
        /// request to the selected backend. If  routeAction specifies any
        /// weightedBackendServices, service must not be set. Conversely if service is set,
        /// routeAction cannot contain any  weightedBackendServices. Only one of routeAction
        /// or urlRedirect must be set.  Structure is documented below.
        /// </summary>
        public readonly Outputs.URLMapPathMatcherPathRuleRouteAction? RouteAction;
        /// <summary>
        /// The backend service or backend bucket link that should be matched by this test.
        /// </summary>
        public readonly string? Service;
        /// <summary>
        /// When this rule is matched, the request is redirected to a URL specified by
        /// urlRedirect. If urlRedirect is specified, service or routeAction must not be
        /// set.  Structure is documented below.
        /// </summary>
        public readonly Outputs.URLMapPathMatcherPathRuleUrlRedirect? UrlRedirect;

        [OutputConstructor]
        private URLMapPathMatcherPathRule(
            ImmutableArray<string> paths,

            Outputs.URLMapPathMatcherPathRuleRouteAction? routeAction,

            string? service,

            Outputs.URLMapPathMatcherPathRuleUrlRedirect? urlRedirect)
        {
            Paths = paths;
            RouteAction = routeAction;
            Service = service;
            UrlRedirect = urlRedirect;
        }
    }
}