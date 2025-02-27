// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/citation-artifact-classifier
*/type CitationArtifactClassifier string

const (
	// The article cited is an audio file.
	CitationArtifactClassifierAudio CitationArtifactClassifier = "audio"
	// Non-periodical written or printed works consisting of sheets of pages fastened or bound together within covers.
	CitationArtifactClassifierD001877 CitationArtifactClassifier = "D001877"
	// The artifact is used for decision support for healthcare decisions.
	CitationArtifactClassifierClinicalDecisionSupportArtifact CitationArtifactClassifier = "cds-artifact"
	// Comment
	CitationArtifactClassifierD016420 CitationArtifactClassifier = "D016420"
	// Citation Resource containing value added data that is openly shared
	CitationArtifactClassifierCommonShare CitationArtifactClassifier = "common-share"
	// A structured file of information or a set of logically related data stored and retrieved using computer-based means.
	CitationArtifactClassifierD019991 CitationArtifactClassifier = "D019991"
	// Works consisting of organized collections of data, which have been stored permanently in a formalized manner suitable for communication, interpretation, or processing.
	CitationArtifactClassifierD064886 CitationArtifactClassifier = "D064886"
	// An organized collection of data that is not stored permanently for communication.
	CitationArtifactClassifierDatasetUnpublished CitationArtifactClassifier = "dataset-unpublished"
	// the journal is published in electronic format only
	CitationArtifactClassifierElectronic CitationArtifactClassifier = "Electronic"
	// used for electronic-only journals that publish individual articles first and then later collect them into an "issue" date that is typically called an eCollection.
	CitationArtifactClassifierElectronicECollection CitationArtifactClassifier = "Electronic-eCollection"
	// the journal is published first in electronic format followed by print (this value is currently used for just one journal, Nucleic Acids Research)
	CitationArtifactClassifierElectronicPrint CitationArtifactClassifier = "Electronic-Print"
	// Executable app
	CitationArtifactClassifierExecutableApp CitationArtifactClassifier = "executable-app"
	// The article cited is a FHIR resource.
	CitationArtifactClassifierFhirResource CitationArtifactClassifier = "fhir-resource"
	// The article cited is an image file.
	CitationArtifactClassifierImage CitationArtifactClassifier = "image"
	// A user interface that supports data entry and data display.
	CitationArtifactClassifierInteractiveForm CitationArtifactClassifier = "interactive-form"
	// Journal Article
	CitationArtifactClassifierD016428 CitationArtifactClassifier = "D016428"
	// Letter
	CitationArtifactClassifierD016422 CitationArtifactClassifier = "D016422"
	// The article cited is machine code.
	CitationArtifactClassifierMachineCode CitationArtifactClassifier = "machine-code"
	// Citation Resource containing only data from Medline
	CitationArtifactClassifierMedlineBase CitationArtifactClassifier = "medline-base"
	// A formula or expression used to calculate an outcome representing a predicted result.
	CitationArtifactClassifierPredictionModel CitationArtifactClassifier = "prediction-model"
	// Scientific manuscript made available prior to PEER REVIEW.
	CitationArtifactClassifierD000076942 CitationArtifactClassifier = "D000076942"
	// the journal is published in print format only
	CitationArtifactClassifierPrint CitationArtifactClassifier = "Print"
	// the journal is published in both print and electronic format
	CitationArtifactClassifierPrintElectronic CitationArtifactClassifier = "Print-Electronic"
	// Citation Resource containing value added data specific to a project
	CitationArtifactClassifierProjectSpecific CitationArtifactClassifier = "project-specific"
	// The article cited is the protocol of an activity and not the results or findings.
	CitationArtifactClassifierProtocol CitationArtifactClassifier = "protocol"
	// A non-executable, human-readable representation of software code.
	CitationArtifactClassifierPseudocode CitationArtifactClassifier = "pseudocode"
	// Published Erratum
	CitationArtifactClassifierD016425 CitationArtifactClassifier = "D016425"
	// An explicit set of requirements for an item, material, component, system or service, often used to define a technical standard which is an established norm or requirement for a repeatable technical task.
	CitationArtifactClassifierStandardSpecification CitationArtifactClassifier = "standard-specification"
	// A structured set of codes and display values, which may be subtyped as a code system, value set, taxonomy, or ontology.
	CitationArtifactClassifierTerminology CitationArtifactClassifier = "terminology"
	// Used with articles which include video files or clips, or for articles which are entirely video.
	CitationArtifactClassifierD059040 CitationArtifactClassifier = "D059040"
	// Webpage
	CitationArtifactClassifierWebpage CitationArtifactClassifier = "webpage"
)
