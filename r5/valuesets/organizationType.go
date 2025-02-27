// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/organization-type
*/type OrganizationType string

const (
	// An organization that provides healthcare services.
	OrganizationTypeProv OrganizationType = "prov"
	// A department or ward within a hospital (Generally is not applicable to top level organizations)
	OrganizationTypeDept OrganizationType = "dept"
	// An organizational team is usually a grouping of practitioners that perform a specific function within an organization (which could be a top level organization, or a department).
	OrganizationTypeTeam OrganizationType = "team"
	// A political body, often used when including organization records for government bodies such as a Federal Government, State or Local Government.
	OrganizationTypeGovt OrganizationType = "govt"
	// A company that provides insurance to its subscribers that may include healthcare related policies.
	OrganizationTypeIns OrganizationType = "ins"
	// A company, charity, or governmental organization, which processes claims and/or issues payments to providers on behalf of patients or groups of patients.
	OrganizationTypePay OrganizationType = "pay"
	// An educational institution that provides education or research facilities.
	OrganizationTypeEdu OrganizationType = "edu"
	// An organization that is identified as a part of a religious institution.
	OrganizationTypeReli OrganizationType = "reli"
	// An organization that is identified as a Pharmaceutical/Clinical Research Sponsor.
	OrganizationTypeCrs OrganizationType = "crs"
	// An un-incorporated community group.
	OrganizationTypeCg OrganizationType = "cg"
	// An organization that is a registered business or corporation but not identified by other types.
	OrganizationTypeBus OrganizationType = "bus"
	// Other type of organization not already specified.
	OrganizationTypeOther OrganizationType = "other"
)
