// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/ChargeItem
type ChargeItem struct {
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
	// Business Identifier for item
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Defining information about the code of this charge item
	DefinitionUri []custom.URI `bson:"definitionUri,omitempty" json:"definitionUri,omitempty"`
	// Resource defining the code of this ChargeItem
	DefinitionCanonical []custom.Canonical[ChargeItemDefinition] `bson:"definitionCanonical,omitempty" json:"definitionCanonical,omitempty"`
	// planned | billable | not-billable | aborted | billed | entered-in-error | unknown
	Status custom.Code `bson:"status" json:"status"`
	// Part of referenced ChargeItem
	PartOf []custom.Reference[ChargeItem] `bson:"partOf,omitempty" json:"partOf,omitempty"`
	// A code that identifies the charge, like a billing code
	Code CodeableConcept `bson:"code" json:"code"`
	// Individual service was done for/to
	Subject custom.Reference[ChargeItemSubject] `bson:"subject" json:"subject"`
	// Encounter associated with this ChargeItem
	Encounter *custom.Reference[Encounter] `bson:"encounter,omitempty" json:"encounter,omitempty"`
	// When the charged service was applied
	Occurrence *custom.DateTime `bson:"occurrence,omitempty" json:"occurrence,omitempty"`
	// Who performed charged service
	Performer []ChargeItemPerformer `bson:"performer,omitempty" json:"performer,omitempty"`
	// Organization providing the charged service
	PerformingOrganization *custom.Reference[Organization] `bson:"performingOrganization,omitempty" json:"performingOrganization,omitempty"`
	// Organization requesting the charged service
	RequestingOrganization *custom.Reference[Organization] `bson:"requestingOrganization,omitempty" json:"requestingOrganization,omitempty"`
	// Organization that has ownership of the (potential, future) revenue
	CostCenter *custom.Reference[Organization] `bson:"costCenter,omitempty" json:"costCenter,omitempty"`
	// Quantity of which the charge item has been serviced
	Quantity *Quantity `bson:"quantity,omitempty" json:"quantity,omitempty"`
	// Anatomical location, if relevant
	Bodysite []CodeableConcept `bson:"bodysite,omitempty" json:"bodysite,omitempty"`
	// Unit price overriding the associated rules
	UnitPriceComponent *MonetaryComponent `bson:"unitPriceComponent,omitempty" json:"unitPriceComponent,omitempty"`
	// Total price overriding the associated rules
	TotalPriceComponent *MonetaryComponent `bson:"totalPriceComponent,omitempty" json:"totalPriceComponent,omitempty"`
	// Reason for overriding the list price/factor
	OverrideReason *CodeableConcept `bson:"overrideReason,omitempty" json:"overrideReason,omitempty"`
	// Individual who was entering
	Enterer *custom.Reference[ChargeItemEnterer] `bson:"enterer,omitempty" json:"enterer,omitempty"`
	// Date the charge item was entered
	EnteredDate *custom.DateTime `bson:"enteredDate,omitempty" json:"enteredDate,omitempty"`
	// Why was the charged  service rendered?
	Reason []CodeableConcept `bson:"reason,omitempty" json:"reason,omitempty"`
	// Which rendered service is being charged?
	Service []custom.CodeableReference[ChargeItemService] `bson:"service,omitempty" json:"service,omitempty"`
	// Product charged
	Product []custom.CodeableReference[ChargeItemProduct] `bson:"product,omitempty" json:"product,omitempty"`
	// Account to place this charge
	Account []custom.Reference[Account] `bson:"account,omitempty" json:"account,omitempty"`
	// Comments made about the ChargeItem
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty"`
	// Further information supporting this charge
	SupportingInformation []custom.Reference[Resource] `bson:"supportingInformation,omitempty" json:"supportingInformation,omitempty"`
}

type ChargeItemPerformer struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// What type of performance was done
	Function *CodeableConcept `bson:"function,omitempty" json:"function,omitempty"`
	// Individual who was performing
	Actor custom.Reference[ChargeItemPerformerActor] `bson:"actor" json:"actor"`
}

type ChargeItemSubject interface {
	gofhirclient.Resource

	Is_ChargeItemSubject()
}

func (p Patient) Is_ChargeItemSubject() {}
func (g Group) Is_ChargeItemSubject()   {}

type ChargeItemPerformerActor interface {
	gofhirclient.Resource

	Is_ChargeItemPerformerActor()
}

func (p Practitioner) Is_ChargeItemPerformerActor()      {}
func (p PractitionerRole) Is_ChargeItemPerformerActor()  {}
func (o Organization) Is_ChargeItemPerformerActor()      {}
func (h HealthcareService) Is_ChargeItemPerformerActor() {}
func (c CareTeam) Is_ChargeItemPerformerActor()          {}
func (p Patient) Is_ChargeItemPerformerActor()           {}
func (d Device) Is_ChargeItemPerformerActor()            {}
func (r RelatedPerson) Is_ChargeItemPerformerActor()     {}

type ChargeItemEnterer interface {
	gofhirclient.Resource

	Is_ChargeItemEnterer()
}

func (p Practitioner) Is_ChargeItemEnterer()     {}
func (p PractitionerRole) Is_ChargeItemEnterer() {}
func (o Organization) Is_ChargeItemEnterer()     {}
func (p Patient) Is_ChargeItemEnterer()          {}
func (d Device) Is_ChargeItemEnterer()           {}
func (r RelatedPerson) Is_ChargeItemEnterer()    {}

type ChargeItemService interface {
	gofhirclient.Resource

	Is_ChargeItemService()
}

func (d DiagnosticReport) Is_ChargeItemService()         {}
func (i ImagingStudy) Is_ChargeItemService()             {}
func (i Immunization) Is_ChargeItemService()             {}
func (m MedicationAdministration) Is_ChargeItemService() {}
func (m MedicationDispense) Is_ChargeItemService()       {}
func (m MedicationRequest) Is_ChargeItemService()        {}
func (o Observation) Is_ChargeItemService()              {}
func (p Procedure) Is_ChargeItemService()                {}
func (s ServiceRequest) Is_ChargeItemService()           {}
func (s SupplyDelivery) Is_ChargeItemService()           {}

type ChargeItemProduct interface {
	gofhirclient.Resource

	Is_ChargeItemProduct()
}

func (d Device) Is_ChargeItemProduct()     {}
func (m Medication) Is_ChargeItemProduct() {}
func (s Substance) Is_ChargeItemProduct()  {}

func (c ChargeItem) ResourceType() string {
	return "StructureDefinition"
}
