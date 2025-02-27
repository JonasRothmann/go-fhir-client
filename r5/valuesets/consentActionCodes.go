// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/consentaction
*/type ConsentActionCodes string

const (
	// Gather retrieved information for storage
	ConsentActionCodesCollect ConsentActionCodes = "collect"
	// Retrieval without permitting collection, use or disclosure. e.g., no screen-scraping for collection, use or disclosure (view-only access)
	ConsentActionCodesAccess ConsentActionCodes = "access"
	// Utilize the retrieved information
	ConsentActionCodesUse ConsentActionCodes = "use"
	// Transfer retrieved information
	ConsentActionCodesDisclose ConsentActionCodes = "disclose"
	// Allow retrieval of a patient's information for the purpose of update or rectify
	ConsentActionCodesCorrect ConsentActionCodes = "correct"
)
