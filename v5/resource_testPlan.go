// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/TestPlan
type TestPlanTestCaseDependency struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Description of the criteria
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// Link to predecessor test plans
	Predecessor *custom.Reference `bson:"predecessor,omitempty" json:"predecessor,omitempty"`
}

type TestPlanTestCaseTestRun struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The narrative description of the tests
	Narrative *custom.Markdown `bson:"narrative,omitempty" json:"narrative,omitempty"`
	// The test cases in a structured language e.g. gherkin, Postman, or FHIR TestScript
	Script *TestPlanTestCaseTestRunScript `bson:"script,omitempty" json:"script,omitempty"`
}

type TestPlanTestCaseTestRunScript struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The language for the test cases e.g. 'gherkin', 'testscript'
	Language *CodeableConcept `bson:"language,omitempty" json:"language,omitempty"`
	// The actual content of the cases - references to TestScripts or externally defined content
	Source *string `bson:"source,omitempty" json:"source,omitempty"`
}

type TestPlanTestCaseTestData struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The type of test data description, e.g. 'synthea'
	Type Coding `bson:"type" json:"type"`
	// The actual test resources when they exist
	Content *custom.Reference `bson:"content,omitempty" json:"content,omitempty"`
	// Pointer to a definition of test resources - narrative or structured e.g. synthetic data generation, etc
	Source *string `bson:"source,omitempty" json:"source,omitempty"`
}

type TestPlanTestCaseAssertion struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Assertion type - for example 'informative' or 'required'
	Type []CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// The focus or object of the assertion
	Object []custom.CodeableReference `bson:"object,omitempty" json:"object,omitempty"`
	// The actual result assertion
	Result []custom.CodeableReference `bson:"result,omitempty" json:"result,omitempty"`
}

type TestPlan struct {
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
	// Canonical identifier for this test plan, represented as a URI (globally unique)
	Url *custom.URI `bson:"url,omitempty" json:"url,omitempty"`
	// Business identifier identifier for the test plan
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Business version of the test plan
	Version *string `bson:"version,omitempty" json:"version,omitempty"`
	// How to compare versions
	VersionAlgorithm *string `bson:"versionAlgorithm,omitempty" json:"versionAlgorithm,omitempty"`
	// Name for this test plan (computer friendly)
	Name *string `bson:"name,omitempty" json:"name,omitempty"`
	// Name for this test plan (human friendly)
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
	// Natural language description of the test plan
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// The context that the content is intended to support
	UseContext []UsageContext `bson:"useContext,omitempty" json:"useContext,omitempty"`
	// Intended jurisdiction where the test plan applies (if applicable)
	Jurisdiction []CodeableConcept `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	// Why this test plan is defined
	Purpose *custom.Markdown `bson:"purpose,omitempty" json:"purpose,omitempty"`
	// Use and/or publishing restrictions
	Copyright *custom.Markdown `bson:"copyright,omitempty" json:"copyright,omitempty"`
	// Copyright holder and year(s)
	CopyrightLabel *string `bson:"copyrightLabel,omitempty" json:"copyrightLabel,omitempty"`
	// The category of the Test Plan - can be acceptance, unit, performance
	Category []CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	// What is being tested with this Test Plan - a conformance resource, or narrative criteria, or an external reference
	Scope []custom.Reference `bson:"scope,omitempty" json:"scope,omitempty"`
	// A description of test tools to be used in the test plan - narrative for now
	TestTools *custom.Markdown `bson:"testTools,omitempty" json:"testTools,omitempty"`
	// The required criteria to execute the test plan - e.g. preconditions, previous tests
	Dependency []TestPlanDependency `bson:"dependency,omitempty" json:"dependency,omitempty"`
	// The threshold or criteria for the test plan to be considered successfully executed - narrative
	ExitCriteria *custom.Markdown `bson:"exitCriteria,omitempty" json:"exitCriteria,omitempty"`
	// The test cases that constitute this plan
	TestCase []TestPlanTestCase `bson:"testCase,omitempty" json:"testCase,omitempty"`
}

type TestPlanDependency struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Description of the dependency criterium
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// Link to predecessor test plans
	Predecessor *custom.Reference `bson:"predecessor,omitempty" json:"predecessor,omitempty"`
}

type TestPlanTestCase struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Sequence of test case in the test plan
	Sequence *int32 `bson:"sequence,omitempty" json:"sequence,omitempty"`
	// The scope or artifact covered by the case
	Scope []custom.Reference `bson:"scope,omitempty" json:"scope,omitempty"`
	// Required criteria to execute the test case
	Dependency []TestPlanTestCaseDependency `bson:"dependency,omitempty" json:"dependency,omitempty"`
	// The actual test to be executed
	TestRun []TestPlanTestCaseTestRun `bson:"testRun,omitempty" json:"testRun,omitempty"`
	// The test data used in the test case
	TestData []TestPlanTestCaseTestData `bson:"testData,omitempty" json:"testData,omitempty"`
	// Test assertions or expectations
	Assertion []TestPlanTestCaseAssertion `bson:"assertion,omitempty" json:"assertion,omitempty"`
}

func (t TestPlan) ResourceType() string {
	return "StructureDefinition"
}
