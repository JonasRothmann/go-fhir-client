// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/certainty-type
*/type EvidenceCertaintyType string

const (
	// Overall certainty of evidence (quality of evidence).
	EvidenceCertaintyTypeOverall EvidenceCertaintyType = "Overall"
	// methodologic concerns reducing internal validity.
	EvidenceCertaintyTypeRiskOfBias EvidenceCertaintyType = "RiskOfBias"
	// concerns that findings are not similar enough to support certainty.
	EvidenceCertaintyTypeInconsistency EvidenceCertaintyType = "Inconsistency"
	// concerns reducing external validity.
	EvidenceCertaintyTypeIndirectness EvidenceCertaintyType = "Indirectness"
	// fuzzy or wide variability.
	EvidenceCertaintyTypeImprecision EvidenceCertaintyType = "Imprecision"
	// likelihood that what is published misrepresents what is available to publish.
	EvidenceCertaintyTypePublicationBias EvidenceCertaintyType = "PublicationBias"
	// higher certainty due to dose response relationship.
	EvidenceCertaintyTypeDoseResponseGradient EvidenceCertaintyType = "DoseResponseGradient"
	// higher certainty due to risk of bias in opposite direction.
	EvidenceCertaintyTypePlausibleConfounding EvidenceCertaintyType = "PlausibleConfounding"
	// higher certainty due to large effect size.
	EvidenceCertaintyTypeLargeEffect EvidenceCertaintyType = "LargeEffect"
)
