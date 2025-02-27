// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Narrative
type Narrative struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// generated | extensions | additional | empty
	Status custom.Code `json:"status"`
	// Limited xhtml content
	Div custom.XHTML `json:"div"`
}

type OtherNarrative Narrative

func (n Narrative) ResourceType() string {
	return "Narrative"
}

func (n Narrative) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherNarrative
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherNarrative: OtherNarrative(n), ResourceType: n.ResourceType()})
}

func UnmarshalNarrative(b []byte) (res Narrative, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
