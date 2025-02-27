// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://dicom.nema.org/resources/ontology/DCM
- http://terminology.hl7.org/CodeSystem/extra-security-role-type
*/type ParticipationRoleType string

const (
	// An entity providing authorization services to enable the electronic sharing of health-related information based on resource owner's preapproved permissions. For example, an UMA Authorization Server[UMA]
	ParticipationRoleTypeAuthserver ParticipationRoleType = "authserver"
	// An entity that collects information over which the data subject may have certain rights under policy or law to control that information's management and distribution by data collectors, including the right to access, retrieve, distribute, or delete that information.
	ParticipationRoleTypeDatacollector ParticipationRoleType = "datacollector"
	// An entity that processes collected information over which the data subject may have certain rights under policy or law to control that information's management and distribution by data processors, including the right to access, retrieve, distribute, or delete that information.
	ParticipationRoleTypeDataprocessor ParticipationRoleType = "dataprocessor"
	// A person whose personal information is collected or processed, and who may have certain rights under policy or law to control that information's management and distribution by data collectors or processors, including the right to access, retrieve, distribute, or delete that information.
	ParticipationRoleTypeDatasubject ParticipationRoleType = "datasubject"
	// The human user that has participated.
	ParticipationRoleTypeHumanuser ParticipationRoleType = "humanuser"
)
