// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/shareabletestscript
type ShareableTestScriptVariable struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Descriptive name for this variable
	Name string `json:"name"`
	// Default, hard-coded, or user-defined value for this variable
	DefaultValue *string `json:"defaultValue,omitempty"`
	// Natural language description of the variable
	Description *string `json:"description,omitempty"`
	// The FHIRPath expression against the fixture body
	Expression *string `json:"expression,omitempty"`
	// HTTP header field name for source
	HeaderField *string `json:"headerField,omitempty"`
	// Hint help text for default value to enter
	Hint *string `json:"hint,omitempty"`
	// XPath or JSONPath against the fixture body
	Path *string `json:"path,omitempty"`
	// Fixture Id of source expression or headerField within this variable
	SourceId *custom.ID `json:"sourceId,omitempty"`
}

type ShareableTestScriptDestination struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The index of the abstract destination server starting at 1
	Index int32 `json:"index"`
	// FHIR-Server | FHIR-SDC-FormManager | FHIR-SDC-FormReceiver | FHIR-SDC-FormProcessor
	Profile Coding `json:"profile"`
	// The url path of the destination server
	Url *custom.URL `json:"url,omitempty"`
}

type ShareableTestScriptMetadataCapability struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Are the capabilities required?
	Required bool `json:"required"`
	// Are the capabilities validated?
	Validated bool `json:"validated"`
	// The expected capabilities of the server
	Description *string `json:"description,omitempty"`
	// Which origin server these requirements apply to
	Origin []int32 `json:"origin,omitempty"`
	// Which server these requirements apply to
	Destination *int32 `json:"destination,omitempty"`
	// Links to the FHIR specification
	Link []custom.URI `json:"link,omitempty"`
	// Required Capability Statement
	Capabilities custom.Canonical[CapabilityStatement] `json:"capabilities"`
}

type ShareableTestScriptSetup struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// A setup operation or assert to perform
	Action []ShareableTestScriptSetupAction `json:"action"`
}

type ShareableTestScriptSetupAction struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The setup operation to perform
	Operation *ShareableTestScriptSetupActionOperation `json:"operation,omitempty"`
	// The assertion to perform
	Assert *ShareableTestScriptSetupActionAssert `json:"assert,omitempty"`
}

type ShareableTestScriptTeardownAction struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The teardown operation to perform
	Operation interface{} `json:"operation"`
}

type ShareableTestScriptOrigin struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The index of the abstract origin server starting at 1
	Index int32 `json:"index"`
	// FHIR-Client | FHIR-SDC-FormFiller
	Profile Coding `json:"profile"`
	// The url path of the origin server
	Url *custom.URL `json:"url,omitempty"`
}

type ShareableTestScriptMetadata struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Links to the FHIR specification
	Link []ShareableTestScriptMetadataLink `json:"link,omitempty"`
	// Capabilities  that are assumed to function correctly on the FHIR server being tested
	Capability []ShareableTestScriptMetadataCapability `json:"capability"`
}

type ShareableTestScriptScope struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The specific conformance artifact being tested
	Artifact custom.Canonical[Resource] `json:"artifact"`
	// required | optional | strict
	Conformance *CodeableConcept `json:"conformance,omitempty"`
	// unit | integration | production
	Phase *CodeableConcept `json:"phase,omitempty"`
}

type ShareableTestScriptFixture struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Whether or not to implicitly create the fixture during setup
	Autocreate bool `json:"autocreate"`
	// Whether or not to implicitly delete the fixture during teardown
	Autodelete bool `json:"autodelete"`
	// Reference of the resource
	Resource *custom.Reference[Resource] `json:"resource,omitempty"`
}

type ShareableTestScriptSetupActionOperationRequestHeader struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// HTTP header field name
	Field string `json:"field"`
	// HTTP headerfield value
	Value string `json:"value"`
}

