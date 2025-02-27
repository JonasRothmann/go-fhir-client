// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/devicedefinition-relationtype
*/type DeviceDefinitionRelationType string

const (
	// Gateway.
	DeviceDefinitionRelationTypeGateway DeviceDefinitionRelationType = "gateway"
	// The current device replaces the linked device.
	DeviceDefinitionRelationTypeReplaces DeviceDefinitionRelationType = "replaces"
	// The current device is a previous device and has been replaced by the linked device.
	DeviceDefinitionRelationTypePrevious DeviceDefinitionRelationType = "previous"
)
