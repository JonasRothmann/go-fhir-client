// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/namingsystem-identifier-type
*/type NamingSystemIdentifierType string

const (
	// An ISO object identifier; e.g. 1.2.3.4.5.
	NamingSystemIdentifierTypeOid NamingSystemIdentifierType = "oid"
	// A universally unique identifier of the form a5afddf4-e880-459b-876e-e4591b0acc11.
	NamingSystemIdentifierTypeUuid NamingSystemIdentifierType = "uuid"
	// A uniform resource identifier (ideally a URL - uniform resource locator); e.g. http://unitsofmeasure.org.
	NamingSystemIdentifierTypeUri NamingSystemIdentifierType = "uri"
	// An IRI string that can be prepended to the code to obtain a concept IRI for RDF applications. This should be a valid, absolute IRI as defined in RFC 3987. See rdf.html#iri-stem for details on how this value may be used.
	NamingSystemIdentifierTypeIriStem NamingSystemIdentifierType = "iri-stem"
	// A short string published by HL7 for use in the V2 family of standsrds to idenfify a code system in the V12 coded data types CWE, CNE, and CF. The code values are also published by HL7 at http://www.hl7.org/Special/committees/vocab/table_0396/index.cfm
	NamingSystemIdentifierTypeV2csmnemonic NamingSystemIdentifierType = "v2csmnemonic"
	// Some other type of unique identifier; e.g. HL7-assigned reserved string such as LN for LOINC.
	NamingSystemIdentifierTypeOther NamingSystemIdentifierType = "other"
)
