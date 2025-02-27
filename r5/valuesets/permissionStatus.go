// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/permission-status
*/type PermissionStatus string

const (
	// Permission is given.
	PermissionStatusActive PermissionStatus = "active"
	// Permission was entered in error and is not active.
	PermissionStatusEnteredInError PermissionStatus = "entered-in-error"
	// Permission is being defined.
	PermissionStatusDraft PermissionStatus = "draft"
	// Permission not granted.
	PermissionStatusRejected PermissionStatus = "rejected"
)
