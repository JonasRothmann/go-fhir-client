// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/resource-validation-mode
*/type ResourceValidationMode string

const (
	// The server checks the content, and then checks that the content would be acceptable as a create (e.g. that the content would not violate any uniqueness constraints).
	ResourceValidationModeCreate ResourceValidationMode = "create"
	// The server checks the content, and then checks that it would accept it as an update against the nominated specific resource (e.g. that there are no changes to immutable fields the server does not allow to change and checking version integrity if appropriate).
	ResourceValidationModeUpdate ResourceValidationMode = "update"
	// The server ignores the content and checks that the nominated resource is allowed to be deleted (e.g. checking referential integrity rules).
	ResourceValidationModeDelete ResourceValidationMode = "delete"
	// The server checks an existing resource (must be nominated by id, not provided as a parameter) as valid against the nominated profile.
	ResourceValidationModeProfile ResourceValidationMode = "profile"
)
