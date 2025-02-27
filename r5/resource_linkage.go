// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Linkage
type Linkage struct {
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
	// Whether this linkage assertion is active or not
	Active *bool `json:"active,omitempty"`
	// Who is responsible for linkages
	Author *custom.Reference[LinkageAuthor] `json:"author,omitempty"`
	// Item to be linked
	Item []LinkageItem `json:"item"`
}

type LinkageItem struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// source | alternate | historical
	Type custom.Code `json:"type"`
	// Resource being linked
	Resource custom.Reference[Resource] `json:"resource"`
}

type OtherLinkage Linkage

type LinkageAuthor interface {
	gofhirclient.Resource

	Is_LinkageAuthor()
}

func (p Practitioner) Is_LinkageAuthor()     {}
func (p PractitionerRole) Is_LinkageAuthor() {}
func (o Organization) Is_LinkageAuthor()     {}

func (l Linkage) ResourceType() string {
	return "Linkage"
}

func (l Linkage) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherLinkage
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherLinkage: OtherLinkage(l), ResourceType: l.ResourceType()})
}

func UnmarshalLinkage(b []byte) (res Linkage, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
