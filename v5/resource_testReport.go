// GENERATED CODE – DO NOT EDIT!

package v5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/TestReport
type TestReportSetupActionAssert struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// pass | skip | fail | warning | error
	Result custom.Code `bson:"result" json:"result"`
	// A message associated with the result
	Message *custom.Markdown `bson:"message,omitempty" json:"message,omitempty"`
	// A link to further details on the result
	Detail *string `bson:"detail,omitempty" json:"detail,omitempty"`
	// Links or references to the testing requirements
	Requirement []TestReportSetupActionAssertRequirement `bson:"requirement,omitempty" json:"requirement,omitempty"`
}

type TestReportSetupActionAssertRequirement struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Link or reference to the testing requirement
	Link *custom.URI `bson:"link,omitempty" json:"link,omitempty"`
}

type TestReportTeardown struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// One or more teardown operations performed
	Action []TestReportTeardownAction `bson:"action" json:"action"`
}

type TestReport struct {
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
	// External identifier
	Identifier *Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Informal name of the executed TestReport
	Name *string `bson:"name,omitempty" json:"name,omitempty"`
	// completed | in-progress | waiting | stopped | entered-in-error
	Status custom.Code `bson:"status" json:"status"`
	// Canonical URL to the  version-specific TestScript that was executed to produce this TestReport
	TestScript custom.Canonical[TestScript] `bson:"testScript" json:"testScript"`
	// pass | fail | pending
	Result custom.Code `bson:"result" json:"result"`
	// The final score (percentage of tests passed) resulting from the execution of the TestScript
	Score *json.Number `bson:"score,omitempty" json:"score,omitempty"`
	// Name of the tester producing this report (Organization or individual)
	Tester *string `bson:"tester,omitempty" json:"tester,omitempty"`
	// When the TestScript was executed and this TestReport was generated
	Issued *custom.DateTime `bson:"issued,omitempty" json:"issued,omitempty"`
	// A participant in the test execution, either the execution engine, a client, or a server
	Participant []TestReportParticipant `bson:"participant,omitempty" json:"participant,omitempty"`
	// The results of the series of required setup operations before the tests were executed
	Setup *TestReportSetup `bson:"setup,omitempty" json:"setup,omitempty"`
	// A test executed from the test script
	Test []TestReportTest `bson:"test,omitempty" json:"test,omitempty"`
	// The results of running the series of required clean up steps
	Teardown *TestReportTeardown `bson:"teardown,omitempty" json:"teardown,omitempty"`
}

type TestReportSetupAction struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The operation to perform
	Operation *TestReportSetupActionOperation `bson:"operation,omitempty" json:"operation,omitempty"`
	// The assertion to perform
	Assert *TestReportSetupActionAssert `bson:"assert,omitempty" json:"assert,omitempty"`
}

type TestReportSetupActionOperation struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// pass | skip | fail | warning | error
	Result custom.Code `bson:"result" json:"result"`
	// A message associated with the result
	Message *custom.Markdown `bson:"message,omitempty" json:"message,omitempty"`
	// A link to further details on the result
	Detail *custom.URI `bson:"detail,omitempty" json:"detail,omitempty"`
}

type TestReportTest struct {
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
	// A test operation or assert that was performed
	Action []TestReportTestAction `bson:"action" json:"action"`
}

type TestReportTestAction struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The operation performed
	Operation *interface{} `bson:"operation,omitempty" json:"operation,omitempty"`
	// The assertion performed
	Assert *interface{} `bson:"assert,omitempty" json:"assert,omitempty"`
}

type TestReportTeardownAction struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The teardown operation performed
	Operation interface{} `bson:"operation" json:"operation"`
}

type TestReportParticipant struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// test-engine | client | server
	Type custom.Code `bson:"type" json:"type"`
	// The uri of the participant. An absolute URL is preferred
	Uri custom.URI `bson:"uri" json:"uri"`
	// The display name of the participant
	Display *string `bson:"display,omitempty" json:"display,omitempty"`
}

type TestReportSetup struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// A setup operation or assert that was executed
	Action []TestReportSetupAction `bson:"action" json:"action"`
}

func (t TestReport) ResourceType() string {
	return "StructureDefinition"
}
