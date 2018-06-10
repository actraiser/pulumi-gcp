# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime

def get_default_service_account(email=None, project=None):
    """
    Use this data source to retrieve default service account for this project
    """
    __args__ = dict()

    __args__['email'] = email
    __args__['project'] = project
    __ret__ = pulumi.runtime.invoke('gcp:compute/getDefaultServiceAccount:getDefaultServiceAccount', __args__)

