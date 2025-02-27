// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/v3-EncounterSpecialCourtesy
  - EXT
  - NRM
  - PRF
  - STF
  - VIP

- http://terminology.hl7.org/CodeSystem/v3-NullFlavor
  - UNK
*/type SpecialCourtesy string

const (
	SpecialCourtesyExtendedCourtesy     SpecialCourtesy = "EXT"
	SpecialCourtesyNormalCourtesy       SpecialCourtesy = "NRM"
	SpecialCourtesyProfessionalCourtesy SpecialCourtesy = "PRF"
	// Courtesies extended to the staff of the entity providing service.
	SpecialCourtesyStaff               SpecialCourtesy = "STF"
	SpecialCourtesyVeryImportantPerson SpecialCourtesy = "VIP"
	// A proper value is applicable, but not known.
	SpecialCourtesyUnknown SpecialCourtesy = "UNK"
)
