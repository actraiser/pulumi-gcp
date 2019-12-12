// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kms

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

//  Creates a device registry in Google's Cloud IoT Core platform. For more information see
// [the official documentation](https://cloud.google.com/iot/docs/) and
// [API](https://cloud.google.com/iot/docs/reference/cloudiot/rest/v1/projects.locations.registries).
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/cloudiot_registry.html.markdown.
type Registry struct {
	s *pulumi.ResourceState
}

// NewRegistry registers a new resource with the given unique name, arguments, and options.
func NewRegistry(ctx *pulumi.Context,
	name string, args *RegistryArgs, opts ...pulumi.ResourceOpt) (*Registry, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["credentials"] = nil
		inputs["eventNotificationConfigs"] = nil
		inputs["httpConfig"] = nil
		inputs["logLevel"] = nil
		inputs["mqttConfig"] = nil
		inputs["name"] = nil
		inputs["project"] = nil
		inputs["region"] = nil
		inputs["stateNotificationConfig"] = nil
	} else {
		inputs["credentials"] = args.Credentials
		inputs["eventNotificationConfigs"] = args.EventNotificationConfigs
		inputs["httpConfig"] = args.HttpConfig
		inputs["logLevel"] = args.LogLevel
		inputs["mqttConfig"] = args.MqttConfig
		inputs["name"] = args.Name
		inputs["project"] = args.Project
		inputs["region"] = args.Region
		inputs["stateNotificationConfig"] = args.StateNotificationConfig
	}
	s, err := ctx.RegisterResource("gcp:kms/registry:Registry", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Registry{s: s}, nil
}

// GetRegistry gets an existing Registry resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegistry(ctx *pulumi.Context,
	name string, id pulumi.ID, state *RegistryState, opts ...pulumi.ResourceOpt) (*Registry, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["credentials"] = state.Credentials
		inputs["eventNotificationConfigs"] = state.EventNotificationConfigs
		inputs["httpConfig"] = state.HttpConfig
		inputs["logLevel"] = state.LogLevel
		inputs["mqttConfig"] = state.MqttConfig
		inputs["name"] = state.Name
		inputs["project"] = state.Project
		inputs["region"] = state.Region
		inputs["stateNotificationConfig"] = state.StateNotificationConfig
	}
	s, err := ctx.ReadResource("gcp:kms/registry:Registry", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Registry{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Registry) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Registry) ID() pulumi.IDOutput {
	return r.s.ID()
}

// List of public key certificates to authenticate devices. Structure is documented below. 
func (r *Registry) Credentials() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["credentials"])
}

// List of configurations for event notification, such as
// PubSub topics to publish device events to. Structure is documented below.
func (r *Registry) EventNotificationConfigs() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["eventNotificationConfigs"])
}

// Activate or deactivate HTTP. Structure is documented below.
func (r *Registry) HttpConfig() pulumi.Output {
	return r.s.State["httpConfig"]
}

func (r *Registry) LogLevel() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["logLevel"])
}

// Activate or deactivate MQTT. Structure is documented below.
func (r *Registry) MqttConfig() pulumi.Output {
	return r.s.State["mqttConfig"]
}

// A unique name for the resource, required by device registry.
// Changing this forces a new resource to be created.
func (r *Registry) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// The project in which the resource belongs. If it is not provided, the provider project is used.
func (r *Registry) Project() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["project"])
}

// The Region in which the created address should reside. If it is not provided, the provider region is used.
func (r *Registry) Region() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["region"])
}

// A PubSub topic to publish device state updates. Structure is documented below.
func (r *Registry) StateNotificationConfig() pulumi.Output {
	return r.s.State["stateNotificationConfig"]
}

// Input properties used for looking up and filtering Registry resources.
type RegistryState struct {
	// List of public key certificates to authenticate devices. Structure is documented below. 
	Credentials interface{}
	// List of configurations for event notification, such as
	// PubSub topics to publish device events to. Structure is documented below.
	EventNotificationConfigs interface{}
	// Activate or deactivate HTTP. Structure is documented below.
	HttpConfig interface{}
	LogLevel interface{}
	// Activate or deactivate MQTT. Structure is documented below.
	MqttConfig interface{}
	// A unique name for the resource, required by device registry.
	// Changing this forces a new resource to be created.
	Name interface{}
	// The project in which the resource belongs. If it is not provided, the provider project is used.
	Project interface{}
	// The Region in which the created address should reside. If it is not provided, the provider region is used.
	Region interface{}
	// A PubSub topic to publish device state updates. Structure is documented below.
	StateNotificationConfig interface{}
}

// The set of arguments for constructing a Registry resource.
type RegistryArgs struct {
	// List of public key certificates to authenticate devices. Structure is documented below. 
	Credentials interface{}
	// List of configurations for event notification, such as
	// PubSub topics to publish device events to. Structure is documented below.
	EventNotificationConfigs interface{}
	// Activate or deactivate HTTP. Structure is documented below.
	HttpConfig interface{}
	LogLevel interface{}
	// Activate or deactivate MQTT. Structure is documented below.
	MqttConfig interface{}
	// A unique name for the resource, required by device registry.
	// Changing this forces a new resource to be created.
	Name interface{}
	// The project in which the resource belongs. If it is not provided, the provider project is used.
	Project interface{}
	// The Region in which the created address should reside. If it is not provided, the provider region is used.
	Region interface{}
	// A PubSub topic to publish device state updates. Structure is documented below.
	StateNotificationConfig interface{}
}
