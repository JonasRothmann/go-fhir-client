// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/cited-artifact-part-type
*/type CitedArtifactPartType string

const (
	// Denotes specific page or pages of an article or artifact.
	CitedArtifactPartTypePages CitedArtifactPartType = "pages"
	// Denotes specific section or sections of an article or artifact.
	CitedArtifactPartTypeSections CitedArtifactPartType = "sections"
	// Denotes specific paragraph or paragraphs of an article or artifact.
	CitedArtifactPartTypeParagraphs CitedArtifactPartType = "paragraphs"
	// Denotes specific line or lines of an article or artifact.
	CitedArtifactPartTypeLines CitedArtifactPartType = "lines"
	// Denotes specific table or tables of an article or artifact.
	CitedArtifactPartTypeTables CitedArtifactPartType = "tables"
	// Denotes specific figure or figures of an article or artifact.
	CitedArtifactPartTypeFigures CitedArtifactPartType = "figures"
	// Used to denote a supplementary file, appendix, or additional part that is not a subpart of the primary article.
	CitedArtifactPartTypeSupplement CitedArtifactPartType = "supplement"
	// Used to denote a subpart within a supplementary file or appendix.
	CitedArtifactPartTypeSupplementOrAppendixSubpart CitedArtifactPartType = "supplement-subpart"
	// Used to distinguish an individual article within an article set where the article set is a base citation.
	CitedArtifactPartTypePartOfAnArticleSet CitedArtifactPartType = "article-set"
)
