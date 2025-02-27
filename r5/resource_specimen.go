// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Specimen
type Specimen struct {
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
	// External Identifier
	Identifier []Identifier `json:"identifier,omitempty"`
	// Identifier assigned by the lab
	AccessionIdentifier *Identifier `json:"accessionIdentifier,omitempty"`
	// available | unavailable | unsatisfactory | entered-in-error
	Status *custom.Code `json:"status,omitempty"`
	// Kind of material that forms the specimen
	Type *CodeableConcept `json:"type,omitempty"`
	// Where the specimen came from. This may be from patient(s), from a location (e.g., the source of an environmental sample), or a sampling of a substance, a biologically-derived product, or a device
	Subject *custom.Reference[SpecimenSubject] `json:"subject,omitempty"`
	// The time when specimen is received by the testing laboratory
	ReceivedTime *custom.DateTime `json:"receivedTime,omitempty"`
	// Specimen from which this specimen originated
	Parent []custom.Reference[Specimen] `json:"parent,omitempty"`
	// Why the specimen was collected
	Request []custom.Reference[ServiceRequest] `json:"request,omitempty"`
	// grouped | pooled
	Combined *custom.Code `json:"combined,omitempty"`
	// The role the specimen serves
	Role []CodeableConcept `json:"role,omitempty"`
	// The physical feature of a specimen
	Feature []SpecimenFeature `json:"feature,omitempty"`
	// Collection details
	Collection *SpecimenCollection `json:"collection,omitempty"`
	// Processing and processing step details
	Processing []SpecimenProcessing `json:"processing,omitempty"`
	// Direct container of specimen (tube/slide, etc.)
	Container []SpecimenContainer `json:"container,omitempty"`
	// State of the specimen
	Condition []CodeableConcept `json:"condition,omitempty"`
	// Comments
	Note []Annotation `json:"note,omitempty"`
}

type SpecimenFeature struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Highlighted feature
	Type CodeableConcept `json:"type"`
	// Information about the feature
	Description string `json:"description"`
}

type SpecimenCollection struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Who collected the specimen
	Collector *custom.Reference[SpecimenCollectionCollector] `json:"collector,omitempty"`
	// Collection time
	CollectedDateTime *custom.DateTime `json:"collectedDateTime,omitempty"`
	// Collection time
	CollectedPeriod *Period `json:"collectedPeriod,omitempty"`
	// How long it took to collect specimen
	Duration *Duration `json:"duration,omitempty"`
	// The quantity of specimen collected
	Quantity *Quantity `json:"quantity,omitempty"`
	// Technique used to perform collection
	Method *CodeableConcept `json:"method,omitempty"`
	// Device used to perform collection
	Device *custom.CodeableReference[Device] `json:"device,omitempty"`
	// The procedure that collects the specimen
	Procedure *custom.Reference[Procedure] `json:"procedure,omitempty"`
	// Anatomical collection site
	BodySite *custom.CodeableReference[BodyStructure] `json:"bodySite,omitempty"`
	// Whether or how long patient abstained from food and/or drink
	FastingStatusCodeableConcept *CodeableConcept `json:"fastingStatusCodeableConcept,omitempty"`
	// Whether or how long patient abstained from food and/or drink
	FastingStatusDuration *Duration `json:"fastingStatusDuration,omitempty"`
}

type SpecimenProcessing struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Textual description of procedure
	Description *string `json:"description,omitempty"`
	// Indicates the treatment step  applied to the specimen
	Method *CodeableConcept `json:"method,omitempty"`
	// Material used in the processing step
	Additive []custom.Reference[Substance] `json:"additive,omitempty"`
	// Date and time of specimen processing
	TimeDateTime *custom.DateTime `json:"timeDateTime,omitempty"`
	// Date and time of specimen processing
	TimePeriod *Period `json:"timePeriod,omitempty"`
}

type SpecimenContainer struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Device resource for the container
	Device custom.Reference[Device] `json:"device"`
	// Where the container is
	Location *custom.Reference[Location] `json:"location,omitempty"`
	// Quantity of specimen within container
	SpecimenQuantity *Quantity `json:"specimenQuantity,omitempty"`
}

type OtherSpecimen Specimen

type SpecimenSubject interface {
	gofhirclient.Resource

	Is_SpecimenSubject()
}

func (p Patient) Is_SpecimenSubject()                    {}
func (g Group) Is_SpecimenSubject()                      {}
func (d Device) Is_SpecimenSubject()                     {}
func (b BiologicallyDerivedProduct) Is_SpecimenSubject() {}
func (s Substance) Is_SpecimenSubject()                  {}
func (l Location) Is_SpecimenSubject()                   {}

type SpecimenCollectionCollector interface {
	gofhirclient.Resource

	Is_SpecimenCollectionCollector()
}

func (p Practitioner) Is_SpecimenCollectionCollector()     {}
func (p PractitionerRole) Is_SpecimenCollectionCollector() {}
func (p Patient) Is_SpecimenCollectionCollector()          {}
func (r RelatedPerson) Is_SpecimenCollectionCollector()    {}

func (s Specimen) ResourceType() string {
	return "Specimen"
}

func (s Specimen) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSpecimen
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherSpecimen: OtherSpecimen(s), ResourceType: s.ResourceType()})
}

func UnmarshalSpecimen(b []byte) (res Specimen, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
