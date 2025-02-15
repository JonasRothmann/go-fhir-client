// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/MarketingStatus
type MarketingStatus struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The country in which the marketing authorization has been granted shall be specified It should be specified using the ISO 3166 ‑ 1 alpha-2 code elements
	Country *CodeableConcept `bson:"country,omitempty" json:"country,omitempty"`
	// Where a Medicines Regulatory Agency has granted a marketing authorization for which specific provisions within a jurisdiction apply, the jurisdiction can be specified using an appropriate controlled terminology The controlled term and the controlled term identifier shall be specified
	Jurisdiction *CodeableConcept `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	// This attribute provides information on the status of the marketing of the medicinal product See ISO/TS 20443 for more information and examples
	Status CodeableConcept `bson:"status" json:"status"`
	// The date when the Medicinal Product is placed on the market by the Marketing Authorization Holder (or where applicable, the manufacturer/distributor) in a country and/or jurisdiction shall be provided A complete date consisting of day, month and year shall be specified using the ISO 8601 date format NOTE “Placed on the market” refers to the release of the Medicinal Product into the distribution chain
	DateRange *Period `bson:"dateRange,omitempty" json:"dateRange,omitempty"`
	// The date when the Medicinal Product is placed on the market by the Marketing Authorization Holder (or where applicable, the manufacturer/distributor) in a country and/or jurisdiction shall be provided A complete date consisting of day, month and year shall be specified using the ISO 8601 date format NOTE “Placed on the market” refers to the release of the Medicinal Product into the distribution chain
	RestoreDate *custom.DateTime `bson:"restoreDate,omitempty" json:"restoreDate,omitempty"`
}

func (m MarketingStatus) ResourceType() string {
	return "StructureDefinition"
}
