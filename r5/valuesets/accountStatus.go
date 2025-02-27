// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/account-status
*/type AccountStatus string

const (
	// This account is active and may be used.
	AccountStatusActive AccountStatus = "active"
	// This account is inactive and should not be used to track financial information.
	AccountStatusInactive AccountStatus = "inactive"
	// This instance should not have been part of this patient's medical record.
	AccountStatusEnteredInError AccountStatus = "entered-in-error"
	// This account is on hold.
	AccountStatusOnHold AccountStatus = "on-hold"
	// The account status is unknown.
	AccountStatusUnknown AccountStatus = "unknown"
)
