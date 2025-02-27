// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/consent-data-meaning
*/type ConsentDataMeaning string

const (
	// The consent applies directly to the instance of the resource.
	ConsentDataMeaningInstance ConsentDataMeaning = "instance"
	// The consent applies directly to the instance of the resource and instances it refers to.
	ConsentDataMeaningRelated ConsentDataMeaning = "related"
	// The consent applies directly to the instance of the resource and instances that refer to it.
	ConsentDataMeaningDependents ConsentDataMeaning = "dependents"
	// The consent applies to instances of resources that are authored by.
	ConsentDataMeaningAuthoredby ConsentDataMeaning = "authoredby"
)
