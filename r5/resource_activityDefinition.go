// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/ActivityDefinition
type ActivityDefinition struct {
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
	// Canonical identifier for this activity definition, represented as a URI (globally unique)
	Url *custom.URI `json:"url,omitempty"`
	// Additional identifier for the activity definition
	Identifier []Identifier `json:"identifier,omitempty"`
	// Business version of the activity definition
	Version *string `json:"version,omitempty"`
	// How to compare versions
	VersionAlgorithmString *string `json:"versionAlgorithmString,omitempty"`
	// How to compare versions
	VersionAlgorithmCoding *Coding `json:"versionAlgorithmCoding,omitempty"`
	// Name for this activity definition (computer friendly)
	Name *string `json:"name,omitempty"`
	// Name for this activity definition (human friendly)
	Title *string `json:"title,omitempty"`
	// Subordinate title of the activity definition
	Subtitle *string `json:"subtitle,omitempty"`
	// draft | active | retired | unknown
	Status custom.Code `json:"status"`
	// For testing purposes, not real usage
	Experimental *bool `json:"experimental,omitempty"`
	// Type of individual the activity definition is intended for
	SubjectCodeableConcept *CodeableConcept `json:"subjectCodeableConcept,omitempty"`
	// Type of individual the activity definition is intended for
	SubjectReference *custom.Reference[ActivityDefinitionSubjectReference] `json:"subjectReference,omitempty"`
	// Type of individual the activity definition is intended for
	SubjectCanonical *custom.Canonical[EvidenceVariable] `json:"subjectCanonical,omitempty"`
	// Date last changed
	Date *custom.DateTime `json:"date,omitempty"`
	// Name of the publisher/steward (organization or individual)
	Publisher *string `json:"publisher,omitempty"`
	// Contact details for the publisher
	Contact []ContactDetail `json:"contact,omitempty"`
	// Natural language description of the activity definition
	Description *custom.Markdown `json:"description,omitempty"`
	// The context that the content is intended to support
	UseContext []UsageContext `json:"useContext,omitempty"`
	// Intended jurisdiction for activity definition (if applicable)
	Jurisdiction []CodeableConcept `json:"jurisdiction,omitempty"`
	// Why this activity definition is defined
	Purpose *custom.Markdown `json:"purpose,omitempty"`
	// Describes the clinical usage of the activity definition
	Usage *custom.Markdown `json:"usage,omitempty"`
	// Use and/or publishing restrictions
	Copyright *custom.Markdown `json:"copyright,omitempty"`
	// Copyright holder and year(s)
	CopyrightLabel *string `json:"copyrightLabel,omitempty"`
	// When the activity definition was approved by publisher
	ApprovalDate *custom.Date `json:"approvalDate,omitempty"`
	// When the activity definition was last reviewed by the publisher
	LastReviewDate *custom.Date `json:"lastReviewDate,omitempty"`
	// When the activity definition is expected to be used
	EffectivePeriod *Period `json:"effectivePeriod,omitempty"`
	// E.g. Education, Treatment, Assessment, etc
	Topic []CodeableConcept `json:"topic,omitempty"`
	// Who authored the content
	Author []ContactDetail `json:"author,omitempty"`
	// Who edited the content
	Editor []ContactDetail `json:"editor,omitempty"`
	// Who reviewed the content
	Reviewer []ContactDetail `json:"reviewer,omitempty"`
	// Who endorsed the content
	Endorser []ContactDetail `json:"endorser,omitempty"`
	// Additional documentation, citations, etc
	RelatedArtifact []RelatedArtifact `json:"relatedArtifact,omitempty"`
	// Logic used by the activity definition
	Library []custom.Canonical[Library] `json:"library,omitempty"`
	// Kind of resource
	Kind *custom.Code `json:"kind,omitempty"`
	// What profile the resource needs to conform to
	Profile *custom.Canonical[StructureDefinition] `json:"profile,omitempty"`
	// Detail type of activity
	Code *CodeableConcept `json:"code,omitempty"`
	// proposal | plan | directive | order | original-order | reflex-order | filler-order | instance-order | option
	Intent *custom.Code `json:"intent,omitempty"`
	// routine | urgent | asap | stat
	Priority *custom.Code `json:"priority,omitempty"`
	// True if the activity should not be performed
	DoNotPerform *bool `json:"doNotPerform,omitempty"`
	// When activity is to occur
	TimingTiming *Timing `json:"timingTiming,omitempty"`
	// When activity is to occur
	TimingAge *Age `json:"timingAge,omitempty"`
	// When activity is to occur
	TimingRange *Range `json:"timingRange,omitempty"`
	// When activity is to occur
	TimingDuration *Duration `json:"timingDuration,omitempty"`
	// Preconditions for service
	AsNeededBoolean *bool `json:"asNeededBoolean,omitempty"`
	// Preconditions for service
	AsNeededCodeableConcept *CodeableConcept `json:"asNeededCodeableConcept,omitempty"`
	// Where it should happen
	Location *custom.CodeableReference[Location] `json:"location,omitempty"`
	// Who should participate in the action
	Participant []ActivityDefinitionParticipant `json:"participant,omitempty"`
	// What's administered/supplied
	ProductReference *custom.Reference[ActivityDefinitionProductReference] `json:"productReference,omitempty"`
	// What's administered/supplied
	ProductCodeableConcept *CodeableConcept `json:"productCodeableConcept,omitempty"`
	// How much is administered/consumed/supplied
	Quantity *Quantity `json:"quantity,omitempty"`
	// Detailed dosage instructions
	Dosage []Dosage `json:"dosage,omitempty"`
	// What part of body to perform on
	BodySite []CodeableConcept `json:"bodySite,omitempty"`
	// What specimens are required to perform this action
	SpecimenRequirement []custom.Canonical[SpecimenDefinition] `json:"specimenRequirement,omitempty"`
	// What observations are required to perform this action
	ObservationRequirement []custom.Canonical[ObservationDefinition] `json:"observationRequirement,omitempty"`
	// What observations must be produced by this action
	ObservationResultRequirement []custom.Canonical[ObservationDefinition] `json:"observationResultRequirement,omitempty"`
	// Transform to apply the template
	Transform *custom.Canonical[StructureMap] `json:"transform,omitempty"`
	// Dynamic aspects of the definition
	DynamicValue []ActivityDefinitionDynamicValue `json:"dynamicValue,omitempty"`
}

