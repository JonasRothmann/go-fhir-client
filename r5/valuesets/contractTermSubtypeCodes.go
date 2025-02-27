// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/contracttermsubtypecodes
*/type ContractTermSubtypeCodes string

const (
	// Terms that go to the very root of a contract.
	ContractTermSubtypeCodesCondition ContractTermSubtypeCodes = "condition"
	// Less imperative than a condition, so the contract will survive a breach
	ContractTermSubtypeCodesWarranty ContractTermSubtypeCodes = "warranty"
	// Breach of which might or might not go to the root of the contract depending upon the nature of the breach
	ContractTermSubtypeCodesInnominate ContractTermSubtypeCodes = "innominate"
)
