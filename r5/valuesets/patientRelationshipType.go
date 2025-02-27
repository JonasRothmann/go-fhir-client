// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/v3-ParticipationType
  - WIT

- http://terminology.hl7.org/CodeSystem/v3-RoleClass
  - NOT
  - ECON
  - NOK
  - GUARD
  - DEPEN
  - EMP
  - GUAR
  - CAREGIVER

- http://terminology.hl7.org/CodeSystem/v2-0131
  - E
  - O
  - U

- http://terminology.hl7.org/CodeSystem/v3-RoleCode
  - INTPRTER
  - POWATT
  - DPOWATT
  - HPOWATT
  - SPOWATT
  - BILL
*/type PatientRelationshipType string

const (
	// Employer
	PatientRelationshipTypeEmployer PatientRelationshipType = "E"
	// Other
	PatientRelationshipTypeOther PatientRelationshipType = "O"
	// Unknown
	PatientRelationshipTypeUnknown PatientRelationshipType = "U"
	// A contact role used to identify a person within a Provider organization that can be contacted for billing purposes (e.g. about the content of the Invoice).
	PatientRelationshipTypeBillingContact PatientRelationshipType = "BILL"
	// An entity which converts spoken or written language into the language of key participants in an event such as when a provider is obtaining a patient's consent to treatment or permission to disclose information.
	PatientRelationshipTypeInterpreter PatientRelationshipType = "INTPRTER"
	// A relationship between two people in which one person authorizes another, usually a family member or relative, to act for him or her in a manner which is a legally binding upon the person giving such authority as if he or she personally were to do the acts that is often limited in the kinds of powers that can be assigned. Unlike ordinary powers of attorney, durable powers can survive for long periods of time, and again, unlike standard powers of attorney, durable powers can continue after incompetency.
	PatientRelationshipTypeDurablePowerOfAttorney PatientRelationshipType = "DPOWATT"
	// A relationship between two people in which one person authorizes another to act for him or her in a manner which is a legally binding upon the person giving such authority as if he or she personally were to do the acts that continues (by its terms) to be effective even though the grantor has become mentally incompetent after signing the document.
	PatientRelationshipTypeHealthcarePowerOfAttorney PatientRelationshipType = "HPOWATT"
	// A relationship between two people in which one person authorizes another to act for him or her in a manner which is a legally binding upon the person giving such authority as if he or she personally were to do the acts that is often limited in the kinds of powers that can be assigned.
	PatientRelationshipTypeSpecialPowerOfAttorney PatientRelationshipType = "SPOWATT"
	// A relationship between two people in which one person authorizes another to act for him in a manner which is a legally binding upon the person giving such authority as if he or she personally were to do the acts.
	PatientRelationshipTypePowerOfAttorney PatientRelationshipType = "POWATT"
)
