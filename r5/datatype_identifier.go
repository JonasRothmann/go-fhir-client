// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Identifier
type Identifier struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// usual | official | temp | secondary | old (If known)
	Use *custom.Code `json:"use,omitempty"`
	// Description of identifier
	Type *CodeableConcept `json:"type,omitempty"`
	// The namespace for the identifier value
	System *custom.URI `json:"system,omitempty"`
	// The value that is unique
	Value *string `json:"value,omitempty"`
	// Time period when id is/was valid for use
	Period *Period `json:"period,omitempty"`
	// Organization that issued id (may be just text)
	Assigner *custom.Reference[Organization] `json:"assigner,omitempty"`
}

type OtherIdentifier Identifier

func (i Identifier) ResourceType() string {
	return "Identifier"
}

func (i Identifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherIdentifier
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherIdentifier: OtherIdentifier(i), ResourceType: i.ResourceType()})
}

func UnmarshalIdentifier(b []byte) (res Identifier, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
