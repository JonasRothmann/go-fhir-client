// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/TestReport
type TestReport struct {
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
	// External identifier
	Identifier *Identifier `json:"identifier,omitempty"`
	// Informal name of the executed TestReport
	Name *string `json:"name,omitempty"`
	// completed | in-progress | waiting | stopped | entered-in-error
	Status custom.Code `json:"status"`
	// Canonical URL to the  version-specific TestScript that was executed to produce this TestReport
	TestScript custom.Canonical[TestScript] `json:"testScript"`
	// pass | fail | pending
	Result custom.Code `json:"result"`
	// The final score (percentage of tests passed) resulting from the execution of the TestScript
	Score *json.Number `json:"score,omitempty"`
	// Name of the tester producing this report (Organization or individual)
	Tester *string `json:"tester,omitempty"`
	// When the TestScript was executed and this TestReport was generated
	Issued *custom.DateTime `json:"issued,omitempty"`
	// A participant in the test execution, either the execution engine, a client, or a server
	Participant []TestReportParticipant `json:"participant,omitempty"`
	// The results of the series of required setup operations before the tests were executed
	Setup *TestReportSetup `json:"setup,omitempty"`
	// A test executed from the test script
	Test []TestReportTest `json:"test,omitempty"`
	// The results of running the series of required clean up steps
	Teardown *TestReportTeardown `json:"teardown,omitempty"`
}

type TestReportSetup struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// A setup operation or assert that was executed
	Action []TestReportSetupAction `json:"action"`
}

type TestReportSetupActionOperation struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// pass | skip | fail | warning | error
	Result custom.Code `json:"result"`
	// A message associated with the result
	Message *custom.Markdown `json:"message,omitempty"`
	// A link to further details on the result
	Detail *custom.URI `json:"detail,omitempty"`
}

type TestReportSetupActionAssert struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// pass | skip | fail | warning | error
	Result custom.Code `json:"result"`
	// A message associated with the result
	Message *custom.Markdown `json:"message,omitempty"`
	// A link to further details on the result
	Detail *string `json:"detail,omitempty"`
	// Links or references to the testing requirements
	Requirement []TestReportSetupActionAssertRequirement `json:"requirement,omitempty"`
}

type TestReportSetupActionAssertRequirement struct {
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

type TestReportTestAction struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The operation performed
	Operation *interface{} `json:"operation,omitempty"`
	// The assertion performed
	Assert *interface{} `json:"assert,omitempty"`
}

type TestReportTeardownAction struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The teardown operation performed
	Operation interface{} `json:"operation"`
}

type TestReportParticipant struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// test-engine | client | server
	Type custom.Code `json:"type"`
	// The uri of the participant. An absolute URL is preferred
	Uri custom.URI `json:"uri"`
	// The display name of the participant
	Display *string `json:"display,omitempty"`
}

type TestReportSetupAction struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The operation to perform
	Operation *TestReportSetupActionOperation `json:"operation,omitempty"`
	// The assertion to perform
	Assert *TestReportSetupActionAssert `json:"assert,omitempty"`
}

type TestReportTest struct {
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
	// A test operation or assert that was performed
	Action []TestReportTestAction `json:"action"`
}

type TestReportTeardown struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// One or more teardown operations performed
	Action []TestReportTeardownAction `json:"action"`
}

type OtherTestReport TestReport

func (t TestReport) ResourceType() string {
	return "TestReport"
}

func (t TestReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherTestReport
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherTestReport: OtherTestReport(t), ResourceType: t.ResourceType()})
}

func UnmarshalTestReport(b []byte) (res TestReport, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
