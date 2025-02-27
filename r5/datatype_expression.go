// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Expression
type Expression struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Natural language description of the condition
	Description *string `json:"description,omitempty"`
	// Short name assigned to expression for reuse
	Name *custom.Code `json:"name,omitempty"`
	// text/cql | text/fhirpath | application/x-fhir-query | etc.
	Language *custom.Code `json:"language,omitempty"`
	// Expression in specified language
	Expression *string `json:"expression,omitempty"`
	// Where the expression is found
	Reference *custom.URI `json:"reference,omitempty"`
}

type OtherExpression Expression

func (e Expression) ResourceType() string {
	return "Expression"
}

func (e Expression) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherExpression
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherExpression: OtherExpression(e), ResourceType: e.ResourceType()})
}

func UnmarshalExpression(b []byte) (res Expression, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
