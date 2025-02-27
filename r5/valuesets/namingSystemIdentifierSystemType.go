// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/v2-0203
*/type NamingSystemIdentifierSystemType string

const (
	// Identifier that has been assigned by an accreditation or certification organization in specific fields, indicating a recognized skill
	NamingSystemIdentifierSystemTypeAccreditationCertificationIdentifier NamingSystemIdentifierSystemType = "AC"
	// Accession Identifier
	NamingSystemIdentifierSystemTypeAccessionId NamingSystemIdentifierSystemType = "ACSN"
	// A numbering system for the official identification of individual animals in the United States that provides a nationally unique identification number for each animal. The first two numbers on a tag are the numbers assigned to a specific State.
	NamingSystemIdentifierSystemTypeAnimalIdentificationNumberUsOfficial NamingSystemIdentifierSystemType = "AIN"
	// American Express
	NamingSystemIdentifierSystemTypeAmericanExpress NamingSystemIdentifierSystemType = "AM"
	// A physician identifier assigned by the AMA.
	NamingSystemIdentifierSystemTypeAmericanMedicalAssociationNumber NamingSystemIdentifierSystemType = "AMA"
	// Account An identifier that is unique to an account.
	NamingSystemIdentifierSystemTypeAccountNumber NamingSystemIdentifierSystemType = "AN"
	// A more precise definition of an account number
	NamingSystemIdentifierSystemTypeAccountNumberCreditor NamingSystemIdentifierSystemType = "ANC"
	// A more precise definition of an account number
	NamingSystemIdentifierSystemTypeAccountNumberDebitor NamingSystemIdentifierSystemType = "AND"
	// An identifier for a living subject whose real identity is protected or suppressed
	NamingSystemIdentifierSystemTypeAnonymousIdentifier NamingSystemIdentifierSystemType = "ANON"
	// Temporary version of an Account Number
	NamingSystemIdentifierSystemTypeTemporaryAccountNumber NamingSystemIdentifierSystemType = "ANT"
	// An identifier that is unique to an advanced practice registered nurse within the jurisdiction of a certifying board
	NamingSystemIdentifierSystemTypeAdvancedPracticeRegisteredNurseNumber NamingSystemIdentifierSystemType = "APRN"
	// A unique identifier for the ancestor specimen.
	NamingSystemIdentifierSystemTypeAncestorSpecimenId NamingSystemIdentifierSystemType = "ASID"
	// Bank Account Number
	NamingSystemIdentifierSystemTypeBankAccountNumber NamingSystemIdentifierSystemType = "BA"
	// An identifier that is unique to a person's bank card
	NamingSystemIdentifierSystemTypeBankCardNumber NamingSystemIdentifierSystemType = "BC"
	// The identifier used within the jurisdictional vital records office file system as an auxiliary means of accessing the record associated with the birth certificate.
	NamingSystemIdentifierSystemTypeBirthCertificateFileNumber NamingSystemIdentifierSystemType = "BCFN"
	// A number associated with a document identifying the event of a person's birth
	NamingSystemIdentifierSystemTypeBirthCertificate NamingSystemIdentifierSystemType = "BCT"
	// An identifier unique within the Assigning Authority that is the official legal record of a person's birth.
	NamingSystemIdentifierSystemTypeBirthRegistryNumber NamingSystemIdentifierSystemType = "BR"
	// Breed Registry Number
	NamingSystemIdentifierSystemTypeBreedRegistryNumber NamingSystemIdentifierSystemType = "BRN"
	// Primary physician office number
	NamingSystemIdentifierSystemTypePrimaryPhysicianOfficeNumber NamingSystemIdentifierSystemType = "BSNR"
	// An identifier for the consumer (e.g., patient, caregiver) for an application such as a portal or App.
	NamingSystemIdentifierSystemTypeConsumerApplicationAccountIdentifier NamingSystemIdentifierSystemType = "CAAI"
	// Cost Center number
	NamingSystemIdentifierSystemTypeCostCenterNumber NamingSystemIdentifierSystemType = "CC"
	// A number associated with a document identifying a person's legal change of name.
	NamingSystemIdentifierSystemTypeChangeOfNameDocument NamingSystemIdentifierSystemType = "CONM"
	// County number
	NamingSystemIdentifierSystemTypeCountyNumber NamingSystemIdentifierSystemType = "CY"
	// A number assigned by a person's country of residence to identify a person's citizenship.
	NamingSystemIdentifierSystemTypeCitizenshipCard NamingSystemIdentifierSystemType = "CZ"
	// The identifier assigned to a death certificate, and printed on the death certificate when issued by a jurisdictional vital records office
	NamingSystemIdentifierSystemTypeDeathCertificateId NamingSystemIdentifierSystemType = "DC"
	// The identifier used within the jurisdictional vital records office file system as an auxiliary means of accessing the record associated with the death certificate.
	NamingSystemIdentifierSystemTypeDeathCertificateFileNumber NamingSystemIdentifierSystemType = "DCFN"
	// An identifier that is unique to a dentist within the jurisdiction of the licensing board
	NamingSystemIdentifierSystemTypeDentistLicenseNumber NamingSystemIdentifierSystemType = "DDS"
	// An identifier for an individual or organization relative to controlled substance regulation and transactions.
	NamingSystemIdentifierSystemTypeDrugEnforcementAdministrationRegistrationNumber NamingSystemIdentifierSystemType = "DEA"
	// An identifier issued to a health care provider authorizing the person to write drug orders
	NamingSystemIdentifierSystemTypeDrugFurnishingOrPrescriptiveAuthorityNumber NamingSystemIdentifierSystemType = "DFN"
	// Diner's Club card
	NamingSystemIdentifierSystemTypeDinerSClubCard NamingSystemIdentifierSystemType = "DI"
	// Driver's license number
	NamingSystemIdentifierSystemTypeDriverSLicenseNumber NamingSystemIdentifierSystemType = "DL"
	// Doctor number
	NamingSystemIdentifierSystemTypeDoctorNumber NamingSystemIdentifierSystemType = "DN"
	// An identifier that is unique to an osteopath within the jurisdiction of a licensing board.
	NamingSystemIdentifierSystemTypeOsteopathicLicenseNumber NamingSystemIdentifierSystemType = "DO"
	// A number assigned to a diplomatic passport.
	NamingSystemIdentifierSystemTypeDiplomaticPassport NamingSystemIdentifierSystemType = "DP"
	// An identifier that is unique to a podiatrist within the jurisdiction of the licensing board.
	NamingSystemIdentifierSystemTypePodiatristLicenseNumber NamingSystemIdentifierSystemType = "DPM"
	// Donor Registration Number
	NamingSystemIdentifierSystemTypeDonorRegistrationNumber NamingSystemIdentifierSystemType = "DR"
	// Discover Card
	NamingSystemIdentifierSystemTypeDiscoverCard NamingSystemIdentifierSystemType = "DS"
	// Unique Identifier that groups several orders that are to be performed together.
	NamingSystemIdentifierSystemTypeDiagnosticStudyGroup NamingSystemIdentifierSystemType = "DSG"
	// A number that uniquely identifies an employee to an employer.
	NamingSystemIdentifierSystemTypeEmployeeNumber NamingSystemIdentifierSystemType = "EI"
	// Employer number
	NamingSystemIdentifierSystemTypeEmployerNumber NamingSystemIdentifierSystemType = "EN"
	// An identifier that is unique to a staff member within an enterprise (as identified by the Assigning Authority).
	NamingSystemIdentifierSystemTypeStaffEnterpriseNumber NamingSystemIdentifierSystemType = "ESN"
	// The identifier assigned to a fetal death report, and printed on the fetal death report when issued by a jurisdictional vital records office
	NamingSystemIdentifierSystemTypeFetalDeathReportId NamingSystemIdentifierSystemType = "FDR"
	// The identifier used within the jurisdictional vital records office file system as an auxiliary means of accessing the record associated with the fetal death report certificate.
	NamingSystemIdentifierSystemTypeFetalDeathReportFileNumber NamingSystemIdentifierSystemType = "FDRFN"
	// Unique identifier assigned to a group of orders by the filler application.
	NamingSystemIdentifierSystemTypeFillerGroupNumber NamingSystemIdentifierSystemType = "FGN"
	// Facility ID
	NamingSystemIdentifierSystemTypeFacilityId NamingSystemIdentifierSystemType = "FI"
	// An identifier for a request where the identifier is issued by the person, or service, that produces the observations or fulfills the request.
	NamingSystemIdentifierSystemTypeFillerIdentifier NamingSystemIdentifierSystemType = "FILL"
	// Guarantor internal identifier
	NamingSystemIdentifierSystemTypeGuarantorInternalIdentifier NamingSystemIdentifierSystemType = "GI"
	// Identifier that can be used to unambiguously describe a specific group of animals.
	NamingSystemIdentifierSystemTypeAnimalGroupIdentifierUsOfficial NamingSystemIdentifierSystemType = "GIN"
	// General ledger number
	NamingSystemIdentifierSystemTypeGeneralLedgerNumber NamingSystemIdentifierSystemType = "GL"
	// Guarantor external  identifier
	NamingSystemIdentifierSystemTypeGuarantorExternalIdentifier NamingSystemIdentifierSystemType = "GN"
	// Health Card Number
	NamingSystemIdentifierSystemTypeHealthCardNumber NamingSystemIdentifierSystemType = "HC"
	// A number assigned to a member of an indigenous or aboriginal group outside of Canada.
	NamingSystemIdentifierSystemTypeIndigenousAboriginal NamingSystemIdentifierSystemType = "IND"
	// An IRI string that can be prepended to the code to obtain a concept IRI for RDF applications. This should be a valid, absolute IRI as defined in RFC 3987. See https://hl7.org/fhir/rdf.html#iri-stem for details on how this value may be used.
	NamingSystemIdentifierSystemTypeAnIriStem NamingSystemIdentifierSystemType = "IRISTEM"
	// Jurisdictional health number
	NamingSystemIdentifierSystemTypeJurisdictionalHealthNumber NamingSystemIdentifierSystemType = "JHN"
	// A laboratory accession id is used in the laboratory domain.
	NamingSystemIdentifierSystemTypeLaboratoryAccessionId NamingSystemIdentifierSystemType = "LACSN"
	// Lifelong physician number
	NamingSystemIdentifierSystemTypeLifelongPhysicianNumber NamingSystemIdentifierSystemType = "LANR"
	// Labor and industries number
	NamingSystemIdentifierSystemTypeLaborAndIndustriesNumber NamingSystemIdentifierSystemType = "LI"
	// Labor and industries number.  Note that this was introduced erroneously (with an ampersand in the code value) many years ago.
	NamingSystemIdentifierSystemTypeLaborAndIndustriesNumber2 NamingSystemIdentifierSystemType = "L&I"
	// License number
	NamingSystemIdentifierSystemTypeLicenseNumber NamingSystemIdentifierSystemType = "LN"
	// Local Registry ID
	NamingSystemIdentifierSystemTypeLocalRegistryId NamingSystemIdentifierSystemType = "LR"
	// Patient Medicaid number
	NamingSystemIdentifierSystemTypePatientMedicaidNumber NamingSystemIdentifierSystemType = "MA"
	// An identifier for the insured of an insurance policy (this insured always has a subscriber), usually assigned by the insurance carrier.
	NamingSystemIdentifierSystemTypeMemberNumber NamingSystemIdentifierSystemType = "MB"
	// Patient's Medicare number
	NamingSystemIdentifierSystemTypePatientSMedicareNumber NamingSystemIdentifierSystemType = "MC"
	// Practitioner Medicaid number
	NamingSystemIdentifierSystemTypePractitionerMedicaidNumber NamingSystemIdentifierSystemType = "MCD"
	// Microchip Number
	NamingSystemIdentifierSystemTypeMicrochipNumber NamingSystemIdentifierSystemType = "MCN"
	// Practitioner Medicare number
	NamingSystemIdentifierSystemTypePractitionerMedicareNumber NamingSystemIdentifierSystemType = "MCR"
	// A number associated with a document identifying the event of a person's marriage.
	NamingSystemIdentifierSystemTypeMarriageCertificate NamingSystemIdentifierSystemType = "MCT"
	// An identifier that is unique to a medical doctor within the jurisdiction of a licensing board.
	NamingSystemIdentifierSystemTypeMedicalLicenseNumber NamingSystemIdentifierSystemType = "MD"
	// A number assigned to an individual who has had military duty, but is not currently on active duty. The number is assigned by the DOD or Veterans' Affairs (VA).
	NamingSystemIdentifierSystemTypeMilitaryIdNumber NamingSystemIdentifierSystemType = "MI"
	// An identifier that is unique to a patient within a set of medical records, not necessarily unique within an application.
	NamingSystemIdentifierSystemTypeMedicalRecordNumber NamingSystemIdentifierSystemType = "MR"
	// Temporary version of a Medical Record Number
	NamingSystemIdentifierSystemTypeTemporaryMedicalRecordNumber NamingSystemIdentifierSystemType = "MRT"
	// MasterCard
	NamingSystemIdentifierSystemTypeMasterCard NamingSystemIdentifierSystemType = "MS"
	// Secondary physician office number
	NamingSystemIdentifierSystemTypeSecondaryPhysicianOfficeNumber NamingSystemIdentifierSystemType = "NBSNR"
	// A number associated with a document identifying a person's retention of citizenship in a particular country.
	NamingSystemIdentifierSystemTypeNaturalizationCertificate NamingSystemIdentifierSystemType = "NCT"
	// National employer identifier
	NamingSystemIdentifierSystemTypeNationalEmployerIdentifier NamingSystemIdentifierSystemType = "NE"
	// National Health Plan Identifier
	NamingSystemIdentifierSystemTypeNationalHealthPlanIdentifier NamingSystemIdentifierSystemType = "NH"
	// National unique individual identifier
	NamingSystemIdentifierSystemTypeNationalUniqueIndividualIdentifier NamingSystemIdentifierSystemType = "NI"
	// National Insurance Organization Identifier
	NamingSystemIdentifierSystemTypeNationalInsuranceOrganizationIdentifier NamingSystemIdentifierSystemType = "NII"
	// National Insurance Payor Identifier (Payor)
	NamingSystemIdentifierSystemTypeNationalInsurancePayorIdentifierPayor NamingSystemIdentifierSystemType = "NIIP"
	// National Person Identifier where the xxx is the ISO table 3166 3-character (alphabetic) country code
	NamingSystemIdentifierSystemTypeNNxxx NamingSystemIdentifierSystemType = "NNxxx"
	// An identifier that is unique to a nurse practitioner within the jurisdiction of a certifying board.
	NamingSystemIdentifierSystemTypeNursePractitionerNumber NamingSystemIdentifierSystemType = "NP"
	// National provider identifier
	NamingSystemIdentifierSystemTypeNationalProviderIdentifier NamingSystemIdentifierSystemType = "NPI"
	// Unique and persistent identifier for an observation instance
	NamingSystemIdentifierSystemTypeObservationInstanceIdentifier NamingSystemIdentifierSystemType = "OBI"
	// A number that is unique to an individual optometrist within the jurisdiction of the licensing board.
	NamingSystemIdentifierSystemTypeOptometristLicenseNumber NamingSystemIdentifierSystemType = "OD"
	// An identifier that is unique to a physician assistant within the jurisdiction of a licensing board
	NamingSystemIdentifierSystemTypePhysicianAssistantNumber NamingSystemIdentifierSystemType = "PA"
	// A number identifying a person on parole.
	NamingSystemIdentifierSystemTypeParoleCard NamingSystemIdentifierSystemType = "PC"
	// A number assigned to individual who is incarcerated.
	NamingSystemIdentifierSystemTypePenitentiaryCorrectionalInstitutionNumber NamingSystemIdentifierSystemType = "PCN"
	// An identifier that is unique to a living subject within an enterprise (as identified by the Assigning Authority).
	NamingSystemIdentifierSystemTypeLivingSubjectEnterpriseNumber NamingSystemIdentifierSystemType = "PE"
	// Pension Number
	NamingSystemIdentifierSystemTypePensionNumber NamingSystemIdentifierSystemType = "PEN"
	// Unique identifier assigned to a group of orders by the placer application.
	NamingSystemIdentifierSystemTypePlacerGroupNumber NamingSystemIdentifierSystemType = "PGN"
	// Identifier assigned to a person during a case investigation as part of a public health event
	NamingSystemIdentifierSystemTypePublicHealthCaseIdentifier NamingSystemIdentifierSystemType = "PHC"
	// Identifier assigned to an event of interest to public health
	NamingSystemIdentifierSystemTypePublicHealthEventIdentifier NamingSystemIdentifierSystemType = "PHE"
	// An identifier for a person working at a public health agency (PHA),  assigned or issued by the agency
	NamingSystemIdentifierSystemTypePublicHealthOfficialId NamingSystemIdentifierSystemType = "PHO"
	// A number that is unique to a patient within an Assigning Authority.
	NamingSystemIdentifierSystemTypePatientInternalIdentifier NamingSystemIdentifierSystemType = "PI"
	// Identifier that uniquely identifies a geographic location in the US.
	NamingSystemIdentifierSystemTypePremisesIdentifierNumberUsOfficial NamingSystemIdentifierSystemType = "PIN"
	// An identifier for a request where the identifier is issued by the person or service making the request.
	NamingSystemIdentifierSystemTypePlacerIdentifier NamingSystemIdentifierSystemType = "PLAC"
	// A number that is unique to a living subject within an Assigning Authority.
	NamingSystemIdentifierSystemTypePersonNumber NamingSystemIdentifierSystemType = "PN"
	// Temporary version of a Living Subject Number.
	NamingSystemIdentifierSystemTypeTemporaryLivingSubjectNumber NamingSystemIdentifierSystemType = "PNT"
	// Medicare/CMS Performing Provider Identification Number
	NamingSystemIdentifierSystemTypeMedicareCmsPerformingProviderIdentificationNumber NamingSystemIdentifierSystemType = "PPIN"
	// A unique number assigned to the document affirming that a person is a citizen of the country.
	NamingSystemIdentifierSystemTypePassportNumber NamingSystemIdentifierSystemType = "PPN"
	// Permanent Resident Card Number
	NamingSystemIdentifierSystemTypePermanentResidentCardNumber NamingSystemIdentifierSystemType = "PRC"
	// A number that is unique to an individual provider, a provider group or an organization within an Assigning Authority.
	NamingSystemIdentifierSystemTypeProviderNumber NamingSystemIdentifierSystemType = "PRN"
	// Patient external identifier
	NamingSystemIdentifierSystemTypePatientExternalIdentifier NamingSystemIdentifierSystemType = "PT"
	// QA number
	NamingSystemIdentifierSystemTypeQaNumber NamingSystemIdentifierSystemType = "QA"
	// A generalized resource identifier.
	NamingSystemIdentifierSystemTypeResourceIdentifier NamingSystemIdentifierSystemType = "RI"
	// An identifier that is unique to a registered nurse within the jurisdiction of the licensing board.
	NamingSystemIdentifierSystemTypeRegisteredNurseNumber NamingSystemIdentifierSystemType = "RN"
	// An identifier that is unique to a pharmacist within the jurisdiction of the licensing board.
	NamingSystemIdentifierSystemTypePharmacistLicenseNumber NamingSystemIdentifierSystemType = "RPH"
	// An identifier for an individual enrolled with the Railroad Retirement Administration.  Analogous to, but distinct from, a Social Security Number
	NamingSystemIdentifierSystemTypeRailroadRetirementNumber NamingSystemIdentifierSystemType = "RR"
	// Regional registry ID
	NamingSystemIdentifierSystemTypeRegionalRegistryId NamingSystemIdentifierSystemType = "RRI"
	// Railroad Retirement Provider
	NamingSystemIdentifierSystemTypeRailroadRetirementProvider NamingSystemIdentifierSystemType = "RRP"
	// The accession number for the BioSample data repository at the National Center for Biotechnology Information (NCBI)
	NamingSystemIdentifierSystemTypeSamnAccessionNumber NamingSystemIdentifierSystemType = "SAMN"
	// An identifier issued by a governmental organization to a person to identify the person should they apply for or receive social services and/or benefits
	NamingSystemIdentifierSystemTypeSocialBeneficiaryIdentifier NamingSystemIdentifierSystemType = "SB"
	// Identifier for a specimen.
	NamingSystemIdentifierSystemTypeSpecimenId NamingSystemIdentifierSystemType = "SID"
	// State license
	NamingSystemIdentifierSystemTypeStateLicense NamingSystemIdentifierSystemType = "SL"
	// An identifier for a subscriber of an insurance policy which is unique for, and usually assigned by, the insurance carrier.
	NamingSystemIdentifierSystemTypeSubscriberNumber NamingSystemIdentifierSystemType = "SN"
	// The identifier on a Newborn Screening Dried Bloodspot (NDBS) card that is assigned by the state which provided the sample collection cards and to whom this information must be reported
	NamingSystemIdentifierSystemTypeStateAssignedNdbsCardIdentifier NamingSystemIdentifierSystemType = "SNBSN"
	// An identifier affixed to an item by the manufacturer when it is first made, where each item has a different identifier.
	NamingSystemIdentifierSystemTypeSerialNumber NamingSystemIdentifierSystemType = "SNO"
	// A number associated with a permit identifying a person who is a resident of a jurisdiction for the purpose of education.
	NamingSystemIdentifierSystemTypeStudyPermit NamingSystemIdentifierSystemType = "SP"
	// State registry ID
	NamingSystemIdentifierSystemTypeStateRegistryId NamingSystemIdentifierSystemType = "SR"
	// The accession number generated by the Sequence Read Archive (SRA) at the National Center for Biotechnology Information (NCBI) when sequence data are uploaded to NCBI.
	NamingSystemIdentifierSystemTypeSraAccessionNumber NamingSystemIdentifierSystemType = "SRX"
	// Social Security number
	NamingSystemIdentifierSystemTypeSocialSecurityNumber NamingSystemIdentifierSystemType = "SS"
	// Identifier assigned to a package being shipped
	NamingSystemIdentifierSystemTypeShipmentTrackingNumber NamingSystemIdentifierSystemType = "STN"
	// Tax ID number
	NamingSystemIdentifierSystemTypeTaxIdNumber NamingSystemIdentifierSystemType = "TAX"
	// A number assigned to a member of an indigenous group in Canada.
	NamingSystemIdentifierSystemTypeTreatyNumberCanada NamingSystemIdentifierSystemType = "TN"
	// A number associated with a document identifying a person's temporary permanent resident status.
	NamingSystemIdentifierSystemTypeTemporaryPermanentResidentCanada NamingSystemIdentifierSystemType = "TPR"
	// The license number used during training.
	NamingSystemIdentifierSystemTypeTrainingLicenseNumber NamingSystemIdentifierSystemType = "TRL"
	// Unspecified identifier
	NamingSystemIdentifierSystemTypeUnspecifiedIdentifier NamingSystemIdentifierSystemType = "U"
	// An identifier assigned to a device using the Unique Device Identification framework as defined by IMDRF (http://imdrf.org).
	NamingSystemIdentifierSystemTypeUniversalDeviceIdentifier NamingSystemIdentifierSystemType = "UDI"
	// An identifier for a provider within the CMS/Medicare program.  A globally unique identifier for the provider in the Medicare program.
	NamingSystemIdentifierSystemTypeMedicareCmsFormerlyHcfasUniversalPhysicianIdentificationNumbers NamingSystemIdentifierSystemType = "UPIN"
	// A unique identifier for a specimen.
	NamingSystemIdentifierSystemTypeUniqueSpecimenId NamingSystemIdentifierSystemType = "USID"
	// Visit number
	NamingSystemIdentifierSystemTypeVisitNumber NamingSystemIdentifierSystemType = "VN"
	// A number associated with a document identifying a person as a visitor of a jurisdiction or country.
	NamingSystemIdentifierSystemTypeVisitorPermit NamingSystemIdentifierSystemType = "VP"
	// VISA
	NamingSystemIdentifierSystemTypeVisa NamingSystemIdentifierSystemType = "VS"
	// WIC identifier
	NamingSystemIdentifierSystemTypeWicIdentifier NamingSystemIdentifierSystemType = "WC"
	// Workers' Comp Number
	NamingSystemIdentifierSystemTypeWorkersCompNumber NamingSystemIdentifierSystemType = "WCN"
	// A number associated with a permit for a person who is granted permission to work in a country for a specified time period.
	NamingSystemIdentifierSystemTypeWorkPermit NamingSystemIdentifierSystemType = "WP"
	// National unique health plan identifier required by the US Department of Health and Human Services, Centers for Medicare and Medicaid Services (CMS) in the US Realm.
	NamingSystemIdentifierSystemTypeHealthPlanIdentifier NamingSystemIdentifierSystemType = "XV"
	// Organization identifier
	NamingSystemIdentifierSystemTypeOrganizationIdentifier NamingSystemIdentifierSystemType = "XX"
)
