// GENERATED CODE – DO NOT EDIT!

package v5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/SampledData
type SampledData struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Zero value and units
	Origin Quantity `bson:"origin" json:"origin"`
	// Number of intervalUnits between samples
	Interval *json.Number `bson:"interval,omitempty" json:"interval,omitempty"`
	// The measurement unit of the interval between samples
	IntervalUnit custom.Code `bson:"intervalUnit" json:"intervalUnit"`
	// Multiply data by this before adding to origin
	Factor *json.Number `bson:"factor,omitempty" json:"factor,omitempty"`
	// Lower limit of detection
	LowerLimit *json.Number `bson:"lowerLimit,omitempty" json:"lowerLimit,omitempty"`
	// Upper limit of detection
	UpperLimit *json.Number `bson:"upperLimit,omitempty" json:"upperLimit,omitempty"`
	// Number of sample points at each time point
	Dimensions uint32 `bson:"dimensions" json:"dimensions"`
	// Defines the codes used in the data
	CodeMap *custom.Canonical[ConceptMap] `bson:"codeMap,omitempty" json:"codeMap,omitempty"`
	// Offsets, typically in time, at which data values were taken
	Offsets *string `bson:"offsets,omitempty" json:"offsets,omitempty"`
	// Decimal values with spaces, or "E" | "U" | "L", or another code
	Data *string `bson:"data,omitempty" json:"data,omitempty"`
}

func (s SampledData) ResourceType() string {
	return "StructureDefinition"
}
