// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/ExtendedContactDetail
type ExtendedContactDetail struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// The type of contact
	Purpose *CodeableConcept `json:"purpose,omitempty"`
	// Name of an individual to contact
	Name []HumanName `json:"name,omitempty"`
	// Contact details (e.g.phone/fax/url)
	Telecom []ContactPoint `json:"telecom,omitempty"`
	// Address for the contact
	Address *Address `json:"address,omitempty"`
	// This contact detail is handled/monitored by a specific organization
	Organization *custom.Reference[Organization] `json:"organization,omitempty"`
	// Period that this contact was valid for usage
	Period *Period `json:"period,omitempty"`
}

type OtherExtendedContactDetail ExtendedContactDetail

func (e ExtendedContactDetail) ResourceType() string {
	return "ExtendedContactDetail"
}

func (e ExtendedContactDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherExtendedContactDetail
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherExtendedContactDetail: OtherExtendedContactDetail(e), ResourceType: e.ResourceType()})
}

func UnmarshalExtendedContactDetail(b []byte) (res ExtendedContactDetail, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
