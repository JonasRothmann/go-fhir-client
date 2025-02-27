// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/medicinal-product-contact-type
*/type ProductContactType string

const (
	// Proposed Marketing Authorization Holder/Person
	ProductContactTypeProposedMah ProductContactType = "ProposedMAH"
	// Person/Company authorised for Communication on behalf of the Applicant during the Procedure
	ProductContactTypeProcedureContactDuring ProductContactType = "ProcedureContactDuring"
	// Person/Company authorised for Communication between MAH and Authorities after Authorization
	ProductContactTypeProcedureContactAfter ProductContactType = "ProcedureContactAfter"
	// Qualified Person Responsible for Pharmacovigilance
	ProductContactTypeQualifiedPersonResponsibleForPharmacovigilance ProductContactType = "QPPV"
	// Pharmacovigilance Enquiry Information
	ProductContactTypePvEnquiries ProductContactType = "PVEnquiries"
)
