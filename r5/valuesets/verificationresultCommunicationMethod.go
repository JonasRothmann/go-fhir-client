// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/verificationresult-communication-method
*/type VerificationresultCommunicationMethod string

const (
	// The information is submitted/retrieved manually (e.g. by phone, fax, paper-based)
	VerificationresultCommunicationMethodManual VerificationresultCommunicationMethod = "manual"
	// The information is submitted/retrieved via a portal
	VerificationresultCommunicationMethodPortal VerificationresultCommunicationMethod = "portal"
	// The information is retrieved (i.e. pulled) from a source (e.g. over an API)
	VerificationresultCommunicationMethodPull VerificationresultCommunicationMethod = "pull"
	// The information is sent (i.e. pushed) from a source (e.g. over an API, asynchronously, secure messaging)
	VerificationresultCommunicationMethodPush VerificationresultCommunicationMethod = "push"
)
