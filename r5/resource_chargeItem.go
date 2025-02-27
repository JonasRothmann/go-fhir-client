// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/ChargeItem
type ChargeItem struct {
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
	// Business Identifier for item
	Identifier []Identifier `json:"identifier,omitempty"`
	// Defining information about the code of this charge item
	DefinitionUri []custom.URI `json:"definitionUri,omitempty"`
	// Resource defining the code of this ChargeItem
	DefinitionCanonical []custom.Canonical[ChargeItemDefinition] `json:"definitionCanonical,omitempty"`
	// planned | billable | not-billable | aborted | billed | entered-in-error | unknown
	Status custom.Code `json:"status"`
	// Part of referenced ChargeItem
	PartOf []custom.Reference[ChargeItem] `json:"partOf,omitempty"`
	// A code that identifies the charge, like a billing code
	Code CodeableConcept `json:"code"`
	// Individual service was done for/to
	Subject custom.Reference[ChargeItemSubject] `json:"subject"`
	// Encounter associated with this ChargeItem
	Encounter *custom.Reference[Encounter] `json:"encounter,omitempty"`
	// When the charged service was applied
	OccurrenceDateTime *custom.DateTime `json:"occurrenceDateTime,omitempty"`
	// When the charged service was applied
	OccurrencePeriod *Period `json:"occurrencePeriod,omitempty"`
	// When the charged service was applied
	OccurrenceTiming *Timing `json:"occurrenceTiming,omitempty"`
	// Who performed charged service
	Performer []ChargeItemPerformer `json:"performer,omitempty"`
	// Organization providing the charged service
	PerformingOrganization *custom.Reference[Organization] `json:"performingOrganization,omitempty"`
	// Organization requesting the charged service
	RequestingOrganization *custom.Reference[Organization] `json:"requestingOrganization,omitempty"`
	// Organization that has ownership of the (potential, future) revenue
	CostCenter *custom.Reference[Organization] `json:"costCenter,omitempty"`
	// Quantity of which the charge item has been serviced
	Quantity *Quantity `json:"quantity,omitempty"`
	// Anatomical location, if relevant
	Bodysite []CodeableConcept `json:"bodysite,omitempty"`
	// Unit price overriding the associated rules
	UnitPriceComponent *MonetaryComponent `json:"unitPriceComponent,omitempty"`
	// Total price overriding the associated rules
	TotalPriceComponent *MonetaryComponent `json:"totalPriceComponent,omitempty"`
	// Reason for overriding the list price/factor
	OverrideReason *CodeableConcept `json:"overrideReason,omitempty"`
	// Individual who was entering
	Enterer *custom.Reference[ChargeItemEnterer] `json:"enterer,omitempty"`
	// Date the charge item was entered
	EnteredDate *custom.DateTime `json:"enteredDate,omitempty"`
	// Why was the charged  service rendered?
	Reason []CodeableConcept `json:"reason,omitempty"`
	// Which rendered service is being charged?
	Service []custom.CodeableReference[ChargeItemService] `json:"service,omitempty"`
	// Product charged
	Product []custom.CodeableReference[ChargeItemProduct] `json:"product,omitempty"`
	// Account to place this charge
	Account []custom.Reference[Account] `json:"account,omitempty"`
	// Comments made about the ChargeItem
	Note []Annotation `json:"note,omitempty"`
	// Further information supporting this charge
	SupportingInformation []custom.Reference[Resource] `json:"supportingInformation,omitempty"`
}

type ChargeItemPerformer struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// What type of performance was done
	Function *CodeableConcept `json:"function,omitempty"`
	// Individual who was performing
	Actor custom.Reference[ChargeItemPerformerActor] `json:"actor"`
}

type OtherChargeItem ChargeItem

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
	return "ChargeItem"
}

func (c ChargeItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherChargeItem
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherChargeItem: OtherChargeItem(c), ResourceType: c.ResourceType()})
}

func UnmarshalChargeItem(b []byte) (res ChargeItem, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
