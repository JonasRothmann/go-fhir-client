// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/permitted-data-type
*/type ObservationDataType string

const (
	// A measured amount.
	ObservationDataTypeQuantity ObservationDataType = "Quantity"
	// A coded concept from a reference terminology and/or text.
	ObservationDataTypeCodeableConcept ObservationDataType = "CodeableConcept"
	// A sequence of Unicode characters.
	ObservationDataTypeString ObservationDataType = "string"
	// true or false.
	ObservationDataTypeBoolean ObservationDataType = "boolean"
	// A signed integer.
	ObservationDataTypeInteger ObservationDataType = "integer"
	// A set of values bounded by low and high.
	ObservationDataTypeRange ObservationDataType = "Range"
	// A ratio of two Quantity values - a numerator and a denominator.
	ObservationDataTypeRatio ObservationDataType = "Ratio"
	// A series of measurements taken by a device.
	ObservationDataTypeSampledData ObservationDataType = "SampledData"
	// A time during the day, in the format hh:mm:ss.
	ObservationDataTypeTime ObservationDataType = "time"
	// A date, date-time or partial date (e.g. just year or year + month) as used in human communication.
	ObservationDataTypeDateTime ObservationDataType = "dateTime"
	// A time range defined by start and end date/time.
	ObservationDataTypePeriod ObservationDataType = "Period"
)
