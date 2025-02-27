// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Binary
type Binary struct {
	// Logical id of this artifact
	Id *string `json:"id,omitempty"`
	// Metadata about the resource
	Meta *Meta `json:"meta,omitempty"`
	// A set of rules under which this content was created
	ImplicitRules *custom.URI `json:"implicitRules,omitempty"`
	// Language of the resource content
	Language *custom.Code `json:"language,omitempty"`
	// MimeType of the binary content
	ContentType custom.Code `json:"contentType"`
	// Identifies another resource to use as proxy when enforcing access control
	SecurityContext *custom.Reference[Resource] `json:"securityContext,omitempty"`
	// The actual content
	Data *custom.Base64binary `json:"data,omitempty"`
}

type OtherBinary Binary

func (b Binary) ResourceType() string {
	return "Binary"
}

func (b Binary) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherBinary
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherBinary: OtherBinary(b), ResourceType: b.ResourceType()})
}

func UnmarshalBinary(b []byte) (res Binary, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