type ShareableTestScriptSetupActionAssert struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Tracking/logging assertion label
	Label *string `json:"label,omitempty"`
	// Tracking/reporting assertion description
	Description *string `json:"description,omitempty"`
	// response | request
	Direction *custom.Code `json:"direction,omitempty"`
	// Id of the source fixture to be evaluated
	CompareToSourceId *string `json:"compareToSourceId,omitempty"`
	// The FHIRPath expression to evaluate against the source fixture
	CompareToSourceExpression *string `json:"compareToSourceExpression,omitempty"`
	// XPath or JSONPath expression to evaluate against the source fixture
	CompareToSourcePath *string `json:"compareToSourcePath,omitempty"`
	// Mime type to compare against the 'Content-Type' header
	ContentType *custom.Code `json:"contentType,omitempty"`
	// fail | pass | skip | stop
	DefaultManualCompletion *custom.Code `json:"defaultManualCompletion,omitempty"`
	// The FHIRPath expression to be evaluated
	Expression *string `json:"expression,omitempty"`
	// HTTP header field name
	HeaderField *string `json:"headerField,omitempty"`
	// Fixture Id of minimum content resource
	MinimumId *string `json:"minimumId,omitempty"`
	// Perform validation on navigation links?
	NavigationLinks *bool `json:"navigationLinks,omitempty"`
	// equals | notEquals | in | notIn | greaterThan | lessThan | empty | notEmpty | contains | notContains | eval | manualEval
	Operator *custom.Code `json:"operator,omitempty"`
	// XPath or JSONPath expression
	Path *string `json:"path,omitempty"`
	// delete | get | options | patch | post | put | head
	RequestMethod *custom.Code `json:"requestMethod,omitempty"`
	// Request URL comparison value
	RequestUrl *string `json:"requestUrl,omitempty"`
	// Resource type
	Resource *custom.URI `json:"resource,omitempty"`
	// continue | switchingProtocols | okay | created | accepted | nonAuthoritativeInformation | noContent | resetContent | partialContent | multipleChoices | movedPermanently | found | seeOther | notModified | useProxy | temporaryRedirect | permanentRedirect | badRequest | unauthorized | paymentRequired | forbidden | notFound | methodNotAllowed | notAcceptable | proxyAuthenticationRequired | requestTimeout | conflict | gone | lengthRequired | preconditionFailed | contentTooLarge | uriTooLong | unsupportedMediaType | rangeNotSatisfiable | expectationFailed | misdirectedRequest | unprocessableContent | upgradeRequired | internalServerError | notImplemented | badGateway | serviceUnavailable | gatewayTimeout | httpVersionNotSupported
	Response *custom.Code `json:"response,omitempty"`
	// HTTP response code to test
	ResponseCode *string `json:"responseCode,omitempty"`
	// Fixture Id of source expression or headerField
	SourceId *custom.ID `json:"sourceId,omitempty"`
	// If this assert fails, will the current test execution stop?
	StopTestOnFail bool `json:"stopTestOnFail"`
	// Profile Id of validation profile reference
	ValidateProfileId *custom.ID `json:"validateProfileId,omitempty"`
	// The value to compare to
	Value *string `json:"value,omitempty"`
	// Will this assert produce a warning only on error?
	WarningOnly bool `json:"warningOnly"`
	// Links or references to the testing requirements
	Requirement []ShareableTestScriptSetupActionAssertRequirement `json:"requirement,omitempty"`
}

type ShareableTestScriptSetupActionAssertRequirement struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Link or reference to the testing requirement
	LinkUri *custom.URI `json:"linkUri,omitempty"`
	// Link or reference to the testing requirement
	LinkCanonical *custom.Canonical[Requirements] `json:"linkCanonical,omitempty"`
}

