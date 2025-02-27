// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/organization-role
*/type OrganizationAffiliationRole string

const (
	// An organization that delivers care services (e.g. hospitals, clinics, community and social services, etc.).
	OrganizationAffiliationRoleProvider OrganizationAffiliationRole = "provider"
	// An organization such as a public health agency, community/social services provider, etc.
	OrganizationAffiliationRoleAgency OrganizationAffiliationRole = "agency"
	// An organization providing research-related services such as conducting research, recruiting research participants, analyzing data, etc.
	OrganizationAffiliationRoleResearch OrganizationAffiliationRole = "research"
	// An organization providing reimbursement, payment, or related services
	OrganizationAffiliationRolePayer OrganizationAffiliationRole = "payer"
	// An organization providing diagnostic testing/laboratory services
	OrganizationAffiliationRoleDiagnostics OrganizationAffiliationRole = "diagnostics"
	// An organization that provides medical supplies (e.g. medical devices, equipment, pharmaceutical products, etc.)
	OrganizationAffiliationRoleSupplier OrganizationAffiliationRole = "supplier"
	// An organization that facilitates electronic clinical data exchange between entities
	OrganizationAffiliationRoleHiehio OrganizationAffiliationRole = "HIE/HIO"
	// A type of non-ownership relationship between entities (encompasses partnerships, collaboration, joint ventures, etc.)
	OrganizationAffiliationRoleMember OrganizationAffiliationRole = "member"
)
