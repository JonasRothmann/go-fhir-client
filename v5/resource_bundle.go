// GENERATED CODE – DO NOT EDIT!

package v5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Bundle
type Bundle struct {
	// Logical id of this artifact
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Metadata about the resource
	Meta *Meta `bson:"meta,omitempty" json:"meta,omitempty"`
	// A set of rules under which this content was created
	ImplicitRules *custom.URI `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	// Language of the resource content
	Language *custom.Code `bson:"language,omitempty" json:"language,omitempty"`
	// Persistent identifier for the bundle
	Identifier *Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// document | message | transaction | transaction-response | batch | batch-response | history | searchset | collection | subscription-notification
	Type custom.Code `bson:"type" json:"type"`
	// When the bundle was assembled
	Timestamp *custom.Instant `bson:"timestamp,omitempty" json:"timestamp,omitempty"`
	// If search, the total number of matches
	Total *uint32 `bson:"total,omitempty" json:"total,omitempty"`
	// Links related to this Bundle
	Link []BundleLink `bson:"link,omitempty" json:"link,omitempty"`
	// Entry in the bundle - will have a resource or information
	Entry []BundleEntry `bson:"entry,omitempty" json:"entry,omitempty"`
	// Digital Signature
	Signature *Signature `bson:"signature,omitempty" json:"signature,omitempty"`
	// Issues with the Bundle
	Issues *Resource `bson:"issues,omitempty" json:"issues,omitempty"`
}

type BundleLink struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// See http://www.iana.org/assignments/link-relations/link-relations.xhtml#link-relations-1
	Relation custom.Code `bson:"relation" json:"relation"`
	// Reference details for the link
	Url custom.URI `bson:"url" json:"url"`
}

type BundleEntry struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Links related to this entry
	Link []interface{} `bson:"link,omitempty" json:"link,omitempty"`
	// URI for resource (e.g. the absolute URL server address, URI for UUID/OID, etc.)
	FullUrl *custom.URI `bson:"fullUrl,omitempty" json:"fullUrl,omitempty"`
	// A resource in the bundle
	Resource *Resource `bson:"resource,omitempty" json:"resource,omitempty"`
	// Search related information
	Search *BundleEntrySearch `bson:"search,omitempty" json:"search,omitempty"`
	// Additional execution information (transaction/batch/history)
	Request *BundleEntryRequest `bson:"request,omitempty" json:"request,omitempty"`
	// Results of execution (transaction/batch/history)
	Response *BundleEntryResponse `bson:"response,omitempty" json:"response,omitempty"`
}

type BundleEntrySearch struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// match | include - why this is in the result set
	Mode *custom.Code `bson:"mode,omitempty" json:"mode,omitempty"`
	// Search ranking (between 0 and 1)
	Score *json.Number `bson:"score,omitempty" json:"score,omitempty"`
}

type BundleEntryRequest struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// GET | HEAD | POST | PUT | DELETE | PATCH
	Method custom.Code `bson:"method" json:"method"`
	// URL for HTTP equivalent of this entry
	Url custom.URI `bson:"url" json:"url"`
	// For managing cache validation
	IfNoneMatch *string `bson:"ifNoneMatch,omitempty" json:"ifNoneMatch,omitempty"`
	// For managing cache currency
	IfModifiedSince *custom.Instant `bson:"ifModifiedSince,omitempty" json:"ifModifiedSince,omitempty"`
	// For managing update contention
	IfMatch *string `bson:"ifMatch,omitempty" json:"ifMatch,omitempty"`
	// For conditional creates
	IfNoneExist *string `bson:"ifNoneExist,omitempty" json:"ifNoneExist,omitempty"`
}

type BundleEntryResponse struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Status response code (text optional)
	Status string `bson:"status" json:"status"`
	// The location (if the operation returns a location)
	Location *custom.URI `bson:"location,omitempty" json:"location,omitempty"`
	// The Etag for the resource (if relevant)
	Etag *string `bson:"etag,omitempty" json:"etag,omitempty"`
	// Server's date time modified
	LastModified *custom.Instant `bson:"lastModified,omitempty" json:"lastModified,omitempty"`
	// OperationOutcome with hints and warnings (for batch/transaction)
	Outcome *Resource `bson:"outcome,omitempty" json:"outcome,omitempty"`
}

func (b Bundle) ResourceType() string {
	return "StructureDefinition"
}
