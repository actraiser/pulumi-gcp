# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime

class Database(pulumi.CustomResource):
    """
    Creates a new Google SQL Database on a Google SQL Database Instance. For more information, see
    the [official documentation](https://cloud.google.com/sql/),
    or the [JSON API](https://cloud.google.com/sql/docs/admin-api/v1beta4/databases).
    """
    def __init__(__self__, __name__, __opts__=None, charset=None, collation=None, instance=None, name=None, project=None):
        """Create a Database resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, basestring):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if charset and not isinstance(charset, basestring):
            raise TypeError('Expected property charset to be a basestring')
        __self__.charset = charset
        """
        The charset value. See MySQL's
        [Supported Character Sets and Collations](https://dev.mysql.com/doc/refman/5.7/en/charset-charsets.html)
        and Postgres' [Character Set Support](https://www.postgresql.org/docs/9.6/static/multibyte.html)
        for more details and supported values. Postgres databases are in [Beta](/docs/providers/google/index.html#beta-features),
        and have limited `charset` support; they only support a value of `UTF8` at creation time.
        """
        __props__['charset'] = charset

        if collation and not isinstance(collation, basestring):
            raise TypeError('Expected property collation to be a basestring')
        __self__.collation = collation
        """
        The collation value. See MySQL's
        [Supported Character Sets and Collations](https://dev.mysql.com/doc/refman/5.7/en/charset-charsets.html)
        and Postgres' [Collation Support](https://www.postgresql.org/docs/9.6/static/collation.html)
        for more details and supported values. Postgres databases are in [Beta](/docs/providers/google/index.html#beta-features),
        and have limited `collation` support; they only support a value of `en_US.UTF8` at creation time.
        """
        __props__['collation'] = collation

        if not instance:
            raise TypeError('Missing required property instance')
        elif not isinstance(instance, basestring):
            raise TypeError('Expected property instance to be a basestring')
        __self__.instance = instance
        """
        The name of containing instance.
        """
        __props__['instance'] = instance

        if name and not isinstance(name, basestring):
            raise TypeError('Expected property name to be a basestring')
        __self__.name = name
        """
        The name of the database.
        """
        __props__['name'] = name

        if project and not isinstance(project, basestring):
            raise TypeError('Expected property project to be a basestring')
        __self__.project = project
        """
        The ID of the project in which the resource belongs. If it
        is not provided, the provider project is used.
        """
        __props__['project'] = project

        __self__.self_link = pulumi.runtime.UNKNOWN
        """
        The URI of the created resource.
        """

        super(Database, __self__).__init__(
            'gcp:sql/database:Database',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'charset' in outs:
            self.charset = outs['charset']
        if 'collation' in outs:
            self.collation = outs['collation']
        if 'instance' in outs:
            self.instance = outs['instance']
        if 'name' in outs:
            self.name = outs['name']
        if 'project' in outs:
            self.project = outs['project']
        if 'selfLink' in outs:
            self.self_link = outs['selfLink']
