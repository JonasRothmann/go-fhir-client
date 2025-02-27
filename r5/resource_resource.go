// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Resource
type Resource struct {
	// Logical id of this artifact
	Id *string `json:"id,omitempty"`
	// Metadata about the resource
	Meta *Meta `json:"meta,omitempty"`
	// A set of rules under which this content was created
	ImplicitRules *custom.URI `json:"implicitRules,omitempty"`
	// Language of the resource content
	Language *custom.Code `json:"language,omitempty"`
}

type OtherResource Resource

func (r Resource) ResourceType() string {
	return "Resource"
}

func (r Resource) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherResource
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherResource: OtherResource(r), ResourceType: r.ResourceType()})
}

func UnmarshalResource(b []byte) (res Resource, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
