// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_reservation.html.markdown.
 */
export class Reservation extends pulumi.CustomResource {
    /**
     * Get an existing Reservation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ReservationState, opts?: pulumi.CustomResourceOptions): Reservation {
        return new Reservation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/reservation:Reservation';

    /**
     * Returns true if the given object is an instance of Reservation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Reservation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Reservation.__pulumiType;
    }

    public /*out*/ readonly commitment!: pulumi.Output<string>;
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly project!: pulumi.Output<string>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    public readonly specificReservation!: pulumi.Output<outputs.compute.ReservationSpecificReservation>;
    public readonly specificReservationRequired!: pulumi.Output<boolean | undefined>;
    public /*out*/ readonly status!: pulumi.Output<string>;
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a Reservation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ReservationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ReservationArgs | ReservationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ReservationState | undefined;
            inputs["commitment"] = state ? state.commitment : undefined;
            inputs["creationTimestamp"] = state ? state.creationTimestamp : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
            inputs["specificReservation"] = state ? state.specificReservation : undefined;
            inputs["specificReservationRequired"] = state ? state.specificReservationRequired : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as ReservationArgs | undefined;
            if (!args || args.specificReservation === undefined) {
                throw new Error("Missing required property 'specificReservation'");
            }
            if (!args || args.zone === undefined) {
                throw new Error("Missing required property 'zone'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["specificReservation"] = args ? args.specificReservation : undefined;
            inputs["specificReservationRequired"] = args ? args.specificReservationRequired : undefined;
            inputs["zone"] = args ? args.zone : undefined;
            inputs["commitment"] = undefined /*out*/;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Reservation.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Reservation resources.
 */
export interface ReservationState {
    readonly commitment?: pulumi.Input<string>;
    readonly creationTimestamp?: pulumi.Input<string>;
    readonly description?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly project?: pulumi.Input<string>;
    /**
     * The URI of the created resource.
     */
    readonly selfLink?: pulumi.Input<string>;
    readonly specificReservation?: pulumi.Input<inputs.compute.ReservationSpecificReservation>;
    readonly specificReservationRequired?: pulumi.Input<boolean>;
    readonly status?: pulumi.Input<string>;
    readonly zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Reservation resource.
 */
export interface ReservationArgs {
    readonly description?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly project?: pulumi.Input<string>;
    readonly specificReservation: pulumi.Input<inputs.compute.ReservationSpecificReservation>;
    readonly specificReservationRequired?: pulumi.Input<boolean>;
    readonly zone: pulumi.Input<string>;
}