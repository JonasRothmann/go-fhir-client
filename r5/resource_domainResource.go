// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/DomainResource
type DomainResource struct {
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
}

type OtherDomainResource DomainResource

func (d DomainResource) ResourceType() string {
	return "DomainResource"
}

func (d DomainResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherDomainResource
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherDomainResource: OtherDomainResource(d), ResourceType: d.ResourceType()})
}

func UnmarshalDomainResource(b []byte) (res DomainResource, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
