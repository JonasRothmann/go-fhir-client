// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/DataRequirement
type DataRequirement struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// The type of the required data
	Type custom.Code `json:"type"`
	// The profile of the required data
	Profile []custom.Canonical[StructureDefinition] `json:"profile,omitempty"`
	// E.g. Patient, Practitioner, RelatedPerson, Organization, Location, Device
	SubjectCodeableConcept *CodeableConcept `json:"subjectCodeableConcept,omitempty"`
	// E.g. Patient, Practitioner, RelatedPerson, Organization, Location, Device
	SubjectReference *custom.Reference[Group] `json:"subjectReference,omitempty"`
	// Indicates specific structure elements that are referenced by the knowledge module
	MustSupport []string `json:"mustSupport,omitempty"`
	// What codes are expected
	CodeFilter []DataRequirementCodeFilter `json:"codeFilter,omitempty"`
	// What dates/date ranges are expected
	DateFilter []DataRequirementDateFilter `json:"dateFilter,omitempty"`
	// What values are expected
	ValueFilter []DataRequirementValueFilter `json:"valueFilter,omitempty"`
	// Number of results
	Limit *uint32 `json:"limit,omitempty"`
	// Order of the results
	Sort []DataRequirementSort `json:"sort,omitempty"`
}

type DataRequirementCodeFilter struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// A code-valued attribute to filter on
	Path *string `json:"path,omitempty"`
	// A coded (token) parameter to search on
	SearchParam *string `json:"searchParam,omitempty"`
	// ValueSet for the filter
	ValueSet *custom.Canonical[ValueSet] `json:"valueSet,omitempty"`
	// What code is expected
	Code []Coding `json:"code,omitempty"`
}

type DataRequirementDateFilter struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// A date-valued attribute to filter on
	Path *string `json:"path,omitempty"`
	// A date valued parameter to search on
	SearchParam *string `json:"searchParam,omitempty"`
	// The value of the filter, as a Period, DateTime, or Duration value
	ValueDateTime *custom.DateTime `json:"valueDateTime,omitempty"`
	// The value of the filter, as a Period, DateTime, or Duration value
	ValuePeriod *Period `json:"valuePeriod,omitempty"`
	// The value of the filter, as a Period, DateTime, or Duration value
	ValueDuration *Duration `json:"valueDuration,omitempty"`
}

type DataRequirementValueFilter struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// An attribute to filter on
	Path *string `json:"path,omitempty"`
	// A parameter to search on
	SearchParam *string `json:"searchParam,omitempty"`
	// eq | gt | lt | ge | le | sa | eb
	Comparator *custom.Code `json:"comparator,omitempty"`
	// The value of the filter, as a Period, DateTime, or Duration value
	ValueDateTime *custom.DateTime `json:"valueDateTime,omitempty"`
	// The value of the filter, as a Period, DateTime, or Duration value
	ValuePeriod *Period `json:"valuePeriod,omitempty"`
	// The value of the filter, as a Period, DateTime, or Duration value
	ValueDuration *Duration `json:"valueDuration,omitempty"`
}

type DataRequirementSort struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// The name of the attribute to perform the sort
	Path string `json:"path"`
	// ascending | descending
	Direction custom.Code `json:"direction"`
}

type OtherDataRequirement DataRequirement

func (d DataRequirement) ResourceType() string {
	return "DataRequirement"
}

func (d DataRequirement) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherDataRequirement
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherDataRequirement: OtherDataRequirement(d), ResourceType: d.ResourceType()})
}

func UnmarshalDataRequirement(b []byte) (res DataRequirement, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
