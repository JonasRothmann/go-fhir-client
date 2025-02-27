// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Requirements
type Requirements struct {
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
	// Canonical identifier for this Requirements, represented as a URI (globally unique)
	Url *custom.URI `json:"url,omitempty"`
	// Additional identifier for the Requirements (business identifier)
	Identifier []Identifier `json:"identifier,omitempty"`
	// Business version of the Requirements
	Version *string `json:"version,omitempty"`
	// How to compare versions
	VersionAlgorithmString *string `json:"versionAlgorithmString,omitempty"`
	// How to compare versions
	VersionAlgorithmCoding *Coding `json:"versionAlgorithmCoding,omitempty"`
	// Name for this Requirements (computer friendly)
	Name *string `json:"name,omitempty"`
	// Name for this Requirements (human friendly)
	Title *string `json:"title,omitempty"`
	// draft | active | retired | unknown
	Status custom.Code `json:"status"`
	// For testing purposes, not real usage
	Experimental *bool `json:"experimental,omitempty"`
	// Date last changed
	Date *custom.DateTime `json:"date,omitempty"`
	// Name of the publisher/steward (organization or individual)
	Publisher *string `json:"publisher,omitempty"`
	// Contact details for the publisher
	Contact []ContactDetail `json:"contact,omitempty"`
	// Natural language description of the requirements
	Description *custom.Markdown `json:"description,omitempty"`
	// The context that the content is intended to support
	UseContext []UsageContext `json:"useContext,omitempty"`
	// Intended jurisdiction for Requirements (if applicable)
	Jurisdiction []CodeableConcept `json:"jurisdiction,omitempty"`
	// Why this Requirements is defined
	Purpose *custom.Markdown `json:"purpose,omitempty"`
	// Use and/or publishing restrictions
	Copyright *custom.Markdown `json:"copyright,omitempty"`
	// Copyright holder and year(s)
	CopyrightLabel *string `json:"copyrightLabel,omitempty"`
	// Other set of Requirements this builds on
	DerivedFrom []custom.Canonical[Requirements] `json:"derivedFrom,omitempty"`
	// External artifact (rule/document etc. that) created this set of requirements
	Reference []custom.URL `json:"reference,omitempty"`
	// Actor for these requirements
	Actor []custom.Canonical[ActorDefinition] `json:"actor,omitempty"`
	// Actual statement as markdown
	Statement []RequirementsStatement `json:"statement,omitempty"`
}

type RequirementsStatement struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Key that identifies this statement
	Key custom.ID `json:"key"`
	// Short Human label for this statement
	Label *string `json:"label,omitempty"`
	// SHALL | SHOULD | MAY | SHOULD-NOT
	Conformance []custom.Code `json:"conformance,omitempty"`
	// Set to true if requirements statement is conditional
	Conditionality *bool `json:"conditionality,omitempty"`
	// The actual requirement
	Requirement custom.Markdown `json:"requirement"`
	// Another statement this clarifies/restricts ([url#]key)
	DerivedFrom *string `json:"derivedFrom,omitempty"`
	// A larger requirement that this requirement helps to refine and enable
	Parent *string `json:"parent,omitempty"`
	// Design artifact that satisfies this requirement
	SatisfiedBy []custom.URL `json:"satisfiedBy,omitempty"`
	// External artifact (rule/document etc. that) created this requirement
	Reference []custom.URL `json:"reference,omitempty"`
	// Who asked for this statement
	Source []custom.Reference[RequirementsStatementSource] `json:"source,omitempty"`
}

type OtherRequirements Requirements

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
	return "Requirements"
}

func (r Requirements) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherRequirements
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherRequirements: OtherRequirements(r), ResourceType: r.ResourceType()})
}

func UnmarshalRequirements(b []byte) (res Requirements, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
