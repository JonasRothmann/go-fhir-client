// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/DataRequirement
type DataRequirementSort struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// The name of the attribute to perform the sort
	Path string `bson:"path" json:"path"`
	// ascending | descending
	Direction custom.Code `bson:"direction" json:"direction"`
}

type DataRequirement struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// The type of the required data
	Type custom.Code `bson:"type" json:"type"`
	// The profile of the required data
	Profile []custom.Canonical[StructureDefinition] `bson:"profile,omitempty" json:"profile,omitempty"`
	// E.g. Patient, Practitioner, RelatedPerson, Organization, Location, Device
	Subject *CodeableConcept `bson:"subject,omitempty" json:"subject,omitempty"`
	// Indicates specific structure elements that are referenced by the knowledge module
	MustSupport []string `bson:"mustSupport,omitempty" json:"mustSupport,omitempty"`
	// What codes are expected
	CodeFilter []DataRequirementCodeFilter `bson:"codeFilter,omitempty" json:"codeFilter,omitempty"`
	// What dates/date ranges are expected
	DateFilter []DataRequirementDateFilter `bson:"dateFilter,omitempty" json:"dateFilter,omitempty"`
	// What values are expected
	ValueFilter []DataRequirementValueFilter `bson:"valueFilter,omitempty" json:"valueFilter,omitempty"`
	// Number of results
	Limit *uint32 `bson:"limit,omitempty" json:"limit,omitempty"`
	// Order of the results
	Sort []DataRequirementSort `bson:"sort,omitempty" json:"sort,omitempty"`
}

type DataRequirementCodeFilter struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// A code-valued attribute to filter on
	Path *string `bson:"path,omitempty" json:"path,omitempty"`
	// A coded (token) parameter to search on
	SearchParam *string `bson:"searchParam,omitempty" json:"searchParam,omitempty"`
	// ValueSet for the filter
	ValueSet *custom.Canonical[ValueSet] `bson:"valueSet,omitempty" json:"valueSet,omitempty"`
	// What code is expected
	Code []Coding `bson:"code,omitempty" json:"code,omitempty"`
}

type DataRequirementDateFilter struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// A date-valued attribute to filter on
	Path *string `bson:"path,omitempty" json:"path,omitempty"`
	// A date valued parameter to search on
	SearchParam *string `bson:"searchParam,omitempty" json:"searchParam,omitempty"`
	// The value of the filter, as a Period, DateTime, or Duration value
	Value *custom.DateTime `bson:"value,omitempty" json:"value,omitempty"`
}

type DataRequirementValueFilter struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// An attribute to filter on
	Path *string `bson:"path,omitempty" json:"path,omitempty"`
	// A parameter to search on
	SearchParam *string `bson:"searchParam,omitempty" json:"searchParam,omitempty"`
	// eq | gt | lt | ge | le | sa | eb
	Comparator *custom.Code `bson:"comparator,omitempty" json:"comparator,omitempty"`
	// The value of the filter, as a Period, DateTime, or Duration value
	Value *custom.DateTime `bson:"value,omitempty" json:"value,omitempty"`
}

func (d DataRequirement) ResourceType() string {
	return "StructureDefinition"
}
