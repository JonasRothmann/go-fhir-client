// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/MarketingStatus
type MarketingStatus struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The country in which the marketing authorization has been granted shall be specified It should be specified using the ISO 3166 ‑ 1 alpha-2 code elements
	Country *CodeableConcept `json:"country,omitempty"`
	// Where a Medicines Regulatory Agency has granted a marketing authorization for which specific provisions within a jurisdiction apply, the jurisdiction can be specified using an appropriate controlled terminology The controlled term and the controlled term identifier shall be specified
	Jurisdiction *CodeableConcept `json:"jurisdiction,omitempty"`
	// This attribute provides information on the status of the marketing of the medicinal product See ISO/TS 20443 for more information and examples
	Status CodeableConcept `json:"status"`
	// The date when the Medicinal Product is placed on the market by the Marketing Authorization Holder (or where applicable, the manufacturer/distributor) in a country and/or jurisdiction shall be provided A complete date consisting of day, month and year shall be specified using the ISO 8601 date format NOTE “Placed on the market” refers to the release of the Medicinal Product into the distribution chain
	DateRange *Period `json:"dateRange,omitempty"`
	// The date when the Medicinal Product is placed on the market by the Marketing Authorization Holder (or where applicable, the manufacturer/distributor) in a country and/or jurisdiction shall be provided A complete date consisting of day, month and year shall be specified using the ISO 8601 date format NOTE “Placed on the market” refers to the release of the Medicinal Product into the distribution chain
	RestoreDate *custom.DateTime `json:"restoreDate,omitempty"`
}

type OtherMarketingStatus MarketingStatus

func (m MarketingStatus) ResourceType() string {
	return "MarketingStatus"
}

func (m MarketingStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMarketingStatus
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherMarketingStatus: OtherMarketingStatus(m), ResourceType: m.ResourceType()})
}

func UnmarshalMarketingStatus(b []byte) (res MarketingStatus, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
