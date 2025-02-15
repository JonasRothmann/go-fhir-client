// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/TestScript
type TestScriptMetadataCapability struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Are the capabilities required?
	Required bool `bson:"required" json:"required"`
	// Are the capabilities validated?
	Validated bool `bson:"validated" json:"validated"`
	// The expected capabilities of the server
	Description *string `bson:"description,omitempty" json:"description,omitempty"`
	// Which origin server these requirements apply to
	Origin []int32 `bson:"origin,omitempty" json:"origin,omitempty"`
	// Which server these requirements apply to
	Destination *int32 `bson:"destination,omitempty" json:"destination,omitempty"`
	// Links to the FHIR specification
	Link []custom.URI `bson:"link,omitempty" json:"link,omitempty"`
	// Required Capability Statement
	Capabilities custom.Canonical[CapabilityStatement] `bson:"capabilities" json:"capabilities"`
}

type TestScriptSetup struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// A setup operation or assert to perform
	Action []TestScriptSetupAction `bson:"action" json:"action"`
}

type TestScriptDestination struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The index of the abstract destination server starting at 1
	Index int32 `bson:"index" json:"index"`
	// FHIR-Server | FHIR-SDC-FormManager | FHIR-SDC-FormReceiver | FHIR-SDC-FormProcessor
	Profile Coding `bson:"profile" json:"profile"`
	// The url path of the destination server
	Url *custom.URL `bson:"url,omitempty" json:"url,omitempty"`
}

type TestScriptScope struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The specific conformance artifact being tested
	Artifact custom.Canonical[Resource] `bson:"artifact" json:"artifact"`
	// required | optional | strict
	Conformance *CodeableConcept `bson:"conformance,omitempty" json:"conformance,omitempty"`
	// unit | integration | production
	Phase *CodeableConcept `bson:"phase,omitempty" json:"phase,omitempty"`
}

type TestScriptSetupActionOperationRequestHeader struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// HTTP header field name
	Field string `bson:"field" json:"field"`
	// HTTP headerfield value
	Value string `bson:"value" json:"value"`
}

type TestScriptSetupActionAssertRequirement struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Link or reference to the testing requirement
	Link *custom.URI `bson:"link,omitempty" json:"link,omitempty"`
}

type TestScriptTeardownAction struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The teardown operation to perform
	Operation interface{} `bson:"operation" json:"operation"`
}

type TestScriptOrigin struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The index of the abstract origin server starting at 1
	Index int32 `bson:"index" json:"index"`
	// FHIR-Client | FHIR-SDC-FormFiller
	Profile Coding `bson:"profile" json:"profile"`
	// The url path of the origin server
	Url *custom.URL `bson:"url,omitempty" json:"url,omitempty"`
}

type TestScriptMetadata struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Links to the FHIR specification
	Link []TestScriptMetadataLink `bson:"link,omitempty" json:"link,omitempty"`
	// Capabilities  that are assumed to function correctly on the FHIR server being tested
	Capability []TestScriptMetadataCapability `bson:"capability" json:"capability"`
}

type TestScriptSetupAction struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The setup operation to perform
	Operation *TestScriptSetupActionOperation `bson:"operation,omitempty" json:"operation,omitempty"`
	// The assertion to perform
	Assert *TestScriptSetupActionAssert `bson:"assert,omitempty" json:"assert,omitempty"`
}

