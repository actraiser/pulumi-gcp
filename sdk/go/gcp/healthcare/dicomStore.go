// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package healthcare

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type DicomStore struct {
	pulumi.CustomResourceState

	// Identifies the dataset addressed by this request. Must be in the format
	// 'projects/{project}/locations/{location}/datasets/{dataset}'
	Dataset pulumi.StringOutput `pulumi:"dataset"`
	// User-supplied key-value pairs used to organize DICOM stores. Label keys must be between 1 and 63 characters long, have a
	// UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression:
	// [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62} Label values are optional, must be between 1 and 63 characters long, have a
	// UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression:
	// [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more than 64 labels can be associated with a given store. An object containing a list of
	// "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The resource name for the DicomStore. ** Changing this property may recreate the Dicom store (removing all data) **
	Name pulumi.StringOutput `pulumi:"name"`
	// A nested object resource
	NotificationConfig DicomStoreNotificationConfigPtrOutput `pulumi:"notificationConfig"`
	// The fully qualified name of this dataset
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
}

// NewDicomStore registers a new resource with the given unique name, arguments, and options.
func NewDicomStore(ctx *pulumi.Context,
	name string, args *DicomStoreArgs, opts ...pulumi.ResourceOption) (*DicomStore, error) {
	if args == nil || args.Dataset == nil {
		return nil, errors.New("missing required argument 'Dataset'")
	}
	if args == nil {
		args = &DicomStoreArgs{}
	}
	var resource DicomStore
	err := ctx.RegisterResource("gcp:healthcare/dicomStore:DicomStore", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDicomStore gets an existing DicomStore resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDicomStore(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DicomStoreState, opts ...pulumi.ResourceOption) (*DicomStore, error) {
	var resource DicomStore
	err := ctx.ReadResource("gcp:healthcare/dicomStore:DicomStore", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DicomStore resources.
type dicomStoreState struct {
	// Identifies the dataset addressed by this request. Must be in the format
	// 'projects/{project}/locations/{location}/datasets/{dataset}'
	Dataset *string `pulumi:"dataset"`
	// User-supplied key-value pairs used to organize DICOM stores. Label keys must be between 1 and 63 characters long, have a
	// UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression:
	// [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62} Label values are optional, must be between 1 and 63 characters long, have a
	// UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression:
	// [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more than 64 labels can be associated with a given store. An object containing a list of
	// "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	Labels map[string]string `pulumi:"labels"`
	// The resource name for the DicomStore. ** Changing this property may recreate the Dicom store (removing all data) **
	Name *string `pulumi:"name"`
	// A nested object resource
	NotificationConfig *DicomStoreNotificationConfig `pulumi:"notificationConfig"`
	// The fully qualified name of this dataset
	SelfLink *string `pulumi:"selfLink"`
}

type DicomStoreState struct {
	// Identifies the dataset addressed by this request. Must be in the format
	// 'projects/{project}/locations/{location}/datasets/{dataset}'
	Dataset pulumi.StringPtrInput
	// User-supplied key-value pairs used to organize DICOM stores. Label keys must be between 1 and 63 characters long, have a
	// UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression:
	// [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62} Label values are optional, must be between 1 and 63 characters long, have a
	// UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression:
	// [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more than 64 labels can be associated with a given store. An object containing a list of
	// "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	Labels pulumi.StringMapInput
	// The resource name for the DicomStore. ** Changing this property may recreate the Dicom store (removing all data) **
	Name pulumi.StringPtrInput
	// A nested object resource
	NotificationConfig DicomStoreNotificationConfigPtrInput
	// The fully qualified name of this dataset
	SelfLink pulumi.StringPtrInput
}

func (DicomStoreState) ElementType() reflect.Type {
	return reflect.TypeOf((*dicomStoreState)(nil)).Elem()
}

type dicomStoreArgs struct {
	// Identifies the dataset addressed by this request. Must be in the format
	// 'projects/{project}/locations/{location}/datasets/{dataset}'
	Dataset string `pulumi:"dataset"`
	// User-supplied key-value pairs used to organize DICOM stores. Label keys must be between 1 and 63 characters long, have a
	// UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression:
	// [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62} Label values are optional, must be between 1 and 63 characters long, have a
	// UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression:
	// [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more than 64 labels can be associated with a given store. An object containing a list of
	// "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	Labels map[string]string `pulumi:"labels"`
	// The resource name for the DicomStore. ** Changing this property may recreate the Dicom store (removing all data) **
	Name *string `pulumi:"name"`
	// A nested object resource
	NotificationConfig *DicomStoreNotificationConfig `pulumi:"notificationConfig"`
}

// The set of arguments for constructing a DicomStore resource.
type DicomStoreArgs struct {
	// Identifies the dataset addressed by this request. Must be in the format
	// 'projects/{project}/locations/{location}/datasets/{dataset}'
	Dataset pulumi.StringInput
	// User-supplied key-value pairs used to organize DICOM stores. Label keys must be between 1 and 63 characters long, have a
	// UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression:
	// [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62} Label values are optional, must be between 1 and 63 characters long, have a
	// UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression:
	// [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more than 64 labels can be associated with a given store. An object containing a list of
	// "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	Labels pulumi.StringMapInput
	// The resource name for the DicomStore. ** Changing this property may recreate the Dicom store (removing all data) **
	Name pulumi.StringPtrInput
	// A nested object resource
	NotificationConfig DicomStoreNotificationConfigPtrInput
}

func (DicomStoreArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dicomStoreArgs)(nil)).Elem()
}

