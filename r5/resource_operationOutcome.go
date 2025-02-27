// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/OperationOutcome
type OperationOutcome struct {
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
	// A single issue associated with the action
	Issue []OperationOutcomeIssue `json:"issue"`
}

type OperationOutcomeIssue struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// fatal | error | warning | information | success
	Severity custom.Code `json:"severity"`
	// Error or warning code
	Code custom.Code `json:"code"`
	// Additional details about the error
	Details *CodeableConcept `json:"details,omitempty"`
	// Additional diagnostic information about the issue
	Diagnostics *string `json:"diagnostics,omitempty"`
	// Deprecated: Path of element(s) related to issue
	Location []string `json:"location,omitempty"`
	// FHIRPath of element(s) related to issue
	Expression []string `json:"expression,omitempty"`
}

type OtherOperationOutcome OperationOutcome

func (o OperationOutcome) ResourceType() string {
	return "OperationOutcome"
}

func (o OperationOutcome) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherOperationOutcome
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherOperationOutcome: OtherOperationOutcome(o), ResourceType: o.ResourceType()})
}

func UnmarshalOperationOutcome(b []byte) (res OperationOutcome, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
