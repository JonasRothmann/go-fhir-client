// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/conditional-read-status
*/type ConditionalReadStatus string

const (
	// No support for conditional reads.
	ConditionalReadStatusNotSupported ConditionalReadStatus = "not-supported"
	// Conditional reads are supported, but only with the If-Modified-Since HTTP Header.
	ConditionalReadStatusIfModifiedSince ConditionalReadStatus = "modified-since"
	// Conditional reads are supported, but only with the If-None-Match HTTP Header.
	ConditionalReadStatusIfNoneMatch ConditionalReadStatus = "not-match"
	// Conditional reads are supported, with both If-Modified-Since and If-None-Match HTTP Headers.
	ConditionalReadStatusFullSupport ConditionalReadStatus = "full-support"
)
