// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/message-transport
*/type MessageTransport string

const (
	// The application sends or receives messages using HTTP POST (may be over http: or https:).
	MessageTransportHttp MessageTransport = "http"
	// The application sends or receives messages using File Transfer Protocol.
	MessageTransportFtp MessageTransport = "ftp"
	// The application sends or receives messages using HL7's Minimal Lower Level Protocol.
	MessageTransportMllp MessageTransport = "mllp"
)
