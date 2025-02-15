// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/OperationOutcome
type OperationOutcomeIssue struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// fatal | error | warning | information | success
	Severity custom.Code `bson:"severity" json:"severity"`
	// Error or warning code
	Code custom.Code `bson:"code" json:"code"`
	// Additional details about the error
	Details *CodeableConcept `bson:"details,omitempty" json:"details,omitempty"`
	// Additional diagnostic information about the issue
	Diagnostics *string `bson:"diagnostics,omitempty" json:"diagnostics,omitempty"`
	// Deprecated: Path of element(s) related to issue
	Location []string `bson:"location,omitempty" json:"location,omitempty"`
	// FHIRPath of element(s) related to issue
	Expression []string `bson:"expression,omitempty" json:"expression,omitempty"`
}

type OperationOutcome struct {
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
	// A single issue associated with the action
	Issue []OperationOutcomeIssue `bson:"issue" json:"issue"`
}

func (o OperationOutcome) ResourceType() string {
	return "StructureDefinition"
}
