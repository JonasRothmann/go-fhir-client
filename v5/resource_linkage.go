// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Linkage
type Linkage struct {
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
	// Whether this linkage assertion is active or not
	Active *bool `bson:"active,omitempty" json:"active,omitempty"`
	// Who is responsible for linkages
	Author *custom.Reference[LinkageAuthor] `bson:"author,omitempty" json:"author,omitempty"`
	// Item to be linked
	Item []LinkageItem `bson:"item" json:"item"`
}

type LinkageItem struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// source | alternate | historical
	Type custom.Code `bson:"type" json:"type"`
	// Resource being linked
	Resource custom.Reference[Resource] `bson:"resource" json:"resource"`
}

type LinkageAuthor interface {
	gofhirclient.Resource

	Is_LinkageAuthor()
}

func (p Practitioner) Is_LinkageAuthor()     {}
func (p PractitionerRole) Is_LinkageAuthor() {}
func (o Organization) Is_LinkageAuthor()     {}

func (l Linkage) ResourceType() string {
	return "StructureDefinition"
}
