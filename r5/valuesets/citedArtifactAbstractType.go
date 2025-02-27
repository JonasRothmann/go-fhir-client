// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/cited-artifact-abstract-type
*/type CitedArtifactAbstractType string

const (
	// Human-friendly main or official abstract
	CitedArtifactAbstractTypePrimaryHumanUse CitedArtifactAbstractType = "primary-human-use"
	// Machine-friendly main or official abstract
	CitedArtifactAbstractTypePrimaryMachineUse CitedArtifactAbstractType = "primary-machine-use"
	// Truncated abstract
	CitedArtifactAbstractTypeTruncated CitedArtifactAbstractType = "truncated"
	// Brief abstract, for use when abstracts are provided in different sizes or lengths
	CitedArtifactAbstractTypeShortAbstract CitedArtifactAbstractType = "short-abstract"
	// Long version of the abstract, for use when abstracts are provided in different sizes or lengths
	CitedArtifactAbstractTypeLongAbstract CitedArtifactAbstractType = "long-abstract"
	// Additional form of abstract written for the general public
	CitedArtifactAbstractTypePlainLanguage CitedArtifactAbstractType = "plain-language"
	// Abstract produced by a different publisher than the cited artifact
	CitedArtifactAbstractTypeDifferentPublisherForAbstract CitedArtifactAbstractType = "different-publisher"
	// Additional form of abstract in a different language
	CitedArtifactAbstractTypeLanguage CitedArtifactAbstractType = "language"
	// Machine translated form of abstract in a different language, language element codes the language into which it was translated by machine
	CitedArtifactAbstractTypeAutotranslated CitedArtifactAbstractType = "autotranslated"
	// Alternative form of abstract in two or more Medline entries
	CitedArtifactAbstractTypeDifferentTextInAdditionalMedlineEntry CitedArtifactAbstractType = "duplicate-pmid"
	// Alternative form of abstract in an earlier version such as epub ahead of print
	CitedArtifactAbstractTypeDifferentTextInAnEarlierVersion CitedArtifactAbstractType = "earlier-abstract"
)
