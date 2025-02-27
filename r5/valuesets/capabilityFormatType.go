// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/fhir-format-type
- urn:ietf:bcp:13
  - application/fhir+json
  - application/fhir+xml
  - application/fhir+ttl
*/type CapabilityFormatType string

const (
	// XML content-type corresponding to the application/fhir+xml mime-type.
	CapabilityFormatTypeXml CapabilityFormatType = "xml"
	// JSON content-type corresponding to the application/fhir+json mime-type.
	CapabilityFormatTypeJson CapabilityFormatType = "json"
	// RDF content-type corresponding to the text/turtle mime-type.
	CapabilityFormatTypeTtl CapabilityFormatType = "ttl"
)
