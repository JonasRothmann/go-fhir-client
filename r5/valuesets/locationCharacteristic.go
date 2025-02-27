// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/location-characteristic
*/type LocationCharacteristic string

const (
	// The location is accessible to thosre requiring wheelchair access
	LocationCharacteristicWheelchair LocationCharacteristic = "wheelchair"
	// The location has translation services available
	LocationCharacteristicTranslationServicesAvailable LocationCharacteristic = "has-translation"
	// The location has oxygen and nitrogen services available
	LocationCharacteristicOxygenNitrogenAvailable LocationCharacteristic = "has-oxy-nitro"
	// The location has negative pressure rooms available
	LocationCharacteristicNegativePressureRoomsAvailable LocationCharacteristic = "has-neg-press"
	// The location is or has an isolation ward
	LocationCharacteristicIsolationWard LocationCharacteristic = "has-iso-ward"
	// The location is or has an intensive care unit
	LocationCharacteristicHasIcu LocationCharacteristic = "has-icu"
)
