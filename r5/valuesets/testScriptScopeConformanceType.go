// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/testscript-scope-conformance-codes
*/type TestScriptScopeConformanceType string

const (
	// All tests are expected to pass. Warning statuses are permitted. This is the default value.
	TestScriptScopeConformanceTypeRequired TestScriptScopeConformanceType = "required"
	// All tests are expected to pass but non-pass statuses may be allowed.
	TestScriptScopeConformanceTypeOptional TestScriptScopeConformanceType = "optional"
	// All tests are expected to pass. Warnings are treated as a failure.
	TestScriptScopeConformanceTypeStrict TestScriptScopeConformanceType = "strict"
)
