// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Specimen
type Specimen struct {
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
	// External Identifier
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Identifier assigned by the lab
	AccessionIdentifier *Identifier `bson:"accessionIdentifier,omitempty" json:"accessionIdentifier,omitempty"`
	// available | unavailable | unsatisfactory | entered-in-error
	Status *custom.Code `bson:"status,omitempty" json:"status,omitempty"`
	// Kind of material that forms the specimen
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// Where the specimen came from. This may be from patient(s), from a location (e.g., the source of an environmental sample), or a sampling of a substance, a biologically-derived product, or a device
	Subject *custom.Reference[SpecimenSubject] `bson:"subject,omitempty" json:"subject,omitempty"`
	// The time when specimen is received by the testing laboratory
	ReceivedTime *custom.DateTime `bson:"receivedTime,omitempty" json:"receivedTime,omitempty"`
	// Specimen from which this specimen originated
	Parent []custom.Reference[Specimen] `bson:"parent,omitempty" json:"parent,omitempty"`
	// Why the specimen was collected
	Request []custom.Reference[ServiceRequest] `bson:"request,omitempty" json:"request,omitempty"`
	// grouped | pooled
	Combined *custom.Code `bson:"combined,omitempty" json:"combined,omitempty"`
	// The role the specimen serves
	Role []CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
	// The physical feature of a specimen
	Feature []SpecimenFeature `bson:"feature,omitempty" json:"feature,omitempty"`
	// Collection details
	Collection *SpecimenCollection `bson:"collection,omitempty" json:"collection,omitempty"`
	// Processing and processing step details
	Processing []SpecimenProcessing `bson:"processing,omitempty" json:"processing,omitempty"`
	// Direct container of specimen (tube/slide, etc.)
	Container []SpecimenContainer `bson:"container,omitempty" json:"container,omitempty"`
	// State of the specimen
	Condition []CodeableConcept `bson:"condition,omitempty" json:"condition,omitempty"`
	// Comments
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty"`
}

type SpecimenFeature struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Highlighted feature
	Type CodeableConcept `bson:"type" json:"type"`
	// Information about the feature
	Description string `bson:"description" json:"description"`
}

type SpecimenCollection struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Who collected the specimen
	Collector *custom.Reference[SpecimenCollectionCollector] `bson:"collector,omitempty" json:"collector,omitempty"`
	// Collection time
	Collected *custom.DateTime `bson:"collected,omitempty" json:"collected,omitempty"`
	// How long it took to collect specimen
	Duration *Duration `bson:"duration,omitempty" json:"duration,omitempty"`
	// The quantity of specimen collected
	Quantity *Quantity `bson:"quantity,omitempty" json:"quantity,omitempty"`
	// Technique used to perform collection
	Method *CodeableConcept `bson:"method,omitempty" json:"method,omitempty"`
	// Device used to perform collection
	Device *custom.CodeableReference[Device] `bson:"device,omitempty" json:"device,omitempty"`
	// The procedure that collects the specimen
	Procedure *custom.Reference[Procedure] `bson:"procedure,omitempty" json:"procedure,omitempty"`
	// Anatomical collection site
	BodySite *custom.CodeableReference[BodyStructure] `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	// Whether or how long patient abstained from food and/or drink
	FastingStatus *CodeableConcept `bson:"fastingStatus,omitempty" json:"fastingStatus,omitempty"`
}

type SpecimenProcessing struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Textual description of procedure
	Description *string `bson:"description,omitempty" json:"description,omitempty"`
	// Indicates the treatment step  applied to the specimen
	Method *CodeableConcept `bson:"method,omitempty" json:"method,omitempty"`
	// Material used in the processing step
	Additive []custom.Reference[Substance] `bson:"additive,omitempty" json:"additive,omitempty"`
	// Date and time of specimen processing
	Time *custom.DateTime `bson:"time,omitempty" json:"time,omitempty"`
}

type SpecimenContainer struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Device resource for the container
	Device custom.Reference[Device] `bson:"device" json:"device"`
	// Where the container is
	Location *custom.Reference[Location] `bson:"location,omitempty" json:"location,omitempty"`
	// Quantity of specimen within container
	SpecimenQuantity *Quantity `bson:"specimenQuantity,omitempty" json:"specimenQuantity,omitempty"`
}

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
	return "StructureDefinition"
}