type TestScriptSetupActionOperation struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The operation code type that will be executed
	Type *Coding `bson:"type,omitempty" json:"type,omitempty"`
	// Resource type
	Resource *custom.URI `bson:"resource,omitempty" json:"resource,omitempty"`
	// Tracking/logging operation label
	Label *string `bson:"label,omitempty" json:"label,omitempty"`
	// Tracking/reporting operation description
	Description *string `bson:"description,omitempty" json:"description,omitempty"`
	// Mime type to accept in the payload of the response, with charset etc
	Accept *custom.Code `bson:"accept,omitempty" json:"accept,omitempty"`
	// Mime type of the request payload contents, with charset etc
	ContentType *custom.Code `bson:"contentType,omitempty" json:"contentType,omitempty"`
	// Server responding to the request
	Destination *int32 `bson:"destination,omitempty" json:"destination,omitempty"`
	// Whether or not to send the request url in encoded format
	EncodeRequestUrl bool `bson:"encodeRequestUrl" json:"encodeRequestUrl"`
	// delete | get | options | patch | post | put | head
	Method *custom.Code `bson:"method,omitempty" json:"method,omitempty"`
	// Server initiating the request
	Origin *int32 `bson:"origin,omitempty" json:"origin,omitempty"`
	// Explicitly defined path parameters
	Params *string `bson:"params,omitempty" json:"params,omitempty"`
	// Each operation can have one or more header elements
	RequestHeader []TestScriptSetupActionOperationRequestHeader `bson:"requestHeader,omitempty" json:"requestHeader,omitempty"`
	// Fixture Id of mapped request
	RequestId *custom.ID `bson:"requestId,omitempty" json:"requestId,omitempty"`
	// Fixture Id of mapped response
	ResponseId *custom.ID `bson:"responseId,omitempty" json:"responseId,omitempty"`
	// Fixture Id of body for PUT and POST requests
	SourceId *custom.ID `bson:"sourceId,omitempty" json:"sourceId,omitempty"`
	// Id of fixture used for extracting the [id],  [type], and [vid] for GET requests
	TargetId *custom.ID `bson:"targetId,omitempty" json:"targetId,omitempty"`
	// Request URL
	Url *string `bson:"url,omitempty" json:"url,omitempty"`
}

type TestScriptTest struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Tracking/logging name of this test
	Name *string `bson:"name,omitempty" json:"name,omitempty"`
	// Tracking/reporting short description of the test
	Description *string `bson:"description,omitempty" json:"description,omitempty"`
	// A test operation or assert to perform
	Action []TestScriptTestAction `bson:"action" json:"action"`
}

type TestScriptTestAction struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The setup operation to perform
	Operation *interface{} `bson:"operation,omitempty" json:"operation,omitempty"`
	// The setup assertion to perform
	Assert *interface{} `bson:"assert,omitempty" json:"assert,omitempty"`
}

type TestScriptTeardown struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// One or more teardown operations to perform
	Action []TestScriptTeardownAction `bson:"action" json:"action"`
}

