// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Creates a new bucket in Google cloud storage service (GCS).
// Once a bucket has been created, its location can't be changed.
// [ACLs](https://cloud.google.com/storage/docs/access-control/lists) can be applied using the `google_storage_bucket_acl` resource.
// For more information see
// [the official documentation](https://cloud.google.com/storage/docs/overview)
// and
// [API](https://cloud.google.com/storage/docs/json_api/v1/buckets).
// 
type Bucket struct {
	s *pulumi.ResourceState
}

// NewBucket registers a new resource with the given unique name, arguments, and options.
func NewBucket(ctx *pulumi.Context,
	name string, args *BucketArgs, opts ...pulumi.ResourceOpt) (*Bucket, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["cors"] = nil
		inputs["forceDestroy"] = nil
		inputs["labels"] = nil
		inputs["lifecycleRules"] = nil
		inputs["location"] = nil
		inputs["logging"] = nil
		inputs["name"] = nil
		inputs["project"] = nil
		inputs["storageClass"] = nil
		inputs["versioning"] = nil
		inputs["websites"] = nil
	} else {
		inputs["cors"] = args.Cors
		inputs["forceDestroy"] = args.ForceDestroy
		inputs["labels"] = args.Labels
		inputs["lifecycleRules"] = args.LifecycleRules
		inputs["location"] = args.Location
		inputs["logging"] = args.Logging
		inputs["name"] = args.Name
		inputs["project"] = args.Project
		inputs["storageClass"] = args.StorageClass
		inputs["versioning"] = args.Versioning
		inputs["websites"] = args.Websites
	}
	inputs["selfLink"] = nil
	inputs["url"] = nil
	s, err := ctx.RegisterResource("gcp:storage/bucket:Bucket", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Bucket{s: s}, nil
}

// GetBucket gets an existing Bucket resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBucket(ctx *pulumi.Context,
	name string, id pulumi.ID, state *BucketState, opts ...pulumi.ResourceOpt) (*Bucket, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["cors"] = state.Cors
		inputs["forceDestroy"] = state.ForceDestroy
		inputs["labels"] = state.Labels
		inputs["lifecycleRules"] = state.LifecycleRules
		inputs["location"] = state.Location
		inputs["logging"] = state.Logging
		inputs["name"] = state.Name
		inputs["project"] = state.Project
		inputs["selfLink"] = state.SelfLink
		inputs["storageClass"] = state.StorageClass
		inputs["url"] = state.Url
		inputs["versioning"] = state.Versioning
		inputs["websites"] = state.Websites
	}
	s, err := ctx.ReadResource("gcp:storage/bucket:Bucket", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Bucket{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Bucket) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Bucket) ID() *pulumi.IDOutput {
	return r.s.ID
}

// The bucket's [Cross-Origin Resource Sharing (CORS)](https://www.w3.org/TR/cors/) configuration. Multiple blocks of this type are permitted. Structure is documented below.
func (r *Bucket) Cors() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["cors"])
}

// When deleting a bucket, this
// boolean option will delete all contained objects. If you try to delete a
// bucket that contains objects, Terraform will fail that run.
func (r *Bucket) ForceDestroy() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["forceDestroy"])
}

// A set of key/value label pairs to assign to the bucket.
func (r *Bucket) Labels() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["labels"])
}

// The bucket's [Lifecycle Rules](https://cloud.google.com/storage/docs/lifecycle#configuration) configuration. Multiple blocks of this type are permitted. Structure is documented below.
func (r *Bucket) LifecycleRules() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["lifecycleRules"])
}

// The [GCS location](https://cloud.google.com/storage/docs/bucket-locations)
func (r *Bucket) Location() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["location"])
}

// The bucket's [Access & Storage Logs](https://cloud.google.com/storage/docs/access-logs) configuration.
func (r *Bucket) Logging() *pulumi.Output {
	return r.s.State["logging"]
}

// The name of the bucket.
func (r *Bucket) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The ID of the project in which the resource belongs. If it
// is not provided, the provider project is used.
func (r *Bucket) Project() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["project"])
}

