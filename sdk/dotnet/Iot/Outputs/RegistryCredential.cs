// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Iot.Outputs
{

    [OutputType]
    public sealed class RegistryCredential
    {
        /// <summary>
        /// The certificate format and data.
        /// </summary>
        public readonly Outputs.RegistryCredentialPublicKeyCertificate PublicKeyCertificate;

        [OutputConstructor]
        private RegistryCredential(Outputs.RegistryCredentialPublicKeyCertificate publicKeyCertificate)
        {
            PublicKeyCertificate = publicKeyCertificate;
        }
    }
}
