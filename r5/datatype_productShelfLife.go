// GENERATED CODE – DO NOT EDIT!

package r5

import "encoding/json"

// http://hl7.org/fhir/StructureDefinition/ProductShelfLife
type ProductShelfLife struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// This describes the shelf life, taking into account various scenarios such as shelf life of the packaged Medicinal Product itself, shelf life after transformation where necessary and shelf life after the first opening of a bottle, etc. The shelf life type shall be specified using an appropriate controlled vocabulary The controlled term and the controlled term identifier shall be specified
	Type *CodeableConcept `json:"type,omitempty"`
	// The shelf life time period can be specified using a numerical value for the period of time and its unit of time measurement The unit of measurement shall be specified in accordance with ISO 11240 and the resulting terminology The symbol and the symbol identifier shall be used
	PeriodDuration *Duration `json:"periodDuration,omitempty"`
	// The shelf life time period can be specified using a numerical value for the period of time and its unit of time measurement The unit of measurement shall be specified in accordance with ISO 11240 and the resulting terminology The symbol and the symbol identifier shall be used
	PeriodString *string `json:"periodString,omitempty"`
	// Special precautions for storage, if any, can be specified using an appropriate controlled vocabulary The controlled term and the controlled term identifier shall be specified
	SpecialPrecautionsForStorage []CodeableConcept `json:"specialPrecautionsForStorage,omitempty"`
}

type OtherProductShelfLife ProductShelfLife

func (p ProductShelfLife) ResourceType() string {
	return "ProductShelfLife"
}

func (p ProductShelfLife) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherProductShelfLife
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherProductShelfLife: OtherProductShelfLife(p), ResourceType: p.ResourceType()})
}

func UnmarshalProductShelfLife(b []byte) (res ProductShelfLife, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