type TestScript struct {
	// Logical id of this artifact
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Metadata about the resource
	Meta *Meta `bson:"meta,omitempty" json:"meta,omitempty"`
	// A set of rules under which this content was created
	ImplicitRules *custom.URI `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	// Language of the resource content
	Language *custom.Code `bson:"language,omitempty" json:"language,omitempty"`
	// Text summary of the resource, for human interpretation
	Text *Narrative `bson:"text,omitempty" json:"text,omitempty"`
	// Contained, inline Resources
	Contained []Resource `bson:"contained,omitempty" json:"contained,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Canonical identifier for this test script, represented as a URI (globally unique)
	Url *custom.URI `bson:"url,omitempty" json:"url,omitempty"`
	// Additional identifier for the test script
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Business version of the test script
	Version *string `bson:"version,omitempty" json:"version,omitempty"`
	// How to compare versions
	VersionAlgorithm *string `bson:"versionAlgorithm,omitempty" json:"versionAlgorithm,omitempty"`
	// Name for this test script (computer friendly)
	Name string `bson:"name" json:"name"`
	// Name for this test script (human friendly)
	Title *string `bson:"title,omitempty" json:"title,omitempty"`
	// draft | active | retired | unknown
	Status custom.Code `bson:"status" json:"status"`
	// For testing purposes, not real usage
	Experimental *bool `bson:"experimental,omitempty" json:"experimental,omitempty"`
	// Date last changed
	Date *custom.DateTime `bson:"date,omitempty" json:"date,omitempty"`
	// Name of the publisher/steward (organization or individual)
	Publisher *string `bson:"publisher,omitempty" json:"publisher,omitempty"`
	// Contact details for the publisher
	Contact []ContactDetail `bson:"contact,omitempty" json:"contact,omitempty"`
	// Natural language description of the test script
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// The context that the content is intended to support
	UseContext []UsageContext `bson:"useContext,omitempty" json:"useContext,omitempty"`
	// Intended jurisdiction for test script (if applicable)
	Jurisdiction []CodeableConcept `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	// Why this test script is defined
	Purpose *custom.Markdown `bson:"purpose,omitempty" json:"purpose,omitempty"`
	// Use and/or publishing restrictions
	Copyright *custom.Markdown `bson:"copyright,omitempty" json:"copyright,omitempty"`
	// Copyright holder and year(s)
	CopyrightLabel *string `bson:"copyrightLabel,omitempty" json:"copyrightLabel,omitempty"`
	// An abstract server representing a client or sender in a message exchange
	Origin []TestScriptOrigin `bson:"origin,omitempty" json:"origin,omitempty"`
	// An abstract server representing a destination or receiver in a message exchange
	Destination []TestScriptDestination `bson:"destination,omitempty" json:"destination,omitempty"`
	// Required capability that is assumed to function correctly on the FHIR server being tested
	Metadata *TestScriptMetadata `bson:"metadata,omitempty" json:"metadata,omitempty"`
	// Indication of the artifact(s) that are tested by this test case
	Scope []TestScriptScope `bson:"scope,omitempty" json:"scope,omitempty"`
	// Fixture in the test script - by reference (uri)
	Fixture []TestScriptFixture `bson:"fixture,omitempty" json:"fixture,omitempty"`
	// Reference of the validation profile
	Profile []custom.Canonical[StructureDefinition] `bson:"profile,omitempty" json:"profile,omitempty"`
	// Placeholder for evaluated elements
	Variable []TestScriptVariable `bson:"variable,omitempty" json:"variable,omitempty"`
	// A series of required setup operations before tests are executed
	Setup *TestScriptSetup `bson:"setup,omitempty" json:"setup,omitempty"`
	// A test in this script
	Test []TestScriptTest `bson:"test,omitempty" json:"test,omitempty"`
	// A series of required clean up steps
	Teardown *TestScriptTeardown `bson:"teardown,omitempty" json:"teardown,omitempty"`
}

type TestScriptMetadataLink struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// URL to the specification
	Url custom.URI `bson:"url" json:"url"`
	// Short description
	Description *string `bson:"description,omitempty" json:"description,omitempty"`
}

type TestScriptFixture struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Whether or not to implicitly create the fixture during setup
	Autocreate bool `bson:"autocreate" json:"autocreate"`
	// Whether or not to implicitly delete the fixture during teardown
	Autodelete bool `bson:"autodelete" json:"autodelete"`
	// Reference of the resource
	Resource *custom.Reference[Resource] `bson:"resource,omitempty" json:"resource,omitempty"`
}

type TestScriptVariable struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Descriptive name for this variable
	Name string `bson:"name" json:"name"`
	// Default, hard-coded, or user-defined value for this variable
	DefaultValue *string `bson:"defaultValue,omitempty" json:"defaultValue,omitempty"`
	// Natural language description of the variable
	Description *string `bson:"description,omitempty" json:"description,omitempty"`
	// The FHIRPath expression against the fixture body
	Expression *string `bson:"expression,omitempty" json:"expression,omitempty"`
	// HTTP header field name for source
	HeaderField *string `bson:"headerField,omitempty" json:"headerField,omitempty"`
	// Hint help text for default value to enter
	Hint *string `bson:"hint,omitempty" json:"hint,omitempty"`
	// XPath or JSONPath against the fixture body
	Path *string `bson:"path,omitempty" json:"path,omitempty"`
	// Fixture Id of source expression or headerField within this variable
	SourceId *custom.ID `bson:"sourceId,omitempty" json:"sourceId,omitempty"`
}

type TestScriptSetupActionAssert struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Tracking/logging assertion label
	Label *string `bson:"label,omitempty" json:"label,omitempty"`
	// Tracking/reporting assertion description
	Description *string `bson:"description,omitempty" json:"description,omitempty"`
	// response | request
	Direction *custom.Code `bson:"direction,omitempty" json:"direction,omitempty"`
	// Id of the source fixture to be evaluated
	CompareToSourceId *string `bson:"compareToSourceId,omitempty" json:"compareToSourceId,omitempty"`
	// The FHIRPath expression to evaluate against the source fixture
	CompareToSourceExpression *string `bson:"compareToSourceExpression,omitempty" json:"compareToSourceExpression,omitempty"`
	// XPath or JSONPath expression to evaluate against the source fixture
	CompareToSourcePath *string `bson:"compareToSourcePath,omitempty" json:"compareToSourcePath,omitempty"`
	// Mime type to compare against the 'Content-Type' header
	ContentType *custom.Code `bson:"contentType,omitempty" json:"contentType,omitempty"`
	// fail | pass | skip | stop
	DefaultManualCompletion *custom.Code `bson:"defaultManualCompletion,omitempty" json:"defaultManualCompletion,omitempty"`
	// The FHIRPath expression to be evaluated
	Expression *string `bson:"expression,omitempty" json:"expression,omitempty"`
	// HTTP header field name
	HeaderField *string `bson:"headerField,omitempty" json:"headerField,omitempty"`
	// Fixture Id of minimum content resource
	MinimumId *string `bson:"minimumId,omitempty" json:"minimumId,omitempty"`
	// Perform validation on navigation links?
	NavigationLinks *bool `bson:"navigationLinks,omitempty" json:"navigationLinks,omitempty"`
	// equals | notEquals | in | notIn | greaterThan | lessThan | empty | notEmpty | contains | notContains | eval | manualEval
	Operator *custom.Code `bson:"operator,omitempty" json:"operator,omitempty"`
	// XPath or JSONPath expression
	Path *string `bson:"path,omitempty" json:"path,omitempty"`
	// delete | get | options | patch | post | put | head
	RequestMethod *custom.Code `bson:"requestMethod,omitempty" json:"requestMethod,omitempty"`
	// Request URL comparison value
	RequestUrl *string `bson:"requestUrl,omitempty" json:"requestUrl,omitempty"`
	// Resource type
	Resource *custom.URI `bson:"resource,omitempty" json:"resource,omitempty"`
	// continue | switchingProtocols | okay | created | accepted | nonAuthoritativeInformation | noContent | resetContent | partialContent | multipleChoices | movedPermanently | found | seeOther | notModified | useProxy | temporaryRedirect | permanentRedirect | badRequest | unauthorized | paymentRequired | forbidden | notFound | methodNotAllowed | notAcceptable | proxyAuthenticationRequired | requestTimeout | conflict | gone | lengthRequired | preconditionFailed | contentTooLarge | uriTooLong | unsupportedMediaType | rangeNotSatisfiable | expectationFailed | misdirectedRequest | unprocessableContent | upgradeRequired | internalServerError | notImplemented | badGateway | serviceUnavailable | gatewayTimeout | httpVersionNotSupported
	Response *custom.Code `bson:"response,omitempty" json:"response,omitempty"`
	// HTTP response code to test
	ResponseCode *string `bson:"responseCode,omitempty" json:"responseCode,omitempty"`
	// Fixture Id of source expression or headerField
	SourceId *custom.ID `bson:"sourceId,omitempty" json:"sourceId,omitempty"`
	// If this assert fails, will the current test execution stop?
	StopTestOnFail bool `bson:"stopTestOnFail" json:"stopTestOnFail"`
	// Profile Id of validation profile reference
	ValidateProfileId *custom.ID `bson:"validateProfileId,omitempty" json:"validateProfileId,omitempty"`
	// The value to compare to
	Value *string `bson:"value,omitempty" json:"value,omitempty"`
	// Will this assert produce a warning only on error?
	WarningOnly bool `bson:"warningOnly" json:"warningOnly"`
	// Links or references to the testing requirements
	Requirement []TestScriptSetupActionAssertRequirement `bson:"requirement,omitempty" json:"requirement,omitempty"`
}

func (t TestScript) ResourceType() string {
	return "StructureDefinition"
}
