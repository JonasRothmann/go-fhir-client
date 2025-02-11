//go:generate go-enum --marshal

package valuesets

/*
https://www.hl7.org/fhir/valueset-marital-status.html

ENUM(

	Unset = "",
	Annulled = "A",
	Divorced = "D",
	Interlocutory = "I",
	LegallySeparated = "L",
	Married = "M",
	CommonLaw = "C",
	Polygamous = "P",
	DomesticPartner = "T",
	Unmarried = "U",
	NeverMarried = "S",
	Widowed = "W",
	Unknown = "UNK"

)
*/
type MaritalStatus string
