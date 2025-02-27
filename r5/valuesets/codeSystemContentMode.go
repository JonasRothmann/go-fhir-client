// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/codesystem-content-mode
*/type CodeSystemContentMode string

const (
	// None of the concepts defined by the code system are included in the code system resource.
	CodeSystemContentModeNotPresent CodeSystemContentMode = "not-present"
	// A subset of the valid externally defined concepts are included in the code system resource. There is no specific purpose or documented intent other than for illustrative purposes.
	CodeSystemContentModeExample CodeSystemContentMode = "example"
	// A subset of the code system concepts are included in the code system resource. This is a curated subset released for a specific purpose under the governance of the code system steward, and that the intent, bounds and consequences of the fragmentation are clearly defined in the fragment or the code system documentation. Fragments are also known as partitions.
	CodeSystemContentModeFragment CodeSystemContentMode = "fragment"
	// All the concepts defined by the code system are included in the code system resource.
	CodeSystemContentModeComplete CodeSystemContentMode = "complete"
	// The resource doesn't define any new concepts; it just provides additional designations and properties to another code system.
	CodeSystemContentModeSupplement CodeSystemContentMode = "supplement"
)
