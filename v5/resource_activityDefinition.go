// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/ActivityDefinition
type ActivityDefinitionParticipant struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// careteam | device | group | healthcareservice | location | organization | patient | practitioner | practitionerrole | relatedperson
	Type *custom.Code `bson:"type,omitempty" json:"type,omitempty"`
	// Who or what can participate
	TypeCanonical *custom.Canonical[CapabilityStatement] `bson:"typeCanonical,omitempty" json:"typeCanonical,omitempty"`
	// Who or what can participate
	TypeReference *custom.Reference[ActivityDefinitionParticipantTypeReference] `bson:"typeReference,omitempty" json:"typeReference,omitempty"`
	// E.g. Nurse, Surgeon, Parent, etc
	Role *CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
	// E.g. Author, Reviewer, Witness, etc
	Function *CodeableConcept `bson:"function,omitempty" json:"function,omitempty"`
}

type ActivityDefinitionDynamicValue struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The path to the element to be set dynamically
	Path string `bson:"path" json:"path"`
	// An expression that provides the dynamic value for the customization
	Expression Expression `bson:"expression" json:"expression"`
}

type ActivityDefinition struct {
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
	// Canonical identifier for this activity definition, represented as a URI (globally unique)
	Url *custom.URI `bson:"url,omitempty" json:"url,omitempty"`
	// Additional identifier for the activity definition
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Business version of the activity definition
	Version *string `bson:"version,omitempty" json:"version,omitempty"`
	// How to compare versions
	VersionAlgorithm *string `bson:"versionAlgorithm,omitempty" json:"versionAlgorithm,omitempty"`
	// Name for this activity definition (computer friendly)
	Name *string `bson:"name,omitempty" json:"name,omitempty"`
	// Name for this activity definition (human friendly)
	Title *string `bson:"title,omitempty" json:"title,omitempty"`
	// Subordinate title of the activity definition
	Subtitle *string `bson:"subtitle,omitempty" json:"subtitle,omitempty"`
	// draft | active | retired | unknown
	Status custom.Code `bson:"status" json:"status"`
	// For testing purposes, not real usage
	Experimental *bool `bson:"experimental,omitempty" json:"experimental,omitempty"`
	// Type of individual the activity definition is intended for
	Subject *CodeableConcept `bson:"subject,omitempty" json:"subject,omitempty"`
	// Date last changed
	Date *custom.DateTime `bson:"date,omitempty" json:"date,omitempty"`
	// Name of the publisher/steward (organization or individual)
	Publisher *string `bson:"publisher,omitempty" json:"publisher,omitempty"`
	// Contact details for the publisher
	Contact []ContactDetail `bson:"contact,omitempty" json:"contact,omitempty"`
	// Natural language description of the activity definition
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// The context that the content is intended to support
	UseContext []UsageContext `bson:"useContext,omitempty" json:"useContext,omitempty"`
	// Intended jurisdiction for activity definition (if applicable)
	Jurisdiction []CodeableConcept `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	// Why this activity definition is defined
	Purpose *custom.Markdown `bson:"purpose,omitempty" json:"purpose,omitempty"`
	// Describes the clinical usage of the activity definition
	Usage *custom.Markdown `bson:"usage,omitempty" json:"usage,omitempty"`
	// Use and/or publishing restrictions
	Copyright *custom.Markdown `bson:"copyright,omitempty" json:"copyright,omitempty"`
	// Copyright holder and year(s)
	CopyrightLabel *string `bson:"copyrightLabel,omitempty" json:"copyrightLabel,omitempty"`
	// When the activity definition was approved by publisher
	ApprovalDate *custom.Date `bson:"approvalDate,omitempty" json:"approvalDate,omitempty"`
	// When the activity definition was last reviewed by the publisher
	LastReviewDate *custom.Date `bson:"lastReviewDate,omitempty" json:"lastReviewDate,omitempty"`
	// When the activity definition is expected to be used
	EffectivePeriod *Period `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	// E.g. Education, Treatment, Assessment, etc
	Topic []CodeableConcept `bson:"topic,omitempty" json:"topic,omitempty"`
	// Who authored the content
	Author []ContactDetail `bson:"author,omitempty" json:"author,omitempty"`
	// Who edited the content
	Editor []ContactDetail `bson:"editor,omitempty" json:"editor,omitempty"`
	// Who reviewed the content
	Reviewer []ContactDetail `bson:"reviewer,omitempty" json:"reviewer,omitempty"`
	// Who endorsed the content
	Endorser []ContactDetail `bson:"endorser,omitempty" json:"endorser,omitempty"`
	// Additional documentation, citations, etc
	RelatedArtifact []RelatedArtifact `bson:"relatedArtifact,omitempty" json:"relatedArtifact,omitempty"`
	// Logic used by the activity definition
	Library []custom.Canonical[Library] `bson:"library,omitempty" json:"library,omitempty"`
	// Kind of resource
	Kind *custom.Code `bson:"kind,omitempty" json:"kind,omitempty"`
	// What profile the resource needs to conform to
	Profile *custom.Canonical[StructureDefinition] `bson:"profile,omitempty" json:"profile,omitempty"`
	// Detail type of activity
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	// proposal | plan | directive | order | original-order | reflex-order | filler-order | instance-order | option
	Intent *custom.Code `bson:"intent,omitempty" json:"intent,omitempty"`
	// routine | urgent | asap | stat
	Priority *custom.Code `bson:"priority,omitempty" json:"priority,omitempty"`
	// True if the activity should not be performed
	DoNotPerform *bool `bson:"doNotPerform,omitempty" json:"doNotPerform,omitempty"`
	// When activity is to occur
	Timing *Timing `bson:"timing,omitempty" json:"timing,omitempty"`
	// Preconditions for service
	AsNeeded *bool `bson:"asNeeded,omitempty" json:"asNeeded,omitempty"`
	// Where it should happen
	Location *custom.CodeableReference[Location] `bson:"location,omitempty" json:"location,omitempty"`
	// Who should participate in the action
	Participant []ActivityDefinitionParticipant `bson:"participant,omitempty" json:"participant,omitempty"`
	// What's administered/supplied
	Product *custom.Reference[ActivityDefinitionProduct] `bson:"product,omitempty" json:"product,omitempty"`
	// How much is administered/consumed/supplied
	Quantity *Quantity `bson:"quantity,omitempty" json:"quantity,omitempty"`
	// Detailed dosage instructions
	Dosage []Dosage `bson:"dosage,omitempty" json:"dosage,omitempty"`
	// What part of body to perform on
	BodySite []CodeableConcept `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	// What specimens are required to perform this action
	SpecimenRequirement []custom.Canonical[SpecimenDefinition] `bson:"specimenRequirement,omitempty" json:"specimenRequirement,omitempty"`
	// What observations are required to perform this action
	ObservationRequirement []custom.Canonical[ObservationDefinition] `bson:"observationRequirement,omitempty" json:"observationRequirement,omitempty"`
	// What observations must be produced by this action
	ObservationResultRequirement []custom.Canonical[ObservationDefinition] `bson:"observationResultRequirement,omitempty" json:"observationResultRequirement,omitempty"`
	// Transform to apply the template
	Transform *custom.Canonical[StructureMap] `bson:"transform,omitempty" json:"transform,omitempty"`
	// Dynamic aspects of the definition
	DynamicValue []ActivityDefinitionDynamicValue `bson:"dynamicValue,omitempty" json:"dynamicValue,omitempty"`
}

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

type ActivityDefinitionProduct interface {
	gofhirclient.Resource

	Is_ActivityDefinitionProduct()
}

func (m Medication) Is_ActivityDefinitionProduct()          {}
func (i Ingredient) Is_ActivityDefinitionProduct()          {}
func (s Substance) Is_ActivityDefinitionProduct()           {}
func (s SubstanceDefinition) Is_ActivityDefinitionProduct() {}

func (a ActivityDefinition) ResourceType() string {
	return "StructureDefinition"
}
