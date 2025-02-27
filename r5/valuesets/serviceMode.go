// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/service-mode
*/type ServiceMode string

const (
	// The service will be provided in person
	ServiceModeInPerson ServiceMode = "in-person"
	// The service will be provided by a teleconferencing facility or regular telephone
	ServiceModeTelephone ServiceMode = "telephone"
	// The service will be provided over a video-conference facility
	ServiceModeVideoconference ServiceMode = "videoconference"
	// This service will be provided via a realtime chat/messaging conversation
	ServiceModeChat ServiceMode = "chat"
)
