// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/FormularyItem
type FormularyItem struct {
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
	// Business identifier for this formulary item
	Identifier []Identifier `json:"identifier,omitempty"`
	// Codes that identify this formulary item
	Code *CodeableConcept `json:"code,omitempty"`
	// active | entered-in-error | inactive
	Status *custom.Code `json:"status,omitempty"`
}

type OtherFormularyItem FormularyItem

func (f FormularyItem) ResourceType() string {
	return "FormularyItem"
}

func (f FormularyItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherFormularyItem
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherFormularyItem: OtherFormularyItem(f), ResourceType: f.ResourceType()})
}

func UnmarshalFormularyItem(b []byte) (res FormularyItem, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
