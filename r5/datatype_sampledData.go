// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/SampledData
type SampledData struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Zero value and units
	Origin Quantity `json:"origin"`
	// Number of intervalUnits between samples
	Interval *json.Number `json:"interval,omitempty"`
	// The measurement unit of the interval between samples
	IntervalUnit custom.Code `json:"intervalUnit"`
	// Multiply data by this before adding to origin
	Factor *json.Number `json:"factor,omitempty"`
	// Lower limit of detection
	LowerLimit *json.Number `json:"lowerLimit,omitempty"`
	// Upper limit of detection
	UpperLimit *json.Number `json:"upperLimit,omitempty"`
	// Number of sample points at each time point
	Dimensions uint32 `json:"dimensions"`
	// Defines the codes used in the data
	CodeMap *custom.Canonical[ConceptMap] `json:"codeMap,omitempty"`
	// Offsets, typically in time, at which data values were taken
	Offsets *string `json:"offsets,omitempty"`
	// Decimal values with spaces, or "E" | "U" | "L", or another code
	Data *string `json:"data,omitempty"`
}

type OtherSampledData SampledData

func (s SampledData) ResourceType() string {
	return "SampledData"
}

func (s SampledData) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSampledData
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherSampledData: OtherSampledData(s), ResourceType: s.ResourceType()})
}

func UnmarshalSampledData(b []byte) (res SampledData, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
