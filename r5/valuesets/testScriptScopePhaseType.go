// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/testscript-scope-phase-codes
*/type TestScriptScopePhaseType string

const (
	// The development or implementation phase.
	TestScriptScopePhaseTypeUnit TestScriptScopePhaseType = "unit"
	// The internal system to system phase.
	TestScriptScopePhaseTypeIntegration TestScriptScopePhaseType = "integration"
	// The live system to system phase (Note, this may involve pii/phi data).
	TestScriptScopePhaseTypeProduction TestScriptScopePhaseType = "production"
)
