// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Meta
type Meta struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Version specific identifier
	VersionId *custom.ID `json:"versionId,omitempty"`
	// When the resource version last changed
	LastUpdated *custom.Instant `json:"lastUpdated,omitempty"`
	// Identifies where the resource comes from
	Source *custom.URI `json:"source,omitempty"`
	// Profiles this resource claims to conform to
	Profile []custom.Canonical[StructureDefinition] `json:"profile,omitempty"`
	// Security Labels applied to this resource
	Security []Coding `json:"security,omitempty"`
	// Tags applied to this resource
	Tag []Coding `json:"tag,omitempty"`
}

type OtherMeta Meta

func (m Meta) ResourceType() string {
	return "Meta"
}

func (m Meta) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMeta
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherMeta: OtherMeta(m), ResourceType: m.ResourceType()})
}

func UnmarshalMeta(b []byte) (res Meta, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
