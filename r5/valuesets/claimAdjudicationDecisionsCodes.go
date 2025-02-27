// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/claim-decision
*/type ClaimAdjudicationDecisionsCodes string

const (
	// The claim, or individual services, are not approved for any payment. This may also be known as 'rejected'.
	ClaimAdjudicationDecisionsCodesDenied ClaimAdjudicationDecisionsCodes = "denied"
	// The claim, or individual services, are approved as submitted.
	ClaimAdjudicationDecisionsCodesApproved ClaimAdjudicationDecisionsCodes = "approved"
	// The claim, or individual services, are approved at an amount less than as submitted.
	ClaimAdjudicationDecisionsCodesPartial ClaimAdjudicationDecisionsCodes = "partial"
	// The adjudication processing is not complete. This may be due to requiring manual review or receipt of additional information.
	ClaimAdjudicationDecisionsCodesPending ClaimAdjudicationDecisionsCodes = "pending"
)
