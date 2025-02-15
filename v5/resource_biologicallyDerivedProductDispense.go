// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/BiologicallyDerivedProductDispense
type BiologicallyDerivedProductDispense struct {
	// Logical id of this artifact
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Metadata about the resource
	Meta *Meta `bson:"meta,omitempty" json:"meta,omitempty"`
	// A set of rules under which this content was created
	ImplicitRules *custom.URI `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	// Language of the resource content
	Language *custom.Code `bson:"language,omitempty" json:"language,omitempty"`
	// Text summary of the resource, for human interpretation
	Text *Narrative `bson:"text,omitempty" json:"text,omitempty"`
	// Contained, inline Resources
	Contained []Resource `bson:"contained,omitempty" json:"contained,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Business identifier for this dispense
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// The order or request that this dispense is fulfilling
	BasedOn []custom.Reference[ServiceRequest] `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	// Short description
	PartOf []custom.Reference[BiologicallyDerivedProductDispense] `bson:"partOf,omitempty" json:"partOf,omitempty"`
	// preparation | in-progress | allocated | issued | unfulfilled | returned | entered-in-error | unknown
	Status custom.Code `bson:"status" json:"status"`
	// Relationship between the donor and intended recipient
	OriginRelationshipType *CodeableConcept `bson:"originRelationshipType,omitempty" json:"originRelationshipType,omitempty"`
	// The BiologicallyDerivedProduct that is dispensed
	Product custom.Reference[BiologicallyDerivedProduct] `bson:"product" json:"product"`
	// The intended recipient of the dispensed product
	Patient custom.Reference[Patient] `bson:"patient" json:"patient"`
	// Indicates the type of matching associated with the dispense
	MatchStatus *CodeableConcept `bson:"matchStatus,omitempty" json:"matchStatus,omitempty"`
	// Indicates who or what performed an action
	Performer []BiologicallyDerivedProductDispensePerformer `bson:"performer,omitempty" json:"performer,omitempty"`
	// Where the dispense occurred
	Location *custom.Reference[Location] `bson:"location,omitempty" json:"location,omitempty"`
	// Amount dispensed
	Quantity *Quantity `bson:"quantity,omitempty" json:"quantity,omitempty"`
	// When product was selected/matched
	PreparedDate *custom.DateTime `bson:"preparedDate,omitempty" json:"preparedDate,omitempty"`
	// When the product was dispatched
	WhenHandedOver *custom.DateTime `bson:"whenHandedOver,omitempty" json:"whenHandedOver,omitempty"`
	// Where the product was dispatched to
	Destination *custom.Reference[Location] `bson:"destination,omitempty" json:"destination,omitempty"`
	// Additional notes
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty"`
	// Specific instructions for use
	UsageInstruction *string `bson:"usageInstruction,omitempty" json:"usageInstruction,omitempty"`
}

type BiologicallyDerivedProductDispensePerformer struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Identifies the function of the performer during the dispense
	Function *CodeableConcept `bson:"function,omitempty" json:"function,omitempty"`
	// Who performed the action
	Actor custom.Reference[Practitioner] `bson:"actor" json:"actor"`
}

func (b BiologicallyDerivedProductDispense) ResourceType() string {
	return "StructureDefinition"
}
