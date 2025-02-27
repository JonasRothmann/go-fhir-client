// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Annotation
type Annotation struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Individual responsible for the annotation
	AuthorReference *custom.Reference[AnnotationAuthorReference] `json:"authorReference,omitempty"`
	// Individual responsible for the annotation
	AuthorString *string `json:"authorString,omitempty"`
	// When the annotation was made
	Time *custom.DateTime `json:"time,omitempty"`
	// The annotation  - text content (as markdown)
	Text custom.Markdown `json:"text"`
}

type OtherAnnotation Annotation

type AnnotationAuthorReference interface {
	gofhirclient.Resource

	Is_AnnotationAuthorReference()
}

func (p Practitioner) Is_AnnotationAuthorReference()     {}
func (p PractitionerRole) Is_AnnotationAuthorReference() {}
func (p Patient) Is_AnnotationAuthorReference()          {}
func (r RelatedPerson) Is_AnnotationAuthorReference()    {}
func (o Organization) Is_AnnotationAuthorReference()     {}

func (a Annotation) ResourceType() string {
	return "Annotation"
}

func (a Annotation) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherAnnotation
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherAnnotation: OtherAnnotation(a), ResourceType: a.ResourceType()})
}

func UnmarshalAnnotation(b []byte) (res Annotation, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
