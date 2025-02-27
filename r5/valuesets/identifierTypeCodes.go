// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/v2-0203
  - DL
  - PPN
  - BRN
  - MR
  - MCN
  - EN
  - TAX
  - NIIP
  - PRN
  - MD
  - DR
  - ACSN
  - UDI
  - SNO
  - SB
  - PLAC
  - FILL
  - JHN
*/type IdentifierTypeCodes string

const (
	// Accession Identifier
	IdentifierTypeCodesAccessionId IdentifierTypeCodes = "ACSN"
	// Breed Registry Number
	IdentifierTypeCodesBreedRegistryNumber IdentifierTypeCodes = "BRN"
	// Driver's license number
	IdentifierTypeCodesDriverSLicenseNumber IdentifierTypeCodes = "DL"
	// Donor Registration Number
	IdentifierTypeCodesDonorRegistrationNumber IdentifierTypeCodes = "DR"
	// Employer number
	IdentifierTypeCodesEmployerNumber IdentifierTypeCodes = "EN"
	// An identifier for a request where the identifier is issued by the person, or service, that produces the observations or fulfills the request.
	IdentifierTypeCodesFillerIdentifier IdentifierTypeCodes = "FILL"
	// Jurisdictional health number
	IdentifierTypeCodesJurisdictionalHealthNumber IdentifierTypeCodes = "JHN"
	// Microchip Number
	IdentifierTypeCodesMicrochipNumber IdentifierTypeCodes = "MCN"
	// An identifier that is unique to a medical doctor within the jurisdiction of a licensing board.
	IdentifierTypeCodesMedicalLicenseNumber IdentifierTypeCodes = "MD"
	// An identifier that is unique to a patient within a set of medical records, not necessarily unique within an application.
	IdentifierTypeCodesMedicalRecordNumber IdentifierTypeCodes = "MR"
	// National Insurance Payor Identifier (Payor)
	IdentifierTypeCodesNationalInsurancePayorIdentifierPayor IdentifierTypeCodes = "NIIP"
	// An identifier for a request where the identifier is issued by the person or service making the request.
	IdentifierTypeCodesPlacerIdentifier IdentifierTypeCodes = "PLAC"
	// A unique number assigned to the document affirming that a person is a citizen of the country.
	IdentifierTypeCodesPassportNumber IdentifierTypeCodes = "PPN"
	// A number that is unique to an individual provider, a provider group or an organization within an Assigning Authority.
	IdentifierTypeCodesProviderNumber IdentifierTypeCodes = "PRN"
	// An identifier issued by a governmental organization to a person to identify the person should they apply for or receive social services and/or benefits
	IdentifierTypeCodesSocialBeneficiaryIdentifier IdentifierTypeCodes = "SB"
	// An identifier affixed to an item by the manufacturer when it is first made, where each item has a different identifier.
	IdentifierTypeCodesSerialNumber IdentifierTypeCodes = "SNO"
	// Tax ID number
	IdentifierTypeCodesTaxIdNumber IdentifierTypeCodes = "TAX"
	// An identifier assigned to a device using the Unique Device Identification framework as defined by IMDRF (http://imdrf.org).
	IdentifierTypeCodesUniversalDeviceIdentifier IdentifierTypeCodes = "UDI"
)