// The URI of the created resource.
func (r *Bucket) SelfLink() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["selfLink"])
}

// The [Storage Class](https://cloud.google.com/storage/docs/storage-classes) of the new bucket. Supported values include: `MULTI_REGIONAL`, `REGIONAL`, `NEARLINE`, `COLDLINE`.
func (r *Bucket) StorageClass() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["storageClass"])
}

// The base URL of the bucket, in the format `gs://<bucket-name>`.
func (r *Bucket) Url() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["url"])
}

// The bucket's [Versioning](https://cloud.google.com/storage/docs/object-versioning) configuration.
func (r *Bucket) Versioning() *pulumi.Output {
	return r.s.State["versioning"]
}

// Configuration if the bucket acts as a website. Structure is documented below.
func (r *Bucket) Websites() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["websites"])
}

// Input properties used for looking up and filtering Bucket resources.
type BucketState struct {
	// The bucket's [Cross-Origin Resource Sharing (CORS)](https://www.w3.org/TR/cors/) configuration. Multiple blocks of this type are permitted. Structure is documented below.
	Cors interface{}
	// When deleting a bucket, this
	// boolean option will delete all contained objects. If you try to delete a
	// bucket that contains objects, Terraform will fail that run.
	ForceDestroy interface{}
	// A set of key/value label pairs to assign to the bucket.
	Labels interface{}
	// The bucket's [Lifecycle Rules](https://cloud.google.com/storage/docs/lifecycle#configuration) configuration. Multiple blocks of this type are permitted. Structure is documented below.
	LifecycleRules interface{}
	// The [GCS location](https://cloud.google.com/storage/docs/bucket-locations)
	Location interface{}
	// The bucket's [Access & Storage Logs](https://cloud.google.com/storage/docs/access-logs) configuration.
	Logging interface{}
	// The name of the bucket.
	Name interface{}
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project interface{}
	// The URI of the created resource.
	SelfLink interface{}
	// The [Storage Class](https://cloud.google.com/storage/docs/storage-classes) of the new bucket. Supported values include: `MULTI_REGIONAL`, `REGIONAL`, `NEARLINE`, `COLDLINE`.
	StorageClass interface{}
	// The base URL of the bucket, in the format `gs://<bucket-name>`.
	Url interface{}
	// The bucket's [Versioning](https://cloud.google.com/storage/docs/object-versioning) configuration.
	Versioning interface{}
	// Configuration if the bucket acts as a website. Structure is documented below.
	Websites interface{}
}

// The set of arguments for constructing a Bucket resource.
type BucketArgs struct {
	// The bucket's [Cross-Origin Resource Sharing (CORS)](https://www.w3.org/TR/cors/) configuration. Multiple blocks of this type are permitted. Structure is documented below.
	Cors interface{}
	// When deleting a bucket, this
	// boolean option will delete all contained objects. If you try to delete a
	// bucket that contains objects, Terraform will fail that run.
	ForceDestroy interface{}
	// A set of key/value label pairs to assign to the bucket.
	Labels interface{}
	// The bucket's [Lifecycle Rules](https://cloud.google.com/storage/docs/lifecycle#configuration) configuration. Multiple blocks of this type are permitted. Structure is documented below.
	LifecycleRules interface{}
	// The [GCS location](https://cloud.google.com/storage/docs/bucket-locations)
	Location interface{}
	// The bucket's [Access & Storage Logs](https://cloud.google.com/storage/docs/access-logs) configuration.
	Logging interface{}
	// The name of the bucket.
	Name interface{}
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project interface{}
	// The [Storage Class](https://cloud.google.com/storage/docs/storage-classes) of the new bucket. Supported values include: `MULTI_REGIONAL`, `REGIONAL`, `NEARLINE`, `COLDLINE`.
	StorageClass interface{}
	// The bucket's [Versioning](https://cloud.google.com/storage/docs/object-versioning) configuration.
	Versioning interface{}
	// Configuration if the bucket acts as a website. Structure is documented below.
	Websites interface{}
}
