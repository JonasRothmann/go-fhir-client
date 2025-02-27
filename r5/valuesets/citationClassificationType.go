// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/citation-classification-type
*/type CitationClassificationType string

const (
	// Citation repository where this citation was created or copied from
	CitationClassificationTypeCitationSource CitationClassificationType = "citation-source"
	// The party responsible for creating and validating the MEDLINE citation
	CitationClassificationTypeMedlineCitationOwner CitationClassificationType = "medline-owner"
	// Used for Citation sharing on the Fast Evidence Interoperability Resources (FEvIR) Platform
	CitationClassificationTypeFEvIrPlatformUse CitationClassificationType = "fevir-platform-use"
)
