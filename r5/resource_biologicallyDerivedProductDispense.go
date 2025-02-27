// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/BiologicallyDerivedProductDispense
type BiologicallyDerivedProductDispense struct {
	// Logical id of this artifact
	Id *string `json:"id,omitempty"`
	// Metadata about the resource
	Meta *Meta `json:"meta,omitempty"`
	// A set of rules under which this content was created
	ImplicitRules *custom.URI `json:"implicitRules,omitempty"`
	// Language of the resource content
	Language *custom.Code `json:"language,omitempty"`
	// Text summary of the resource, for human interpretation
	Text *Narrative `json:"text,omitempty"`
	// Contained, inline Resources
	Contained []Resource `json:"contained,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Business identifier for this dispense
	Identifier []Identifier `json:"identifier,omitempty"`
	// The order or request that this dispense is fulfilling
	BasedOn []custom.Reference[ServiceRequest] `json:"basedOn,omitempty"`
	// Short description
	PartOf []custom.Reference[BiologicallyDerivedProductDispense] `json:"partOf,omitempty"`
	// preparation | in-progress | allocated | issued | unfulfilled | returned | entered-in-error | unknown
	Status custom.Code `json:"status"`
	// Relationship between the donor and intended recipient
	OriginRelationshipType *CodeableConcept `json:"originRelationshipType,omitempty"`
	// The BiologicallyDerivedProduct that is dispensed
	Product custom.Reference[BiologicallyDerivedProduct] `json:"product"`
	// The intended recipient of the dispensed product
	Patient custom.Reference[Patient] `json:"patient"`
	// Indicates the type of matching associated with the dispense
	MatchStatus *CodeableConcept `json:"matchStatus,omitempty"`
	// Indicates who or what performed an action
	Performer []BiologicallyDerivedProductDispensePerformer `json:"performer,omitempty"`
	// Where the dispense occurred
	Location *custom.Reference[Location] `json:"location,omitempty"`
	// Amount dispensed
	Quantity *Quantity `json:"quantity,omitempty"`
	// When product was selected/matched
	PreparedDate *custom.DateTime `json:"preparedDate,omitempty"`
	// When the product was dispatched
	WhenHandedOver *custom.DateTime `json:"whenHandedOver,omitempty"`
	// Where the product was dispatched to
	Destination *custom.Reference[Location] `json:"destination,omitempty"`
	// Additional notes
	Note []Annotation `json:"note,omitempty"`
	// Specific instructions for use
	UsageInstruction *string `json:"usageInstruction,omitempty"`
}

type BiologicallyDerivedProductDispensePerformer struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Identifies the function of the performer during the dispense
	Function *CodeableConcept `json:"function,omitempty"`
	// Who performed the action
	Actor custom.Reference[Practitioner] `json:"actor"`
}

type OtherBiologicallyDerivedProductDispense BiologicallyDerivedProductDispense

func (b BiologicallyDerivedProductDispense) ResourceType() string {
	return "BiologicallyDerivedProductDispense"
}

func (b BiologicallyDerivedProductDispense) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherBiologicallyDerivedProductDispense
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherBiologicallyDerivedProductDispense: OtherBiologicallyDerivedProductDispense(b), ResourceType: b.ResourceType()})
}

func UnmarshalBiologicallyDerivedProductDispense(b []byte) (res BiologicallyDerivedProductDispense, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
