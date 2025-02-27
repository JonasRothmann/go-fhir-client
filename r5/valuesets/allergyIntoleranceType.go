// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/allergy-intolerance-type
*/type AllergyIntoleranceType string

const (
	// A propensity for hypersensitive reaction(s) to a substance.  These reactions are most typically type I hypersensitivity, plus other "allergy-like" reactions, including pseudoallergy.
	AllergyIntoleranceTypeAllergy AllergyIntoleranceType = "allergy"
	// A propensity for adverse reactions to a substance that is judged to be not allergic or "allergy-like".  These reactions are typically (but not necessarily) non-immune.  They are to some degree idiosyncratic and/or patient-specific (i.e. are not a reaction that is expected to occur with most or all patients given similar circumstances).
	AllergyIntoleranceTypeIntolerance AllergyIntoleranceType = "intolerance"
)