type ShareableTestScript struct {
	// Logical id of this artifact
	Id *string `json:"id,omitempty"`
	// Metadata about the resource
	Meta *Meta `json:"meta,omitempty"`
	// A set of rules under which this content was created
	ImplicitRules *custom.URI `json:"implicitRules,omitempty"`
	// Language of the resource content
	Language *custom.Code `json:"language,omitempty"`
	// Text summary of the resource, for human interpretation
	Text *Narrative `json:"text,omitempty"`
	// Contained, inline Resources
	Contained []Resource `json:"contained,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Canonical identifier for this test script, represented as a URI (globally unique)
	Url custom.URI `json:"url"`
	// Additional identifier for the test script
	Identifier []Identifier `json:"identifier,omitempty"`
	// Business version of the test script
	Version string `json:"version"`
	// How to compare versions
	VersionAlgorithmString *string `json:"versionAlgorithmString,omitempty"`
	// How to compare versions
	VersionAlgorithmCoding *Coding `json:"versionAlgorithmCoding,omitempty"`
	// Name for this test script (computer friendly)
	Name string `json:"name"`
	// Name for this test script (human friendly)
	Title *string `json:"title,omitempty"`
	// draft | active | retired | unknown
	Status custom.Code `json:"status"`
	// For testing purposes, not real usage
	Experimental bool `json:"experimental"`
	// Date last changed
	Date *custom.DateTime `json:"date,omitempty"`
	// Name of the publisher/steward (organization or individual)
	Publisher string `json:"publisher"`
	// Contact details for the publisher
	Contact []ContactDetail `json:"contact,omitempty"`
	// Natural language description of the test script
	Description custom.Markdown `json:"description"`
	// The context that the content is intended to support
	UseContext []UsageContext `json:"useContext,omitempty"`
	// Intended jurisdiction for test script (if applicable)
	Jurisdiction []CodeableConcept `json:"jurisdiction,omitempty"`
	// Why this test script is defined
	Purpose *custom.Markdown `json:"purpose,omitempty"`
	// Use and/or publishing restrictions
	Copyright *custom.Markdown `json:"copyright,omitempty"`
	// Copyright holder and year(s)
	CopyrightLabel *string `json:"copyrightLabel,omitempty"`
	// An abstract server representing a client or sender in a message exchange
	Origin []ShareableTestScriptOrigin `json:"origin,omitempty"`
	// An abstract server representing a destination or receiver in a message exchange
	Destination []ShareableTestScriptDestination `json:"destination,omitempty"`
	// Required capability that is assumed to function correctly on the FHIR server being tested
	Metadata *ShareableTestScriptMetadata `json:"metadata,omitempty"`
	// Indication of the artifact(s) that are tested by this test case
	Scope []ShareableTestScriptScope `json:"scope,omitempty"`
	// Fixture in the test script - by reference (uri)
	Fixture []ShareableTestScriptFixture `json:"fixture,omitempty"`
	// Reference of the validation profile
	Profile []custom.Canonical[StructureDefinition] `json:"profile,omitempty"`
	// Placeholder for evaluated elements
	Variable []ShareableTestScriptVariable `json:"variable,omitempty"`
	// A series of required setup operations before tests are executed
	Setup *ShareableTestScriptSetup `json:"setup,omitempty"`
	// A test in this script
	Test []ShareableTestScriptTest `json:"test,omitempty"`
	// A series of required clean up steps
	Teardown *ShareableTestScriptTeardown `json:"teardown,omitempty"`
}

type ShareableTestScriptMetadataLink struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// URL to the specification
	Url custom.URI `json:"url"`
	// Short description
	Description *string `json:"description,omitempty"`
}

type ShareableTestScriptSetupActionOperation struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The operation code type that will be executed
	Type *Coding `json:"type,omitempty"`
	// Resource type
	Resource *custom.URI `json:"resource,omitempty"`
	// Tracking/logging operation label
	Label *string `json:"label,omitempty"`
	// Tracking/reporting operation description
	Description *string `json:"description,omitempty"`
	// Mime type to accept in the payload of the response, with charset etc
	Accept *custom.Code `json:"accept,omitempty"`
	// Mime type of the request payload contents, with charset etc
	ContentType *custom.Code `json:"contentType,omitempty"`
	// Server responding to the request
	Destination *int32 `json:"destination,omitempty"`
	// Whether or not to send the request url in encoded format
	EncodeRequestUrl bool `json:"encodeRequestUrl"`
	// delete | get | options | patch | post | put | head
	Method *custom.Code `json:"method,omitempty"`
	// Server initiating the request
	Origin *int32 `json:"origin,omitempty"`
	// Explicitly defined path parameters
	Params *string `json:"params,omitempty"`
	// Each operation can have one or more header elements
	RequestHeader []ShareableTestScriptSetupActionOperationRequestHeader `json:"requestHeader,omitempty"`
	// Fixture Id of mapped request
	RequestId *custom.ID `json:"requestId,omitempty"`
	// Fixture Id of mapped response
	ResponseId *custom.ID `json:"responseId,omitempty"`
	// Fixture Id of body for PUT and POST requests
	SourceId *custom.ID `json:"sourceId,omitempty"`
	// Id of fixture used for extracting the [id],  [type], and [vid] for GET requests
	TargetId *custom.ID `json:"targetId,omitempty"`
	// Request URL
	Url *string `json:"url,omitempty"`
}

type ShareableTestScriptTest struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Tracking/logging name of this test
	Name *string `json:"name,omitempty"`
	// Tracking/reporting short description of the test
	Description *string `json:"description,omitempty"`
	// A test operation or assert to perform
	Action []ShareableTestScriptTestAction `json:"action"`
}

type ShareableTestScriptTestAction struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The setup operation to perform
	Operation *interface{} `json:"operation,omitempty"`
	// The setup assertion to perform
	Assert *interface{} `json:"assert,omitempty"`
}

type ShareableTestScriptTeardown struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// One or more teardown operations to perform
	Action []ShareableTestScriptTeardownAction `json:"action"`
}

type OtherShareableTestScript ShareableTestScript

func (s ShareableTestScript) ResourceType() string {
	return "ShareableTestScript"
}

func (s ShareableTestScript) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherShareableTestScript
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherShareableTestScript: OtherShareableTestScript(s), ResourceType: s.ResourceType()})
}

func UnmarshalShareableTestScript(b []byte) (res ShareableTestScript, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
