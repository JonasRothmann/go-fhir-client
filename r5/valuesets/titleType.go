// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/title-type
*/type TitleType string

const (
	// Main title for common use. The primary title used for representation if multiple titles exist.
	TitleTypePrimary TitleType = "primary"
	// The official or authoritative title.
	TitleTypeOfficial TitleType = "official"
	// Title using scientific terminology.
	TitleTypeScientific TitleType = "scientific"
	// Title using language common to lay public discourse.
	TitleTypePlainLanguageTitle TitleType = "plain-language"
	// Subtitle or secondary title.
	TitleTypeSubtitle TitleType = "subtitle"
	// Brief title (e.g. 'running title' or title used in page headers)
	TitleTypeShortTitle TitleType = "short-title"
	// Abbreviation used as title
	TitleTypeAcronym TitleType = "acronym"
	// Alternative form of title in an earlier version such as epub ahead of print.
	TitleTypeDifferentTextInAnEarlierVersion TitleType = "earlier-title"
	// Additional form of title in a different language.
	TitleTypeLanguage TitleType = "language"
	// Machine translated form of title in a different language, language element codes the language into which it was translated by machine.
	TitleTypeAutotranslated TitleType = "autotranslated"
	// Human-friendly title
	TitleTypeHumanUse TitleType = "human-use"
	// Machine-friendly title
	TitleTypeMachineUse TitleType = "machine-use"
	// An alternative form of the title in two or more entries, e.g. in multiple medline entries
	TitleTypeDifferentTextForTheSameObjectWithADifferentIdentifier TitleType = "duplicate-uid"
)
