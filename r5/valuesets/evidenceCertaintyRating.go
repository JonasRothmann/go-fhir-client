// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/certainty-rating
*/type EvidenceCertaintyRating string

const (
	// High quality evidence.
	EvidenceCertaintyRatingHigh EvidenceCertaintyRating = "high"
	// Moderate quality evidence.
	EvidenceCertaintyRatingModerate EvidenceCertaintyRating = "moderate"
	// Low quality evidence.
	EvidenceCertaintyRatingLow EvidenceCertaintyRating = "low"
	// Very low quality evidence.
	EvidenceCertaintyRatingVeryLowQuality EvidenceCertaintyRating = "very-low"
	// no serious concern.
	EvidenceCertaintyRatingNoSeriousConcern EvidenceCertaintyRating = "no-concern"
	// serious concern.
	EvidenceCertaintyRatingSeriousConcern EvidenceCertaintyRating = "serious-concern"
	// very serious concern.
	EvidenceCertaintyRatingVerySeriousConcern EvidenceCertaintyRating = "very-serious-concern"
	// extremely serious concern.
	EvidenceCertaintyRatingExtremelySeriousConcern EvidenceCertaintyRating = "extremely-serious-concern"
	// possible reason for increasing quality rating was checked and found to be present.
	EvidenceCertaintyRatingPresent EvidenceCertaintyRating = "present"
	// possible reason for increasing quality rating was checked and found to be absent.
	EvidenceCertaintyRatingAbsent EvidenceCertaintyRating = "absent"
	// no change to quality rating.
	EvidenceCertaintyRatingNoChangeToRating EvidenceCertaintyRating = "no-change"
	// reduce quality rating by 1.
	EvidenceCertaintyRatingDowncode1 EvidenceCertaintyRating = "downcode1"
	// reduce quality rating by 2.
	EvidenceCertaintyRatingDowncode2 EvidenceCertaintyRating = "downcode2"
	// reduce quality rating by 3.
	EvidenceCertaintyRatingDowncode3 EvidenceCertaintyRating = "downcode3"
	// increase quality rating by 1.
	EvidenceCertaintyRatingUpcode1 EvidenceCertaintyRating = "upcode1"
	// increase quality rating by 2.
	EvidenceCertaintyRatingUpcode2 EvidenceCertaintyRating = "upcode2"
)
