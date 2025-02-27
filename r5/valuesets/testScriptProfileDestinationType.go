// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/testscript-profile-destination-types
*/type TestScriptProfileDestinationType string

const (
	// General FHIR server used to respond to operations sent from a FHIR client.
	TestScriptProfileDestinationTypeFhirServer TestScriptProfileDestinationType = "FHIR-Server"
	// A FHIR server acting as a Structured Data Capture Form Manager.
	TestScriptProfileDestinationTypeFhirsdcFormManager TestScriptProfileDestinationType = "FHIR-SDC-FormManager"
	// A FHIR server acting as a Structured Data Capture Form Processor.
	TestScriptProfileDestinationTypeFhirsdcFormProcessor TestScriptProfileDestinationType = "FHIR-SDC-FormProcessor"
	// A FHIR server acting as a Structured Data Capture Form Receiver.
	TestScriptProfileDestinationTypeFhirsdcFormReceiver TestScriptProfileDestinationType = "FHIR-SDC-FormReceiver"
)
