// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/clinical-use-definition-type
*/type ClinicalUseDefinitionType string

const (
	// A reason for giving the medication.
	ClinicalUseDefinitionTypeIndication ClinicalUseDefinitionType = "indication"
	// A reason for not giving the medication.
	ClinicalUseDefinitionTypeContraindication ClinicalUseDefinitionType = "contraindication"
	// Interactions between the medication and other substances.
	ClinicalUseDefinitionTypeInteraction ClinicalUseDefinitionType = "interaction"
	// Side effects or adverse effects associated with the medication.
	ClinicalUseDefinitionTypeUndesirableEffect ClinicalUseDefinitionType = "undesirable-effect"
	// A general warning or issue that is not specifically one of the other types.
	ClinicalUseDefinitionTypeWarning ClinicalUseDefinitionType = "warning"
)
