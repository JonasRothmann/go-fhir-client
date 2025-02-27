// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/specimen-role
*/type SpecimenRole string

const (
	// Used to test the validity of the measurement process, where the composition of the sample is unknown except to the person submitting it.
	SpecimenRoleB SpecimenRole = "b"
	// Used for initial setting of calibration of the instrument.
	SpecimenRoleC SpecimenRole = "c"
	// Used with manufactured reference providing signals that simulate QC results.
	SpecimenRoleE SpecimenRole = "e"
	// Specimen used for testing proficiency of the organization performing the testing (Filler).
	SpecimenRoleF SpecimenRole = "f"
	// Specimen used for testing Operator Proficiency.
	SpecimenRoleO SpecimenRole = "o"
	// Used for any patient sample.
	SpecimenRoleP SpecimenRole = "p"
	// Used when specimen is the control specimen (either positive or negative).
	SpecimenRoleQ SpecimenRole = "q"
	// Used when a patient sample is re-run as a control for a repeat test.
	SpecimenRoleR SpecimenRole = "r"
	// Used for periodic calibration checks.
	SpecimenRoleV SpecimenRole = "v"
)
