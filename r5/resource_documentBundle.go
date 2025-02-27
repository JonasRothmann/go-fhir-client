// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/document-bundle
type DocumentBundleLink struct {
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

type DocumentBundleEntry struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Links related to this entry
	Link []interface{} `json:"link,omitempty"`
	// URI for resource (e.g. the absolute URL server address, URI for UUID/OID, etc.)
	FullUrl custom.URI `json:"fullUrl"`
	// A resource in the bundle
	Resource Resource `json:"resource"`
	// Search related information
	Search *DocumentBundleEntrySearch `json:"search,omitempty"`
	// Additional execution information (transaction/batch/history)
	Request *DocumentBundleEntryRequest `json:"request,omitempty"`
	// Results of execution (transaction/batch/history)
	Response *DocumentBundleEntryResponse `json:"response,omitempty"`
}

type DocumentBundleEntrySearch struct {
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

type DocumentBundleEntryRequest struct {
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

type DocumentBundleEntryResponse struct {
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

type DocumentBundle struct {
	// Logical id of this artifact
	Id *string `json:"id,omitempty"`
	// Metadata about the resource
	Meta *Meta `json:"meta,omitempty"`
	// A set of rules under which this content was created
	ImplicitRules *custom.URI `json:"implicitRules,omitempty"`
	// Language of the resource content
	Language *custom.Code `json:"language,omitempty"`
	// Persistent identifier for the bundle
	Identifier Identifier `json:"identifier"`
	// document | message | transaction | transaction-response | batch | batch-response | history | searchset | collection | subscription-notification
	Type custom.Code `json:"type"`
	// When the bundle was assembled
	Timestamp custom.Instant `json:"timestamp"`
	// If search, the total number of matches
	Total *uint32 `json:"total,omitempty"`
	// Links related to this Bundle
	Link []DocumentBundleLink `json:"link,omitempty"`
	// Entry in the bundle - will have a resource or information
	Entry []DocumentBundleEntry `json:"entry"`
	// Digital Signature
	Signature *Signature `json:"signature,omitempty"`
	// Issues with the Bundle
	Issues *Resource `json:"issues,omitempty"`
}

type DocumentBundleIdentifier struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// usual | official | temp | secondary | old (If known)
	Use *custom.Code `json:"use,omitempty"`
	// Description of identifier
	Type *CodeableConcept `json:"type,omitempty"`
	// The namespace for the identifier value
	System custom.URI `json:"system"`
	// The value that is unique
	Value string `json:"value"`
	// Time period when id is/was valid for use
	Period *Period `json:"period,omitempty"`
	// Organization that issued id (may be just text)
	Assigner *custom.Reference[Organization] `json:"assigner,omitempty"`
}

type OtherDocumentBundle DocumentBundle

func (d DocumentBundle) ResourceType() string {
	return "DocumentBundle"
}

func (d DocumentBundle) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherDocumentBundle
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherDocumentBundle: OtherDocumentBundle(d), ResourceType: d.ResourceType()})
}

func UnmarshalDocumentBundle(b []byte) (res DocumentBundle, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
