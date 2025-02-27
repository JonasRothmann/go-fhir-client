// GENERATED CODE – DO NOT EDIT!

package r5

import "encoding/json"

// http://hl7.org/fhir/StructureDefinition/ContactDetail
type ContactDetail struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Name of an individual to contact
	Name *string `json:"name,omitempty"`
	// Contact details for individual or organization
	Telecom []ContactPoint `json:"telecom,omitempty"`
}

type OtherContactDetail ContactDetail

func (c ContactDetail) ResourceType() string {
	return "ContactDetail"
}

func (c ContactDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherContactDetail
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherContactDetail: OtherContactDetail(c), ResourceType: c.ResourceType()})
}

func UnmarshalContactDetail(b []byte) (res ContactDetail, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
