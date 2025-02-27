// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/payeetype
*/type ClaimPayeeTypeCodes string

const (
	// The subscriber (policy holder) will be reimbursed.
	ClaimPayeeTypeCodesSubscriber ClaimPayeeTypeCodes = "subscriber"
	// Any benefit payable will be paid to the provider (Assignment of Benefit).
	ClaimPayeeTypeCodesProvider ClaimPayeeTypeCodes = "provider"
	// The beneficiary (patient) will be reimbursed.
	ClaimPayeeTypeCodesBeneficiary ClaimPayeeTypeCodes = "beneficiary"
	// Any benefit payable will be paid to a third party such as a guarantor.
	ClaimPayeeTypeCodesOther ClaimPayeeTypeCodes = "other"
)
