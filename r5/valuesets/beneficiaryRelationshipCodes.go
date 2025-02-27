// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/relationship
*/type BeneficiaryRelationshipCodes string

const (
	// The patient is the subscriber (policy holder)
	BeneficiaryRelationshipCodesSelf BeneficiaryRelationshipCodes = "1"
	// The patient is the spouse or equivalent of the subscriber (policy holder)
	BeneficiaryRelationshipCodesSpouse BeneficiaryRelationshipCodes = "2"
	// The patient is the child of the subscriber (policy holder)
	BeneficiaryRelationshipCodesChild BeneficiaryRelationshipCodes = "3"
	// The patient is the common law spouse of the subscriber (policy holder)
	BeneficiaryRelationshipCodesCommonLawSpouse BeneficiaryRelationshipCodes = "4"
	// The patient has some other relationship, such as parent, to the subscriber (policy holder)
	BeneficiaryRelationshipCodesOther BeneficiaryRelationshipCodes = "5"
)
