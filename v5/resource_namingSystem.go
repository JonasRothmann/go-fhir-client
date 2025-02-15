// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/NamingSystem
type NamingSystemUniqueId struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// oid | uuid | uri | iri-stem | v2csmnemonic | other
	Type custom.Code `bson:"type" json:"type"`
	// The unique identifier
	Value string `bson:"value" json:"value"`
	// Is this the id that should be used for this type
	Preferred *bool `bson:"preferred,omitempty" json:"preferred,omitempty"`
	// Notes about identifier usage
	Comment *string `bson:"comment,omitempty" json:"comment,omitempty"`
	// When is identifier valid?
	Period *Period `bson:"period,omitempty" json:"period,omitempty"`
	// Whether the identifier is authoritative
	Authoritative *bool `bson:"authoritative,omitempty" json:"authoritative,omitempty"`
}

type NamingSystem struct {
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
	// Canonical identifier for this naming system, represented as a URI (globally unique)
	Url *custom.URI `bson:"url,omitempty" json:"url,omitempty"`
	// Additional identifier for the naming system (business identifier)
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Business version of the naming system
	Version *string `bson:"version,omitempty" json:"version,omitempty"`
	// How to compare versions
	VersionAlgorithm *string `bson:"versionAlgorithm,omitempty" json:"versionAlgorithm,omitempty"`
	// Name for this naming system (computer friendly)
	Name string `bson:"name" json:"name"`
	// Title for this naming system (human friendly)
	Title *string `bson:"title,omitempty" json:"title,omitempty"`
	// draft | active | retired | unknown
	Status custom.Code `bson:"status" json:"status"`
	// codesystem | identifier | root
	Kind custom.Code `bson:"kind" json:"kind"`
	// For testing purposes, not real usage
	Experimental *bool `bson:"experimental,omitempty" json:"experimental,omitempty"`
	// Date last changed
	Date custom.DateTime `bson:"date" json:"date"`
	// Name of the publisher/steward (organization or individual)
	Publisher *string `bson:"publisher,omitempty" json:"publisher,omitempty"`
	// Contact details for the publisher
	Contact []ContactDetail `bson:"contact,omitempty" json:"contact,omitempty"`
	// Who maintains system namespace?
	Responsible *string `bson:"responsible,omitempty" json:"responsible,omitempty"`
	// e.g. driver,  provider,  patient, bank etc
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// Natural language description of the naming system
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// The context that the content is intended to support
	UseContext []UsageContext `bson:"useContext,omitempty" json:"useContext,omitempty"`
	// Intended jurisdiction for naming system (if applicable)
	Jurisdiction []CodeableConcept `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	// Why this naming system is defined
	Purpose *custom.Markdown `bson:"purpose,omitempty" json:"purpose,omitempty"`
	// Use and/or publishing restrictions
	Copyright *custom.Markdown `bson:"copyright,omitempty" json:"copyright,omitempty"`
	// Copyright holder and year(s)
	CopyrightLabel *string `bson:"copyrightLabel,omitempty" json:"copyrightLabel,omitempty"`
	// When the NamingSystem was approved by publisher
	ApprovalDate *custom.Date `bson:"approvalDate,omitempty" json:"approvalDate,omitempty"`
	// When the NamingSystem was last reviewed by the publisher
	LastReviewDate *custom.Date `bson:"lastReviewDate,omitempty" json:"lastReviewDate,omitempty"`
	// When the NamingSystem is expected to be used
	EffectivePeriod *Period `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	// E.g. Education, Treatment, Assessment, etc
	Topic []CodeableConcept `bson:"topic,omitempty" json:"topic,omitempty"`
	// Who authored the CodeSystem
	Author []ContactDetail `bson:"author,omitempty" json:"author,omitempty"`
	// Who edited the NamingSystem
	Editor []ContactDetail `bson:"editor,omitempty" json:"editor,omitempty"`
	// Who reviewed the NamingSystem
	Reviewer []ContactDetail `bson:"reviewer,omitempty" json:"reviewer,omitempty"`
	// Who endorsed the NamingSystem
	Endorser []ContactDetail `bson:"endorser,omitempty" json:"endorser,omitempty"`
	// Additional documentation, citations, etc
	RelatedArtifact []RelatedArtifact `bson:"relatedArtifact,omitempty" json:"relatedArtifact,omitempty"`
	// How/where is it used
	Usage *string `bson:"usage,omitempty" json:"usage,omitempty"`
	// Unique identifiers used for system
	UniqueId []NamingSystemUniqueId `bson:"uniqueId" json:"uniqueId"`
}

func (n NamingSystem) ResourceType() string {
	return "StructureDefinition"
}
