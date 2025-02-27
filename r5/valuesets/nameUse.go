// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/name-use
*/type NameUse string

const (
	// Known as/conventional/the one you normally use.
	NameUseUsual NameUse = "usual"
	// The formal name as registered in an official (government) registry, but which name might not be commonly used. May be called "legal name".
	NameUseOfficial NameUse = "official"
	// A temporary name. Name.period can provide more detailed information. This may also be used for temporary names assigned at birth or in emergency situations.
	NameUseTemp NameUse = "temp"
	// A name that is used to address the person in an informal manner, but is not part of their formal or usual name.
	NameUseNickname NameUse = "nickname"
	// Anonymous assigned name, alias, or pseudonym (used to protect a person's identity for privacy reasons).
	NameUseAnonymous NameUse = "anonymous"
	// This name is no longer in use (or was never correct, but retained for records).
	NameUseOld NameUse = "old"
)
