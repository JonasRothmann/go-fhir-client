// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/flag-status
*/type FlagStatus string

const (
	// A current flag that should be displayed to a user. A system may use the category to determine which user roles should view the flag.
	FlagStatusActive FlagStatus = "active"
	// The flag no longer needs to be displayed.
	FlagStatusInactive FlagStatus = "inactive"
	// The flag was added in error and should no longer be displayed.
	FlagStatusEnteredInError FlagStatus = "entered-in-error"
)
