// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/namingsystem-type
*/type NamingSystemType string

const (
	// The naming system is used to define concepts and symbols to represent those concepts; e.g. UCUM, LOINC, NDC code, local lab codes, etc.
	NamingSystemTypeCodesystem NamingSystemType = "codesystem"
	// The naming system is used to manage identifiers (e.g. license numbers, order numbers, etc.).
	NamingSystemTypeIdentifier NamingSystemType = "identifier"
	// The naming system is used as the root for other identifiers and naming systems.
	NamingSystemTypeRoot NamingSystemType = "root"
)
