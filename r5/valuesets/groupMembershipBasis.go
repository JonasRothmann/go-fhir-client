// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/group-membership-basis
*/type GroupMembershipBasis string

const (
	// The Group.characteristics specified are both necessary and sufficient to determine membership. All entities that meet the criteria are considered to be members of the group, whether referenced by the group or not. If members are present, they are individuals that happen to be known as meeting the Group.characteristics. The list cannot be presumed to be complete.
	GroupMembershipBasisDefinitional GroupMembershipBasis = "definitional"
	// The Group.characteristics are necessary but not sufficient to determine membership. Membership is determined by being listed as one of the Group.member.
	GroupMembershipBasisEnumerated GroupMembershipBasis = "enumerated"
)
