// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/artifact-url-classifier
*/type ArtifactUrlClassifier string

const (
	// The URL will reach a brief summary for the article.
	ArtifactUrlClassifierAbstract ArtifactUrlClassifier = "abstract"
	// The URL will reach the full-text of the article.
	ArtifactUrlClassifierFullText ArtifactUrlClassifier = "full-text"
	// The URL will reach a supplement, appendix, or additional supporting information for the article.
	ArtifactUrlClassifierSupplement ArtifactUrlClassifier = "supplement"
	// The URL will reach a webpage related to the article, where the content is not easily classified as abstract, full-text or supplement.
	ArtifactUrlClassifierWebpage ArtifactUrlClassifier = "webpage"
	// The URL will reach a file directory.
	ArtifactUrlClassifierFileDirectory ArtifactUrlClassifier = "file-directory"
	// File archive and web hosting facility for source code of software, documentation, web pages, and other works.
	ArtifactUrlClassifierCodeRepository ArtifactUrlClassifier = "code-repository"
	// The URL content has restricted access (e.g. subcription required).
	ArtifactUrlClassifierRestricted ArtifactUrlClassifier = "restricted"
	// Compressed archive file (e.g. a zip file) that contains multiple files
	ArtifactUrlClassifierCompressedFile ArtifactUrlClassifier = "compressed-file"
	// The URL is derived from the Digital Object Identifier (DOI).
	ArtifactUrlClassifierDoiBased ArtifactUrlClassifier = "doi-based"
	// The URL will reach content in PDF form.
	ArtifactUrlClassifierPdf ArtifactUrlClassifier = "pdf"
	// The URL will reach content in JSON format.
	ArtifactUrlClassifierJson ArtifactUrlClassifier = "json"
	// The URL will reach content in XML format.
	ArtifactUrlClassifierXml ArtifactUrlClassifier = "xml"
	// The URL will reach content that is a specific version of the article.
	ArtifactUrlClassifierVersionSpecific ArtifactUrlClassifier = "version-specific"
	// The URL will reach content that is machine-interpretable.
	ArtifactUrlClassifierComputableResource ArtifactUrlClassifier = "computable-resource"
	// Used when URL classifier is not specified but expected in a system.
	ArtifactUrlClassifierNotSpecified ArtifactUrlClassifier = "not-specified"
)
