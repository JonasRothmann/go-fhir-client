// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/identifier-use
*/type IdentifierUse string

const (
	// The identifier recommended for display and use in real-world interactions which should be used when such identifier is different from the "official" identifier.
	IdentifierUseUsual IdentifierUse = "usual"
	// The identifier considered to be most trusted for the identification of this item. Sometimes also known as "primary" and "main". The determination of "official" is subjective and implementation guides often provide additional guidelines for use.
	IdentifierUseOfficial IdentifierUse = "official"
	// A temporary identifier.
	IdentifierUseTemp IdentifierUse = "temp"
	// An identifier that was assigned in secondary use - it serves to identify the object in a relative context, but cannot be consistently assigned to the same object again in a different context.
	IdentifierUseSecondary IdentifierUse = "secondary"
	// The identifier id no longer considered valid, but may be relevant for search purposes.  E.g. Changes to identifier schemes, account merges, etc.
	IdentifierUseOld IdentifierUse = "old"
)
