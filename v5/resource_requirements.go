// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Requirements
type Requirements struct {
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
	// Canonical identifier for this Requirements, represented as a URI (globally unique)
	Url *custom.URI `bson:"url,omitempty" json:"url,omitempty"`
	// Additional identifier for the Requirements (business identifier)
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Business version of the Requirements
	Version *string `bson:"version,omitempty" json:"version,omitempty"`
	// How to compare versions
	VersionAlgorithm *string `bson:"versionAlgorithm,omitempty" json:"versionAlgorithm,omitempty"`
	// Name for this Requirements (computer friendly)
	Name *string `bson:"name,omitempty" json:"name,omitempty"`
	// Name for this Requirements (human friendly)
	Title *string `bson:"title,omitempty" json:"title,omitempty"`
	// draft | active | retired | unknown
	Status custom.Code `bson:"status" json:"status"`
	// For testing purposes, not real usage
	Experimental *bool `bson:"experimental,omitempty" json:"experimental,omitempty"`
	// Date last changed
	Date *custom.DateTime `bson:"date,omitempty" json:"date,omitempty"`
	// Name of the publisher/steward (organization or individual)
	Publisher *string `bson:"publisher,omitempty" json:"publisher,omitempty"`
	// Contact details for the publisher
	Contact []ContactDetail `bson:"contact,omitempty" json:"contact,omitempty"`
	// Natural language description of the requirements
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// The context that the content is intended to support
	UseContext []UsageContext `bson:"useContext,omitempty" json:"useContext,omitempty"`
	// Intended jurisdiction for Requirements (if applicable)
	Jurisdiction []CodeableConcept `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	// Why this Requirements is defined
	Purpose *custom.Markdown `bson:"purpose,omitempty" json:"purpose,omitempty"`
	// Use and/or publishing restrictions
	Copyright *custom.Markdown `bson:"copyright,omitempty" json:"copyright,omitempty"`
	// Copyright holder and year(s)
	CopyrightLabel *string `bson:"copyrightLabel,omitempty" json:"copyrightLabel,omitempty"`
	// Other set of Requirements this builds on
	DerivedFrom []custom.Canonical[Requirements] `bson:"derivedFrom,omitempty" json:"derivedFrom,omitempty"`
	// External artifact (rule/document etc. that) created this set of requirements
	Reference []custom.URL `bson:"reference,omitempty" json:"reference,omitempty"`
	// Actor for these requirements
	Actor []custom.Canonical[ActorDefinition] `bson:"actor,omitempty" json:"actor,omitempty"`
	// Actual statement as markdown
	Statement []RequirementsStatement `bson:"statement,omitempty" json:"statement,omitempty"`
}

type RequirementsStatement struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Key that identifies this statement
	Key custom.ID `bson:"key" json:"key"`
	// Short Human label for this statement
	Label *string `bson:"label,omitempty" json:"label,omitempty"`
	// SHALL | SHOULD | MAY | SHOULD-NOT
	Conformance []custom.Code `bson:"conformance,omitempty" json:"conformance,omitempty"`
	// Set to true if requirements statement is conditional
	Conditionality *bool `bson:"conditionality,omitempty" json:"conditionality,omitempty"`
	// The actual requirement
	Requirement custom.Markdown `bson:"requirement" json:"requirement"`
	// Another statement this clarifies/restricts ([url#]key)
	DerivedFrom *string `bson:"derivedFrom,omitempty" json:"derivedFrom,omitempty"`
	// A larger requirement that this requirement helps to refine and enable
	Parent *string `bson:"parent,omitempty" json:"parent,omitempty"`
	// Design artifact that satisfies this requirement
	SatisfiedBy []custom.URL `bson:"satisfiedBy,omitempty" json:"satisfiedBy,omitempty"`
	// External artifact (rule/document etc. that) created this requirement
	Reference []custom.URL `bson:"reference,omitempty" json:"reference,omitempty"`
	// Who asked for this statement
	Source []custom.Reference[RequirementsStatementSource] `bson:"source,omitempty" json:"source,omitempty"`
}

type RequirementsStatementSource interface {
	gofhirclient.Resource

	Is_RequirementsStatementSource()
}

func (c CareTeam) Is_RequirementsStatementSource()          {}
func (d Device) Is_RequirementsStatementSource()            {}
func (g Group) Is_RequirementsStatementSource()             {}
func (h HealthcareService) Is_RequirementsStatementSource() {}
func (o Organization) Is_RequirementsStatementSource()      {}
func (p Patient) Is_RequirementsStatementSource()           {}
func (p Practitioner) Is_RequirementsStatementSource()      {}
func (p PractitionerRole) Is_RequirementsStatementSource()  {}
func (r RelatedPerson) Is_RequirementsStatementSource()     {}

func (r Requirements) ResourceType() string {
	return "StructureDefinition"
}
