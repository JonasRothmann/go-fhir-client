// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/contract-content-derivative
*/type ContractContentDerivationCodes string

const (
	// Content derivative that conveys sufficient information needed to register the source basal content from which it is derived.  This derivative content may be used to register the basal content as it changes status in its lifecycle.  For example, content registration may occur when the basal content is created, updated, inactive, or deleted.
	ContractContentDerivationCodesRegistration ContractContentDerivationCodes = "registration"
	// A content derivative that conveys sufficient information to locate and retrieve the content.
	ContractContentDerivationCodesRetrieval ContractContentDerivationCodes = "retrieval"
	// Content derivative that has less than full fidelity to the basal information source from which it was 'transcribed'. It provides recipients with the full content representation they may require for compliance purposes, and typically include a reference to or an attached unstructured representation for recipients needing an exact copy of the legal agreement.
	ContractContentDerivationCodesStatement ContractContentDerivationCodes = "statement"
	// A Content Derivative that conveys sufficient information to determine the authorized entities with which the content may be shared.
	ContractContentDerivationCodesShareable ContractContentDerivationCodes = "shareable"
)
