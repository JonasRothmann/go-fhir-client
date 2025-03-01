// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/guide-page-generation
*/type GuidePageGeneration string

const (
	// Page is proper xhtml with no templating.  Will be brought across unchanged for standard post-processing.
	GuidePageGenerationHtml GuidePageGeneration = "html"
	// Page is markdown with templating.  Will use the template to create a file that imports the markdown file prior to post-processing.
	GuidePageGenerationMarkdown GuidePageGeneration = "markdown"
	// Page is xml with templating.  Will use the template to create a file that imports the source file and run the nominated XSLT transform (see parameters) if present prior to post-processing.
	GuidePageGenerationXml GuidePageGeneration = "xml"
	// Page will be generated by the publication process - no source to bring across.
	GuidePageGenerationGenerated GuidePageGeneration = "generated"
)
