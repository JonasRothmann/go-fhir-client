// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/v3-ActCode
- http://terminology.hl7.org/CodeSystem/v3-ActCode
- http://terminology.hl7.org/CodeSystem/v3-ActCode
- http://terminology.hl7.org/CodeSystem/v3-ActCode
*/type DetectedIssueCategory string

const (
	// Administration of the proposed therapy may be inappropriate or contraindicated as proposed
	DetectedIssueCategoryAdministrationDetectedIssueCode  DetectedIssueCategory = "_AdministrationDetectedIssueCode"
	DetectedIssueCategoryAppropriatenessDetectedIssueCode DetectedIssueCategory = "_AppropriatenessDetectedIssueCode"
	// There may be an issue with the patient complying with the intentions of the proposed therapy
	DetectedIssueCategoryComplianceAlert DetectedIssueCategory = "COMPLY"
	// Proposed therapy may be contraindicated or ineffective based on an existing or recent drug therapy.
	DetectedIssueCategoryDrugActionDetectedIssue DetectedIssueCategory = "DACT"
	// Proposed dosage instructions for therapy differ from standard practice.
	DetectedIssueCategoryDosageProblem DetectedIssueCategory = "DOSE"
	// The proposed therapy appears to duplicate an existing therapy
	DetectedIssueCategoryDuplicateTherapyAlert DetectedIssueCategory = "DUPTHPY"
	// Proposed therapy may be inappropriate or ineffective based on the proposed start or end time.
	DetectedIssueCategoryTimingDetectedIssue DetectedIssueCategory = "TIME"
	// Proposed therapy may be contraindicated or ineffective based on an existing or recent drug therapy
	DetectedIssueCategoryDrugActionDetectedIssueCode DetectedIssueCategory = "_DrugActionDetectedIssueCode"
	// Proposed therapy may be inappropriate or ineffective based on the proposed start or end time.
	DetectedIssueCategoryTimingDetectedIssueCode      DetectedIssueCategory = "_TimingDetectedIssueCode"
	DetectedIssueCategoryInteractionDetectedIssueCode DetectedIssueCategory = "_InteractionDetectedIssueCode"
	// Proposed therapy may be inappropriate or contraindicated due to conditions or characteristics of the patient
	DetectedIssueCategoryObservationAlert      DetectedIssueCategory = "OBSA"
	DetectedIssueCategoryPreviouslyIneffective DetectedIssueCategory = "PREVINEF"
	// Proposed therapy may interact with certain foods
	DetectedIssueCategoryFoodInteractionAlert DetectedIssueCategory = "FOOD"
	// Proposed therapy may interact with an existing or recent therapeutic product
	DetectedIssueCategoryTherapeuticProductAlert DetectedIssueCategory = "TPROD"
	// Proposed therapy may be inappropriate or ineffective because the end of administration is too close to another planned therapy
	DetectedIssueCategoryEndTooLateAlert DetectedIssueCategory = "ENDLATE"
	// Proposed therapy may be inappropriate or ineffective because the start of administration is too late after the onset of the condition
	DetectedIssueCategoryStartTooLateAlert DetectedIssueCategory = "STRTLATE"
	// Proposed therapy is outside of the standard practice for an adult patient.
	DetectedIssueCategoryAdultAlert DetectedIssueCategory = "ADALRT"
	// Proposed dosage exceeds standard practice for the patient's age
	DetectedIssueCategoryHighDoseForAgeAlert DetectedIssueCategory = "DOSEHINDA"
	// Proposed dosage is below suggested therapeutic levels for the patient's age
	DetectedIssueCategoryLowDoseForAgeAlert DetectedIssueCategory = "DOSELINDA"
	// Proposed therapy is outside of standard practice for a geriatric patient.
	DetectedIssueCategoryGeriatricAlert DetectedIssueCategory = "GEALRT"
	// Proposed therapy is outside of the standard practice for a pediatric patient.
	DetectedIssueCategoryPediatricAlert DetectedIssueCategory = "PEALRT"
	// The proposed therapy is frequently misused or abused and therefore should be used with caution and/or monitoring.
	DetectedIssueCategoryCommonlyAbusedMisusedAlert DetectedIssueCategory = "ABUSE"
	// The request is suspected to have a fraudulent basis.
	DetectedIssueCategoryPotentialFraud DetectedIssueCategory = "FRAUD"
	// A similar or identical therapy was recently ordered by a different practitioner.
	DetectedIssueCategoryPolyOrdererAlert DetectedIssueCategory = "PLYDOC"
	// This patient was recently supplied a similar or identical therapy from a different pharmacy or supplier.
	DetectedIssueCategoryPolySupplierAlert DetectedIssueCategory = "PLYPHRM"
	DetectedIssueCategoryHght              DetectedIssueCategory = "HGHT"
	// Proposed therapy may be inappropriate or contraindicated when breast-feeding
	DetectedIssueCategoryLactationAlert DetectedIssueCategory = "LACT"
	// Proposed therapy may be inappropriate or contraindicated during pregnancy
	DetectedIssueCategoryPregnancyAlert DetectedIssueCategory = "PREG"
	DetectedIssueCategoryWght           DetectedIssueCategory = "WGHT"
	// Proposed dosage is inappropriate due to patient's medical condition.
	DetectedIssueCategoryDosageConditionAlert DetectedIssueCategory = "DOSECOND"
	// Proposed length of therapy differs from standard practice.
	DetectedIssueCategoryDoseDurationAlert DetectedIssueCategory = "DOSEDUR"
	// Proposed dosage exceeds standard practice
	DetectedIssueCategoryHighDoseAlert DetectedIssueCategory = "DOSEH"
	// Proposed dosage interval/timing differs from standard practice
	DetectedIssueCategoryDoseIntervalAlert DetectedIssueCategory = "DOSEIVL"
	// Proposed dosage is below suggested therapeutic levels
	DetectedIssueCategoryLowDoseAlert DetectedIssueCategory = "DOSEL"
	// The maximum quantity of this drug allowed to be administered within a particular time-range (month, year, lifetime) has been reached or exceeded.
	DetectedIssueCategoryMaximumDosageReached DetectedIssueCategory = "MDOSE"
	// Proposed length of therapy is longer than standard practice
	DetectedIssueCategoryDoseDurationHighAlert DetectedIssueCategory = "DOSEDURH"
	// Proposed length of therapy is shorter than that necessary for therapeutic effect
	DetectedIssueCategoryDoseDurationLowAlert DetectedIssueCategory = "DOSEDURL"
	// Proposed length of therapy is longer than standard practice for the identified indication or diagnosis
	DetectedIssueCategoryDoseDurationHighForIndicationAlert DetectedIssueCategory = "DOSEDURHIND"
	// Proposed length of therapy is shorter than standard practice for the identified indication or diagnosis
	DetectedIssueCategoryDoseDurationLowForIndicationAlert DetectedIssueCategory = "DOSEDURLIND"
	DetectedIssueCategoryHighDoseForIndicationAlert        DetectedIssueCategory = "DOSEHIND"
	// Proposed dosage exceeds standard practice for the patient's height or body surface area
	DetectedIssueCategoryHighDoseForHeightSurfaceAreaAlert DetectedIssueCategory = "DOSEHINDSA"
	// Proposed dosage exceeds standard practice for the patient's weight
	DetectedIssueCategoryHighDoseForWeightAlert DetectedIssueCategory = "DOSEHINDW"
	// Proposed dosage interval/timing differs from standard practice for the identified indication or diagnosis
	DetectedIssueCategoryDoseIntervalForIndicationAlert DetectedIssueCategory = "DOSEIVLIND"
	DetectedIssueCategoryLowDoseForIndicationAlert      DetectedIssueCategory = "DOSELIND"
	// Proposed dosage is below suggested therapeutic levels for the patient's height or body surface area
	DetectedIssueCategoryLowDoseForHeightSurfaceAreaAlert DetectedIssueCategory = "DOSELINDSA"
	// Proposed dosage is below suggested therapeutic levels for the patient's weight
	DetectedIssueCategoryLowDoseForWeightAlert DetectedIssueCategory = "DOSELINDW"
	// The proposed therapy appears to have the same intended therapeutic benefit as an existing therapy, though the specific mechanisms of action vary.
	DetectedIssueCategoryDuplicateTherapeuticAlassAlert DetectedIssueCategory = "DUPTHPCLS"
	// The proposed therapy appears to have the same intended therapeutic benefit as an existing therapy and uses the same mechanisms of action as the existing therapy.
	DetectedIssueCategoryDuplicateGenericAlert DetectedIssueCategory = "DUPTHPGEN"
	// Proposed therapy may be inappropriate or contraindicated due to patient age
	DetectedIssueCategoryAgeAlert DetectedIssueCategory = "AGE"
	// Proposed therapy may be inappropriate or contraindicated due to an existing/recent patient condition or diagnosis
	DetectedIssueCategoryConditionAlert DetectedIssueCategory = "COND"
	// Proposed therapy may be inappropriate or contraindicated because of a common but non-patient specific reaction to the product.
	DetectedIssueCategoryCommonReactionAlert DetectedIssueCategory = "CREACT"
	// Proposed therapy may be inappropriate or contraindicated due to patient genetic indicators.
	DetectedIssueCategoryGeneticAlert DetectedIssueCategory = "GEN"
	// Proposed therapy may be inappropriate or contraindicated due to patient gender.
	DetectedIssueCategoryGenderAlert DetectedIssueCategory = "GEND"
	// Proposed therapy may be inappropriate or contraindicated due to recent lab test results
	DetectedIssueCategoryLabAlert DetectedIssueCategory = "LAB"
	// Proposed therapy may be inappropriate or contraindicated based on the potential for a patient reaction to the proposed product
	DetectedIssueCategoryReactionAlert DetectedIssueCategory = "REACT"
	// Proposed therapy may be inappropriate or contraindicated because of a potential patient reaction to a cross-sensitivity related product.
	DetectedIssueCategoryRelatedReactionAlert DetectedIssueCategory = "RREACT"
	// Proposed therapy may be inappropriate or contraindicated because of a recorded patient allergy to the proposed product. (Allergies are immune based reactions.)
	DetectedIssueCategoryAllergyAlert DetectedIssueCategory = "ALGY"
	// Proposed therapy may be inappropriate or contraindicated because of a recorded patient intolerance to the proposed product. (Intolerances are non-immune based sensitivities.)
	DetectedIssueCategoryIntoleranceAlert DetectedIssueCategory = "INT"
	// Proposed therapy may be inappropriate or contraindicated because of a recorded patient allergy to a cross-sensitivity related product. (Allergies are immune based reactions.)
	DetectedIssueCategoryRelatedAllergyAlert DetectedIssueCategory = "RALG"
	// Proposed therapy may be inappropriate or contraindicated because of a recorded prior adverse reaction to a cross-sensitivity related product.
	DetectedIssueCategoryRelatedPriorReactionAlert DetectedIssueCategory = "RAR"
	// Proposed therapy may be inappropriate or contraindicated because of a recorded patient intolerance to a cross-sensitivity related product. (Intolerances are non-immune based sensitivities.)
	DetectedIssueCategoryRelatedIntoleranceAlert DetectedIssueCategory = "RINT"
	DetectedIssueCategoryEndTooLateAlert2        DetectedIssueCategory = "ALRTENDLATE"
	DetectedIssueCategoryStartTooLateAlert2      DetectedIssueCategory = "ALRTSTRTLATE"
	// Proposed therapy may interact with an existing or recent drug therapy
	DetectedIssueCategoryDrugInteractionAlert DetectedIssueCategory = "DRG"
	// Proposed therapy may interact with existing or recent natural health product therapy
	DetectedIssueCategoryNaturalHealthProductAlert DetectedIssueCategory = "NHP"
	// Proposed therapy may interact with a non-prescription drug (e.g. alcohol, tobacco, Aspirin)
	DetectedIssueCategoryNonPrescriptionInteractionAlert DetectedIssueCategory = "NONRX"
	// Administration of the proposed therapy may be inappropriate or contraindicated as proposed
	DetectedIssueCategoryAdministrationDetectedIssueCode2 DetectedIssueCategory = "_AdministrationDetectedIssueCode"
	// Supplying the product at this time may be inappropriate or indicate compliance issues with the associated therapy
	DetectedIssueCategorySupplyDetectedIssueCode DetectedIssueCategory = "_SupplyDetectedIssueCode"
	// While the record was accepted in the repository, there is a more recent version of a record of this type.
	DetectedIssueCategoryRecordRecordedAsHistorical        DetectedIssueCategory = "HISTORIC"
	DetectedIssueCategoryViolatesStatedPreferences         DetectedIssueCategory = "PATPREF"
	DetectedIssueCategoryAppropriatenessDetectedIssueCode2 DetectedIssueCategory = "_AppropriatenessDetectedIssueCode"
	// There may be an issue with the patient complying with the intentions of the proposed therapy
	DetectedIssueCategoryComplianceAlert2 DetectedIssueCategory = "COMPLY"
	// Proposed therapy may be contraindicated or ineffective based on an existing or recent drug therapy.
	DetectedIssueCategoryDrugActionDetectedIssue2 DetectedIssueCategory = "DACT"
	// Proposed dosage instructions for therapy differ from standard practice.
	DetectedIssueCategoryDosageProblem2 DetectedIssueCategory = "DOSE"
	// The proposed therapy appears to duplicate an existing therapy
	DetectedIssueCategoryDuplicateTherapyAlert2 DetectedIssueCategory = "DUPTHPY"
	// Proposed therapy may be inappropriate or ineffective based on the proposed start or end time.
	DetectedIssueCategoryTimingDetectedIssue2 DetectedIssueCategory = "TIME"
	// Proposed therapy may be contraindicated or ineffective based on an existing or recent drug therapy
	DetectedIssueCategoryDrugActionDetectedIssueCode2 DetectedIssueCategory = "_DrugActionDetectedIssueCode"
	// Proposed therapy may be inappropriate or ineffective based on the proposed start or end time.
	DetectedIssueCategoryTimingDetectedIssueCode2      DetectedIssueCategory = "_TimingDetectedIssueCode"
	DetectedIssueCategoryInteractionDetectedIssueCode2 DetectedIssueCategory = "_InteractionDetectedIssueCode"
	// Proposed therapy may be inappropriate or contraindicated due to conditions or characteristics of the patient
	DetectedIssueCategoryObservationAlert2      DetectedIssueCategory = "OBSA"
	DetectedIssueCategoryPreviouslyIneffective2 DetectedIssueCategory = "PREVINEF"
	// Proposed therapy may interact with certain foods
	DetectedIssueCategoryFoodInteractionAlert2 DetectedIssueCategory = "FOOD"
	// Proposed therapy may interact with an existing or recent therapeutic product
	DetectedIssueCategoryTherapeuticProductAlert2 DetectedIssueCategory = "TPROD"
	DetectedIssueCategoryAlreadyPerformed         DetectedIssueCategory = "ALLDONE"
	DetectedIssueCategoryFulfillmentAlert         DetectedIssueCategory = "FULFIL"
	DetectedIssueCategoryHeldSuspendedAlert       DetectedIssueCategory = "HELD"
	// The patient is receiving a subsequent fill significantly later than would be expected based on the amount previously supplied and the therapy dosage instructions
	DetectedIssueCategoryRefillTooLateAlert DetectedIssueCategory = "TOOLATE"
	// The patient is receiving a subsequent fill significantly earlier than would be expected based on the amount previously supplied and the therapy dosage instructions
	DetectedIssueCategoryRefillTooSoonAlert DetectedIssueCategory = "TOOSOON"
	// Proposed therapy may be inappropriate or ineffective because the end of administration is too close to another planned therapy
	DetectedIssueCategoryEndTooLateAlert3 DetectedIssueCategory = "ENDLATE"
	// Proposed therapy may be inappropriate or ineffective because the start of administration is too late after the onset of the condition
	DetectedIssueCategoryStartTooLateAlert3 DetectedIssueCategory = "STRTLATE"
	// Proposed therapy is outside of the standard practice for an adult patient.
	DetectedIssueCategoryAdultAlert2 DetectedIssueCategory = "ADALRT"
	// Proposed dosage exceeds standard practice for the patient's age
	DetectedIssueCategoryHighDoseForAgeAlert2 DetectedIssueCategory = "DOSEHINDA"
	// Proposed dosage is below suggested therapeutic levels for the patient's age
	DetectedIssueCategoryLowDoseForAgeAlert2 DetectedIssueCategory = "DOSELINDA"
	// Proposed therapy is outside of standard practice for a geriatric patient.
	DetectedIssueCategoryGeriatricAlert2 DetectedIssueCategory = "GEALRT"
	// Proposed therapy is outside of the standard practice for a pediatric patient.
	DetectedIssueCategoryPediatricAlert2 DetectedIssueCategory = "PEALRT"
	// The proposed therapy is frequently misused or abused and therefore should be used with caution and/or monitoring.
	DetectedIssueCategoryCommonlyAbusedMisusedAlert2 DetectedIssueCategory = "ABUSE"
	// The request is suspected to have a fraudulent basis.
	DetectedIssueCategoryPotentialFraud2 DetectedIssueCategory = "FRAUD"
	// A similar or identical therapy was recently ordered by a different practitioner.
	DetectedIssueCategoryPolyOrdererAlert2 DetectedIssueCategory = "PLYDOC"
	// This patient was recently supplied a similar or identical therapy from a different pharmacy or supplier.
	DetectedIssueCategoryPolySupplierAlert2 DetectedIssueCategory = "PLYPHRM"
	DetectedIssueCategoryHght2              DetectedIssueCategory = "HGHT"
	// Proposed therapy may be inappropriate or contraindicated when breast-feeding
	DetectedIssueCategoryLactationAlert2 DetectedIssueCategory = "LACT"
	// Proposed therapy may be inappropriate or contraindicated during pregnancy
	DetectedIssueCategoryPregnancyAlert2 DetectedIssueCategory = "PREG"
	DetectedIssueCategoryWght2           DetectedIssueCategory = "WGHT"
	// Proposed dosage is inappropriate due to patient's medical condition.
	DetectedIssueCategoryDosageConditionAlert2 DetectedIssueCategory = "DOSECOND"
	// Proposed length of therapy differs from standard practice.
	DetectedIssueCategoryDoseDurationAlert2 DetectedIssueCategory = "DOSEDUR"
	// Proposed dosage exceeds standard practice
	DetectedIssueCategoryHighDoseAlert2 DetectedIssueCategory = "DOSEH"
	// Proposed dosage interval/timing differs from standard practice
	DetectedIssueCategoryDoseIntervalAlert2 DetectedIssueCategory = "DOSEIVL"
	// Proposed dosage is below suggested therapeutic levels
	DetectedIssueCategoryLowDoseAlert2 DetectedIssueCategory = "DOSEL"
	// The maximum quantity of this drug allowed to be administered within a particular time-range (month, year, lifetime) has been reached or exceeded.
	DetectedIssueCategoryMaximumDosageReached2 DetectedIssueCategory = "MDOSE"
	// Proposed length of therapy is longer than standard practice
	DetectedIssueCategoryDoseDurationHighAlert2 DetectedIssueCategory = "DOSEDURH"
	// Proposed length of therapy is shorter than that necessary for therapeutic effect
	DetectedIssueCategoryDoseDurationLowAlert2 DetectedIssueCategory = "DOSEDURL"
	// Proposed length of therapy is longer than standard practice for the identified indication or diagnosis
	DetectedIssueCategoryDoseDurationHighForIndicationAlert2 DetectedIssueCategory = "DOSEDURHIND"
	// Proposed length of therapy is shorter than standard practice for the identified indication or diagnosis
	DetectedIssueCategoryDoseDurationLowForIndicationAlert2 DetectedIssueCategory = "DOSEDURLIND"
	DetectedIssueCategoryHighDoseForIndicationAlert2        DetectedIssueCategory = "DOSEHIND"
	// Proposed dosage exceeds standard practice for the patient's height or body surface area
	DetectedIssueCategoryHighDoseForHeightSurfaceAreaAlert2 DetectedIssueCategory = "DOSEHINDSA"
	// Proposed dosage exceeds standard practice for the patient's weight
	DetectedIssueCategoryHighDoseForWeightAlert2 DetectedIssueCategory = "DOSEHINDW"
	// Proposed dosage interval/timing differs from standard practice for the identified indication or diagnosis
	DetectedIssueCategoryDoseIntervalForIndicationAlert2 DetectedIssueCategory = "DOSEIVLIND"
	DetectedIssueCategoryLowDoseForIndicationAlert2      DetectedIssueCategory = "DOSELIND"
	// Proposed dosage is below suggested therapeutic levels for the patient's height or body surface area
	DetectedIssueCategoryLowDoseForHeightSurfaceAreaAlert2 DetectedIssueCategory = "DOSELINDSA"
	// Proposed dosage is below suggested therapeutic levels for the patient's weight
	DetectedIssueCategoryLowDoseForWeightAlert2 DetectedIssueCategory = "DOSELINDW"
	// The proposed therapy appears to have the same intended therapeutic benefit as an existing therapy, though the specific mechanisms of action vary.
	DetectedIssueCategoryDuplicateTherapeuticAlassAlert2 DetectedIssueCategory = "DUPTHPCLS"
	// The proposed therapy appears to have the same intended therapeutic benefit as an existing therapy and uses the same mechanisms of action as the existing therapy.
	DetectedIssueCategoryDuplicateGenericAlert2    DetectedIssueCategory = "DUPTHPGEN"
	DetectedIssueCategoryNoLongerActionable        DetectedIssueCategory = "NOTACTN"
	DetectedIssueCategoryNotEquivalentAlert        DetectedIssueCategory = "NOTEQUIV"
	DetectedIssueCategoryEventTimingIncorrectAlert DetectedIssueCategory = "TIMING"
	// Identifies types of detected issues regarding the administration or supply of an item to a patient.
	DetectedIssueCategoryActSuppliedItemDetectedIssueCode  DetectedIssueCategory = "_ActSuppliedItemDetectedIssueCode"
	DetectedIssueCategoryNotGenericallyEquivalentAlert     DetectedIssueCategory = "NOTEQUIVGEN"
	DetectedIssueCategoryNotTherapeuticallyEquivalentAlert DetectedIssueCategory = "NOTEQUIVTHER"
	// Proposed therapy may be inappropriate or contraindicated due to patient age
	DetectedIssueCategoryAgeAlert2 DetectedIssueCategory = "AGE"
	// Proposed therapy may be inappropriate or contraindicated due to an existing/recent patient condition or diagnosis
	DetectedIssueCategoryConditionAlert2 DetectedIssueCategory = "COND"
	// Proposed therapy may be inappropriate or contraindicated because of a common but non-patient specific reaction to the product.
	DetectedIssueCategoryCommonReactionAlert2 DetectedIssueCategory = "CREACT"
	// Proposed therapy may be inappropriate or contraindicated due to patient genetic indicators.
	DetectedIssueCategoryGeneticAlert2 DetectedIssueCategory = "GEN"
	// Proposed therapy may be inappropriate or contraindicated due to patient gender.
	DetectedIssueCategoryGenderAlert2 DetectedIssueCategory = "GEND"
	// Proposed therapy may be inappropriate or contraindicated due to recent lab test results
	DetectedIssueCategoryLabAlert2 DetectedIssueCategory = "LAB"
	// Proposed therapy may be inappropriate or contraindicated based on the potential for a patient reaction to the proposed product
	DetectedIssueCategoryReactionAlert2 DetectedIssueCategory = "REACT"
	// Proposed therapy may be inappropriate or contraindicated because of a potential patient reaction to a cross-sensitivity related product.
	DetectedIssueCategoryRelatedReactionAlert2                       DetectedIssueCategory = "RREACT"
	DetectedIssueCategoryViolatesStatedPreferencesAlternateAvailable DetectedIssueCategory = "PATPREFALT"
	// Proposed therapy may be inappropriate or contraindicated because of a recorded patient allergy to the proposed product. (Allergies are immune based reactions.)
	DetectedIssueCategoryAllergyAlert2 DetectedIssueCategory = "ALGY"
	// Proposed therapy may be inappropriate or contraindicated because of a recorded patient intolerance to the proposed product. (Intolerances are non-immune based sensitivities.)
	DetectedIssueCategoryIntoleranceAlert2 DetectedIssueCategory = "INT"
	// Proposed therapy may be inappropriate or contraindicated because of a recorded patient allergy to a cross-sensitivity related product. (Allergies are immune based reactions.)
	DetectedIssueCategoryRelatedAllergyAlert2 DetectedIssueCategory = "RALG"
	// Proposed therapy may be inappropriate or contraindicated because of a recorded prior adverse reaction to a cross-sensitivity related product.
	DetectedIssueCategoryRelatedPriorReactionAlert2 DetectedIssueCategory = "RAR"
	// Proposed therapy may be inappropriate or contraindicated because of a recorded patient intolerance to a cross-sensitivity related product. (Intolerances are non-immune based sensitivities.)
	DetectedIssueCategoryRelatedIntoleranceAlert2              DetectedIssueCategory = "RINT"
	DetectedIssueCategoryEndTooLateAlert4                      DetectedIssueCategory = "ALRTENDLATE"
	DetectedIssueCategoryStartTooLateAlert4                    DetectedIssueCategory = "ALRTSTRTLATE"
	DetectedIssueCategoryOutsideRequestedTime                  DetectedIssueCategory = "INTERVAL"
	DetectedIssueCategoryTooSoonWithinFrequencyBasedOnTheUsage DetectedIssueCategory = "MINFREQ"
	// Proposed therapy may interact with an existing or recent drug therapy
	DetectedIssueCategoryDrugInteractionAlert2 DetectedIssueCategory = "DRG"
	// Proposed therapy may interact with existing or recent natural health product therapy
	DetectedIssueCategoryNaturalHealthProductAlert2 DetectedIssueCategory = "NHP"
	// Proposed therapy may interact with a non-prescription drug (e.g. alcohol, tobacco, Aspirin)
	DetectedIssueCategoryNonPrescriptionInteractionAlert2 DetectedIssueCategory = "NONRX"
	// Identifies types of issues detected regarding the performance of a clinical action on a patient.
	DetectedIssueCategoryClinicalActionDetectedIssueCode DetectedIssueCategory = "_ClinicalActionDetectedIssueCode"
	// Supplying the product at this time may be inappropriate or indicate compliance issues with the associated therapy
	DetectedIssueCategorySupplyDetectedIssueCode2 DetectedIssueCategory = "_SupplyDetectedIssueCode"
	DetectedIssueCategoryAlreadyPerformed2        DetectedIssueCategory = "ALLDONE"
	DetectedIssueCategoryFulfillmentAlert2        DetectedIssueCategory = "FULFIL"
	DetectedIssueCategoryHeldSuspendedAlert2      DetectedIssueCategory = "HELD"
	// The patient is receiving a subsequent fill significantly later than would be expected based on the amount previously supplied and the therapy dosage instructions
	DetectedIssueCategoryRefillTooLateAlert2 DetectedIssueCategory = "TOOLATE"
	// The patient is receiving a subsequent fill significantly earlier than would be expected based on the amount previously supplied and the therapy dosage instructions
	DetectedIssueCategoryRefillTooSoonAlert2                    DetectedIssueCategory = "TOOSOON"
	DetectedIssueCategoryNoLongerActionable2                    DetectedIssueCategory = "NOTACTN"
	DetectedIssueCategoryNotEquivalentAlert2                    DetectedIssueCategory = "NOTEQUIV"
	DetectedIssueCategoryEventTimingIncorrectAlert2             DetectedIssueCategory = "TIMING"
	DetectedIssueCategoryNotGenericallyEquivalentAlert2         DetectedIssueCategory = "NOTEQUIVGEN"
	DetectedIssueCategoryNotTherapeuticallyEquivalentAlert2     DetectedIssueCategory = "NOTEQUIVTHER"
	DetectedIssueCategoryOutsideRequestedTime2                  DetectedIssueCategory = "INTERVAL"
	DetectedIssueCategoryTooSoonWithinFrequencyBasedOnTheUsage2 DetectedIssueCategory = "MINFREQ"
)
