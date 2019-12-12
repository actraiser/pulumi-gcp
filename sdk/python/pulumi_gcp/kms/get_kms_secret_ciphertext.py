# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetKMSSecretCiphertextResult:
    """
    A collection of values returned by getKMSSecretCiphertext.
    """
    def __init__(__self__, ciphertext=None, crypto_key=None, plaintext=None, id=None):
        if ciphertext and not isinstance(ciphertext, str):
            raise TypeError("Expected argument 'ciphertext' to be a str")
        __self__.ciphertext = ciphertext
        """
        Contains the result of encrypting the provided plaintext, encoded in base64.
        """
        if crypto_key and not isinstance(crypto_key, str):
            raise TypeError("Expected argument 'crypto_key' to be a str")
        __self__.crypto_key = crypto_key
        if plaintext and not isinstance(plaintext, str):
            raise TypeError("Expected argument 'plaintext' to be a str")
        __self__.plaintext = plaintext
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
class AwaitableGetKMSSecretCiphertextResult(GetKMSSecretCiphertextResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetKMSSecretCiphertextResult(
            ciphertext=self.ciphertext,
            crypto_key=self.crypto_key,
            plaintext=self.plaintext,
            id=self.id)

def get_kms_secret_ciphertext(crypto_key=None,plaintext=None,opts=None):
    """
    Use this data source to access information about an existing resource.
    
    :param str crypto_key: The id of the CryptoKey that will be used to
           encrypt the provided plaintext. This is represented by the format
           `{projectId}/{location}/{keyRingName}/{cryptoKeyName}`.
    :param str plaintext: The plaintext to be encrypted

    > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/kms_secret_ciphertext.html.markdown.
    """
    __args__ = dict()

    __args__['cryptoKey'] = crypto_key
    __args__['plaintext'] = plaintext
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:kms/getKMSSecretCiphertext:getKMSSecretCiphertext', __args__, opts=opts).value

    return AwaitableGetKMSSecretCiphertextResult(
        ciphertext=__ret__.get('ciphertext'),
        crypto_key=__ret__.get('cryptoKey'),
        plaintext=__ret__.get('plaintext'),
        id=__ret__.get('id'))
