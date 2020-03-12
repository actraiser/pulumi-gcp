// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.IdentityPlatform
{
    public partial class TenantDefaultSupportedIdpConfig : Pulumi.CustomResource
    {
        /// <summary>
        /// OAuth client ID
        /// </summary>
        [Output("clientId")]
        public Output<string> ClientId { get; private set; } = null!;

        /// <summary>
        /// OAuth client secret
        /// </summary>
        [Output("clientSecret")]
        public Output<string> ClientSecret { get; private set; } = null!;

        /// <summary>
        /// If this IDP allows the user to sign in
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// ID of the IDP. Possible values include: * 'apple.com' * 'facebook.com' * 'gc.apple.com' * 'github.com' *
        /// 'google.com' * 'linkedin.com' * 'microsoft.com' * 'playgames.google.com' * 'twitter.com' * 'yahoo.com'
        /// </summary>
        [Output("idpId")]
        public Output<string> IdpId { get; private set; } = null!;

        /// <summary>
        /// The name of the default supported IDP config resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The name of the tenant where this DefaultSupportedIdpConfig resource exists
        /// </summary>
        [Output("tenant")]
        public Output<string> Tenant { get; private set; } = null!;


        /// <summary>
        /// Create a TenantDefaultSupportedIdpConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TenantDefaultSupportedIdpConfig(string name, TenantDefaultSupportedIdpConfigArgs args, CustomResourceOptions? options = null)
            : base("gcp:identityplatform/tenantDefaultSupportedIdpConfig:TenantDefaultSupportedIdpConfig", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private TenantDefaultSupportedIdpConfig(string name, Input<string> id, TenantDefaultSupportedIdpConfigState? state = null, CustomResourceOptions? options = null)
            : base("gcp:identityplatform/tenantDefaultSupportedIdpConfig:TenantDefaultSupportedIdpConfig", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing TenantDefaultSupportedIdpConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TenantDefaultSupportedIdpConfig Get(string name, Input<string> id, TenantDefaultSupportedIdpConfigState? state = null, CustomResourceOptions? options = null)
        {
            return new TenantDefaultSupportedIdpConfig(name, id, state, options);
        }
    }

    public sealed class TenantDefaultSupportedIdpConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// OAuth client ID
        /// </summary>
        [Input("clientId", required: true)]
        public Input<string> ClientId { get; set; } = null!;

        /// <summary>
        /// OAuth client secret
        /// </summary>
        [Input("clientSecret", required: true)]
        public Input<string> ClientSecret { get; set; } = null!;

        /// <summary>
        /// If this IDP allows the user to sign in
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// ID of the IDP. Possible values include: * 'apple.com' * 'facebook.com' * 'gc.apple.com' * 'github.com' *
        /// 'google.com' * 'linkedin.com' * 'microsoft.com' * 'playgames.google.com' * 'twitter.com' * 'yahoo.com'
        /// </summary>
        [Input("idpId", required: true)]
        public Input<string> IdpId { get; set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The name of the tenant where this DefaultSupportedIdpConfig resource exists
        /// </summary>
        [Input("tenant", required: true)]
        public Input<string> Tenant { get; set; } = null!;

        public TenantDefaultSupportedIdpConfigArgs()
        {
        }
    }

    public sealed class TenantDefaultSupportedIdpConfigState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// OAuth client ID
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        /// <summary>
        /// OAuth client secret
        /// </summary>
        [Input("clientSecret")]
        public Input<string>? ClientSecret { get; set; }

        /// <summary>
        /// If this IDP allows the user to sign in
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// ID of the IDP. Possible values include: * 'apple.com' * 'facebook.com' * 'gc.apple.com' * 'github.com' *
        /// 'google.com' * 'linkedin.com' * 'microsoft.com' * 'playgames.google.com' * 'twitter.com' * 'yahoo.com'
        /// </summary>
        [Input("idpId")]
        public Input<string>? IdpId { get; set; }

        /// <summary>
        /// The name of the default supported IDP config resource
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The name of the tenant where this DefaultSupportedIdpConfig resource exists
        /// </summary>
        [Input("tenant")]
        public Input<string>? Tenant { get; set; }

        public TenantDefaultSupportedIdpConfigState()
        {
        }
    }
}
