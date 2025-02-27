// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/ContactPoint
type ContactPoint struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// phone | fax | email | pager | url | sms | other
	System *custom.Code `json:"system,omitempty"`
	// The actual contact point details
	Value *string `json:"value,omitempty"`
	// home | work | temp | old | mobile - purpose of this contact point
	Use *custom.Code `json:"use,omitempty"`
	// Specify preferred order of use (1 = highest)
	Rank *uint32 `json:"rank,omitempty"`
	// Time period when the contact point was/is in use
	Period *Period `json:"period,omitempty"`
}

type OtherContactPoint ContactPoint

func (c ContactPoint) ResourceType() string {
	return "ContactPoint"
}

func (c ContactPoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherContactPoint
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherContactPoint: OtherContactPoint(c), ResourceType: c.ResourceType()})
}

func UnmarshalContactPoint(b []byte) (res ContactPoint, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
