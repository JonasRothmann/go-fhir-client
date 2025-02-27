// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/RelatedArtifact
type RelatedArtifact struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// documentation | justification | citation | predecessor | successor | derived-from | depends-on | composed-of | part-of | amends | amended-with | appends | appended-with | cites | cited-by | comments-on | comment-in | contains | contained-in | corrects | correction-in | replaces | replaced-with | retracts | retracted-by | signs | similar-to | supports | supported-with | transforms | transformed-into | transformed-with | documents | specification-of | created-with | cite-as
	Type custom.Code `json:"type"`
	// Additional classifiers
	Classifier []CodeableConcept `json:"classifier,omitempty"`
	// Short label
	Label *string `json:"label,omitempty"`
	// Brief description of the related artifact
	Display *string `json:"display,omitempty"`
	// Bibliographic citation for the artifact
	Citation *custom.Markdown `json:"citation,omitempty"`
	// What document is being referenced
	Document *Attachment `json:"document,omitempty"`
	// What artifact is being referenced
	Resource *custom.Canonical[Resource] `json:"resource,omitempty"`
	// What artifact, if not a conformance resource
	ResourceReference *custom.Reference[Resource] `json:"resourceReference,omitempty"`
	// draft | active | retired | unknown
	PublicationStatus *custom.Code `json:"publicationStatus,omitempty"`
	// Date of publication of the artifact being referred to
	PublicationDate *custom.Date `json:"publicationDate,omitempty"`
}

type OtherRelatedArtifact RelatedArtifact

func (r RelatedArtifact) ResourceType() string {
	return "RelatedArtifact"
}

func (r RelatedArtifact) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherRelatedArtifact
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherRelatedArtifact: OtherRelatedArtifact(r), ResourceType: r.ResourceType()})
}

func UnmarshalRelatedArtifact(b []byte) (res RelatedArtifact, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
