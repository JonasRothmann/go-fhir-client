// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/therapy-relationship-type
*/type TherapyRelationshipType string

const (
	// Only contraindicated if the other therapy is given
	TherapyRelationshipTypeOnlyContraindicatedIfTheOtherTherapyIsGiven TherapyRelationshipType = "contraindicated-only-with"
	// Contraindicated unless the other therapy is given
	TherapyRelationshipTypeContraindicatedUnlessTheOtherTherapyIsGiven TherapyRelationshipType = "contraindicated-except-with"
	// Indicated only when the other therapy is given (co-occurrent)
	TherapyRelationshipTypeIndicatedOnlyWhenTheOtherTherapyIsGivenCoOccurrent TherapyRelationshipType = "indicated-only-with"
	// Indicated except when the other therapy is given
	TherapyRelationshipTypeIndicatedExceptWhenTheOtherTherapyIsGiven TherapyRelationshipType = "indicated-except-with"
	// Indicated only if the other therapy is planned to be given afterwards (prep)
	TherapyRelationshipTypeIndicatedOnlyIfTheOtherTherapyIsPlannedToBeGivenAfterwardsPrep TherapyRelationshipType = "indicated-only-after"
	// Indicated only if the other therapy was given before (follow-up)
	TherapyRelationshipTypeIndicatedOnlyIfTheOtherTherapyWasGivenBeforeFollowUp TherapyRelationshipType = "indicated-only-before"
	// Indicated to replace the other therapy
	TherapyRelationshipTypeIndicatedToReplaceTheOtherTherapy TherapyRelationshipType = "replace-other-therapy"
	// Indicated to replace the other contraindicated therapy.
	TherapyRelationshipTypeIndicatedToReplaceTheOtherContraindicatedTherapy TherapyRelationshipType = "replace-other-therapy-contraindicated"
	// Indicated to replace the other therapy not well tolerated by patient
	TherapyRelationshipTypeIndicatedToReplaceTheOtherTherapyNotWellToleratedByPatient TherapyRelationshipType = "replace-other-therapy-not-tolerated"
	// Indicated to replace the other therapy not effective on patient
	TherapyRelationshipTypeIndicatedToReplaceTheOtherTherapyNotEffectiveOnPatient TherapyRelationshipType = "replace-other-therapy-not-effective"
)
