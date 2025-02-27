// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/UsageContext
type UsageContext struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Type of context being specified
	Code Coding `json:"code"`
	// Value that defines the context
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty"`
	// Value that defines the context
	ValueQuantity *Quantity `json:"valueQuantity,omitempty"`
	// Value that defines the context
	ValueRange *Range `json:"valueRange,omitempty"`
	// Value that defines the context
	ValueReference *custom.Reference[UsageContextValueReference] `json:"valueReference,omitempty"`
}

type OtherUsageContext UsageContext

type UsageContextValueReference interface {
	gofhirclient.Resource

	Is_UsageContextValueReference()
}

func (p PlanDefinition) Is_UsageContextValueReference()    {}
func (r ResearchStudy) Is_UsageContextValueReference()     {}
func (i InsurancePlan) Is_UsageContextValueReference()     {}
func (h HealthcareService) Is_UsageContextValueReference() {}
func (g Group) Is_UsageContextValueReference()             {}
func (l Location) Is_UsageContextValueReference()          {}
func (o Organization) Is_UsageContextValueReference()      {}

func (u UsageContext) ResourceType() string {
	return "UsageContext"
}

func (u UsageContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherUsageContext
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherUsageContext: OtherUsageContext(u), ResourceType: u.ResourceType()})
}

func UnmarshalUsageContext(b []byte) (res UsageContext, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
