// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/code-search-support
*/type CodeSearchSupport string

const (
	// The search for code on ValueSet returns ValueSet resources where the code is included in the extensional definition of the ValueSet.
	CodeSearchSupportInCompose CodeSearchSupport = "in-compose"
	// The search for code on ValueSet returns ValueSet resources where the code is contained in the  ValueSet expansion.
	CodeSearchSupportInExpansion CodeSearchSupport = "in-expansion"
	// The search for code on ValueSet returns ValueSet resources where the code is included in the extensional definition or contained in the ValueSet expansion.
	CodeSearchSupportInComposeOrExpansion CodeSearchSupport = "in-compose-or-expansion"
)
