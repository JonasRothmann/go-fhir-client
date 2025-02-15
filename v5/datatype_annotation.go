// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Annotation
type Annotation struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Individual responsible for the annotation
	Author *custom.Reference[AnnotationAuthor] `bson:"author,omitempty" json:"author,omitempty"`
	// When the annotation was made
	Time *custom.DateTime `bson:"time,omitempty" json:"time,omitempty"`
	// The annotation  - text content (as markdown)
	Text custom.Markdown `bson:"text" json:"text"`
}

type AnnotationAuthor interface {
	gofhirclient.Resource

	Is_AnnotationAuthor()
}

func (p Practitioner) Is_AnnotationAuthor()     {}
func (p PractitionerRole) Is_AnnotationAuthor() {}
func (p Patient) Is_AnnotationAuthor()          {}
func (r RelatedPerson) Is_AnnotationAuthor()    {}
func (o Organization) Is_AnnotationAuthor()     {}

func (a Annotation) ResourceType() string {
	return "StructureDefinition"
}
