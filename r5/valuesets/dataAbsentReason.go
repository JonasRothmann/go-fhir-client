// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/data-absent-reason
*/type DataAbsentReason string

const (
	// The value is expected to exist but is not known.
	DataAbsentReasonUnknown DataAbsentReason = "unknown"
	// The information is not available due to security, privacy or related reasons.
	DataAbsentReasonMasked DataAbsentReason = "masked"
	// There is no proper value for this element (e.g. last menstrual period for a male).
	DataAbsentReasonNotApplicable DataAbsentReason = "not-applicable"
	// The source system wasn't capable of supporting this element.
	DataAbsentReasonUnsupported DataAbsentReason = "unsupported"
	// The content of the data is represented in the resource narrative.
	DataAbsentReasonAsText DataAbsentReason = "as-text"
	// Some system or workflow process error means that the information is not available.
	DataAbsentReasonError DataAbsentReason = "error"
	// The value is not available because the observation procedure (test, etc.) was not performed.
	DataAbsentReasonNotPerformed DataAbsentReason = "not-performed"
	// The value is not permitted in this context (e.g. due to profiles, or the base data types).
	DataAbsentReasonNotPermitted DataAbsentReason = "not-permitted"
)
