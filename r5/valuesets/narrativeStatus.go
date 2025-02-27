// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/narrative-status
*/type NarrativeStatus string

const (
	// The contents of the narrative are entirely generated from the core elements in the content.
	NarrativeStatusGenerated NarrativeStatus = "generated"
	// The contents of the narrative are entirely generated from the core elements in the content and some of the content is generated from extensions. The narrative SHALL reflect the impact of all modifier extensions.
	NarrativeStatusExtensions NarrativeStatus = "extensions"
	// The contents of the narrative may contain additional information not found in the structured data. Note that there is no computable way to determine what the extra information is, other than by human inspection.
	NarrativeStatusAdditional NarrativeStatus = "additional"
	// The contents of the narrative are some equivalent of "No human-readable text provided in this case".
	NarrativeStatusEmpty NarrativeStatus = "empty"
)
