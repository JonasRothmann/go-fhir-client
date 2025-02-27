// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/claim-decision-reason
*/type ClaimAdjudicationDecisionReasonCodes string

const (
	// The payer has determined this product, service, or procedure as not medically necessary.
	ClaimAdjudicationDecisionReasonCodesNotMedicallyNecessary ClaimAdjudicationDecisionReasonCodes = "0001"
	// Prior authorization was not obtained prior to providing the product, service, or procedure.
	ClaimAdjudicationDecisionReasonCodesPriorAuthorizationNotObtained ClaimAdjudicationDecisionReasonCodes = "0002"
	// This provider is considered out-of-network by the payer for this plan.
	ClaimAdjudicationDecisionReasonCodesProviderOutOfNetwork ClaimAdjudicationDecisionReasonCodes = "0003"
	// The payer has determined this product, service, or procedure is not consistent with the patient's age.
	ClaimAdjudicationDecisionReasonCodesServiceInconsistentWithPatientAge ClaimAdjudicationDecisionReasonCodes = "0004"
	// The patient or subscriber benefit's have been exceeded.
	ClaimAdjudicationDecisionReasonCodesBenefitLimitsExceeded ClaimAdjudicationDecisionReasonCodes = "0005"
)
