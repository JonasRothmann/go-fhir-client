// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/Binary
type Binary struct {
	// Logical id of this artifact
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Metadata about the resource
	Meta *Meta `bson:"meta,omitempty" json:"meta,omitempty"`
	// A set of rules under which this content was created
	ImplicitRules *custom.URI `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	// Language of the resource content
	Language *custom.Code `bson:"language,omitempty" json:"language,omitempty"`
	// MimeType of the binary content
	ContentType custom.Code `bson:"contentType" json:"contentType"`
	// Identifies another resource to use as proxy when enforcing access control
	SecurityContext *custom.Reference[Resource] `bson:"securityContext,omitempty" json:"securityContext,omitempty"`
	// The actual content
	Data *custom.Base64binary `bson:"data,omitempty" json:"data,omitempty"`
}

func (b Binary) ResourceType() string {
	return "StructureDefinition"
}
