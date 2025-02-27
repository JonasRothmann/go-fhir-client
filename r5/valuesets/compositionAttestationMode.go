// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/composition-attestation-mode
*/type CompositionAttestationMode string

const (
	// The person authenticated the content in their personal capacity.
	CompositionAttestationModePersonal CompositionAttestationMode = "personal"
	// The person authenticated the content in their professional capacity.
	CompositionAttestationModeProfessional CompositionAttestationMode = "professional"
	// The person authenticated the content and accepted legal responsibility for its content.
	CompositionAttestationModeLegal CompositionAttestationMode = "legal"
	// The organization authenticated the content as consistent with their policies and procedures.
	CompositionAttestationModeOfficial CompositionAttestationMode = "official"
)
