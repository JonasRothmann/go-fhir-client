// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/Meta
type Meta struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Version specific identifier
	VersionId *custom.ID `bson:"versionId,omitempty" json:"versionId,omitempty"`
	// When the resource version last changed
	LastUpdated *custom.Instant `bson:"lastUpdated,omitempty" json:"lastUpdated,omitempty"`
	// Identifies where the resource comes from
	Source *custom.URI `bson:"source,omitempty" json:"source,omitempty"`
	// Profiles this resource claims to conform to
	Profile []custom.Canonical[StructureDefinition] `bson:"profile,omitempty" json:"profile,omitempty"`
	// Security Labels applied to this resource
	Security []Coding `bson:"security,omitempty" json:"security,omitempty"`
	// Tags applied to this resource
	Tag []Coding `bson:"tag,omitempty" json:"tag,omitempty"`
}

func (m Meta) ResourceType() string {
	return "StructureDefinition"
}
