// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/RelatedArtifact
type RelatedArtifact struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// documentation | justification | citation | predecessor | successor | derived-from | depends-on | composed-of | part-of | amends | amended-with | appends | appended-with | cites | cited-by | comments-on | comment-in | contains | contained-in | corrects | correction-in | replaces | replaced-with | retracts | retracted-by | signs | similar-to | supports | supported-with | transforms | transformed-into | transformed-with | documents | specification-of | created-with | cite-as
	Type custom.Code `bson:"type" json:"type"`
	// Additional classifiers
	Classifier []CodeableConcept `bson:"classifier,omitempty" json:"classifier,omitempty"`
	// Short label
	Label *string `bson:"label,omitempty" json:"label,omitempty"`
	// Brief description of the related artifact
	Display *string `bson:"display,omitempty" json:"display,omitempty"`
	// Bibliographic citation for the artifact
	Citation *custom.Markdown `bson:"citation,omitempty" json:"citation,omitempty"`
	// What document is being referenced
	Document *Attachment `bson:"document,omitempty" json:"document,omitempty"`
	// What artifact is being referenced
	Resource *custom.Canonical[Resource] `bson:"resource,omitempty" json:"resource,omitempty"`
	// What artifact, if not a conformance resource
	ResourceReference *custom.Reference[Resource] `bson:"resourceReference,omitempty" json:"resourceReference,omitempty"`
	// draft | active | retired | unknown
	PublicationStatus *custom.Code `bson:"publicationStatus,omitempty" json:"publicationStatus,omitempty"`
	// Date of publication of the artifact being referred to
	PublicationDate *custom.Date `bson:"publicationDate,omitempty" json:"publicationDate,omitempty"`
}

func (r RelatedArtifact) ResourceType() string {
	return "StructureDefinition"
}
