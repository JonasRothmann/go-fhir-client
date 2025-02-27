// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/allergyintolerance-verification
*/type AllergyIntoleranceVerificationStatus string

const (
	// The propensity for a reaction to the identified substance has not been objectively verified.
	AllergyIntoleranceVerificationStatusUnconfirmed AllergyIntoleranceVerificationStatus = "unconfirmed"
	// The propensity for a reaction to the identified substance has been objectively verified (which may include clinical evidence by testing, rechallenge, or observation).
	AllergyIntoleranceVerificationStatusConfirmed AllergyIntoleranceVerificationStatus = "confirmed"
	// A propensity for a reaction to the identified substance has been disputed or disproven with a sufficient level of clinical certainty to justify invalidating the assertion. This might or might not include testing or rechallenge.
	AllergyIntoleranceVerificationStatusRefuted AllergyIntoleranceVerificationStatus = "refuted"
	// The statement was entered in error and is not valid.
	AllergyIntoleranceVerificationStatusEnteredInError AllergyIntoleranceVerificationStatus = "entered-in-error"
)
