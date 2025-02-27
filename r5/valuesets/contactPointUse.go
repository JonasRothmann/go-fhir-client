// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/contact-point-use
*/type ContactPointUse string

const (
	// A communication contact point at a home; attempted contacts for business purposes might intrude privacy and chances are one will contact family or other household members instead of the person one wishes to call. Typically used with urgent cases, or if no other contacts are available.
	ContactPointUseHome ContactPointUse = "home"
	// An office contact point. First choice for business related contacts during business hours.
	ContactPointUseWork ContactPointUse = "work"
	// A temporary contact point. The period can provide more detailed information.
	ContactPointUseTemp ContactPointUse = "temp"
	// This contact point is no longer in use (or was never correct, but retained for records).
	ContactPointUseOld ContactPointUse = "old"
	// A telecommunication device that moves and stays with its owner. May have characteristics of all other use codes, suitable for urgent matters, not the first choice for routine business.
	ContactPointUseMobile ContactPointUse = "mobile"
)
