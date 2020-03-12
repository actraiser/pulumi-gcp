// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package compute

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type Reservation struct {
	pulumi.CustomResourceState

	// Full or partial URL to a parent commitment. This field displays for reservations that are tied to a commitment.
	Commitment pulumi.StringOutput `pulumi:"commitment"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
	// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Reservation for instances with specific machine shapes.
	SpecificReservation ReservationSpecificReservationOutput `pulumi:"specificReservation"`
	// When set to true, only VMs that target this reservation by name can consume this reservation. Otherwise, it can be
	// consumed by VMs with affinity for any reservation. Defaults to false.
	SpecificReservationRequired pulumi.BoolPtrOutput `pulumi:"specificReservationRequired"`
	// The status of the reservation.
	Status pulumi.StringOutput `pulumi:"status"`
	// The zone where the reservation is made.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewReservation registers a new resource with the given unique name, arguments, and options.
func NewReservation(ctx *pulumi.Context,
	name string, args *ReservationArgs, opts ...pulumi.ResourceOption) (*Reservation, error) {
	if args == nil || args.SpecificReservation == nil {
		return nil, errors.New("missing required argument 'SpecificReservation'")
	}
	if args == nil || args.Zone == nil {
		return nil, errors.New("missing required argument 'Zone'")
	}
	if args == nil {
		args = &ReservationArgs{}
	}
	var resource Reservation
	err := ctx.RegisterResource("gcp:compute/reservation:Reservation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReservation gets an existing Reservation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReservation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReservationState, opts ...pulumi.ResourceOption) (*Reservation, error) {
	var resource Reservation
	err := ctx.ReadResource("gcp:compute/reservation:Reservation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Reservation resources.
type reservationState struct {
	// Full or partial URL to a parent commitment. This field displays for reservations that are tied to a commitment.
	Commitment *string `pulumi:"commitment"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
	// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
	// Reservation for instances with specific machine shapes.
	SpecificReservation *ReservationSpecificReservation `pulumi:"specificReservation"`
	// When set to true, only VMs that target this reservation by name can consume this reservation. Otherwise, it can be
	// consumed by VMs with affinity for any reservation. Defaults to false.
	SpecificReservationRequired *bool `pulumi:"specificReservationRequired"`
	// The status of the reservation.
	Status *string `pulumi:"status"`
	// The zone where the reservation is made.
	Zone *string `pulumi:"zone"`
}

type ReservationState struct {
	// Full or partial URL to a parent commitment. This field displays for reservations that are tied to a commitment.
	Commitment pulumi.StringPtrInput
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
	// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
	// Reservation for instances with specific machine shapes.
	SpecificReservation ReservationSpecificReservationPtrInput
	// When set to true, only VMs that target this reservation by name can consume this reservation. Otherwise, it can be
	// consumed by VMs with affinity for any reservation. Defaults to false.
	SpecificReservationRequired pulumi.BoolPtrInput
	// The status of the reservation.
	Status pulumi.StringPtrInput
	// The zone where the reservation is made.
	Zone pulumi.StringPtrInput
}

func (ReservationState) ElementType() reflect.Type {
	return reflect.TypeOf((*reservationState)(nil)).Elem()
}

type reservationArgs struct {
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
	// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// Reservation for instances with specific machine shapes.
	SpecificReservation ReservationSpecificReservation `pulumi:"specificReservation"`
	// When set to true, only VMs that target this reservation by name can consume this reservation. Otherwise, it can be
	// consumed by VMs with affinity for any reservation. Defaults to false.
	SpecificReservationRequired *bool `pulumi:"specificReservationRequired"`
	// The zone where the reservation is made.
	Zone string `pulumi:"zone"`
}

// The set of arguments for constructing a Reservation resource.
type ReservationArgs struct {
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
	// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// Reservation for instances with specific machine shapes.
	SpecificReservation ReservationSpecificReservationInput
	// When set to true, only VMs that target this reservation by name can consume this reservation. Otherwise, it can be
	// consumed by VMs with affinity for any reservation. Defaults to false.
	SpecificReservationRequired pulumi.BoolPtrInput
	// The zone where the reservation is made.
	Zone pulumi.StringInput
}

func (ReservationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*reservationArgs)(nil)).Elem()
}

