// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/TestPlan
type TestPlanTestCase struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Sequence of test case in the test plan
	Sequence *int32 `json:"sequence,omitempty"`
	// The scope or artifact covered by the case
	Scope []custom.Reference[gofhirclient.Resource] `json:"scope,omitempty"`
	// Required criteria to execute the test case
	Dependency []TestPlanTestCaseDependency `json:"dependency,omitempty"`
	// The actual test to be executed
	TestRun []TestPlanTestCaseTestRun `json:"testRun,omitempty"`
	// The test data used in the test case
	TestData []TestPlanTestCaseTestData `json:"testData,omitempty"`
	// Test assertions or expectations
	Assertion []TestPlanTestCaseAssertion `json:"assertion,omitempty"`
}

type TestPlanTestCaseDependency struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Description of the criteria
	Description *custom.Markdown `json:"description,omitempty"`
	// Link to predecessor test plans
	Predecessor *custom.Reference[gofhirclient.Resource] `json:"predecessor,omitempty"`
}

type TestPlanTestCaseTestRun struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The narrative description of the tests
	Narrative *custom.Markdown `json:"narrative,omitempty"`
	// The test cases in a structured language e.g. gherkin, Postman, or FHIR TestScript
	Script *TestPlanTestCaseTestRunScript `json:"script,omitempty"`
}

type TestPlanTestCaseTestRunScript struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The language for the test cases e.g. 'gherkin', 'testscript'
	Language *CodeableConcept `json:"language,omitempty"`
	// The actual content of the cases - references to TestScripts or externally defined content
	SourceString *string `json:"sourceString,omitempty"`
	// The actual content of the cases - references to TestScripts or externally defined content
	SourceReference *custom.Reference[gofhirclient.Resource] `json:"sourceReference,omitempty"`
}

type TestPlanTestCaseTestData struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The type of test data description, e.g. 'synthea'
	Type Coding `json:"type"`
	// The actual test resources when they exist
	Content *custom.Reference[gofhirclient.Resource] `json:"content,omitempty"`
	// Pointer to a definition of test resources - narrative or structured e.g. synthetic data generation, etc
	SourceString *string `json:"sourceString,omitempty"`
	// Pointer to a definition of test resources - narrative or structured e.g. synthetic data generation, etc
	SourceReference *custom.Reference[gofhirclient.Resource] `json:"sourceReference,omitempty"`
}

type TestPlanTestCaseAssertion struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Assertion type - for example 'informative' or 'required'
	Type []CodeableConcept `json:"type,omitempty"`
	// The focus or object of the assertion
	Object []custom.CodeableReference[gofhirclient.Resource] `json:"object,omitempty"`
	// The actual result assertion
	Result []custom.CodeableReference[gofhirclient.Resource] `json:"result,omitempty"`
}

type TestPlan struct {
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
	// Canonical identifier for this test plan, represented as a URI (globally unique)
	Url *custom.URI `json:"url,omitempty"`
	// Business identifier identifier for the test plan
	Identifier []Identifier `json:"identifier,omitempty"`
	// Business version of the test plan
	Version *string `json:"version,omitempty"`
	// How to compare versions
	VersionAlgorithmString *string `json:"versionAlgorithmString,omitempty"`
	// How to compare versions
	VersionAlgorithmCoding *Coding `json:"versionAlgorithmCoding,omitempty"`
	// Name for this test plan (computer friendly)
	Name *string `json:"name,omitempty"`
	// Name for this test plan (human friendly)
	Title *string `json:"title,omitempty"`
	// draft | active | retired | unknown
	Status custom.Code `json:"status"`
	// For testing purposes, not real usage
	Experimental *bool `json:"experimental,omitempty"`
	// Date last changed
	Date *custom.DateTime `json:"date,omitempty"`
	// Name of the publisher/steward (organization or individual)
	Publisher *string `json:"publisher,omitempty"`
	// Contact details for the publisher
	Contact []ContactDetail `json:"contact,omitempty"`
	// Natural language description of the test plan
	Description *custom.Markdown `json:"description,omitempty"`
	// The context that the content is intended to support
	UseContext []UsageContext `json:"useContext,omitempty"`
	// Intended jurisdiction where the test plan applies (if applicable)
	Jurisdiction []CodeableConcept `json:"jurisdiction,omitempty"`
	// Why this test plan is defined
	Purpose *custom.Markdown `json:"purpose,omitempty"`
	// Use and/or publishing restrictions
	Copyright *custom.Markdown `json:"copyright,omitempty"`
	// Copyright holder and year(s)
	CopyrightLabel *string `json:"copyrightLabel,omitempty"`
	// The category of the Test Plan - can be acceptance, unit, performance
	Category []CodeableConcept `json:"category,omitempty"`
	// What is being tested with this Test Plan - a conformance resource, or narrative criteria, or an external reference
	Scope []custom.Reference[gofhirclient.Resource] `json:"scope,omitempty"`
	// A description of test tools to be used in the test plan - narrative for now
	TestTools *custom.Markdown `json:"testTools,omitempty"`
	// The required criteria to execute the test plan - e.g. preconditions, previous tests
	Dependency []TestPlanDependency `json:"dependency,omitempty"`
	// The threshold or criteria for the test plan to be considered successfully executed - narrative
	ExitCriteria *custom.Markdown `json:"exitCriteria,omitempty"`
	// The test cases that constitute this plan
	TestCase []TestPlanTestCase `json:"testCase,omitempty"`
}

type TestPlanDependency struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Description of the dependency criterium
	Description *custom.Markdown `json:"description,omitempty"`
	// Link to predecessor test plans
	Predecessor *custom.Reference[gofhirclient.Resource] `json:"predecessor,omitempty"`
}

type OtherTestPlan TestPlan

func (t TestPlan) ResourceType() string {
	return "TestPlan"
}

func (t TestPlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherTestPlan
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherTestPlan: OtherTestPlan(t), ResourceType: t.ResourceType()})
}

func UnmarshalTestPlan(b []byte) (res TestPlan, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
