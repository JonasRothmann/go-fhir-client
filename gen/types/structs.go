package types

import (
	"encoding/json"
	"fmt"

	"github.com/jonasrothmann/go-fhir-client/gen/overwrites"
	"github.com/rs/zerolog/log"
)

type StructureDefinitionKind string

const (
	StructureDefinitionKindComplex   StructureDefinitionKind = "complex-type"
	StructureDefinitionKindPrimitive StructureDefinitionKind = "primitive-type"
	StructureDefinitionKindResource  StructureDefinitionKind = "resource"
)

func (s StructureDefinitionKind) Path() string {
	return fmt.Sprintf("github.com/jonasrothmann/go-fhir-client/%s", s.Name())
}

func (s StructureDefinitionKind) Name() string {
	switch s {
	case StructureDefinitionKindComplex, StructureDefinitionKindPrimitive:
		return "datatypes"
	case StructureDefinitionKindResource:
		return "resources"
	default:
		log.Fatal().Msg("unknown StructureDefinitionKind")
	}

	return ""
}

type StructureDefinition struct {
	ResourceType string                  `json:"resourceType"`
	Name         *string                 `json:"name,omitempty"`
	Url          *string                 `json:"url,omitempty"`
	Kind         StructureDefinitionKind `json:"kind,omitempty"`
	Snapshot     SnapshotDefinition      `json:"snapshot"`
}

type SnapshotDefinition struct {
	Element []ElementDefinition `json:"element"`
}

type ElementDefinition struct {
	Path string    `json:"path"`
	Min  *int      `json:"min,omitempty"`
	Max  *string   `json:"max,omitempty"`
	Type []TypeRef `json:"type,omitempty"`
	Code string    `json:"code,omitempty"`
	// Additional fields (e.g. Binding, ContentReference) can be added if needed.
}

type TypeRef struct {
	Code          string   `json:"code"`
	TargetProfile []string `json:"targetProfile"`
	// Additional fields (Profile, TargetProfile, etc.) can be added if required.
}

type ResourceCollections struct {
	StructureDefinitions map[string]StructureDefinition
	// You could add ValueSets and CodeSystems here as needed.
}

type Bundle struct {
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	//Meta          *overwrites.Meta      `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules *overwrites.URI  `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language      *overwrites.Code `bson:"language,omitempty" json:"language,omitempty"`
	//Identifier    *datatypes.Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Type      overwrites.Code     `bson:"type" json:"type"`
	Timestamp *overwrites.Instant `bson:"timestamp,omitempty" json:"timestamp,omitempty"`
	Total     *uint32             `bson:"total,omitempty" json:"total,omitempty"`
	Link      []BundleLink        `bson:"link,omitempty" json:"link,omitempty"`
	Entry     []BundleEntry       `bson:"entry,omitempty" json:"entry,omitempty"`
	//Signature     *datatypes.Signature  `bson:"signature,omitempty" json:"signature,omitempty"`
	Issues *json.RawMessage `bson:"issues,omitempty" json:"issues,omitempty"`
}
type BundleLink struct {
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	//Extension         []datatypes.Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	//ModifierExtension []datatypes.Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Relation overwrites.Code `bson:"relation" json:"relation"`
	Url      overwrites.URI  `bson:"url" json:"url"`
}
type BundleEntry struct {
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	//Extension         []datatypes.Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	//ModifierExtension []datatypes.Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Link     []interface{}        `bson:"link,omitempty" json:"link,omitempty"`
	FullUrl  *overwrites.URI      `bson:"fullUrl,omitempty" json:"fullUrl,omitempty"`
	Resource *json.RawMessage     `bson:"resource,omitempty" json:"resource,omitempty"`
	Search   *BundleEntrySearch   `bson:"search,omitempty" json:"search,omitempty"`
	Request  *BundleEntryRequest  `bson:"request,omitempty" json:"request,omitempty"`
	Response *BundleEntryResponse `bson:"response,omitempty" json:"response,omitempty"`
}
type BundleEntrySearch struct {
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	//Extension         []datatypes.Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	//ModifierExtension []datatypes.Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Mode  *overwrites.Code `bson:"mode,omitempty" json:"mode,omitempty"`
	Score *json.Number     `bson:"score,omitempty" json:"score,omitempty"`
}
type BundleEntryRequest struct {
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	//Extension         []datatypes.Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	//ModifierExtension []datatypes.Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Method          overwrites.Code     `bson:"method" json:"method"`
	Url             overwrites.URI      `bson:"url" json:"url"`
	IfNoneMatch     *string             `bson:"ifNoneMatch,omitempty" json:"ifNoneMatch,omitempty"`
	IfModifiedSince *overwrites.Instant `bson:"ifModifiedSince,omitempty" json:"ifModifiedSince,omitempty"`
	IfMatch         *string             `bson:"ifMatch,omitempty" json:"ifMatch,omitempty"`
	IfNoneExist     *string             `bson:"ifNoneExist,omitempty" json:"ifNoneExist,omitempty"`
}
type BundleEntryResponse struct {
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	//Extension         []datatypes.Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	//ModifierExtension []datatypes.Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Status       string              `bson:"status" json:"status"`
	Location     *overwrites.URI     `bson:"location,omitempty" json:"location,omitempty"`
	Etag         *string             `bson:"etag,omitempty" json:"etag,omitempty"`
	LastModified *overwrites.Instant `bson:"lastModified,omitempty" json:"lastModified,omitempty"`
	Outcome      *json.RawMessage    `bson:"outcome,omitempty" json:"outcome,omitempty"`
}
