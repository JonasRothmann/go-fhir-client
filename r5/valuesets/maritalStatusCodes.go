// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/v3-MaritalStatus
- http://terminology.hl7.org/CodeSystem/v3-NullFlavor
  - UNK
*/type MaritalStatusCodes string

const (
	// Marriage contract has been declared null and to not have existed
	MaritalStatusCodesAnnulled MaritalStatusCodes = "A"
	// Marriage contract has been declared dissolved and inactive
	MaritalStatusCodesDivorced MaritalStatusCodes = "D"
	// Subject to an Interlocutory Decree.
	MaritalStatusCodesInterlocutory    MaritalStatusCodes = "I"
	MaritalStatusCodesLegallySeparated MaritalStatusCodes = "L"
	// A current marriage contract is active
	MaritalStatusCodesMarried MaritalStatusCodes = "M"
	// More than 1 current spouse
	MaritalStatusCodesPolygamous MaritalStatusCodes = "P"
	// Person declares that a domestic partner relationship exists.
	MaritalStatusCodesDomesticPartner MaritalStatusCodes = "T"
	// Currently not in a marriage contract.
	MaritalStatusCodesUnmarried MaritalStatusCodes = "U"
	// The spouse has died
	MaritalStatusCodesWidowed MaritalStatusCodes = "W"
	// A proper value is applicable, but not known.
	MaritalStatusCodesUnknown MaritalStatusCodes = "UNK"
)
