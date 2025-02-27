// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/search-set-bundle
type SearchSetBundleLink struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// See http://www.iana.org/assignments/link-relations/link-relations.xhtml#link-relations-1
	Relation custom.Code `json:"relation"`
	// Reference details for the link
	Url custom.URI `json:"url"`
}

type SearchSetBundleEntry struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Links related to this entry
	Link []interface{} `json:"link,omitempty"`
	// URI for resource (e.g. the absolute URL server address, URI for UUID/OID, etc.)
	FullUrl *custom.URI `json:"fullUrl,omitempty"`
	// A resource in the bundle
	Resource *Resource `json:"resource,omitempty"`
	// Search related information
	Search *SearchSetBundleEntrySearch `json:"search,omitempty"`
	// Additional execution information (transaction/batch/history)
	Request *SearchSetBundleEntryRequest `json:"request,omitempty"`
	// Results of execution (transaction/batch/history)
	Response *SearchSetBundleEntryResponse `json:"response,omitempty"`
}

type SearchSetBundleEntrySearch struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// match | include - why this is in the result set
	Mode *custom.Code `json:"mode,omitempty"`
	// Search ranking (between 0 and 1)
	Score *json.Number `json:"score,omitempty"`
}

type SearchSetBundleEntryRequest struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// GET | HEAD | POST | PUT | DELETE | PATCH
	Method custom.Code `json:"method"`
	// URL for HTTP equivalent of this entry
	Url custom.URI `json:"url"`
	// For managing cache validation
	IfNoneMatch *string `json:"ifNoneMatch,omitempty"`
	// For managing cache currency
	IfModifiedSince *custom.Instant `json:"ifModifiedSince,omitempty"`
	// For managing update contention
	IfMatch *string `json:"ifMatch,omitempty"`
	// For conditional creates
	IfNoneExist *string `json:"ifNoneExist,omitempty"`
}

type SearchSetBundleEntryResponse struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Status response code (text optional)
	Status string `json:"status"`
	// The location (if the operation returns a location)
	Location *custom.URI `json:"location,omitempty"`
	// The Etag for the resource (if relevant)
	Etag *string `json:"etag,omitempty"`
	// Server's date time modified
	LastModified *custom.Instant `json:"lastModified,omitempty"`
	// OperationOutcome with hints and warnings (for batch/transaction)
	Outcome *Resource `json:"outcome,omitempty"`
}

type SearchSetBundle struct {
	// Logical id of this artifact
	Id *string `json:"id,omitempty"`
	// Metadata about the resource
	Meta *Meta `json:"meta,omitempty"`
	// A set of rules under which this content was created
	ImplicitRules *custom.URI `json:"implicitRules,omitempty"`
	// Language of the resource content
	Language *custom.Code `json:"language,omitempty"`
	// Persistent identifier for the bundle
	Identifier *Identifier `json:"identifier,omitempty"`
	// document | message | transaction | transaction-response | batch | batch-response | history | searchset | collection | subscription-notification
	Type custom.Code `json:"type"`
	// When the bundle was assembled
	Timestamp *custom.Instant `json:"timestamp,omitempty"`
	// If search, the total number of matches
	Total *uint32 `json:"total,omitempty"`
	// Links related to this Bundle
	Link []SearchSetBundleLink `json:"link,omitempty"`
	// Entry in the bundle - will have a resource or information
	Entry []SearchSetBundleEntry `json:"entry,omitempty"`
	// Digital Signature
	Signature *Signature `json:"signature,omitempty"`
	// Issues with the Bundle
	Issues *Resource `json:"issues,omitempty"`
}

type OtherSearchSetBundle SearchSetBundle

func (s SearchSetBundle) ResourceType() string {
	return "SearchSetBundle"
}

func (s SearchSetBundle) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSearchSetBundle
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherSearchSetBundle: OtherSearchSetBundle(s), ResourceType: s.ResourceType()})
}

func UnmarshalSearchSetBundle(b []byte) (res SearchSetBundle, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