type ActivityDefinitionParticipant struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// careteam | device | group | healthcareservice | location | organization | patient | practitioner | practitionerrole | relatedperson
	Type *custom.Code `json:"type,omitempty"`
	// Who or what can participate
	TypeCanonical *custom.Canonical[CapabilityStatement] `json:"typeCanonical,omitempty"`
	// Who or what can participate
	TypeReference *custom.Reference[ActivityDefinitionParticipantTypeReference] `json:"typeReference,omitempty"`
	// E.g. Nurse, Surgeon, Parent, etc
	Role *CodeableConcept `json:"role,omitempty"`
	// E.g. Author, Reviewer, Witness, etc
	Function *CodeableConcept `json:"function,omitempty"`
}

type ActivityDefinitionDynamicValue struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The path to the element to be set dynamically
	Path string `json:"path"`
	// An expression that provides the dynamic value for the customization
	Expression Expression `json:"expression"`
}

type OtherActivityDefinition ActivityDefinition

type ActivityDefinitionSubjectReference interface {
	gofhirclient.Resource

	Is_ActivityDefinitionSubjectReference()
}

func (g Group) Is_ActivityDefinitionSubjectReference()                          {}
func (m MedicinalProductDefinition) Is_ActivityDefinitionSubjectReference()     {}
func (s SubstanceDefinition) Is_ActivityDefinitionSubjectReference()            {}
func (a AdministrableProductDefinition) Is_ActivityDefinitionSubjectReference() {}
func (m ManufacturedItemDefinition) Is_ActivityDefinitionSubjectReference()     {}
func (p PackagedProductDefinition) Is_ActivityDefinitionSubjectReference()      {}

type ActivityDefinitionParticipantTypeReference interface {
	gofhirclient.Resource

	Is_ActivityDefinitionParticipantTypeReference()
}

func (c CareTeam) Is_ActivityDefinitionParticipantTypeReference()          {}
func (d Device) Is_ActivityDefinitionParticipantTypeReference()            {}
func (d DeviceDefinition) Is_ActivityDefinitionParticipantTypeReference()  {}
func (e Endpoint) Is_ActivityDefinitionParticipantTypeReference()          {}
func (g Group) Is_ActivityDefinitionParticipantTypeReference()             {}
func (h HealthcareService) Is_ActivityDefinitionParticipantTypeReference() {}
func (l Location) Is_ActivityDefinitionParticipantTypeReference()          {}
func (o Organization) Is_ActivityDefinitionParticipantTypeReference()      {}
func (p Patient) Is_ActivityDefinitionParticipantTypeReference()           {}
func (p Practitioner) Is_ActivityDefinitionParticipantTypeReference()      {}
func (p PractitionerRole) Is_ActivityDefinitionParticipantTypeReference()  {}
func (r RelatedPerson) Is_ActivityDefinitionParticipantTypeReference()     {}

type ActivityDefinitionProductReference interface {
	gofhirclient.Resource

	Is_ActivityDefinitionProductReference()
}

func (m Medication) Is_ActivityDefinitionProductReference()          {}
func (i Ingredient) Is_ActivityDefinitionProductReference()          {}
func (s Substance) Is_ActivityDefinitionProductReference()           {}
func (s SubstanceDefinition) Is_ActivityDefinitionProductReference() {}

func (a ActivityDefinition) ResourceType() string {
	return "ActivityDefinition"
}

func (a ActivityDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherActivityDefinition
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherActivityDefinition: OtherActivityDefinition(a), ResourceType: a.ResourceType()})
}

func UnmarshalActivityDefinition(b []byte) (res ActivityDefinition, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
